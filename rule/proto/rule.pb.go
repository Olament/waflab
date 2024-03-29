// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rule.proto

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

type Rule struct {
	// Rule's attributes
	// Reference: https://github.com/SpiderLabs/ModSecurity/wiki/Reference-Manual-%28v2.x%29
	Maturity int32  `protobuf:"varint,1,opt,name=maturity,proto3" json:"maturity,omitempty"`
	Phase    int32  `protobuf:"varint,2,opt,name=phase,proto3" json:"phase,omitempty"`
	Rev      string `protobuf:"bytes,3,opt,name=rev,proto3" json:"rev,omitempty"`
	Id       int64  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Accuracy int32  `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	Ver      string `protobuf:"bytes,6,opt,name=ver,proto3" json:"ver,omitempty"`
	// Information of Marker.
	// Reference: https://github.com/SpiderLabs/ModSecurity/wiki/Reference-Manual-%28v2.x%29#secmarker
	Marker    string `protobuf:"bytes,7,opt,name=marker,proto3" json:"marker,omitempty"`
	SecMarker bool   `protobuf:"varint,8,opt,name=sec_marker,json=secMarker,proto3" json:"sec_marker,omitempty"`
	// Reference: https://github.com/SpiderLabs/ModSecurity/wiki/Reference-Manual-%28v2.x%29#severity
	Severity int32 `protobuf:"varint,9,opt,name=severity,proto3" json:"severity,omitempty"`
	// Reference: https://github.com/SpiderLabs/ModSecurity/wiki/Reference-Manual-%28v2.x%29#chain
	Chained          bool  `protobuf:"varint,10,opt,name=chained,proto3" json:"chained,omitempty"`
	ChainedRuleChild *Rule `protobuf:"bytes,11,opt,name=chained_rule_child,json=chainedRuleChild,proto3" json:"chained_rule_child,omitempty"`
	// Debug log file name.
	FileName string `protobuf:"bytes,12,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	// Line number of specific debug log.
	LineNumber int32 `protobuf:"varint,13,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	// Block Actions.
	ActionsRuntimePos []*Action `protobuf:"bytes,14,rep,name=actions_runtime_pos,json=actionsRuntimePos,proto3" json:"actions_runtime_pos,omitempty"`
	// Actions belongs to runtime before match attempt kind.
	ActionsRuntimePre []*Action `protobuf:"bytes,15,rep,name=actions_runtime_pre,json=actionsRuntimePre,proto3" json:"actions_runtime_pre,omitempty"`
	// Rule's operation.
	Op *Operator `protobuf:"bytes,16,opt,name=op,proto3" json:"op,omitempty"`
	// Indicate whether this rule does not contain operation.
	Unconditional bool        `protobuf:"varint,17,opt,name=unconditional,proto3" json:"unconditional,omitempty"`
	Variables     []*Variable `protobuf:"bytes,18,rep,name=variables,proto3" json:"variables,omitempty"`
	Setvar        []*SetVar   `protobuf:"bytes,19,rep,name=setvar,proto3" json:"setvar,omitempty"`
	// More information about action's type:
	// https://github.com/SpiderLabs/ModSecurity/wiki/Reference-Manual-%28v2.x%29#Actions
	ContiansCaptureAction     bool     `protobuf:"varint,20,opt,name=contians_capture_action,json=contiansCaptureAction,proto3" json:"contians_capture_action,omitempty"`
	ContiansMultimatchAction  bool     `protobuf:"varint,21,opt,name=contians_multimatch_action,json=contiansMultimatchAction,proto3" json:"contians_multimatch_action,omitempty"`
	ContiansStaticBlockAction bool     `protobuf:"varint,22,opt,name=contians_static_block_action,json=contiansStaticBlockAction,proto3" json:"contians_static_block_action,omitempty"`
	DisruptiveAction          *Action  `protobuf:"bytes,23,opt,name=disruptive_action,json=disruptiveAction,proto3" json:"disruptive_action,omitempty"`
	Logdata                   *LogData `protobuf:"bytes,24,opt,name=logdata,proto3" json:"logdata,omitempty"`
	Msg                       *Msg     `protobuf:"bytes,25,opt,name=msg,proto3" json:"msg,omitempty"`
	Tag                       []*Tag   `protobuf:"bytes,26,rep,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_07e8e0fa338d4596, []int{0}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetMaturity() int32 {
	if m != nil {
		return m.Maturity
	}
	return 0
}

