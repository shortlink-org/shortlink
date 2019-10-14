package logger

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

//Logger is our contract for the logger
type Logger interface {
	Info(msg string, fields Fields)

	Warn(msg string, fields Fields)

	Error(msg string, fields Fields)

	Fatal(msg string, fields Fields)
}
