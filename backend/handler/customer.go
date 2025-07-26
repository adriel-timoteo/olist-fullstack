package handler

import (
	"github.com/adriel-timoteo/olist-fullstack/backend/usecase"
)

type CustomerHandler struct {
	cuc usecase.CustomerUsecaseItf
}

func NewCustomerHandler(cuc usecase.CustomerUsecaseItf) CustomerHandler {
	return CustomerHandler{
		cuc: cuc,
	}
}
