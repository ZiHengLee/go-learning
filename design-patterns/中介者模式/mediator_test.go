package 中介者模式

import (
	"testing"
)

func TestMediator(t *testing.T) {
	stationManager := newStationManger()

	trainA := &TrainA{
		mediator: stationManager,
	}
	trainB := &TrainB{
		mediator: stationManager,
	}

	trainA.arrive()
	trainB.arrive()
	trainA.depart()
}
