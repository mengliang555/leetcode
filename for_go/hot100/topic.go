package hot100

import (
	"fmt"
	"sort"
	"strings"
)

func moveZeroes(nums []int) {
	zeroIndex := 0
	for i, v := range nums {
		if v != 0 {
			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			zeroIndex++
		}
	}
}

func GetSet[key comparable](val []key) map[key]struct{} {
	valMap := make(map[key]struct{})
	for _, v := range val {
		valMap[v] = struct{}{}
	}
	return valMap
}

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	ans := 0
	for _, v := range nums {
		numMap[v] = true
	}
	for _, v := range nums {
		if numMap[v-1] {
			continue
		}

		matchVal := v
		length := 0
		for numMap[matchVal] {
			matchVal++
			length++
		}

		if length > ans {
			ans = length
		}
	}

	return ans
}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, v := range nums {
		if num, ok := numsMap[target-v]; ok && target-v != v {
			return []int{num, i}
		}
		numsMap[v] = i
	}
	return []int{-1, -1}
}

func groupAnagrams(strs []string) [][]string {
	valMap := make(map[string][]string)

	intJoin := func(frequency []int) string {
		ans := strings.Builder{}
		for i, v := range frequency {
			if v == 0 {
				continue
			}
			ans.WriteString(fmt.Sprintf("%b%d", i+'a', v))
		}
		return ans.String()
	}

	calculateStr := func(val string) string {
		charFrequency := make([]int, 26)
		for _, v := range val {
			charFrequency[v-'a']++
		}
		return intJoin(charFrequency)
	}

	for _, v := range strs {
		mapK := calculateStr(v)
		if _, ok := valMap[mapK]; !ok {
			valMap[mapK] = make([]string, 0)
		}
		valMap[mapK] = append(valMap[mapK], v)
	}

	ans := make([][]string, 0, len(valMap))

	for _, v := range valMap {
		ans = append(ans, v)
	}

	return ans

}

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	ans := 0

	maxF := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for start < end {
		width := end - start
		var maxHeight int
		if height[start] > height[end] {
			maxHeight = height[end]
			end--
		} else {
			maxHeight = height[start]
			start++
		}
		ans = maxF(ans, width*maxHeight)

	}
	return ans
}

func threeSum(nums []int) [][]int {
	left, right := 0, len(nums)-1
	ans := make([][]int, 0)
	sort.Ints(nums)
	for left < right-1 {
		current := nums[left] + nums[right]
		middle := (left + right) / 2
		if current+nums[middle] >= 0 {
			for middle > left && middle < right {
				if current+nums[middle] == 0 {
					ans = append(ans, []int{nums[left], nums[middle], nums[right]})
					break
				}
				if middle-1 > left && nums[middle-1] == nums[middle] {
					continue
				}
				middle--
			}
		} else {
			for middle > left && middle < right {
				if current+nums[middle] == 0 {
					ans = append(ans, []int{nums[left], nums[middle], nums[right]})
					break
				}
				if middle+1 < right && nums[middle+1] == nums[middle] {
					continue
				}
				middle++
			}
		}

		if current > 0 {
			for right-1 > left && nums[right-1] == nums[right] {
				right -= 1
			}
			right -= 1
		} else {
			for left+1 < right && nums[left+1] == nums[left] {
				left += 1
			}
			left += 1
		}

	}
	return ans
}
