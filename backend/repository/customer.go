package repository

import (
	"context"
	"database/sql"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
)

type CustomerRepoItf interface {
	SelectTopCities(context.Context, int) ([]entity.CustomerCityCount, error)
	SelectTotalUniqueCustomer(context.Context) (*entity.Count, error)
	SelectRepeatPurchaseRate(context.Context) (*entity.Rate, error)
}

type CustomerRepoImpl struct {
}

func NewCustomerRepo() CustomerRepoImpl {
	return CustomerRepoImpl{}
}

func (cr CustomerRepoImpl) SelectTopCities(ctx context.Context, limit int) ([]entity.CustomerCityCount, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	q := `
		SELECT
			customer_city,
			COUNT(DISTINCT customer_id) AS customer_count
		FROM customers
		GROUP BY customer_city
		ORDER BY customer_count DESC
		LIMIT $1;
	`

	rows, err := tx.QueryContext(ctx, q, limit)
	if err != nil {
		return nil, ce.NewError(ce.DatabaseError, "query execution failed")
	}
	defer rows.Close()

	var results []entity.CustomerCityCount
	for rows.Next() {
		var record entity.CustomerCityCount
		if err := rows.Scan(&record.City, &record.CustomerCount); err != nil {
			return nil, ce.NewError(ce.DatabaseError, "failed to scan row")
		}
		results = append(results, record)
	}

	if err = rows.Err(); err != nil {
		return nil, ce.NewError(ce.DatabaseError, "row iteration error")
	}

	return results, nil
}

func (cr CustomerRepoImpl) SelectTotalUniqueCustomer(ctx context.Context) (*entity.Count, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	q := `
		SELECT
			COUNT(DISTINCT customer_unique_id)
		AS total_customers
		FROM customers;
	`

	var count int
	err := tx.QueryRowContext(ctx, q).Scan(&count)
	if err != nil {
		return nil, ce.NewError(ce.DatabaseError, "query execution failed")
	}

	return &entity.Count{Count: count}, nil
}

func (cr CustomerRepoImpl) SelectRepeatPurchaseRate(ctx context.Context) (*entity.Rate, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	q := `
		SELECT 
			100.0 * COUNT(*) FILTER (WHERE order_count > 1) / COUNT(*) AS repeat_purchase_rate_percent
		FROM (
			SELECT customer_unique_id, COUNT(o.order_id) AS order_count
			FROM customers c
			JOIN orders o ON c.customer_id = o.customer_id
			GROUP BY customer_unique_id
		) AS sub;
	`

	var rate float64
	err := tx.QueryRowContext(ctx, q).Scan(&rate)
	if err != nil {
		return nil, ce.NewError(ce.DatabaseError, "query execution failed")
	}

	return &entity.Rate{Rate: rate}, nil
}
