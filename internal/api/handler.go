package api

import (
	"net/http"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Handler struct {
	tracer trace.Tracer
	mu     sync.Mutex
}

func NewHandler() *Handler {
	return &Handler{
		tracer: otel.Tracer("go-bp"),
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, span := h.tracer.Start(r.Context(), "HealthCheckHandler")
	defer span.End()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
