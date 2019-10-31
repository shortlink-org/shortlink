package logger

type key int

const (
	keyLogger key = iota
)

const (
	Zap    int = iota // nolint unused
	Logrus            // nolint unused
)

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{} //nolint unused

//Logger is our contract for the logger
type Logger interface { //nolint unused
	Init() error

	Info(msg string, fields ...Fields)

	Warn(msg string, fields ...Fields)

	Error(msg string, fields ...Fields)

	Fatal(msg string, fields ...Fields)

	Close()
}

// Configuration - options for logger
type Configuration struct{} // nolint unused
