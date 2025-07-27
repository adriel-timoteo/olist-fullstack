package ce

import "net/http"

type CustomError struct {
	ErrorCode    string
	ErrorMessage string
}

func (err CustomError) Error() string {
	return err.ErrorMessage
}

func NewError(errCode string, errMsg string) *CustomError {
	return &CustomError{
		ErrorCode:    errCode,
		ErrorMessage: errMsg,
	}
}

func (err CustomError) GetHTTPErrorCode() int {
	switch err.ErrorCode {
	case ValidationError:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case NotFound:
		return http.StatusNotFound
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case ServiceUnavailable:
		return http.StatusServiceUnavailable
	case TimeoutError:
		return http.StatusGatewayTimeout
	default:
		return http.StatusInternalServerError
	}
}

const (
	ValidationError = "VALIDATION_ERROR"
	Unauthorized    = "UNAUTHORIZED"
	Forbidden       = "FORBIDDEN"
	NotFound        = "NOT_FOUND"
	Conflict        = "CONFLICT"
	BadRequest      = "BAD_REQUEST"

	InternalError      = "INTERNAL_ERROR"
	DatabaseError      = "DATABASE_ERROR"
	ServiceUnavailable = "SERVICE_UNAVAILABLE"
	TimeoutError       = "TIMEOUT_ERROR"
)
