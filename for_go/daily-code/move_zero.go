package dailycode

func moveZeroes(nums []int) {
	zeroMaxLength, zeroIndex := 0, 0
	for i, v := range nums {
		if v == 0 {
			zeroMaxLength++
			zeroIndex = i
			continue
		}
		if zeroMaxLength != 0 {
			nums[zeroIndex-zeroMaxLength+1], nums[i] = v, 0
			zeroIndex = i
		}
	}
}
