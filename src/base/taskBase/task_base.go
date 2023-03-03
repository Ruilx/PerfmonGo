package taskBase

import (
	"PerfmonGo/util"
	"PerfmonGo/util/logger"
	"context"
	"errors"
	"os"
)

//const validExpectEnum = []string{"int", "intOrNull", "real", "realOrNull", "string", "stringOrNull", "null"}

const (
	TimeoutType_ContextBackground = iota
	TimeoutType_CustomType
)

type TaskBaseI interface {
	checkProcess()
	setup()
	run()
	join()
}

type TaskBase struct {
	validExpectEnum []string
	name            string
	config          map[string]any

	method  string
	format  []string
	expect  string
	timeout float64

	retry int

	params map[string]any
	value  any
	error  string

	context       context.Context
	contextCancel context.CancelFunc

	timeoutType int
	log         logger.Logger
}

func New(name string, config map[string]any) (*TaskBase, error) {
	t := &TaskBase{}
	t.buildValidExpectEnum()
	t.name = name
	t.config = config

	t.log = logger.New(name, os.Stderr, logger.Debug)

	t.timeoutType = TimeoutType_ContextBackground

	method, err := util.CheckKey("method", config, util.StringType, "task", false)
	if err != nil {
		return nil, err
	}
	t.method = method.(string)

	t.format = make([]string, 0, 1)
	format, err := util.CheckKey("format", config, util.StringType, "task", true)
	if err != nil {
		format, err := util.CheckKey("format", config, util.ListType, "task", true)
		if err != nil {
			return nil, errors.New("task.format cannot convert either string and list")
		}
		formatList := format.([]any)
		for _, f := range formatList {
			if fStr, ok := f.(string); ok {
				t.format = append(t.format, fStr)
			} else {
				t.log.Error("task.format list has a non-string type value and not inserted to format list.")
			}
		}
	} else {
		t.format = append(t.format, format)
	}

	expect, err := util.CheckKey("expect", config, util.StringType, "task", false)
	if err != nil {
		return nil, err
	}
	t.expect = expect.(string)

	timeout, err := util.CheckKey("timeout", config, util.FloatType, "task", false)
	if err != nil {
		return nil, err
	}
	t.timeout = timeout.(float64)

	retry, err := util.CheckKey("retry", config, util.IntType, "task", false)
	if err != nil {
		retry = 3
	}
	t.retry = retry

	err = util.CheckValueEnumT(t.expect, t.validExpectEnum, false, "expect")
	if err != nil {
		return nil, err
	}

	t.params = make(map[string]any)
	t.value = nil
	t.error = ""

	t.context, t.contextCancel = context.WithCancel(context.Background())

}

func (c *TaskBase) buildValidExpectEnum() {
	c.validExpectEnum = make([]string, 0, 7)
	c.validExpectEnum = append(c.validExpectEnum, "int", "intOrNull", "real", "realOrNull", "string", "stringOrNull", "null")
}
