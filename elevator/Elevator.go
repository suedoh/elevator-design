package elevator

import (
    "fmt"
)

type Elevator struct {
    CurrentFloor        int
    CurrentDirection    Direction
    CurrentState        State
}

func (e *Elevator) String() string {
    return fmt.Sprintf("Floor: %d, Direction: %v, State: %v", e.CurrentFloor, e.CurrentDirection, e.CurrentState)
}
