package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lopesmarcello/money-transfer/internal/store/pgstore"
)

type CurrencyService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewCurrencyService(pool *pgxpool.Pool) CurrencyService {
	return CurrencyService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (cs *CurrencyService) GetSaldoPessoaFisica(ctx context.Context, userID int32) (float64, error) {
	saldo, err := cs.queries.GetSaldoPessoaFisicaByID(ctx, userID)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}

func (cs *CurrencyService) GetSaldoPessoaJuridica(ctx context.Context, userID int32) (float64, error) {
	saldo, err := cs.queries.GetSaldoPessoaJuridicaByID(ctx, userID)
	if err != nil {
		return 0, err
	}
	return saldo, nil
}
