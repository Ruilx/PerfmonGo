package formats

import (
	"PerfmonGo/format"
	"errors"
	"strconv"
	"strings"
)

var IsNullValue = errors.New("is Null")

func init() {
	formatInstance := format.Instance()
	formatInstance.RegisterFormat("toInt", ToInt)
	formatInstance.RegisterFormat("toIntOrNull", ToIntOrNull)
	formatInstance.RegisterFormat("toFloat", ToFloat)
	formatInstance.RegisterFormat("toFloatOrNull", ToFloatOrNull)
	formatInstance.RegisterFormat("trim", Trim)
}

func ToInt(value string) (any, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	return i, err
}

func ToIntOrNull(value string) (any, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, IsNullValue
	}
	return i, err
}

func ToFloat(value string) (any, error) {
	i, err := strconv.ParseFloat(value, 64)
	return i, err
}

func ToFloatOrNull(value string) (any, error) {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0, IsNullValue
	}
	return f, err
}

func Trim(value string) (any, error) {
	return strings.TrimSpace(value), nil
}
