package lua

import "fmt"

type Hooker interface {
	call(L *LState, cf *callFrame)
	String() string
}

type LHook struct {
	callback *LFunction
	line     int
}

func newLHook(callback *LFunction, line int) *LHook {
	return &LHook{
		callback: callback,
		line:     line,
	}
}

func (lh *LHook) call(L *LState, cf *callFrame) {
	currentline := cf.Fn.Proto.DbgSourcePositions[cf.Pc-1]
	if currentline != 0 && cf.Fn != lh.callback && currentline != L.prevline {
		L.reg.Push(lh.callback)
		L.reg.Push(LString("l"))
		L.reg.Push(LNumber(currentline))
		L.callR(2, 0, -1)
		L.prevline = currentline
	}
}

func (lh *LHook) String() string {
	return fmt.Sprintf("hook: %p", lh)
}

type CHook struct {
	callback *LFunction
}

func newCHook(callback *LFunction) *CHook {
	return &CHook{
		callback: callback,
	}
}

func (ch *CHook) call(L *LState, cf *callFrame) {

}

func (ch *CHook) String() string {
	return fmt.Sprintf("hook: %p", ch)
}

type CTHook struct {
	callback *LFunction
	count    int
	counter  int
}

func newCTHook(callback *LFunction, count int) *CTHook {
	return &CTHook{
		callback: callback,
		count:    count,
		counter:  0,
	}
}

func (ct *CTHook) call(L *LState, cf *callFrame) {
	if ct.count <= 0 {
		return
	}
	ct.counter++
	if ct.counter >= ct.count {
		ct.counter = 0
		L.reg.Push(ct.callback)
		L.reg.Push(LString("c"))
		L.callR(1, 0, -1)
	}
}

func (ct *CTHook) String() string {
	return fmt.Sprintf("hook: %p", ct)
}

type RHook struct {
	callback *LFunction
}

func newRHook(callback *LFunction) *RHook {
	return &RHook{
		callback: callback,
	}
}

func (rh *RHook) call(L *LState, cf *callFrame) {

}

func (rh *RHook) String() string {
	return fmt.Sprintf("hook: %p", rh)
}
