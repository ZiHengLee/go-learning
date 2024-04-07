package 访问者模式

import (
	"reflect"
	"testing"
)

func TestVisitorSimple(t *testing.T) {
	//已知类型
	apple := Apple{price: 1}
	banana := Banana{price: 1}
	v := NormalVisitor{}
	v.BuyBanana(banana)
	v.BuyApple(apple)
}

func TestVisitorFruits(t *testing.T) {
	//未知类型
	f := make([]Fruit, 2)
	f[0] = Apple{price: 1}
	f[1] = Banana{price: 1}
	v := NormalVisitor{}
	for _, fruit := range f {
		typ := reflect.TypeOf(fruit)
		switch typ.Name() {
		case "Apple":
			v.BuyApple(fruit.(Apple))
		case "Banana":
			v.BuyBanana(fruit.(Banana))
		}
	}
	//这种方式在增加类型的时候在使用处也需要修改
}

func TestVisitorFruitsNew(t *testing.T) {
	//未知类型，交给访问者自己处理，无论增加多少种类型，这种方式不需要在使用处进行修改
	f := make([]NewFruit, 2)
	f[0] = NewApple{price: 1}
	f[1] = NewBanana{price: 1}
	v := NewNormalVisitor{}
	for _, fruit := range f {
		fruit.Accept(v)
	}
}
