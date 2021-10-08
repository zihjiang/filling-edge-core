
package fieldremover

import (
	"fmt"
	"regexp"

	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/stagelibrary"
)

const (
	LIBRARY         = "streamsets-datacollector-basic-lib"
	STAGE_NAME      = "com_streamsets_pipeline_stage_processor_fieldfilter_FieldFilterDProcessor"
	KEEP            = "KEEP"
	REMOVE          = "REMOVE"
	REMOVE_NULL     = "REMOVE_NULL"
	FIELDS          = "fields"
	FILTEROPERATION = "filterOperation"
	VERSION         = 1
)

type FieldRemoverProcessor struct {
	*common.BaseStage
	Fields          []interface{} `ConfigDef:"type=LIST,required=true"`
	FilterOperation string        `ConfigDef:"type=STRING,required=true"`
	fieldList       []*regexp.Regexp
}

func init() {
	stagelibrary.SetCreator(LIBRARY, STAGE_NAME, func() api.Stage {
		return &FieldRemoverProcessor{BaseStage: &common.BaseStage{}}
	})
}

func (f *FieldRemoverProcessor) Init(stageContext api.StageContext) []validation.Issue {
	issues := f.BaseStage.Init(stageContext)

	f.fieldList = make([]*regexp.Regexp, len(f.Fields))
	for i, field := range f.Fields {
		fieldPath, ok := field.(string)
		if !ok {
			issues = append(issues, stageContext.CreateConfigIssue("Unexpected field list value"))
			return issues
		}

		if re, err := regexp.Compile(fieldPath); err != nil {
			issues = append(issues, stageContext.CreateConfigIssue("Field path %s cannot be compiled to a regular expression: %s", fieldPath, err.Error()))
		} else {
			f.fieldList[i] = re
		}
	}

	if f.FilterOperation != KEEP && f.FilterOperation != REMOVE && f.FilterOperation != REMOVE_NULL {
		issues = append(
			issues,
			stageContext.CreateConfigIssue("Unsupported field FilterOperation: "+f.FilterOperation),
		)
		return issues
	}
	return issues
}

func filterPaths(paths map[string]bool, patterns []*regexp.Regexp) map[string]bool {
	filtered := make(map[string]bool)

OUTER:
	for path := range paths {
		// ignore the empty string path
		if path == "" {
			continue
		}
		for _, pattern := range patterns {
			if pattern.MatchString(path) {
				filtered[path] = true
				continue OUTER
			}
		}
	}

	return filtered
}

func (f *FieldRemoverProcessor) Process(batch api.Batch, batchMaker api.BatchMaker) error {
	for _, record := range batch.GetRecords() {
		recordPaths := record.GetFieldPaths()
		filteredPaths := filterPaths(recordPaths, f.fieldList)
		var err error

		for path := range recordPaths {
			if path == "" { // ignore the empty field path
				continue
			}
			if _, ok := filteredPaths[path]; ok != (f.FilterOperation == KEEP) {
				var skip bool
				if f.FilterOperation == REMOVE_NULL {
					if field, err := record.Get(path); err == nil {
						skip = field.Value != "" // check value for "null"
					}
				}
				if !skip {
					_, err = record.Delete(path)
				}
				if err != nil {
					err = fmt.Errorf("Error removing field : %s. Reason : %s", path, err.Error())
					break
				}
			}
		}

		if err == nil {
			batchMaker.AddRecord(record)
		} else {
			f.GetStageContext().ToError(err, record)
		}
	}
	return nil
}

func (f *FieldRemoverProcessor) getParentFields(fieldPath string) []string {
	parentFields := []string{}
	for index := 0; index < len(fieldPath); {
		c := fieldPath[index]
		switch c {
		case '/':
			fallthrough
		case '[':
			parentFields = append(parentFields, fieldPath[:index])
			break
		default:
			break
		}
		index++
	}
	return parentFields
}
