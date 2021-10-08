
package el

import (
	"reflect"
	"testing"
)

func TestMapListEL_EmptyMap(t *testing.T) {
	evaluator, _ := NewEvaluator(
		"emptyMap",
		nil,
		[]Definitions{&MapListEL{}},
	)
	result, err := evaluator.Evaluate("${emptyMap()")
	if err != nil {
		t.Errorf("Failed to create Map : %s", err)
	}

	if result == nil {
		t.Error("Failed to create Map : result is nil")
	}

	if reflect.ValueOf(result).Kind() != reflect.Map {
		t.Errorf("Failed to create Map, returned value type : %s", reflect.ValueOf(result).Kind())
	}
}

func TestMapListEL_EmptyList(t *testing.T) {
	evaluator, _ := NewEvaluator(
		"emptyList",
		nil,
		[]Definitions{&MapListEL{}},
	)
	result, err := evaluator.Evaluate("${emptyList()")
	if err != nil {
		t.Errorf("Failed to create List : %s", err)
	}

	if result == nil {
		t.Error("Failed to create List : result is nil")
	}

	if reflect.ValueOf(result).Kind() != reflect.Slice {
		t.Errorf("Failed to create List, returned value type : %s", reflect.ValueOf(result).Kind())
	}
}

func TestMapListEL(test *testing.T) {
	evaluationTests := []EvaluationTest{
		{
			Name:       "Test function size 1",
			Expression: "${size(MAP_PARAM)}",
			Parameters: map[string]interface{}{
				"MAP_PARAM": map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			Expected: 2,
		},
		{
			Name:       "Test function size - Error 1",
			Expression: "${size()}",
			Expected:   "The function 'size' requires 1 arguments but was passed 0",
			ErrorCase:  true,
		},
		{
			Name:       "Test function size - Error 1",
			Expression: "${size(STRING_PARAM)}",
			Parameters: map[string]interface{}{
				"STRING_PARAM": "stringValue",
			},
			Expected:  "Unsupported Field Type: string for EL function size()",
			ErrorCase: true,
		},

		{
			Name:       "Test function isEmptyMap 1",
			Expression: "${isEmptyMap(MAP_PARAM)}",
			Parameters: map[string]interface{}{
				"MAP_PARAM": map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			Expected: false,
		},
		{
			Name:       "Test function isEmptyMap 2",
			Expression: "${isEmptyMap(MAP_PARAM)}",
			Parameters: map[string]interface{}{
				"MAP_PARAM": map[string]interface{}{},
			},
			Expected: true,
		},
		{
			Name:       "Test function isEmptyMap - Error 1",
			Expression: "${isEmptyMap()}",
			Expected:   "The function 'isEmptyMap' requires 1 arguments but was passed 0",
			ErrorCase:  true,
		},
		{
			Name:       "Test function size - Error 1",
			Expression: "${isEmptyMap(STRING_PARAM)}",
			Parameters: map[string]interface{}{
				"STRING_PARAM": "stringValue",
			},
			Expected:  "Unsupported Field Type: string for EL function isEmptyMap()",
			ErrorCase: true,
		},

		{
			Name:       "Test function length 1",
			Expression: "${length(LIST_PARAM)}",
			Parameters: map[string]interface{}{
				"LIST_PARAM": []string{
					"value1",
					"value2",
				},
			},
			Expected: 2,
		},
		{
			Name:       "Test function length - Error 1",
			Expression: "${length()}",
			Expected:   "The function 'length' requires 1 arguments but was passed 0",
			ErrorCase:  true,
		},
		{
			Name:       "Test function length - Error 1",
			Expression: "${length(STRING_PARAM)}",
			Parameters: map[string]interface{}{
				"STRING_PARAM": "stringValue",
			},
			Expected:  "Unsupported Field Type: string for EL function length()",
			ErrorCase: true,
		},

		{
			Name:       "Test function isEmptyList 1",
			Expression: "${isEmptyList(LIST_PARAM)}",
			Parameters: map[string]interface{}{
				"LIST_PARAM": []string{
					"value1",
					"value2",
				},
			},
			Expected: false,
		},
		{
			Name:       "Test function isEmptyList 2",
			Expression: "${isEmptyList(LIST_PARAM)}",
			Parameters: map[string]interface{}{
				"LIST_PARAM": []string{},
			},
			Expected: true,
		},
		{
			Name:       "Test function isEmptyList - Error 1",
			Expression: "${isEmptyList()}",
			Expected:   "The function 'isEmptyList' requires 1 arguments but was passed 0",
			ErrorCase:  true,
		},
		{
			Name:       "Test function size - Error 1",
			Expression: "${isEmptyList(STRING_PARAM)}",
			Parameters: map[string]interface{}{
				"STRING_PARAM": "stringValue",
			},
			Expected:  "Unsupported Field Type: string for EL function isEmptyList()",
			ErrorCase: true,
		},
	}
	RunEvaluationTests(evaluationTests, []Definitions{&MapListEL{}}, test)
}
