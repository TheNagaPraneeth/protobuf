// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imp/imp.proto

package imp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ImportedMessage_Owner int32

const (
	ImportedMessage_DAVE ImportedMessage_Owner = 1
	ImportedMessage_MIKE ImportedMessage_Owner = 2
)

var ImportedMessage_Owner_name = map[int32]string{
	1: "DAVE",
	2: "MIKE",
}
var ImportedMessage_Owner_value = map[string]int32{
	"DAVE": 1,
	"MIKE": 2,
}

func (x ImportedMessage_Owner) Enum() *ImportedMessage_Owner {
	p := new(ImportedMessage_Owner)
	*p = x
	return p
}
func (x ImportedMessage_Owner) String() string {
	return proto.EnumName(ImportedMessage_Owner_name, int32(x))
}
func (x *ImportedMessage_Owner) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImportedMessage_Owner_value, data, "ImportedMessage_Owner")
	if err != nil {
		return err
	}
	*x = ImportedMessage_Owner(value)
	return nil
}
func (ImportedMessage_Owner) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type ImportedMessage struct {
	Field *int64 `protobuf:"varint,1,req,name=field" json:"field,omitempty"`
	// The forwarded getters for these fields are fiddly to get right.
	LocalMsg   *ImportedMessage2       `protobuf:"bytes,2,opt,name=local_msg,json=localMsg" json:"local_msg,omitempty"`
	ForeignMsg *ForeignImportedMessage `protobuf:"bytes,3,opt,name=foreign_msg,json=foreignMsg" json:"foreign_msg,omitempty"`
	EnumField  *ImportedMessage_Owner  `protobuf:"varint,4,opt,name=enum_field,json=enumField,enum=imp.ImportedMessage_Owner" json:"enum_field,omitempty"`
	// Types that are valid to be assigned to Union:
	//	*ImportedMessage_State
	Union                        isImportedMessage_Union      `protobuf_oneof:"union"`
	Name                         []string                     `protobuf:"bytes,5,rep,name=name" json:"name,omitempty"`
	Boss                         []ImportedMessage_Owner      `protobuf:"varint,6,rep,name=boss,enum=imp.ImportedMessage_Owner" json:"boss,omitempty"`
	Memo                         []*ImportedMessage2          `protobuf:"bytes,7,rep,name=memo" json:"memo,omitempty"`
	MsgMap                       map[string]*ImportedMessage2 `protobuf:"bytes,8,rep,name=msg_map,json=msgMap" json:"msg_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral         struct{}                     `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *ImportedMessage) Reset()                    { *m = ImportedMessage{} }
func (m *ImportedMessage) String() string            { return proto.CompactTextString(m) }
func (*ImportedMessage) ProtoMessage()               {}
func (*ImportedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

var extRange_ImportedMessage = []proto.ExtensionRange{
	{90, 100},
}

func (*ImportedMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ImportedMessage
}
func (m *ImportedMessage) Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportedMessage.Unmarshal(m, b)
}
func (m *ImportedMessage) Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportedMessage.Marshal(b, m, deterministic)
}
func (dst *ImportedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportedMessage.Merge(dst, src)
}
func (m *ImportedMessage) XXX_Size() int {
	return xxx_messageInfo_ImportedMessage.Size(m)
}
func (m *ImportedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ImportedMessage proto.InternalMessageInfo

type isImportedMessage_Union interface {
	isImportedMessage_Union()
}

type ImportedMessage_State struct {
	State int32 `protobuf:"varint,9,opt,name=state,oneof"`
}

func (*ImportedMessage_State) isImportedMessage_Union() {}

func (m *ImportedMessage) GetUnion() isImportedMessage_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *ImportedMessage) GetField() int64 {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return 0
}

func (m *ImportedMessage) GetLocalMsg() *ImportedMessage2 {
	if m != nil {
		return m.LocalMsg
	}
	return nil
}

func (m *ImportedMessage) GetForeignMsg() *ForeignImportedMessage {
	if m != nil {
		return m.ForeignMsg
	}
	return nil
}

func (m *ImportedMessage) GetEnumField() ImportedMessage_Owner {
	if m != nil && m.EnumField != nil {
		return *m.EnumField
	}
	return ImportedMessage_DAVE
}

func (m *ImportedMessage) GetState() int32 {
	if x, ok := m.GetUnion().(*ImportedMessage_State); ok {
		return x.State
	}
	return 0
}

func (m *ImportedMessage) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ImportedMessage) GetBoss() []ImportedMessage_Owner {
	if m != nil {
		return m.Boss
	}
	return nil
}

func (m *ImportedMessage) GetMemo() []*ImportedMessage2 {
	if m != nil {
		return m.Memo
	}
	return nil
}

