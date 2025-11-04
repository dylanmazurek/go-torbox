package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	method string
	path   string

	url       *url.URL
	header    *http.Header
	urlParams *url.Values
	body      io.ReadCloser
}

// NewRequest creates a new HTTP request with:
// method, path, and options.
func NewRequest(m string, p string, opts ...Option) (*Request, error) {
	requestOptions := defaultOptions()
	for _, opt := range opts {
		opt(&requestOptions)
	}

	path := fmt.Sprintf("%s/%s", requestOptions.baseUrl, p)
	newRequest := &Request{
		method: m,
		path:   path,
	}

	if requestOptions.body != nil {
		body := requestOptions.body

		bodyType := requestOptions.bodyType

		var contentType string
		switch bodyType {
		case BodyTypeForm:
			bodyBuffer, ok := body.(*bytes.Buffer)
			if !ok {
				return nil, fmt.Errorf("expected *bytes.Buffer for form body type, got %T", body)
			}

			newRequest.body = io.NopCloser(bodyBuffer)
			contentType = "application/x-www-form-urlencoded"
		case "file":
			bodyBuffer, ok := body.(*bytes.Buffer)
			if !ok {
				return nil, fmt.Errorf("expected *bytes.Buffer for file body type, got %T", body)
			}

			newRequest.body = io.NopCloser(bodyBuffer)

			if requestOptions.urlParams != nil {
				contentType = requestOptions.urlParams.Get("Content-Type")
			}
		default:
			bodyBytes, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}

			bodyBuffer := bytes.NewBuffer(bodyBytes)
			newRequest.body = io.NopCloser(bodyBuffer)

			contentType = "application/json"
		}

		newRequest.header.Set("Content-Type", contentType)
	}

	if requestOptions.urlParams != nil {
		newRequest.urlParams = requestOptions.urlParams
	}

	return newRequest, nil
}

func (r *Request) ToHTTPRequest() (*http.Request, error) {
	httpReq, err := http.NewRequest(r.method, r.path, r.body)
	if err != nil {
		return nil, err
	}

	if r.header != nil {
		httpReq.Header = *r.header
	}

	if r.urlParams != nil {
		httpReq.URL.RawQuery = r.urlParams.Encode()
	}

	return httpReq, nil
}
