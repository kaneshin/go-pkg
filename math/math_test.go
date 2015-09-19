package math

import (
	"testing"

	"github.com/kaneshin/go-pkg/testing/assert"
)

func TestMax(t *testing.T) {
	assert.Equal(t, -3.0, Min(-3))
	assert.Equal(t, 123.0, Min(123, 1234, 10000, 123455))
	assert.Equal(t, -23.0, Min(123, 1234, -23, 123455))
}

// TestRound ...
func TestRound(t *testing.T) {
	assert.Equal(t, 0.0, Round(0.0))

	assert.Equal(t, 0.0, Round(0.4))
	assert.Equal(t, 1.0, Round(0.5))

	assert.Equal(t, 5.0, Round(5.4))
	assert.Equal(t, 6.0, Round(5.5))

	assert.Equal(t, -5.0, Round(-5.4))
	assert.Equal(t, -6.0, Round(-5.5))
}
