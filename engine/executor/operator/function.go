package operator

import "github.com/openGemini/openGemini/open_src/influx/influxql"

// type CompileFunc func(expr *influxql.Call) error
type CallTypeFunc func(name string, args []influxql.DataType) (influxql.DataType, error)
type CallFunc func(name string, args []interface{}) (interface{}, bool)

type Function struct {
	Name         string
	CallTypeFunc CallTypeFunc
	CallFunc     CallFunc
}
