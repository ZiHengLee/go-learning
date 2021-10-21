package main

import (
	"fmt"
	"reflect"
	"errors"
)

type Person struct {
	Name     string
	Sex      string
}
func (p *Person) Add(a, b string) {
	fmt.Printf("me %+v\n", a+b)
}
func (p *Person) AddRes(a, b string) (string, error){
	if b == "boy" {
		return a + b, errors.New("its a boy")
	}
	return a+b, nil
}
func main() {
	funcName := "Add"
	typeT := &Person{}
	a := reflect.ValueOf("li")
	b := reflect.ValueOf("femal")
	in := []reflect.Value{a, b}
	reflect.ValueOf(typeT).MethodByName(funcName).Call(in)//如果是没有参数的函数就需要传nil
	//me lifemal

	c := reflect.ValueOf("J")
	d := reflect.ValueOf("boy")
	inr := []reflect.Value{c, d}
	funcName = "AddRes"
	ret := reflect.ValueOf(typeT).MethodByName("AddRes").Call(inr)
	fmt.Printf("ret is %+v\n", ret)
	//ret is [Jboy <error Value>]
	for i := 0; i < len(ret); i++ {
		fmt.Printf("ret index:%+v, type:%+v, value:%+v\n", i, ret[i].Kind(), ret[i].Interface())
	}
	//ret index:0, type:string, value:Jboy
	//ret index:1, type:interface, value:its a boy
}