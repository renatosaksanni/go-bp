# go-bp
Boilerplate for any future Go language development purposes

Project structure:

go-bp/
├── cmd/
│   └── main.go
├── config/
│   ├── config.go
│   └── vault.go
├── internal/
│   ├── api/
│   │   └── handler.go
│   ├── app/
│   │   └── app.go
│   ├── di/
│   │   └── container.go
│   ├── infra/
│   │   ├── repository/
│   │   │   └── repository.go
│   │   ├── database.go
│   │   ├── redis.go
│   │   └── telemetry.go
├── pkg/
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── logging.go
│   │   ├── recovery.go
│   │   ├── cors.go
│   │   └── ratelimit.go
│   └── utils/
│       └── utils.go
├── go.mod
└── go.sum
