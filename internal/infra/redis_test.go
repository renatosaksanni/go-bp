package infra

import (
	"testing"
)

func TestInitRedis(t *testing.T) {
	rdb := InitRedis()
	if rdb == nil {
		t.Errorf("Failed to initialize Redis client")
	}
}
