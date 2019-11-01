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
	init(Configuration) error

	Info(msg string, fields ...Fields)

	Warn(msg string, fields ...Fields)

	Error(msg string, fields ...Fields)

	Fatal(msg string, fields ...Fields)

	SetConfig(Configuration) error

	Close()
}

// level config
const (
	LOG_LEVEL_INFO  int = iota // nolint unused
	LOG_LEVEL_DEBUG            // nolint unused
)

// Configuration - options for logger
type Configuration struct { // nolint unused
	Level int
}
