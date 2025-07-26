package repository

import (
	"context"
	"database/sql"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
)

type CustomerRepoItf interface {
	SelectTopCities(context.Context, int) ([]entity.CustomerCityCount, error)
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
