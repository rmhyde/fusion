package helpers

import (
	"context"
	"testing"

	"github.com/rs/zerolog"
)

func NewTestWriterContext(t *testing.T) context.Context {
	logger := zerolog.New(zerolog.TestWriter{T: t}).With().Timestamp().Logger()
	return logger.WithContext(context.Background())
}
