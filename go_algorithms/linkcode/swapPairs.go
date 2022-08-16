package linkcode

//非递归
func swapPairs(head *ListNode) *ListNode {
	ret := &ListNode{Val: -1, Next: nil}
	ret.Next = head
	cur := ret
	for cur.Next != nil && cur.Next.Next != nil {
		before := cur.Next
		after := cur.Next.Next
		before.Next = after.Next
		after.Next = before
		cur.Next = after
		cur = before
	}
	return ret.Next
}

//递归
func swapPairsv1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsv1(next.Next)
	next.Next = head
	return next
}
