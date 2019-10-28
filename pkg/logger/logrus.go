package logger

type logrusLogger struct{}

func (log *logrusLogger) Info(msg string, fields ...Field) {}

func (log *logrusLogger) Warn(msg string, fields ...Field) {}

func (log *logrusLogger) Error(msg string, fields ...Field) {}

func (log *logrusLogger) Fatal(msg string, fields ...Field) {}

func (log *logrusLogger) Close() {}
