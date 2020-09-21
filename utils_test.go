package berrylan

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// Test_boolToInt tests the boolToInt method.
func Test_boolToInt(t *testing.T) {
	a := true
	b := false

	assert.Equal(t, 1, boolToInt(a))
	assert.Equal(t, 0, boolToInt(b))
}
