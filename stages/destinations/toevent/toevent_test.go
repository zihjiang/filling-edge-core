
package toevent

import (
	"datacollector-edge/api"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"datacollector-edge/container/execution/runner"
	"testing"
)

func getStageContext() (*common.StageContextImpl, *common.ErrorSink, *common.EventSink) {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = Library
	stageConfig.StageName = StageName
	errorSink := common.NewErrorSink()
	eventSink := common.NewEventSink()
	return &common.StageContextImpl{
		StageConfig:       &stageConfig,
		ErrorSink:         errorSink,
		ErrorRecordPolicy: common.ErrorRecordPolicyStage,
		EventSink:         eventSink,
	}, errorSink, eventSink
}

func TestDestination(t *testing.T) {
	stageContext, errorSink, eventSink := getStageContext()
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage.(*Destination)
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	stageInstance.Init(stageContext)

	records := make([]api.Record, 2)
	records[0], _ = stageContext.CreateRecord(
		"abc",
		map[string]interface{}{
			"a": float64(2.55),
			"b": float64(3.55),
			"c": "random",
		},
	)
	records[1], _ = stageContext.CreateRecord(
		"abc",
		map[string]interface{}{
			"a": float64(2.55),
			"b": float64(3.55),
			"c": "random",
		},
	)
	batch := runner.NewBatchImpl("toEvent", records, nil)
	err = stageInstance.Write(batch)
	if err != nil {
		t.Fatal("Error when writing: " + err.Error())
	}

	if errorSink.GetTotalErrorRecords() != 0 {
		t.Fatal("Failed to write records to event sink")
	}

	if len(eventSink.GetStageEvents("")) != 2 {
		t.Fatal("Failed to write records to event sink")
	}

	stageInstance.Destroy()
}
