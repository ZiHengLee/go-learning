package 状态模式

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	machine := InitialMarioStateMachine()
	fmt.Println(machine.State(), machine.Score())
	machine.ObtainMushRoom()
	fmt.Println(machine.State(), machine.Score())
	machine.ObtainFireFlower()
	fmt.Println(machine.State(), machine.Score())
	machine.MeetMonster()
	fmt.Println(machine.State(), machine.Score())
}
