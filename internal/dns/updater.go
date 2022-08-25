package dns

type updater struct {
	dns       DNSProvider
	statsRepo TrackRepository
	ip        PublicIPAPI
}

func NewUpdater(dns DNSProvider, stats TrackRepository, ip PublicIPAPI) *updater {
	return &updater{
		dns:       dns,
		statsRepo: stats,
		ip:        ip,
	}
}

func (u *updater) Sync() (interface{}, error) {
	return nil, nil
}
