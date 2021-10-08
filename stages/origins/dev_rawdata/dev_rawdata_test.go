
package dev_random

import (
	"datacollector-edge/api"
	"datacollector-edge/api/dataformats"
	"datacollector-edge/api/fieldtype"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"datacollector-edge/container/execution/runner"
	"datacollector-edge/stages/services"
	"testing"
)

const (
	ConfRawData    = "rawData"
	ConfDataFormat = "dataFormat"
)

func getStageContext(
	rawData string,
	dataFormat string,
	parameters map[string]interface{},
) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = Library
	stageConfig.StageName = StageName
	stageConfig.Configuration = []common.Config{
		{
			Name:  ConfRawData,
			Value: rawData,
		},
	}

	serviceConfig := &common.ServiceConfiguration{}
	serviceConfig.Service = dataformats.DataFormatParserServiceName
	serviceConfig.Configuration = []common.Config{
		{
			Name:  ConfDataFormat,
			Value: dataFormat,
		},
	}
	stageConfig.Services = []*common.ServiceConfiguration{serviceConfig}

	return &common.StageContextImpl{
		StageConfig: &stageConfig,
		Parameters:  parameters,
	}
}

func TestDevRawDataDSource_Init(t *testing.T) {
	rawData := "text1\n text2"
	stageContext := getStageContext(rawData, "TEXT", nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
	}
	stageInstance := stageBean.Stage

	if stageInstance.(*DevRawDataDSource).RawData != rawData {
		t.Error("Failed to inject config value for rawData")
	}
}

func TestDevRandomOrigin_TextFormat(t *testing.T) {
	rawData := "text1\n text2"
	stageContext := getStageContext(rawData, "TEXT", nil)

	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
	}
	stageInstance := stageBean.Stage
	serviceInstance := stageBean.Services[0].Service

	stageContext.Services = map[string]api.Service{
		services.GetDataFormatParserServiceName(): serviceInstance,
	}

	// initialize service instance
	issues := serviceInstance.Init(stageContext)
	if len(issues) != 0 {
		t.Error(issues[0].Message)
	}

	issues = stageInstance.Init(stageContext)
	if len(issues) != 0 {
		t.Error(issues[0].Message)
	}

	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(&randomOffset, 5, batchMaker)
	if err != nil {
		t.Error("Err :", err)
	}

	records := batchMaker.GetStageOutput()
	if len(records) != 2 {
		t.Error("Excepted 2 records but got - ", len(records))
	}

	for _, record := range records {
		rootField, _ := record.Get()
		if rootField.Type != fieldtype.MAP {
			t.Error("Excepted Map field type but got - ", rootField.Type, " Value: ", rootField.Value)
			return
		}
	}
	stageInstance.Destroy()
}

func TestDevRandomOrigin_JsonFormat(t *testing.T) {
	rawData := "{\n  \"f1\": \"abc\"\n}\n{\n  \"f1\": \"xyz\"\n}"
	stageContext := getStageContext(rawData, "JSON", nil)

	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
	}
	stageInstance := stageBean.Stage
	serviceInstance := stageBean.Services[0].Service

	stageContext.Services = map[string]api.Service{
		services.GetDataFormatParserServiceName(): serviceInstance,
	}

	// initialize service instance
	issues := serviceInstance.Init(stageContext)
	if len(issues) != 0 {
		t.Error(issues[0].Message)
	}

	issues = stageInstance.Init(stageContext)
	if len(issues) != 0 {
		t.Error(issues[0].Message)
	}

	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(&randomOffset, 5, batchMaker)
	if err != nil {
		t.Error("Err :", err)
	}

	records := batchMaker.GetStageOutput()
	if len(records) != 2 {
		t.Error("Excepted 2 records but got - ", len(records))
		return
	}

	record1 := records[0]
	rootField, _ := record1.Get()
	if rootField.Type != fieldtype.MAP {
		t.Error("Excepted Map field type but got - ", rootField.Type, " Value: ", rootField.Value)
		return
	}
	rootFieldValue := rootField.Value.(map[string]*api.Field)
	for key, field := range rootFieldValue {
		if field.Type != fieldtype.STRING {
			t.Error("Excepted String field type but got - ", field.Type, " Value: ", field.Value)
			return
		}
		if key != "f1" {
			t.Error("Invalid key", key)
		}
		if field.Value != "abc" {
			t.Error("Invalid value", field.Value)
		}
	}

	stageInstance.Destroy()
}
