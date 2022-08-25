package dns

//go:generate mockgen --package=dns --source=contract.go --destination=contract_mock.go DNSRecord,StatsRepository,PublicIP

type DNSProvider interface {
	GetRecord(name string) (string, error)
	UpdateRecord(name string, ip string) error
}

type TrackRepository interface {
	LastEvent() (interface{}, error)
	Save(stats interface{}) error
}

type PublicIPAPI interface {
	Get() (string, error)
}
