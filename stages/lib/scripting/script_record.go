
package scripting

import "datacollector-edge/api"

func NewScriptRecord(record api.Record, scriptObject interface{}) (map[string]interface{}, error) {
	var err error
	scriptRecord := map[string]interface{}{
		"record":               record,
		"value":                scriptObject,
		"stageCreator":         record.GetHeader().GetStageCreator(),
		"sourceId":             record.GetHeader().GetSourceId(),
		"previousTrackingId":   record.GetHeader().GetPreviousTrackingId(),
		"attributes":           make(map[string]string),
		"errorDataCollectorId": record.GetHeader().GetErrorDataCollectorId(),
		"errorPipelineName":    record.GetHeader().GetErrorPipelineName(),
		"errorCode":            record.GetHeader().GetErrorMessage(),
		"errorMessage":         record.GetHeader().GetErrorMessage(),
		"errorStage":           record.GetHeader().GetErrorStage(),
		"errorTimestamp":       record.GetHeader().GetErrorTimestamp(),
		"errorStackTrace":      record.GetHeader().GetErrorMessage(),
	}

	attributes := scriptRecord["attributes"].(map[string]string)
	for _, key := range record.GetHeader().GetAttributeNames() {
		attributes[key] = record.GetHeader().GetAttribute(key).(string)
	}

	return scriptRecord, err
}
