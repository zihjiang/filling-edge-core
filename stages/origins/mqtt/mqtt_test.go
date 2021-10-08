
package mqtt

import (
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"testing"
)

func getStageContext(
	brokerUrl string,
	clientId string,
	qos string,
	topicFilters []string,
	dataFormat string,
	parameters map[string]interface{},
) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = Library
	stageConfig.StageName = StageName
	stageConfig.Configuration = []common.Config{
		{
			Name:  "commonConf.brokerUrl",
			Value: brokerUrl,
		},
		{
			Name:  "commonConf.clientId",
			Value: clientId,
		},
		{
			Name:  "commonConf.qos",
			Value: qos,
		},
		{
			Name:  "subscriberConf.topicFilters",
			Value: topicFilters,
		},
		{
			Name:  "subscriberConf.dataFormat",
			Value: dataFormat,
		},
	}
	return &common.StageContextImpl{
		StageConfig: &stageConfig,
		Parameters:  parameters,
	}
}

func TestMqttClientSource_Init(t *testing.T) {
	brokerUrl := "http://test:9000"
	clientId := "clientId"
	qos := "AT_LEAST_ONCE"
	topicFilters := []string{"Sample/Topic"}
	dataFormat := "JSON"

	stageContext := getStageContext(brokerUrl, clientId, qos, topicFilters, dataFormat, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	if stageInstance.(*Origin).CommonConf.BrokerUrl != brokerUrl {
		t.Error("Failed to inject config value for brokerUrl")
	}

	if stageInstance.(*Origin).CommonConf.ClientId != clientId {
		t.Error("Failed to inject config value for clientId")
	}

	if stageInstance.(*Origin).CommonConf.Qos != qos {
		t.Error("Failed to inject config value for qos")
	}

	if len(stageInstance.(*Origin).SubscriberConf.TopicFilters) != len(topicFilters) {
		t.Error("Failed to inject config value for topicFilters")
		return
	}

	if stageInstance.(*Origin).SubscriberConf.TopicFilters[0] != topicFilters[0] {
		t.Error("Failed to inject config value for topicFilters")
		return
	}

	if stageInstance.(*Origin).SubscriberConf.DataFormat != dataFormat {
		t.Error("Failed to inject config value for dataFormat")
	}
}
