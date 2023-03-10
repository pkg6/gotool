package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

func New() *Logger {
	return &Logger{l: log.Default()}
}

type Logger struct {
	l   *log.Logger
	Ctx context.Context
}

func (l *Logger) SetOutput(w io.Writer) {
	l.l.SetOutput(w)
}
func (l *Logger) SetPrefix(prefix string) {
	l.l.SetPrefix(prefix)
}
func (l *Logger) SetFlags(flag int) {
	l.l.SetFlags(flag)
}

func (l *Logger) SetContext(ctx context.Context) {
	l.Ctx = ctx
}
func (l *Logger) Emergency(message string, context any) {
	l.Log(EMERGENCY, message, context)
}

func (l *Logger) Alert(message string, context any) {
	l.Log(ALERT, message, context)
}

func (l *Logger) Critical(message string, context any) {
	l.Log(CRITICAL, message, context)
}

func (l *Logger) Error(message string, context any) {
	l.Log(ERROR, message, context)
}

func (l *Logger) Warning(message string, context any) {
	l.Log(WARNING, message, context)
}

func (l *Logger) Notice(message string, context any) {
	l.Log(NOTICE, message, context)
}

func (l *Logger) Info(message string, context any) {
	l.Log(INFO, message, context)
}
func (l *Logger) Debug(message string, context any) {
	l.Log(DEBUG, message, context)
}
func (l *Logger) Log(level, message string, context any) {
	s := level
	s += " " + message

	if context != nil {
		t := reflect.TypeOf(context)
		switch t.Kind() {
		case reflect.Struct, reflect.Slice, reflect.Func, reflect.Map, reflect.Interface:
			s += " " + fmt.Sprintf("%#v", context)
		}
	} else {
		s += " " + fmt.Sprintf("%v", context)
	}
	if l.Ctx != nil {
		s += " " + fmt.Sprintf("%#v", l.Ctx)
	}
	_ = l.l.Output(2, s)
	if level == ERROR {
		os.Exit(1)
	}
}
