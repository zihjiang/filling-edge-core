
package el

import (
	"context"
	"strings"
)

const (
	NAMESPACE_FN_SEPARATOR = ":"
)

func IsElString(configValue string) bool {
	return strings.HasPrefix(configValue, PARAMETER_PREFIX) &&
		strings.HasSuffix(configValue, PARAMETER_SUFFIX)
}

func Evaluate(
	value string,
	configName string,
	parameters map[string]interface{},
	elContext context.Context,
) (interface{}, error) {
	evaluator, _ := NewEvaluator(
		configName,
		parameters,
		[]Definitions{
			&StringEL{},
			&MathEL{},
			&MapListEL{},
			&PipelineEL{Context: elContext},
			&JobEL{Context: elContext},
			&SdcEL{},
		},
	)
	return evaluator.Evaluate(value)
}
