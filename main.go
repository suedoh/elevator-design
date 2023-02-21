package main

import "fmt"

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
   Direction Direction
   State State
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

type ExternalRequester interface {
   SetDirectionToGo()
   GetDirectionToGo()
   GetSourceFloor()
   SetSourceFloor()
}

type ExternalRequest struct {
   DirectionToGo Direction
   SourceFloor int
}

func (e *ExternalRequest) SetDirectionToGo(d Direction)  {
   e.DirectionToGo = d
}

func (e ExternalRequest) GetDirectionToGo() Direction {
   return e.DirectionToGo
}

func (e ExternalRequest) GetSourceFloor() int {
   return e.SourceFloor 
}

func (e *ExternalRequest) SetSourceFloor(s int) {
    e.SourceFloor = s
}

func main()  {
   e := ExternalRequest{UP, 2}
   fmt.Println(e.GetDirectionToGo())
   e.SetDirectionToGo(DOWN) 
   fmt.Println(e.GetDirectionToGo())
//   fmt.Println(UP, DOWN)
}

