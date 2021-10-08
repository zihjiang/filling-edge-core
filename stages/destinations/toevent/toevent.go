
package toevent

import (
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/stagelibrary"
)

const (
	Library   = "streamsets-datacollector-dev-lib"
	StageName = "com_streamsets_pipeline_stage_destination_ToEventTarget"
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
	counter := 1
	for _, record := range batch.GetRecords() {
		rootField, _ := record.Get()
		recordId := common.CreateRecordId("event-target", counter)
		if eventRecord, err := d.GetStageContext().CreateEventRecord(
			recordId,
			nil,
			"event-target",
			1,
		); err == nil {
			eventRecord.Set(rootField)
			d.GetStageContext().ToEvent(eventRecord)
		} else {
			d.GetStageContext().ToError(err, eventRecord)
		}
		counter++
	}
	return nil
}
