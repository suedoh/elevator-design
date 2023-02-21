package elevator

type InternalRequest struct {
	DestinationFloor int
}

type InternalRequester interface {
	GetDetinationFloor()
	SetDetinationFloor()
}

func (i InternalRequest) GetDetinationFloor() int {
	return i.DestinationFloor
}

func (i *InternalRequest) SetDetinationFloor(d int) {
	i.DestinationFloor = d
}
