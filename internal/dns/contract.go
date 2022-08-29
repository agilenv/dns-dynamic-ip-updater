package dns

import (
	"context"

	"github.com/agilenv/linkip/internal/dns/track"
)

//go:generate mockgen --package=dns --source=contract.go --destination=contract_mock.go DNSRecord,StatsRepository,PublicIP

type DNSProvider interface {
	GetRecord(ctx context.Context) (string, error)
	UpdateRecord(ctx context.Context, ip string) error
}

type TrackRepository interface {
	LastEvent() track.Event
	Save(event track.Event) error
}

type PublicIPAPI interface {
	Get(ctx context.Context) (string, error)
	Name() string
}
