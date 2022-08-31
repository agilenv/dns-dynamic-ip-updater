package dns

import (
	"context"

	"github.com/agilenv/linkip/internal/dns/track"
)

//go:generate mockgen --package=dns --source=updater.go --destination=updater_mock.go StatsUsecase

type StatsUsecase interface {
	Save(event track.Event) error
	LastExecution() *track.Event
}

type Updater struct {
	provider    DNSProvider
	publicIPAPI PublicIPAPI
	stats       StatsUsecase
}

func NewUpdater(provider DNSProvider, stats StatsUsecase, publicIPAPI PublicIPAPI) *Updater {
	return &Updater{
		provider:    provider,
		stats:       stats,
		publicIPAPI: publicIPAPI,
	}
}

func (u Updater) SearchForChanges(ctx context.Context) (bool, string, error) {
	publicIP, err := u.publicIPAPI.Get(ctx)
	if err != nil {
		return false, "", err
	}
	event := u.stats.LastExecution()
	if event != nil && event.IP == publicIP {
		return false, publicIP, nil
	}
	return true, publicIP, nil
}

func (u Updater) Update(ctx context.Context, ip string) error {
	if err := u.provider.UpdateRecord(ctx, ip); err != nil {
		return err
	}
	_ = u.stats.Save(track.Event{IP: ip, PublicAPI: u.publicIPAPI.Name()})
	return nil
}
