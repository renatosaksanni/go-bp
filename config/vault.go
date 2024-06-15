package config

import (
	"os"

	"github.com/hashicorp/vault/api"
)

func GetVaultClient() (*api.Client, error) {
	config := api.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	token := os.Getenv("VAULT_TOKEN")
	client.SetToken(token)
	return client, nil
}
