package linkcode

//思路：这个题简化很多情况，刚好倒着输出就行了，基本思路尾插法，进位保留到下一位
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := l1
	s2 := l2
	pre := 0
	res := new(ListNode)
	help := res
	prenode := res
	for i := 0; s1 != nil || s2 != nil; i++ {
		val1, val2 := 0, 0
		if s1 != nil {
			val1 = s1.Val
		}
		if s2 != nil {
			val2 = s2.Val
		}
		now := (val1 + val2 + pre) % 10
		pre = (val1 + val2 + pre) / 10
		help.Val = now
		help.Next = new(ListNode)
		prenode = help
		help = help.Next
		if s1 != nil {
			s1 = s1.Next
		}
		if s2 != nil {
			s2 = s2.Next
		}
	}
	if pre > 0 {
		help.Val = pre
	} else {
		prenode.Next = nil
	}
	return res
}
