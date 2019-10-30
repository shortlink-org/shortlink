package logger

import "fmt"

type logrusLogger struct{} // nolint unused

func (log *logrusLogger) Info(msg string, fields ...Field) {
	fmt.Println(msg)
}

func (log *logrusLogger) Warn(msg string, fields ...Field) {
	fmt.Println(msg)
}

func (log *logrusLogger) Error(msg string, fields ...Field) {
	fmt.Println(msg)
}

func (log *logrusLogger) Fatal(msg string, fields ...Field) {
	fmt.Println(msg)
}

func (log *logrusLogger) Close() {}
