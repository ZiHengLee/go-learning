package 迭代器模式

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	var datas []*Row
	datas = append(datas, &Row{data: "hello"})
	datas = append(datas, &Row{data: "wolrd"})
	iterator := CollectionIterator{Rows: datas, index: 0}

	for ; iterator.hasNext(); iterator.getNext() {
		fmt.Println(iterator.CurrentItem())
	}
}
