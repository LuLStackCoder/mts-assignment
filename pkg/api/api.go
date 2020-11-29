package api

// URLData ...
type URLData struct {
	URL  string `json:"url"`
	Body string `json:"body"`
}

// HandleUrlsRequest ...
//easyjson:json
type HandleUrlsRequest []string

// HandleUrlsResponse ...
//easyjson:json
type HandleUrlsResponse struct {
	Data []URLData `json:"data"`

	ErrorFlag bool `json:"error"`

	ErrorText string `json:"errorText"`
}
