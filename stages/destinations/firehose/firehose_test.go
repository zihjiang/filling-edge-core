// +build aws


package firehose

import (
	"fmt"
	"datacollector-edge/api"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"datacollector-edge/container/execution/runner"
	"testing"
)

func getStageContext(
	stageConfigurationList []common.Config,
	parameters map[string]interface{},
) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = Library
	stageConfig.StageName = StageName
	stageConfig.Configuration = stageConfigurationList
	errorSink := common.NewErrorSink()
	return &common.StageContextImpl{
		StageConfig:       &stageConfig,
		Parameters:        parameters,
		ErrorSink:         errorSink,
		ErrorRecordPolicy: common.ErrorRecordPolicyStage,
	}
}

func getTestConfig(
	awsAccessKeyId string,
	awsSecretAccessKey string,
	streamName string,
) []common.Config {
	configuration := []common.Config{
		{
			Name:  "kinesisConfig.awsConfig.awsAccessKeyId",
			Value: awsAccessKeyId,
		},
		{
			Name:  "kinesisConfig.awsConfig.awsSecretAccessKey",
			Value: awsSecretAccessKey,
		},
		{
			Name:  "kinesisConfig.region",
			Value: "US_WEST_2",
		},
		{
			Name:  "kinesisConfig.streamName",
			Value: streamName,
		},
		{
			Name:  "kinesisConfig.dataFormat",
			Value: "JSON",
		},
	}

	return configuration
}

func TestDestination_Init(t *testing.T) {
	config := getTestConfig(
		"awsAccessKeyId",
		"awsSecretAccessKey",
		"sampleStreamName",
	)
	stageContext := getStageContext(config, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	if stageInstance.(*Destination).KinesisConfig.Region != "US_WEST_2" {
		t.Error("Failed to inject config value for Region")
	}

	if stageInstance.(*Destination).KinesisConfig.StreamName != "sampleStreamName" {
		t.Error("Failed to inject config value for Stream Name")
	}
}

func TestDestination_Write_PreserveOrdering(t *testing.T) {
	config := getTestConfig(
		"invalidAccessKeyId",
		"invalid",
		"invalid",
	)
	stageContext := getStageContext(config, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		t.Error(issues)
		return
	}

	records := make([]api.Record, 1)
	records[0], _ = stageContext.CreateRecord("1", map[string]interface{}{
		"index": "test data",
	})

	batch := runner.NewBatchImpl("random", records, nil)
	err = stageInstance.(api.Destination).Write(batch)

	if len(stageContext.ErrorSink.GetErrorRecords()[""]) != 1 {
		t.Errorf("Expected 1 error recors with an invalid AWS credential, but got")
		return
	}
}

func TestDestination_Write_NoPreserveOrdering(t *testing.T) {
	config := getTestConfig(
		"invalidAccessKeyId",
		"invalid",
		"invalid",
	)
	stageContext := getStageContext(config, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		t.Error(issues)
		return
	}

	records := make([]api.Record, 1)
	records[0], _ = stageContext.CreateRecord("1", map[string]interface{}{
		"index": "test data",
	})

	batch := runner.NewBatchImpl("random", records, nil)
	err = stageInstance.(api.Destination).Write(batch)

	if len(stageContext.ErrorSink.GetErrorRecords()[""]) != 1 {
		t.Errorf("Expected 1 error recors with an invalid AWS credential, but got")
		return
	}
}

func _TestDestination_WriteUsingTestAccount(t *testing.T) {
	config := getTestConfig(
		"sdfsd",
		"ds",
		"dfs",
	)
	stageContext := getStageContext(config, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		t.Error(issues)
		return
	}

	records := make([]api.Record, 10)

	for i := 0; i < 10; i++ {
		records[i], _ = stageContext.CreateRecord("1", map[string]interface{}{
			"index": fmt.Sprintf("test data %d", i),
		})
	}

	batch := runner.NewBatchImpl("random", records, nil)
	err = stageInstance.(api.Destination).Write(batch)

	if stageContext.ErrorSink.GetTotalErrorMessages() != 0 {
		t.Errorf(
			"Expected no stage errors, but encountered error: %s",
			stageContext.ErrorSink.GetStageErrorMessages("")[0].LocalizableMessage,
		)
		return
	}
}
