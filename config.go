package framework

import "time"

var defaultConfig = Config{
	ROUTER: ConfigRouter{
		Port:       3000,
		Host:       "0.0.0.0",
		StaticDirs: map[string]string{},
	},
	LOGGER: ConfigLogger{
		TimestampFormat: time.RFC3339,
		LogLevel:        uint8(InfoLevel),
	},
	RENDERER: ConfigRenderer{
		ViewsDir: "views",
	},
}

type Config map[string]interface{}

type ConfigRouter struct {
	Port       int
	Host       string
	StaticDirs map[string]string
}

type ConfigLogger struct {
	TimestampFormat string
	LogLevel        uint8
}

type ConfigRenderer struct {
	ViewsDir string
}

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

// Level type
type Level uint8

// Convert the Level to a string. E.g. PanicLevel becomes "panic".
func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warning"
	case ErrorLevel:
		return "error"
	case FatalLevel:
		return "fatal"
	case PanicLevel:
		return "panic"
	}

	return "unknown"
}
