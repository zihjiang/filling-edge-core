
package store

import (
	"encoding/json"
	"datacollector-edge/container/common"
	"io/ioutil"
	"os"
	"strings"
)

var BaseDir = "."

const (
	OFFSET_FILE               = "offset.json"
	PIPELINES_RUN_INFO_FOLDER = "/data/runInfo/"
)

func GetOffset(pipelineId string) (common.SourceOffset, error) {
	defaultSourceOffset := common.GetDefaultOffset()
	if _, err := os.Stat(getPipelineOffsetFile(pipelineId)); os.IsNotExist(err) {
		return defaultSourceOffset, nil
	} else {
		file, readError := ioutil.ReadFile(getPipelineOffsetFile(pipelineId))

		if readError != nil {
			return defaultSourceOffset, readError
		}

		var sourceOffset common.SourceOffset
		json.Unmarshal(file, &sourceOffset)
		return sourceOffset, nil
	}
}

func SaveOffset(pipelineId string, sourceOffset common.SourceOffset) error {
	var err error
	var offsetJson []byte
	if offsetJson, err = json.Marshal(sourceOffset); err == nil {
		err = ioutil.WriteFile(getPipelineOffsetFile(pipelineId), offsetJson, 0644)
	}
	return err
}

func ResetOffset(pipelineId string) error {
	return SaveOffset(pipelineId, common.GetDefaultOffset())
}

func getPipelineOffsetFile(pipelineId string) string {
	return getRunInfoDir(pipelineId) + OFFSET_FILE
}

func getRunInfoDir(pipelineId string) string {
	validPipelineId := strings.Replace(pipelineId, ":", "", -1)
	return BaseDir + PIPELINES_RUN_INFO_FOLDER + validPipelineId + "/"
}