func (m *ImportedMessage) GetMsgMap() map[string]*ImportedMessage2 {
	if m != nil {
		return m.MsgMap
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ImportedMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ImportedMessage_OneofMarshaler, _ImportedMessage_OneofUnmarshaler, _ImportedMessage_OneofSizer, []interface{}{
		(*ImportedMessage_State)(nil),
	}
}

func _ImportedMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ImportedMessage)
	// union
	switch x := m.Union.(type) {
	case *ImportedMessage_State:
		b.EncodeVarint(9<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.State))
	case nil:
	default:
		return fmt.Errorf("ImportedMessage.Union has unexpected type %T", x)
	}
	return nil
}

func _ImportedMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ImportedMessage)
	switch tag {
	case 9: // union.state
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Union = &ImportedMessage_State{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _ImportedMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ImportedMessage)
	// union
	switch x := m.Union.(type) {
	case *ImportedMessage_State:
		n += proto.SizeVarint(9<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.State))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ImportedMessage2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImportedMessage2) Reset()                    { *m = ImportedMessage2{} }
func (m *ImportedMessage2) String() string            { return proto.CompactTextString(m) }
func (*ImportedMessage2) ProtoMessage()               {}
func (*ImportedMessage2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }
func (m *ImportedMessage2) Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportedMessage2.Unmarshal(m, b)
}
func (m *ImportedMessage2) Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportedMessage2.Marshal(b, m, deterministic)
}
func (dst *ImportedMessage2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportedMessage2.Merge(dst, src)
}
func (m *ImportedMessage2) XXX_Size() int {
	return xxx_messageInfo_ImportedMessage2.Size(m)
}
func (m *ImportedMessage2) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportedMessage2.DiscardUnknown(m)
}

var xxx_messageInfo_ImportedMessage2 proto.InternalMessageInfo

