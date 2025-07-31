package usecase

import (
	"context"
	"time"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/constant"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
	"github.com/adriel-timoteo/olist-fullstack/backend/repository"
)

type OrderUsecaseItf interface {
	GetDeliveredTrend(context.Context, time.Time, time.Time, constant.Interval) ([]entity.DeliveredOrderTrend, error)
	GetOrderStatusTrend(context.Context, time.Time, time.Time) ([]entity.OrderStatusTrend, error)
	GetOnTimeDeliveryRate(context.Context) (*entity.Rate, error)
}

type OrderUsecaseImpl struct {
	or  repository.OrderRepoItf
	trx repository.Transactor
}

func NewOrderUsecaseImpl(or repository.OrderRepoItf, trx repository.Transactor) OrderUsecaseImpl {
	return OrderUsecaseImpl{
		or:  or,
		trx: trx,
	}
}

func (puc OrderUsecaseImpl) GetDeliveredTrend(ctx context.Context, start, end time.Time, interval constant.Interval) ([]entity.DeliveredOrderTrend, error) {
	data, err := puc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return puc.or.SelectDeliveredOrdersTrend(ctx, start, end, interval)
	})
	if err != nil {
		return nil, err
	}

	trends, ok := data.([]entity.DeliveredOrderTrend)
	if !ok {
		return nil, ce.NewError(ce.InternalError, "error occurred")
	}

	return trends, nil
}

func (puc OrderUsecaseImpl) GetOrderStatusTrend(ctx context.Context, start, end time.Time) ([]entity.OrderStatusTrend, error) {
	data, err := puc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return puc.or.SelectOrderStatusTrend(ctx, start, end)
	})
	if err != nil {
		return nil, err
	}

	statuses, ok := data.([]entity.OrderStatusTrend)
	if !ok {
		return nil, ce.NewError(ce.InternalError, "error occurred")
	}

	return statuses, nil
}

func (puc OrderUsecaseImpl) GetOnTimeDeliveryRate(ctx context.Context) (*entity.Rate, error) {
	data, err := puc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return puc.or.SelectOnTimeDeliveryRate(ctx)
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
