[![GoDoc](https://godoc.org/github.com/renatosaksanni/go-bp?status.svg)](https://godoc.org/github.com/renatosaksanni/go-bp)
[![codecov](https://codecov.io/gh/renatosaksanni/go-bp/branch/main/graph/badge.svg)](https://codecov.io/gh/renatosaksanni/go-bp)
![Build Status](https://img.shields.io/github/actions/workflow/status/renatosaksanni/go-bp/ci.yml?branch=main)
![Go Report Card](https://goreportcard.com/badge/github.com/renatosaksanni/go-bp)
![License](https://img.shields.io/github/license/renatosaksanni/go-bp.svg)
![Issues](https://img.shields.io/github/issues/renatosaksanni/go-bp.svg)


# go-bp
Boilerplate for any future Go language development purposes

### Folders and Files

### cmd/
- `main.go`: The main file that runs the application.

### config/
- `config.go`: General configuration file.
- `vault.go`: Configuration for integration with HashiCorp Vault.

### internal/
#### api/
- `handler.go`: Handler for the API.

#### app/
- `app.go`: Main application logic.

#### di/
- `container.go`: Dependency injection configuration.

#### infra/
##### repository/
- `repository.go`: Repository implementation.

- `database.go`: Database configuration.
- `redis.go`: Redis configuration.
- `telemetry.go`: Telemetry configuration.

### pkg/
#### middleware/
- `auth.go`: Middleware for authentication.
- `logging.go`: Middleware for logging.
- `recovery.go`: Middleware for panic recovery.
- `cors.go`: Middleware for CORS.
- `ratelimit.go`: Middleware for rate limiting.

#### utils/
- `utils.go`: General utility functions.

### Root
- `go.mod`: Go module file that defines project dependencies.
- `go.sum`: Checksum file to ensure the integrity of downloaded dependencies.