func (m *Rule) GetPhase() int32 {
	if m != nil {
		return m.Phase
	}
	return 0
}

func (m *Rule) GetRev() string {
	if m != nil {
		return m.Rev
	}
	return ""
}

func (m *Rule) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Rule) GetAccuracy() int32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func (m *Rule) GetVer() string {
	if m != nil {
		return m.Ver
	}
	return ""
}

func (m *Rule) GetMarker() string {
	if m != nil {
		return m.Marker
	}
	return ""
}

func (m *Rule) GetSecMarker() bool {
	if m != nil {
		return m.SecMarker
	}
	return false
}

func (m *Rule) GetSeverity() int32 {
	if m != nil {
		return m.Severity
	}
	return 0
}

func (m *Rule) GetChained() bool {
	if m != nil {
		return m.Chained
	}
	return false
}

func (m *Rule) GetChainedRuleChild() *Rule {
	if m != nil {
		return m.ChainedRuleChild
	}
	return nil
}

func (m *Rule) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Rule) GetLineNumber() int32 {
	if m != nil {
		return m.LineNumber
	}
	return 0
}

func (m *Rule) GetActionsRuntimePos() []*Action {
	if m != nil {
		return m.ActionsRuntimePos
	}
	return nil
}

func (m *Rule) GetActionsRuntimePre() []*Action {
	if m != nil {
		return m.ActionsRuntimePre
	}
	return nil
}

func (m *Rule) GetOp() *Operator {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *Rule) GetUnconditional() bool {
	if m != nil {
		return m.Unconditional
	}
	return false
}

func (m *Rule) GetVariables() []*Variable {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *Rule) GetSetvar() []*SetVar {
	if m != nil {
		return m.Setvar
	}
	return nil
}

func (m *Rule) GetContiansCaptureAction() bool {
	if m != nil {
		return m.ContiansCaptureAction
	}
	return false
}

func (m *Rule) GetContiansMultimatchAction() bool {
	if m != nil {
		return m.ContiansMultimatchAction
	}
	return false
}

func (m *Rule) GetContiansStaticBlockAction() bool {
	if m != nil {
		return m.ContiansStaticBlockAction
	}
	return false
}

func (m *Rule) GetDisruptiveAction() *Action {
	if m != nil {
		return m.DisruptiveAction
	}
	return nil
}

func (m *Rule) GetLogdata() *LogData {
	if m != nil {
		return m.Logdata
	}
	return nil
}

func (m *Rule) GetMsg() *Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *Rule) GetTag() []*Tag {
	if m != nil {
		return m.Tag
	}
	return nil
}

