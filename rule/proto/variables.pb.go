// Code generated by protoc-gen-go. DO NOT EDIT.
// source: variables.proto

package parser_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KeyExclusion_Type int32

const (
	KeyExclusion_STRING KeyExclusion_Type = 0
	KeyExclusion_REGEX  KeyExclusion_Type = 1
)

var KeyExclusion_Type_name = map[int32]string{
	0: "STRING",
	1: "REGEX",
}

var KeyExclusion_Type_value = map[string]int32{
	"STRING": 0,
	"REGEX":  1,
}

func (x KeyExclusion_Type) String() string {
	return proto.EnumName(KeyExclusion_Type_name, int32(x))
}

func (KeyExclusion_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b8b958d8129f2ed, []int{1, 0}
}

type Variable struct {
	// Variable's collection name.
	CollectionName string `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	// Varibale's name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates whether use "&" in variables.
	IsCount bool `protobuf:"varint,3,opt,name=is_count,json=isCount,proto3" json:"is_count,omitempty"`
	// Collection of variables that use "!".
	KeyExclusion         []*KeyExclusion `protobuf:"bytes,4,rep,name=key_exclusion,json=keyExclusion,proto3" json:"key_exclusion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Variable) Reset()         { *m = Variable{} }
func (m *Variable) String() string { return proto.CompactTextString(m) }
func (*Variable) ProtoMessage()    {}
func (*Variable) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b8b958d8129f2ed, []int{0}
}

func (m *Variable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Variable.Unmarshal(m, b)
}
func (m *Variable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Variable.Marshal(b, m, deterministic)
}
func (m *Variable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Variable.Merge(m, src)
}
func (m *Variable) XXX_Size() int {
	return xxx_messageInfo_Variable.Size(m)
}
func (m *Variable) XXX_DiscardUnknown() {
	xxx_messageInfo_Variable.DiscardUnknown(m)
}

var xxx_messageInfo_Variable proto.InternalMessageInfo

func (m *Variable) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *Variable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Variable) GetIsCount() bool {
	if m != nil {
		return m.IsCount
	}
	return false
}

func (m *Variable) GetKeyExclusion() []*KeyExclusion {
	if m != nil {
		return m.KeyExclusion
	}
	return nil
}

type KeyExclusion struct {
	Type                 KeyExclusion_Type `protobuf:"varint,1,opt,name=type,proto3,enum=parser.proto.KeyExclusion_Type" json:"type,omitempty"`
	Param                string            `protobuf:"bytes,2,opt,name=param,proto3" json:"param,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *KeyExclusion) Reset()         { *m = KeyExclusion{} }
func (m *KeyExclusion) String() string { return proto.CompactTextString(m) }
func (*KeyExclusion) ProtoMessage()    {}
func (*KeyExclusion) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b8b958d8129f2ed, []int{1}
}

func (m *KeyExclusion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyExclusion.Unmarshal(m, b)
}
func (m *KeyExclusion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyExclusion.Marshal(b, m, deterministic)
}
func (m *KeyExclusion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyExclusion.Merge(m, src)
}
func (m *KeyExclusion) XXX_Size() int {
	return xxx_messageInfo_KeyExclusion.Size(m)
}
func (m *KeyExclusion) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyExclusion.DiscardUnknown(m)
}

var xxx_messageInfo_KeyExclusion proto.InternalMessageInfo

func (m *KeyExclusion) GetType() KeyExclusion_Type {
	if m != nil {
		return m.Type
	}
	return KeyExclusion_STRING
}

func (m *KeyExclusion) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

type VariableList struct {
	Item                 []*Variable `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VariableList) Reset()         { *m = VariableList{} }
func (m *VariableList) String() string { return proto.CompactTextString(m) }
func (*VariableList) ProtoMessage()    {}
func (*VariableList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b8b958d8129f2ed, []int{2}
}

func (m *VariableList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VariableList.Unmarshal(m, b)
}
func (m *VariableList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VariableList.Marshal(b, m, deterministic)
}
func (m *VariableList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VariableList.Merge(m, src)
}
func (m *VariableList) XXX_Size() int {
	return xxx_messageInfo_VariableList.Size(m)
}
func (m *VariableList) XXX_DiscardUnknown() {
	xxx_messageInfo_VariableList.DiscardUnknown(m)
}

