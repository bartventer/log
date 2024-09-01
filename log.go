// Package log provides a structured logger with context support.
package log

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync"
	"sync/atomic"

	"github.com/charmbracelet/log"
)

// DefaultStyles returns the default styles.
// It applies custom styles to the [log.DefaultStyles].
func DefaultStyles() *Styles {
	defaultStylesOnce.Do(func() {
		defaultStylesOnce.s = log.DefaultStyles()

		// ========= Level styles =========
		// Disable truncation for all levels
		defaultStylesOnce.s.Levels[log.DebugLevel] = defaultStylesOnce.s.Levels[log.DebugLevel].MaxWidth(5)
		defaultStylesOnce.s.Levels[log.InfoLevel] = defaultStylesOnce.s.Levels[log.InfoLevel].MaxWidth(4)
		defaultStylesOnce.s.Levels[log.WarnLevel] = defaultStylesOnce.s.Levels[log.WarnLevel].MaxWidth(4)
		defaultStylesOnce.s.Levels[log.ErrorLevel] = defaultStylesOnce.s.Levels[log.ErrorLevel].MaxWidth(5)
		defaultStylesOnce.s.Levels[log.FatalLevel] = defaultStylesOnce.s.Levels[log.FatalLevel].MaxWidth(5)
	})

	return defaultStylesOnce.s
}

var defaultStylesOnce struct {
	sync.Once
	s *log.Styles
}

var defaultOnce struct {
	sync.Once
	l atomic.Pointer[slog.Logger]
}

// Default returns the default logger.
func Default() *slog.Logger {
	defaultOnce.Do(func() {
		if defaultOnce.l.Load() != nil {
			return
		}
		defaultOnce.l.Store(New(WithDefault()))
	})
	return defaultOnce.l.Load()
}

// SetOutput sets the writer for the default logger.
func SetOutput(l *slog.Logger, w io.Writer) {
	l.Handler().(*log.Logger).SetOutput(w)
}

// Logger defines the [slog.Logger] interface.
type Logger interface {
	Debug(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	Enabled(ctx context.Context, level slog.Level) bool
	Error(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
	Handler() slog.Handler
	Info(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	Log(ctx context.Context, level slog.Level, msg string, args ...any)
	LogAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr)
	Warn(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	With(args ...any) *slog.Logger
	WithGroup(name string) *slog.Logger
}

// New creates a new logger with the given options.
func New(opts ...Option) *slog.Logger {
	o := &Options{
		LogOptions: &LogOptions{
			Level: log.InfoLevel,
		},
		Writer: os.Stderr,
		Styles: DefaultStyles(),
	}
	o.apply(opts...)

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
