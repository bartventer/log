package log

import "log/slog"

type ValueFunc func() slog.Value

func (f ValueFunc) LogValue() slog.Value {
	return f()
}
