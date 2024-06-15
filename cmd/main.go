package main

import (
	"log"
	"net/http"
	"os"

	"github.com/renatosaksanni/go-bp/internal/di"
	"github.com/renatosaksanni/go-bp/internal/infra"
	"github.com/renatosaksanni/go-bp/pkg/middleware"

	"github.com/gorilla/mux"
)

func run() error {
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
	return http.ListenAndServe(":8080", r)
}

func main() {
	if err := run(); err != nil {
		log.Printf("Server failed: %s\n", err)
		os.Exit(1)
	}
}
