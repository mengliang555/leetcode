package array_string

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

func removeDuplicates2(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	left, right := 2, 2
	for right < len(nums) {
		if nums[right] != nums[left-2] {
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
