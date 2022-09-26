package linkcode

import (
	"fmt"
	"sync"
	"testing"
	"time"
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
	a := []int{1, 2, 3, 4}
	node := createListNodeTail(a)
	reorderList(node)
	printLists(node)
}

func TestReverseKGroup(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	node := createListNodeTail(a)
	//printLists(node)
	newNode := reverseKGroup(node.Next, 2)
	printLists(newNode)
}

func TestReverseBetween(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	node := createListNodeTail(a)
	//printLists(node)
	newNode := reverseBetween(node.Next, 2, 4)
	printLists(newNode)
}

func TestDetectCycle(t *testing.T) {
	a := &ListNode{3, nil}
	b := &ListNode{2, nil}
	c := &ListNode{0, nil}
	d := &ListNode{-4, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b
	res := detectCycle(a)
	fmt.Println(res.Val)
}

func TestMergeKLists(t *testing.T) {
	lists := make([]*ListNode, 0)

	a := &ListNode{1, nil}
	b := &ListNode{4, nil}
	c := &ListNode{5, nil}
	a.Next = b
	b.Next = c
	d := &ListNode{1, nil}
	e := &ListNode{3, nil}
	f := &ListNode{4, nil}
	d.Next = e
	e.Next = f
	g := &ListNode{2, nil}
	h := &ListNode{6, nil}
	g.Next = h

	lists = append(lists, a, d, g)
	res := mergeKLists(lists)
	printLists(res)
}

func TestQuickSortLists(t *testing.T) {
	a := &ListNode{1, nil}
	b := &ListNode{4, nil}
	c := &ListNode{3, nil}
	d := &ListNode{2, nil}
	e := &ListNode{5, nil}
	f := &ListNode{2, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	//f.Next = g
	//printLists(a)
	printLists(a)
	fmt.Println()
	res := partitionv1(a, 3)
	printLists(res)
}

func TestLlLL(t *testing.T) {
	var a map[string]string
	a = make(map[string]string)
	a["a"] = "b"
	fmt.Println(a)
}

func TestSpiralMatrix(t *testing.T) {
	a := []int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0}
	node := createListNodeTail(a)
	ret := spiralMatrix(3, 5, node.Next)
	fmt.Println(ret)
}

func TestTT(t *testing.T) {
	wg := sync.WaitGroup{}
	c := make(chan int, 0)
	wg.Add(2)
	var data int
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		data = <-c
		print(data)
	}()
	go func() {
		defer wg.Done()
		c <- 1
	}()
	wg.Wait()
	print("hell")
}
