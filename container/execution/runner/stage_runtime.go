
package runner

import (
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
)

type StageRuntime struct {
	pipelineBean creation.PipelineBean
	config       *common.StageConfiguration
	stageBean    creation.StageBean
	stageContext api.StageContext
}

func (s *StageRuntime) Init() []validation.Issue {
	issues := make([]validation.Issue, 0)
	if s.stageBean.Services != nil {
		for _, serviceBean := range s.stageBean.Services {
			serviceIssues := serviceBean.Service.Init(s.stageContext)
			issues = append(issues, serviceIssues...)
		}
	}
	stageIssues := s.stageBean.Stage.Init(s.stageContext)
	return append(issues, stageIssues...)
}

func (s *StageRuntime) Execute(
	previousOffset *string,
	batchSize int,
	batch *BatchImpl,
	batchMaker *BatchMakerImpl,
) (*string, error) {
	var newOffset *string
	var err error
	if s.stageBean.IsSource() {
		newOffset, err = s.stageBean.Stage.(api.Origin).Produce(previousOffset, batchSize, batchMaker)
	} else if s.stageBean.IsProcessor() {
		err = s.stageBean.Stage.(api.Processor).Process(batch, batchMaker)
	} else if s.stageBean.IsTarget() {
		err = s.stageBean.Stage.(api.Destination).Write(batch)
	}
	return newOffset, err
}

func (s *StageRuntime) Destroy() {
	if s.stageBean.Services != nil {
		for _, serviceBean := range s.stageBean.Services {
			_ = serviceBean.Service.Destroy()
		}
	}
	_ = s.stageBean.Stage.Destroy()
}

func (s *StageRuntime) GetInstanceName() string {
	return s.config.InstanceName
}

func NewStageRuntime(
	pipelineBean creation.PipelineBean,
	stageBean creation.StageBean,
	stageContext api.StageContext,
) StageRuntime {
	return StageRuntime{
		pipelineBean: pipelineBean,
		config:       stageBean.Config,
		stageBean:    stageBean,
		stageContext: stageContext,
	}
}
