package httpserver

import (
	"context"
	"net/http/pprof"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const (
	httpMethodHandleUrls = "POST"
	uriPathHandleUrls    = "/api/v1/handle"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

// New ...
func New(router *fasthttprouter.Router, svc service, timeout time.Duration, errorProcessor errorProcessor) {

	handleUrlsTransport := NewHandleUrlsTransport()

	router.Handle(httpMethodHandleUrls, uriPathHandleUrls, NewHandleUrls(handleUrlsTransport, svc, timeout, errorProcessor))

	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
}