type ImportedExtendable struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `protobuf_messageset:"1" json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *ImportedExtendable) Reset()                    { *m = ImportedExtendable{} }
func (m *ImportedExtendable) String() string            { return proto.CompactTextString(m) }
func (*ImportedExtendable) ProtoMessage()               {}
func (*ImportedExtendable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ImportedExtendable) MarshalJSON() ([]byte, error) {
	return proto.MarshalMessageSetJSON(&m.XXX_InternalExtensions)
}
func (m *ImportedExtendable) UnmarshalJSON(buf []byte) error {
	return proto.UnmarshalMessageSetJSON(buf, &m.XXX_InternalExtensions)
}

// ensure ImportedExtendable satisfies proto.Unmarshaler
var _ proto.Unmarshaler = (*ImportedExtendable)(nil)

var extRange_ImportedExtendable = []proto.ExtensionRange{
	{100, 2147483646},
}

func (*ImportedExtendable) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ImportedExtendable
}
func (m *ImportedExtendable) Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportedExtendable.Unmarshal(m, b)
}
func (m *ImportedExtendable) Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportedExtendable.Marshal(b, m, deterministic)
}
func (dst *ImportedExtendable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportedExtendable.Merge(dst, src)
}
func (m *ImportedExtendable) XXX_Size() int {
	return xxx_messageInfo_ImportedExtendable.Size(m)
}
func (m *ImportedExtendable) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportedExtendable.DiscardUnknown(m)
}

var xxx_messageInfo_ImportedExtendable proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ImportedMessage)(nil), "imp.ImportedMessage")
	proto.RegisterMapType((map[string]*ImportedMessage2)(nil), "imp.ImportedMessage.MsgMapEntry")
	proto.RegisterType((*ImportedMessage2)(nil), "imp.ImportedMessage2")
	proto.RegisterType((*ImportedExtendable)(nil), "imp.ImportedExtendable")
	proto.RegisterEnum("imp.ImportedMessage_Owner", ImportedMessage_Owner_name, ImportedMessage_Owner_value)
}

func init() { proto.RegisterFile("imp/imp.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x4d, 0xff, 0xec, 0xb4, 0xef, 0xe0, 0x5a, 0x82, 0x4a, 0x99, 0xbd, 0x84, 0x9e, 0xea,
	0xca, 0x76, 0xa0, 0x22, 0xba, 0x8b, 0x17, 0x17, 0x67, 0x71, 0x91, 0xa2, 0xf4, 0xe0, 0x61, 0x2f,
	0x43, 0x66, 0x9a, 0x89, 0xc5, 0x26, 0x29, 0x4d, 0xaa, 0xee, 0xf7, 0xf0, 0xfb, 0x56, 0x9a, 0xae,
	0x22, 0xc3, 0xe8, 0xde, 0x9e, 0xe7, 0xe1, 0xf7, 0xe4, 0x4d, 0x9b, 0x17, 0x1e, 0xd6, 0xa2, 0x5d,
	0xd6, 0xa2, 0xcd, 0xda, 0x4e, 0x19, 0x85, 0xdd, 0x5a, 0xb4, 0x8b, 0xe3, 0xbb, 0x2c, 0x9f, 0xc2,
	0x3f, 0xfe, 0xc5, 0xe4, 0x93, 0x9f, 0x1e, 0x3c, 0xba, 0x16, 0xad, 0xea, 0x0c, 0xab, 0x0a, 0xa6,
	0x35, 0xe5, 0x0c, 0x3f, 0x06, 0x7f, 0x57, 0xb3, 0xa6, 0x8a, 0x11, 0x71, 0x52, 0xb7, 0x9c, 0x0c,
	0xce, 0x21, 0x6c, 0xd4, 0x96, 0x36, 0x6b, 0xa1, 0x79, 0xec, 0x10, 0x94, 0xce, 0xf3, 0x27, 0xd9,
	0x38, 0x6d, 0xaf, 0x9e, 0x97, 0x81, 0xe5, 0x0a, 0xcd, 0xf1, 0x1b, 0x98, 0xef, 0x54, 0xc7, 0x6a,
	0x2e, 0x6d, 0xcb, 0xb5, 0xad, 0x13, 0xdb, 0xba, 0x9a, 0xf2, 0xbd, 0x72, 0x09, 0x77, 0xfc, 0xd8,
	0x3e, 0x07, 0x60, 0xb2, 0x17, 0xeb, 0xe9, 0x32, 0x1e, 0x41, 0xe9, 0x71, 0xbe, 0x38, 0x34, 0x32,
	0xfb, 0xf8, 0x5d, 0xb2, 0xae, 0x0c, 0x47, 0xfa, 0xca, 0x5e, 0xf6, 0x29, 0xf8, 0xda, 0x50, 0xc3,
	0xe2, 0x90, 0xa0, 0xd4, 0x7f, 0xff, 0xa0, 0x9c, 0x2c, 0xc6, 0xe0, 0x49, 0x2a, 0x58, 0xec, 0x13,
	0x37, 0x0d, 0x4b, 0xab, 0x71, 0x06, 0xde, 0x46, 0x69, 0x1d, 0x1f, 0x11, 0xf7, 0x9e, 0x01, 0x96,
	0xc3, 0xcf, 0xc0, 0x13, 0x4c, 0xa8, 0x78, 0x46, 0xdc, 0x7f, 0xff, 0x03, 0x8b, 0xe0, 0x73, 0x98,
	0x09, 0xcd, 0xd7, 0x82, 0xb6, 0x71, 0x60, 0x69, 0x72, 0xf0, 0xf4, 0x42, 0xf3, 0x82, 0xb6, 0x2b,
	0x69, 0xba, 0xdb, 0xf2, 0x48, 0x58, 0xb3, 0xf8, 0x04, 0xf3, 0xbf, 0x62, 0x1c, 0x81, 0xfb, 0x95,
	0xdd, 0xc6, 0x88, 0xa0, 0x34, 0x2c, 0x47, 0x89, 0x9f, 0x83, 0xff, 0x8d, 0x36, 0x3d, 0xfb, 0xff,
	0x5b, 0x4c, 0xcc, 0x85, 0xf3, 0x1a, 0x25, 0x27, 0xe0, 0xdb, 0xcf, 0xc0, 0x01, 0x78, 0xef, 0xde,
	0x7e, 0x5e, 0x45, 0x68, 0x54, 0xc5, 0xf5, 0x87, 0x55, 0xe4, 0x9c, 0x7a, 0xc1, 0x4d, 0xc4, 0x2e,
	0x67, 0xe0, 0xf7, 0xb2, 0x56, 0x32, 0xc1, 0x10, 0xed, 0x1f, 0x95, 0x24, 0x80, 0x7f, 0x67, 0xab,
	0x1f, 0x86, 0xc9, 0x8a, 0x6e, 0x1a, 0x76, 0x1a, 0x04, 0x55, 0x34, 0x0c, 0xc3, 0x30, 0xbb, 0x70,
	0x02, 0x74, 0xf9, 0xea, 0xe6, 0x25, 0xaf, 0xcd, 0x97, 0x7e, 0x93, 0x6d, 0x95, 0x58, 0x72, 0xd5,
	0x50, 0xc9, 0x97, 0x76, 0xd3, 0x36, 0xfd, 0x6e, 0x12, 0xdb, 0x33, 0xce, 0xe4, 0x19, 0x57, 0x4b,
	0xc3, 0xb4, 0xa9, 0xa8, 0xa1, 0xe3, 0x3a, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xe4, 0x5b,
	0xa3, 0xbc, 0x02, 0x00, 0x00,
}
