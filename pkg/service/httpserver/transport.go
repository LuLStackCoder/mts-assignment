package httpserver

import (
	"context"

	"github.com/pkg/errors"

	"github.com/LuLStackCoder/mts-assignment/pkg/api"
	"github.com/LuLStackCoder/mts-assignment/pkg/service/httperrors"

	"github.com/valyala/fasthttp"
)

// HandleUrlsTransport transport interface
type HandleUrlsTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (urls []string, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, data []api.URLData, errorFlag bool, errorText string) (err error)
}

type handleUrlsTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *handleUrlsTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (urls []string, err error) {

	var request api.HandleUrlsRequest
	if err = request.UnmarshalJSON(r.Body()); err != nil {
		err = errors.Wrap(httperrors.ErrFailedToDecodeJSON, err.Error())
		return
	}

	urls = request

	return
}

// EncodeResponse method for encoding response on server side
func (t *handleUrlsTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, data []api.URLData, errorFlag bool, errorText string) (err error) {

	r.Header.Set("Content-Type", "application/json")
	var theResponse api.HandleUrlsResponse

	theResponse.Data = data

	theResponse.ErrorFlag = errorFlag

	theResponse.ErrorText = errorText

	body, err := theResponse.MarshalJSON()
	if err != nil {
		err = errors.Wrap(httperrors.ErrFailedToEncodeJSON, err.Error())
		return
	}
	r.SetBody(body)

	r.Header.SetStatusCode(200)
	return
}

// NewHandleUrlsTransport the transport creator for http requests
func NewHandleUrlsTransport() HandleUrlsTransport {
	return &handleUrlsTransport{}
}
