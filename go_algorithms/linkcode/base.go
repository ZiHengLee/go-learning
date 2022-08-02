package linkcode

import "fmt"

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
	for temp1 != nil || temp2 != nil {
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
		} else if temp2 != nil {
			tempNode.Val = temp2.Val
			temp2 = temp2.Next
		}
		tempStart.Next = tempNode
		tempStart = tempNode
	}
	return newNode.Next
}

//删除倒数第k个节点
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	left := head
//	right := head
//	i := 0
//	for ; i < n; i++ {
//		if right != nil {
//			right = right.Next
//		}
//	}
//	if right == nil {
//		if i == n {
//			return head.Next
//		} else {
//			return head
//		}
//	}
//	for right.Next != nil {
//		left = left.Next
//		right = right.Next
//	}
//	left.Next = left.Next.Next
//	return head
//}

//遍历链表
func printLists(head *ListNode){
	temp := head
	for temp != nil{
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}

//头插法建立链表
func createListNodeHead(a []int) *ListNode{
	retList := &ListNode{-1,nil}
	cursor := retList
	for _, v := range a{
		newNode := &ListNode{v,nil}
		newNode.Next = cursor.Next
		cursor.Next = newNode
	}
	return retList
}



//尾插法建立链表
func createListNodeTail(a []int) *ListNode{
	retList := &ListNode{-1,nil}
	cursor := retList
	for _, v := range a{
		newNode := &ListNode{v,nil}
		newNode.Next = cursor.Next
		cursor.Next = newNode
		cursor = cursor.Next
	}
	return retList
}

//随意给两个位置旋转
//头插法尾插法结合，关键在于找起始点和结尾点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	retList := &ListNode{-1,nil}
	cur := head
	curHead := retList
	curTail := retList
	start := 0
	for cur != nil {
		newNode := &ListNode{cur.Val,nil}
		if start>=left-1 && start<right{
			newNode.Next = curHead.Next
			curHead.Next = newNode
			if start == left-1{
				curTail = newNode
			}
			//fmt.Println(curHead.Val,curTail.Val,"f")
		}else{
			newNode.Next = curTail.Next
			curTail.Next = newNode

			curHead = newNode
			curTail = newNode
			//fmt.Println(curHead.Val,curTail.Val,"s")
		}

		cur = cur.Next
		start++
	}
	return retList.Next
}