package dailycode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	if val := longestPalindrome("babad"); val != "bab" && val != "aba" {
		t.Error("err", val)
	}
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}
