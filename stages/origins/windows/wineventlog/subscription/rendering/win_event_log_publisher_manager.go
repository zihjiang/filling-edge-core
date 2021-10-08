// +build 386 windows,amd64 windows


package rendering

import (
	"errors"
	winevtcommon "datacollector-edge/stages/origins/windows/wineventlog/common"
)

type winEventLogPublisherManager struct {
	providerToPublisherMetadataHandle map[string]winevtcommon.PublisherMetadataHandle
}

func (welpm *winEventLogPublisherManager) GetPublisherHandle(
	provider string,
) (winevtcommon.PublisherMetadataHandle, error) {
	var err error
	providerHandle := winevtcommon.PublisherMetadataHandle(0)
	if provider != "" {
		var ok bool
		providerHandle, ok = welpm.providerToPublisherMetadataHandle[provider]
		if !ok {
			providerHandle, err = winevtcommon.EvtOpenPublisherMetadata(provider)
		}
	} else {
		err = errors.New("invalid arg - provider empty")
	}
	return providerHandle, err
}

func (welpm *winEventLogPublisherManager) Close() {
	for _, publisherMetadataHandle := range welpm.providerToPublisherMetadataHandle {
		publisherMetadataHandle.Close()
	}
}
