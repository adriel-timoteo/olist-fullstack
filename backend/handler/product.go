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
		ctx.Error(ce.NewError(ce.InvalidAction, "start date not valid"))
		return
	}

	// Parse end date
	end, err := time.Parse(constant.DateFormat, endParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.InvalidAction, "end date not valid"))
		return
	}

	var interval constant.Interval
	switch intervalParam {
	case "day":
		interval = constant.IntervalDay
	case "month":
		interval = constant.IntervalMonth
	default:
		ctx.Error(ce.NewError(ce.InvalidAction, "interval must be 'day' or 'month'"))
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

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "success",
		Error:   nil,
		Data:    trendsDto,
	})
}

func (ph ProductHandler) GetProductStatusSnapshot(ctx *gin.Context) {
	startParam := ctx.DefaultQuery("start", time.Now().Format(constant.DateFormat))
	endParam := ctx.DefaultQuery("end", time.Now().Format(constant.DateFormat))

	// Parse start date
	start, err := time.Parse(constant.DateFormat, startParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.InvalidAction, "start date not valid"))
		return
	}

	// Parse end date
	end, err := time.Parse(constant.DateFormat, endParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.InvalidAction, "end date not valid"))
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
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "success",
		Error:   nil,
		Data:    statusesDto,
	})
}

func (ph ProductHandler) GetTopCategories(ctx *gin.Context) {
	limitParam := ctx.DefaultQuery("limit", "5")

	// Turn limit to int
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		ctx.Error(ce.NewError(ce.InvalidAction, "id not valid"))
		return
	}

	// Call usecase
	cities, err := ph.puc.GetTopCategories(ctx, limit)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Map to DTO
	var citiesDto []dto.ProductCategoryCount
	for _, c := range cities {
		citiesDto = append(citiesDto, dto.ProductCategoryCount{
			Category: c.Category,
			Count:    c.ProductCount,
		})
	}

	// Return as JSON
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "success",
		Error:   nil,
		Data:    citiesDto,
	})
}
