package log_test

import (
	"bytes"
	"log/slog"
	"testing"
	"time"

	"github.com/bartventer/log"
	"github.com/stretchr/testify/assert"
)

func TestDefaultStyles(t *testing.T) {
	styles := log.DefaultStyles()
	assert.NotNil(t, styles)
	assert.Equal(t, 5, styles.Levels[log.DebugLevel].GetMaxWidth())
	assert.Equal(t, 4, styles.Levels[log.InfoLevel].GetMaxWidth())
	assert.Equal(t, 4, styles.Levels[log.WarnLevel].GetMaxWidth())
	assert.Equal(t, 5, styles.Levels[log.ErrorLevel].GetMaxWidth())
	assert.Equal(t, 5, styles.Levels[log.FatalLevel].GetMaxWidth())
}

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(
		log.UseOutput(&buf),
		log.UseLevel(log.DebugLevel),
		log.UsePrefix("TEST"),
		log.UseReportTimestamp(true),
		log.UseReportCaller(true),
		log.UseCallerFormatter(log.ShortCallerFormatter),
		log.UseFields(map[string]slog.Value{"key": slog.StringValue("value")}),
		log.UseFormatter(log.JSONFormatter),
		log.UseCallerOffset(2),
		log.UseStyles(log.DefaultStyles()),
	)
	assert.NotNil(t, logger)

	logger.Debug("test message")
	assert.Contains(t, buf.String(), log.DebugLevel.String())
	assert.Contains(t, buf.String(), "TEST")
	assert.Contains(t, buf.String(), "key")
	assert.Contains(t, buf.String(), "value")
	assert.Contains(t, buf.String(), "test message")

}

func TestLogFunctions(t *testing.T) {
	t.Run("SetCallerFormatter", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetCallerFormatter(log.ShortCallerFormatter)
		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetCallerOffset", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetCallerOffset(2)
		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetFormatter", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetFormatter(log.JSONFormatter)

		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetLevel", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetLevel(log.DebugLevel)

		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetOutput", func(t *testing.T) {
		var buf bytes.Buffer
		logger := log.New(log.AsDefault())
		log.SetOutput(&buf, logger)

		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetPrefix", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetPrefix("TEST")

		log.Info("test message")
		assert.Contains(t, buf.String(), "TEST")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("WithPrefix", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetPrefix("TEST")

		log.Info("test message")
		assert.Contains(t, buf.String(), "TEST")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetStyles", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetStyles(log.DefaultStyles())

		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetReportCaller", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetReportCaller(true)

		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetReportTimestamp", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetReportTimestamp(true)

		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetTimeFormat", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetTimeFormat("2006-01-02")

		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("SetTimeFunction", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetTimeFunction(func(t time.Time) time.Time { return t.Add(time.Hour) })

		log.Info("test message")
		assert.Contains(t, buf.String(), "test message")
	})

	t.Run("Debug", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetLevel(log.DebugLevel)
		log.Debug("debug message")
		assert.Contains(t, buf.String(), "debug message")
	})

	t.Run("Debugf", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.SetLevel(log.DebugLevel)
		log.Debugf("debug %s", "message")
		assert.Contains(t, buf.String(), "debug message")
	})

	t.Run("Info", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.Info("info message")
		assert.Contains(t, buf.String(), "info message")
	})

	t.Run("Infof", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.Infof("info %s", "message")
		assert.Contains(t, buf.String(), "info message")
	})

	t.Run("Warn", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.Warn("warn message")
		assert.Contains(t, buf.String(), "warn message")
	})

	t.Run("Warnf", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.Warnf("warn %s", "message")
		assert.Contains(t, buf.String(), "warn message")
	})

	t.Run("Error", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.Error("error message")
		assert.Contains(t, buf.String(), "error message")
	})

	t.Run("Errorf", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.Errorf("error %s", "message")
		assert.Contains(t, buf.String(), "error message")
	})

	t.Run("Print", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.Print("print message")
		assert.Contains(t, buf.String(), "print message")
	})

	t.Run("Log", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.Log(log.InfoLevel, "log message", slog.Bool("key", true))
		assert.Contains(t, buf.String(), "log message")
		assert.Contains(t, buf.String(), "key")
		assert.Contains(t, buf.String(), "true")
	})

	t.Run("Logf", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		log.Logf(log.InfoLevel, "log %s", "message")
		assert.Contains(t, buf.String(), "log message")
	})
}
