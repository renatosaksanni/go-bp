package main

import (
	"go-bp/internal/di"
	"go-bp/internal/infra"
	"go-bp/pkg/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	tp := infra.InitTracer()
	defer infra.ShutdownTracer(tp)

	container := di.BuildContainer()

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.RecoveryMiddleware)
	r.Use(middleware.CORSMiddleware)

	rateLimiter := middleware.NewRateLimiter(1, 5) // 1 request per second with burst size of 5
	r.Use(rateLimiter.RateLimitMiddleware)

	r.HandleFunc("/health", container.APIHandler.HealthCheck).Methods("GET")
	r.Use(container.AuthMiddleware.ServeHTTP)

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
