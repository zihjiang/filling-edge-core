
package el

import (
	"context"
	"errors"
	"fmt"
	"github.com/madhukard/govaluate"
	"time"
)

const (
	PipelineElContextVar        = "PIPELINE_EL"
	PipelineIdContextVar        = "PIPELINE_ID"
	PipelineTitleContextVar     = "PIPELINE_TITLE"
	PipelineUserContextVar      = "PIPELINE_USER"
	PipelineStartTimeContextVar = "PIPELINE_START_TIME"
	UndefinedValue              = "UNDEFINED"
)

type PipelineEL struct {
	Context context.Context
}

func (p *PipelineEL) GetId(args ...interface{}) (interface{}, error) {
	if len(args) != 0 {
		return "", errors.New(
			fmt.Sprintf("The function 'pipeline:id' requires 0 arguments but was passed %d", len(args)),
		)
	}

	if p.Context != nil && p.Context.Value(PipelineElContextVar) != nil {
		pipelineELContextValues := p.Context.Value(PipelineElContextVar).(map[string]interface{})
		return pipelineELContextValues[PipelineIdContextVar], nil
	}

	return UndefinedValue, nil
}

func (p *PipelineEL) GetTitle(args ...interface{}) (interface{}, error) {
	if len(args) != 0 {
		return "", errors.New(
			fmt.Sprintf("The function 'pipeline:title' requires 0 arguments but was passed %d", len(args)),
		)
	}

	if p.Context != nil && p.Context.Value(PipelineElContextVar) != nil {
		pipelineELContextValues := p.Context.Value(PipelineElContextVar).(map[string]interface{})
		return pipelineELContextValues[PipelineTitleContextVar], nil
	}

	return UndefinedValue, nil
}

func (p *PipelineEL) GetUser(args ...interface{}) (interface{}, error) {
	if len(args) != 0 {
		return "", errors.New(
			fmt.Sprintf("The function 'pipeline:user' requires 0 arguments but was passed %d", len(args)),
		)
	}

	if p.Context != nil && p.Context.Value(PipelineElContextVar) != nil {
		pipelineELContextValues := p.Context.Value(PipelineElContextVar).(map[string]interface{})
		return pipelineELContextValues[PipelineUserContextVar], nil
	}

	return UndefinedValue, nil
}

func (p *PipelineEL) GetStartTime(args ...interface{}) (interface{}, error) {
	if len(args) != 0 {
		return "", errors.New(
			fmt.Sprintf("The function 'pipeline:startTime' requires 0 arguments but was passed %d", len(args)),
		)
	}

	if p.Context != nil && p.Context.Value(PipelineElContextVar) != nil {
		pipelineELContextValues := p.Context.Value(PipelineElContextVar).(map[string]interface{})
		return pipelineELContextValues[PipelineStartTimeContextVar], nil
	}

	return time.Now(), nil
}

func (p *PipelineEL) GetELFunctionDefinitions() map[string]govaluate.ExpressionFunction {
	functions := map[string]govaluate.ExpressionFunction{
		"pipeline:id":        p.GetId,
		"pipeline:title":     p.GetTitle,
		"pipeline:user":      p.GetUser,
		"pipeline:startTime": p.GetStartTime,
	}
	return functions
}
