package executor

import (
	"github.com/openGemini/openGemini/engine/hybridqp"
	"github.com/openGemini/openGemini/open_src/influx/influxql"
)

type NewRoutinue func(inRowDataType, outRowDataType hybridqp.RowDataType, opt hybridqp.ExprOptions, isSingleCall bool, auxProcessor []*AuxProcessor, interval hybridqp.Interval) (Routine, error)
type CallTypeFunc func(name string, args []influxql.DataType) (influxql.DataType, error)

type Operator struct {
	Name         string
	CallTypeFunc CallTypeFunc
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
