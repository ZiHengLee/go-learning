package linkcode

func splitListToParts(head *ListNode, k int) []*ListNode {
	cur := head
	allLen := 0
	for cur != nil {
		cur = cur.Next
		allLen += 1
	}
	m, n := allLen/k, allLen%k
	res := make([]*ListNode, k)
	if head == nil {
		return res
	}
	cur = head
	for i := 0; i < k; i++ {
		tempI := &ListNode{-1, nil}
		rlen := m
		if i < n {
			rlen += 1
		}
		curTemp := tempI
		for j := 0; j < rlen; j++ {
			if cur == nil {
				break
			}
			curTemp.Next = cur
			curTemp = curTemp.Next
			cur = cur.Next
		}
		curTemp.Next = nil
		res[i] = tempI.Next
		if cur == nil {
			break
		}
	}
	return res
}
