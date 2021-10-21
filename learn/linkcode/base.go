package linkcode

type ListNodeIface interface {
	Create(array []int64) *ListNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) Create(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	node := new(ListNode)
	node.Val = array[0]
	preNode := node
	for i := 1; i < len(array); i++ {
		nodeNext := new(ListNode)
		nodeNext.Val = array[i]
		preNode.Next = nodeNext
		preNode = nodeNext
	}
	return node
}

func LenListNode(l *ListNode) int {
	len := 0
	for i := l; i != nil; i = i.Next {
		len += 1
	}
	return len
}

//合并有序链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := new(ListNode)
	tempStart := newNode
	temp1, temp2 := l1, l2
	for ; temp1 != nil || temp2 != nil; {
		tempNode := new(ListNode)
		if temp1 != nil && temp2 != nil {
			if temp1.Val < temp2.Val {
				tempNode.Val = temp1.Val
				temp1 = temp1.Next
			} else {
				tempNode.Val = temp2.Val
				temp2 = temp2.Next
			}
		} else if temp1 != nil {
			tempNode.Val = temp1.Val
			temp1 = temp1.Next
		} else if temp2 != nil{
			tempNode.Val = temp2.Val
			temp2 = temp2.Next
		}
		tempStart.Next = tempNode
		tempStart = tempNode
	}
	return newNode.Next
}

//删除倒数第k个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left := head
	right := head
	i := 0
	for ; i < n; i++ {
		if right != nil {
			right = right.Next
		}
	}
	if right == nil {
		if i == n{
			return head.Next
		}else{
			return head
		}
	}
	for ; right.Next != nil; {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return head
}
