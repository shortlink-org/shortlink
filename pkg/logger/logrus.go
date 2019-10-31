package logger

import "fmt"

type logrusLogger struct{} // nolint unused

func (log *logrusLogger) Init() error {
	return nil
}

func (log *logrusLogger) Info(msg string, fields ...Fields) {
	fmt.Println(msg)
}

func (log *logrusLogger) Warn(msg string, fields ...Fields) {
	fmt.Println(msg)
}

func (log *logrusLogger) Error(msg string, fields ...Fields) {
	fmt.Println(msg)
}

func (log *logrusLogger) Fatal(msg string, fields ...Fields) {
	fmt.Println(msg)
}

func (log *logrusLogger) Close() {}
