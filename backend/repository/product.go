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
	SelectTopCategories(context.Context, int) ([]entity.ProductCategoryCount, error)
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

func (pr ProductRepoImpl) SelectTopCategories(ctx context.Context, limit int) ([]entity.ProductCategoryCount, error) {

	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	q := `
		SELECT
			p.product_category_name, count(*) AS purchase_count
		FROM order_items oi
		JOIN products p ON oi.product_id = p.product_id 
		GROUP BY p.product_category_name
		ORDER BY purchase_count DESC
		LIMIT $1;
	`

	rows, err := tx.QueryContext(ctx, q, limit)
	if err != nil {
		return nil, ce.NewError(ce.DatabaseError, "query execution failed")
	}
	defer rows.Close()

	var results []entity.ProductCategoryCount
	for rows.Next() {
		var record entity.ProductCategoryCount
		if err := rows.Scan(&record.Category, &record.ProductCount); err != nil {
			return nil, ce.NewError(ce.DatabaseError, "failed to scan row")
		}
		results = append(results, record)
	}

	if err = rows.Err(); err != nil {
		return nil, ce.NewError(ce.DatabaseError, "row iteration error")
	}

	return results, nil
}
