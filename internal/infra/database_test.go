package infra

import (
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	err := os.Setenv("DATABASE_DSN", "user:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}
	db := InitDB()
	if db == nil {
		t.Errorf("Database connection is nil")
	}
}
