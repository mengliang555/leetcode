package hot100

import (
	"fmt"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}

func Test_groupAnagrams(t *testing.T) {
	fmt.Printf("%+v", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func Test_maxArea(t *testing.T) {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func Test_threeSum(t *testing.T) {
	fmt.Printf("%+v", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("%+v", threeSum([]int{0, 1, 1}))
	fmt.Printf("%+v", threeSum([]int{0, 0, 0}))

}
