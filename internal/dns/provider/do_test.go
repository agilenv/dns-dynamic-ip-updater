package provider

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"

	"github.com/agilenv/linkip/pkg/rest"
)

func loadTestdata(filename string) []byte {
	b, _ := ioutil.ReadFile("./testdata/" + filename)
	return b
}

func Test_do_GetRecord(t *testing.T) {
	type fields struct {
		http   *rest.Client
		config DigitaloceanConfig
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected func(t *testing.T, f fields)
		want     string
		wantErr  bool
	}{
		{
			name: "should get an ip from dns provider successful response",
			fields: fields{
				http: rest.NewClient(),
				config: DigitaloceanConfig{
					DomainName: "foo",
					RecordID:   "bar",
					Token:      "token",
				},
			},
			args: args{ctx: context.Background()},
			expected: func(t *testing.T, f fields) {
				httpmock.ActivateNonDefault(f.http.GetClient())
				httpmock.RegisterResponder(
					http.MethodGet,
					prepareEndpoint("foo", "bar"),
					func(req *http.Request) (*http.Response, error) {
						if req.Header.Get("Authorization") != "Bearer "+f.config.Token {
							t.Errorf("Authorization header must exists and has to be Bearer <token>")
						}
						response := httpmock.NewBytesResponse(http.StatusOK, loadTestdata("do_get_record.json"))
						return response, nil
					})
			},
			want:    "1.2.3.4",
			wantErr: false,
		},
		{
			name: "should get an error when response returns error",
			fields: fields{
				http: rest.NewClient(),
				config: DigitaloceanConfig{
					DomainName: "foo",
					RecordID:   "bar",
					Token:      "token",
				},
			},
			args: args{ctx: context.Background()},
			expected: func(t *testing.T, f fields) {
				httpmock.ActivateNonDefault(f.http.GetClient())
				responder := httpmock.NewErrorResponder(errors.New("fail"))
				httpmock.RegisterResponder(http.MethodGet, prepareEndpoint("foo", "bar"), responder)
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDigitaloceanProvider(tt.fields.http, tt.fields.config)
			tt.expected(t, tt.fields)
			got, err := d.GetRecord(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_do_GetRecordStatusCodeError(t *testing.T) {
	respBody := loadTestdata("do_error_response.json")
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		response *http.Response
		args     args
		want     string
		wantErr  bool
	}{
		{
			name:     "should return an err when response is 401",
			response: httpmock.NewBytesResponse(http.StatusUnauthorized, respBody),
			args:     args{ctx: context.Background()},
			want:     "",
			wantErr:  true,
		},
		{
			name:     "should return an err when response is 404",
			response: httpmock.NewBytesResponse(http.StatusNotFound, respBody),
			args:     args{ctx: context.Background()},
			want:     "",
			wantErr:  true,
		},
		{
			name:     "should return an err when response is 429",
			response: httpmock.NewBytesResponse(http.StatusTooManyRequests, respBody),
			args:     args{ctx: context.Background()},
			want:     "",
			wantErr:  true,
		},
		{
			name:     "should return an err when responseis 500",
			response: httpmock.NewBytesResponse(http.StatusInternalServerError, respBody),
			args:     args{ctx: context.Background()},
			want:     "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rest.NewClient()
			d := NewDigitaloceanProvider(r, DigitaloceanConfig{})
			httpmock.ActivateNonDefault(r.GetClient())
			responder := httpmock.ResponderFromResponse(tt.response)
			httpmock.RegisterResponder(http.MethodGet, prepareEndpoint("", ""), responder)

			got, err := d.GetRecord(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_do_UpdateRecord(t *testing.T) {
	type fields struct {
		http   *rest.Client
		config DigitaloceanConfig
	}
	type args struct {
		ctx context.Context
		ip  string
	}
	tests := []struct {
		name     string
		fields   fields
		expected func(t *testing.T, f fields)
		args     args
		wantErr  bool
	}{
		{
			name: "should succeed when modifying the dns record on provider gets a successful response",
			fields: fields{http: rest.NewClient(), config: DigitaloceanConfig{
				DomainName: "foo",
				RecordID:   "bar",
				Token:      "token",
			}},
			expected: func(t *testing.T, f fields) {
				httpmock.ActivateNonDefault(f.http.GetClient())
				httpmock.RegisterResponder(
					http.MethodPut,
					prepareEndpoint("foo", "bar"),
					func(req *http.Request) (*http.Response, error) {
						if req.Header.Get("Authorization") != "Bearer "+f.config.Token {
							t.Errorf("AAuthorization header must exists and has to be Bearer <token>")
						}
						body, _ := ioutil.ReadAll(req.Body)
						reqBody := doUpdateRecord{}
						expectedBody := doUpdateRecord{}
						if err := json.Unmarshal(body, &reqBody); err != nil {
							t.Errorf("unmarshal err request body %s", err)
						}
						mockReqBody := loadTestdata("do_update_record.json")
						if err := json.Unmarshal(mockReqBody, &expectedBody); err != nil {
							t.Errorf("unmarshal mock request body %s", err)
						}
						if reqBody.Data != expectedBody.Data {
							t.Errorf("body request expects to be %s, got %s", string(mockReqBody), string(body))
						}
						response := httpmock.NewBytesResponse(http.StatusOK, nil)
						return response, nil
					})
			},
			args:    args{ctx: context.Background(), ip: "4.3.2.1"},
			wantErr: false,
		},
		{
			name: "should error when modifying the dns record on provider gets an error",
			fields: fields{http: rest.NewClient(), config: DigitaloceanConfig{
				DomainName: "foo",
				RecordID:   "bar",
				Token:      "token",
			}},
			expected: func(t *testing.T, f fields) {
				httpmock.ActivateNonDefault(f.http.GetClient())
				httpmock.RegisterResponder(
					http.MethodPut,
					prepareEndpoint("foo", "bar"),
					func(req *http.Request) (*http.Response, error) {
						return nil, errors.New("failed")
					},
				)
			},
			args:    args{ctx: context.Background(), ip: "4.3.2.1"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewDigitaloceanProvider(tt.fields.http, tt.fields.config)
			tt.expected(t, tt.fields)
			err := d.UpdateRecord(tt.args.ctx, tt.args.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_do_UpdateRecordStatusCodeError(t *testing.T) {
	respBody := loadTestdata("do_error_response.json")
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		response *http.Response
		args     args
		want     string
		wantErr  bool
	}{
		{
			name:     "should return an err when response is 401",
			response: httpmock.NewBytesResponse(http.StatusUnauthorized, respBody),
			args:     args{ctx: context.Background()},
			want:     "",
			wantErr:  true,
		},
		{
			name:     "should return an err when response is 404",
			response: httpmock.NewBytesResponse(http.StatusNotFound, respBody),
			args:     args{ctx: context.Background()},
			want:     "",
			wantErr:  true,
		},
		{
			name:     "should return an err when response is 429",
			response: httpmock.NewBytesResponse(http.StatusTooManyRequests, respBody),
			args:     args{ctx: context.Background()},
			want:     "",
			wantErr:  true,
		},
		{
			name:     "should return an err when responseis 500",
			response: httpmock.NewBytesResponse(http.StatusInternalServerError, respBody),
			args:     args{ctx: context.Background()},
			want:     "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rest.NewClient()
			d := NewDigitaloceanProvider(r, DigitaloceanConfig{})
			httpmock.ActivateNonDefault(r.GetClient())
			responder := httpmock.ResponderFromResponse(tt.response)
			httpmock.RegisterResponder(http.MethodPut, prepareEndpoint("", ""), responder)

			got, err := d.GetRecord(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}
