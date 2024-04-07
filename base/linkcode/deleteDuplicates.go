package linkcode

//head = [1,2,3,3,4,4,5]
//[1,2,5]
//删除排列重复的元素
//最初思路，记录连续三个节点,发现在连接段不太好处理
//正确思路，用递归的思想去做，简化到其实每次处理两个连续的节点就行

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		tmp := head.Next
		for tmp != nil && head.Val == tmp.Val {
			tmp = tmp.Next
		}
		return deleteDuplicates(tmp)
	} else {
		head.Next = deleteDuplicates(head.Next)
	}
	return head
}
