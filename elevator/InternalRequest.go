package elevator

type InternalRequest struct {
	destinationFloor int
}

func NewInternalRequest(df int) *InternalRequest  {
    return &InternalRequest{
        destinationFloor: df,
    }
}

type InternalRequester interface {
	GetDestinationFloor()
	SetDestinationFloor()
}

func (i InternalRequest) GetDestinationFloor() int {
	return i.destinationFloor
}

func (i *InternalRequest) SetDestinationFloor(d int) {
	i.destinationFloor = d
}
