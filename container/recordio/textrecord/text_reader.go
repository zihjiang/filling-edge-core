
package textrecord

import (
	"bufio"
	"datacollector-edge/api"
	"datacollector-edge/api/dataformats"
	"datacollector-edge/container/common"
	"datacollector-edge/container/recordio"
	"datacollector-edge/container/util"
	"io"
	"strings"
)

type TextReaderFactoryImpl struct {
	recordio.AbstractRecordReaderFactory
	TextMaxLineLen int
}

func (j *TextReaderFactoryImpl) CreateReader(
	context api.StageContext,
	reader io.Reader,
	messageId string,
) (dataformats.RecordReader, error) {
	var recordReader dataformats.RecordReader
	recordReader = newRecordReader(context, reader, messageId, j.TextMaxLineLen)
	return recordReader, nil
}

type TextReaderImpl struct {
	context        api.StageContext
	reader         *bufio.Reader
	messageId      string
	counter        int
	textMaxLineLen int
}

func (textReader *TextReaderImpl) ReadRecord() (api.Record, error) {
	var err error
	line, err := textReader.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}
	if len(line) > 0 {
		line = util.TruncateString(strings.TrimRight(line, "\r\n"), textReader.textMaxLineLen)
		recordValue := map[string]interface{}{"text": line}
		textReader.counter++
		sourceId := common.CreateRecordId(textReader.messageId, textReader.counter)
		return textReader.context.CreateRecord(sourceId, recordValue)
	}
	return nil, nil
}

func (textReader *TextReaderImpl) Close() error {
	return recordio.Close(textReader.reader)
}

func newRecordReader(context api.StageContext, reader io.Reader, messageId string, textMaxLineLen int) *TextReaderImpl {
	return &TextReaderImpl{
		context:        context,
		reader:         bufio.NewReader(reader),
		messageId:      messageId,
		counter:        0,
		textMaxLineLen: textMaxLineLen,
	}
}
