package logger

import "io"

var Default = New()

func Reset() {
	Default = New()
}
func SetOutput(w io.Writer) {
	Default.SetOutput(w)
}
func SetPrefix(s string) {
	Default.SetPrefix(s)
}
func Emergency(message string, context any) {
	Default.Emergency(message, context)
}
func Alert(message string, context any) {
	Default.Alert(message, context)
}
func Critical(message string, context any) {
	Default.Critical(message, context)
}
func Error(message string, context any) {
	Default.Error(message, context)
}
func Warning(message string, context any) {
	Default.Warning(message, context)
}
func Notice(message string, context any) {
	Default.Notice(message, context)
}
func Info(message string, context any) {
	Default.Info(message, context)
}
func Debug(message string, context any) {
	Default.Debug(message, context)
}
func Log(level, message string, context any) {
	Default.Log(level, message, context)
}
