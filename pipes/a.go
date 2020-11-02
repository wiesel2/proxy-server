package pipes

// MsgState export
type MsgState int

//
const (
	msgStateOpen MsgState = 1 << iota
	msgStateClosed
)

// pipes

// Msg export
type Msg struct {
	//
	body interface{}
	//
}

// MsgPipe export
type MsgPipe struct {
	cIn   chan *Msg
	cOut  chan *Msg
	name  string
	state MsgState
	done  chan struct{}
}

// NewMsgPipe export
func NewMsgPipe() *MsgPipe {
	return new(MsgPipe)
}

// Close export
func (mp *MsgPipe) Close() {
	defer func() {
		if r := recover(); r != nil {
			// TODO, add log here
		}
	}()
	close(mp.done)
	mp.state = msgStateClosed
}

// Closed export
func (mp *MsgPipe) Closed() bool {
	select {
	case <-mp.done:
		return true
	default:
		return false
	}
}

func (mp *MsgPipe) GetIn() *Msg {
	return mp.get(false)
}

func (mp *MsgPipe) get(out bool) *Msg {
	var c chan *Msg
	if out {
		c = mp.cOut
	} else {
		c = mp.cIn
	}

	for {
		if mp.Closed() {
			return nil
		}
		select {
		case m := <-c:
			return m
		}
	}
}

func (mp *MsgPipe) GetOut() *Msg {
	return mp.get(true)
}
