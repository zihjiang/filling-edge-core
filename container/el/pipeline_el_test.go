
package el

import (
	"context"
	"testing"
	"time"
)

func TestPipelineEL(test *testing.T) {
	pipelineId := "samplePipelineId"
	pipelineTitle := "Sample Pipeline"
	pipelineUser := "admin"
	pipelineStartTime := time.Now()
	evaluationTests := []EvaluationTest{
		{
			Name:       "Test pipeline:id()",
			Expression: "${pipeline:id()}",
			Expected:   pipelineId,
		},
		{
			Name:       "Test function pipeline:id() - Error 1",
			Expression: "${pipeline:id('invalid param')}",
			Expected:   "The function 'pipeline:id' requires 0 arguments but was passed 1",
			ErrorCase:  true,
		},

		{
			Name:       "Test pipeline:title()",
			Expression: "${pipeline:title()}",
			Expected:   pipelineTitle,
		},
		{
			Name:       "Test function pipeline:title() - Error 1",
			Expression: "${pipeline:title('invalid param')}",
			Expected:   "The function 'pipeline:title' requires 0 arguments but was passed 1",
			ErrorCase:  true,
		},

		{
			Name:       "Test pipeline:user()",
			Expression: "${pipeline:user()}",
			Expected:   pipelineUser,
		},
		{
			Name:       "Test function pipeline:user() - Error 1",
			Expression: "${pipeline:user('invalid param')}",
			Expected:   "The function 'pipeline:user' requires 0 arguments but was passed 1",
			ErrorCase:  true,
		},

		{
			Name:       "Test pipeline:startTime()",
			Expression: "${pipeline:startTime()}",
			Expected:   pipelineStartTime,
		},
		{
			Name:       "Test function pipeline:startTime() - Error 1",
			Expression: "${pipeline:startTime('invalid param')}",
			Expected:   "The function 'pipeline:startTime' requires 0 arguments but was passed 1",
			ErrorCase:  true,
		},
	}

	pipelineELContextValues := map[string]interface{}{
		PipelineIdContextVar:        pipelineId,
		PipelineTitleContextVar:     pipelineTitle,
		PipelineUserContextVar:      pipelineUser,
		PipelineStartTimeContextVar: pipelineStartTime,
	}
	pipelineElContext := context.WithValue(context.Background(), PipelineElContextVar, pipelineELContextValues)

	RunEvaluationTests(evaluationTests, []Definitions{&PipelineEL{Context: pipelineElContext}}, test)
}

func TestPipelineELUndefinedValues(test *testing.T) {
	evaluationTests := []EvaluationTest{
		{
			Name:       "Test pipeline:id()",
			Expression: "${pipeline:id()}",
			Expected:   UndefinedValue,
		},
		{
			Name:       "Test pipeline:title()",
			Expression: "${pipeline:title()}",
			Expected:   UndefinedValue,
		},
		{
			Name:       "Test pipeline:user()",
			Expression: "${pipeline:user()}",
			Expected:   UndefinedValue,
		},
	}
	RunEvaluationTests(evaluationTests, []Definitions{&PipelineEL{Context: context.Background()}}, test)
}
