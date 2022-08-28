package dns

import (
	"context"

	"github.com/agilenv/dns-dynamic-ip-updater/internal/dns/track"
)

type updater struct {
	provider    DNSProvider
	stats       TrackRepository
	publicIPAPI PublicIPAPI
}

func NewUpdater(provider DNSProvider, stats TrackRepository, publicIPAPI PublicIPAPI) *updater {
	return &updater{
		provider:    provider,
		stats:       stats,
		publicIPAPI: publicIPAPI,
	}
}

func (u *updater) Sync(ctx context.Context) error {
	event := u.stats.LastEvent()
	publicIP, err := u.publicIPAPI.Get(ctx)
	if err != nil {
		return err
	}
	return u.handleChanges(ctx, publicIP, event.IP)
}

func (u *updater) handleChanges(ctx context.Context, publicIP, lastSavedIP string) error {
	if publicIP != lastSavedIP {
		if err := u.provider.UpdateRecord(ctx, publicIP); err != nil {
			return err
		}
		_ = u.stats.Save(track.Event{IP: publicIP, PublicAPI: u.publicIPAPI.Name()})
	}
	return nil
}
