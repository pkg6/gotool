package i

import "io"

type ILogger interface {
	SetOutput(w io.Writer)
	SetPrefix(prefix string)
	SetFlags(flag int)
	Emergency(message string, context any)
	Alert(message string, context any)
	Critical(message string, context any)
	Error(message string, context any)
	Warning(message string, context any)
	Notice(message string, context any)
	Info(message string, context any)
	Debug(message string, context any)
	Log(level, message string, context any)
}
