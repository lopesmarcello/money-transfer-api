package services

import (
	"context"
	"errors"

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

func (cs *CurrencyService) Deposit(ctx context.Context, isPessoaFisica bool, destinationID int32, amount float64) (float64, error) {
	var updatedAmount float64

	if isPessoaFisica {
		currentAmount, err := cs.queries.GetSaldoPessoaFisicaByID(ctx, destinationID)
		if err != nil {
			return 0, err
		}
		args := pgstore.UpdateSaldoFromPessoaFisicaParams{
			ID:    destinationID,
			Saldo: currentAmount + amount,
		}
		updatedPerson, err := cs.queries.UpdateSaldoFromPessoaFisica(ctx, args)
		if err != nil {
			return 0, err
		}
		updatedAmount = updatedPerson.Saldo
	}

	if !isPessoaFisica {
		currentAmount, err := cs.queries.GetSaldoPessoaJuridicaByID(ctx, destinationID)
		if err != nil {
			return 0, err
		}
		args := pgstore.UpdateSaldoFromPessoaJuridicaParams{
			ID:    destinationID,
			Saldo: currentAmount + amount,
		}
		updatedPerson, err := cs.queries.UpdateSaldoFromPessoaJuridica(ctx, args)
		if err != nil {
			return 0, err
		}
		updatedAmount = updatedPerson.Saldo
	}

	return updatedAmount, nil
}

var ErrNotEnoughCurrency = errors.New("not enough currency")

func (cs *CurrencyService) Withdraw(ctx context.Context, isPessoaFisica bool, originID int32, amount float64) (float64, error) {
	var updatedAmount float64

	if isPessoaFisica {
		currentAmount, err := cs.queries.GetSaldoPessoaFisicaByID(ctx, originID)
		if err != nil {
			return 0, err
		}

		updatedAmount = currentAmount - amount

		if updatedAmount < 0 {
			return 0, ErrNotEnoughCurrency
		}
		args := pgstore.UpdateSaldoFromPessoaFisicaParams{
			ID:    originID,
			Saldo: updatedAmount,
		}
		updatedPerson, err := cs.queries.UpdateSaldoFromPessoaFisica(ctx, args)
		if err != nil {
			return 0, err
		}
		updatedAmount = updatedPerson.Saldo
	}

	if !isPessoaFisica {
		currentAmount, err := cs.queries.GetSaldoPessoaJuridicaByID(ctx, originID)
		if err != nil {
			return 0, err
		}
		updatedAmount = currentAmount - amount
		if updatedAmount < 0 {
			return 0, ErrNotEnoughCurrency
		}

		args := pgstore.UpdateSaldoFromPessoaJuridicaParams{
			ID:    originID,
			Saldo: updatedAmount,
		}

		updatedPerson, err := cs.queries.UpdateSaldoFromPessoaJuridica(ctx, args)
		if err != nil {
			return 0, err
		}

		updatedAmount = updatedPerson.Saldo
	}

	return updatedAmount, nil
}

func (cs *CurrencyService) Transfer(ctx context.Context, isDestinationPessoaFisica, isOriginPessoaFisica bool, destinationID, originID int32, amount float64) (map[string]any, error) {
	result := make(map[string]any)

	var originCurrency float64
	var destinationCurrency float64

	// ORIGIN
	if isOriginPessoaFisica {
		currentAmount, err := cs.queries.GetSaldoPessoaFisicaByID(ctx, originID)
		if err != nil {
			return result, err
		}
		originCurrency = currentAmount
	} else {
		currentAmount, err := cs.queries.GetSaldoPessoaJuridicaByID(ctx, originID)
		if err != nil {
			return result, err
		}
		originCurrency = currentAmount
	}

	if originCurrency-amount < 0 {
		return result, ErrNotEnoughCurrency
	}

	// DESTINATION
	if isDestinationPessoaFisica {
		currentAmount, err := cs.queries.GetSaldoPessoaFisicaByID(ctx, originID)
		if err != nil {
			return result, err
		}
		destinationCurrency = currentAmount
	} else {
		currentAmount, err := cs.queries.GetSaldoPessoaJuridicaByID(ctx, originID)
		if err != nil {
			return result, err
		}
		destinationCurrency = currentAmount
	}

	if isOriginPessoaFisica {
		argsOriginPF := pgstore.UpdateSaldoFromPessoaFisicaParams{
			Saldo: originCurrency - amount,
			ID:    originID,
		}
		updatedOrigin, err := cs.queries.UpdateSaldoFromPessoaFisica(ctx, argsOriginPF)
		if err != nil {
			return result, err
		}

		result["origin"] = updatedOrigin
	}

	if !isOriginPessoaFisica {
		argsOriginPJ := pgstore.UpdateSaldoFromPessoaJuridicaParams{
			Saldo: originCurrency - amount,
			ID:    originID,
		}

		updatedOrigin, err := cs.queries.UpdateSaldoFromPessoaJuridica(ctx, argsOriginPJ)
		if err != nil {
			return result, err
		}

		result["origin"] = updatedOrigin
	}

	if isDestinationPessoaFisica {
		argsDestinationPF := pgstore.UpdateSaldoFromPessoaFisicaParams{
			Saldo: destinationCurrency + amount,
			ID:    destinationID,
		}

		updatedDestination, err := cs.queries.UpdateSaldoFromPessoaFisica(ctx, argsDestinationPF)
		if err != nil {
			return result, err
		}

		result["destination"] = updatedDestination
	}

	if !isDestinationPessoaFisica {
		argsDestinationPJ := pgstore.UpdateSaldoFromPessoaJuridicaParams{
			Saldo: destinationCurrency + amount,
			ID:    destinationID,
		}

		updatedDestination, err := cs.queries.UpdateSaldoFromPessoaJuridica(ctx, argsDestinationPJ)
		if err != nil {
			return result, err
		}

		result["destination"] = updatedDestination
	}

	return result, nil
}
