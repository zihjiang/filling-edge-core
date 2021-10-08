// +build 386 windows,amd64 windows


package subscription

import (
	log "github.com/sirupsen/logrus"
	wineventsyscall "datacollector-edge/stages/origins/windows/wineventlog/common"
	"syscall"
	"unsafe"
)

type pushWinEventSubscriber struct {
	*baseWinEventSubscriber
}

func (pwes *pushWinEventSubscriber) Subscribe() error {
	pwes.subscriptionCallback = func(
		action wineventsyscall.EvtSubscribeNotifyAction,
		userContext unsafe.Pointer,
		eventHandle wineventsyscall.EventHandle,
	) syscall.Errno {
		var returnStatus syscall.Errno
		switch action {
		case wineventsyscall.EvtSubscribeActionError:
			if wineventsyscall.ErrorEvtQueryResultStale == returnStatus {
				log.Error("The subscription callback was notified that eventHandle records are missing")
			} else {
				log.WithError(syscall.Errno(eventHandle)).Error("The subscription callback received the following Win32 error")
			}
		case wineventsyscall.EvtSubscribeActionDeliver:
			eventRecord, err := pwes.renderer.RenderEvent(pwes.stageContext, eventHandle, pwes.bookMarkHandle)
			if err == nil {
				pwes.eventsQueue.Put(eventRecord)
			} else {
				log.WithError(err).Errorf("Error rendering from event handle %d", eventHandle)
			}
		}
		return returnStatus
	}
	return pwes.baseWinEventSubscriber.Subscribe()
}
