package usecase

import (
	"context"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
	"github.com/adriel-timoteo/olist-fullstack/backend/repository"
)

type CustomerUsecaseItf interface {
	GetTopCities(context.Context, int) ([]entity.CustomerCityCount, error)
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
		return nil, ce.NewError(ce.CommonErr, "error occurred")
	}

	return cities, nil
}
