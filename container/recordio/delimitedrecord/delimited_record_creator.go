
package delimitedrecord

import (
	"datacollector-edge/api"
	"strings"
)

type RecordCreator struct {
	CsvFileFormat      string
	CsvCustomDelimiter string
	CsvRecordType      string
}

func (r *RecordCreator) CreateRecord(
	context api.StageContext,
	lineText string,
	messageId string,
	headers []*api.Field,
) (api.Record, error) {
	sep := ","
	if r.CsvFileFormat == Custom && len(r.CsvCustomDelimiter) > 0 {
		sep = r.CsvCustomDelimiter
	}
	columns := strings.Split(lineText, sep)
	return createRecord(
		context,
		messageId,
		1,
		r.CsvRecordType,
		columns,
		headers,
	)
}
