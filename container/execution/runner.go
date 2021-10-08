
package execution

import (
	"github.com/rcrowley/go-metrics"
	"datacollector-edge/api"
	"datacollector-edge/container/common"
)

type Runner interface {
	GetPipelineConfig() common.PipelineConfiguration
	GetStatus() (*common.PipelineState, error)
	GetHistory() ([]*common.PipelineState, error)
	GetMetrics() (metrics.Registry, error)
	StartPipeline(runtimeParameters map[string]interface{}) (*common.PipelineState, error)
	StopPipeline() (*common.PipelineState, error)
	ResetOffset() error
	CommitOffset(sourceOffset common.SourceOffset) error
	GetOffset() (common.SourceOffset, error)
	IsRemotePipeline() bool
	GetErrorRecords(stageInstanceName string, size int) ([]api.Record, error)
	GetErrorMessages(stageInstanceName string, size int) ([]api.ErrorMessage, error)
}
