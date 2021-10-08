
package websocketclient

import (
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
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
	}
}

func TestHttpClientOrigin_Init(t *testing.T) {
	resourceUrl := "ws://localhost:18630/rest/v1/webSocket?type=status"
	headers := make([]interface{}, 2)
	headers[0] = map[string]interface{}{
		"key":   "X-SDC-APPLICATION-ID",
		"value": "SDC Edge",
	}
	headers[1] = map[string]interface{}{
		"key":   "DUMMY-HEADER",
		"value": "DUMMY",
	}

	configuration := []common.Config{
		{
			Name:  "conf.resourceUrl",
			Value: resourceUrl,
		},
		{
			Name:  "conf.headers",
			Value: headers,
		},
		{
			Name:  "conf.dataFormat",
			Value: "JSON",
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	if stageInstance.(*Origin).Conf.ResourceUrl != resourceUrl {
		t.Error("Failed to inject config value for resourceUrl")
	}

	if stageInstance.(*Origin).Conf.Headers == nil {
		t.Error("Failed to inject config value for Headers")
		return
	}
}

func TestHttpClientOrigin_InvalidResourceURL(t *testing.T) {
	configuration := []common.Config{
		{
			Name:  "conf.resourceUrl",
			Value: "dfsd",
		},
		{
			Name:  "conf.dataFormat",
			Value: "JSON",
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	issues := stageInstance.Init(stageContext)
	if len(issues) == 0 {
		t.Error("Expected error related to invalid URL")
		return
	}

}
