package handler

import (
	"net/http"
	"strconv"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
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
