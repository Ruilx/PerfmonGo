package main

import (
	"PerfmonGo/util/logger"
	"os"
)

func main() {
	//file, err := os.Create("log.log")
	//defer file.Close()
	file := os.Stderr
	var err error = nil
	if err != nil {
		panic("cannot open")
	}
	l := logger.New("a", file, logger.Debug)
	l.Debugf("%s", "Hello, world!")
	l.Debugf("%s", "Hello, world!")
}
