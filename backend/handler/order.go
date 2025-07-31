package handler

import (
	"net/http"
	"time"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/constant"
	"github.com/adriel-timoteo/olist-fullstack/backend/dto"
	"github.com/adriel-timoteo/olist-fullstack/backend/usecase"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	ouc usecase.OrderUsecaseItf
}

func NewOrderHandler(ouc usecase.OrderUsecaseItf) OrderHandler {
	return OrderHandler{
		ouc: ouc,
	}
}

func (oh OrderHandler) GetDeliveredTrend(ctx *gin.Context) {
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
	trends, err := oh.ouc.GetDeliveredTrend(ctx, start, end, interval)
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

func (oh OrderHandler) GetOrderStatusSnapshot(ctx *gin.Context) {
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
	statuses, err := oh.ouc.GetOrderStatusTrend(ctx, start, end)
	if err != nil {
		ctx.Error(err)
		return
	}

	// Map to DTO
	var statusesDto []dto.OrderStatusSnapshot
	for _, s := range statuses {
		statusesDto = append(statusesDto, dto.OrderStatusSnapshot{
			Time:   s.OrderTime.Format(constant.DateFormat),
			Status: s.Status,
			Count:  s.OrderCount,
		})
	}

	// Return as JSON
	ctx.JSON(http.StatusOK, dto.SuccessResponse(statusesDto, "success"))
}

func (oh OrderHandler) GetOnTimeDeliveryRate(ctx *gin.Context) {
	// Call usecase
	rate, err := oh.ouc.GetOnTimeDeliveryRate(ctx)
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

func (oh OrderHandler) GetTotalRevenue(ctx *gin.Context) {
	// Call usecase
	count, err := oh.ouc.GetTotalRevenue(ctx)
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

func (oh OrderHandler) GetAverageOrderValue(ctx *gin.Context) {
	// Call usecase
	count, err := oh.ouc.GetAverageOrderValue(ctx)
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
