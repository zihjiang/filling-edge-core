
package textrecord

import (
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"datacollector-edge/api"
	"datacollector-edge/api/dataformats"
	"datacollector-edge/container/recordio"
	"io"
)

type TextWriterFactoryImpl struct {
	TextFieldPath string
}

func (t *TextWriterFactoryImpl) CreateWriter(
	context api.StageContext,
	writer io.Writer,
) (dataformats.RecordWriter, error) {
	var recordWriter dataformats.RecordWriter
	recordWriter = newRecordWriter(context, writer, t.TextFieldPath)
	return recordWriter, nil
}

type TextWriterImpl struct {
	context       api.StageContext
	writer        *bufio.Writer
	textFieldPath string
}

func (textWriter *TextWriterImpl) WriteRecord(r api.Record) error {
	if textFieldValue, err := r.Get(textWriter.textFieldPath); err != nil {
		return err
	} else if _, err = fmt.Fprintln(textWriter.writer, cast.ToString(textFieldValue.Value)); err != nil {
		return err
	}
	return nil
}

func (textWriter *TextWriterImpl) Flush() error {
	return recordio.Flush(textWriter.writer)
}

func (textWriter *TextWriterImpl) Close() error {
	return recordio.Close(textWriter.writer)
}

func newRecordWriter(
	context api.StageContext,
	writer io.Writer,
	textFieldPath string,
) *TextWriterImpl {
	if len(textFieldPath) == 0 {
		textFieldPath = DefaultTextFieldPath
	}
	return &TextWriterImpl{
		context:       context,
		writer:        bufio.NewWriter(writer),
		textFieldPath: textFieldPath,
	}
}
