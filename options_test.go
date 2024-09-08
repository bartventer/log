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
	UseTimeFunction(timeFunc)(options)
	UseTimeFormat(timeFormat)(options)
	UseLevel(level)(options)
	UsePrefix(prefix)(options)
	UseReportTimestamp(reportTimestamp)(options)
	UseReportCaller(reportCaller)(options)
	UseCallerFormatter(callerFormatter)(options)
	UseFields(fields)(options)
	UseFormatter(formatter)(options)
	UseCallerOffset(callerOffset)(options)
	UseOutput(&buf)(options)
	UseStyles(styles)(options)
	AsDefault()(options)

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
