package logging

import (
	"fmt"
	"io"
	"log"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

type logger struct {
	logLevel LogLevel
	logger   *log.Logger
}

func NewLogger(out io.Writer, logLevel LogLevel) Logger {
	return &logger{
		logLevel: logLevel,
		logger:   log.New(out, "", log.LstdFlags),
	}
}

func (l *logger) Debugf(format string, v ...interface{}) {
	if l.logLevel <= LogLevelDebug {
		l.logger.Output(2, fmt.Sprintf("[DEBUG] "+format, v...))
	}
}

func (l *logger) Infof(format string, v ...interface{}) {
	if l.logLevel <= LogLevelInfo {
		l.logger.Output(2, fmt.Sprintf("[INFO] "+format, v...))
	}
}

func (l *logger) Warnf(format string, v ...interface{}) {
	if l.logLevel <= LogLevelWarn {
		l.logger.Output(2, fmt.Sprintf("[WARN] "+format, v...))
	}
}

func (l *logger) Errorf(format string, v ...interface{}) {
	if l.logLevel <= LogLevelError {
		l.logger.Output(2, fmt.Sprintf("[ERROR] "+format, v...))
	}
}
