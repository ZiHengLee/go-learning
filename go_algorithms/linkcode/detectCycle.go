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
		if fast == slow {
			fast := head
			//减去共同的那一段，剩下的两段是相等的
			slow = slow.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
		slow = slow.Next
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
