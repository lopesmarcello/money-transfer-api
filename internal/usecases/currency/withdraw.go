package currency

import (
	"context"

	"github.com/lopesmarcello/money-transfer/internal/utils/validator"
)

type WithdrawReq struct {
	IsPessoaFisica bool    `json:"is_pessoa_fisica"`
	Amount         float64 `json:"amount"`
}

func (req WithdrawReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.AssertBool(req.IsPessoaFisica), "is_pessoa_fisica", "must be true or false")
	eval.CheckField(req.Amount > 0, "amount", "amount should be positive")

	return eval
}
