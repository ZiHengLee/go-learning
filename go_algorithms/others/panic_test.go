package others

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("正常退出")
	}()
	defer func() {
		err := recover() // recover()内置函数,可以捕获到异常
		if err != nil {  // 捕获到异常
			fmt.Println("err=", err)
			return
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}
