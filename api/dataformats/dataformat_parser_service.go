
package dataformats

import (
	"datacollector-edge/api"
	"io"
)

const (
	DataFormatParserServiceName = "com.streamsets.pipeline.api.service.dataformats.DataFormatParserService"
)

type DataFormatParserService interface {
	GetParser(messageId string, reader io.Reader) (RecordReader, error)
}

type RecordReader interface {
	ReadRecord() (api.Record, error)
	Close() error
}
