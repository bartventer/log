package log

import (
	"io"
	"log/slog"
	"time"

	"github.com/charmbracelet/log"
)

// Type aliases
type (
	LogOptions      = log.Options
	Level           = log.Level
	CallerFormatter = log.CallerFormatter
	Formatter       = log.Formatter
	Styles          = log.Styles
	TimeFunction    = func(time.Time) time.Time
)

// Levels
const (
	DebugLevel = log.DebugLevel
	InfoLevel  = log.InfoLevel
	WarnLevel  = log.WarnLevel
	ErrorLevel = log.ErrorLevel
	FatalLevel = log.FatalLevel
)

// Caller Formatters
var (
	ShortCallerFormatter = log.ShortCallerFormatter
	LongCallerFormatter  = log.LongCallerFormatter
)

// Formatters
const (
	TextFormatter   = log.TextFormatter
	JSONFormatter   = log.JSONFormatter
	LogfmtFormatter = log.LogfmtFormatter
)

//  +------------------------------------------------------------+
//  | Options 												 	 |
//  +------------------------------------------------------------+

// Options is the logger options.
type Options struct {
	*LogOptions
	Writer  io.Writer   // Writer is the writer for the logger. Default is [os.Stderr].
	Styles  *log.Styles // Styles is the styles for the logger. Default is [DefaultStyles].
	Default bool        // Default is whether the logger is the default logger. Default is false.
}

func (o *Options) Apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

// Option is a logger option.
type Option func(*Options)

// UseTimeFunction sets the time function option. Default is [time.Now].
func UseTimeFunction(f TimeFunction) Option {
	return func(o *Options) {
		o.TimeFunction = f
	}
}

// UseTimeFormat sets the time format option. Default is [log.DefaultTimeFormat].
func UseTimeFormat(f string) Option {
	return func(o *Options) {
		o.TimeFormat = f
	}
}

// UseLevel sets the level option. Default is [log.InfoLevel].
func UseLevel(l Level) Option {
	return func(o *Options) {
		o.Level = l
	}
}

// UsePrefix sets the prefix option. Default is no prefix.
func UsePrefix(p string) Option {
	return func(o *Options) {
		o.Prefix = p
	}
}

// UseReportTimestamp sets the report timestamp option. Default is false.
func UseReportTimestamp(r bool) Option {
	return func(o *Options) {
		o.ReportTimestamp = r
	}
}

// UseReportCaller sets the report caller option. Default is false.
func UseReportCaller(r bool) Option {
	return func(o *Options) {
		o.ReportCaller = r
	}
}

// UseCallerFormatter sets the caller formatter option. Default is [log.ShortCallerFormatter].
func UseCallerFormatter(f CallerFormatter) Option {
	return func(o *Options) {
		o.CallerFormatter = f
	}
}

// UseFields sets the fields option. Default is no fields.
func UseFields(fields map[string]slog.Value) Option {
	return func(o *Options) {
		for k, v := range fields {
			o.Fields = append(o.Fields, k, v)
		}
	}
}

// UseFormatter sets the formatter option. Default is [TextFormatter].
func UseFormatter(f Formatter) Option {
	return func(o *Options) {
		o.Formatter = f
	}
}

// UseCallerOffset sets the caller offset option. Default is 0.
func UseCallerOffset(offset int) Option {
	return func(o *Options) {
		o.CallerOffset = offset
	}
}

// UseOutput sets the writer option. Default is [os.Stderr].
func UseOutput(w io.Writer) Option {
	return func(o *Options) {
		o.Writer = w
	}
}

// UseStyles sets the styles option. Default is [DefaultStyles].
func UseStyles(s *Styles) Option {
	return func(o *Options) {
		o.Styles = s
	}
}

// AsDefault sets the logger as the default logger. Default is false.
func AsDefault() Option {
	return func(o *Options) {
		o.Default = true
	}
}
