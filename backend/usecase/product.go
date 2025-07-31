package usecase

import (
	"context"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
	"github.com/adriel-timoteo/olist-fullstack/backend/repository"
)

type ProductUsecaseItf interface {
	GetTopCategories(context.Context, int) ([]entity.ProductCategoryCount, error)
}

type ProductUsecaseImpl struct {
	pr  repository.ProductRepoItf
	trx repository.Transactor
}

func NewProductUsecaseImpl(pr repository.ProductRepoItf, trx repository.Transactor) ProductUsecaseImpl {
	return ProductUsecaseImpl{
		pr:  pr,
		trx: trx,
	}
}

func (puc ProductUsecaseImpl) GetTopCategories(ctx context.Context, limit int) ([]entity.ProductCategoryCount, error) {
	data, err := puc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return puc.pr.SelectTopCategories(ctx, limit)
	})
	if err != nil {
		return nil, err
	}

	categories, ok := data.([]entity.ProductCategoryCount)
	if !ok {
		return nil, ce.NewError(ce.InternalError, "error occurred")
	}

	return categories, nil
}
