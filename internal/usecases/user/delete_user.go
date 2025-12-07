package user

import (
	"context"

	"github.com/lopesmarcello/money-transfer/internal/utils/validator"
)

type DeleteUserReq struct {
	IsPessoaFisica bool `json:"is_pessoa_fisica"`
}

func (req DeleteUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.AssertBool(req.IsPessoaFisica), "is_pessoa_fisica", "must be true or false")

	return eval
}
