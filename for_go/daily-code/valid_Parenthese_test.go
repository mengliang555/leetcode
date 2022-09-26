package dailycode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, true, isValid("()[]{}"))
	assert.Equal(t, false, isValid("(]"))
}
