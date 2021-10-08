
package delimitedrecord

import (
	"datacollector-edge/api"
	"datacollector-edge/api/fieldtype"
	"datacollector-edge/api/linkedhashmap"
	"testing"
)

func TestReadDelimitedRecord_CreateRecord_Custom_Format(t *testing.T) {
	stageContext := CreateStageContext()
	readerFactoryImpl := &RecordCreator{
		CsvFileFormat:      Custom,
		CsvCustomDelimiter: "|",
		CsvRecordType:      ListMap,
	}
	record, err := readerFactoryImpl.CreateRecord(stageContext, "a|b|c", "m", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rootField, _ := record.Get()
	if rootField.Type != fieldtype.LIST_MAP {
		t.Errorf("Excpeted record type : LIST_MAP, but received: %s", rootField.Type)
	}

	listMapField := rootField.Value.(*linkedhashmap.Map)

	if val, found := listMapField.Get("0"); !found {
		t.Errorf("Missing 0 key")
	} else {
		valField := val.(*api.Field)
		if valField.Value != "a" {
			t.Errorf("Excpeted field value a, but received: %s", val)
		}
	}

	if val, found := listMapField.Get("1"); !found {
		t.Errorf("Missing 1 key")
	} else {
		valField := val.(*api.Field)
		if valField.Value != "b" {
			t.Errorf("Excpeted field value b, but received: %s", val)
		}
	}

	if val, found := listMapField.Get("2"); !found {
		t.Errorf("Missing 2 key")
	} else {
		valField := val.(*api.Field)
		if valField.Value != "c" {
			t.Errorf("Excpeted field value c, but received: %s", val)
		}
	}

	// check order
	keys := listMapField.Keys()

	if keys[0] != "0" {
		t.Errorf("Expected column 0 in first position")
	}

	if keys[1] != "1" {
		t.Errorf("Expected column 1 in first position")
	}

	if keys[2] != "2" {
		t.Errorf("Expected column 2 in first position")
	}
}
