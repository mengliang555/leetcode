package dailycode

func longestPalindrome(s string) string {
	startAns, TotalLen := 0, 0
	for i := range s {
		v1 := startWithOne(i, i, s)
		v2 := startWithOne(i, i+1, s)
		if val := max(v1, v2); val > TotalLen {
			startAns, TotalLen = i, val
		}
	}
	if TotalLen%2 == 0 {
		return s[startAns-(TotalLen-1)/2 : startAns+TotalLen/2+1]
	} else {
		return s[startAns-TotalLen/2 : startAns+TotalLen/2+1]
	}

}

func startWithOne(center1, center2 int, s string) int {
	left, right := center1, center2
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			return right - left - 1
		}
		left--
		right++
	}
	return right - left - 1
}
