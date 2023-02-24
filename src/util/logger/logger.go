package logger

import (
	"fmt"
	"io"
	"log"
	"path"
	"runtime"
	"time"
)

const (
	Debug = iota
	Log
	Info
	Warning
	Error
	Fatal
)

const WiredTime = "2006-01-02 15:04:05.000"

type Logger struct {
	logger *log.Logger
	device io.Writer
	level  int
	name   string
}

func funcInfo(skip int) (string, string, int) {
	pc := make([]uintptr, 1)
	ret := runtime.Callers(skip, pc)
	if ret == 0 {
		return "<NotValid>", "<NotValid>", 0
	}
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return f.Name(), file, line
}

func now() string {
	return time.Now().Format(WiredTime)
}

func levelName(level int) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Log:
		return "LOG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func prefix(level int, name string, funcName string, file string, line int) string {
	return fmt.Sprintf("[%s] %s(%s): %s:%d: %s:",
		now(),
		levelName(level),
		name,
		file, line, funcName)
}

func basename(filepath string) string {
	return path.Base(filepath)
}

func New(name string, device io.Writer, level int) *Logger {
	return &Logger{
		logger: log.New(device, "", log.Lmsgprefix),
		level:  level,
		name:   name,
	}
}

func (l *Logger) printLog(funcInfoSkip int, level int, v ...any) {
	funcName, file, line := funcInfo(funcInfoSkip)
	p := make([]any, 0, len(v)+1)
	p = append(p, prefix(level, l.name, funcName, basename(file), line))
	p = append(p, v...)

	l.logger.Println(p...)
}

func (l *Logger) printLogf(funcInfoSkip int, level int, format string, v ...any) {
	funcName, file, line := funcInfo(funcInfoSkip)
	p := make([]any, 0, len(v)+1)
	p = append(p, prefix(level, l.name, funcName, basename(file), line))
	p = append(p, fmt.Sprintf(format, v...))

	l.logger.Println(p...)
}

func (l *Logger) Debug(v ...any) {
	if l.level > Debug {
		return
	}
	l.printLog(4, Debug, v...)
}

func (l *Logger) Debugf(format string, v ...any) {
	if l.level > Debug {
		return
	}
	l.printLogf(4, Debug, format, v...)
}

func (l *Logger) Log(v ...any) {
	if l.level > Log {
		return
	}
	l.printLog(4, Log, v...)
}

func (l *Logger) Logf(format string, v ...any) {
	if l.level > Log {
		return
	}
	l.printLogf(4, Log, format, v...)
}

func (l *Logger) Info(v ...any) {
	if l.level > Info {
		return
	}
	l.printLog(4, Info, v...)
}

func (l *Logger) Infof(format string, v ...any) {
	if l.level > Info {
		return
	}
	l.printLogf(4, Info, format, v...)
}

func (l *Logger) Warning(v ...any) {
	if l.level > Warning {
		return
	}
	l.printLog(4, Warning, v...)
}

func (l *Logger) Warningf(format string, v ...any) {
	if l.level > Warning {
		return
	}
	l.printLogf(4, Warning, format, v...)
}

func (l *Logger) Error(v ...any) {
	if l.level > Error {
		return
	}
	l.printLog(4, Error, v...)
}

func (l *Logger) Errorf(format string, v ...any) {
	if l.level > Error {
		return
	}
	l.printLogf(4, Error, format, v...)
}

func (l *Logger) Fatal(v ...any) {
	l.printLog(4, Fatal, v...)
	l.logger.Fatal(v...)
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.printLogf(4, Fatal, format, v...)
	l.logger.Fatalf(format, v...)
}

func (l *Logger) Panic(v ...any) {
	l.logger.Panic(v...)
}

func (l *Logger) Panicf(format string, v ...any) {
	l.logger.Panicf(format, v...)
}
