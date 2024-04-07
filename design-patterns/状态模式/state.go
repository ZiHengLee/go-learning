package 状态模式

type State string
type MarioStateMachine struct {
	score      int
	marioState IMarioState
}

const (
	StateSmall     = State("Small")
	StateSuper     = State("Super")
	StateSuperFire = State("SuperFire")
	StateDead      = State("Dead")
)

var (
	SmallMarioInstance     = &SmallMario{}
	SuperMarioInstance     = &SuperMario{}
	SuperFireMarioInstance = &SuperFireMario{}
	DeadMarioInstance      = &DeadMario{}
)

func InitialMarioStateMachine() *MarioStateMachine {
	return &MarioStateMachine{
		score:      0,
		marioState: &SmallMario{},
	}
}

func (m *MarioStateMachine) ObtainMushRoom() {
	m.marioState.obtainMushRoom(m)
}
func (m *MarioStateMachine) ObtainFireFlower() {
	m.marioState.obtainFireFlower(m)
}
func (m *MarioStateMachine) MeetMonster() {
	m.marioState.meetMonster(m)
}
func (m *MarioStateMachine) Score() int {
	return m.score
}
func (m *MarioStateMachine) State() State {
	return m.marioState.getState()
}

type IMarioState interface {
	getState() State
	obtainMushRoom(*MarioStateMachine)
	obtainFireFlower(*MarioStateMachine)
	meetMonster(*MarioStateMachine)
}

type SmallMario struct{}

func (m *SmallMario) getState() State {
	return StateSmall
}

func (m *SmallMario) obtainMushRoom(sm *MarioStateMachine) {
	sm.score += 10
	sm.marioState = SuperMarioInstance
}

func (m *SmallMario) obtainFireFlower(sm *MarioStateMachine) {
	// will not happen
}

func (m *SmallMario) meetMonster(sm *MarioStateMachine) {
	sm.marioState = DeadMarioInstance
}

type SuperMario struct{}

func (m *SuperMario) getState() State {
	return StateSuper
}

func (m *SuperMario) obtainMushRoom(sm *MarioStateMachine) {
	sm.score += 20
}

func (m *SuperMario) obtainFireFlower(sm *MarioStateMachine) {
	sm.score += 40
	sm.marioState = SuperFireMarioInstance
}

func (m *SuperMario) meetMonster(sm *MarioStateMachine) {
	sm.marioState = SmallMarioInstance
}

type SuperFireMario struct {
}

func (m *SuperFireMario) getState() State {
	return StateSuperFire
}

func (m *SuperFireMario) obtainMushRoom(sm *MarioStateMachine) {
	sm.score += 40
}

func (m *SuperFireMario) obtainFireFlower(sm *MarioStateMachine) {
	sm.score += 80
}

func (m *SuperFireMario) meetMonster(sm *MarioStateMachine) {
	sm.marioState = SuperMarioInstance
}

type DeadMario struct {
}

func (m *DeadMario) getState() State {
	return StateDead
}

func (m *DeadMario) obtainMushRoom(sm *MarioStateMachine) {
}

func (m *DeadMario) obtainFireFlower(sm *MarioStateMachine) {
}

func (m *DeadMario) meetMonster(sm *MarioStateMachine) {

}
