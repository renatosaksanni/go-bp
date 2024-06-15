package di

import (
	"github.com/renatosaksanni/go-bp/config"
	"github.com/renatosaksanni/go-bp/internal/api"
	"github.com/renatosaksanni/go-bp/internal/infra/repository"
	"github.com/renatosaksanni/go-bp/pkg/middleware"

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
