

package api

import "datacollector-edge/api/validation"

type Service interface {
	Init(stageContext StageContext) []validation.Issue
	Destroy() error
}
