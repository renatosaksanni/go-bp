package di

import (
	"testing"
)

func TestBuildContainer(t *testing.T) {
	container := BuildContainer()
	if container.APIHandler == nil {
		t.Errorf("APIHandler is nil")
	}
	if container.AuthMiddleware == nil {
		t.Errorf("AuthMiddleware is nil")
	}
}
