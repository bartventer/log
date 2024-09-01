package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithContext(t *testing.T) {
	logger := New()
	ctx := WithContext(context.Background(), logger)
	assert.Equal(t, logger, ctx.Value(ContextKey))
}

func TestFromContext(t *testing.T) {
	defaultLogger := FromContext(context.Background())
	assert.NotNil(t, defaultLogger)

	logger := New()
	ctx := WithContext(context.Background(), logger)
	loggerFromCtx := FromContext(ctx)
	assert.Equal(t, logger, loggerFromCtx)

}
