// +build 386 windows,amd64 windows


package wineventlog

import (
	"fmt"
	"github.com/AllenDang/w32"
	log "github.com/sirupsen/logrus"
	"datacollector-edge/api"
	"datacollector-edge/container/common"
	wincommon "datacollector-edge/stages/origins/windows/common"
	winevtsubscription "datacollector-edge/stages/origins/windows/wineventlog/subscription"
	winevtrender "datacollector-edge/stages/origins/windows/wineventlog/subscription/rendering"
	"time"
)

type windowsEventLogReader struct {
	*common.BaseStage
	*wincommon.BaseEventLogReader
	eventSubscriber winevtsubscription.WinEventSubscriber
	offset          string
	handle          w32.HANDLE
}

func (welr *windowsEventLogReader) Open() error {
	err := welr.eventSubscriber.Subscribe()
	if err != nil {
		log.WithError(err).Error("Error subscribing")
	}
	return err
}

func (welr *windowsEventLogReader) Read() ([]api.Record, error) {
	eventRecords, err := welr.eventSubscriber.GetRecords()
	if err != nil {
		log.WithError(err).Error("Error reading from windows event log")
	}
	return eventRecords, err
}

func (welr *windowsEventLogReader) GetCurrentOffset() string {
	return welr.eventSubscriber.GetBookmark()
}

func (welr *windowsEventLogReader) Close() error {
	welr.eventSubscriber.Close()
	return nil
}

func NewWindowsEventLogReader(
	baseStage *common.BaseStage,
	logName string,
	mode wincommon.EventLogReaderMode,
	bufferSize int,
	maxBatchSize int,
	lastSourceOffset string,
	winEventLogConf wincommon.WinEventLogConf,
) (wincommon.EventLogReader, error) {
	subscriptionMode := winevtsubscription.SubscriptionMode(winEventLogConf.SubscriptionMode)
	rawEventPopulationStrategy := winevtrender.RawEventPopulationStrategy(winEventLogConf.RawEventPopulationStrategy)

	query := fmt.Sprintf(`<QueryList> <Query Id="0"> <Select Path="%s">*</Select> </Query></QueryList>`, logName)
	log.Debugf("Querying windows Event log with %s", logName)
	return &windowsEventLogReader{
		BaseStage:          baseStage,
		BaseEventLogReader: &wincommon.BaseEventLogReader{Log: logName, Mode: mode},
		eventSubscriber: winevtsubscription.NewWinEventSubscriber(
			baseStage.GetStageContext(),
			subscriptionMode,
			rawEventPopulationStrategy,
			query,
			uint32(maxBatchSize),
			lastSourceOffset,
			mode,
			bufferSize,
			time.Duration(int64(winEventLogConf.MaxWaitTimeSecs))*time.Second,
		),
		offset: lastSourceOffset,
	}, nil
}
