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
		ctx.Error(ce.NewError(ce.ValidationError, "id not valid"))
		return
	}

	// Call usecase
	cities, err := ch.cuc.GetTopCities(ctx, limit)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Map to DTO
	var citiesDto []dto.CustomerCityCount
	for _, c := range cities {
		citiesDto = append(citiesDto, dto.CustomerCityCount{
			City:  c.City,
			Count: c.CustomerCount,
		})
	}

	// Return as JSON
	ctx.JSON(http.StatusOK, dto.SuccessResponse(citiesDto, "success"))
}

func (ch CustomerHandler) GetTotalUniqueCustomers(ctx *gin.Context) {
	// Call usecase
	count, err := ch.cuc.GetTotalUniqueCustomers(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Map to DTO
	countDto := dto.Count{
		Count: count.Count,
	}

	// Return as JSON
	ctx.JSON(http.StatusOK, dto.SuccessResponse(countDto, "success"))
}

func (ch CustomerHandler) GetRepeatPurchaseRate(ctx *gin.Context) {
	// Call usecase
	rate, err := ch.cuc.GetRepeatPurchaseRate(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Map to DTO
	rateDto := dto.Rate{
		Rate: rate.Rate,
	}

	// Return as JSON
	ctx.JSON(http.StatusOK, dto.SuccessResponse(rateDto, "success"))
}
