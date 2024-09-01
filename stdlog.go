package log

import (
	stdlog "log"
	"log/slog"

	"github.com/charmbracelet/log"
)

type (
	// StandardLogOptions can be used to configure the standard log adapter.
	StandardLogOptions struct {
		log.StandardLogOptions
		Logger *slog.Logger // Log is the logger to use. Default is the default logger.
	}

	// StandardLogOption is a standard logger option.
	StandardLogOption func(*StandardLogOptions)
)

// StandardLog creates a new standard logger with the given options.
func StandardLog(opts ...StandardLogOption) *stdlog.Logger {
	o := &StandardLogOptions{}
	for _, opt := range opts {
		opt(o)
	}

	if o.Logger != nil {
		return o.Logger.Handler().(*log.Logger).StandardLog(o.StandardLogOptions)
	}

	if defaultOnce.l.Load() == nil {
		_ = Default()
	}

	return log.StandardLog(o.StandardLogOptions)
}
