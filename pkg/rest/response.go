package rest

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

type response struct {
	code        int
	body        []byte
	rawResponse *http.Response
}

func newResponse(resp *resty.Response) *response {
	return &response{
		code:        resp.StatusCode(),
		body:        resp.Body(),
		rawResponse: resp.RawResponse,
	}
}

func (r *response) StatusCode() int {
	return r.code
}

func (r *response) Body() []byte {
	return r.body
}

func (r *response) Raw() *http.Response {
	return r.rawResponse
}
