package logger

import (
	"context"
	"io"
)

var l = New()

func Reset() {
	l = New()
}
func SetOutput(w io.Writer) {
	l.SetOutput(w)
}
func SetPrefix(s string) {
	l.SetPrefix(s)
}
func SetContext(ctx context.Context) {
	l.SetContext(ctx)
}
func Emergency(message string, context any) {
	l.Emergency(message, context)
}
func Alert(message string, context any) {
	l.Alert(message, context)
}
func Critical(message string, context any) {
	l.Critical(message, context)
}
func Error(message string, context any) {
	l.Error(message, context)
}
func Warning(message string, context any) {
	l.Warning(message, context)
}
func Notice(message string, context any) {
	l.Notice(message, context)
}
func Info(message string, context any) {
	l.Info(message, context)
}
func Debug(message string, context any) {
	l.Debug(message, context)
}
func Log(level, message string, context any) {
	l.Log(level, message, context)
}
