package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/entity"
)

type UserRepoItf interface {
	CheckIsEmailExist(context.Context, entity.ReqRegisterUser) (bool, error)
	RegisterUser(context.Context, entity.ReqRegisterUser) (*entity.User, error)
	GetUserByEmail(context.Context, entity.ReqLoginUser) (*entity.User, error)
}

type UserRepoImpl struct {
}

func NewUserRepo() UserRepoImpl {
	return UserRepoImpl{}
}

func (ur UserRepoImpl) CheckIsEmailExist(ctx context.Context, req entity.ReqRegisterUser) (bool, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return false, ce.NewError(ce.DatabaseError, "internal server error")
	}

	q := `
		SELECT 
			id 
		FROM 
			users
		WHERE 
			email = $1`

	err := tx.QueryRowContext(ctx, q, req.Email).Scan(new(string))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		return false, ce.NewError(ce.DatabaseError, "internal server error")
	}

	return true, nil
}

func (ur UserRepoImpl) RegisterUser(ctx context.Context, req entity.ReqRegisterUser) (*entity.User, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	var user entity.User

	q := `
		INSERT INTO 
			users (email, password, created_at, updated_at) 
		VALUES 
			($1, $2, NOW(), NOW())
		RETURNING 
			id, email`

	err := tx.QueryRowContext(ctx, q, req.Email, req.Password).Scan(&user.Id, &user.Email)

	if err != nil {
		return nil, ce.NewError(ce.DatabaseError, "error register user")
	}

	return &user, nil
}

func (ur UserRepoImpl) GetUserByEmail(ctx context.Context, req entity.ReqLoginUser) (*entity.User, error) {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if !ok {
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	var user entity.User

	q := `
		select 
			id, email, password
		from 
			users
		where 
			email = $1`

	err := tx.QueryRowContext(ctx, q, req.Email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ce.NewError(ce.InvalidAction, "invalid credentials")
		}
		return nil, ce.NewError(ce.DatabaseError, "internal server error")
	}

	return &user, nil
}
