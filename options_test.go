package log

import (
	"bytes"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestApplyOptions(t *testing.T) {
	timeFunc := func(t time.Time) time.Time { return t.Add(time.Hour) }
	timeFormat := time.RFC3339
	level := DebugLevel
	prefix := "PREFIX"
	reportTimestamp := true
	reportCaller := true
	callerFormatter := ShortCallerFormatter
	fields := map[string]slog.Value{
		"key": slog.StringValue("value"),
	}
	formatter := JSONFormatter
	callerOffset := 2
	var buf bytes.Buffer
	styles := DefaultStyles()

	// Apply the options
	options := &Options{
		LogOptions: &LogOptions{},
	}
	WithTimeFunction(timeFunc)(options)
	WithTimeFormat(timeFormat)(options)
	WithLevel(level)(options)
	WithPrefix(prefix)(options)
	WithReportTimestamp(reportTimestamp)(options)
	WithReportCaller(reportCaller)(options)
	WithCallerFormatter(callerFormatter)(options)
	WithFields(fields)(options)
	WithFormatter(formatter)(options)
	WithCallerOffset(callerOffset)(options)
	WithWriter(&buf)(options)
	WithStyles(styles)(options)
	WithDefault()(options)

	// Verify the options
	assert.NotNil(t, options.TimeFunction)
	assert.Equal(t, timeFormat, options.TimeFormat)
	assert.Equal(t, level, options.Level)
	assert.Equal(t, prefix, options.Prefix)
	assert.True(t, options.ReportTimestamp)
	assert.True(t, options.ReportCaller)
	assert.NotNil(t, options.CallerFormatter)
	assert.Contains(t, options.Fields, "key")
	assert.Equal(t, formatter, options.Formatter)
	assert.Equal(t, callerOffset, options.CallerOffset)
	assert.Equal(t, &buf, options.Writer)
	assert.Equal(t, styles, options.Styles)
	assert.True(t, options.Default)
}
