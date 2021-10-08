
package runner

import (
	"datacollector-edge/api"
	"datacollector-edge/container/common"
)

type BatchMakerImpl struct {
	stagePipe           StagePipe
	stageOutput         map[string][]api.Record
	singleOutputLane    string
	size                int64
	StageOutputSnapshot map[string][]api.Record
}

func (b *BatchMakerImpl) GetLanes() []string {
	return b.stagePipe.OutputLanes
}

func (b *BatchMakerImpl) AddRecord(record api.Record, outputLanes ...string) {
	if len(outputLanes) > 0 {
		// This is a bit costly we are cloning all records (greedy) before going to the stages
		// We can do better by simply cloning for the immediate stage which is going to process
		// To save up on memory
		for _, outputLane := range outputLanes {
			recordCopy := b.getRecordForBatchMaker(record)
			b.stageOutput[outputLane] = append(b.stageOutput[outputLane], recordCopy)
		}
	} else {
		recordCopy := b.getRecordForBatchMaker(record)
		b.stageOutput[b.singleOutputLane] = append(b.stageOutput[b.singleOutputLane], recordCopy)
	}
	b.size++

	if b.StageOutputSnapshot != nil {
		recordCopy := b.getRecordForBatchMaker(record)
		if len(outputLanes) > 0 {
			for _, outputLane := range outputLanes {
				b.StageOutputSnapshot[outputLane] = append(b.StageOutputSnapshot[outputLane], recordCopy)
			}
		} else {
			b.StageOutputSnapshot[b.singleOutputLane] = append(b.StageOutputSnapshot[b.singleOutputLane], recordCopy)
		}
	}
}

func (b *BatchMakerImpl) GetStageOutput(outputLane ...string) []api.Record {
	if len(outputLane) > 0 {
		return b.stageOutput[outputLane[0]]
	}
	return b.stageOutput[b.singleOutputLane]
}

func (b *BatchMakerImpl) GetSize() int64 {
	return b.size
}

func (b *BatchMakerImpl) getRecordForBatchMaker(record api.Record) api.Record {
	recordCopy := record.Clone()
	if b.stagePipe.Stage.config != nil {
		headerImplForRecord := recordCopy.GetHeader().(*common.HeaderImpl)
		common.AddStageToStagePath(headerImplForRecord, b.stagePipe.Stage.config.InstanceName)
		common.CreateTrackingId(headerImplForRecord)
	}
	return recordCopy
}

func NewBatchMakerImpl(stagePipe StagePipe, keepSnapshot bool) *BatchMakerImpl {
	batchMaker := &BatchMakerImpl{stagePipe: stagePipe}
	batchMaker.stageOutput = make(map[string][]api.Record)
	for _, outputLane := range stagePipe.OutputLanes {
		batchMaker.stageOutput[outputLane] = make([]api.Record, 0)
	}
	if len(stagePipe.OutputLanes) > 0 {
		batchMaker.singleOutputLane = stagePipe.OutputLanes[0]
	}
	if keepSnapshot {
		batchMaker.StageOutputSnapshot = make(map[string][]api.Record)
	}
	return batchMaker
}
