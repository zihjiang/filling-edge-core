// +build 386 windows,amd64 windows



package windows

import (
	"datacollector-edge/api"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"datacollector-edge/container/execution/runner"
	"testing"
)

func createStageContext(logName string) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = Library
	stageConfig.StageName = StageName
	stageConfig.Configuration = make([]common.Config, 2)

	stageConfig.Configuration[0] = common.Config{
		Name:  "logName",
		Value: logName,
	}
	stageConfig.Configuration[1] = common.Config{
		Name:  "readMode",
		Value: "ALL",
	}
	return &common.StageContextImpl{
		StageConfig: &stageConfig,
		Parameters:  nil,
	}
}

func testWindowsEventLogRead(t *testing.T, logName string, maxBatchSize int) {
	stageContext := createStageContext(logName)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
	}
	stageInstance := stageBean.Stage

	issues := stageInstance.Init(stageContext)

	if len(issues) > 0 {
		t.Fatalf("Error when Initializing stage %s", issues[0].Message)
	}

	defer stageInstance.Destroy()
	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)

	_, err = stageInstance.(api.Origin).Produce(nil, maxBatchSize, batchMaker)

	if err != nil {
		t.Fatalf("Error when Producing %s", err.Error())
	}

	records := batchMaker.GetStageOutput()

	if len(records) <= 0 {
		t.Fatalf("Did not read any records")
	} else {
		for _, event := range records {
			rootField, _ := event.Get()
			rootFieldValue := rootField.Value.(map[string](*api.Field))
			actualLogName := rootFieldValue["LogName"].Value
			if actualLogName != logName {
				t.Fatalf("Wrong Log Name. Expected : %s, Actual : %s", logName, actualLogName)
			}
		}
	}
}

func TestWindowsApplicationLogRead(t *testing.T) {
	testWindowsEventLogRead(t, Application, 1)
}

func TestWindowsSystemLogRead(t *testing.T) {
	testWindowsEventLogRead(t, System, 1)
}
