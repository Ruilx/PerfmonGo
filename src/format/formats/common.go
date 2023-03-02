package formats

import (
	"PerfmonGo/format"
	"strconv"
)

func init() {
	formatInstance := format.Instance()
	formatInstance.RegisterFormat("toInt", ToInt)
}

func ToInt(value string) (any, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	return i, err
}
