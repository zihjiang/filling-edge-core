
package controlhub

const (
	DefaultBaseUrl              = "http://localhost:18631"
	AllLabel                    = "all"
	JobRunnerApp                = "jobrunner-app"
	TimeSeriesApp               = "timeseries-app"
	DefaultPingFrequency        = 5000
	DefaultStatusEventsInterval = 60000
)

type Config struct {
	Enabled                bool     `toml:"enabled"`
	BaseUrl                string   `toml:"base-url"`
	AppAuthToken           string   `toml:"app-auth-token"`
	JobLabels              []string `toml:"job-labels"`
	EventsRecipient        string   `toml:"events-recipient"`
	ProcessEventsRecipient []string `toml:"process-events-recipients"`
	PingFrequency          int      `toml:"ping-frequency"`
	StatusEventsInterval   int      `toml:"status-events-interval"`
}

// NewConfig returns a new Config with default settings.
func NewConfig() Config {
	return Config{
		Enabled:                false,
		BaseUrl:                DefaultBaseUrl,
		AppAuthToken:           "",
		JobLabels:              []string{AllLabel},
		EventsRecipient:        JobRunnerApp,
		ProcessEventsRecipient: []string{JobRunnerApp, TimeSeriesApp},
		PingFrequency:          DefaultPingFrequency,
		StatusEventsInterval:   DefaultStatusEventsInterval,
	}
}
