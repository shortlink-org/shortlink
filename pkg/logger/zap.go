package logger

import "fmt"

type zapLogger struct{} // nolint unused

func (log *zapLogger) Info(msg string, fields ...Field) {
	fmt.Println(msg)
}

func (log *zapLogger) Warn(msg string, fields ...Field) {
	fmt.Println(msg)
}

func (log *zapLogger) Error(msg string, fields ...Field) {
	fmt.Println(msg)
}

func (log *zapLogger) Fatal(msg string, fields ...Field) {
	fmt.Println(msg)
}

func (log *zapLogger) Close() {}
