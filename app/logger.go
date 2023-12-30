package app

import (
	"github.com/fatih/color"
	"log"
)

type LoggerInterface interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type LogLevel int

const (
	None  LogLevel = iota
	Error LogLevel = iota
	Warn
	Info
	Debug
)

func NewLogger(logLevel LogLevel, nestedLogger *log.Logger) *Logger {
	return &Logger{
		nestedLogger,
		logLevel,
	}
}

type Logger struct {
	*log.Logger
	level LogLevel
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) Level() LogLevel {
	return l.level
}

func (l *Logger) Error(args ...interface{}) {
	if l.level < Error {
		return
	}
	l.Logger.Println(color.RedString("[ERROR] "), args)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.level < Error {
		return
	}
	l.Logger.Printf(color.RedString("[ERROR] ")+format, args)
}

func (l *Logger) Warn(args ...interface{}) {
	if l.level < Warn {
		return
	}
	l.Logger.Println(color.YellowString("[WARN] "), args)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.level < Warn {
		return
	}
	l.Logger.Printf(color.YellowString("[WARN] ")+format, args)
}

func (l *Logger) Info(args ...interface{}) {
	if l.level < Info {
		return
	}
	l.Logger.Println("[INFO] ", args)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level < Info {
		return
	}
	l.Logger.Printf("[INFO] "+format, args)
}

func (l *Logger) Debug(args ...interface{}) {
	if l.level < Debug {
		return
	}
	l.Logger.Println(color.BlueString("[DEBUG] "), args)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level < Debug {
		return
	}
	l.Logger.Printf(color.BlueString("[DEBUG] ")+format, args)
}
