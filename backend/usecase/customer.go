package usecase

import (
	"github.com/adriel-timoteo/olist-fullstack/backend/repository"
)

type CustomerUsecaseItf interface {
}

type CustomerUsecaseImpl struct {
	cr  repository.CustomerRepoItf
	trx repository.Transactor
}

func NewCustomerUsecaseImpl(cr repository.CustomerRepoItf, trx repository.Transactor) CustomerUsecaseImpl {
	return CustomerUsecaseImpl{
		cr:  cr,
		trx: trx,
	}
}
