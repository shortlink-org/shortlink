package logger

type zapLogger struct{}

func (log *zapLogger) Info(msg string, fields ...Field) {}

func (log *zapLogger) Warn(msg string, fields ...Field) {}

func (log *zapLogger) Error(msg string, fields ...Field) {}

func (log *zapLogger) Fatal(msg string, fields ...Field) {}

func (log *zapLogger) Close() {}
