package httperrors

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

type errorResponse struct {
	Error     bool      `json:"error"`
	ErrorText string    `json:"errorText"`
	Data      *struct{} `json:"data"`
}

// ErrorProcessor ...
type ErrorProcessor struct {
	errors map[error]int
}

//Encode writes a svc error to the given http.ResponseWriter.
func (e *ErrorProcessor) Encode(ctx context.Context, r *fasthttp.Response, err error) {
	errorText := err.Error()
	er := errors.Unwrap(errors.Unwrap(err))

	res := errorResponse{
		Error:     true,
		ErrorText: errorText,
	}
	r.SetStatusCode(e.errors[er])
	r.Header.Set("Content-Type", "application/json")
	body, err := json.Marshal(res)
	if err != nil {
		return
	}
	r.SetBody(body)
}

// NewErrorProcessor ...
func NewErrorProcessor(errors map[error]int) *ErrorProcessor {
	return &ErrorProcessor{
		errors: errors,
	}
}
