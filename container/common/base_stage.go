
package common

import (
	"datacollector-edge/api"
	"datacollector-edge/api/dataformats"
	"datacollector-edge/api/validation"
)

type BaseStage struct {
	stageContext api.StageContext
}

func (b *BaseStage) GetStageContext() api.StageContext {
	return b.stageContext
}

func (b *BaseStage) Init(stageContext api.StageContext) []validation.Issue {
	issues := make([]validation.Issue, 0)
	b.stageContext = stageContext
	return issues
}

func (b *BaseStage) Destroy() error {
	// No OP Destroy
	return nil
}

func (b *BaseStage) GetStageConfig() *StageConfiguration {
	return b.stageContext.(*StageContextImpl).StageConfig
}

func (b *BaseStage) GetDataParserService() (dataformats.DataFormatParserService, error) {
	dataParserService, err := b.GetStageContext().GetService(dataformats.DataFormatParserServiceName)
	if err != nil {
		return nil, err
	}
	return dataParserService.(dataformats.DataFormatParserService), err
}

func (b *BaseStage) GetDataGeneratorService() (dataformats.DataFormatGeneratorService, error) {
	dataGeneratorService, err := b.GetStageContext().GetService(dataformats.DataFormatGeneratorServiceName)
	if err != nil {
		return nil, err
	}
	return dataGeneratorService.(dataformats.DataFormatGeneratorService), err
}
