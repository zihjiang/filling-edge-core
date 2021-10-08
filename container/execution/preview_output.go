
package execution

import "datacollector-edge/api/validation"

type PreviewOutput struct {
	PreviewStatus string              `json:"status"`
	Issues        *validation.Issues  `json:"issues"`
	Output        [][]StageOutputJson `json:"batchesOutput"`
	Message       string              `json:"message"`
}

func NewPreviewOutput(batchOutputs [][]StageOutput) ([][]StageOutputJson, error) {
	batchOutputsJson := make([][]StageOutputJson, len(batchOutputs))
	for batchIndex, batchOutput := range batchOutputs {
		batchOutputJson := make([]StageOutputJson, len(batchOutput))
		for stageIndex, stageOutput := range batchOutput {
			stageOutputJson, err := NewStageOutputJson(stageOutput)
			if err != nil {
				return batchOutputsJson, err
			}
			batchOutputJson[stageIndex] = *stageOutputJson
		}
		batchOutputsJson[batchIndex] = batchOutputJson
	}

	return batchOutputsJson, nil
}
