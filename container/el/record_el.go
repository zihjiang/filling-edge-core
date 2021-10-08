
package el

import (
	"context"
	"errors"
	"fmt"
	"github.com/madhukard/govaluate"
	"github.com/spf13/cast"
	"datacollector-edge/api"
	"datacollector-edge/container/util"
)

const (
	RecordContextVar = "record"
)

type RecordEL struct {
	Context context.Context
}

func (r *RecordEL) GetType(args ...interface{}) (interface{}, error) {
	if len(args) < 1 {
		return "", errors.New(
			fmt.Sprintf("The function 'record:type' requires 1 arguments but was passed %d", len(args)),
		)
	}

	fieldPath := cast.ToString(args[0])

	record, err := r.getRecordInContext()
	if err != nil {
		return nil, err
	}

	field, err := record.Get(fieldPath)
	if err != nil {
		return nil, err
	}

	return field.Type, nil
}

func (r *RecordEL) GetValue(args ...interface{}) (interface{}, error) {
	if len(args) < 1 {
		return "", errors.New(
			fmt.Sprintf("The function 'record:value' requires 1 arguments but was passed %d", len(args)),
		)
	}

	fieldPath := cast.ToString(args[0])

	record, err := r.getRecordInContext()
	if err != nil {
		return nil, err
	}

	field, err := record.Get(fieldPath)
	if err != nil {
		return nil, err
	}

	// govaluate library only officially deals with four types; float64, bool, string, and arrays.
	// https://github.com/Knetic/govaluate/blob/master/MANUAL.md
	// so cast all numeric values to float64
	return util.CastToFloat64(field.Value), nil
}

func (r *RecordEL) GetValueOrDefault(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return "", errors.New(
			fmt.Sprintf("The function 'record:valueOrDefault' requires 2 arguments but was passed %d",
				len(args),
			),
		)
	}

	fieldPath := cast.ToString(args[0])
	defaultValue := args[1]

	record, err := r.getRecordInContext()
	if err != nil {
		return nil, err
	}

	field, err := record.Get(fieldPath)
	if err != nil {
		return nil, err
	}

	if field.Value != nil {
		return field.Value, nil
	}
	return defaultValue, nil
}

func (r *RecordEL) GetAttribute(args ...interface{}) (interface{}, error) {
	if len(args) < 1 {
		return "", errors.New(
			fmt.Sprintf("The function 'record:attribute' requires 1 arguments but was passed %d", len(args)),
		)
	}

	attributeName := cast.ToString(args[0])

	record, err := r.getRecordInContext()
	if err != nil {
		return nil, err
	}

	attributeValue := record.GetHeader().GetAttribute(attributeName)
	return attributeValue, nil
}

func (r *RecordEL) GetAttributeOrDefault(args ...interface{}) (interface{}, error) {
	if len(args) < 2 {
		return "", errors.New(
			fmt.Sprintf("The function 'record:attributeOrDefault' requires 2 arguments but was passed %d",
				len(args),
			),
		)
	}

	attributeName := cast.ToString(args[0])
	defaultValue := args[1]

	record, err := r.getRecordInContext()
	if err != nil {
		return nil, err
	}

	attributeValue := record.GetHeader().GetAttribute(attributeName)

	if attributeValue != nil {
		return attributeValue, nil
	}

	return defaultValue, nil
}

func (r *RecordEL) Exists(args ...interface{}) (interface{}, error) {
	if len(args) < 1 {
		return false, errors.New(
			fmt.Sprintf("The function 'record:exists' requires 1 arguments but was passed %d",
				len(args),
			),
		)
	}

	fieldPath := cast.ToString(args[0])

	record, err := r.getRecordInContext()
	if err != nil {
		return nil, err
	}

	field, err := record.Get(fieldPath)

	if field != nil && len(field.Type) > 0 {
		return true, nil
	}
	return false, nil
}

func (r *RecordEL) getRecordInContext() (api.Record, error) {
	if r.Context != nil {
		record := r.Context.Value(RecordContextVar).(api.Record)
		if record != nil {
			return record, nil
		}
	}
	return nil, errors.New("record context is not set")
}

func (r *RecordEL) GetELFunctionDefinitions() map[string]govaluate.ExpressionFunction {
	functions := map[string]govaluate.ExpressionFunction{
		"record:type":               r.GetType,
		"record:value":              r.GetValue,
		"record:valueOrDefault":     r.GetValueOrDefault,
		"record:attribute":          r.GetAttribute,
		"record:attributeOrDefault": r.GetAttributeOrDefault,
		"record:exists":             r.Exists,
		// TODO: SDCE-63 Add remaining record el functions
	}
	return functions
}
