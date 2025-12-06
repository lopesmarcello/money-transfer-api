package currency

import (
	"context"

	"github.com/lopesmarcello/money-transfer/internal/usecases/user"
	"github.com/lopesmarcello/money-transfer/internal/utils/validator"
)

type GetSaldoReq struct {
	TipoPessoa user.Pessoa `json:"tipo_pessoa"`
}

func (req GetSaldoReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(req.TipoPessoa != 0 || req.TipoPessoa != 1, "tipo_pessoa", "must be 0 or 1")

	return eval
}
