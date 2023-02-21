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
	DirectionToGo Direction
	SourceFloor   int
}

func (e *ExternalRequest) SetDirectionToGo(d Direction) {
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
