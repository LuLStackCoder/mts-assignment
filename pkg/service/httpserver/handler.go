package httpserver

import (
	"context"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// ErrorProcessor implements error handling
type ErrorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

// New ...
func New(router *fasthttprouter.Router, svc Service, timeout time.Duration, errorProcessor ErrorProcessor) {

	handleUrlsTransport := NewHandleUrlsTransport()

	router.Handle(httpMethodHandleUrls, uriPathHandleUrls, NewHandleUrls(handleUrlsTransport, svc, timeout, errorProcessor))
}
