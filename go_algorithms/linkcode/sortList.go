package linkcode

import (
	"math/rand"
	"time"
)

//对链表进行排序
//是归并排序，关键在于每次找到中间的分割点
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	start := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		start += 1
	}
	return mergeSort(head, start)
}

func mergeSort(first *ListNode, len int) *ListNode {
	if len == 1 {
		return first
	} else {
		mid := &ListNode{Val: -1}
		cur := first
		for i := 0; i < len; i++ {
			if i == len/2 {
				break
			}
			mid = cur
			cur = cur.Next
		}
		newLeft := mergeSort(first, len/2)
		right := mid.Next
		mid.Next = nil
		var newRight *ListNode
		if len%2 == 0 {
			newRight = mergeSort(right, len/2)
		} else {
			newRight = mergeSort(right, len/2+1)
		}
		return mergeTwoLists(newLeft, newRight)
	}
}

//方法2，使用快排，小的链接到前一段，大的链接后一段
func sortListv1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := &ListNode{Val: -1}
	currr := ret
	//ret.Next = head
	tempArr := make([]int, 0)
	cur := head
	for cur != nil {
		tempArr = append(tempArr, cur.Val)
		cur = cur.Next
	}
	shuffing(tempArr)
	for _, a := range tempArr {
		b := &ListNode{Val: a}
		currr.Next = b
		currr = currr.Next
	}

	quickSort(ret, head, nil)
	return ret.Next
}

func quickSort(ret, head, tail *ListNode) {
	if head != tail {
		swap(head, tail)
		mid := partition(ret, head)
		quickSort(ret, ret.Next, mid)
		quickSort(mid, mid.Next, tail)
	}
}

func partition(ret, head *ListNode) *ListNode {
	firstVal := head.Val
	mid := &ListNode{Val: firstVal}
	leftH, rightH := &ListNode{Val: -1}, &ListNode{Val: -1}
	left, right, cur := leftH, rightH, head.Next
	for cur != nil {
		if cur.Val < firstVal {
			left.Next = cur
			left = left.Next
		} else {
			right.Next = cur
			right = right.Next
		}
		cur = cur.Next
	}
	right.Next = nil
	//链接中间节点
	mid.Next = rightH.Next
	left.Next = mid
	ret.Next = leftH.Next
	return mid
}

func swap(start, end *ListNode) {
	slow, fast := start, start.Next
	for fast != end && fast.Next != end {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Val
	slow.Val = start.Val
	start.Val = tmp
}

//func swap(start,end *ListNode){
//	cur := start
//	len := 0
//	for cur.Next!= end{
//		cur = cur.Next
//		// if rand.Intn(100) > 90{
//		// 	break
//		// }
//		len+=1
//	}
//	cur = start
//	for i:=0;i<len/2;i++{
//		cur = cur.Next
//	}
//
//	tmp := start.Val
//	start.Val = cur.Val
//	cur.Val = tmp
//}

//func swap(start,end *ListNode){
//	slow,fast := start,start.Next
//	for fast != end && fast.Next != end && fast.Next.Next != end{
//		slow = slow.Next
//		fast = fast.Next.Next.Next
//	}
//
//	tmp := slow.Val
//	slow.Val = start.Val
//	start.Val = tmp
//}

func shuffing(arr []int) {
	seed := time.Now().UTC().UnixNano()
	r := rand.New(rand.NewSource(seed))

	N := len(arr)
	for i := 0; i < N; i++ {
		t := r.Intn(N - i)
		arr[i], arr[i+t] = arr[i+t], arr[i]
	}
}
