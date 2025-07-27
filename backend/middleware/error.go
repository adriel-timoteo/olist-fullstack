package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors[0]
		timestamp := time.Now().UTC().Format(time.RFC3339)

		// Handle validation errors
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			fieldErrors := make([]dto.FieldError, 0, len(ve))
			for _, fe := range ve {
				fieldErrors = append(fieldErrors, dto.FieldError{
					Field:   fe.Field(),
					Message: fmt.Sprintf("Invalid input on field %s", fe.Field()),
				})
			}

			ctx.JSON(http.StatusBadRequest, dto.Response{
				Success: false,
				Error: &dto.ErrorRes{
					Code:    ce.ValidationError,
					Message: "One or more fields are invalid",
					Fields:  fieldErrors,
				},
				Timestamp: timestamp,
			})
			return
		}

		// Handle custom application errors
		var customErr *ce.CustomError
		if errors.As(err, &customErr) {
			ctx.JSON(customErr.GetHTTPErrorCode(), dto.Response{
				Success: false,
				Error: &dto.ErrorRes{
					Code:    customErr.ErrorCode,
					Message: customErr.Error(),
				},
				Timestamp: timestamp,
			})
			return
		}

		// Handle unexpected errors
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error: &dto.ErrorRes{
				Code:    ce.InternalError,
				Message: "An unexpected error occurred",
			},
			Timestamp: timestamp,
		})
	}
}
