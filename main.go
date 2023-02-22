package main

import (
	"fmt"

	"github.com/suedoh/elevator-design/elevator"
)

const (
	MOVING elevator.State = iota
	IDLE
	STOPPED
)

const (
	DOWN elevator.Direction = iota
	UP
)

type Elevator struct {
	CurrentFloor int
	Direction    elevator.Direction
	State        elevator.State
}

type Request interface {
	InternalRequest()
	ExternalRequest()
	Request()
	GetInternalRequest()
	SetInternalRequest()
	GetExternalRequest()
	SetExternalRequest()
	CompareTo()
}

func main() {
	e := elevator.ExternalRequest{DirectionToGo: UP, SourceFloor: 2}
	fmt.Println(e.GetDirectionToGo())
	e.SetDirectionToGo(DOWN)
	fmt.Println(e.GetDirectionToGo())
	// fmt.Println(UP, DOWN)
}
