package util

import "reflect"

var StringType = reflect.TypeOf("")
var IntType = reflect.TypeOf(int64(0))
var FloatType = reflect.TypeOf(float64(0.0))
var NullType = reflect.TypeOf(nil)
var ListType = reflect.TypeOf([]any{})
var DictType = reflect.TypeOf(map[string]any{})
