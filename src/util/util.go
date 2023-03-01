package util

import (
	"reflect"
)

func CheckKey(key string, cfg map[string]any, typ reflect.Type, cfgName string, canBeNil bool) (any, error) {
	if v, ok := cfg[key]; ok {
		rev := reflect.ValueOf(v)
		if canBeNil && rev.IsNil() {
			return nil, nil
		}
		if rev.CanConvert(typ) {

		}
	}
}
