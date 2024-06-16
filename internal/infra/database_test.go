package infra

import (
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	os.Setenv("DATABASE_DSN", "host=localhost user=postgres password=postgres dbname=testdb port=5432 sslmode=disable")
	db := InitDB()
	if db == nil {
		t.Errorf("Database connection is nil")
	}
}
