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
type Field map[string]interface{} //nolint unused

//Logger is our contract for the logger
type Logger interface { //nolint unused
	Info(msg string, fields ...Field)

	Warn(msg string, fields ...Field)

	Error(msg string, fields ...Field)

	Fatal(msg string, fields ...Field)

	Close()
}

// Configuration - options for logger
type Configuration struct{} // nolint unused
