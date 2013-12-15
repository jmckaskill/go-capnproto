package capn

// AUTO GENERATED - DO NOT EDIT
import (
	C "github.com/jmckaskill/go-capnproto"
	"unsafe"
)

type SturdyRef C.Struct
type SturdyRef_which uint16

const (
	STURDYREF_EXTERNAL SturdyRef_which = 0
	STURDYREF_CONFINED                 = 1
)

func NewSturdyRef(s *C.Segment) SturdyRef  { return SturdyRef(s.NewStruct(8, 1)) }
func ReadSturdyRef(s *C.Segment) SturdyRef { return SturdyRef(s.Root(0).ToStruct()) }
func (s SturdyRef) which() SturdyRef_which { return SturdyRef_which(C.Struct(s).Get16(0)) }
func (s SturdyRef) External() C.Object     { return C.Struct(s).GetObject(0) }
func (s SturdyRef) SetExternal(v C.Object) { C.Struct(s).Set16(0, 0); C.Struct(s).SetObject(0, v) }
func (s SturdyRef) Confined() C.Object     { return C.Struct(s).GetObject(0) }
func (s SturdyRef) SetConfined(v C.Object) { C.Struct(s).Set16(0, 1); C.Struct(s).SetObject(0, v) }

type SturdyRef_List C.PointerList

func NewSturdyRefList(s *C.Segment, sz int) SturdyRef_List {
	return SturdyRef_List(s.NewCompositeList(8, 1, sz))
}
func (s SturdyRef_List) Len() int           { return C.PointerList(s).Len() }
func (s SturdyRef_List) At(i int) SturdyRef { return SturdyRef(C.PointerList(s).At(i).ToStruct()) }
func (s SturdyRef_List) ToArray() []SturdyRef {
	return *(*[]SturdyRef)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type ProvisionId C.Struct

func NewProvisionId(s *C.Segment) ProvisionId  { return ProvisionId(s.NewStruct(8, 0)) }
func ReadProvisionId(s *C.Segment) ProvisionId { return ProvisionId(s.Root(0).ToStruct()) }
func (s ProvisionId) JoinId() uint32           { return C.Struct(s).Get32(0) }
func (s ProvisionId) SetJoinId(v uint32)       { C.Struct(s).Set32(0, v) }

type ProvisionId_List C.PointerList

func NewProvisionIdList(s *C.Segment, sz int) ProvisionId_List {
	return ProvisionId_List(s.NewCompositeList(8, 0, sz))
}
func (s ProvisionId_List) Len() int             { return C.PointerList(s).Len() }
func (s ProvisionId_List) At(i int) ProvisionId { return ProvisionId(C.PointerList(s).At(i).ToStruct()) }
func (s ProvisionId_List) ToArray() []ProvisionId {
	return *(*[]ProvisionId)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type RecipientId C.Struct

func NewRecipientId(s *C.Segment) RecipientId  { return RecipientId(s.NewStruct(0, 0)) }
func ReadRecipientId(s *C.Segment) RecipientId { return RecipientId(s.Root(0).ToStruct()) }

type RecipientId_List C.PointerList

func NewRecipientIdList(s *C.Segment, sz int) RecipientId_List {
	return RecipientId_List(s.NewCompositeList(0, 0, sz))
}
func (s RecipientId_List) Len() int             { return C.PointerList(s).Len() }
func (s RecipientId_List) At(i int) RecipientId { return RecipientId(C.PointerList(s).At(i).ToStruct()) }
func (s RecipientId_List) ToArray() []RecipientId {
	return *(*[]RecipientId)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type ThirdPartyCapId C.Struct

func NewThirdPartyCapId(s *C.Segment) ThirdPartyCapId  { return ThirdPartyCapId(s.NewStruct(0, 0)) }
func ReadThirdPartyCapId(s *C.Segment) ThirdPartyCapId { return ThirdPartyCapId(s.Root(0).ToStruct()) }

type ThirdPartyCapId_List C.PointerList

func NewThirdPartyCapIdList(s *C.Segment, sz int) ThirdPartyCapId_List {
	return ThirdPartyCapId_List(s.NewCompositeList(0, 0, sz))
}
func (s ThirdPartyCapId_List) Len() int { return C.PointerList(s).Len() }
func (s ThirdPartyCapId_List) At(i int) ThirdPartyCapId {
	return ThirdPartyCapId(C.PointerList(s).At(i).ToStruct())
}
func (s ThirdPartyCapId_List) ToArray() []ThirdPartyCapId {
	return *(*[]ThirdPartyCapId)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type JoinKeyPart C.Struct

func NewJoinKeyPart(s *C.Segment) JoinKeyPart  { return JoinKeyPart(s.NewStruct(8, 0)) }
func ReadJoinKeyPart(s *C.Segment) JoinKeyPart { return JoinKeyPart(s.Root(0).ToStruct()) }
func (s JoinKeyPart) JoinId() uint32           { return C.Struct(s).Get32(0) }
func (s JoinKeyPart) SetJoinId(v uint32)       { C.Struct(s).Set32(0, v) }
func (s JoinKeyPart) PartCount() uint16        { return C.Struct(s).Get16(4) }
func (s JoinKeyPart) SetPartCount(v uint16)    { C.Struct(s).Set16(4, v) }
func (s JoinKeyPart) PartNum() uint16          { return C.Struct(s).Get16(6) }
func (s JoinKeyPart) SetPartNum(v uint16)      { C.Struct(s).Set16(6, v) }

type JoinKeyPart_List C.PointerList

func NewJoinKeyPartList(s *C.Segment, sz int) JoinKeyPart_List {
	return JoinKeyPart_List(s.NewCompositeList(8, 0, sz))
}
func (s JoinKeyPart_List) Len() int             { return C.PointerList(s).Len() }
func (s JoinKeyPart_List) At(i int) JoinKeyPart { return JoinKeyPart(C.PointerList(s).At(i).ToStruct()) }
func (s JoinKeyPart_List) ToArray() []JoinKeyPart {
	return *(*[]JoinKeyPart)(unsafe.Pointer(C.PointerList(s).ToArray()))
}

type JoinAnswer C.Struct

func NewJoinAnswer(s *C.Segment) JoinAnswer  { return JoinAnswer(s.NewStruct(8, 0)) }
func ReadJoinAnswer(s *C.Segment) JoinAnswer { return JoinAnswer(s.Root(0).ToStruct()) }
func (s JoinAnswer) JoinId() uint32          { return C.Struct(s).Get32(0) }
func (s JoinAnswer) SetJoinId(v uint32)      { C.Struct(s).Set32(0, v) }
func (s JoinAnswer) Succeeded() bool         { return C.Struct(s).Get1(32) }
func (s JoinAnswer) SetSucceeded(v bool)     { C.Struct(s).Set1(32, v) }

type JoinAnswer_List C.PointerList

func NewJoinAnswerList(s *C.Segment, sz int) JoinAnswer_List {
	return JoinAnswer_List(s.NewCompositeList(8, 0, sz))
}
func (s JoinAnswer_List) Len() int            { return C.PointerList(s).Len() }
func (s JoinAnswer_List) At(i int) JoinAnswer { return JoinAnswer(C.PointerList(s).At(i).ToStruct()) }
func (s JoinAnswer_List) ToArray() []JoinAnswer {
	return *(*[]JoinAnswer)(unsafe.Pointer(C.PointerList(s).ToArray()))
}
