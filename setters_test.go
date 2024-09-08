package log_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/bartventer/log"
	"github.com/stretchr/testify/assert"
)

func TestSetCallerFormatter(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetCallerFormatter(log.ShortCallerFormatter, logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetCallerOffset(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetCallerOffset(2, logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetFormatter(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetFormatter(log.JSONFormatter, logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetLevel(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetLevel(log.DebugLevel, logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetOutput(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New()
	log.SetOutput(&buf, logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetPrefix(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetPrefix("TEST", logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "TEST")
	assert.Contains(t, buf.String(), "test message")
}

func TestWithPrefix(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	logger = log.WithPrefix(logger, "TEST")

	logger.Info("test message")
	assert.Contains(t, buf.String(), "TEST")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetStyles(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetStyles(log.DefaultStyles(), logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetReportCaller(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetReportCaller(true, logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetReportTimestamp(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetReportTimestamp(true, logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetTimeFormat(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetTimeFormat("2006-01-02", logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}

func TestSetTimeFunction(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(log.UseOutput(&buf))
	log.SetTimeFunction(func(t time.Time) time.Time { return t.Add(time.Hour) }, logger)

	logger.Info("test message")
	assert.Contains(t, buf.String(), "test message")
}
