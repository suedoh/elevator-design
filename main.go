package main

import(
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

func main() {
    elevator := elevator.Elevator{0, UP, IDLE}
	fmt.Println(elevator.String())
}

