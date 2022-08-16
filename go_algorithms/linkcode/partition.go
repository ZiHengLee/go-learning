package linkcode

func partitionv1(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left := &ListNode{Val: -1, Next: nil}
	right := &ListNode{Val: -1, Next: nil}
	cur := head
	leftCur := left
	rightCur := right
	for cur != nil {
		if cur.Val < x {
			leftCur.Next = cur
			leftCur = leftCur.Next
			//fmt.Println(cur.Val)
		}else{
			//fmt.Println(cur.Val)
			rightCur.Next = cur
			rightCur = rightCur.Next
		}
		cur = cur.Next
	}
	leftCur.Next = right.Next
	//关键在于这个地方一定要断开，不然会指向left，然后不断循环
	rightCur.Next = nil
	return left.Next
}
