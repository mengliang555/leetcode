package target

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	ans := twoSum([]int{2, 7, 11, 15}, 9)
	actual := []int{0, 1}
	assert.Equal(t, actual[0], ans[0])
	assert.Equal(t, actual[1], ans[1])
}

func generateListNode(nums []int) *ListNode {
	header := new(ListNode)
	current := header

	for i, v := range nums {
		current.Val = v
		if i == len(nums)-1 {
			break
		}
		current.Next = new(ListNode)
		current = current.Next
	}
	return header
}
func TestAddTwoNumbers(t *testing.T) {
	ans := addTwoNumbers(generateListNode([]int{2, 4, 9}), generateListNode([]int{5, 6, 4, 9}))
	actual := []int{7, 0, 4, 0, 1}
	for _, v := range actual {
		assert.Equal(t, v, ans.Val)
		ans = ans.Next
	}
}

func TestLongestPalindrome(t *testing.T) {
	fmt.Printf("longestPalindrome(): %v\n", longestPalindrome("aaaa"))
}

func Test_ThreeSum(t *testing.T) {
	val := threeSum([]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0})
	fmt.Printf("val: %v\n", val)
}

func Test_LetterCombinations(t *testing.T) {
	fmt.Printf("letterCombinations(\"23\"): %v\n", letterCombinations("23"))
}

func Test_RemoveNthFromEnd(t *testing.T) {
	val := removeNthFromEnd(generateListNode([]int{1, 2, 3, 4, 5}), 2)
	fmt.Printf("val: %v\n", val)
}

func Test_Function(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 6}
	val := a[0]

	for i := 1; i < len(a); i++ {
		val &= a[i]
	}
	fmt.Printf("val: %v\n", val)
}
