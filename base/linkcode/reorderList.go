package linkcode

import "fmt"

//交替进行插入
//L0 → L1 → … → Ln - 1 → Ln
//请将其重新排列后变为：
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//[1,2,3,4]
//[1,4,2,3]

//思路，如果直接用数组，其实就是头部插一个尾部再插入一个

//但是如果必须用链表的话
//拆分，反转，合并
func reorderList(head *ListNode)  {
	if head == nil || head.Next==nil{return}
	midNode := getMidNode(head)
	before := head
	after := midNode.Next
	midNode.Next = nil
	after = reverseList(after)
	head = simpleMergeList(before,after)
}

func getMidNode(head *ListNode) *ListNode{
	slow,fast := head,head
	fmt.Println(slow,fast.Next)
	for slow.Next != nil && fast.Next != nil &&fast.Next.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
//一前一后的顺序合并
func simpleMergeList(before,after *ListNode) *ListNode{
	curBefor,curAfter := before,after
	retList := &ListNode{-1,nil}
	cur := retList
	for curBefor != nil || curAfter != nil{
		if curBefor != nil{
			cur.Next = curBefor
			curBefor = curBefor.Next
			cur =cur.Next
		}
		if curAfter != nil{
			cur.Next = curAfter
			curAfter = curAfter.Next
			cur =cur.Next
		}
	}
	return retList.Next
}