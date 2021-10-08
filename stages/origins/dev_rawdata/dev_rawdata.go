
package dev_random

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/stagelibrary"
)

const (
	Library   = "streamsets-datacollector-dev-lib"
	StageName = "com_streamsets_pipeline_stage_devtest_rawdata_RawDataDSource"
)

var randomOffset = "random"

type DevRawDataDSource struct {
	*common.BaseStage
	RawData             string `ConfigDef:"type=STRING,required=true"`
	StopAfterFirstBatch bool   `ConfigDef:"type=BOOLEAN,required=true"`
}

func init() {
	stagelibrary.SetCreator(Library, StageName, func() api.Stage {
		return &DevRawDataDSource{BaseStage: &common.BaseStage{}}
	})
}

func (d *DevRawDataDSource) Init(stageContext api.StageContext) []validation.Issue {
	issues := d.BaseStage.Init(stageContext)
	log.Debug("DevRawDataDSource Init method")
	return issues
}

func (d *DevRawDataDSource) Produce(
	lastSourceOffset *string,
	maxBatchSize int,
	batchMaker api.BatchMaker,
) (*string, error) {

	dataParserService, err := d.GetDataParserService()
	if err != nil {
		log.WithError(err).Error("Failed to get DataParserService")
		return nil, err
	}
	recordReader, err := dataParserService.GetParser("rawData", bytes.NewBufferString(d.RawData))
	if err != nil {
		log.WithError(err).Error("Failed to create record reader")
		return nil, err
	}

	defer recordReader.Close()
	for {
		record, err := recordReader.ReadRecord()
		if err != nil {
			log.WithError(err).Error("Failed to parse raw data")
			d.GetStageContext().ReportError(err)
			return nil, nil
		}

		if record == nil {
			break
		}
		batchMaker.AddRecord(record)
	}

	if d.StopAfterFirstBatch {
		return nil, nil
	}

	return &randomOffset, nil
}
