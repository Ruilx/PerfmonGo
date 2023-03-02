package format

import (
	_ "PerfmonGo/format/formats"
	"PerfmonGo/util/logger"
	"os"
)

type FormatFunc func(value string) (any, error)

type Format struct {
	instance *Format
	formats  map[string]FormatFunc
	log      *logger.Logger
}

var instance *Format

func (c *Format) RegisterFormat(name string, f FormatFunc) {
	if _, ok := c.formats[name]; ok {
		c.log.Warningf("name '%s' has already registered for format, will replaced to the new one.", name)
	}
	c.formats[name] = f
}

func Instance() *Format {
	if instance == nil {
		instance := &Format{
			formats: make(map[string]FormatFunc),
			log:     logger.New("format", os.Stderr, logger.Debug),
		}
		instance.instance = instance
	}
	return instance
}
