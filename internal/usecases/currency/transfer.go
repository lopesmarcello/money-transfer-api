package currency

import (
	"context"

	"github.com/lopesmarcello/money-transfer/internal/utils/validator"
)

type TransferReq struct {
	IsOriginPessoaFisica      bool    `json:"is_origin_pessoa_fisica"`
	OriginID                  int32   `json:"origin_id"`
	IsDestinationPessoaFisica bool    `json:"is_destination_pessoa_fisica"`
	DestinationID             int32   `json:"destination_id"`
	Amount                    float64 `json:"amount"`
}

func (req TransferReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator
	eval.CheckField(validator.AssertBool(req.IsOriginPessoaFisica), "origin_tipo_pessoa", "must be 0 or 1")
	eval.CheckField(req.OriginID != 0, "origin_id", "invalid origin id")
	eval.CheckField(validator.AssertBool(req.IsDestinationPessoaFisica), "origin_tipo_pessoa", "must be 0 or 1")
	eval.CheckField(req.DestinationID != 0, "origin_id", "invalid origin id")
	eval.CheckField(req.Amount > 0, "amount", "amount should be positive")

	return eval
}
