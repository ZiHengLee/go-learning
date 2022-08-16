package linkcode

//识别环形链表并记录初始位置
//快慢节点
//3 2 0 -4
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head.Next
	for fast != nil {
		//fmt.Println(slow.Val,fast.Val)
		if fast == slow {
			fast := head
			//slow = slow.next其实有点不是特别能理解
			slow = slow.Next
			//fmt.Println(fast.Val,slow.Val)
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
		slow = slow.Next
		//fmt.Println(slow.Val)
		if slow == nil {
			return nil
		}
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == nil {
			return nil
		}
	}
	return nil
}
