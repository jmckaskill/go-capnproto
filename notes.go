/*
Call
- has a question ID
- Can target either:
	- known cap (import)
	- promised answer
		- includes question ID
		- + path to cap

Return
- refers to a question ID

Finish
indicates we've received the return so that the caps can be optionally
auto-released
*/



// From a user perspective:

root := get_root()

// want only a single value returned so things can be chained
promise := root.callFunc(...)

// promise hasn't been resolved at this point

// foo is still a promise
foo := promise.FooStruct()

// cap is still a promise
cap := foo.BarCap()

// this can proceed immediately
cap.callOtherFunc()


// calls/returns come in with a separate table of cap descriptors
// these cap descriptors can indicate a promise

// also when making a call, the call is sent out and a promise is
// generated which will be later filled out when the question returns


root.callFunc() returns

obj := Object {
	Segment: nil,
	promise: {
		type: receiverAnswer,
		id: questionId,
		cond: sync.NewCond(),
		paramlk: nil,
		err: nil,
	}
	ops: nil,
}

obj.FooStruct() then returns
Object {
	Segment: nil,
	promise: {
		type: receiverAnswer,
		id: questionId,
	}
	ops: []uint16{fooOffset},
}

// on the call to Resolve, the user calls
promise.lk.Lock()
for !promise.type == receiverAnswer {
	promise.cond.Wait()
}
promise.lk.Unlock()

// the receiver thread calls
promise.lk.Lock()
promise.type = cap.type
promise.id = cap.id
promise.err = cap.err
// etc
promise.cond.Broadcast()
promise.lk.Unlock()


// For server usage
need a means of looking up persistent objects

SetPersistentLookup(...)
type SturdyLookup interface {
	RestoreSturdy(SturdyRef) (Capability, error)
	DeleteSturdy(SturdyRef) error
}

// saving is done via a special interface
type Saver interface {
	SaveSturdy() (SturdyRef, error)
}

// should errors be chained or not?


conn := NewConnection(net.Accept()) ...
conn.SetSturdyRestorer(...)

// ....
// client sends Restore(foo)

Restore(foo) -> (foo, nil)
foo is a Capability
// it must minimally meet
type Exporter interface {
	ExportInterface() InterfaceId
}
// or must it? This seems to be already indicated by the schema.

// it gets assigned an export Id (say 0)
// and we send out a Resolve with a CapDescriptor{senderHosted: 0}
// that then gets put in the export table with a ref count of 1

// we maintain a table of exports to ids
map[Capability]*export
// and id to export
[]*export

type Interface func(ctx CallContext, obj Capability, method uint16, params Object) error

type export struct {
	ref int
	iface Interface
	obj Capability
}

type Capability interface {
	// This way objects exporting multiple interfaces are forced
	// into using proxy object. Otherwise wouldn't be able to track
	// what interfaces are allowed for a given export.
	Interface() Interface
}

type MyInterface interface {
	Capability
	... my interface functions
}

// on the client side we have to maintain
// an import table
map[ImportId]*import

type import struct {
	ref int
	id int
}

func (*import) Call(iface uint64, method uint16, params Object)

// a promised answer table
[]*promise
