package log_test

import (
	"bytes"
	"log/slog"
	"testing"

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
		log.WithWriter(&buf),
		log.WithLevel(log.DebugLevel),
		log.WithPrefix("TEST"),
		log.WithReportTimestamp(true),
		log.WithReportCaller(true),
		log.WithCallerFormatter(log.ShortCallerFormatter),
		log.WithFields(map[string]slog.Value{"key": slog.StringValue("value")}),
		log.WithFormatter(log.JSONFormatter),
		log.WithCallerOffset(2),
		log.WithStyles(log.DefaultStyles()),
	)
	assert.NotNil(t, logger)

	logger.Debug("test message")
	assert.Contains(t, buf.String(), log.DebugLevel.String())
	assert.Contains(t, buf.String(), "TEST")
	assert.Contains(t, buf.String(), "key")
	assert.Contains(t, buf.String(), "value")
	assert.Contains(t, buf.String(), "test message")

}
