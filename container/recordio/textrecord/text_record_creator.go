
package textrecord

import (
	"datacollector-edge/api"
	"datacollector-edge/container/util"
)

const (
	DefaultTextField     = "text"
	DefaultTextFieldPath = "/text"
)

type RecordCreator struct {
	TextMaxLineLen int
}

func (r *RecordCreator) CreateRecord(
	context api.StageContext,
	lineText string,
	messageId string,
	headers []*api.Field,
) (api.Record, error) {
	return context.CreateRecord(messageId, map[string]interface{}{
		DefaultTextField: util.TruncateString(lineText, r.TextMaxLineLen),
	})
}
