package query

import "github.com/openGemini/openGemini/open_src/influx/influxql"

func init() {
	var (
		_ = RegistryAggregateFunction("last", &LastFunc{
			BaseInfo:      BaseInfo{FuncType: AGG_NORMAL},
			onlySelectors: false,
		})
		_ = RegistryAggregateFunction("max", &LastFunc{
			BaseInfo:      BaseInfo{FuncType: AGG_NORMAL},
			onlySelectors: false,
		})
	)
}

// last
type LastFunc struct {
	onlySelectors bool
	BaseInfo
}

func (l *LastFunc) OnlySelectors() bool {
	return l.onlySelectors
}

func (l *LastFunc) CompileFunc(expr *influxql.Call) error {
	return nil
}

func (l *LastFunc) CallTypeFunc(name string, args []influxql.DataType) (influxql.DataType, error) {
	return args[0], nil
}

// max
type MaxFunc struct {
	BaseInfo
	onlySelectors bool
}

func (m *MaxFunc) OnlySelectors() bool {
	return m.onlySelectors
}

func (m *MaxFunc) CompileFunc(expr *influxql.Call) error {
	return nil
}

func (m *MaxFunc) CallTypeFunc(name string, args []influxql.DataType) (influxql.DataType, error) {
	return args[0], nil
}

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
	if aggFunc, ok := GetFunctionFactoryInstance().FindAggFunc(name); ok {
		return aggFunc.CallTypeFunc(name, args)
	}
	return influxql.Unknown, nil
}
