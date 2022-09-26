package linkcode

//https://leetcode.cn/problems/reverse-nodes-in-k-group/
//每k个数反转一次，最初思路就是头插法

func reverseKGroup(head *ListNode, k int) *ListNode {
	var lenList int
	temp := head
	for temp != nil {
		temp = temp.Next
		lenList++
	}
	retHeader := &ListNode{-1, nil}
	curOrigin := head
	curNow := retHeader
	for i := 0; i < lenList/k; i++ {
		var newTemp *ListNode
		lastNode := curNow
		for j := 0; j < k; j++ {
			if j == 0 {
				lastNode = curOrigin
			}

			newTemp = curOrigin.Next
			curOrigin.Next = curNow.Next
			curNow.Next = curOrigin
			curOrigin = newTemp
		}
		curNow = lastNode
	}
	for curOrigin != nil {
		curNow.Next = curOrigin
		curOrigin = curOrigin.Next
		curNow = curNow.Next
	}
	return retHeader.Next
}

func reverseList(head *ListNode) *ListNode {
	retList := &ListNode{-1, nil}
	curOrigin := head
	for curOrigin != nil {
		temp := curOrigin.Next
		curOrigin.Next = retList.Next
		retList.Next = curOrigin
		curOrigin = temp
	}
	return retList.Next
}
