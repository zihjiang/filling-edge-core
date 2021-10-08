
package http

const (
	DefaultBindAddress = ":18633"
)

type Config struct {
	Enabled     bool   `toml:"enabled"`
	BindAddress string `toml:"bind-address"`
	BaseHttpUrl string `toml:"base-http-url"`
}

// NewConfig returns a new Config with default settings.
func NewConfig() Config {
	return Config{
		Enabled:     true,
		BindAddress: DefaultBindAddress,
	}
}
