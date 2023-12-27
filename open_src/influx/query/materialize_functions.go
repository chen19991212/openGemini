package query

import (
	"sync"

	"github.com/openGemini/openGemini/open_src/influx/influxql"
)

const (
	STRING = iota
	MATH
	AGG_NORMAL
	AGG_SLICE
)

type BaseInfo struct {
	FuncType int
}

func (b *BaseInfo) GetFuncType() int {
	return b.FuncType
}

type MaterializeFunc interface {
	CompileFunc(expr *influxql.Call) error
	CallTypeFunc(name string, args []influxql.DataType) (influxql.DataType, error)
	CallFunc(name string, args []interface{}) (interface{}, bool)
	GetFuncType() int
}

type AggregateFunc interface {
	CompileFunc(expr *influxql.Call) error
	CallTypeFunc(name string, args []influxql.DataType) (influxql.DataType, error)
	GetFuncType() int
	OnlySelectors() bool
}

type FunctionFactory struct {
	materialize map[string]MaterializeFunc
	aggregate   map[string]AggregateFunc
}

func NewFunctionFactory() *FunctionFactory {
	return &FunctionFactory{
		materialize: make(map[string]MaterializeFunc),
		aggregate:   make(map[string]AggregateFunc),
	}
}

func RegistryMaterializeFunction(name string, function MaterializeFunc) bool {
	factory := GetFunctionFactoryInstance()
	_, ok := factory.FindMaterFunc(name)

	if ok {
		return ok
	}

	factory.AddMaterFunc(name, function)
	return ok
}

func (r *FunctionFactory) AddMaterFunc(name string, function MaterializeFunc) {
	r.materialize[name] = function
}

func (r *FunctionFactory) FindMaterFunc(name string) (MaterializeFunc, bool) {
	function, ok := r.materialize[name]
	return function, ok
}

func RegistryAggregateFunction(name string, function AggregateFunc) bool {
	factory := GetFunctionFactoryInstance()
	_, ok := factory.FindAggFunc(name)

	if ok {
		return ok
	}

	factory.AddAggFunc(name, function)
	return ok
}

func (r *FunctionFactory) AddAggFunc(name string, function AggregateFunc) {
	r.aggregate[name] = function
}

func (r *FunctionFactory) FindAggFunc(name string) (AggregateFunc, bool) {
	function, ok := r.aggregate[name]
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
