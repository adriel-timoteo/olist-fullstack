package middleware

import (
	"errors"
	"strconv"
	"strings"

	"github.com/adriel-timoteo/olist-fullstack/backend/ce"
	"github.com/adriel-timoteo/olist-fullstack/backend/constant"
	"github.com/adriel-timoteo/olist-fullstack/backend/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authHeader) == 0 {
			ctx.Error(ce.NewError(ce.Unauthorized, "authorization header not found"))
			ctx.Abort()
			return
		}

		splitHeader := strings.Split(ctx.GetHeader(authorizationHeaderKey), " ")
		if len(splitHeader) != 2 {
			ctx.Error(ce.NewError(ce.Unauthorized, "invalid token format"))
			ctx.Abort()
			return
		}

		authType := strings.ToLower(splitHeader[0])
		if authType != authorizationTypeBearer {
			ctx.Error(ce.NewError(ce.Unauthorized, "unsupported authorization type"))
			ctx.Abort()
			return
		}

		token := splitHeader[1]

		claims, err := util.ParseJWTToken(token)
		if err != nil {
			errMsg := "cannot parse token"
			if errors.Is(err, jwt.ErrTokenExpired) {
				errMsg = "token has expired"
			}
			ctx.Error(ce.NewError(ce.Unauthorized, errMsg))
			ctx.Abort()
			return
		}

		userId, err := claims.GetSubject()
		if err != nil {
			ctx.Error(ce.NewError(ce.Unauthorized, "auth error"))
			ctx.Abort()
			return
		}

		intUserId, err := strconv.Atoi(userId)
		if err != nil {
			ctx.Error(ce.NewError(ce.CommonErr, "auth error"))
			ctx.Abort()
			return
		}

		ctx.Set(constant.RequestUserId, intUserId)
		ctx.Next()
	}
}
