
package toerror

import (
	"errors"
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/stagelibrary"
)

const (
	Library   = "streamsets-datacollector-basic-lib"
	StageName = "com_streamsets_pipeline_stage_destination_toerror_ToErrorDTarget"
)

type Destination struct {
	*common.BaseStage
}

func init() {
	stagelibrary.SetCreator(Library, StageName, func() api.Stage {
		return &Destination{BaseStage: &common.BaseStage{}}
	})
}

func (d *Destination) Init(stageContext api.StageContext) []validation.Issue {
	return d.BaseStage.Init(stageContext)
}

func (d *Destination) Write(batch api.Batch) error {
	for _, record := range batch.GetRecords() {
		d.GetStageContext().ToError(errors.New("error target"), record)
	}
	return nil
}
