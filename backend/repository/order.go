package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/constant"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
)

type OrderRepoItf interface {
	SelectDeliveredOrdersTrend(context.Context, time.Time, time.Time, constant.Interval) ([]entity.DeliveredOrderTrend, error)
	SelectOrderStatusTrend(context.Context, time.Time, time.Time) ([]entity.OrderStatusTrend, error)
	SelectOnTimeDeliveryRate(context.Context) (*entity.Rate, error)
}

type OrderRepoImpl struct {
}

func NewOrderRepo() OrderRepoImpl {
	return OrderRepoImpl{}
}

func (or OrderRepoImpl) SelectDeliveredOrdersTrend(ctx context.Context, start, end time.Time, interval constant.Interval) ([]entity.DeliveredOrderTrend, error) {

	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	groupByClause := fmt.Sprintf("DATE_TRUNC('%s', order_delivered_customer_date)", interval)
	q := fmt.Sprintf(`
    SELECT 
			%s AS delivery_period,
			COUNT(order_id) AS total_delivered_orders
		FROM orders
		WHERE order_delivered_customer_date IS NOT NULL
			AND order_delivered_customer_date >= $1
			AND order_delivered_customer_date <= $2
		GROUP BY delivery_period
		ORDER BY delivery_period;`, groupByClause)

	rows, err := tx.QueryContext(ctx, q, start, end)
	if err != nil {
		return nil, ce.NewError(ce.DatabaseError, err.Error())
	}
	defer rows.Close()

	var results []entity.DeliveredOrderTrend
	for rows.Next() {
		var record entity.DeliveredOrderTrend
		if err := rows.Scan(&record.DeliveryTime, &record.TotalDeliveredCount); err != nil {
			return nil, ce.NewError(ce.DatabaseError, "failed to scan row")
		}
		results = append(results, record)
	}

	if err = rows.Err(); err != nil {
		return nil, ce.NewError(ce.DatabaseError, "row iteration error")
	}

	return results, nil
}

func (or OrderRepoImpl) SelectOrderStatusTrend(ctx context.Context, start, end time.Time) ([]entity.OrderStatusTrend, error) {

	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	q := `
		SELECT
			DATE(o.order_purchase_timestamp) AS order_date,
			o.order_status,
			COUNT(o.order_id) AS order_count
		FROM orders o
		WHERE DATE(o.order_purchase_timestamp) BETWEEN $1 AND $2
		GROUP BY order_date, o.order_status
		ORDER BY order_date ASC, order_count DESC;
	`

	rows, err := tx.QueryContext(ctx, q, start, end)
	if err != nil {
		return nil, ce.NewError(ce.DatabaseError, "query execution failed")
	}
	defer rows.Close()

	var results []entity.OrderStatusTrend
	for rows.Next() {
		var record entity.OrderStatusTrend
		if err := rows.Scan(&record.OrderTime, &record.Status, &record.OrderCount); err != nil {
			return nil, ce.NewError(ce.DatabaseError, "failed to scan row")
		}
		results = append(results, record)
	}

	if err = rows.Err(); err != nil {
		return nil, ce.NewError(ce.DatabaseError, "row iteration error")
	}

	return results, nil
}

func (or OrderRepoImpl) SelectOnTimeDeliveryRate(ctx context.Context) (*entity.Rate, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	q := `
		SELECT 
			COUNT(*) FILTER (WHERE order_delivered_customer_date <= order_estimated_delivery_date)::float / COUNT(*) AS on_time_delivery_rate_percent
		FROM orders
		WHERE order_delivered_customer_date IS NOT NULL;
	`

	var rate float64
	err := tx.QueryRowContext(ctx, q).Scan(&rate)
	if err != nil {
		return nil, ce.NewError(ce.DatabaseError, "query execution failed")
	}

	return &entity.Rate{Rate: rate}, nil
}
