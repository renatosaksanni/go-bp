package config

import (
	"github.com/renatosaksanni/go-bp/internal/infra"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return infra.InitDB()
}

var ProviderSet = wire.NewSet(GetDB, GetVaultClient)
