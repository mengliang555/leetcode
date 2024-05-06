package array_string

import (
	"fmt"
	"testing"
)

func Print(v any) {
	fmt.Printf("%+v\n", v)
}

func Test_merge(t *testing.T) {
	v := []int{1, 2, 3, 0, 0, 0}
	merge(v, 3, []int{2, 5, 6}, 3)
	Print(v)
}

func Test_removeElement(t *testing.T) {
	v := []int{3, 2, 2, 3}
	println(removeElement(v, 3))
	Print(v)

	v1 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	println(removeElement(v1, 2))
	Print(v1)
}

func Test_removeDuplicates(t *testing.T) {
	v := []int{1, 1, 2}
	println(removeDuplicates(v))
	Print(v)
	v1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	println(removeDuplicates(v1))
	Print(v1)
}

func Test_majorityElement(t *testing.T) {
	println(majorityElement([]int{
		2, 2, 1, 1, 1, 2, 2}))

}

func Test_rotate(t *testing.T) {
	v := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(v, 3)
	Print(v)
}

func Test_maxProfit(t *testing.T) {
	println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func Test_maxProfit2(t *testing.T) {
	println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}
