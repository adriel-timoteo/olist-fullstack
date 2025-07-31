package repository

import (
	"context"
	"database/sql"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
)

type ProductRepoItf interface {
	SelectTopCategories(context.Context, int) ([]entity.ProductCategoryCount, error)
}

type ProductRepoImpl struct {
}

func NewProductRepo() ProductRepoImpl {
	return ProductRepoImpl{}
}

func (pr ProductRepoImpl) SelectTopCategories(ctx context.Context, limit int) ([]entity.ProductCategoryCount, error) {

	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	q := `
		SELECT
			ct.product_category_name_english, count(*) AS purchase_count
		FROM order_items oi
		JOIN products p ON oi.product_id = p.product_id
		JOIN category_translations ct ON p.product_category_name = ct.product_category_name
		WHERE p.product_category_name IS NOT NULL
		GROUP BY ct.product_category_name_english
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
			return nil, ce.NewError(ce.DatabaseError, err.Error())
		}
		results = append(results, record)
	}

	if err = rows.Err(); err != nil {
		return nil, ce.NewError(ce.DatabaseError, "row iteration error")
	}

	return results, nil
}
