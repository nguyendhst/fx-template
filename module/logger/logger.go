package logger

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Fatal(msg string)
	Error(msg string)
	LogRequest(...interface{})
}
