package service

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/LuLStackCoder/mts-assignment/pkg/api"
	"github.com/LuLStackCoder/mts-assignment/pkg/service/httperrors"
)

// URLClient interface represent the client for getting data
type URLClient interface {
	GetData(ctx context.Context, url string) (data []byte, err error)
}

// Service implements the service logic
type Service interface {
	HandleUrls(ctx context.Context, urls []string) (data []api.URLData, err error)
}

type service struct {
	client  URLClient
	maxUrls int
}

func (s *service) HandleUrls(ctx context.Context, urls []string) (data []api.URLData, err error) {
	// check len urls
	if len(urls) == 0 {
		err = httperrors.ErrZeroURLs
		return
	}

	if len(urls) > s.maxUrls {
		err = httperrors.ErrLimitURLs
		return
	}

	// validating urls
	for i := range urls {
		_, err = url.Parse(urls[i])
		if err != nil {
			err = errors.Wrap(httperrors.ErrParseURL, err.Error())
			return
		}
	}

	// creation errgroup for convenient goroutine handling
	g, ctx := errgroup.WithContext(ctx)

	data = make([]api.URLData, len(urls))

	for i := range urls {
		iter := i // closure feature
		g.Go(func() error {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			// requesting body of every url
			body, err := s.client.GetData(ctx, urls[iter])
			if err != nil {
				return errors.Wrap(httperrors.ErrURLHandle, err.Error())
			}

			data[iter] = api.URLData{URL: urls[iter], Body: string(body)}
			return nil
		})
	}

	err = g.Wait()

	return
}

// NewService constructor
func NewService(client URLClient, maxUrls int) Service {
	return &service{
		client:  client,
		maxUrls: maxUrls,
	}
}
