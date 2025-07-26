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

type ProductRepoItf interface {
	SelectDeliveredProductsTrend(context.Context, time.Time, time.Time, constant.Interval) ([]entity.DeliveredProductTrend, error)
	SelectProductStatusTrend(context.Context, time.Time, time.Time) ([]entity.ProductStatusTrend, error)
}

type ProductRepoImpl struct {
}

func NewProductRepo() ProductRepoImpl {
	return ProductRepoImpl{}
}

func (pr ProductRepoImpl) SelectDeliveredProductsTrend(ctx context.Context, start, end time.Time, interval constant.Interval) ([]entity.DeliveredProductTrend, error) {

	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	groupByClause := fmt.Sprintf("DATE_TRUNC('%s', o.order_delivered_customer_date)", interval)
	q := fmt.Sprintf(`
    SELECT 
        %s AS delivery_period,
        COUNT(oi.order_id) AS total_delivered_products
    FROM orders o
    JOIN order_items oi ON o.order_id = oi.order_id
    WHERE o.order_delivered_customer_date IS NOT NULL
        AND o.order_delivered_customer_date >= $1
        AND o.order_delivered_customer_date <= $2
    GROUP BY delivery_period
    ORDER BY delivery_period;`, groupByClause)

	rows, err := tx.QueryContext(ctx, q, start, end)
	if err != nil {
		return nil, ce.NewError(ce.DatabaseError, "query execution failed")
	}
	defer rows.Close()

	var results []entity.DeliveredProductTrend
	for rows.Next() {
		var record entity.DeliveredProductTrend
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

func (pr ProductRepoImpl) SelectProductStatusTrend(ctx context.Context, start, end time.Time) ([]entity.ProductStatusTrend, error) {

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

	var results []entity.ProductStatusTrend
	for rows.Next() {
		var record entity.ProductStatusTrend
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