var xxx_messageInfo_VariableList proto.InternalMessageInfo

func (m *VariableList) GetItem() []*Variable {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterEnum("parser.proto.KeyExclusion_Type", KeyExclusion_Type_name, KeyExclusion_Type_value)
	proto.RegisterType((*Variable)(nil), "parser.proto.Variable")
	proto.RegisterType((*KeyExclusion)(nil), "parser.proto.KeyExclusion")
	proto.RegisterType((*VariableList)(nil), "parser.proto.VariableList")
}

func init() { proto.RegisterFile("variables.proto", fileDescriptor_3b8b958d8129f2ed) }

var fileDescriptor_3b8b958d8129f2ed = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x8d, 0xeb, 0x66, 0xf7, 0xac, 0xdb, 0x78, 0x88, 0x54, 0x41, 0x2c, 0xbd, 0x58, 0x3c,
	0xf4, 0xb0, 0xdd, 0xbc, 0x78, 0x90, 0x32, 0x44, 0xd9, 0x21, 0x0e, 0xf1, 0x56, 0xb2, 0xf2, 0x0e,
	0x61, 0x6d, 0x53, 0x9a, 0x4c, 0xd6, 0xcf, 0xe3, 0x17, 0x95, 0xa5, 0x2b, 0xd6, 0x83, 0xb7, 0xfc,
	0xde, 0xff, 0x47, 0xf2, 0xcf, 0x83, 0xe9, 0x97, 0xa8, 0xa5, 0xd8, 0xe4, 0xa4, 0xe3, 0xaa, 0x56,
	0x46, 0xa1, 0x57, 0x89, 0x5a, 0x53, 0xdd, 0x52, 0xf8, 0xcd, 0xc0, 0xfd, 0x38, 0x1a, 0x78, 0x0f,
	0xd3, 0x4c, 0xe5, 0x39, 0x65, 0x46, 0xaa, 0x32, 0x2d, 0x45, 0x41, 0x3e, 0x0b, 0x58, 0x34, 0xe6,
	0x93, 0xdf, 0xf1, 0x4a, 0x14, 0x84, 0x08, 0x8e, 0x4d, 0x4f, 0x6d, 0x6a, 0xcf, 0x78, 0x0d, 0xae,
	0xd4, 0x69, 0xa6, 0x76, 0xa5, 0xf1, 0x07, 0x01, 0x8b, 0x5c, 0x7e, 0x26, 0xf5, 0xf3, 0x01, 0xf1,
	0x09, 0x2e, 0xb6, 0xd4, 0xa4, 0xb4, 0xcf, 0xf2, 0x9d, 0x96, 0xaa, 0xf4, 0x9d, 0x60, 0x10, 0x9d,
	0xcf, 0x6f, 0xe2, 0x7e, 0x95, 0xf8, 0x95, 0x9a, 0xa4, 0x33, 0xb8, 0xb7, 0xed, 0x51, 0xb8, 0x07,
	0xaf, 0x9f, 0xe2, 0x02, 0x1c, 0xd3, 0x54, 0x6d, 0xbb, 0xc9, 0xfc, 0xee, 0xff, 0x7b, 0xe2, 0x75,
	0x53, 0x11, 0xb7, 0x32, 0x5e, 0xc2, 0xb0, 0x12, 0xb5, 0x28, 0x8e, 0xad, 0x5b, 0x08, 0x6f, 0xc1,
	0x39, 0x38, 0x08, 0x30, 0x7a, 0x5f, 0xf3, 0x97, 0xd5, 0x72, 0x76, 0x82, 0x63, 0x18, 0xf2, 0x64,
	0x99, 0x7c, 0xce, 0x58, 0xf8, 0x08, 0x5e, 0xb7, 0x9e, 0x37, 0xa9, 0x0d, 0x3e, 0x80, 0x23, 0x0d,
	0x15, 0x3e, 0xb3, 0x3f, 0xb8, 0xfa, 0xfb, 0x72, 0x67, 0x72, 0xeb, 0x6c, 0x46, 0x76, 0xba, 0xf8,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xd9, 0x8e, 0x9a, 0x83, 0x01, 0x00, 0x00,
}
