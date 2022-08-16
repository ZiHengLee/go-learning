package linkcode

//奇偶链表
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	oddHead,evenHead := head,head.Next
	curOdd,curEven := oddHead,evenHead
	for curEven != nil{
		curOdd.Next = curEven.Next
		curOdd = curOdd.Next

		curEven.Next = curOdd.Next
		curEven = curEven.Next

	}
	curOdd.Next = evenHead
	return oddHead
}