
package manager

import (
	"datacollector-edge/container/common"
	"datacollector-edge/container/execution"
)

type Manager interface {
	CreatePreviewer(pipelineId string) (execution.Previewer, error)
	GetPreviewer(previewerId string) (execution.Previewer, error)
	GetRunner(pipelineId string) execution.Runner
	StartPipeline(
		pipelineId string,
		runtimeParameters map[string]interface{},
	) (*common.PipelineState, error)
	StopPipeline(pipelineId string) (*common.PipelineState, error)
	ResetOffset(pipelineId string) error
}
