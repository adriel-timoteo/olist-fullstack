package dto

import "time"

type Response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	Error     *ErrorRes   `json:"error,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp string      `json:"timestamp"`
}

type ErrorRes struct {
	Code    string       `json:"code"`
	Message string       `json:"message,omitempty"`
	Fields  []FieldError `json:"fields,omitempty"` // For validation errors
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type GetPageResponse struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	CountData int `json:"count_data"`
	Data      any `json:"data,omitempty"`
}

func SuccessResponse(data interface{}, message string) Response {
	return Response{
		Success:   true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
}
