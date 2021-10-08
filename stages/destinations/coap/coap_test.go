
package coap

import (
	"datacollector-edge/api"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"datacollector-edge/container/execution/runner"
	"testing"
)

func getStageContext(resourceUrl string, coapMethod string, messageType string) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = LIBRARY
	stageConfig.StageName = STAGE_NAME
	stageConfig.Configuration = make([]common.Config, 3)
	stageConfig.Configuration[0] = common.Config{
		Name:  CONF_RESOURCE_URL,
		Value: resourceUrl,
	}
	stageConfig.Configuration[1] = common.Config{
		Name:  CONF_COAP_METHOD,
		Value: coapMethod,
	}
	stageConfig.Configuration[2] = common.Config{
		Name:  CONF_RESOURCE_TYPE,
		Value: messageType,
	}

	return &common.StageContextImpl{
		StageConfig: &stageConfig,
		Parameters:  nil,
	}
}

func TestConfirmableMessage(t *testing.T) {
	stageContext := getStageContext("coap://localhost:56831/sdc", POST, CONFIRMABLE)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
	}
	stageInstance := stageBean.Stage

	if stageInstance.(*CoapClientDestination).Conf.ResourceUrl != "coap://localhost:56831/sdc" {
		t.Error("Failed to inject config value for ResourceUrl")
	}

	if stageInstance.(*CoapClientDestination).Conf.CoapMethod != POST {
		t.Error("Failed to inject config value for CoapMethod")
	}

	if stageInstance.(*CoapClientDestination).Conf.RequestType != CONFIRMABLE {
		t.Error("Failed to inject config value for RequestType")
	}

	stageInstance.Init(stageContext)
	records := make([]api.Record, 1)
	records[0], _ = stageContext.CreateRecord("1", "TestData")
	batch := runner.NewBatchImpl("random", records, nil)
	err = stageInstance.(api.Destination).Write(batch)
	if err == nil {
		t.Error("Excepted error message for invalid CoAP URL with confirmable message")
	}
	stageInstance.Destroy()
}

func TestNonConfirmableMessage(t *testing.T) {
	stageContext := getStageContext("coap://localhost:45/sdc", POST, NONCONFIRMABLE)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
	}
	stageInstance := stageBean.Stage
	records := make([]api.Record, 1)
	records[0], _ = stageContext.CreateRecord("1", "test data")
	batch := runner.NewBatchImpl("random", records, nil)

	stageInstance.Init(stageContext)
	err = stageInstance.(api.Destination).Write(batch)
	if err != nil {
		t.Error("Not excepted error message for invalid CoAP URL with non confirmable message")
	}
	stageInstance.Destroy()
}
