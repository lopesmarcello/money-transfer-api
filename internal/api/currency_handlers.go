package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lopesmarcello/money-transfer/internal/usecases/currency"
	"github.com/lopesmarcello/money-transfer/internal/utils"
)

func (api *API) handleGetSaldo(w http.ResponseWriter, r *http.Request) {
	var saldo float64

	data, problems, err := utils.DecodeValidJSON[currency.GetSaldoReq](r)
	if len(problems) > 0 {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	if err != nil {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity,
			utils.JSONmsg("error", "internal server error", "message", "payload error"))
		return
	}

	stringUserID := chi.URLParam(r, "id")
	userID64, _ := strconv.ParseInt(stringUserID, 10, 32)
	userID := int32(userID64)

	if data.IsPessoaFisica {
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
		utils.JSONmsg("success", "user found with success", "userID", stringUserID, "saldo", saldo))
}

func (api *API) handleDeposit(w http.ResponseWriter, r *http.Request) {
	stringUserID := chi.URLParam(r, "id")
	userID64, _ := strconv.ParseInt(stringUserID, 10, 32)
	userID := int32(userID64)

	data, problems, err := utils.DecodeValidJSON[currency.DepositReq](r)

	if len(problems) > 0 {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}
	if err != nil {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	updatedAmount, err := api.CurrencyService.Deposit(r.Context(), data.IsPessoaFisica, userID, data.Amount)
	if err != nil {
		utils.EncodeJSON(w, r, http.StatusBadRequest, utils.JSONmsg("error", "bad request", "errors", err))
		return
	}

	utils.EncodeJSON(w, r, http.StatusOK, utils.JSONmsg("success", "your deposit has been done sucessfully", "updated_amount", updatedAmount, "destination_id", userID))
}

func (api *API) handleWithdraw(w http.ResponseWriter, r *http.Request) {
	stringUserID := chi.URLParam(r, "id")
	userID64, _ := strconv.ParseInt(stringUserID, 10, 32)
	userID := int32(userID64)

	data, problems, err := utils.DecodeValidJSON[currency.WithdrawReq](r)

	if len(problems) > 0 {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}
	if err != nil {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	updatedAmount, err := api.CurrencyService.Withdraw(r.Context(), data.IsPessoaFisica, userID, data.Amount)
	if err != nil {
		utils.EncodeJSON(w, r, http.StatusBadRequest, utils.JSONmsg("error", "bad request", "errors", err))
		return
	}

	utils.EncodeJSON(w, r, http.StatusOK, utils.JSONmsg("success", "your withdraw has been done sucessfully", "updated_amount", updatedAmount, "destination_id", userID))
}

func (api *API) handleTransfer(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJSON[currency.TransferReq](r)
	if len(problems) > 0 {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}
	if err != nil {
		utils.EncodeJSON(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	result, err := api.CurrencyService.Transfer(
		r.Context(),
		data.IsDestinationPessoaFisica,
		data.IsOriginPessoaFisica,
		data.DestinationID,
		data.OriginID,
		data.Amount,
	)
	if err != nil {
		utils.EncodeJSON(w, r, http.StatusInternalServerError, err)
		return
	}

	utils.EncodeJSON(w, r, http.StatusOK, result)
}
