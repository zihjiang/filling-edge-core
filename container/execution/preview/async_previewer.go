
package preview

import (
	"github.com/satori/go.uuid"
	"datacollector-edge/container/execution"
	pipelineStore "datacollector-edge/container/store"
)

type AsyncPreviewer struct {
	syncPreviewer SyncPreviewer
}

func (p *AsyncPreviewer) GetId() string {
	return p.syncPreviewer.GetId()
}

func (p *AsyncPreviewer) ValidateConfigs(timeoutMillis int64) error {
	return nil
}

func (p *AsyncPreviewer) Start(
	batches int,
	batchSize int,
	skipTargets bool,
	stopStage string,
	stagesOverride []execution.StageOutputJson,
	timeoutMillis int64,
	testOrigin bool,
) error {
	go p.syncPreviewer.Start(batches, batchSize, skipTargets, stopStage, stagesOverride, timeoutMillis, testOrigin)
	return nil
}

func (p *AsyncPreviewer) Stop() error {
	return p.syncPreviewer.Stop()
}

func (p *AsyncPreviewer) GetStatus() string {
	return p.syncPreviewer.GetStatus()
}

func (p *AsyncPreviewer) GetOutput() execution.PreviewOutput {
	return p.syncPreviewer.GetOutput()
}

func NewAsyncPreviewer(
	pipelineId string,
	config execution.Config,
	pipelineStoreTask pipelineStore.PipelineStoreTask,
) (execution.Previewer, error) {
	syncPreviewer := &SyncPreviewer{
		pipelineId:        pipelineId,
		previewerId:       uuid.NewV4().String(),
		config:            config,
		pipelineStoreTask: pipelineStoreTask,
		previewOutput:     execution.PreviewOutput{},
	}
	return syncPreviewer, nil
}
