package rpc

type ReplyWriter interface {
	Segment() *capn.Segment
	RetainCap(obj interface{})
	ExportCap(obj interface{}) CapDescriptor
	SetAnswer(c capn.Object)
}

type Interface interface {
	Id() uint64
	Dispatch(w ReplyWriter, method uint16, args capn.Object)
}

type LocalCapability interface {
	Interfaces() []Interface
}

type Export struct {
	Id uint32
	Dispatcher Dispatcher
}

type answer struct {
	promise chan Call
	answer Dispatcher
	exports []Export
}

type export struct {
	id uint32
	ref int
}

type Connection struct {
	questions []Interface
	answers map[uint32]*answer
	exports []Dispatcher
	exports map[Interface]*export
	imports map[uint32]interface{}
	rw io.ReadWriteCloser
}

func NewConnection(rw io.ReadWriteCloser) *Connection {
}

func (c *Connection) close() {
	c.rw.Close()
	c.questions = nil
	c.answers = nil
	c.exports = nil
	c.imports = nil
}

func (c *Connection) sendUnimplemented(rx Message) error {
	buf := capn.NewBuffer(nil)
	tx := NewMessage(buf)
	tx.SetUnimplemented(rx)
	_, err := txbuf.WriteTo(c.rw)
	return err
}

func (c *connectionSession) AddExport(d Dispatcher) uint32 {
}

func (c *Connection) reader() {
	rxbuf := bytes.Buffer{}
	txbuf := capn.NewBuffer(nil)

	defer c.close()

	for {
		rxbuf.Reset()
		txbuf.Data = txbuf.Data[:0]

		seg, err := capn.ReadFromStream(c.rw, &rxbuf)
		if err != nil {
			return
		}

		switch rx := ReadMessage(seg); rx.which() {
		case MESSAGE_UNIMPLEMENTED:
		case MESSAGE_ABORT:
			return

		case MESSAGE_CALL:
			call := rx.Call()
			iface := Interface(nil)

			switch call.Target().which() {
			case CALLTARGET_EXPORTEDCAP:
				if call.Target().ExportedCap() < len(c.exports) {
					iface = c.exports[call.Target().ExportedCap()]
				}
			case CALLTARGET_PROMISEDANSWER:
				if call.Target().ExportedCap() < len(c.questions) {
					iface = c.questions[call.Target().PromisedAnswer()]
				}
			default:
				if c.sendUnimplemented(rx) != nil {
					return
				}
			}

			tx := txbuf.NewRoot()
			msg := NewMessage(txbuf)
			ret := NewReturn(txbuf)
			tx.Set(msg)
			msg.SetReturn(ret)
			msg.SetQuestionId(call.QuestionId())

			if iface != nil {
				obj, exports := iface.Dispatch(ret, call.MethodId(), call.Request())
				c.answers[call.QuestionId()] = answer{obj, exports}
			} else {
				ret.SetCanceled()
			}

			if _, err := txbuf.WriteTo(c.rw); err != nil {
				return
			}

		case MESSAGE_FINISH:
			fin := rx.Finish()
			if fin.QuestionId() < len(c.answers) {
				ans := c.answers[fin.QuestionId()]
				c.answers[fin.QuestionId()] = answer{}
				for _, d := range ans.exports {

					retain := false
					for _, id := range fin.RetainedCaps().ToArray() {
						if id == d.id {
							retain = true
							break
						}
					}

					if !retain {
						c.exports[d.Id] = Export{}
						d.disposer.Dispose()
					}
				}
			}

		case MESSAGE_RETURN:
		case MESSAGE_RESOLVE:
		case MESSAGE_RELEASE:
		default:
			if c.sendUnimplemented(rx) != nil {
				return
			}
		}
	}
}

