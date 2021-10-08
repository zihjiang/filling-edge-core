
package dev_random

import (
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/stagelibrary"
	"math/rand"
	"strings"
	"time"
)

const (
	LIBRARY                  = "streamsets-datacollector-dev-lib"
	STAGE_NAME               = "com_streamsets_pipeline_stage_devtest_RandomSource"
	ConfFields               = "fields"
	ConfDelay                = "delay"
	ConfMaxRecordsToGenerate = "maxRecordsToGenerate"
)

var randomOffset = "random"

type DevRandom struct {
	*common.BaseStage
	Fields               string  `ConfigDef:"type=STRING,required=true"`
	Delay                float64 `ConfigDef:"type=NUMBER,required=true"`
	MaxRecordsToGenerate float64 `ConfigDef:"type=NUMBER,required=true"`
	fieldsList           []string
	recordsProduced      float64
}

func init() {
	stagelibrary.SetCreator(LIBRARY, STAGE_NAME, func() api.Stage {
		return &DevRandom{BaseStage: &common.BaseStage{}}
	})
}

func (d *DevRandom) Init(stageContext api.StageContext) []validation.Issue {
	issues := d.BaseStage.Init(stageContext)
	d.fieldsList = strings.Split(d.Fields, ",")
	d.recordsProduced = 0
	return issues
}

func (d *DevRandom) Produce(lastSourceOffset *string, maxBatchSize int, batchMaker api.BatchMaker) (*string, error) {
	r := rand.New(rand.NewSource(99))
	time.Sleep(time.Duration(d.Delay) * time.Millisecond)
	for i := 0; i < maxBatchSize && d.recordsProduced < d.MaxRecordsToGenerate; i++ {
		var recordValue = make(map[string]interface{})
		for _, field := range d.fieldsList {
			recordValue[field] = r.Int63()
		}
		recordId := common.CreateRecordId("dev-random", i)
		if record, err := d.GetStageContext().CreateRecord(recordId, recordValue); err == nil {
			batchMaker.AddRecord(record)
		} else {
			d.GetStageContext().ToError(err, record)
		}
		d.recordsProduced++
	}
	return &randomOffset, nil
}
