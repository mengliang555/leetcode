package dailycode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZero(t *testing.T) {
	ans := []int{0, 1, 0, 3, 12}
	moveZeroes(ans)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, ans)
}
