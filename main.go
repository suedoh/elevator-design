package main

import (
	"fmt"
)

type State int
type Direction int

const (
	MOVING State = iota
	IDLE
	STOPPED
)

const (
	DOWN Direction = iota
	UP
)

type Elevator struct {
	CurrentFloor int
	Direction    Direction
	State        State
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
	e := helper.ExternalRequest{UP, 2}
	fmt.Println(e.GetDirectionToGo())
	e.SetDirectionToGo(DOWN)
	fmt.Println(e.GetDirectionToGo())
	// fmt.Println(UP, DOWN)
}
