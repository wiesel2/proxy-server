package pipes

import "context"

// Middelware interface

// MiddelwareInt interface
type MiddelwareInt interface {
	ProcessIn(m interface{}) (bool, interface{})
	ProcessOut(m interface{}) (bool, interface{})
	Run()
	Closed() bool
	Close()
	Name() string
}

const msgCountMWName = "count middleware"

// MsgCountMW export
type MsgCountMW struct {
	TotalInCount  int64
	TotalOutCount int64
	inC           chan interface{}
	outC          chan interface{}
	done          chan struct{}
	ctx           context.Context
}

// Name export
func (mc *MsgCountMW) Name() string {
	return msgCountMWName
}

//
func (mc *MsgCountMW) ProcessIn(m interface{}) (bool, interface{}) {
	mc.TotalInCount++
	return true, nil
}

//
func (mc *MsgCountMW) ProcessOut(m interface{}) (bool, interface{}) {
	mc.TotalOutCount++
	return true, nil
}

//
func (mc *MsgCountMW) Closed() bool {
	select {
	case <-mc.done:
		return true
	default:
		return false
	}
}

//
func (mc *MsgCountMW) Close() {
	close(mc.done)
}

//
func (mc *MsgCountMW) Run() {
	for {
		if mc.Closed() {
			// TODO, add log here
			return
		}
		select {
		case d := <-mc.inC:
			mc.ProcessIn(d)
		case d := <-mc.outC:
			mc.ProcessOut(d)
		default:
			// do nothing?
		}
	}
}
