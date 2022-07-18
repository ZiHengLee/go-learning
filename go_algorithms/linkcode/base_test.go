package linkcode

import "testing"

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
	a := []int{1,2,3,4,5,6}
	node := createListNodeTail(a)
	printLists(node)
}

func TestReverseKGroup(t *testing.T) {
	a := []int{1,2,3,4,5}
	node := createListNodeTail(a)
	//printLists(node)
	newNode := reverseKGroup(node.Next,2)
	printLists(newNode)
}