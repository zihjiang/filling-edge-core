
package common

import (
	"datacollector-edge/api"
)

type ErrorSink struct {
	stageErrorMessages map[string][]api.ErrorMessage
	stageErrorRecords  map[string][]api.Record
	totalErrorRecords  int64
	totalErrorMessages int64
}

func NewErrorSink() *ErrorSink {
	errorSink := &ErrorSink{}
	errorSink.ClearErrorRecordsAndMessages()
	errorSink.totalErrorMessages = 0
	errorSink.totalErrorRecords = 0
	return errorSink
}

//After each batch call this function to clear current batch error messages/records
func (e *ErrorSink) ClearErrorRecordsAndMessages() {
	e.stageErrorMessages = make(map[string][]api.ErrorMessage)
	e.stageErrorRecords = make(map[string][]api.Record)
	e.totalErrorMessages = 0
	e.totalErrorRecords = 0
}

func (e *ErrorSink) GetStageErrorMessages(stageIns string) []api.ErrorMessage {
	return e.stageErrorMessages[stageIns]
}

func (e *ErrorSink) GetStageErrorRecords(stageIns string) []api.Record {
	return e.stageErrorRecords[stageIns]
}

func (e *ErrorSink) GetTotalErrorMessages() int64 {
	return e.totalErrorMessages
}

func (e *ErrorSink) GetTotalErrorRecords() int64 {
	return e.totalErrorRecords
}

func (e *ErrorSink) GetErrorRecords() map[string][]api.Record {
	return e.stageErrorRecords
}

func (e *ErrorSink) GetErrorMessages() map[string][]api.ErrorMessage {
	return e.stageErrorMessages
}

func (e *ErrorSink) ReportError(stageIns string, errorMessage api.ErrorMessage) {
	var errorMessages []api.ErrorMessage
	var keyExists bool
	errorMessages, keyExists = e.stageErrorMessages[stageIns]

	if !keyExists {
		errorMessages = make([]api.ErrorMessage, 0)
	}

	errorMessages = append(errorMessages, errorMessage)
	e.stageErrorMessages[stageIns] = errorMessages
	e.totalErrorMessages += 1
}

func (e *ErrorSink) ToError(stageIns string, record api.Record) {
	var errorRecords []api.Record
	var keyExists bool
	errorRecords, keyExists = e.stageErrorRecords[stageIns]

	if !keyExists {
		errorRecords = []api.Record{}
	}
	errorRecords = append(errorRecords, record)
	e.stageErrorRecords[stageIns] = errorRecords
	e.totalErrorRecords += 1
}
