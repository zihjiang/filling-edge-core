
package trash

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/stagelibrary"
)

const (
	LIBRARY                       = "streamsets-datacollector-basic-lib"
	ERROR_STAGE_NAME              = "com_streamsets_pipeline_stage_destination_devnull_ToErrorNullDTarget"
	NULL_STAGE_NAME               = "com_streamsets_pipeline_stage_destination_devnull_NullDTarget"
	STATS_NULL_STAGE_NAME         = "com_streamsets_pipeline_stage_destination_devnull_StatsNullDTarget"
	STATS_DPM_DIRECTLY_STAGE_NAME = "com_streamsets_pipeline_stage_destination_devnull_StatsDpmDirectlyDTarget"
)

type TrashDestination struct {
	*common.BaseStage
}

func init() {
	stagelibrary.SetCreator(LIBRARY, ERROR_STAGE_NAME, func() api.Stage {
		return &TrashDestination{BaseStage: &common.BaseStage{}}
	})
	stagelibrary.SetCreator(LIBRARY, NULL_STAGE_NAME, func() api.Stage {
		return &TrashDestination{BaseStage: &common.BaseStage{}}
	})
	stagelibrary.SetCreator(LIBRARY, STATS_NULL_STAGE_NAME, func() api.Stage {
		return &TrashDestination{BaseStage: &common.BaseStage{}}
	})
	stagelibrary.SetCreator(LIBRARY, STATS_DPM_DIRECTLY_STAGE_NAME, func() api.Stage {
		return &TrashDestination{BaseStage: &common.BaseStage{}}
	})
}

func (t *TrashDestination) Init(stageContext api.StageContext) []validation.Issue {
	return t.BaseStage.Init(stageContext)
}

func (t *TrashDestination) Write(batch api.Batch) error {
	for _, record := range batch.GetRecords() {
		recordValue, _ := record.Get()
		jsonValue, err := json.Marshal(recordValue.Value)
		if err != nil {
			log.WithError(err).Error("Json Serialization Error")
			t.GetStageContext().ToError(err, record)
		}
		log.WithField("record", string(jsonValue)).Debug("Trashed record")
	}
	return nil
}
