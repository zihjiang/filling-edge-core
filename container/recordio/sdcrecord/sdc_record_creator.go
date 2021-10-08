
package sdcrecord

import (
	"errors"
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
	return nil, errors.New("not supported")
}
