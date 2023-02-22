package elevator

import (
	"fmt"
	"time"
)

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
	CurrentFloor     int
	CurrentDirection Direction
	CurrentState     State
	CurrentJobs      TreeSet
	UpPendingJobs    TreeSet
	DownPendingJobs  TreeSet
}

func (e *Elevator) String() string {
	return fmt.Sprintf("Floor: %d, Direction: %v, State: %v", e.CurrentFloor, e.CurrentDirection, e.CurrentState)
}

func (e *Elevator) StartElevator() {
	for {
		if e.CheckIfJob() {
			if e.CurrentDirection == UP {
				r := e.CurrentJobs.tree[0]
				e.ProcessUpRequest(r)
				if !e.CheckIfJob() {
					e.AddPendingDownJobsToCurrentJobs()
				}
			}

			if e.CurrentDirection == DOWN {
				r := e.CurrentJobs.tree[0]
				e.ProcessDownRequest(r)
				if !e.CheckIfJob() {
					e.AddPendingDownJobsToCurrentJobs()
				}
			}
		}

	}
}

func (e *Elevator) ProcessUpRequest(r Request) {
	startFloor := e.CurrentFloor
	if startFloor < r.GetExternalRequest().GetSourceFloor() {
		for i := startFloor; i <= r.GetExternalRequest().GetSourceFloor(); i++ {
			time.Sleep(10 * time.Second)
			fmt.Printf("We have reached floor -- %d\n", i)
			e.CurrentFloor = i
		}
	}

	fmt.Println("Reached Source Floor--Opening Door")

	startFloor = e.CurrentFloor
	if startFloor < r.GetInternalRequest().GetDestinationFloor() {
		for i := startFloor; i <= r.GetInternalRequest().GetDestinationFloor(); i++ {
			time.Sleep(10 * time.Second)
			fmt.Printf("We have reached floor -- %d\n", i)
			e.CurrentFloor = i
			if e.CheckIfNewJobCanBeProcessed(r) {
				break
			}
		}
	}

}

func (e *Elevator) ProcessDownRequest(r Request) {

}

func (e *Elevator) CheckIfNewJobCanBeProcessed(r Request) bool {
	return false

}

func (e *Elevator) AddPendingDownJobsToCurrentJobs() {

}

func (e *Elevator) CheckIfJob() bool {
	return e.CurrentJobs.length > 0
}
