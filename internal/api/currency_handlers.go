package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lopesmarcello/money-transfer/internal/usecases/currency"
	"github.com/lopesmarcello/money-transfer/internal/utils"
)

func (api *API) handleGetSaldo(w http.ResponseWriter, r *http.Request) {
	var saldo float64

	data, problems, err := utils.DecodeValidJSON[currency.GetSaldoReq](r)
	slog.Info(fmt.Sprintf("%d problems", len(problems)))
	if len(problems) > 0 {
		for problem := range problems {
			fmt.Println("Problem")
			fmt.Println(problem)
		}
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	if err != nil {
		slog.Error("Error decoding JSON:", err)
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity,
			utils.JSONmsg("error", "internal server error", "message", "payload error"))
		return
	}

	userIdString := chi.URLParam(r, "id")
	userID64, _ := strconv.ParseInt(userIdString, 10, 32)
	userID := int32(userID64)

	isPessoaFisica := data.TipoPessoa == 0

	if isPessoaFisica {
		saldo, err = api.CurrencyService.GetSaldoPessoaFisica(r.Context(), userID)
		if err != nil {
			utils.EncodeJSON(w, r, http.StatusUnprocessableEntity,
				utils.JSONmsg("error", "unprocessable entity", "message", "error while getting saldo pessoa fisica"))
			return
		}
	} else {
		saldo, err = api.CurrencyService.GetSaldoPessoaJuridica(r.Context(), userID)
		if err != nil {
			utils.EncodeJSON(w, r, http.StatusUnprocessableEntity,
				utils.JSONmsg("error", "unprocessable entity", "message", "error while getting saldo pessoa juridica"))
			return
		}
	}

	utils.EncodeJSON(w, r, http.StatusOK,
		utils.JSONmsg("success", "user found with success", "userID", userIdString, "saldo", saldo))
}
