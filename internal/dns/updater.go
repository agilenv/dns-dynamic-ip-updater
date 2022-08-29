package dns

import (
	"context"

	"github.com/agilenv/linkip/internal/dns/track"
)

type Updater struct {
	provider    DNSProvider
	stats       TrackRepository
	publicIPAPI PublicIPAPI
}

func NewUpdater(provider DNSProvider, stats TrackRepository, publicIPAPI PublicIPAPI) *Updater {
	return &Updater{
		provider:    provider,
		stats:       stats,
		publicIPAPI: publicIPAPI,
	}
}

func (u *Updater) SearchForChanges(ctx context.Context) (bool, string, error) {
	event := u.stats.LastEvent()
	publicIP, err := u.publicIPAPI.Get(ctx)
	if err != nil {
		return false, "", err
	}
	if publicIP != event.IP {
		return true, publicIP, nil
	}
	return false, event.IP, nil
}

func (u *Updater) Update(ctx context.Context, ip string) error {
	if err := u.provider.UpdateRecord(ctx, ip); err != nil {
		return err
	}
	_ = u.stats.Save(track.Event{IP: ip, PublicAPI: u.publicIPAPI.Name()})
	return nil
}

func (u *Updater) LastExecution() *track.Event {
	if e := u.stats.LastEvent(); e.IP != "" {
		return &e
	}
	return nil
}
