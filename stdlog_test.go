package log_test

import (
	"bytes"
	"testing"

	"github.com/bartventer/log"
	"github.com/stretchr/testify/assert"
)

func TestStandardLog(t *testing.T) {
	var buf bytes.Buffer
	l := log.New(log.UseOutput(&buf))
	stdLogger := log.StandardLog(func(slo *log.StandardLogOptions) {
		slo.ForceLevel = log.InfoLevel
		slo.Logger = l
	})
	stdLogger.Println("test message")
	assert.Contains(t, buf.String(), "INFO")
	assert.Contains(t, buf.String(), "test message")

	// Test with default logger
	buf.Reset()
	log.SetOutput(&buf)
	stdLogger = log.StandardLog()
	stdLogger.Println("test message")
	assert.Contains(t, buf.String(), "INFO")
	assert.Contains(t, buf.String(), "test message")
}
