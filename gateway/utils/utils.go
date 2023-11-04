package utils

import (
	"errors"
)

var (
	BadRequest = errors.New("bad request")
	NotFound   = errors.New("not found")
  InternalServer = errors.New("internal server error")
)

func ErrorResponse(err error) map[string]any {
	return map[string]any{
		"error": err.Error(),
	}
}

func ToResponseBody(data interface{}) map[string]any {
	return map[string]any{
		"data": data,
	}
}
