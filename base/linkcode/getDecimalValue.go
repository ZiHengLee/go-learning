package linkcode

func getDecimalValue(head *ListNode) int {
	if head == nil {
		return 0
	}
	k := make([]int, 0)
	cur := head
	for cur != nil {
		k = append(k, cur.Val)
		cur = cur.Next
	}
	res := 0
	l := len(k)
	for index, data := range k {
		//fmt.Println(l-index-1, 1<<(l-index-1), data)
		res += (1 << (l - index - 1)) * data
	}
	return res
}
