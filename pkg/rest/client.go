package rest

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	defaultTimeout = 10 * time.Second
)

type Client struct {
	r *resty.Client
}

func NewClient() *Client {
	cli := resty.New()
	cli.SetTimeout(defaultTimeout)
	return &Client{
		r: cli,
	}
}

func (c *Client) Do(req *request) (*response, error) {
	resp, err := prepareRequest(c, req).Send()
	return newResponse(resp), err
}

func (c *Client) GetClient() *http.Client {
	return c.r.GetClient()
}

func prepareRequest(client *Client, req *request) *resty.Request {
	request_ := client.r.R()
	request_.Method = req.method
	request_.URL = req.url
	request_.SetHeaders(req.headers)
	request_.SetBody(req.body)
	request_.SetContext(req.ctx)
	if req.timeout != 0 {
		client.r.SetTimeout(req.timeout)
	}
	return request_
}
