package linkcode

//删除链表倒数第k个节点
//前后节点，前节点提前移动
//主要是考虑几种极端情况
//n>len(list)
//n=len(list)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre,after := head,head
	if head == nil{
		return nil
	}
	if head.Next == nil && n == 1{
		return nil
	}
	start := 0
	for after.Next!= nil{
		after = after.Next
		start++
		if start == n{
			break
		}
	}
	if start == n-1 && after.Next==nil{
		return head.Next
	}

	//fmt.Println(pre.Val,after.Val)
	for after.Next != nil{
		pre = pre.Next
		after = after.Next
	}
	if pre != nil && pre.Next != nil{
		pre.Next = pre.Next.Next
	}
	return head
}
