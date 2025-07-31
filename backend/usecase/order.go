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
	GetTotalRevenue(context.Context) (*entity.Count, error)
	GetAverageOrderValue(context.Context) (*entity.Count, error)
	GetAverageDeliveryTime(context.Context) (*entity.Count, error)
	GetOrdersByHour(context.Context) ([]entity.OrderByHour, error)
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

func (ouc OrderUsecaseImpl) GetDeliveredTrend(ctx context.Context, start, end time.Time, interval constant.Interval) ([]entity.DeliveredOrderTrend, error) {
	data, err := ouc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return ouc.or.SelectDeliveredOrdersTrend(ctx, start, end, interval)
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

func (ouc OrderUsecaseImpl) GetOrderStatusTrend(ctx context.Context, start, end time.Time) ([]entity.OrderStatusTrend, error) {
	data, err := ouc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return ouc.or.SelectOrderStatusTrend(ctx, start, end)
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

func (ouc OrderUsecaseImpl) GetOnTimeDeliveryRate(ctx context.Context) (*entity.Rate, error) {
	data, err := ouc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return ouc.or.SelectOnTimeDeliveryRate(ctx)
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

func (ouc OrderUsecaseImpl) GetTotalRevenue(ctx context.Context) (*entity.Count, error) {
	data, err := ouc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return ouc.or.SelectGrossRevenue(ctx)
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

func (ouc OrderUsecaseImpl) GetAverageOrderValue(ctx context.Context) (*entity.Count, error) {
	data, err := ouc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return ouc.or.SelectAverageOrderValue(ctx)
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

func (ouc OrderUsecaseImpl) GetAverageDeliveryTime(ctx context.Context) (*entity.Count, error) {
	data, err := ouc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return ouc.or.SelectAverageDeliveryTime(ctx)
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

func (ouc OrderUsecaseImpl) GetOrdersByHour(ctx context.Context) ([]entity.OrderByHour, error) {
	data, err := ouc.trx.WithinTransaction(ctx, func(ctx context.Context) (any, error) {
		return ouc.or.SelectOrdersByHour(ctx)
	})
	if err != nil {
		return nil, err
	}

	orders, ok := data.([]entity.OrderByHour)
	if !ok {
		return nil, ce.NewError(ce.InternalError, "error occurred")
	}

	return orders, nil
}
