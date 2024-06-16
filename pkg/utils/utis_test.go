package utils

import (
	"errors"
	"testing"
)

func TestLogError(t *testing.T) {
	err := errors.New("Test error")
	LogError(err)
	// No assertion needed, just ensure no panic
}

func TestValidateData(t *testing.T) {
	valid := ValidateData(nil)
	if !valid {
		t.Errorf("Expected valid data")
	}
}
