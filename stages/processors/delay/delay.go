
package delay

import (
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/stagelibrary"
	"time"
)

const (
	LIBRARY    = "streamsets-datacollector-basic-lib"
	STAGE_NAME = "com_streamsets_pipeline_stage_processor_delay_DelayProcessor"
	VERSION    = 1
)

type DelayProcessor struct {
	*common.BaseStage
	Delay float64 `ConfigDef:"type=NUMBER,required=true"`
}

func init() {
	stagelibrary.SetCreator(LIBRARY, STAGE_NAME, func() api.Stage {
		return &DelayProcessor{BaseStage: &common.BaseStage{}}
	})
}

func (d *DelayProcessor) Init(stageContext api.StageContext) []validation.Issue {
	return d.BaseStage.Init(stageContext)
}

func (d *DelayProcessor) Process(batch api.Batch, batchMaker api.BatchMaker) error {
	time.Sleep(time.Duration(d.Delay) * time.Millisecond)
	for _, record := range batch.GetRecords() {
		batchMaker.AddRecord(record)
	}
	return nil
}
