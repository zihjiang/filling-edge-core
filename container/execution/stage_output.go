
package execution

import (
	"datacollector-edge/api"
	"datacollector-edge/container/recordio/sdcrecord"
)

type StageOutput struct {
	InstanceName string
	Output       map[string][]api.Record
	EventRecords []api.Record
	ErrorRecords []api.Record
	StageErrors  []api.ErrorMessage
}

type StageOutputJson struct {
	InstanceName string                           `json:"instanceName"`
	Output       map[string][]sdcrecord.SDCRecord `json:"output"`
	EventRecords []sdcrecord.SDCRecord            `json:"eventRecords"`
	ErrorRecords []sdcrecord.SDCRecord            `json:"errorRecords"`
	StageErrors  []api.ErrorMessage               `json:"stageErrors"`
}

func NewStageOutputJson(stageOutput StageOutput) (*StageOutputJson, error) {
	stageOutputJson := &StageOutputJson{
		InstanceName: stageOutput.InstanceName,
		StageErrors:  stageOutput.StageErrors,
		Output:       make(map[string][]sdcrecord.SDCRecord),
		ErrorRecords: make([]sdcrecord.SDCRecord, len(stageOutput.ErrorRecords)),
		EventRecords: make([]sdcrecord.SDCRecord, len(stageOutput.EventRecords)),
	}

	for outpuLane, records := range stageOutput.Output {
		stageOutputJson.Output[outpuLane] = make([]sdcrecord.SDCRecord, len(records))
		for i, record := range records {
			sdcRecord, err := sdcrecord.NewSdcRecordFromRecord(record)
			if err != nil {
				return nil, err
			}
			stageOutputJson.Output[outpuLane][i] = *sdcRecord
		}
	}

	for i, record := range stageOutput.EventRecords {
		sdcRecord, err := sdcrecord.NewSdcRecordFromRecord(record)
		if err != nil {
			return nil, err
		}
		stageOutputJson.EventRecords[i] = *sdcRecord
	}

	for i, record := range stageOutput.ErrorRecords {
		sdcRecord, err := sdcrecord.NewSdcRecordFromRecord(record)
		if err != nil {
			return nil, err
		}
		stageOutputJson.ErrorRecords[i] = *sdcRecord
	}

	return stageOutputJson, nil
}

func NewStageOutput(stageContext api.StageContext, stageOutputJson StageOutputJson) (*StageOutput, error) {
	stageOutput := &StageOutput{
		InstanceName: stageOutputJson.InstanceName,
		StageErrors:  stageOutputJson.StageErrors,
		Output:       make(map[string][]api.Record),
		ErrorRecords: make([]api.Record, len(stageOutputJson.ErrorRecords)),
	}

	for stageInstanceName, records := range stageOutputJson.Output {
		stageOutput.Output[stageInstanceName] = make([]api.Record, len(records))
		for i, record := range records {
			sdcRecord, err := sdcrecord.NewRecordFromSDCRecord(stageContext, &record)
			if err != nil {
				return nil, err
			}
			stageOutput.Output[stageInstanceName][i] = sdcRecord
		}
	}

	for i, record := range stageOutput.EventRecords {
		sdcRecord, err := sdcrecord.NewSdcRecordFromRecord(record)
		if err != nil {
			return nil, err
		}
		stageOutputJson.EventRecords[i] = *sdcRecord
	}

	for i, record := range stageOutput.ErrorRecords {
		sdcRecord, err := sdcrecord.NewSdcRecordFromRecord(record)
		if err != nil {
			return nil, err
		}
		stageOutputJson.ErrorRecords[i] = *sdcRecord
	}

	return stageOutput, nil
}
