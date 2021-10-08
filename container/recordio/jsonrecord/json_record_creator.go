
package jsonrecord

import (
	"bytes"
	"encoding/json"
	"datacollector-edge/api"
)

type RecordCreator struct {
}

func (r *RecordCreator) CreateRecord(
	context api.StageContext,
	lineText string,
	messageId string,
	headers []*api.Field,
) (api.Record, error) {
	recordBuffer := bytes.NewBufferString(lineText)
	decoder := json.NewDecoder(recordBuffer)
	var recordValue interface{}
	err := decoder.Decode(&recordValue)
	if err != nil {
		return nil, err
	}
	return context.CreateRecord(messageId, recordValue)
}
