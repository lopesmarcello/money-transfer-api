package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/lopesmarcello/money-transfer/internal/usecases/user"
	"github.com/lopesmarcello/money-transfer/internal/utils"
)

func (api *API) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJSON[user.CreateUserReq](r)
	slog.Info(fmt.Sprintf("%d problems", len(problems)))

	if len(problems) > 0 {
		for problem := range problems {
			fmt.Println("Problem:")
			fmt.Println(problem)
		}
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	if err != nil {
		slog.Error("Error decoding JSON:", err)
		utils.EncodeJSON(w, r, http.StatusInternalServerError,
			utils.JSONmsg("error", "internal server error", "message", "error decoding json"))
		return
	}

	// data.TipoPessoa

	id, err := api.UserService.CreateUserPessoaFisica(r.Context(),
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

	utils.EncodeJSON(w, r, http.StatusCreated,
		utils.JSONmsg("sucess", "user created", "user_id", string(id)))
}
