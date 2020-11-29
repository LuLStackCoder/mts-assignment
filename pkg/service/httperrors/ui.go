package httperrors

import (
	"context"
	"errors"

	"github.com/valyala/fasthttp"
)

// ErrorProcessor ...
type ErrorProcessor struct {
	errors map[error]int
}

//Encode writes a svc error to the given http.ResponseWriter.
func (e *ErrorProcessor) Encode(ctx context.Context, r *fasthttp.Response, err error) {
	er := errors.Unwrap(errors.Unwrap(err))

	r.SetStatusCode(e.errors[er])
	r.Header.Set("Content-Type", "text/plain")

	r.SetBody([]byte(err.Error()))
}

// NewErrorProcessor ...
func NewErrorProcessor(errors map[error]int) *ErrorProcessor {
	return &ErrorProcessor{
		errors: errors,
	}
}
