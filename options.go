package log

import (
	"io"
	"log/slog"
	"time"

	"github.com/charmbracelet/log"
)

// Type Aliases
type (
	LogOptions      = log.Options
	Level           = log.Level
	CallerFormatter = log.CallerFormatter
	Formatter       = log.Formatter
	Styles          = log.Styles
)

// Options is the logger options.
type Options struct {
	*LogOptions
	Writer  io.Writer   // Writer is the writer for the logger. Default is [os.Stderr].
	Styles  *log.Styles // Styles is the styles for the logger. Default is [DefaultStyles].
	Default bool        // Default is whether the logger is the default logger. Default is false.
}

func (o *Options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// Option is a logger option.
type Option func(*Options)

// Log Levels
const (
	DebugLevel Level = log.DebugLevel
	InfoLevel  Level = log.InfoLevel
	WarnLevel  Level = log.WarnLevel
	ErrorLevel Level = log.ErrorLevel
	FatalLevel Level = log.FatalLevel
)

// Caller Formatters
var (
	ShortCallerFormatter = log.ShortCallerFormatter
	LongCallerFormatter  = log.LongCallerFormatter
)

// Formatters
const (
	TextFormatter   Formatter = log.TextFormatter
	JSONFormatter   Formatter = log.JSONFormatter
	LogfmtFormatter Formatter = log.LogfmtFormatter
)

//  +------------------------------------------------------------+
//  | Option Functions 										 	 |
//  +------------------------------------------------------------+

// WithTimeFunction sets the time function option. Default is [time.Now].
func WithTimeFunction(f func(time.Time) time.Time) Option {
	return func(o *Options) {
		o.TimeFunction = f
	}
}

// WithTimeFormat sets the time format option. Default is [log.DefaultTimeFormat].
func WithTimeFormat(f string) Option {
	return func(o *Options) {
		o.TimeFormat = f
	}
}

// WithLevel sets the level option. Default is [log.InfoLevel].
func WithLevel(l Level) Option {
	return func(o *Options) {
		o.Level = l
	}
}

// WithPrefix sets the prefix option. Default is no prefix.
func WithPrefix(p string) Option {
	return func(o *Options) {
		o.Prefix = p
	}
}

// WithReportTimestamp sets the report timestamp option. Default is false.
func WithReportTimestamp(r bool) Option {
	return func(o *Options) {
		o.ReportTimestamp = r
	}
}

// WithReportCaller sets the report caller option. Default is false.
func WithReportCaller(r bool) Option {
	return func(o *Options) {
		o.ReportCaller = r
	}
}

// WithCallerFormatter sets the caller formatter option. Default is [log.ShortCallerFormatter].
func WithCallerFormatter(f CallerFormatter) Option {
	return func(o *Options) {
		o.CallerFormatter = f
	}
}

// WithFields sets the fields option. Default is no fields.
func WithFields(fields map[string]slog.Value) Option {
	return func(o *Options) {
		for k, v := range fields {
			o.Fields = append(o.Fields, k, v)
		}
	}
}

// WithFormatter sets the formatter option. Default is [TextFormatter].
func WithFormatter(f Formatter) Option {
	return func(o *Options) {
		o.Formatter = f
	}
}

// WithCallerOffset sets the caller offset option. Default is 0.
func WithCallerOffset(offset int) Option {
	return func(o *Options) {
		o.CallerOffset = offset
	}
}

// WithWriter sets the writer option. Default is [os.Stderr].
func WithWriter(w io.Writer) Option {
	return func(o *Options) {
		o.Writer = w
	}
}

// WithStyles sets the styles option. Default is [DefaultStyles].
func WithStyles(s *Styles) Option {
	return func(o *Options) {
		o.Styles = s
	}
}

// WithDefault sets the default option to true.
func WithDefault() Option {
	return func(o *Options) {
		o.Default = true
	}
}
