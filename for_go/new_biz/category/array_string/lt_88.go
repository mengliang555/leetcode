package array_string

//https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
//88. 合并两个有序数组
//给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i, v := range nums2 {
			nums1[i] = v
		}
		return
	}
	if n == 0 {
		return
	}

	ans := nums1
	num1P, num2P, lastIndex := m-1, n-1, m+n-1

	for num1P >= 0 && num2P >= 0 {
		if nums1[num1P] > nums2[num2P] {
			ans[lastIndex] = nums1[num1P]
			num1P--
		} else {
			ans[lastIndex] = nums2[num2P]
			num2P--
		}
		lastIndex--
	}

	for num1P >= 0 {
		ans[lastIndex] = nums1[num1P]
		num1P--
		lastIndex--
	}

	for num2P >= 0 {
		ans[lastIndex] = nums2[num2P]
		num2P--
		lastIndex--
	}

}
