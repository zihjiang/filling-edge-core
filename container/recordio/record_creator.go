
package recordio

import (
	"datacollector-edge/api"
)

type RecordCreator interface {
	CreateRecord(
		context api.StageContext,
		lineText string,
		messageId string,
		headers []*api.Field,
	) (api.Record, error)
}
