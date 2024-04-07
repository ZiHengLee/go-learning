package commonmistake_test

import (
	"fmt"
	"testing"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

// for语句的闭包中使用迭代变量
func TestRangeFor(t *testing.T) {
	data := []field{{"one"}, {"two"}, {"three"}}
	for _, v := range data {
		// 解决办法：添加如下语句
		// v := v
		go v.print()
	}
	//output: three, three, three
	time.Sleep(time.Second)

	data2 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data2 {
		go v.print() // go执行是函数，函数执行之前，函数的接受对象已经传过来
	}
	//output: one, two, three(随机顺序)
	time.Sleep(time.Second)
}
