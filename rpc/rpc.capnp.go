package capn

// AUTO GENERATED - DO NOT EDIT
import (
	C "github.com/jmckaskill/go-capnproto"
	"unsafe"
)

type Message C.Struct
type Message_which uint16

const (
	MESSAGE_UNIMPLEMENTED Message_which = 0
	MESSAGE_ABORT                       = 1
	MESSAGE_CALL                        = 2
	MESSAGE_RETURN                      = 3
	MESSAGE_FINISH                      = 4
	MESSAGE_RESOLVE                     = 5
	MESSAGE_RELEASE                     = 6
	MESSAGE_SAVE                        = 7
	MESSAGE_RESTORE                     = 8
	MESSAGE_DELETE                      = 9
	MESSAGE_PROVIDE                     = 10
	MESSAGE_ACCEPT                      = 11
	MESSAGE_JOIN                        = 12
)

func NewMessage(s *C.Segment) Message    { return Message(s.NewStruct(8, 1)) }
func ReadMessage(s *C.Segment) Message   { return Message(s.Root(0).ToStruct()) }
func (s Message) which() Message_which   { return Message_which(C.Struct(s).Get16(0)) }
func (s Message) Unimplemented() Message { return Message(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetUnimplemented(v Message) {
	C.Struct(s).Set16(0, 0)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Abort() Exception     { return Exception(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetAbort(v Exception) { C.Struct(s).Set16(0, 1); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Call() Call           { return Call(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetCall(v Call)       { C.Struct(s).Set16(0, 2); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Return() Return       { return Return(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetReturn(v Return)   { C.Struct(s).Set16(0, 3); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Finish() Finish       { return Finish(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetFinish(v Finish)   { C.Struct(s).Set16(0, 4); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Resolve() Resolve     { return Resolve(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetResolve(v Resolve) { C.Struct(s).Set16(0, 5); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Release() Release     { return Release(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetRelease(v Release) { C.Struct(s).Set16(0, 6); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Save() Save           { return Save(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetSave(v Save)       { C.Struct(s).Set16(0, 7); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Restore() Restore     { return Restore(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetRestore(v Restore) { C.Struct(s).Set16(0, 8); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Delete() Delete       { return Delete(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetDelete(v Delete)   { C.Struct(s).Set16(0, 9); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Provide() Provide     { return Provide(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetProvide(v Provide) {
	C.Struct(s).Set16(0, 10)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Accept() Accept     { return Accept(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetAccept(v Accept) { C.Struct(s).Set16(0, 11); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Join() Join         { return Join(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetJoin(v Join)     { C.Struct(s).Set16(0, 12); C.Struct(s).SetObject(0, C.Object(v)) }

type Message_List C.PointerList

func NewMessageList(s *C.Segment, sz int) Message_List {
	return Message_List(s.NewCompositeList(8, 1, sz))
}
func (s Message_List) Len() int         { return C.PointerList(s).Len() }
func (s Message_List) At(i int) Message { return Message(C.PointerList(s).At(i).ToStruct()) }
func (s Message_List) ToArray() []Message {
	return *(*[]Message)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Call C.Struct
type CallTarget Call
type CallTarget_which uint16

const (
	CALLTARGET_EXPORTEDCAP    CallTarget_which = 0
	CALLTARGET_PROMISEDANSWER                  = 1
)

func NewCall(s *C.Segment) Call              { return Call(s.NewStruct(24, 2)) }
func ReadCall(s *C.Segment) Call             { return Call(s.Root(0).ToStruct()) }
func (s Call) QuestionId() uint32            { return C.Struct(s).Get32(0) }
func (s Call) SetQuestionId(v uint32)        { C.Struct(s).Set32(0, v) }
func (s Call) Target() CallTarget            { return CallTarget(s) }
func (s CallTarget) which() CallTarget_which { return CallTarget_which(C.Struct(s).Get16(8)) }
func (s CallTarget) ExportedCap() uint32     { return C.Struct(s).Get32(4) }
func (s CallTarget) SetExportedCap(v uint32) { C.Struct(s).Set16(8, 0); C.Struct(s).Set32(4, v) }
func (s CallTarget) PromisedAnswer() PromisedAnswer {
	return PromisedAnswer(C.Struct(s).GetObject(0).ToStruct())
}
func (s CallTarget) SetPromisedAnswer(v PromisedAnswer) {
	C.Struct(s).Set16(8, 1)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Call) InterfaceId() uint64     { return C.Struct(s).Get64(16) }
func (s Call) SetInterfaceId(v uint64) { C.Struct(s).Set64(16, v) }
func (s Call) MethodId() uint16        { return C.Struct(s).Get16(10) }
func (s Call) SetMethodId(v uint16)    { C.Struct(s).Set16(10, v) }
func (s Call) Request() C.Object       { return C.Struct(s).GetObject(1) }
func (s Call) SetRequest(v C.Object)   { C.Struct(s).SetObject(1, v) }

type Call_List C.PointerList

func NewCallList(s *C.Segment, sz int) Call_List { return Call_List(s.NewCompositeList(24, 2, sz)) }
func (s Call_List) Len() int                     { return C.PointerList(s).Len() }
func (s Call_List) At(i int) Call                { return Call(C.PointerList(s).At(i).ToStruct()) }
func (s Call_List) ToArray() []Call              { return *(*[]Call)(unsafe.Pointer(C.PointerList(s).ToArray())) }

type Return C.Struct
type Return_which uint16

const (
	RETURN_ANSWER                Return_which = 0
	RETURN_EXCEPTION                          = 1
	RETURN_CANCELED                           = 2
	RETURN_UNSUPPORTEDPIPELINEOP              = 3
)

func NewReturn(s *C.Segment) Return          { return Return(s.NewStruct(8, 2)) }
func ReadReturn(s *C.Segment) Return         { return Return(s.Root(0).ToStruct()) }
func (s Return) which() Return_which         { return Return_which(C.Struct(s).Get16(4)) }
func (s Return) QuestionId() uint32          { return C.Struct(s).Get32(0) }
func (s Return) SetQuestionId(v uint32)      { C.Struct(s).Set32(0, v) }
func (s Return) RetainedCaps() C.ListU32     { return C.ListU32(C.Struct(s).GetObject(0)) }
func (s Return) SetRetainedCaps(v C.ListU32) { C.Struct(s).SetObject(0, C.Object(v)) }
func (s Return) Answer() C.Object            { return C.Struct(s).GetObject(1) }
func (s Return) SetAnswer(v C.Object)        { C.Struct(s).Set16(4, 0); C.Struct(s).SetObject(1, v) }
func (s Return) Exception() Exception        { return Exception(C.Struct(s).GetObject(1).ToStruct()) }
func (s Return) SetException(v Exception) {
	C.Struct(s).Set16(4, 1)
	C.Struct(s).SetObject(1, C.Object(v))
}

type Return_List C.PointerList

func NewReturnList(s *C.Segment, sz int) Return_List { return Return_List(s.NewCompositeList(8, 2, sz)) }
func (s Return_List) Len() int                       { return C.PointerList(s).Len() }
func (s Return_List) At(i int) Return                { return Return(C.PointerList(s).At(i).ToStruct()) }
func (s Return_List) ToArray() []Return {
	return *(*[]Return)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Finish C.Struct

func NewFinish(s *C.Segment) Finish          { return Finish(s.NewStruct(8, 1)) }
func ReadFinish(s *C.Segment) Finish         { return Finish(s.Root(0).ToStruct()) }
func (s Finish) QuestionId() uint32          { return C.Struct(s).Get32(0) }
func (s Finish) SetQuestionId(v uint32)      { C.Struct(s).Set32(0, v) }
func (s Finish) RetainedCaps() C.ListU32     { return C.ListU32(C.Struct(s).GetObject(0)) }
func (s Finish) SetRetainedCaps(v C.ListU32) { C.Struct(s).SetObject(0, C.Object(v)) }

type Finish_List C.PointerList

func NewFinishList(s *C.Segment, sz int) Finish_List { return Finish_List(s.NewCompositeList(8, 1, sz)) }
func (s Finish_List) Len() int                       { return C.PointerList(s).Len() }
func (s Finish_List) At(i int) Finish                { return Finish(C.PointerList(s).At(i).ToStruct()) }
func (s Finish_List) ToArray() []Finish {
	return *(*[]Finish)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Resolve C.Struct
type Resolve_which uint16

const (
	RESOLVE_CAP       Resolve_which = 0
	RESOLVE_EXCEPTION               = 1
	RESOLVE_CANCELED                = 2
)

func NewResolve(s *C.Segment) Resolve   { return Resolve(s.NewStruct(8, 1)) }
func ReadResolve(s *C.Segment) Resolve  { return Resolve(s.Root(0).ToStruct()) }
func (s Resolve) which() Resolve_which  { return Resolve_which(C.Struct(s).Get16(4)) }
func (s Resolve) PromiseId() uint32     { return C.Struct(s).Get32(0) }
func (s Resolve) SetPromiseId(v uint32) { C.Struct(s).Set32(0, v) }
func (s Resolve) Cap() CapDescriptor    { return CapDescriptor(C.Struct(s).GetObject(0).ToStruct()) }
func (s Resolve) SetCap(v CapDescriptor) {
	C.Struct(s).Set16(4, 0)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Resolve) Exception() Exception { return Exception(C.Struct(s).GetObject(0).ToStruct()) }
func (s Resolve) SetException(v Exception) {
	C.Struct(s).Set16(4, 1)
	C.Struct(s).SetObject(0, C.Object(v))
}

type Resolve_List C.PointerList

func NewResolveList(s *C.Segment, sz int) Resolve_List {
	return Resolve_List(s.NewCompositeList(8, 1, sz))
}
func (s Resolve_List) Len() int         { return C.PointerList(s).Len() }
func (s Resolve_List) At(i int) Resolve { return Resolve(C.PointerList(s).At(i).ToStruct()) }
func (s Resolve_List) ToArray() []Resolve {
	return *(*[]Resolve)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Release C.Struct

func NewRelease(s *C.Segment) Release        { return Release(s.NewStruct(8, 0)) }
func ReadRelease(s *C.Segment) Release       { return Release(s.Root(0).ToStruct()) }
func (s Release) Id() uint32                 { return C.Struct(s).Get32(0) }
func (s Release) SetId(v uint32)             { C.Struct(s).Set32(0, v) }
func (s Release) ReferenceCount() uint32     { return C.Struct(s).Get32(4) }
func (s Release) SetReferenceCount(v uint32) { C.Struct(s).Set32(4, v) }

type Release_List C.PointerList

func NewReleaseList(s *C.Segment, sz int) Release_List {
	return Release_List(s.NewCompositeList(8, 0, sz))
}
func (s Release_List) Len() int         { return C.PointerList(s).Len() }
func (s Release_List) At(i int) Release { return Release(C.PointerList(s).At(i).ToStruct()) }
func (s Release_List) ToArray() []Release {
	return *(*[]Release)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Save C.Struct
type SaveTarget Save
type SaveTarget_which uint16

const (
	SAVETARGET_EXPORTEDCAP    SaveTarget_which = 0
	SAVETARGET_PROMISEDANSWER                  = 1
)

func NewSave(s *C.Segment) Save              { return Save(s.NewStruct(16, 1)) }
func ReadSave(s *C.Segment) Save             { return Save(s.Root(0).ToStruct()) }
func (s Save) QuestionId() uint32            { return C.Struct(s).Get32(0) }
func (s Save) SetQuestionId(v uint32)        { C.Struct(s).Set32(0, v) }
func (s Save) Target() SaveTarget            { return SaveTarget(s) }
func (s SaveTarget) which() SaveTarget_which { return SaveTarget_which(C.Struct(s).Get16(8)) }
func (s SaveTarget) ExportedCap() uint32     { return C.Struct(s).Get32(4) }
func (s SaveTarget) SetExportedCap(v uint32) { C.Struct(s).Set16(8, 0); C.Struct(s).Set32(4, v) }
func (s SaveTarget) PromisedAnswer() PromisedAnswer {
	return PromisedAnswer(C.Struct(s).GetObject(0).ToStruct())
}
func (s SaveTarget) SetPromisedAnswer(v PromisedAnswer) {
	C.Struct(s).Set16(8, 1)
	C.Struct(s).SetObject(0, C.Object(v))
}

type Save_List C.PointerList

func NewSaveList(s *C.Segment, sz int) Save_List { return Save_List(s.NewCompositeList(16, 1, sz)) }
func (s Save_List) Len() int                     { return C.PointerList(s).Len() }
func (s Save_List) At(i int) Save                { return Save(C.PointerList(s).At(i).ToStruct()) }
func (s Save_List) ToArray() []Save              { return *(*[]Save)(unsafe.Pointer(C.PointerList(s).ToArray())) }

type Restore C.Struct

func NewRestore(s *C.Segment) Restore    { return Restore(s.NewStruct(8, 1)) }
func ReadRestore(s *C.Segment) Restore   { return Restore(s.Root(0).ToStruct()) }
func (s Restore) QuestionId() uint32     { return C.Struct(s).Get32(0) }
func (s Restore) SetQuestionId(v uint32) { C.Struct(s).Set32(0, v) }
func (s Restore) Ref() C.Object          { return C.Struct(s).GetObject(0) }
func (s Restore) SetRef(v C.Object)      { C.Struct(s).SetObject(0, v) }

type Restore_List C.PointerList

func NewRestoreList(s *C.Segment, sz int) Restore_List {
	return Restore_List(s.NewCompositeList(8, 1, sz))
}
func (s Restore_List) Len() int         { return C.PointerList(s).Len() }
func (s Restore_List) At(i int) Restore { return Restore(C.PointerList(s).At(i).ToStruct()) }
func (s Restore_List) ToArray() []Restore {
	return *(*[]Restore)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Delete C.Struct

func NewDelete(s *C.Segment) Delete     { return Delete(s.NewStruct(8, 1)) }
func ReadDelete(s *C.Segment) Delete    { return Delete(s.Root(0).ToStruct()) }
func (s Delete) QuestionId() uint32     { return C.Struct(s).Get32(0) }
func (s Delete) SetQuestionId(v uint32) { C.Struct(s).Set32(0, v) }
func (s Delete) Ref() C.Object          { return C.Struct(s).GetObject(0) }
func (s Delete) SetRef(v C.Object)      { C.Struct(s).SetObject(0, v) }

type Delete_List C.PointerList

func NewDeleteList(s *C.Segment, sz int) Delete_List { return Delete_List(s.NewCompositeList(8, 1, sz)) }
func (s Delete_List) Len() int                       { return C.PointerList(s).Len() }
func (s Delete_List) At(i int) Delete                { return Delete(C.PointerList(s).At(i).ToStruct()) }
func (s Delete_List) ToArray() []Delete {
	return *(*[]Delete)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Provide C.Struct
type ProvideTarget Provide
type ProvideTarget_which uint16

const (
	PROVIDETARGET_EXPORTEDCAP    ProvideTarget_which = 0
	PROVIDETARGET_PROMISEDANSWER                     = 1
)

func NewProvide(s *C.Segment) Provide              { return Provide(s.NewStruct(16, 2)) }
func ReadProvide(s *C.Segment) Provide             { return Provide(s.Root(0).ToStruct()) }
func (s Provide) QuestionId() uint32               { return C.Struct(s).Get32(0) }
func (s Provide) SetQuestionId(v uint32)           { C.Struct(s).Set32(0, v) }
func (s Provide) Target() ProvideTarget            { return ProvideTarget(s) }
func (s ProvideTarget) which() ProvideTarget_which { return ProvideTarget_which(C.Struct(s).Get16(8)) }
func (s ProvideTarget) ExportedCap() uint32        { return C.Struct(s).Get32(4) }
func (s ProvideTarget) SetExportedCap(v uint32)    { C.Struct(s).Set16(8, 0); C.Struct(s).Set32(4, v) }
func (s ProvideTarget) PromisedAnswer() PromisedAnswer {
	return PromisedAnswer(C.Struct(s).GetObject(0).ToStruct())
}
func (s ProvideTarget) SetPromisedAnswer(v PromisedAnswer) {
	C.Struct(s).Set16(8, 1)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Provide) Recipient() C.Object     { return C.Struct(s).GetObject(1) }
func (s Provide) SetRecipient(v C.Object) { C.Struct(s).SetObject(1, v) }

type Provide_List C.PointerList

func NewProvideList(s *C.Segment, sz int) Provide_List {
	return Provide_List(s.NewCompositeList(16, 2, sz))
}
func (s Provide_List) Len() int         { return C.PointerList(s).Len() }
func (s Provide_List) At(i int) Provide { return Provide(C.PointerList(s).At(i).ToStruct()) }
func (s Provide_List) ToArray() []Provide {
	return *(*[]Provide)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Accept C.Struct

func NewAccept(s *C.Segment) Accept      { return Accept(s.NewStruct(8, 1)) }
func ReadAccept(s *C.Segment) Accept     { return Accept(s.Root(0).ToStruct()) }
func (s Accept) QuestionId() uint32      { return C.Struct(s).Get32(0) }
func (s Accept) SetQuestionId(v uint32)  { C.Struct(s).Set32(0, v) }
func (s Accept) Provision() C.Object     { return C.Struct(s).GetObject(0) }
func (s Accept) SetProvision(v C.Object) { C.Struct(s).SetObject(0, v) }

type Accept_List C.PointerList

func NewAcceptList(s *C.Segment, sz int) Accept_List { return Accept_List(s.NewCompositeList(8, 1, sz)) }
func (s Accept_List) Len() int                       { return C.PointerList(s).Len() }
func (s Accept_List) At(i int) Accept                { return Accept(C.PointerList(s).At(i).ToStruct()) }
func (s Accept_List) ToArray() []Accept {
	return *(*[]Accept)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Join C.Struct

func NewJoin(s *C.Segment) Join       { return Join(s.NewStruct(8, 1)) }
func ReadJoin(s *C.Segment) Join      { return Join(s.Root(0).ToStruct()) }
func (s Join) QuestionId() uint32     { return C.Struct(s).Get32(0) }
func (s Join) SetQuestionId(v uint32) { C.Struct(s).Set32(0, v) }
func (s Join) CapId() uint32          { return C.Struct(s).Get32(4) }
func (s Join) SetCapId(v uint32)      { C.Struct(s).Set32(4, v) }
func (s Join) KeyPart() C.Object      { return C.Struct(s).GetObject(0) }
func (s Join) SetKeyPart(v C.Object)  { C.Struct(s).SetObject(0, v) }

type Join_List C.PointerList

func NewJoinList(s *C.Segment, sz int) Join_List { return Join_List(s.NewCompositeList(8, 1, sz)) }
func (s Join_List) Len() int                     { return C.PointerList(s).Len() }
func (s Join_List) At(i int) Join                { return Join(C.PointerList(s).At(i).ToStruct()) }
func (s Join_List) ToArray() []Join              { return *(*[]Join)(unsafe.Pointer(C.PointerList(s).ToArray())) }

type CapDescriptor C.Struct
type CapDescriptorSenderHosted CapDescriptor
type CapDescriptor_which uint16

const (
	CAPDESCRIPTOR_SENDERHOSTED     CapDescriptor_which = 0
	CAPDESCRIPTOR_SENDERPROMISE                        = 1
	CAPDESCRIPTOR_RECEIVERHOSTED                       = 2
	CAPDESCRIPTOR_RECEIVERANSWER                       = 3
	CAPDESCRIPTOR_THIRDPARTYHOSTED                     = 4
)

func NewCapDescriptor(s *C.Segment) CapDescriptor               { return CapDescriptor(s.NewStruct(8, 1)) }
func ReadCapDescriptor(s *C.Segment) CapDescriptor              { return CapDescriptor(s.Root(0).ToStruct()) }
func (s CapDescriptor) which() CapDescriptor_which              { return CapDescriptor_which(C.Struct(s).Get16(4)) }
func (s CapDescriptor) SenderHosted() CapDescriptorSenderHosted { return CapDescriptorSenderHosted(s) }
func (s CapDescriptor) SetSenderHosted()                        { C.Struct(s).Set16(4, 0) }
func (s CapDescriptorSenderHosted) Id() uint32                  { return C.Struct(s).Get32(0) }
func (s CapDescriptorSenderHosted) SetId(v uint32)              { C.Struct(s).Set32(0, v) }
func (s CapDescriptorSenderHosted) Interfaces() C.ListU64       { return C.ListU64(C.Struct(s).GetObject(0)) }
func (s CapDescriptorSenderHosted) SetInterfaces(v C.ListU64)   { C.Struct(s).SetObject(0, C.Object(v)) }
func (s CapDescriptor) SenderPromise() uint32                   { return C.Struct(s).Get32(0) }
func (s CapDescriptor) SetSenderPromise(v uint32)               { C.Struct(s).Set16(4, 1); C.Struct(s).Set32(0, v) }
func (s CapDescriptor) ReceiverHosted() uint32                  { return C.Struct(s).Get32(0) }
func (s CapDescriptor) SetReceiverHosted(v uint32)              { C.Struct(s).Set16(4, 2); C.Struct(s).Set32(0, v) }
func (s CapDescriptor) ReceiverAnswer() PromisedAnswer {
	return PromisedAnswer(C.Struct(s).GetObject(0).ToStruct())
}
func (s CapDescriptor) SetReceiverAnswer(v PromisedAnswer) {
	C.Struct(s).Set16(4, 3)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s CapDescriptor) ThirdPartyHosted() ThirdPartyCapDescriptor {
	return ThirdPartyCapDescriptor(C.Struct(s).GetObject(0).ToStruct())
}
func (s CapDescriptor) SetThirdPartyHosted(v ThirdPartyCapDescriptor) {
	C.Struct(s).Set16(4, 4)
	C.Struct(s).SetObject(0, C.Object(v))
}

type CapDescriptor_List C.PointerList

func NewCapDescriptorList(s *C.Segment, sz int) CapDescriptor_List {
	return CapDescriptor_List(s.NewCompositeList(8, 1, sz))
}
func (s CapDescriptor_List) Len() int { return C.PointerList(s).Len() }
func (s CapDescriptor_List) At(i int) CapDescriptor {
	return CapDescriptor(C.PointerList(s).At(i).ToStruct())
}
func (s CapDescriptor_List) ToArray() []CapDescriptor {
	return *(*[]CapDescriptor)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type PromisedAnswer C.Struct

func NewPromisedAnswer(s *C.Segment) PromisedAnswer  { return PromisedAnswer(s.NewStruct(8, 1)) }
func ReadPromisedAnswer(s *C.Segment) PromisedAnswer { return PromisedAnswer(s.Root(0).ToStruct()) }
func (s PromisedAnswer) QuestionId() uint32          { return C.Struct(s).Get32(0) }
func (s PromisedAnswer) SetQuestionId(v uint32)      { C.Struct(s).Set32(0, v) }
func (s PromisedAnswer) Transform() PromisedAnswerOp_List {
	return PromisedAnswerOp_List(C.Struct(s).GetObject(0))
}
func (s PromisedAnswer) SetTransform(v PromisedAnswerOp_List) { C.Struct(s).SetObject(0, C.Object(v)) }

type PromisedAnswer_List C.PointerList

func NewPromisedAnswerList(s *C.Segment, sz int) PromisedAnswer_List {
	return PromisedAnswer_List(s.NewCompositeList(8, 1, sz))
}
func (s PromisedAnswer_List) Len() int { return C.PointerList(s).Len() }
func (s PromisedAnswer_List) At(i int) PromisedAnswer {
	return PromisedAnswer(C.PointerList(s).At(i).ToStruct())
}
func (s PromisedAnswer_List) ToArray() []PromisedAnswer {
	return *(*[]PromisedAnswer)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type PromisedAnswerOp C.Struct
type PromisedAnswerOp_which uint16

const (
	PROMISEDANSWEROP_NOOP            PromisedAnswerOp_which = 0
	PROMISEDANSWEROP_GETPOINTERFIELD                        = 1
)

func NewPromisedAnswerOp(s *C.Segment) PromisedAnswerOp { return PromisedAnswerOp(s.NewStruct(8, 0)) }
func ReadPromisedAnswerOp(s *C.Segment) PromisedAnswerOp {
	return PromisedAnswerOp(s.Root(0).ToStruct())
}
func (s PromisedAnswerOp) which() PromisedAnswerOp_which {
	return PromisedAnswerOp_which(C.Struct(s).Get16(0))
}
func (s PromisedAnswerOp) GetPointerField() uint16 { return C.Struct(s).Get16(2) }
func (s PromisedAnswerOp) SetGetPointerField(v uint16) {
	C.Struct(s).Set16(0, 1)
	C.Struct(s).Set16(2, v)
}

type PromisedAnswerOp_List C.PointerList

func NewPromisedAnswerOpList(s *C.Segment, sz int) PromisedAnswerOp_List {
	return PromisedAnswerOp_List(s.NewCompositeList(8, 0, sz))
}
func (s PromisedAnswerOp_List) Len() int { return C.PointerList(s).Len() }
func (s PromisedAnswerOp_List) At(i int) PromisedAnswerOp {
	return PromisedAnswerOp(C.PointerList(s).At(i).ToStruct())
}
func (s PromisedAnswerOp_List) ToArray() []PromisedAnswerOp {
	return *(*[]PromisedAnswerOp)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type ThirdPartyCapDescriptor C.Struct

func NewThirdPartyCapDescriptor(s *C.Segment) ThirdPartyCapDescriptor {
	return ThirdPartyCapDescriptor(s.NewStruct(8, 1))
}
func ReadThirdPartyCapDescriptor(s *C.Segment) ThirdPartyCapDescriptor {
	return ThirdPartyCapDescriptor(s.Root(0).ToStruct())
}
func (s ThirdPartyCapDescriptor) Id() C.Object       { return C.Struct(s).GetObject(0) }
func (s ThirdPartyCapDescriptor) SetId(v C.Object)   { C.Struct(s).SetObject(0, v) }
func (s ThirdPartyCapDescriptor) VineId() uint32     { return C.Struct(s).Get32(0) }
func (s ThirdPartyCapDescriptor) SetVineId(v uint32) { C.Struct(s).Set32(0, v) }

type ThirdPartyCapDescriptor_List C.PointerList

func NewThirdPartyCapDescriptorList(s *C.Segment, sz int) ThirdPartyCapDescriptor_List {
	return ThirdPartyCapDescriptor_List(s.NewCompositeList(8, 1, sz))
}
func (s ThirdPartyCapDescriptor_List) Len() int { return C.PointerList(s).Len() }
func (s ThirdPartyCapDescriptor_List) At(i int) ThirdPartyCapDescriptor {
	return ThirdPartyCapDescriptor(C.PointerList(s).At(i).ToStruct())
}
func (s ThirdPartyCapDescriptor_List) ToArray() []ThirdPartyCapDescriptor {
	return *(*[]ThirdPartyCapDescriptor)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type Exception C.Struct

func NewException(s *C.Segment) Exception    { return Exception(s.NewStruct(8, 1)) }
func ReadException(s *C.Segment) Exception   { return Exception(s.Root(0).ToStruct()) }
func (s Exception) Reason() string           { return C.Struct(s).GetObject(0).ToString() }
func (s Exception) SetReason(v string)       { C.Struct(s).SetObject(0, s.Segment.NewString(v)) }
func (s Exception) IsCallersFault() bool     { return C.Struct(s).Get1(0) }
func (s Exception) SetIsCallersFault(v bool) { C.Struct(s).Set1(0, v) }
func (s Exception) IsPermanent() bool        { return C.Struct(s).Get1(1) }
func (s Exception) SetIsPermanent(v bool)    { C.Struct(s).Set1(1, v) }
func (s Exception) IsOverloaded() bool       { return C.Struct(s).Get1(2) }
func (s Exception) SetIsOverloaded(v bool)   { C.Struct(s).Set1(2, v) }

type Exception_List C.PointerList

func NewExceptionList(s *C.Segment, sz int) Exception_List {
	return Exception_List(s.NewCompositeList(8, 1, sz))
}
func (s Exception_List) Len() int           { return C.PointerList(s).Len() }
func (s Exception_List) At(i int) Exception { return Exception(C.PointerList(s).At(i).ToStruct()) }
func (s Exception_List) ToArray() []Exception {
	return *(*[]Exception)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
