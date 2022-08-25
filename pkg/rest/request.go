package rest

import (
	"context"
	"time"
)

type request struct {
	method,
	url string
	headers map[string]string
	body    []byte
	ctx     context.Context
	timeout time.Duration
}

func NewRequest(method, url string) *request {
	return &request{
		method: method,
		url:    url,
	}
}

func (r *request) WithHeaders(headers map[string]string) *request {
	r.headers = headers
	return r
}

func (r *request) WithBody(body []byte) *request {
	r.body = body
	return r
}

func (r *request) WithContext(ctx context.Context) *request {
	r.ctx = ctx
	return r
}

func (r *request) WithTimeout(duration time.Duration) *request {
	r.timeout = duration
	return r
}
