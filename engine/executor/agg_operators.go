package executor

import (
	"github.com/openGemini/openGemini/engine/executor/operator"
	"github.com/openGemini/openGemini/engine/hybridqp"
	"github.com/openGemini/openGemini/open_src/influx/influxql"
)

type NewRoutinue func(inRowDataType, outRowDataType hybridqp.RowDataType, opt hybridqp.ExprOptions, isSingleCall bool, auxProcessor []*AuxProcessor, interval hybridqp.Interval) (Routine, error)

type Operator struct {
	Name         string
	CallTypeFunc operator.CallTypeFunc
	NewRoutinue  NewRoutinue
}

var AggOperator = map[string]*Operator{
	"last": {
		Name:         "last",
		CallTypeFunc: lastCallType,
		// NewRoutinue:  NewLastRoutineImpl,
	},
	"max": {
		Name:         "max",
		CallTypeFunc: maxCallType,
		NewRoutinue:  NewMaxRoutineImpl,
	},
	// "mean": {
	// 	Name:         "mean",
	// 	CallTypeFunc: meanCallType,
	// 	NewRoutinue: NewMean,
	// },
	"rate2": {
		Name:         "rate",
		CallTypeFunc: rateCallType,
		NewRoutinue:  NewRate2RoutineImpl,
	},
}

func lastCallType(name string, args []influxql.DataType) (influxql.DataType, error) {
	return args[0], nil
}

func maxCallType(name string, args []influxql.DataType) (influxql.DataType, error) {
	return args[0], nil
}

func rateCallType(name string, args []influxql.DataType) (influxql.DataType, error) {
	return influxql.Float, nil
}

// func CommonCompileFunc(expr *influxql.Call) error {
// 	if exp, got := 1, len(expr.Args); exp != got {
// 		return fmt.Errorf("invalid number of arguments for %s, expected %d, got %d", expr.Name, exp, got)
// 	}
// 	return nil
// }

// calltype
type CallTypeMapper struct{}

func (CallTypeMapper) MapType(measurement *influxql.Measurement, field string) influxql.DataType {
	return influxql.Unknown
}

func (CallTypeMapper) MapTypeBatch(measurement *influxql.Measurement, field map[string]*influxql.FieldNameSpace, schema *influxql.Schema) error {
	return nil
}

func (CallTypeMapper) CallType(name string, args []influxql.DataType) (influxql.DataType, error) {
	// If the function is not implemented by the embedded field mapper, then
	// see if we implement the function and return the type here.
	if _, ok := AggOperator[name]; ok {
		return AggOperator[name].CallTypeFunc(name, args)
	}
	// switch name {
	// case "mean":
	// 	return influxql.Float, nil
	// case "count":
	// 	return influxql.Integer, nil
	// case "min", "max", "sum", "first", "last":
	// 	// TODO(jsternberg): Verify the input type.
	// 	return args[0], nil
	// }
	return influxql.Unknown, nil
}
