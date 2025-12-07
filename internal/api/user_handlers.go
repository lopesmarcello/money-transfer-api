package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/lopesmarcello/money-transfer/internal/usecases/user"
	"github.com/lopesmarcello/money-transfer/internal/utils"
)

// Deletes user
func (api *API) handleCloseAccount(w http.ResponseWriter, r *http.Request) {
	stringID := chi.URLParam(r, "id")
	userID64, err := strconv.Atoi(stringID)
	if err != nil {
		utils.EncodeJSON(w, r, http.StatusBadRequest, utils.JSONmsg("error", "error parsing ID", "user_id", stringID))
		return
	}
	userID := int32(userID64)

	data, problems, err := utils.DecodeValidJSON[user.DeleteUserReq](r)
	if len(problems) > 0 {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
	}
	if err != nil {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, err)
	}

	api.UserService.DeleteUser(r.Context(), data.IsPessoaFisica, userID)
}

// Creates new user
func (api *API) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJSON[user.CreateUserReq](r)

	if len(problems) > 0 {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	if err != nil {
		utils.EncodeJSON(w, r, http.StatusInternalServerError,
			utils.JSONmsg("error", "internal server error", "message", "error decoding json"))
		return
	}

	var id int32
	isPessoaFisica := data.TipoPessoa == 0

	if isPessoaFisica {
		id, err = api.UserService.CreateUserPessoaFisica(r.Context(),
			data.RendaMensal,
			int32(data.Idade),
			data.NomeCompleto,
			data.Email,
			data.Celular,
			data.Categoria,
		)
		if err != nil {
			utils.EncodeJSON(w, r, http.StatusUnprocessableEntity,
				utils.JSONmsg("error", "unprocessable entity", "message", "error while creating pessoa fisica"))
			return
		}
	} else {
		id, err = api.UserService.CreateUserPessoaJuridica(r.Context(),
			data.Faturamento,
			data.NomeFantasia,
			data.Email,
			data.Celular,
			data.Categoria,
			data.Saldo)
		if err != nil {
			utils.EncodeJSON(w, r, http.StatusUnprocessableEntity,
				utils.JSONmsg("error", "unprocessable entity", "message", "error while creating pessoa juridica"))
			return
		}
	}

	utils.EncodeJSON(w, r, http.StatusCreated, utils.JSONmsg("sucess", "user created", "user_id", strconv.Itoa(int(id))))
}

func (api *API) getCsrfToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	log.Printf("Generated CSRF token: %s", token) // Log to console
	if token == "" {
		log.Println("Warning: Empty CSRF token - middleware may not have run")
	}
	// Return the token as JSON for easy consumption by clients
	w.Header().Set("Content-Type", "application/json")
	utils.EncodeJSON(w, r, http.StatusOK,
		utils.JSONmsg("csrf_token", token))

	// Optionally, set it as a header for SPAs (e.g., X-CSRF-Token)
	w.Header().Set("X-CSRF-Token", token)
}