type RuleList struct {
	Item                 []*Rule  `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuleList) Reset()         { *m = RuleList{} }
func (m *RuleList) String() string { return proto.CompactTextString(m) }
func (*RuleList) ProtoMessage()    {}
func (*RuleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_07e8e0fa338d4596, []int{1}
}

func (m *RuleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuleList.Unmarshal(m, b)
}
func (m *RuleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuleList.Marshal(b, m, deterministic)
}
func (m *RuleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuleList.Merge(m, src)
}
func (m *RuleList) XXX_Size() int {
	return xxx_messageInfo_RuleList.Size(m)
}
func (m *RuleList) XXX_DiscardUnknown() {
	xxx_messageInfo_RuleList.DiscardUnknown(m)
}

var xxx_messageInfo_RuleList proto.InternalMessageInfo

func (m *RuleList) GetItem() []*Rule {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterType((*Rule)(nil), "parser.proto.Rule")
	proto.RegisterType((*RuleList)(nil), "parser.proto.RuleList")
}

func init() { proto.RegisterFile("rule.proto", fileDescriptor_07e8e0fa338d4596) }

var fileDescriptor_07e8e0fa338d4596 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x5d, 0x4f, 0xdb, 0x3e,
	0x14, 0xc6, 0x95, 0x16, 0x4a, 0x7a, 0xca, 0x4b, 0x6b, 0xde, 0x4c, 0xff, 0xff, 0x69, 0x15, 0x9b,
	0x50, 0x2f, 0x26, 0x26, 0xb1, 0x69, 0x57, 0x93, 0x36, 0x06, 0x97, 0xc0, 0x26, 0x33, 0x71, 0x1b,
	0xb9, 0xce, 0x59, 0x6b, 0x91, 0xc4, 0x91, 0xed, 0x44, 0xe2, 0x1b, 0xed, 0x63, 0x4e, 0x76, 0x9c,
	0xb2, 0xa2, 0xde, 0xec, 0x2e, 0xcf, 0x79, 0x7e, 0xcf, 0xe9, 0xf1, 0xb1, 0x0b, 0xa0, 0xab, 0x0c,
	0xcf, 0x4b, 0xad, 0xac, 0x22, 0xdb, 0x25, 0xd7, 0x06, 0x75, 0xa3, 0xc6, 0x7b, 0x35, 0xd7, 0x92,
	0xcf, 0x32, 0x34, 0xa1, 0xb0, 0xc3, 0x85, 0x95, 0xaa, 0x68, 0xe5, 0x9e, 0x2a, 0x51, 0x73, 0xab,
	0x74, 0x28, 0x9c, 0xfe, 0x8e, 0x61, 0x83, 0x55, 0x19, 0x92, 0x31, 0xc4, 0x39, 0xb7, 0x95, 0x96,
	0xf6, 0x89, 0x46, 0x93, 0x68, 0xba, 0xc9, 0x96, 0x9a, 0x1c, 0xc0, 0x66, 0xb9, 0xe0, 0x06, 0x69,
	0xc7, 0x1b, 0x8d, 0x20, 0x43, 0xe8, 0x6a, 0xac, 0x69, 0x77, 0x12, 0x4d, 0xfb, 0xcc, 0x7d, 0x92,
	0x5d, 0xe8, 0xc8, 0x94, 0x6e, 0x4c, 0xa2, 0x69, 0x97, 0x75, 0x64, 0xea, 0x7a, 0x72, 0x21, 0x2a,
	0xcd, 0xc5, 0x13, 0xdd, 0x6c, 0x7a, 0xb6, 0xda, 0xa5, 0x6b, 0xd4, 0xb4, 0xd7, 0xa4, 0x6b, 0xd4,
	0xe4, 0x08, 0x7a, 0x39, 0xd7, 0x8f, 0xa8, 0xe9, 0x96, 0x2f, 0x06, 0x45, 0x5e, 0x01, 0x18, 0x14,
	0x49, 0xf0, 0xe2, 0x49, 0x34, 0x8d, 0x59, 0xdf, 0xa0, 0xb8, 0x6d, 0xec, 0x31, 0xc4, 0x06, 0x6b,
	0xf4, 0x83, 0xf7, 0x9b, 0x1f, 0x69, 0x35, 0xa1, 0xb0, 0x25, 0x16, 0x5c, 0x16, 0x98, 0x52, 0xf0,
	0xb9, 0x56, 0x92, 0xaf, 0x40, 0xc2, 0x67, 0xe2, 0x96, 0x99, 0x88, 0x85, 0xcc, 0x52, 0x3a, 0x98,
	0x44, 0xd3, 0xc1, 0x05, 0x39, 0xff, 0x7b, 0xa7, 0xe7, 0x6e, 0x3d, 0x6c, 0x18, 0x68, 0x27, 0xae,
	0x1c, 0x4b, 0xfe, 0x83, 0xfe, 0x2f, 0x99, 0x61, 0x52, 0xf0, 0x1c, 0xe9, 0xb6, 0x9f, 0x38, 0x76,
	0x85, 0x3b, 0x9e, 0x23, 0x79, 0x0d, 0x83, 0x4c, 0x16, 0x98, 0x14, 0x55, 0x3e, 0x43, 0x4d, 0x77,
	0xfc, 0x5c, 0xe0, 0x4a, 0x77, 0xbe, 0x42, 0xae, 0x61, 0x3f, 0xdc, 0x4c, 0xa2, 0xab, 0xc2, 0xca,
	0x1c, 0x93, 0x52, 0x19, 0xba, 0x3b, 0xe9, 0x4e, 0x07, 0x17, 0x07, 0xab, 0x03, 0x5c, 0x7a, 0x90,
	0x8d, 0x42, 0x80, 0x35, 0xfc, 0x0f, 0x65, 0xd6, 0x76, 0xd1, 0x48, 0xf7, 0xfe, 0xa1, 0x8b, 0x46,
	0x72, 0x06, 0x1d, 0x55, 0xd2, 0xa1, 0x3f, 0xfb, 0xd1, 0x6a, 0xe8, 0x7b, 0x78, 0x2e, 0xac, 0xa3,
	0x4a, 0xf2, 0x16, 0x76, 0xaa, 0x42, 0xa8, 0x22, 0x95, 0xae, 0x03, 0xcf, 0xe8, 0xc8, 0xef, 0x74,
	0xb5, 0x48, 0x3e, 0x42, 0x7f, 0xf9, 0x08, 0x29, 0xf1, 0x93, 0xbc, 0x68, 0xfa, 0x10, 0x6c, 0xf6,
	0x0c, 0x92, 0x77, 0xd0, 0x33, 0x68, 0x6b, 0xae, 0xe9, 0xfe, 0xba, 0xe1, 0xef, 0xd1, 0x3e, 0x70,
	0xcd, 0x02, 0x43, 0x3e, 0xc1, 0xb1, 0x50, 0x85, 0x95, 0xbc, 0x30, 0x89, 0xe0, 0xa5, 0xad, 0x34,
	0x26, 0xcd, 0xb9, 0xe8, 0x81, 0x9f, 0xe9, 0xb0, 0xb5, 0xaf, 0x1a, 0xb7, 0x39, 0x3c, 0xf9, 0x0c,
	0xe3, 0x65, 0x2e, 0xaf, 0x32, 0x2b, 0x73, 0x6e, 0xc5, 0xa2, 0x8d, 0x1e, 0xfa, 0x28, 0x6d, 0x89,
	0xdb, 0x25, 0x10, 0xd2, 0x5f, 0xe0, 0xff, 0x65, 0xda, 0x58, 0x6e, 0xa5, 0x48, 0x66, 0x99, 0x12,
	0x8f, 0x6d, 0xfe, 0xc8, 0xe7, 0x4f, 0x5a, 0xe6, 0xde, 0x23, 0xdf, 0x1c, 0x11, 0x1a, 0x5c, 0xc2,
	0x28, 0x95, 0x46, 0x57, 0xa5, 0x95, 0xf5, 0x72, 0xe0, 0x63, 0xbf, 0xf7, 0xf5, 0x97, 0x35, 0x7c,
	0xc6, 0x43, 0x8b, 0xf7, 0xb0, 0x95, 0xa9, 0x79, 0xca, 0x2d, 0xa7, 0xd4, 0x07, 0x0f, 0x57, 0x83,
	0x37, 0x6a, 0x7e, 0xcd, 0x2d, 0x67, 0x2d, 0x45, 0xde, 0x40, 0x37, 0x37, 0x73, 0x7a, 0xe2, 0xe1,
	0xd1, 0x2a, 0x7c, 0x6b, 0xe6, 0xcc, 0xb9, 0x0e, 0xb2, 0x7c, 0x4e, 0xc7, 0x7e, 0xf5, 0x2f, 0xa0,
	0x9f, 0x7c, 0xce, 0x9c, 0x7b, 0x7a, 0x01, 0xb1, 0x7b, 0xfd, 0x37, 0xd2, 0x58, 0x72, 0x06, 0x1b,
	0xd2, 0x62, 0x4e, 0x23, 0x9f, 0x58, 0xf7, 0x87, 0xf1, 0xfe, 0xac, 0xe7, 0x2b, 0x1f, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xd7, 0x15, 0xb9, 0xae, 0xb2, 0x04, 0x00, 0x00,
}
