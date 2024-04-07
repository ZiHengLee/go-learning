package linkcode

//https://leetcode.cn/problems/next-greater-node-in-linked-list/

//链表中更大的节点
func nextLargerNodes(head *ListNode) []int {
	if head == nil{
		return nil
	}
	headNow := ListNode{Val: -1,Next: nil}
	cur := head
	llist := 0
	for cur!= nil{
		temp := cur.Next
		cur.Next = headNow.Next
		headNow.Next = cur
		cur = temp
		llist+=1
	}
	//fmt.Println(llist)
	curNew := headNow.Next
	res := make([]int,llist)
	myStack := make([]int,0)
	myStack = append(myStack, curNew.Val)
	res[llist-1] = 0
	curNew = curNew.Next
	start := 1
	for curNew != nil{
		start+=1
		top := myStack[len(myStack)-1]
		for curNew.Val >= top && len(myStack) > 0{
			myStack = myStack[:len(myStack)-1]
			if len(myStack) > 0{
				top = myStack[len(myStack)-1]
			}
		}
		if len(myStack) > 0{
			res[llist-start] = top
		}else{
			res[llist-start] = 0
		}
		myStack = append(myStack, curNew.Val)
		curNew = curNew.Next
	}
	return res
}