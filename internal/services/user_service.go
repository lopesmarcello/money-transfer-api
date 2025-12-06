// Package services
package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lopesmarcello/money-transfer/internal/store/pgstore"
)

type UserService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (us *UserService) CreateUserPessoaFisica(
	ctx context.Context,
	rendaMensal float64,
	idade int32,
	nomeCompleto,
	email,
	celular,
	categoria string,
) (int32, error) {
	args := pgstore.CreateUserPessoaFisicaParams{
		RendaMensal:  rendaMensal,
		Saldo:        0,
		NomeCompleto: nomeCompleto,
		Email:        email,
		Idade:        idade,
		Celular:      celular,
		Categoria:    categoria,
	}

	id, err := us.queries.CreateUserPessoaFisica(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, errors.New("duplicated email")
		}
		return 0, err
	}

	return id, nil
}

func (us *UserService) CreateUserPessoaJuridica(ctx context.Context,
	faturamento float64,
	nomeFantasia string,
	email string,
	celular string,
	categoria string,
	saldo float64,
) (int32, error) {
	args := pgstore.CreateUserPessoaJuridicaParams{
		Faturamento:      faturamento,
		NomeFantasia:     nomeFantasia,
		EmailCorporativo: email,
		Celular:          celular,
		Categoria:        categoria,
		Saldo:            saldo,
	}

	id, err := us.queries.CreateUserPessoaJuridica(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, errors.New("duplicated email")
		}
		return 0, err

	}

	return id, nil
}
