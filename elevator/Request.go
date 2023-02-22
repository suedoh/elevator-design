package elevator

type Requester interface {
	GetInternalRequest()
	SetInternalRequest()
	GetExternalRequest()
	SetExternalRequest()
	CompareTo()
}

type Request struct {
    internalRequest *InternalRequest
    externalRequest *ExternalRequest
}

func NewRequest(ir *InternalRequest, er *ExternalRequest) *Request  {
    return &Request{
        internalRequest: ir,
        externalRequest: er,
    } 
}

func (r *Request) GetInternalRequest() *InternalRequest  {
    return r.internalRequest
}

func (r *Request) SetInternalRequest(ir *InternalRequest)  {
    r.internalRequest = ir
}

func (r *Request) GetExternalRequest() *ExternalRequest  {
    return r.externalRequest
}

func (r *Request) SetExternalRequest(er *ExternalRequest)  {
    r.externalRequest = er
}

func (r *Request) CompareTo(req *Request) int {
    switch {
    case r.GetInternalRequest().GetDestinationFloor() == req.GetInternalRequest().GetDestinationFloor():
        return 0 
    case r.GetInternalRequest().GetDestinationFloor() > req.GetInternalRequest().GetDestinationFloor():
        return 1
    default:
        return -1
    }
}
