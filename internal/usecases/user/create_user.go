// Package user to control user related requisitions
package user

import (
	"context"
	"log/slog"

	"github.com/lopesmarcello/money-transfer/internal/utils/validator"
)

type Pessoa int

const (
	PessoaFisica Pessoa = iota
	PessoaJuridica
)

type CreateUserReq struct {
	TipoPessoa Pessoa `json:"tipo_pessoa"`

	Idade     int     `json:"idade"`
	Celular   string  `json:"celular"`
	Categoria string  `json:"categoria"`
	Email     string  `json:"email"`
	Saldo     float64 `json:"saldo"`

	// Fisica
	RendaMensal  float64 `json:"renda_mensal"`
	NomeCompleto string  `json:"nome_completo"`

	// Juridica
	Faturamento  float64 `json:"faturamento"`
	NomeFantasia string  `json:"nome_fantasia"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {
	var (
		eval           validator.Evaluator
		isPessoaFisica = req.TipoPessoa == 0
	)

	slog.Info("isPessoaFisica", isPessoaFisica)
	// Both
	eval.CheckField(req.Idade > 0, "idade", "this field cannot be negative")

	eval.CheckField(validator.NotBlank(req.Celular), "celular", "this field cannot be empty")
	eval.CheckField(validator.MinChars(req.Celular, 8) && validator.MaxChars(req.Celular, 20), "celular", "this field needs to have  a length between 8 and 20")

	eval.CheckField(validator.NotBlank(req.Email), "email", "this field cannot be empty")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "this field must be a valid email")
	eval.CheckField(validator.MinChars(req.Email, 8) && validator.MaxChars(req.Email, 255), "email", "this field needs to have  a length between 8 and 255")

	eval.CheckField(validator.NotBlank(req.Categoria), "categoria", "this field cannot be empty")
	eval.CheckField(validator.MinChars(req.Email, 0) && validator.MaxChars(req.Email, 50), "categoria", "this field needs to have a length limit of 50")

	// Fisica
	if isPessoaFisica {
		eval.CheckField(req.RendaMensal >= 0, "renda_mensal", "this field cannot be negative")

		eval.CheckField(validator.NotBlank(req.NomeCompleto), "nome_completo", "this field cannot be empty")
		eval.CheckField(validator.MinChars(req.NomeCompleto, 4) && validator.MaxChars(req.NomeCompleto, 255), "nome_completo", "this field needs to have  a length between 4 and 255")
	}

	// Juridica
	if !isPessoaFisica {
		eval.CheckField(req.Faturamento >= 0, "faturamento", "this field cannot be negative")

		eval.CheckField(validator.NotBlank(req.NomeFantasia), "nome_fantasia", "this field cannot be empty")
		eval.CheckField(validator.MinChars(req.NomeFantasia, 4) && validator.MaxChars(req.NomeFantasia, 255), "nome_fantasia", "this field needs to have  a length between 4 and 255")
	}

	return eval
}
