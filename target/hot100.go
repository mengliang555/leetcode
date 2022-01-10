package target

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	valMap := make(map[int]int)

	for i, v := range nums {
		if val, ok := valMap[target-v]; ok {
			return []int{val, i}
		}
		valMap[v] = i
	}
	return []int{-1, -1}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) print() {
	if list == nil {
		fmt.Println("empty listNode")
		return
	}
	fmt.Println("begin print listNode")
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
	fmt.Println("close print listNode")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	flag := 0
	header := new(ListNode)
	current := header
	for l1 != nil && l2 != nil {
		current.Val = l1.Val + l2.Val + flag
		flag = current.Val / 10
		current.Val %= 10
		if l1.Next != nil || l2.Next != nil {
			current.Next = new(ListNode)
		} else {
			l1 = l1.Next
			l2 = l2.Next
			break
		}
		l1 = l1.Next
		l2 = l2.Next
		current = current.Next
	}

	l1.print()
	l2.print()
	for l1 != nil {
		current.Val = l1.Val + flag
		flag = current.Val / 10
		current.Val %= 10
		if l1.Next != nil {
			current.Next = new(ListNode)
		} else {
			break
		}
		current = current.Next
		l1 = l1.Next
	}

	for l2 != nil {
		current.Val = l2.Val + flag
		flag = current.Val / 10
		current.Val %= 10
		if l2.Next != nil {
			current.Next = new(ListNode)
		} else {
			break
		}
		current = current.Next
		l2 = l2.Next
	}

	if flag > 0 {
		current.Next = &ListNode{
			Val: flag,
		}
	}
	return header
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// fdjlbcba
func lengthOfLongestSubstring(s string) int {
	val := new([128]int)

	for i := range val {
		val[i] = -1
	}
	start, res := 0, 0
	for i, v := range s {
		index := v
		start = MaxInt(start, val[index]+1)
		res = MaxInt(res, i-start+1)
		val[index] = i
	}
	return res
}

// dp[i][j] means i->j all is palid or not
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	maxLength := 0
	left, right := 0, 0
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] && ((j-i <= 2) || dp[i+1][j-1]) {
				dp[i][j] = true
				if maxLength < j-i+1 {
					maxLength = j - i + 1
					left, right = i, j
				}
			}
		}
	}
	return string(s[left : right+1])
}

func maxArea(height []int) int {
	ans := 0
	left, right := 0, len(height)-1
	for left < right {
		val := MinInt(height[left], height[right]) * (right - left)
		ans = MaxInt(ans, val)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func threeSum(nums []int) [][]int {
	if len(nums) <= 2 {
		return nil
	}
	sort.Ints(nums)
	fmt.Printf("nums: %v\n", nums)

	ans := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		} else if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left, right := i+1, len(nums)-1; left < right; {
			val := nums[left] + nums[right] + nums[i]
			if val < 0 {
				for left += 1; left < right && nums[left] == nums[left-1]; left += 1 {
				}
			} else if val > 0 {
				for right -= 1; left < right && nums[right] == nums[right+1]; right -= 1 {
				}
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left += 1; left < right && nums[left] == nums[left-1]; left++ {
				}
				for right -= 1; left < right && nums[right] == nums[right+1]; right-- {
				}
			}
		}
	}
	return ans
}

func letterCombinations(digits string) []string {
	valMap := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	ans := make([]string, 0)
	dfsLetterCombinations(digits, "", 0, valMap, &ans)
	return ans
}

func dfsLetterCombinations(originStr, currentStr string, index int, valMap map[byte]string, ans *[]string) {
	if index >= len(originStr) {
		*ans = append(*ans, currentStr)
		return
	}
	if val, ok := valMap[originStr[index]]; ok {
		for j := 0; j < len(val); j++ {
			dfsLetterCombinations(originStr, currentStr+string(val[j]), index+1, valMap, ans)
		}
	}
}

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func dayOfTheWeek(day int, month int, year int) string {
	monthSet := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	days := (year-1971)*365 + (year-1969)/4
	for _, d := range monthSet[:month] {
		days += d
	}
	if month >= 2 && (year%400 == 0 || (year%4 == 1 && year%100 != 0)) {
		day += 1
	}
	return week[(days+3)%7]
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// reverse the List
	// remove the N Node
	length := 0
	if head == nil {
		return head
	} else if head.Next == nil {
		return nil
	}
	current := head
	for current != nil {
		current = current.Next
		length++
	}
	if length == n {
		return head.Next
	}

	current = head
	for i := 0; i < length-n-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	return head
}

func isValid(s string) bool {
	val := make([]rune, 0, 0)
	for _, v := range s {

		switch v {
		case '(', '[', '{':
			val = append(val, v)
		case ')':
			if len(val) == 0 || val[len(val)-1] != '(' {
				return false
			} else {
				val = val[:len(val)-1]
			}
		case ']':
			if len(val) == 0 || val[len(val)-1] != '[' {
				return false
			} else {
				val = val[:len(val)-1]
			}
		case '}':
			if len(val) == 0 || val[len(val)-1] != '{' {
				return false
			} else {
				val = val[:len(val)-1]
			}
		}
	}
	return len(val) == 0
}
