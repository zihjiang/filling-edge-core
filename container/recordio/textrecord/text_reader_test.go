
package textrecord

import (
	"bytes"
	"datacollector-edge/api"
	"datacollector-edge/api/fieldtype"
	"datacollector-edge/container/common"
	"strconv"
	"testing"
)

func CreateStageContext() api.StageContext {
	return &common.StageContextImpl{
		StageConfig: &common.StageConfiguration{InstanceName: "Dummy Stage"},
		Parameters:  nil,
	}
}

func TestReadTextRecord(t *testing.T) {
	sampleTextData := bytes.NewBuffer([]byte("test data 1\r\ntest data 2\ntest data 3"))
	testReadTextRecord(t, sampleTextData)
}

func TestReadTextRecordMaxLen(t *testing.T) {
	sampleTextData := bytes.NewBuffer([]byte("test data 1 extra texta\r\ntest data 2 extra\ntest data 3 extra"))
	testReadTextRecord(t, sampleTextData)
}

func testReadTextRecord(t *testing.T, sampleTextData *bytes.Buffer) {
	stageContext := CreateStageContext()
	readerFactoryImpl := &TextReaderFactoryImpl{TextMaxLineLen: 11}
	recordReader, err := readerFactoryImpl.CreateReader(stageContext, sampleTextData, "m")
	if err != nil {
		t.Fatal(err.Error())
	}

	recordCount := 0
	for {
		record, err := recordReader.ReadRecord()
		if err != nil {
			t.Fatal(err.Error())
		}

		if record == nil {
			break
		}

		rootField, _ := record.Get()
		if rootField.Type != fieldtype.MAP {
			t.Errorf("Excpeted record type : Map, but received: %s", rootField.Type)
		}

		mapField := rootField.Value.(map[string]*api.Field)
		testData := "test data " + strconv.Itoa(recordCount+1)
		if mapField["text"].Value.(string) != testData {
			t.Errorf("Excpeted field value %s, but received: %s", testData, mapField["text"].Value)
		}
		recordCount++
	}

	if recordCount != 3 {
		t.Errorf("Excpeted 3 records, but received: %d", recordCount)
	}

	_ = recordReader.Close()
}
