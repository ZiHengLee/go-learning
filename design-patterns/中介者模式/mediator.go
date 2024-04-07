package 中介者模式

import "fmt"

type Train interface {
	arrive()
	depart()
	permitArrival()
}

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManger() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

type TrainA struct {
	mediator Mediator
}

func (g *TrainA) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("A: Arrival blocked, waiting")
		return
	}
	fmt.Println("A: Arrived")
}

func (g *TrainA) depart() {
	fmt.Println("A: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *TrainA) permitArrival() {
	fmt.Println("A: Arrival permitted, arriving")
	g.arrive()
}

type TrainB struct {
	mediator Mediator
}

func (g *TrainB) arrive() {
	if !g.mediator.canArrive(g) {
		fmt.Println("B: Arrival blocked, waiting")
		return
	}
	fmt.Println("B: Arrived")
}

func (g *TrainB) depart() {
	fmt.Println("B: Leaving")
	g.mediator.notifyAboutDeparture()
}

func (g *TrainB) permitArrival() {
	fmt.Println("B: Arrival permitted, arriving")
	g.arrive()
}
