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
					middle--
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
					middle++
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildListNode(nums []int) *ListNode {
	root := &ListNode{}
	temp := root
	for _, v := range nums {
		temp.Next = &ListNode{Val: v}
		temp = temp.Next
	}
	return root.Next
}

func PrintListNode(root *ListNode) {
	nums := make([]int, 0)
	for root != nil {
		nums = append(nums, root.Val)
		root = root.Next
	}
	fmt.Printf("%+v\n", nums)
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	replaceNode := func(root, targetNode *ListNode) *ListNode {
		temp := root
		if targetNode.Val > temp.Val {
			return targetNode
		}
		for temp.Next != nil {
			if temp.Next.Val < targetNode.Val {
				temp.Next = targetNode
				return root
			}
			temp = temp.Next
		}
		temp.Next = targetNode
		return root
	}
	root := head
	lastNum := head.Val
	for head.Next != nil {
		head = head.Next
		if head.Val > lastNum {
			root = replaceNode(root, head)
		}
		lastNum = head.Val
	}
	return root
}

func maxI(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minI(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 辗转相除法
func gcd(a, b int) int {
	maxV, minV := maxI(a, b), minI(a, b)
	for val := maxV % minV; val != 0; {
		maxV = minV
		minV = val
		val = maxV % minV
	}
	return minV
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	root := head
	for head.Next != nil {
		if head.Val == 10 {
			fmt.Println(head.Val)
		}
		temp := head.Next
		gcdV := gcd(head.Val, temp.Val)
		head.Next = &ListNode{
			Val:  gcdV,
			Next: temp,
		}
		head = temp
	}
	return root
}

func canSeePersonsCount(heights []int) []int {
	if len(heights) <= 1 {
		return []int{0}
	}
	ans := make([]int, len(heights))
	maxH := make([]int, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		for len(maxH) > 0 && maxH[len(maxH)-1] < heights[i] {
			maxH = maxH[0 : len(maxH)-1]
			ans[i] += 1
		}
		if len(maxH) > 0 {
			ans[i]++
		}
		maxH = append(maxH, heights[i])
	}
	return ans
}
