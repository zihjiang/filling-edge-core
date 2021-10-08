
package identity

import (
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/stagelibrary"
)

const (
	LIBRARY    = "streamsets-datacollector-dev-lib"
	STAGE_NAME = "com_streamsets_pipeline_stage_processor_identity_IdentityProcessor"
	VERSION    = 1
)

type IdentityProcessor struct {
	*common.BaseStage
}

func init() {
	stagelibrary.SetCreator(LIBRARY, STAGE_NAME, func() api.Stage {
		return &IdentityProcessor{BaseStage: &common.BaseStage{}}
	})
}

func (i *IdentityProcessor) Init(stageContext api.StageContext) []validation.Issue {
	return i.BaseStage.Init(stageContext)
}

func (i *IdentityProcessor) Process(batch api.Batch, batchMaker api.BatchMaker) error {
	for _, record := range batch.GetRecords() {
		batchMaker.AddRecord(record)
	}
	return nil
}
