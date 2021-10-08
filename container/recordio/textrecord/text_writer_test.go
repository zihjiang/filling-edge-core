
package textrecord

import (
	"bytes"
	"datacollector-edge/api/linkedhashmap"
	"testing"
)

func TestWriteTextRecord(t *testing.T) {
	testWriteTextRecord(t, DefaultTextField)
	testWriteTextRecord(t, "newTextFieldName")
}

func testWriteTextRecord(t *testing.T, textFieldName string) {
	stageContext := CreateStageContext()
	record1, err := stageContext.CreateRecord("Id1", map[string]interface{}{textFieldName: "log line 1"})
	record2, err := stageContext.CreateRecord("Id2", map[string]interface{}{textFieldName: "log line 2"})
	listMapValue := linkedhashmap.New()
	listMapValue.Put(textFieldName, "log line 3")
	record3, err := stageContext.CreateRecord("Id3", listMapValue)

	bufferWriter := bytes.NewBuffer([]byte{})
	recordWriterFactory := &TextWriterFactoryImpl{TextFieldPath: "/" + textFieldName}
	recordWriter, err := recordWriterFactory.CreateWriter(stageContext, bufferWriter)
	if err != nil {
		t.Fatal(err)
	}

	err = recordWriter.WriteRecord(record1)
	if err != nil {
		t.Fatal(err)
	}

	err = recordWriter.WriteRecord(record2)
	if err != nil {
		t.Fatal(err)
	}

	err = recordWriter.WriteRecord(record3)
	if err != nil {
		t.Fatal(err)
	}

	_ = recordWriter.Flush()
	_ = recordWriter.Close()

	testData := "log line 1\nlog line 2\nlog line 3\n"
	if bufferWriter.String() != "log line 1\nlog line 2\nlog line 3\n" {
		t.Errorf("Excpeted field value %s, but received: %s", testData, bufferWriter.String())
	}
}
