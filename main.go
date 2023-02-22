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
    er := elevator.NewExternalRequest(UP, 0)
    ir := elevator.NewInternalRequest(5)
    request := elevator.NewRequest(ir, er)

	fmt.Println(request.GetExternalRequest())
	fmt.Println(request.GetInternalRequest())
    elevator := elevator.Elevator{0, UP, IDLE}
	fmt.Println(elevator.String())
}

