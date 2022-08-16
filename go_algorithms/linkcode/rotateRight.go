package linkcode

//旋转链表，向后移动几位

//思路

//最开始想直接用倒数第k个节点那种方式，找到开始节点，再把首尾节点连接，就可以了，但是这里的k是可以大于链表长度的，那么可以直接先连接
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	cur := head
	llen := 1
	for cur.Next != nil{
		cur = cur.Next
		llen += 1
	}
	//fmt.Println(llen)
	tail := cur
	cur.Next = head
	tK := k%llen
	after := head
	for i:=0;i<tK;i++{
		after = after.Next
	}
	cur = tail
	for after != head{
		after = after.Next
		cur = cur.Next
	}
	ret := cur.Next
	cur.Next = nil
	return ret
}
