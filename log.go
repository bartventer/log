// Package log provides a structured logger with context support.
package log

import (
	"log/slog"
	"os"
	"sync"
	"sync/atomic"

	"github.com/charmbracelet/log"
)

// Logger is an alias for [slog.Logger].
type Logger = slog.Logger

var (
	defaultStylesOnce struct {
		sync.Once
		s *log.Styles
	}

	defaultOnce struct {
		sync.Once
		l atomic.Pointer[slog.Logger]
	}
)

// DefaultStyles returns the default styles.
// It applies custom styles to the [log.DefaultStyles].
func DefaultStyles() *Styles {
	defaultStylesOnce.Do(func() {
		defaultStylesOnce.s = log.DefaultStyles()

		// ========= Custom styles =========
		for _, s := range []struct {
			Level    Level
			MaxWidth int // to avoid truncation
			// ... more custom styles
		}{
			{DebugLevel, 5},
			{InfoLevel, 4},
			{WarnLevel, 4},
			{ErrorLevel, 5},
			{FatalLevel, 5},
		} {
			defaultStylesOnce.s.Levels[s.Level] = defaultStylesOnce.s.Levels[s.Level].
				MaxWidth(s.MaxWidth)
		}
	})

	return defaultStylesOnce.s
}

// Default returns the default logger.
func Default() *slog.Logger {
	defaultOnce.Do(func() {
		if defaultOnce.l.Load() != nil {
			return
		}
		defaultOnce.l.Store(New(AsDefault()))
	})
	return defaultOnce.l.Load()
}

//  +------------------------------------------------------------+
//  | Helpers 												 	 |
//  +------------------------------------------------------------+

// handler returns the default logger's handler.
func handler() *log.Logger {
	return loggerHandler(Default())
}

// loggerHandler returns the logger's handler.
func loggerHandler(l *slog.Logger) *log.Logger {
	return l.Handler().(*log.Logger)
}

// DefaultOptions returns the default options.
func DefaultOptions() *Options {
	return &Options{
		LogOptions: &LogOptions{
			Level: log.InfoLevel,
		},
		Writer: os.Stderr,
		Styles: DefaultStyles(),
	}
}

//  +------------------------------------------------------------+
//  | Loggers 												 	 |
//  +------------------------------------------------------------+

// New creates a new logger with the given options.
func New(opts ...Option) *slog.Logger {
	o := DefaultOptions()
	o.Apply(opts...)

	handler := log.NewWithOptions(o.Writer, *o.LogOptions)

	if o.Styles != nil {
		handler.SetStyles(o.Styles)
	}

	l := slog.New(handler)

	if o.Default {
		log.SetDefault(handler)
		slog.SetDefault(l)
		defaultOnce.l.Store(l)
	}

	return l
}

//  +------------------------------------------------------------+
//  | Logging 												 	 |
//  +------------------------------------------------------------+

// Debug logs a message with level Debug.
func Debug(msg string, args ...any) {
	handler().Debug(msg, args...)
}

// Debugf logs a formatted message with level Debug.
func Debugf(format string, args ...any) {
	handler().Debugf(format, args...)
}

// Info logs a message with level Info.
func Info(msg string, args ...any) {
	handler().Info(msg, args...)
}

// Infof logs a formatted message with level Info.
func Infof(format string, args ...any) {
	handler().Infof(format, args...)
}

// Warn logs a message with level Warn.
func Warn(msg string, args ...any) {
	handler().Warn(msg, args...)
}

// Warnf logs a formatted message with level Warn.
func Warnf(format string, args ...any) {
	handler().Warnf(format, args...)
}

// Error logs a message with level Error.
func Error(msg string, args ...any) {
	handler().Error(msg, args...)
}

// Errorf logs a formatted message with level Error.
func Errorf(format string, args ...any) {
	handler().Errorf(format, args...)
}

// Fatal logs a message with level Fatal and exits with status code 1.
func Fatal(msg any, keyvals ...any) {
	handler().Fatal(msg, keyvals...)
}

// Fatalf logs a formatted message with level Fatal and exits with status code 1.
func Fatalf(format string, args ...any) {
	handler().Fatalf(format, args...)
}

// Print logs a message with no level.
func Print(msg string, args ...any) {
	handler().Print(msg, args...)
}

// Log logs a message with the given level.
func Log(level Level, msg string, args ...any) {
	handler().Log(level, msg, args...)
}

// Logf logs a formatted message with the given level.
func Logf(level Level, format string, args ...any) {
	handler().Logf(level, format, args...)
}
