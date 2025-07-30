package usecase

import (
	"context"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
	"github.com/adriel-timoteo/olist-fullstack/backend/repository"
)

type CustomerUsecaseItf interface {
	GetTopCities(context.Context, int) ([]entity.CustomerCityCount, error)
	GetTotalUniqueCustomers(context.Context) (*entity.Count, error)
	GetRepeatPurchaseRate(context.Context) (*entity.Rate, error)
}

type CustomerUsecaseImpl struct {
	cr  repository.CustomerRepoItf
	trx repository.Transactor
}

func NewCustomerUsecaseImpl(cr repository.CustomerRepoItf, trx repository.Transactor) CustomerUsecaseImpl {
	return CustomerUsecaseImpl{
		cr:  cr,
		trx: trx,
	}
}

func (cuc CustomerUsecaseImpl) GetTopCities(ctx context.Context, limit int) ([]entity.CustomerCityCount, error) {
	data, err := cuc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return cuc.cr.SelectTopCities(ctx, limit)
	})
	if err != nil {
		return nil, err
	}

	cities, ok := data.([]entity.CustomerCityCount)
	if !ok {
		return nil, ce.NewError(ce.InternalError, "error occurred")
	}

	return cities, nil
}

func (cuc CustomerUsecaseImpl) GetTotalUniqueCustomers(ctx context.Context) (*entity.Count, error) {
	data, err := cuc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return cuc.cr.SelectTotalUniqueCustomer(ctx)
	})
	if err != nil {
		return nil, err
	}

	count, ok := data.(*entity.Count)
	if !ok {
		return nil, ce.NewError(ce.InternalError, "error occurred")
	}

	return count, nil
}

func (cuc CustomerUsecaseImpl) GetRepeatPurchaseRate(ctx context.Context) (*entity.Rate, error) {
	data, err := cuc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return cuc.cr.SelectRepeatPurchaseRate(ctx)
	})
	if err != nil {
		return nil, err
	}

	rate, ok := data.(*entity.Rate)
	if !ok {
		return nil, ce.NewError(ce.InternalError, "error occurred")
	}

	return rate, nil
}
