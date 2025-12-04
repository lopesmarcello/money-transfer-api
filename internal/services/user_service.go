package services

import (
	"context"

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

func (us *UserService) CreateUserPessoaFisica(ctx context.Context, rendaMensal float64, idade int32, nomeCompleto, email, celular, categoria string) (int32, error) {
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
			return uuid.UUID{}, ErrDuplicatedEmailOrUsername
		}
		return uuid.UUID{}, err
	}

	return id, nil
}
