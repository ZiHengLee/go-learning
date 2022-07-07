package main

import (
	"fmt"
	"reflect"
)

type temp interface {
	f()
}
type Father struct {
}
// 实现 f() 方法
func (father Father) f() {
	fmt.Println("father transfer f()")
}
//可以理解为继承
type Son struct {
	temp
	Name string
}

//不实现就直接调用用Father的方法
//func (son Son) f() {
// fmt.Println("son transfer f()")
//}

func main() {
	//如果没有把Father作为匿名参数，编译期不会报错，运行时会报错
	var d temp = Son{}
	fmt.Println(reflect.TypeOf(d))
	d.f()
	//fmt.Println(reflect.TypeOf(d))
}