package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/agilenv/linkip/pkg/rest"
)

const (
	doEndpoint = "https://api.digitalocean.com/v2/domains/{{domain}}/records/{{record_id}}"
)

type do struct {
	http   *rest.Client
	config DigitaloceanConfig
}

type DigitaloceanConfig struct {
	DomainName string
	RecordID   string
	Token      string
}

type doUpdateRecord struct {
	Data string `json:"data"`
}

type doResponse struct {
	DomainRecord struct {
		Data string `json:"data"`
	} `json:"domain_record"`
}

func NewDigitaloceanProvider(http *rest.Client, config DigitaloceanConfig) *do {
	return &do{
		http:   http,
		config: config,
	}
}

func (d *do) GetRecord(ctx context.Context) (string, error) {
	endpoint := prepareEndpoint(d.config.DomainName, d.config.RecordID)
	resp, err := d.http.Do(rest.NewRequest(http.MethodGet, endpoint).
		WithContext(ctx).
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + d.config.Token,
		}),
	)
	if err != nil {
		return "", err
	}
	if resp.StatusCode() == http.StatusOK {
		r := &doResponse{}
		if err = json.Unmarshal(resp.Body(), &r); err != nil {
			return "", err
		}
		return r.DomainRecord.Data, nil
	}
	return "", fmt.Errorf("unexpected response from digitalocean provider")
}

func (d *do) UpdateRecord(ctx context.Context, ip string) error {
	endpoint := prepareEndpoint(d.config.DomainName, d.config.RecordID)
	resp, err := d.http.Do(rest.NewRequest(http.MethodPut, endpoint).
		WithContext(ctx).
		WithHeaders(map[string]string{
			"Authorization": "Bearer " + d.config.Token,
		}).
		WithBody(func() []byte {
			body := &doUpdateRecord{Data: ip}
			data, err := json.Marshal(body)
			if err != nil {
				return []byte{}
			}
			return data
		}()),
	)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return d.handleError(resp.StatusCode())
	}
	return nil
}

func prepareEndpoint(domain, recordID string) string {
	endpoint := strings.Replace(doEndpoint, "{{domain}}", domain, 1)
	endpoint = strings.Replace(endpoint, "{{record_id}}", recordID, 1)
	return endpoint
}

func (d *do) handleError(code int) error {
	var s string
	switch code {
	case 401:
		s = "Unauthorized. token appears to be invalid, are you sure it is correct?"
	case 404:
		s = "Resource not found. Are you sure domain name and domain record id environment variables are valid?"
	case 429:
		s = "API rate limit exceeded. There is a limitation on the number of request you can do in an hour."
	case 500:
		s = "Server error. Try again in a few minutes."
	default:
		s = "Unexpected error."
	}
	return fmt.Errorf("digital ocean => %s", s)
}
