package main

import (
	"fmt"

	"github.com/suedoh/elevator-design/elevator"
)

func main() {
	er := elevator.NewExternalRequest(elevator.UP, 0)
	ir := elevator.NewInternalRequest(5)
	request := elevator.NewRequest(ir, er)

	fmt.Println(request.GetExternalRequest())
	fmt.Println(request.GetInternalRequest())
	// elevator := elevator.Elevator{0, elevator.UP, elevator.IDLE}
	// fmt.Println(elevator.String())
}
