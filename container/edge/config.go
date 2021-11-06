package edge

import (
	"datacollector-edge/container/controlhub"
	"datacollector-edge/container/execution"
	"datacollector-edge/container/http"
	"datacollector-edge/container/process"
	"datacollector-edge/container/util"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"os"
)

// Config represents the configuration format for the Data Collector Edge binary.
type Config struct {
	LogDir    string `toml:"log-dir"`
	Execution execution.Config
	Http      http.Config
	SCH       controlhub.Config
	Process   process.Config
}

// NewConfig returns a new Config with default settings.
func NewConfig() *Config {
	c := &Config{}
	c.Execution = execution.NewConfig()
	c.Http = http.NewConfig()
	c.SCH = controlhub.NewConfig()
	c.Process = process.NewConfig()
	return c
}

// FromTomlFile loads the config from a TOML file.
func (c *Config) FromTomlFile(fPath string) error {
	if _, err := toml.DecodeFile(fPath, c); err != nil {
		return err
	}
	return nil
}

func (c *Config) ToTomlFile(fPath string) error {
	fi, err := os.OpenFile(fPath, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_ = fi.Truncate(0)
	defer util.CloseFile(fi)

	if err := toml.NewEncoder(fi).Encode(c); err != nil {
		log.WithError(err).Error()
		return err
	}
	return nil
}
