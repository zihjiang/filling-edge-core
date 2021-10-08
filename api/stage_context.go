
package api

import (
	"context"
	"github.com/rcrowley/go-metrics"
	"datacollector-edge/api/validation"
)

type StageContext interface {
	// If we plan to support ELs later, we should remove and provide in build support for this
	GetResolvedValue(configValue interface{}) (interface{}, error)
	CreateRecord(recordSourceId string, value interface{}) (Record, error)
	CreateEventRecord(recordSourceId string, value interface{}, eventType string, eventVersion int) (Record, error)
	GetMetrics() metrics.Registry
	ToError(err error, record Record)
	ToEvent(record Record)
	ReportError(err error)
	GetOutputLanes() []string
	Evaluate(value string, configName string, ctx context.Context) (interface{}, error)
	IsErrorStage() bool
	CreateConfigIssue(error string, optional ...interface{}) validation.Issue
	GetService(serviceName string) (Service, error)
	IsPreview() bool
	GetPipelineParameters() map[string]interface{}
	SetStop()
	IsStopped() bool
}
