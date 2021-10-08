
package recordio

import (
	"errors"
	"datacollector-edge/api"
	"datacollector-edge/api/dataformats"
	"io"
)

const (
	CompressedNone = "NONE"
	CompressedFile = "COMPRESSED_FILE"
)

type RecordReaderFactory interface {
	CreateReader(context api.StageContext, reader io.Reader, messageId string) (dataformats.RecordReader, error)
	CreateWholeFileReader(
		context api.StageContext,
		messageId string,
		metadata map[string]interface{},
		fileRef api.FileRef,
	) (dataformats.RecordReader, error)
}

type AbstractRecordReaderFactory struct {
}

func (*AbstractRecordReaderFactory) CreateWholeFileReader(
	context api.StageContext,
	messageId string,
	metadata map[string]interface{},
	fileRef api.FileRef,
) (dataformats.RecordReader, error) {
	return nil, errors.New("not supported operation")
}

func (*AbstractRecordReaderFactory) CreateReader(
	context api.StageContext,
	reader io.Reader,
	messageId string,
) (dataformats.RecordReader, error) {
	return nil, errors.New("not supported operation")
}

type RecordWriterFactory interface {
	CreateWriter(context api.StageContext, writer io.Writer) (dataformats.RecordWriter, error)
}

type Flusher interface {
	Flush() error
}

func Flush(v interface{}) error {
	c, ok := v.(Flusher)
	if ok {
		return c.Flush()
	}
	return nil
}

func Close(v interface{}) error {
	c, ok := v.(io.Closer)
	if ok {
		return c.Close()
	}
	return nil
}
