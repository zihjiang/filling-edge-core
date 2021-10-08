
package dataformats

import (
	"datacollector-edge/api"
	"io"
)

const (
	DataFormatGeneratorServiceName = "com.streamsets.pipeline.api.service.dataformats.DataFormatGeneratorService"
)

type DataFormatGeneratorService interface {
	GetGenerator(writer io.Writer) (RecordWriter, error)
	IsWholeFileFormat() bool
	GetWholeFileName(record api.Record) (string, error)
	GetWholeFileExistsAction() string
	GetIncludeChecksumInTheEvents() bool
	GetChecksumAlgorithm() string
}

type RecordWriter interface {
	WriteRecord(r api.Record) error
	Flush() error
	Close() error
}
