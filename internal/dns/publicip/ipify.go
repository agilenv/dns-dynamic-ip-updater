package publicip

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/agilenv/linkip/pkg/rest"
)

const (
	ipifyURL = "https://api.ipify.org/?format=text"
)

type ipify struct {
	http *rest.Client
	name string
}

func NewIpifyPublicIPAPI(http *rest.Client) *ipify {
	return &ipify{
		http: http,
		name: "Ipify",
	}
}

func (i *ipify) Get(ctx context.Context) (string, error) {
	resp, err := i.http.Do(rest.NewRequest(http.MethodGet, ipifyURL).WithContext(ctx))
	if err != nil {
		return "", err
	}
	if resp.StatusCode() == http.StatusOK {
		ip := strings.TrimSpace(string(resp.Body()))
		if err = validate(ip); err != nil {
			return "", err
		}
		return ip, nil
	}
	return "", fmt.Errorf("cannot get public ID from API")
}

func (i *ipify) Name() string {
	return i.name
}
