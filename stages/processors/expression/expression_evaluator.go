
package expression

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/container/el"
	"datacollector-edge/stages/stagelibrary"
)

const (
	LIBRARY    = "streamsets-datacollector-basic-lib"
	STAGE_NAME = "com_streamsets_pipeline_stage_processor_expression_ExpressionDProcessor"
	EXPRESSION = "expression"
)

type ExpressionProcessor struct {
	*common.BaseStage
	ExpressionProcessorConfigs []FieldValueConfig      `ConfigDef:"type=MODEL,evaluation=EXPLICIT" ListBeanModel:"name=expressionProcessorConfigs"`
	HeaderAttributeConfigs     []HeaderAttributeConfig `ConfigDef:"type=MODEL,evaluation=EXPLICIT" ListBeanModel:"name=headerAttributeConfigs"`
	FieldAttributeConfigs      []FieldAttributeConfig  `ConfigDef:"type=MODEL,evaluation=EXPLICIT" ListBeanModel:"name=fieldAttributeConfigs"`
	//TODO: Add support for field attributes in SDCE
}

type FieldValueConfig struct {
	FieldToSet string `ConfigDef:"type=STRING,required=true"`
	Expression string `ConfigDef:"type=STRING,evaluation=EXPLICIT,required=true"`
}

type HeaderAttributeConfig struct {
	AttributeToSet string `ConfigDef:"type=STRING,required=true"`
	Expression     string `ConfigDef:"type=STRING,evaluation=EXPLICIT,required=true"`
}

type FieldAttributeConfig struct {
	FieldToSet     string `ConfigDef:"type=STRING,required=true"`
	AttributeToSet string `ConfigDef:"type=STRING,required=true"`
	Expression     string `ConfigDef:"type=STRING,evaluation=EXPLICIT,required=true"`
}

func init() {
	stagelibrary.SetCreator(LIBRARY, STAGE_NAME, func() api.Stage {
		return &ExpressionProcessor{BaseStage: &common.BaseStage{}}
	})
}

func (f *ExpressionProcessor) Init(stageContext api.StageContext) []validation.Issue {
	return f.BaseStage.Init(stageContext)
}

func (f *ExpressionProcessor) Process(batch api.Batch, batchMaker api.BatchMaker) error {
	for _, record := range batch.GetRecords() {
		recordContext := context.WithValue(context.Background(), el.RecordContextVar, record)
		var err error
		var evaluatedRes interface{}
		for _, exprProcessorConfig := range f.ExpressionProcessorConfigs {
			evaluatedRes, err = f.GetStageContext().Evaluate(exprProcessorConfig.Expression, EXPRESSION, recordContext)
			if err == nil {
				var evalField *api.Field
				if evalField, err = api.CreateFieldFromSDCField(evaluatedRes); err == nil {
					record.SetField(exprProcessorConfig.FieldToSet, evalField)
				}
			}
			if err != nil {
				err = errors.New(
					fmt.Sprintf(
						"Error when setting field '%s' with expression : '%s'. Reason : '%s'",
						exprProcessorConfig.FieldToSet, exprProcessorConfig.Expression, err.Error()))
				break
			}
		}

		if err == nil {
			for _, headerAttrConfig := range f.HeaderAttributeConfigs {
				evaluatedRes, err = f.GetStageContext().Evaluate(headerAttrConfig.Expression, EXPRESSION, recordContext)
				if err == nil {
					record.GetHeader().SetAttribute(headerAttrConfig.AttributeToSet, evaluatedRes.(string))
				} else {
					err = errors.New(
						fmt.Sprintf(
							"Error when setting attribute '%s' with expression : '%s'. Reason : '%s'",
							headerAttrConfig.AttributeToSet, headerAttrConfig.Expression, err.Error()))
					break
				}
			}
		}

		if err != nil {
			log.WithError(err).Error("Error evaluating record")
			f.GetStageContext().ToError(err, record)
		} else {
			batchMaker.AddRecord(record)
		}
	}
	return nil
}
