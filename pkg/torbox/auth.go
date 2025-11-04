package torbox

import (
	"fmt"
	"net/http"
)

type addAuthHeaderTransport struct {
	T      http.RoundTripper
	APIKey string
}

func (adt *addAuthHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	bearer := fmt.Sprintf("Bearer %s", adt.APIKey)
	req.Header.Add("Authorization", bearer)

	resp, err := adt.T.RoundTrip(req)
	return resp, err
}
