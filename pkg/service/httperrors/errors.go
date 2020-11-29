package httperrors

import (
	"errors"
)

// Transport errors
var (
	ErrFailedToDecodeJSON = errors.New("failed to decode JSON response: %s")
	ErrFailedToEncodeJSON = errors.New("failed to encode JSON response: %s")
)

// Service errors
var (
	ErrZeroURLs  = errors.New("number of urls is zero")
	ErrLimitURLs = errors.New("number of urls over maxUrls")
	ErrURLHandle = errors.New("error to handle url")
)

// StatusMap for status codes
var (
	StatusMap = map[error]int{
		ErrFailedToDecodeJSON: 400,
		ErrFailedToEncodeJSON: 400,
		ErrZeroURLs:           400,
		ErrLimitURLs:          400,
		ErrURLHandle:          500,
	}
)
