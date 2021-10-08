
package jsonrecord

import (
	"bytes"
	"encoding/json"
	"datacollector-edge/api/linkedhashmap"
	"testing"
)

func TestWriteMapRecord(t *testing.T) {
	stageContext := CreateStageContext()
	commits := map[string]interface{}{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	record1, err := stageContext.CreateRecord("Id1", commits)
	if err != nil {
		t.Fatal(err)
	}
	record1.GetHeader().SetAttribute("Sample Attribute", "Sample Value1")

	bufferWriter := bytes.NewBuffer([]byte{})

	recordWriterFactory := &JsonWriterFactoryImpl{Mode: MultipleObjects}
	recordWriter, err := recordWriterFactory.CreateWriter(stageContext, bufferWriter)

	if err != nil {
		t.Fatal(err)
	}

	err = recordWriter.WriteRecord(record1)
	if err != nil {
		t.Fatal(err)
	}

	recordWriter.Flush()
	recordWriter.Close()

	decoder := json.NewDecoder(bufferWriter)
	var recordObject = make(map[string]int)
	err = decoder.Decode(&recordObject)
	if err != nil {
		t.Fatal(err)
	}

	if recordObject["rsc"] != commits["rsc"] {
		t.Errorf("Excepted: %d, but got: %d", commits["rsc"], recordObject["rsc"])
	}

	if recordObject["r"] != commits["r"] {
		t.Errorf("Excepted: %d, but got: %d", commits["r"], recordObject["r"])
	}

	if recordObject["gri"] != commits["gri"] {
		t.Errorf("Excepted: %d, but got: %d", commits["gri"], recordObject["gri"])
	}

	if recordObject["adg"] != commits["adg"] {
		t.Errorf("Excepted: %d, but got: %d", commits["adg"], recordObject["adg"])
	}
}

func TestWriteListMapRecord(t *testing.T) {
	stageContext := CreateStageContext()
	commits := linkedhashmap.New()
	commits.Put("rsc", 3711)
	commits.Put("r", 2138)
	commits.Put("gri", 1908)
	commits.Put("adg", 912)

	record1, err := stageContext.CreateRecord("Id1", commits)
	if err != nil {
		t.Fatal(err)
	}
	record1.GetHeader().SetAttribute("Sample Attribute", "Sample Value1")

	bufferWriter := bytes.NewBuffer([]byte{})

	recordWriterFactory := &JsonWriterFactoryImpl{Mode: MultipleObjects}
	recordWriter, err := recordWriterFactory.CreateWriter(stageContext, bufferWriter)

	if err != nil {
		t.Fatal(err)
	}

	err = recordWriter.WriteRecord(record1)
	if err != nil {
		t.Fatal(err)
	}

	recordWriter.Flush()
	recordWriter.Close()

	decoder := json.NewDecoder(bufferWriter)
	var recordObject = make(map[string]int)
	err = decoder.Decode(&recordObject)
	if err != nil {
		t.Fatal(err)
	}

	if recordObject["rsc"] != 3711 {
		t.Errorf("Excepted: %d, but got: %d", 3711, recordObject["rsc"])
	}

	if recordObject["r"] != 2138 {
		t.Errorf("Excepted: %d, but got: %d", 2138, recordObject["r"])
	}

	if recordObject["gri"] != 1908 {
		t.Errorf("Excepted: %d, but got: %d", 1908, recordObject["gri"])
	}

	if recordObject["adg"] != 912 {
		t.Errorf("Excepted: %d, but got: %d", 912, recordObject["adg"])
	}
}

func TestWriteListRecord(t *testing.T) {
	stageContext := CreateStageContext()
	stringSlice := []interface{}{"apple", "orange", "banana"}
	record1, err := stageContext.CreateRecord("Id1", stringSlice)
	if err != nil {
		t.Fatal(err)
	}
	record1.GetHeader().SetAttribute("Sample Attribute", "Sample Value1")
	bufferWriter := bytes.NewBuffer([]byte{})

	//
	recordWriterFactory := &JsonWriterFactoryImpl{Mode: MultipleObjects}
	recordWriter, err := recordWriterFactory.CreateWriter(stageContext, bufferWriter)
	if err != nil {
		t.Fatal(err)
	}
	err = recordWriter.WriteRecord(record1)
	if err != nil {
		t.Fatal(err)
	}
	recordWriter.Flush()
	recordWriter.Close()

	decoder := json.NewDecoder(bufferWriter)
	var listRecordObject = []interface{}{}
	err = decoder.Decode(&listRecordObject)
	if err != nil {
		t.Fatal(err)
	}

	if len(listRecordObject) != 3 {
		t.Errorf("Excepted length: 3, but got: %d", len(listRecordObject))
	}

	if listRecordObject[0] != stringSlice[0] {
		t.Errorf("Excepted: %s, but got: %s", stringSlice[0], listRecordObject[0])
	}
}

func TestWriteMapRecord_ArrayObjects(t *testing.T) {
	stageContext := CreateStageContext()
	commits := map[string]interface{}{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	record1, err := stageContext.CreateRecord("Id1", commits)
	if err != nil {
		t.Fatal(err)
	}
	record1.GetHeader().SetAttribute("Sample Attribute", "Sample Value1")

	bufferWriter := bytes.NewBuffer([]byte{})

	recordWriterFactory := &JsonWriterFactoryImpl{Mode: ArrayObjects}
	recordWriter, err := recordWriterFactory.CreateWriter(stageContext, bufferWriter)

	if err != nil {
		t.Fatal(err)
	}

	err = recordWriter.WriteRecord(record1)
	if err != nil {
		t.Fatal(err)
	}

	recordWriter.Flush()
	recordWriter.Close()

	decoder := json.NewDecoder(bufferWriter)
	var recordObjectList = make([]map[string]int, 1)
	err = decoder.Decode(&recordObjectList)
	if err != nil {
		t.Fatal(err)
	}

	recordObject := recordObjectList[0]

	if recordObject["rsc"] != commits["rsc"] {
		t.Errorf("Excepted: %d, but got: %d", commits["rsc"], recordObject["rsc"])
	}

	if recordObject["r"] != commits["r"] {
		t.Errorf("Excepted: %d, but got: %d", commits["r"], recordObject["r"])
	}

	if recordObject["gri"] != commits["gri"] {
		t.Errorf("Excepted: %d, but got: %d", commits["gri"], recordObject["gri"])
	}

	if recordObject["adg"] != commits["adg"] {
		t.Errorf("Excepted: %d, but got: %d", commits["adg"], recordObject["adg"])
	}
}
