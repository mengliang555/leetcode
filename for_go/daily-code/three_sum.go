package dailycode

import "sort"

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp > 0 {
				right--
			} else if tmp < 0 {
				left++
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				right--
			}

		}
	}
	return ans
}
