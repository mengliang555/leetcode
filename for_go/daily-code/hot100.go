package dailycode

func TwoSum(nums []int, target int) []int {
	ans := make(map[int]int)
	for i, v := range nums {
		if val, ok := ans[target-v]; ok {
			return []int{val, i}
		}
		ans[v] = i
	}
	return []int{-1, -1}
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}

	return -1
}
