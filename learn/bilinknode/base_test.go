package bilinknode

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(2, 1)
	//fmt.Println(lRUCache.head.val,lRUCache.head.next.val,lRUCache.tail.pre.val,lRUCache.tail.val)
	lRUCache.Put(2, 2)
	fmt.Println(lRUCache.Get(2))
	//fmt.Println(lRUCache.tail.pre.val)
	lRUCache.Put(1, 1)
	//ListPrint(lRUCache.head)
	//fmt.Println(lRUCache.Get(2))
	lRUCache.Put(4, 1)
	//ListPrint(lRUCache.head)
	fmt.Println(lRUCache.Get(1))
	fmt.Println(lRUCache.Get(3))
	fmt.Println(lRUCache.Get(4))
}
func ListPrint(node *LinkNode){
	for i := node;i!= nil;{
		fmt.Print(i.val)
		fmt.Print("-")
		i=i.next
	}
	fmt.Println()
}