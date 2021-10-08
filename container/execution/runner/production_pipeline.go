
package runner

import (
	"github.com/rcrowley/go-metrics"
	log "github.com/sirupsen/logrus"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/container/execution"
)

const (
	IssueErrorTemplate = "Initialization Error '%s' on Instance : '%s' "
)

type ProductionPipeline struct {
	PipelineConfig common.PipelineConfiguration
	Pipeline       *Pipeline
	MetricRegistry metrics.Registry
}

func (p *ProductionPipeline) Init() []validation.Issue {
	issues := p.Pipeline.Init()
	if len(issues) != 0 {
		for _, issue := range issues {
			log.Printf("[ERROR] "+IssueErrorTemplate, issue.Message, issue.InstanceName)
		}
	}
	return issues
}

func (p *ProductionPipeline) Run() {
	log.Debug("Production Pipeline Run")
	p.Pipeline.Run()
}

func (p *ProductionPipeline) Stop() {
	log.Debug("Production Pipeline Stop")
	p.Pipeline.Stop()
}

func NewProductionPipeline(
	pipelineId string,
	config execution.Config,
	runner execution.Runner,
	pipelineConfiguration common.PipelineConfiguration,
	runtimeParameters map[string]interface{},
) (*ProductionPipeline, []validation.Issue) {
	if sourceOffsetTracker, err := NewProductionSourceOffsetTracker(pipelineId); err == nil {
		metricRegistry := metrics.NewRegistry()
		pipeline, issues := NewPipeline(
			config,
			runner.GetPipelineConfig(),
			sourceOffsetTracker,
			runtimeParameters,
			metricRegistry,
		)
		return &ProductionPipeline{
			PipelineConfig: pipelineConfiguration,
			Pipeline:       pipeline,
			MetricRegistry: metricRegistry,
		}, issues
	} else {
		issues := make([]validation.Issue, 0)
		issues = append(issues, validation.Issue{
			Count:   1,
			Message: err.Error(),
		})
		return nil, issues
	}
}
