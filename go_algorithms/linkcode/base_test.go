package linkcode

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	node := ListNode{}
	a := []int{1, 2, 4}
	first := node.Create(a)
	b := []int{1, 2, 3}
	second := node.Create(b)
	//mergeTwoLists(first, second)
	res := mergeTwoLists(first, second)
	print(res.Val)
	for i := res; i != nil; i = i.Next {
		if i.Next != nil {
			print(i.Val)
		}
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	node := ListNode{}
	a := []int{1}
	first := node.Create(a)
	res := removeNthFromEnd(first, 1)
	print(res.Val)
	for i := first; i != nil; i = i.Next {
		if i.Next != nil {
			print(i.Val)
		}
	}
}

func TestCreatLinkListHead(t *testing.T) {
	a := []int{1,2,3,4}
	node := createListNodeTail(a)
	reorderList(node)
	printLists(node)
}

func TestReverseKGroup(t *testing.T) {
	a := []int{1,2,3,4,5}
	node := createListNodeTail(a)
	//printLists(node)
	newNode := reverseKGroup(node.Next,2)
	printLists(newNode)
}

func TestReverseBetween(t *testing.T) {
	a :=[]int{1,2,3,4,5}
	node := createListNodeTail(a)
	//printLists(node)
	newNode := reverseBetween(node.Next,2,4)
	printLists(newNode)
}

func TestDetectCycle(t *testing.T) {
	a := &ListNode{3,nil}
	b := &ListNode{2,nil}
	c := &ListNode{0,nil}
	d := &ListNode{-4,nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b
	res := detectCycle(a)
	fmt.Println(res.Val)
}

func TestMergeKLists(t *testing.T) {
	lists := make([]*ListNode,0)

	a := &ListNode{1,nil}
	b := &ListNode{4,nil}
	c := &ListNode{5,nil}
	a.Next = b
	b.Next = c
	d := &ListNode{1,nil}
	e := &ListNode{3,nil}
	f := &ListNode{4,nil}
	d.Next = e
	e.Next = f
	g := &ListNode{2,nil}
	h := &ListNode{6,nil}
	g.Next =h

	lists = append(lists, a,d,g)
	res := mergeKLists(lists)
	printLists(res)
}

func TestQuickSortLists(t *testing.T) {
	a := &ListNode{1,nil}
	cur := a
	for i:=2;i<10;i++{
		t := &ListNode{i,nil}
		cur.Next = t
		cur=cur.Next
	}
	//b := &ListNode{4,nil}
	//c := &ListNode{5,nil}
	//d := &ListNode{6,nil}
	//e := &ListNode{2,nil}
	//f := &ListNode{9,nil}
	//g := &ListNode{8,nil}
	//a.Next = b
	//b.Next = c
	//c.Next = d
	//d.Next = e
	//e.Next = f
	//f.Next = g


	res := sortListv1(a)
	printLists(res)
}