
package common

import (
	"datacollector-edge/api"
)

type EventSink struct {
	eventRecords map[string][]api.Record
}

func NewEventSink() *EventSink {
	eventSink := &EventSink{}
	eventSink.ClearEventRecords()
	return eventSink
}

func (e *EventSink) ClearEventRecords() {
	e.eventRecords = make(map[string][]api.Record)
}

func (e *EventSink) GetStageEvents(stageIns string) []api.Record {
	return e.eventRecords[stageIns]
}

func (e *EventSink) AddEvent(stageIns string, record api.Record) {
	var eventRecords []api.Record
	var keyExists bool
	eventRecords, keyExists = e.eventRecords[stageIns]

	if !keyExists {
		eventRecords = []api.Record{}
	}
	eventRecords = append(eventRecords, record)
	e.eventRecords[stageIns] = eventRecords
}
