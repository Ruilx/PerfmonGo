package util

import (
	"errors"
	"fmt"
	url2 "net/url"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const WiredTime = "2006-01-02 15:04:05.000"

func CheckKey(key string, cfg map[string]any, typ reflect.Type, cfgName string, canBeNil bool) (any, error) {
	if v, ok := cfg[key]; ok {
		rev := reflect.ValueOf(v)
		if canBeNil && rev.IsNil() {
			return nil, nil
		}
		if rev.CanConvert(typ) {
			val := rev.Convert(typ)
			return val, nil
		} else {
			return nil, errors.New(fmt.Sprintf("'%s' cannot convert key '%s' to '%s'", cfgName, key, typ.Name()))
		}
	} else {
		return nil, errors.New(fmt.Sprintf("'%s' item need key named '%s' with type '%s'", cfgName, key, typ.Name()))
	}
}

func CheckValueEnum(value any, valueMustInList []any, valueCanBeNone bool, valueName string) error {
	rev := reflect.ValueOf(value)
	if valueCanBeNone && rev.IsNil() {
		return nil
	} else {
		if !rev.IsNil() {
			for _, v := range valueMustInList {
				if v == value {
					return nil
				}
			}
			return errors.New(fmt.Sprintf("value '%s' not exist in list", value))
		} else {
			return errors.New(fmt.Sprintf("value '%s' is Nil", value))
		}
	}
}

func CheckValueEnumT[T comparable](value T, valueMustInList []T, valueCanBeNone bool, valueName string) error {
	if valueCanBeNone && value == nil {
		return nil
	} else {
		if value != nil {
			for _, v := range valueMustInList {
				if v == value {
					return nil
				}
			}
			return errors.New(fmt.Sprintf("value '%s' not exist in list", value))
		} else {
			return errors.New(fmt.Sprintf("value '%s' is Nil", value))
		}
	}
}

func CheckUrl(url string) error {
	urlT, err := url2.Parse(url)
	if err != nil {
		return err
	}
	scheme := strings.ToLower(urlT.Scheme)
	if scheme != "http" && scheme != "https" {
		return errors.New(fmt.Sprintf("server scheme only support 'HTTP' or 'HTTPS', but '%s' found", scheme))
	}
	return nil
}

func Now() string {
	return time.Now().Format(WiredTime)
}

func Timestamp() int64 {
	return time.Now().Unix()
}

func CpuCount() int {
	return runtime.NumCPU()
}
