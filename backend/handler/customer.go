package handler

import (
	"net/http"
	"strconv"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/dto"
	"github.com/adriel-timoteo/olist-fullstack/backend/usecase"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	cuc usecase.CustomerUsecaseItf
}

func NewCustomerHandler(cuc usecase.CustomerUsecaseItf) CustomerHandler {
	return CustomerHandler{
		cuc: cuc,
	}
}

func (ch CustomerHandler) GetTopCities(ctx *gin.Context) {
	limitParam := ctx.DefaultQuery("limit", "5")

	// Turn limit to int
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.InvalidAction, "id not valid"))
		return
	}

	// Call usecase
	cities, err := ch.cuc.GetTopCities(ctx, limit)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Map to DTO
	var citiesDto []dto.CustomerCity
	for _, c := range cities {
		citiesDto = append(citiesDto, dto.CustomerCity{
			City:  c.City,
			Count: c.CustomerCount,
		})
	}

	// Return as JSON
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully retrieved product status snapshot",
		Error:   nil,
		Data:    citiesDto,
	})
}
