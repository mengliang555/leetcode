package dailycode

type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers 用例已经通过
func _(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := new(ListNode)
	root := ans
	i1, i2, val := 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			i1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			i2 = l2.Val
			l2 = l2.Next
		}
		root.Next = &ListNode{Val: (i1 + i2 + val) % 10}
		i1, i2, val = 0, 0, (i1+i2+val)/10
		root = root.Next
	}
	if val != 0 {
		root.Next = &ListNode{Val: val}
	}
	return ans.Next
}
