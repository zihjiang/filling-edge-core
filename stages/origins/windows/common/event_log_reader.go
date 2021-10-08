// +build 386 windows,amd64 windows


package common

import (
	"datacollector-edge/api"
)

type EventLogReader interface {
	Open() error
	Read() ([]api.Record, error)
	GetCurrentOffset() string
	Close() error
}

type BaseEventLogReader struct {
	Log          string
	Mode         EventLogReaderMode
	MaxBatchSize int
}

type EventLogReaderMode string
type EventLogReaderAPIType string

const (
	ReadAll                      = EventLogReaderMode("ALL")
	ReadNew                      = EventLogReaderMode("NEW")
	ReaderAPITypeEventLogging    = EventLogReaderAPIType("EVENT_LOGGING")
	ReaderAPITypeWindowsEventLog = EventLogReaderAPIType("WINDOWS_EVENT_LOG")
)

type CommonConf struct {
	LogName       string  `ConfigDef:"type=STRING,required=true"`
	ReadMode      string  `ConfigDef:"type=STRING,required=true"`
	CustomLogName string  `ConfigDef:"type=STRING,required=true"`
	BufferSize    float64 `ConfigDef:"type=NUMBER,required=true"`
}

type WinEventLogConf struct {
	SubscriptionMode           string  `ConfigDef:"type=STRING,required=true"`
	MaxWaitTimeSecs            float64 `ConfigDef:"type=NUMBER,required=true"`
	RawEventPopulationStrategy string  `ConfigDef:"type=STRING,required=true"`
}
