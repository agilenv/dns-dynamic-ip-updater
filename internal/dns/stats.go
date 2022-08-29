package dns

import (
	"github.com/agilenv/linkip/internal/dns/track"
)

type Stats struct {
	tracks TrackRepository
}

func NewStats(tracks TrackRepository) *Stats {
	return &Stats{
		tracks: tracks,
	}
}

func (u *Stats) LastExecution() *track.Event {
	if e := u.tracks.LastEvent(); e.IP != "" {
		return &e
	}
	return nil
}

func (u *Stats) Save(event track.Event) error {
	return u.tracks.Save(event)
}
