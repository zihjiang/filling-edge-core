
package jsonrecord

import (
	"encoding/json"
	"datacollector-edge/api"
	"datacollector-edge/api/dataformats"
	"datacollector-edge/container/common"
	"datacollector-edge/container/recordio"
	"io"
)

type JsonReaderFactoryImpl struct {
	recordio.AbstractRecordReaderFactory
	// TODO: Add needed configs
}

func (j *JsonReaderFactoryImpl) CreateReader(
	context api.StageContext,
	reader io.Reader,
	messageId string,
) (dataformats.RecordReader, error) {
	var recordReader dataformats.RecordReader
	recordReader = newRecordReader(context, reader, messageId)
	return recordReader, nil
}

type JsonReaderImpl struct {
	context   api.StageContext
	reader    io.Reader
	decoder   *json.Decoder
	messageId string
	counter   int
}

func (jsonReader *JsonReaderImpl) ReadRecord() (api.Record, error) {
	var f interface{}
	err := jsonReader.decoder.Decode(&f)
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}
	jsonReader.counter++
	sourceId := common.CreateRecordId(jsonReader.messageId, jsonReader.counter)
	return jsonReader.context.CreateRecord(sourceId, f)
}

func (jsonReader *JsonReaderImpl) Close() error {
	return recordio.Close(jsonReader.reader)
}

func newRecordReader(context api.StageContext, reader io.Reader, messageId string) *JsonReaderImpl {
	return &JsonReaderImpl{
		context:   context,
		reader:    reader,
		decoder:   json.NewDecoder(reader),
		messageId: messageId,
		counter:   0,
	}
}
