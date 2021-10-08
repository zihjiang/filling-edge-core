
package dev_random

import (
	"datacollector-edge/api"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"datacollector-edge/container/execution/runner"
	"testing"
)

func getStageContext(
	configuration []common.Config,
	parameters map[string]interface{},
) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = Library
	stageConfig.StageName = StageName
	stageConfig.Configuration = configuration
	errorSink := common.NewErrorSink()
	return &common.StageContextImpl{
		StageConfig: &stageConfig,
		Parameters:  parameters,
		ErrorSink:   errorSink,
		EventSink:   common.NewEventSink(),
	}
}

func TestOrigin_Init(t *testing.T) {
	stageContext := getStageContext(getDefaultTestConfigs(), nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	if stageInstance.(*Origin).Delay != float64(1000) {
		t.Error("Failed to inject config value for delay")
	}

	if stageInstance.(*Origin).DataGenConfigs == nil {
		t.Error("Failed to inject config value for DataGenConfigs")
	}
}

func TestOrigin_Produce(t *testing.T) {
	stageContext := getStageContext(getDefaultTestConfigs(), nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	stageInstance.Init(stageContext)
	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(nil, 1, batchMaker)
	if err != nil {
		t.Error("Err :", err)
		return
	}
}

func getDefaultTestConfigs() []common.Config {
	dataGeneratorConfigList := []interface{}{
		map[string]interface{}{
			"field": "stringField",
			"type":  STRING,
		},
		map[string]interface{}{
			"field": "integerField",
			"type":  INTEGER,
		},
		map[string]interface{}{
			"field": "longField",
			"type":  LONG,
		},
		map[string]interface{}{
			"field": "floatField",
			"type":  FLOAT,
		},
		map[string]interface{}{
			"field": "doubleField",
			"type":  DOUBLE,
		},
		map[string]interface{}{
			"field": "boolField",
			"type":  BOOLEAN,
		},
		map[string]interface{}{
			"field": "dateTimeField",
			"type":  DATETIME,
		},
		map[string]interface{}{
			"field": "decimalField",
			"type":  DECIMAL,
		},
	}

	configuration := []common.Config{
		{
			Name:  "delay",
			Value: float64(1000),
		},
		{
			Name:  "dataGenConfigs",
			Value: dataGeneratorConfigList,
		},
	}

	return configuration
}
