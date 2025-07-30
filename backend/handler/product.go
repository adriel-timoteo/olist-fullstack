package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/constant"
	"github.com/adriel-timoteo/olist-fullstack/backend/dto"
	"github.com/adriel-timoteo/olist-fullstack/backend/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	puc usecase.ProductUsecaseItf
}

func NewProductHandler(puc usecase.ProductUsecaseItf) ProductHandler {
	return ProductHandler{
		puc: puc,
	}
}

func (ph ProductHandler) GetDeliveredTrend(ctx *gin.Context) {
	intervalParam := ctx.DefaultQuery("interval", "day")
	startParam := ctx.Query("start")
	endParam := ctx.Query("end")

	// Parse start date
	start, err := time.Parse(constant.DateFormat, startParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.ValidationError, "start date not valid"))
		return
	}

	// Parse end date
	end, err := time.Parse(constant.DateFormat, endParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.ValidationError, "end date not valid"))
		return
	}

	var interval constant.Interval
	switch intervalParam {
	case "day":
		interval = constant.IntervalDay
	case "month":
		interval = constant.IntervalMonth
	default:
		ctx.Error(ce.NewError(ce.ValidationError, "interval must be 'day' or 'month'"))
		return
	}

	// Call usecase
	trends, err := ph.puc.GetDeliveredTrend(ctx, start, end, interval)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Map to DTO
	var trendsDto []dto.DeliveredTrend
	for _, t := range trends {
		trendsDto = append(trendsDto, dto.DeliveredTrend{
			Time:  t.DeliveryTime.Format(constant.DateFormat),
			Count: t.TotalDeliveredCount,
		})
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse(trendsDto, "success"))
}

func (ph ProductHandler) GetProductStatusSnapshot(ctx *gin.Context) {
	startParam := ctx.DefaultQuery("start", time.Now().Format(constant.DateFormat))
	endParam := ctx.DefaultQuery("end", time.Now().Format(constant.DateFormat))

	// Parse start date
	start, err := time.Parse(constant.DateFormat, startParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.ValidationError, "start date not valid"))
		return
	}

	// Parse end date
	end, err := time.Parse(constant.DateFormat, endParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.ValidationError, "end date not valid"))
		return
	}

	// Call usecase
	statuses, err := ph.puc.GetProductStatusTrend(ctx, start, end)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Map to DTO
	var statusesDto []dto.ProductStatusSnapshot
	for _, s := range statuses {
		statusesDto = append(statusesDto, dto.ProductStatusSnapshot{
			Time:   s.OrderTime.Format(constant.DateFormat),
			Status: s.Status,
			Count:  s.OrderCount,
		})
	}

	// Return as JSON
	ctx.JSON(http.StatusOK, dto.SuccessResponse(statusesDto, "success"))
}

func (ph ProductHandler) GetTopCategories(ctx *gin.Context) {
	limitParam := ctx.DefaultQuery("limit", "5")

	// Turn limit to int
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.ValidationError, "id not valid"))
		return
	}

	// Call usecase
	categories, err := ph.puc.GetTopCategories(ctx, limit)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Map to DTO
	var catDto []dto.ProductCategoryCount
	for _, c := range categories {
		catDto = append(catDto, dto.ProductCategoryCount{
			Category: c.Category,
			Count:    c.ProductCount,
		})
	}

	// Return as JSON
	ctx.JSON(http.StatusOK, dto.SuccessResponse(catDto, "success"))
}

func (ph ProductHandler) GetOnTimeDeliveryRate(ctx *gin.Context) {
	// Call usecase
	rate, err := ph.puc.GetOnTimeDeliveryRate(ctx)
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
