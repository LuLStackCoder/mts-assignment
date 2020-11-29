package httpserver

import (
	"context"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/LuLStackCoder/mts-assignment/pkg/api"
)

type service interface {
	HandleUrls(ctx context.Context, urls []string) (data []api.URLData, errorFlag bool, errorText string, err error)
}

type handleUrls struct {
	transport      HandleUrlsTransport
	timeout        time.Duration
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *handleUrls) ServeHTTP(ctx *fasthttp.RequestCtx) {
	var (
		urls      []string
		data      []api.URLData
		errorFlag bool
		errorText string
		err       error
	)
	// propagate timeout for request
	tCtx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	// decode request from json
	urls, err = s.transport.DecodeRequest(tCtx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(tCtx, &ctx.Response, err)
		return
	}
	// service logic
	data, errorFlag, errorText, err = s.service.HandleUrls(tCtx, urls)
	if err != nil {
		s.errorProcessor.Encode(tCtx, &ctx.Response, err)
		return
	}
	// encode from data to json
	if err = s.transport.EncodeResponse(tCtx, &ctx.Response, data, errorFlag, errorText); err != nil {
		s.errorProcessor.Encode(tCtx, &ctx.Response, err)
		return
	}
}

// NewHandleUrls the server creator
func NewHandleUrls(transport HandleUrlsTransport, service service, timeout time.Duration, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := handleUrls{
		transport:      transport,
		timeout:        timeout,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}
