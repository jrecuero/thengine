package main

type StateMachineType int

const (
	WaitingSM StateMachineType = iota
	StartSM
	PlayerSM
	EnemySM
	EndSM
)

var (
	TheStateMachine *StateMachine = NewStateMachine("state-machine/1")
)

type StateMachine struct {
	name  string
	state StateMachineType
}

func NewStateMachine(name string) *StateMachine {
	return &StateMachine{
		name:  name,
		state: WaitingSM,
	}
}

func (s *StateMachine) GetState() StateMachineType {
	return s.state
}

func (s *StateMachine) IsEnd() bool {
	return s.state == EndSM
}

func (s *StateMachine) IsEnemy() bool {
	return s.state == EnemySM
}

func (s *StateMachine) IsPlayer() bool {
	return s.state == PlayerSM
}

func (s *StateMachine) IsStart() bool {
	return s.state == StartSM
}
func (s *StateMachine) IsWaiting() bool {
	return s.state == WaitingSM
}

func (s *StateMachine) Next() StateMachineType {
	if s.state == EndSM {
		s.state = WaitingSM
	} else {
		s.state++
	}
	return s.state
}
