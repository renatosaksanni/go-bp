package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func init() {
	// Initialize a no-op tracer provider to avoid errors during tests
	tp := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(tp)
}

func TestHealthCheck(t *testing.T) {
	handler := NewHandler()

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.HealthCheck(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHealthCheck_WriteError(t *testing.T) {
	handler := NewHandler()

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a custom ResponseRecorder to inject the failWriter
	rr := httptest.NewRecorder()
	fw := &failWriter{ResponseRecorder: rr}
	handler.HealthCheck(fw, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expected := "Failed to write response\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

// failWriter is a custom http.ResponseWriter that fails on write
type failWriter struct {
	*httptest.ResponseRecorder
}

func (fw *failWriter) Write(p []byte) (int, error) {
	return 0, http.ErrAbortHandler
}
