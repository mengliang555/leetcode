package hot100

import (
	"container/heap"
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
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		k := len(nums) - 1

		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for j < k && nums[i]+nums[j]+nums[k] > 0 {
				k--
			}

			if j == k {
				break
			}

			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
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

func minExtraChar(s string, dictionary []string) int {
	d := make([]int, 0, len(s)) // d[i] 表示s前缀[0...i-1] 的额外字符长度
	valM := make(map[string]struct{})
	for _, v := range dictionary {
		valM[v] = struct{}{}
	}
	for i := 1; i <= len(s); i++ {
		d[i] = d[i-1] + 1
		for j := i - 1; j >= 0; j-- {
			if _, ok := valM[s[j:i]]; ok {
				d[i] = min(d[i], d[j])
			}
		}
	}
	return d[len(s)]
}

func subarraySum(nums []int, k int) int {
	totalC := 0
	for i := range nums {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				totalC++
			}
		}
	}
	return totalC
}

func subarraySumWithPrefix(nums []int, k int) int {
	totalC := 0
	prefixSum := map[int]int{0: 1} // 当前缀和[0:i]等于k的时候, 加1
	sum := 0
	for i := range nums {
		sum += nums[i]
		if val, ok := prefixSum[sum-k]; ok {
			totalC += val
		}
		prefixSum[sum] = prefixSum[sum] + 1
	}
	return totalC
}

type hp struct {
	sort.IntSlice
}

var a []int

func (h *hp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	val := h.IntSlice
	v := val[len(val)-1]
	h.IntSlice = val[:len(val)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums

	hpp := &hp{make([]int, k)}

	for i := 0; i < k; i++ {
		hpp.IntSlice[i] = i
	}

	heap.Init(hpp)

	println(hpp.IntSlice[0])

	ans := make([]int, 1, len(nums)-k+1)

	ans[0] = nums[hpp.IntSlice[0]]

	for i := k; i < len(nums); i++ {
		heap.Push(hpp, i)

		for hpp.IntSlice[0] <= i-k {
			heap.Pop(hpp)
		}

		ans = append(ans, nums[hpp.IntSlice[0]])
	}
	return ans
}
