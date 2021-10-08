
package runner

import (
	"datacollector-edge/container/common"
	"datacollector-edge/container/execution/store"
	"time"
)

type ProductionSourceOffsetTracker struct {
	pipelineId    string
	currentOffset common.SourceOffset
	newOffset     *string
	finished      bool
	lastBatchTime time.Time
}

var emptyOffset = ""

func (o *ProductionSourceOffsetTracker) IsFinished() bool {
	return o.finished
}

func (o *ProductionSourceOffsetTracker) SetOffset(newOffset *string) {
	o.newOffset = newOffset
}

func (o *ProductionSourceOffsetTracker) CommitOffset() error {
	o.currentOffset.Offset[common.PollSourceOffsetKey] = o.newOffset
	o.finished = o.currentOffset.Offset[common.PollSourceOffsetKey] == nil
	o.newOffset = &emptyOffset
	return store.SaveOffset(o.pipelineId, o.currentOffset)
}

func (o *ProductionSourceOffsetTracker) GetOffset() *string {
	return o.currentOffset.Offset[common.PollSourceOffsetKey]
}

func (o *ProductionSourceOffsetTracker) GetLastBatchTime() time.Time {
	return o.lastBatchTime
}

func NewProductionSourceOffsetTracker(pipelineId string) (*ProductionSourceOffsetTracker, error) {
	if sourceOffset, err := store.GetOffset(pipelineId); err == nil {
		return &ProductionSourceOffsetTracker{
			pipelineId:    pipelineId,
			currentOffset: sourceOffset,
		}, nil
	} else {
		return nil, err
	}
}
