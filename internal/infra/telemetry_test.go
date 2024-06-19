package infra

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInitTracer(t *testing.T) {
	// Set up the environment variable for the OTLP endpoint
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_ENDPOINT")

	tp := InitTracer()

	assert.NotNil(t, tp, "TracerProvider should not be nil")

	// Test if a tracer can be created from the provider
	tracer := tp.Tracer("test-tracer")
	assert.NotNil(t, tracer, "Tracer should not be nil")

	// Shutdown the tracer provider to ensure it shuts down correctly
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := tp.Shutdown(ctx)
	assert.NoError(t, err, "Failed to shutdown TracerProvider")
}

func TestShutdownTracer(t *testing.T) {
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317")
	defer os.Unsetenv("OTEL_EXPORTER_OTLP_ENDPOINT")

	tp := InitTracer()
	assert.NotNil(t, tp, "TracerProvider should not be nil")

	// Call the ShutdownTracer function and verify there are no errors
	ShutdownTracer(tp)
}
