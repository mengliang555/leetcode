package dailycode

func maxArea(height []int) int {
	left, right, maxA := 0, len(height)-1, 0
	for left < right {
		maxA = max(maxA, min(height[left], height[right])*(right-left))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxA
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
