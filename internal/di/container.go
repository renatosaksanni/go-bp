package di

import (
	"go-bp/config"
	"go-bp/internal/api"
	"go-bp/internal/infra/repository"
	"go-bp/pkg/middleware"

	"github.com/google/wire"
)

type Container struct {
	APIHandler     *api.Handler
	AuthMiddleware *middleware.AuthMiddleware
}

func BuildContainer() Container {
	wire.Build(
		config.GetDB,
		config.GetVaultClient,
		repository.NewRepository,
		api.NewHandler,
		middleware.NewAuthMiddleware,
		wire.Struct(new(Container), "APIHandler", "AuthMiddleware"),
	)
	return Container{}
}
