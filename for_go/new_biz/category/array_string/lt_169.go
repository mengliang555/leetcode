package array_string

//https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150

//给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。

func majorityElement(nums []int) int {
	var num int
	var count int

	for _, v := range nums {
		if num == v {
			count++
		} else {
			count--
			if count < 0 {
				num = v
				count = 0
			}
		}
	}

	return num
}
