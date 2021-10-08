
package preview

import (
	"datacollector-edge/container/common"
	"datacollector-edge/container/execution"
	"time"
)

type PreviewSourceOffsetTracker struct {
	pipelineId    string
	currentOffset common.SourceOffset
	newOffset     *string
	finished      bool
	lastBatchTime time.Time
}

var emptyOffset = ""

func (o *PreviewSourceOffsetTracker) IsFinished() bool {
	return false
}

func (o *PreviewSourceOffsetTracker) SetOffset(newOffset *string) {
	o.newOffset = newOffset
}

func (o *PreviewSourceOffsetTracker) CommitOffset() error {
	o.currentOffset.Offset[common.PollSourceOffsetKey] = o.newOffset
	o.finished = o.currentOffset.Offset[common.PollSourceOffsetKey] == &emptyOffset
	o.newOffset = &emptyOffset
	return nil
}

func (o *PreviewSourceOffsetTracker) GetOffset() *string {
	return o.currentOffset.Offset[common.PollSourceOffsetKey]
}

func (o *PreviewSourceOffsetTracker) GetLastBatchTime() time.Time {
	return o.lastBatchTime
}

func NewPreviewSourceOffsetTracker(pipelineId string) execution.SourceOffsetTracker {
	return &PreviewSourceOffsetTracker{
		pipelineId:    pipelineId,
		currentOffset: common.GetDefaultOffset(),
	}
}
