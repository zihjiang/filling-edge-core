
package manager

import (
	"errors"
	"fmt"
	"datacollector-edge/container/common"
	"datacollector-edge/container/execution"
	"datacollector-edge/container/execution/preview"
	"datacollector-edge/container/execution/runner"
	"datacollector-edge/container/store"
)

type PipelineManager struct {
	config            execution.Config
	runnerMap         map[string]execution.Runner
	previewerMap      map[string]execution.Previewer
	runtimeInfo       *common.RuntimeInfo
	pipelineStoreTask store.PipelineStoreTask
}

func (p *PipelineManager) CreatePreviewer(pipelineId string) (execution.Previewer, error) {
	previewer, err := preview.NewAsyncPreviewer(pipelineId, p.config, p.pipelineStoreTask)
	if err != nil {
		return nil, err
	}
	p.previewerMap[previewer.GetId()] = previewer
	return previewer, nil
}

func (p *PipelineManager) GetPreviewer(previewerId string) (execution.Previewer, error) {
	if p.previewerMap[previewerId] == nil {
		return nil, errors.New(fmt.Sprintf("Cannot find the previewer in cache for id: %s", previewerId))
	}
	return p.previewerMap[previewerId], nil
}

func (p *PipelineManager) GetRunner(pipelineId string) execution.Runner {
	if p.runnerMap[pipelineId] == nil {
		pRunner, err := runner.NewEdgeRunner(pipelineId, p.config, p.runtimeInfo, p.pipelineStoreTask)
		if err != nil {
			panic(err)
		}
		p.runnerMap[pipelineId] = pRunner
	}
	return p.runnerMap[pipelineId]
}

func (p *PipelineManager) StartPipeline(
	pipelineId string,
	runtimeParameters map[string]interface{},
) (*common.PipelineState, error) {
	return p.GetRunner(pipelineId).StartPipeline(runtimeParameters)
}

func (p *PipelineManager) StopPipeline(pipelineId string) (*common.PipelineState, error) {
	return p.GetRunner(pipelineId).StopPipeline()
}

func (p *PipelineManager) ResetOffset(pipelineId string) error {
	return p.GetRunner(pipelineId).ResetOffset()
}

func NewManager(
	config execution.Config,
	runtimeInfo *common.RuntimeInfo,
	pipelineStoreTask store.PipelineStoreTask,
) (Manager, error) {
	pipelineManager := PipelineManager{
		config:            config,
		runnerMap:         make(map[string]execution.Runner),
		previewerMap:      make(map[string]execution.Previewer),
		runtimeInfo:       runtimeInfo,
		pipelineStoreTask: pipelineStoreTask,
	}
	return &pipelineManager, nil
}
