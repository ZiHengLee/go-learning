package commonmistake_test

import (
	"fmt"
	"runtime"
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

func f(x int) func(int) int {
	g := func(y int) int {
		return x + y
	}
	// 返回闭包
	return g
}

// 测试闭包
func TestClosure(t *testing.T) {
	// 将函数的返回结果"闭包"赋值给变量a
	a := f(3) //3是外部函数的值
	// 调用存储在变量中的闭包函数
	res := a(5) //5是内部函数的值
	fmt.Println(res)
	// 可以直接调用闭包
	// 因为闭包没有赋值给变量，所以它称为匿名闭包
	fmt.Println(f(3)(5))
}

// 测试协程调度
// https://learnku.com/articles/85761
// 总是输出最后一个
func TestGoroutine(t *testing.T) {
	runtime.GOMAXPROCS(1)
	for n := 0; n <= 10; n++ {
		ch := make(chan int, 10)
		for i := 1; i <= 10; i++ {
			go func(i int) {
				ch <- i
			}(i)
		}
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		fmt.Println("done")
	}
}

//https://learnku.com/articles/56077
//http不关闭的内存泄漏
