package publicip

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/agilenv/linkip/pkg/rest"
)

func successfulFakeResponse(t *testing.T) []byte {
	r, err := ioutil.ReadFile("./testdata/ipify")
	if err != nil {
		t.Errorf("response mock file err: %v", err)
	}
	return r
}

func Test_ipify_Get(t *testing.T) {
	type fields struct {
		http *rest.Client
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected func(t *testing.T, client *rest.Client)
		want     string
		wantErr  bool
	}{
		{
			name:   "should return ip address when get a successful response from API",
			fields: fields{http: rest.NewClient()},
			args:   args{ctx: context.Background()},
			expected: func(t *testing.T, c *rest.Client) {
				httpmock.ActivateNonDefault(c.GetClient())
				response := httpmock.NewBytesResponse(http.StatusOK, successfulFakeResponse(t))
				responder := httpmock.ResponderFromResponse(response)
				httpmock.RegisterResponder(http.MethodGet, ipifyURL, responder)
			},
			want:    "1.2.3.4",
			wantErr: false,
		},
		{
			name:   "should return err when get a failed response from API",
			fields: fields{http: rest.NewClient()},
			args:   args{ctx: context.Background()},
			expected: func(t *testing.T, c *rest.Client) {
				httpmock.ActivateNonDefault(c.GetClient())
				responder := httpmock.NewErrorResponder(errors.New("failed"))
				httpmock.RegisterResponder(http.MethodGet, ipifyURL, responder)
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "should return error when resource API not found",
			fields: fields{http: rest.NewClient()},
			args:   args{ctx: context.Background()},
			expected: func(t *testing.T, c *rest.Client) {
				httpmock.ActivateNonDefault(c.GetClient())
				response := httpmock.NewBytesResponse(http.StatusNotFound, nil)
				responder := httpmock.ResponderFromResponse(response)
				httpmock.RegisterResponder(http.MethodGet, ipifyURL, responder)
			},
			want:    "",
			wantErr: true,
		},
		{
			name:   "should return error when API returns error 500",
			fields: fields{http: rest.NewClient()},
			args:   args{ctx: context.Background()},
			expected: func(t *testing.T, c *rest.Client) {
				httpmock.ActivateNonDefault(c.GetClient())
				response := httpmock.NewBytesResponse(http.StatusInternalServerError, nil)
				responder := httpmock.ResponderFromResponse(response)
				httpmock.RegisterResponder(http.MethodGet, ipifyURL, responder)
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &ipify{
				http: tt.fields.http,
			}
			tt.expected(t, i.http)
			got, err := i.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
