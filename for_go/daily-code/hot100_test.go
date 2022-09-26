package dailycode

import "testing"

func TestTwoSum(t *testing.T) {
	ans := TwoSum([]int{3, 3}, 6)
	if len(ans) < 2 {
		t.Error("the length is error")
	}
	if ans[0] != 0 || ans[1] != 1 {
		t.Errorf("the value is error %d,%d", ans[0], ans[1])
	}
}
