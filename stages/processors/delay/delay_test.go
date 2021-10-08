
package delay

import (
	"datacollector-edge/api"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"datacollector-edge/container/execution/runner"
	"testing"
)

func getStageContext(delay float64) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = LIBRARY
	stageConfig.StageName = STAGE_NAME
	stageConfig.Configuration = []common.Config{
		{
			Name:  "delay",
			Value: delay,
		},
	}
	return &common.StageContextImpl{
		StageConfig: &stageConfig,
		Parameters:  nil,
	}
}

func TestDelayProcessor(t *testing.T) {
	delay := float64(100)
	stageContext := getStageContext(delay)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
	}

	stageInstance := stageBean.Stage.(*DelayProcessor)
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	if stageInstance.Delay != delay {
		t.Error("Failed to inject config value for delay - ", stageInstance.Delay)
		return
	}

	stageInstance.Init(stageContext)
	records := make([]api.Record, 1)
	records[0], _ = stageContext.CreateRecord("1", "TestData")
	batch := runner.NewBatchImpl("random", records, nil)
	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)

	err = stageInstance.Process(batch, batchMaker)
	if err != nil {
		t.Error("Error in Identity Processor")
	}

	outputRecords := batchMaker.GetStageOutput()
	if len(outputRecords) != 1 {
		t.Error("Excepted 1 records but got - ", len(records))
	}

	rootField, _ := records[0].Get()
	if rootField.Value != "TestData" {
		t.Error("Excepted 'TestData' but got - ", rootField.Value)
	}

	stageInstance.Destroy()
}
