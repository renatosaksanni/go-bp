package di

import (
	"go-bp/config"
	"go-bp/internal/api"
	"go-bp/internal/infra/repository"
	"go-bp/pkg/middleware"

	gwire "github.com/google/wire"
)

type Container struct {
	APIHandler     *api.Handler
	AuthMiddleware *middleware.AuthMiddleware
}

func BuildContainer() Container {
	gwire.Build(
		config.GetDB,
		config.GetVaultClient,
		repository.NewRepository,
		api.NewHandler,
		middleware.NewAuthMiddleware,
		gwire.Struct(new(Container), "APIHandler", "AuthMiddleware"),
	)
	return Container{}
}
