package query

import (
	"sync"

	"github.com/openGemini/openGemini/open_src/influx/influxql"
)

type MaterializeFunc interface {
	CompileFunc(expr *influxql.Call) error
	CallTypeFunc(name string, args []influxql.DataType) (influxql.DataType, error)
	CallFunc(name string, args []interface{}) (interface{}, bool)
}

func RegistryFunction(name string, function MaterializeFunc) bool {
	factory := GetFunctionFactoryInstance()
	_, ok := factory.Find(name)

	if ok {
		return ok
	}

	factory.Add(name, function)
	return ok
}

type FunctionFactory struct {
	functions map[string]MaterializeFunc
}

func NewFunctionFactory() *FunctionFactory {
	return &FunctionFactory{
		functions: make(map[string]MaterializeFunc),
	}
}

func (r *FunctionFactory) Add(name string, function MaterializeFunc) {
	r.functions[name] = function
}

func (r *FunctionFactory) Find(name string) (MaterializeFunc, bool) {
	function, ok := r.functions[name]
	return function, ok
}

var instance *FunctionFactory
var once sync.Once

func GetFunctionFactoryInstance() *FunctionFactory {
	once.Do(func() {
		instance = NewFunctionFactory()
	})
	return instance
}
