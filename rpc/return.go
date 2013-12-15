
type ReturnWriter struct {
	c *Connection
	ret Return
	exports []Export
	retained []uint32
	promise Dispatcher
	next chan Call
	cancel func()
	answer *answer
}

type Capability struct {
	c *Connection
	id uint32
}

type CallersFaultError interface {
	IsCallersFault() bool
}

type PermanentError interface {
	IsPermanent() bool
}

type OverloadedError interface {
	IsOverloaded() bool
}

func (r *ReturnWriter) AddExport(d Dispatcher) uint32 {
	for i, d := range r.c.exports {
		if d == nil {
			r.c.exports[i] = d
			r.exports = append(r.exports, Export{i, d})
			return i
		}
	}

	i := len(r.c.exports)
	r.c.exports = append(r.c.exports, d)
	r.exports = append(r.exports, Export{i, d})
	return i
}

func (r *ReturnWriter) Buffer() *capn.Segment { return r.ret.Segment }
func (r *ReturnWriter) SetAnswer(ans capn.Object) { r.ret.SetAnswer(ans) }

func (r *ReturnWriter) SetException(err error) {
	exc := NewException(r.ret.Segment)
	exc.SetReason(err.Error())

	if q, ok := err.(CallersFaultError); ok {
		exc.SetIsCallersFault(q.IsCallersFault())
	}
	if q, ok := err.(PermanentError); ok {
		exc.SetIsPermanent(q.IsPermanent())
	}
	if q, ok := err.(OverloadedError); ok {
		exc.SetIsOverloaded(q.IsOverloaded())
	}

	r.ret.SetException(exc)
}

func (r *ReturnWriter) SetPromise(d Dispatcher) {
	r.promise = d
}

func (r *ReturnWriter) RetainCapability(c capn.Capability) {
	r.retain = append(r.retain, c.id)
}

func (r *ReturnWriter) ReplyLater(cancel func()) {
	r.cancel = cancel
	r.next = make(chan Call, 10)
	r.answer.promise = make(chan Call, 10)
}

func (r *ReturnWriter) Send() {
	if len(r.retain) > 0 {
		caps := NewListU32(len(r.retain))
		for i := range caps {
			caps[i] = r.retain.id
		}
		r.ret.SetRetainedCaps(caps)
	}

	r.c.sendlk.Lock()
	r.ret.Segment.WriteTo(r.c.rw)
	r.c.sendlk.Unlock()

	r.c.recvlk.Lock()
	r.answer.answer = r.promise
	r.c.recvlk.Unlock()

	for c := range r.next {
	}
}

	Dispatch(w ReplyWriter, method uint16, args capn.Object)

type myfoo interface {
	dofoo(w ReplyWriter, v int) (myfoo, error)
}

func DispatchMyfoo(obj interface{}, w ReplyWriter, method uint16, args capn.Object) {
	switch method {
	case MYFOO_DOFOO:
		myfoo, err := obj.(myfoo).dofoo(w, args.Get32(0))
		if err != nil {
			w.SetException(err)
		}
	}
}
