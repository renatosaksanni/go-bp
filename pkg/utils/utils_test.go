package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogError(t *testing.T) {
	t.Run("No error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			LogError(nil)
		})
	})

	t.Run("With error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			LogError(assert.AnError)
		})
	})
}

func TestValidateData(t *testing.T) {
	t.Run("Valid data", func(t *testing.T) {
		assert.True(t, ValidateData(nil))
		assert.True(t, ValidateData("test"))
	})
}
