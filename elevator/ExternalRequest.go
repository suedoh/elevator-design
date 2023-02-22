package elevator

type State int
type Direction int

type ExternalRequester interface {
	SetDirectionToGo()
	GetDirectionToGo()
	GetSourceFloor()
	SetSourceFloor()
}

type ExternalRequest struct {
	directionToGo Direction
	sourceFloor   int
}

func NewExternalRequest(dtg Direction, sf int) *ExternalRequest  {
    return &ExternalRequest{
        directionToGo: dtg,
        sourceFloor: sf,
    }
}

func (e *ExternalRequest) SetDirectionToGo(d Direction) {
	e.directionToGo = d
}

func (e ExternalRequest) GetDirectionToGo() Direction {
	return e.directionToGo
}

func (e ExternalRequest) GetSourceFloor() int {
	return e.sourceFloor
}

func (e *ExternalRequest) SetSourceFloor(s int) {
	e.sourceFloor = s
}
