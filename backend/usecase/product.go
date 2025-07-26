package usecase

import (
	"context"
	"time"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/constant"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
	"github.com/adriel-timoteo/olist-fullstack/backend/repository"
)

type ProductUsecaseItf interface {
	GetDeliveredTrend(context.Context, time.Time, time.Time, constant.Interval) ([]entity.DeliveredProductTrend, error)
	GetProductStatusTrend(context.Context, time.Time, time.Time) ([]entity.ProductStatusTrend, error)
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

func (puc ProductUsecaseImpl) GetDeliveredTrend(ctx context.Context, start, end time.Time, interval constant.Interval) ([]entity.DeliveredProductTrend, error) {
	data, err := puc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return puc.pr.SelectDeliveredProductsTrend(ctx, start, end, interval)
	})
	if err != nil {
		return nil, err
	}

	trends, ok := data.([]entity.DeliveredProductTrend)
	if !ok {
		return nil, ce.NewError(ce.CommonErr, "error occurred")
	}

	return trends, nil
}

func (puc ProductUsecaseImpl) GetProductStatusTrend(ctx context.Context, start, end time.Time) ([]entity.ProductStatusTrend, error) {
	data, err := puc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return puc.pr.SelectProductStatusTrend(ctx, start, end)
	})
	if err != nil {
		return nil, err
	}

	statuses, ok := data.([]entity.ProductStatusTrend)
	if !ok {
		return nil, ce.NewError(ce.CommonErr, "error occurred")
	}

	return statuses, nil
}
