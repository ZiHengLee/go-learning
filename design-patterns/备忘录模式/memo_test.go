package 备忘录模式

import (
	"fmt"
	"testing"
)

func TestMemo(t *testing.T) {
	cal := Calculator{result: 0}
	cal.Add(10)
	cal.Sub(2)
	fmt.Println(cal.GetResult())
	calMemo := cal.Save()

	newCal := Calculator{result: 0}
	newCal.Restore(calMemo)
	fmt.Println(newCal.GetResult())
}
