package logger

type key int

const (
	keyLogger key = iota
)

const (
	Zap    int = iota // nolint unused
	Logrus            // nolint unused
)

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{} //nolint unused

// Logger is our contract for the logger
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
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PANIC_LEVEL int = iota // nolint unused
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FATAL_LEVEL // nolint unused
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ERROR_LEVEL // nolint unused
	// WarnLevel level. Non-critical entries that deserve eyes.
	WARN_LEVEL // nolint unused
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	INFO_LEVEL // nolint unused
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DEBUG_LEVEL // nolint unused
)

// Configuration - options for logger
type Configuration struct { // nolint unused
	Level int
}
