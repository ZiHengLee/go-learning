package linkcode

//思路：这个题简化很多情况，刚好倒着输出就行了，基本思路尾插法，进位保留到下一位
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := l1, l2
	pre := 0
	head := new(ListNode)
	cur := head
	prenode := cur
	for s1 != nil || s2 != nil {
		val1, val2 := 0, 0
		if s1 != nil {
			val1 = s1.Val
			s1 = s1.Next
		}
		if s2 != nil {
			val2 = s2.Val
			s2 = s2.Next
		}
		now := (val1 + val2 + pre) % 10
		pre = (val1 + val2 + pre) / 10
		cur.Val = now
		cur.Next = new(ListNode)
		prenode = cur
		cur = cur.Next
	}
	//很关键，这就是为什么要保留一个prenode的原因
	if pre > 0 {
		cur.Val = pre
	} else {
		prenode.Next = nil
	}
	return head
}
