package array_string

// https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150

// 189. 轮转数组
// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数
func rotate(nums []int, k int) {
	k = k % len(nums)
	if len(nums) <= 1 {
		return
	}

	reverseNum(nums)
	reverseNum(nums[:k])
	reverseNum(nums[k:])
	return
}

func reverseNum(nums []int) {
	length := len(nums)

	if length <= 1 {
		return
	}

	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
}
