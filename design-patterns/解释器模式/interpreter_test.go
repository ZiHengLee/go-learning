package 解释器模式

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	add := AddExpression{}
	mul := MulExpression{}
	p := Pythagoras{expr1: add, expr2: mul}
	fmt.Println(p.interperter(3, 4))
}
