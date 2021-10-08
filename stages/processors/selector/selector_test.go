
package selector

import (
	"datacollector-edge/api"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"datacollector-edge/container/execution/runner"
	"testing"
)

func getStageContext() *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = LIBRARY
	stageConfig.StageName = STAGE_NAME

	lane1 := map[string]interface{}{
		"outputLane": "lane1",
		"predicate":  "${record:value('/a') != NULL}",
	}
	lane2 := map[string]interface{}{
		"outputLane": "lane2",
		"predicate":  "default",
	}
	predicateValueList := make([]interface{}, 0)
	predicateValueList = append(predicateValueList, lane1)
	predicateValueList = append(predicateValueList, lane2)
	stageConfig.Configuration = []common.Config{
		{
			Name:  "lanePredicates",
			Value: predicateValueList,
		},
	}
	stageConfig.OutputLanes = []string{
		"lane1",
		"lane2",
	}
	return &common.StageContextImpl{
		StageConfig: &stageConfig,
		Parameters:  nil,
	}
}

func TestHttpServerOrigin_Init(t *testing.T) {
	stageContext := getStageContext()
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
	}

	stageInstance := stageBean.Stage.(*SelectorProcessor)
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	if stageInstance.LanePredicates == nil {
		t.Error("Failed to inject config value for lane predicate")
		return
	}

	if len(stageInstance.LanePredicates) != 2 {
		t.Error("Failed to inject config value for lane predicate")
	}

	if stageInstance.LanePredicates[0]["predicate"] != "${record:value('/a') != NULL}" {
		t.Error("Failed to inject config value for lane predicate")
	}

	if stageInstance.LanePredicates[1]["predicate"] != "default" {
		t.Error("Failed to inject config value for lane predicate")
	}

}

func TestSelectorProcessor(t *testing.T) {
	stageContext := getStageContext()
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
	}
	stageInstance := stageBean.Stage
	issues := stageInstance.Init(stageContext)
	if len(issues) != 0 {
		t.Error(issues[0].Message)
		return
	}
	records := make([]api.Record, 1)
	records[0], _ = stageContext.CreateRecord("1", map[string]interface{}{"a": "sample"})
	batch := runner.NewBatchImpl("random", records, nil)
	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)

	err = stageInstance.(api.Processor).Process(batch, batchMaker)
	if err != nil {
		t.Error("Error in Identity Processor")
	}

	lane1OutputRecords := batchMaker.GetStageOutput("lane1")
	if len(lane1OutputRecords) != 1 {
		t.Error("Excepted 1 records but got - ", len(lane1OutputRecords))
		return
	}

	recordValue, err := lane1OutputRecords[0].Get("/a")
	if recordValue.Value != "sample" {
		t.Error("Excepted 'sample' but got - ", recordValue.Value)
	}

	lane2OutputRecords := batchMaker.GetStageOutput("lane2")
	if len(lane2OutputRecords) != 0 {
		t.Error("Excepted 0 records but got - ", len(lane2OutputRecords))
		return
	}

	stageInstance.Destroy()
}
