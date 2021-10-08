
package identity

import (
	"errors"
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/stagelibrary"
	"math/rand"
)

const (
	Library   = "streamsets-datacollector-dev-lib"
	StageName = "com_streamsets_pipeline_stage_devtest_RandomErrorProcessor"
)

var randomError = errors.New("random error")

type Processor struct {
	*common.BaseStage
}

func init() {
	stagelibrary.SetCreator(Library, StageName, func() api.Stage {
		return &Processor{BaseStage: &common.BaseStage{}}
	})
}

func (p *Processor) Init(stageContext api.StageContext) []validation.Issue {
	return p.BaseStage.Init(stageContext)
}

func (p *Processor) Process(batch api.Batch, batchMaker api.BatchMaker) error {
	for _, record := range batch.GetRecords() {
		if rand.Float32() < 0.5 {
			batchMaker.AddRecord(record)
		} else {
			p.GetStageContext().ToError(randomError, record)
		}
	}

	if rand.Float32() < 0.5 {
		p.GetStageContext().ReportError(randomError)
	}
	return nil
}
