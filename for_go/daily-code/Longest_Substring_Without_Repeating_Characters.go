package dailycode

func lengthOfLongestSubstring(s string) int {
	var valMap [256]bool
	left, right, maxLength := 0, 0, 0
	for left < len(s) && right < len(s) {
		if !valMap[s[right]-'a'] {
			valMap[s[right]-'a'] = true
			right++
		} else if valMap[s[left]-'a'] {
			valMap[s[left]-'a'] = false
			left++
		}
		maxLength = max(maxLength, right-left)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
