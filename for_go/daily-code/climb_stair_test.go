package dailycode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStair2(t *testing.T) {
	assert.Equal(t, 3, climbStairs2(3))
	assert.Equal(t, 5, climbStairs2(4))
}

func TestClimbStair(t *testing.T) {
	assert.Equal(t, 3, climbStairs(3))
	assert.Equal(t, 5, climbStairs(4))
	assert.Equal(t, 8, climbStairs(5))
	assert.Equal(t, 13, climbStairs(6))
}
