package dailycode

import (
	"strconv"
	"strings"
)

func maxLengthBetweenEqualCharacters(s string) int {
	ans := 0
	maxF := func(a, b int) int {
		if a <= b {
			return b
		}
		return a
	}
	charMap := make(map[rune]int)
	for i, v := range s {
		if singC, ok := charMap[v]; ok {
			ans = maxF(ans, i-singC-1)
		} else {
			charMap[v] = i
		}
	}
	return ans
}

//func lengthOfLongestSubstring(s string) int {
//	charCount := make([]int, 26)
//	charLoc := make([]int, 26)
//	maxF := func(a, b int) int {
//		if a <= b {
//			return b
//		}
//		return a
//	}
//
//	ans := 1
//	start, end := 0, 0
//	for ; end < len(s); end++ {
//		charCount[end]++
//		if charCount[end] > 0 {
//			for i := start; i < end; i++ {
//				charCount[s[i]-'a']--
//			}
//			start = charLoc[s[end-'a']] + 1
//		} else {
//			ans = maxF(ans, end-start)
//		}
//
//	}
//
//	return ans
//}

func dayOfYear(date string) int {
	dataMap := strings.Split(date, "-")

	calculateRain := func(val string) bool {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return false
		}

		if valInt%100 == 0 {
			return valInt%400 == 0
		} else {
			return valInt%4 == 0
		}
	}

	isRain := calculateRain(dataMap[0])

	calCul := func(isRain bool, month int) int {
		switch month {
		case 0:
			return 0
		case 1, 3, 5, 7, 8, 10, 12:
			return 31
		case 2:
			if isRain {
				return 29
			}
			return 28
		default:
			return 30
		}
	}

	valMonth, err := strconv.Atoi(dataMap[1])

	if err != nil {
		return 0
	}

	totalDay, err := strconv.Atoi(dataMap[2])

	if err != nil {
		return 0
	}

	for i := 0; i < valMonth; i++ {
		totalDay += calCul(isRain, i)
	}
	return totalDay
}

//计算平闰

func diagonalSum(mat [][]int) int {
	ans := 0
	for i, eachLine := range mat {
		ans += eachLine[i]
		if len(eachLine)-1-i != i {
			ans += eachLine[len(eachLine)-1-i]
		}
	}
	return ans
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i] = s[len(s)-i], s[i]
	}
}

/**
* Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val = root2.Val + root1.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
