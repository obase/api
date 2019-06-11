// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/obase/api/x.proto

package x

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Body int32

const (
	Body_JSON Body = 0
	Body_FORM Body = 1
)

var Body_name = map[int32]string{
	0: "JSON",
	1: "FORM",
}

var Body_value = map[string]int32{
	"JSON": 0,
	"FORM": 1,
}

func (x Body) String() string {
	return proto.EnumName(Body_name, int32(x))
}

func (Body) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e7c1a25b3c76cdf, []int{0}
}

type PackFunc struct {
	Pack                 string   `protobuf:"bytes,1,opt,name=pack,proto3" json:"pack,omitempty" bson:"pack"`
	Func                 string   `protobuf:"bytes,2,opt,name=func,proto3" json:"func,omitempty" bson:"func"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *PackFunc) Reset()         { *m = PackFunc{} }
func (m *PackFunc) String() string { return proto.CompactTextString(m) }
func (*PackFunc) ProtoMessage()    {}
func (*PackFunc) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7c1a25b3c76cdf, []int{0}
}

func (m *PackFunc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackFunc.Unmarshal(m, b)
}
func (m *PackFunc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackFunc.Marshal(b, m, deterministic)
}
func (m *PackFunc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackFunc.Merge(m, src)
}
func (m *PackFunc) XXX_Size() int {
	return xxx_messageInfo_PackFunc.Size(m)
}
func (m *PackFunc) XXX_DiscardUnknown() {
	xxx_messageInfo_PackFunc.DiscardUnknown(m)
}

var xxx_messageInfo_PackFunc proto.InternalMessageInfo

func (m *PackFunc) GetPack() string {
	if m != nil {
		return m.Pack
	}
	return ""
}

func (m *PackFunc) GetFunc() string {
	if m != nil {
		return m.Func
	}
	return ""
}

type Group struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty" bson:"path"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty" bson:"desc"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7c1a25b3c76cdf, []int{1}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Group) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type Handle struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty" bson:"path"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty" bson:"desc"`
	Body                 Body     `protobuf:"varint,3,opt,name=body,proto3,enum=api.Body" json:"body,omitempty" bson:"body"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Handle) Reset()         { *m = Handle{} }
func (m *Handle) String() string { return proto.CompactTextString(m) }
func (*Handle) ProtoMessage()    {}
func (*Handle) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7c1a25b3c76cdf, []int{2}
}

func (m *Handle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Handle.Unmarshal(m, b)
}
func (m *Handle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Handle.Marshal(b, m, deterministic)
}
func (m *Handle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handle.Merge(m, src)
}
func (m *Handle) XXX_Size() int {
	return xxx_messageInfo_Handle.Size(m)
}
func (m *Handle) XXX_DiscardUnknown() {
	xxx_messageInfo_Handle.DiscardUnknown(m)
}

var xxx_messageInfo_Handle proto.InternalMessageInfo

func (m *Handle) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Handle) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Handle) GetBody() Body {
	if m != nil {
		return m.Body
	}
	return Body_JSON
}

type Socket struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty" bson:"path"`
	Desc                 string   `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty" bson:"desc"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Socket) Reset()         { *m = Socket{} }
func (m *Socket) String() string { return proto.CompactTextString(m) }
func (*Socket) ProtoMessage()    {}
func (*Socket) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7c1a25b3c76cdf, []int{3}
}

func (m *Socket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Socket.Unmarshal(m, b)
}
func (m *Socket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Socket.Marshal(b, m, deterministic)
}
func (m *Socket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Socket.Merge(m, src)
}
func (m *Socket) XXX_Size() int {
	return xxx_messageInfo_Socket.Size(m)
}
func (m *Socket) XXX_DiscardUnknown() {
	xxx_messageInfo_Socket.DiscardUnknown(m)
}

var xxx_messageInfo_Socket proto.InternalMessageInfo

func (m *Socket) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Socket) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type Field struct {
	Json                 string   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty" bson:"json"`
	Bson                 string   `protobuf:"bytes,2,opt,name=bson,proto3" json:"bson,omitempty" bson:"bson"`
	File                 bool     `protobuf:"varint,3,opt,name=file,proto3" json:"file,omitempty" bson:"file"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7c1a25b3c76cdf, []int{4}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func (m *Field) GetBson() string {
	if m != nil {
		return m.Bson
	}
	return ""
}

func (m *Field) GetFile() bool {
	if m != nil {
		return m.File
	}
	return false
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7c1a25b3c76cdf, []int{5}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

var E_ServerOption = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: ([]*PackFunc)(nil),
	Field:         90001,
	Name:          "api.server_option",
	Tag:           "bytes,90001,rep,name=server_option",
	Filename:      "github.com/obase/api/x.proto",
}

var E_MiddleFilter = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: ([]*PackFunc)(nil),
	Field:         90002,
	Name:          "api.middle_filter",
	Tag:           "bytes,90002,rep,name=middle_filter",
	Filename:      "github.com/obase/api/x.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Field)(nil),
	Field:         90001,
	Name:          "api.field",
	Tag:           "bytes,90001,opt,name=field",
	Filename:      "github.com/obase/api/x.proto",
}

var E_Group = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*Group)(nil),
	Field:         90001,
	Name:          "api.group",
	Tag:           "bytes,90001,opt,name=group",
	Filename:      "github.com/obase/api/x.proto",
}

var E_GroupFilter = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: ([]*PackFunc)(nil),
	Field:         90002,
	Name:          "api.group_filter",
	Tag:           "bytes,90002,rep,name=group_filter",
	Filename:      "github.com/obase/api/x.proto",
}

var E_Handle = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Handle)(nil),
	Field:         90001,
	Name:          "api.handle",
	Tag:           "bytes,90001,opt,name=handle",
	Filename:      "github.com/obase/api/x.proto",
}

var E_HandleFilter = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: ([]*PackFunc)(nil),
	Field:         90002,
	Name:          "api.handle_filter",
	Tag:           "bytes,90002,rep,name=handle_filter",
	Filename:      "github.com/obase/api/x.proto",
}

var E_Socket = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*Socket)(nil),
	Field:         90003,
	Name:          "api.socket",
	Tag:           "bytes,90003,opt,name=socket",
	Filename:      "github.com/obase/api/x.proto",
}

var E_SocketFilter = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: ([]*PackFunc)(nil),
	Field:         90004,
	Name:          "api.socket_filter",
	Tag:           "bytes,90004,rep,name=socket_filter",
	Filename:      "github.com/obase/api/x.proto",
}

func init() {
	proto.RegisterEnum("api.Body", Body_name, Body_value)
	proto.RegisterType((*PackFunc)(nil), "api.PackFunc")
	proto.RegisterType((*Group)(nil), "api.Group")
	proto.RegisterType((*Handle)(nil), "api.Handle")
	proto.RegisterType((*Socket)(nil), "api.Socket")
	proto.RegisterType((*Field)(nil), "api.Field")
	proto.RegisterType((*Void)(nil), "api.Void")
	proto.RegisterExtension(E_ServerOption)
	proto.RegisterExtension(E_MiddleFilter)
	proto.RegisterExtension(E_Field)
	proto.RegisterExtension(E_Group)
	proto.RegisterExtension(E_GroupFilter)
	proto.RegisterExtension(E_Handle)
	proto.RegisterExtension(E_HandleFilter)
	proto.RegisterExtension(E_Socket)
	proto.RegisterExtension(E_SocketFilter)
}

func init() { proto.RegisterFile("github.com/obase/api/x.proto", fileDescriptor_5e7c1a25b3c76cdf) }

var fileDescriptor_5e7c1a25b3c76cdf = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x6e, 0x9b, 0x40,
	0x14, 0x85, 0x4b, 0x0d, 0xc8, 0xb9, 0x8e, 0xab, 0x88, 0x45, 0x85, 0xa2, 0xa4, 0xb5, 0x58, 0x59,
	0x5d, 0x40, 0xe5, 0xee, 0xd8, 0x35, 0x51, 0x69, 0x55, 0x29, 0x75, 0x35, 0xae, 0xba, 0xe8, 0x26,
	0x02, 0x66, 0x30, 0x53, 0x13, 0x06, 0xc1, 0x10, 0x35, 0x8f, 0xd1, 0x9f, 0xf7, 0xe9, 0xab, 0x55,
	0x73, 0x07, 0xe4, 0x46, 0xb6, 0x85, 0x77, 0x47, 0x67, 0x38, 0x9f, 0xcf, 0x9d, 0xb9, 0x32, 0x5c,
	0xac, 0xb9, 0xcc, 0xdb, 0xc4, 0x4f, 0xc5, 0x5d, 0x20, 0x92, 0xb8, 0x61, 0x41, 0x5c, 0xf1, 0xe0,
	0x87, 0x5f, 0xd5, 0x42, 0x0a, 0x67, 0x14, 0x57, 0xfc, 0x7c, 0xb6, 0x16, 0x62, 0x5d, 0xb0, 0x00,
	0xad, 0xa4, 0xcd, 0x02, 0xca, 0x9a, 0xb4, 0xe6, 0x95, 0x14, 0xb5, 0xfe, 0xcc, 0x5b, 0xc0, 0xf8,
	0x73, 0x9c, 0x6e, 0xa2, 0xb6, 0x4c, 0x1d, 0x07, 0xcc, 0x2a, 0x4e, 0x37, 0xae, 0x31, 0x33, 0xe6,
	0x27, 0x04, 0xb5, 0xf2, 0xb2, 0xb6, 0x4c, 0xdd, 0xa7, 0xda, 0x53, 0xda, 0x0b, 0xc0, 0x7a, 0x5f,
	0x8b, 0xb6, 0xd2, 0x01, 0x99, 0x6f, 0x03, 0x32, 0x57, 0x9e, 0xfa, 0x91, 0x3e, 0xa0, 0xb4, 0xb7,
	0x04, 0xfb, 0x43, 0x5c, 0xd2, 0x82, 0x1d, 0x9b, 0x70, 0x2e, 0xc1, 0x4c, 0x04, 0x7d, 0x70, 0x47,
	0x33, 0x63, 0xfe, 0x6c, 0x71, 0xe2, 0xc7, 0x15, 0xf7, 0xaf, 0x04, 0x7d, 0x20, 0x68, 0x7b, 0xaf,
	0xc1, 0x5e, 0x89, 0x74, 0xc3, 0xe4, 0xd1, 0x15, 0xae, 0xc1, 0x8a, 0x38, 0x2b, 0xa8, 0x3a, 0xfc,
	0xde, 0x88, 0xb2, 0x0f, 0x28, 0xad, 0xbc, 0x44, 0x79, 0x5d, 0x20, 0xe9, 0xbc, 0x8c, 0x17, 0x0c,
	0x1b, 0x8c, 0x09, 0x6a, 0xcf, 0x06, 0xf3, 0xab, 0xe0, 0xf4, 0xd5, 0x39, 0x98, 0xaa, 0x8c, 0x33,
	0x06, 0xf3, 0xe3, 0x6a, 0xf9, 0xe9, 0xec, 0x89, 0x52, 0xd1, 0x92, 0xdc, 0x9c, 0x19, 0x21, 0x81,
	0x69, 0xc3, 0xea, 0x7b, 0x56, 0xdf, 0x8a, 0x4a, 0x72, 0x51, 0x3a, 0x17, 0xbe, 0x7e, 0x04, 0xbf,
	0x7f, 0x04, 0x3f, 0xe2, 0x05, 0x5b, 0xe2, 0x61, 0xe3, 0xfe, 0xfc, 0x6b, 0xcd, 0x46, 0xf3, 0xc9,
	0x62, 0x8a, 0x23, 0xf6, 0x4f, 0x41, 0x4e, 0x35, 0x43, 0x7f, 0xa5, 0x98, 0x77, 0x9c, 0xd2, 0x82,
	0xdd, 0x66, 0xbc, 0x90, 0xac, 0x1e, 0x60, 0xfe, 0x3a, 0xc0, 0xd4, 0x8c, 0x08, 0x11, 0xe1, 0x5b,
	0xb0, 0x32, 0xbc, 0x90, 0xcb, 0x3d, 0x2c, 0x56, 0xd0, 0xff, 0x0b, 0x1a, 0xf3, 0xc9, 0x02, 0x10,
	0x86, 0x47, 0x44, 0x27, 0xc3, 0x6b, 0xb0, 0xd6, 0xb8, 0x07, 0x2f, 0x77, 0x10, 0x2b, 0x56, 0xdf,
	0xf3, 0x94, 0xed, 0x87, 0xe0, 0xf2, 0x10, 0x9d, 0x0d, 0x57, 0x70, 0x8a, 0xa2, 0x1f, 0x6d, 0x90,
	0x75, 0x60, 0xba, 0x09, 0x52, 0xba, 0xe1, 0xde, 0x81, 0x9d, 0xeb, 0x85, 0x7b, 0xb1, 0x83, 0xbb,
	0x61, 0x32, 0x17, 0x3b, 0xe3, 0x4d, 0x90, 0xa6, 0xb7, 0x94, 0x74, 0xe1, 0xf0, 0x0b, 0x4c, 0xb5,
	0xea, 0xcb, 0x0d, 0xd1, 0x0e, 0xdd, 0xbc, 0xa6, 0x6c, 0xcb, 0x35, 0x7a, 0x79, 0x87, 0x70, 0xbf,
	0x1f, 0x95, 0xd3, 0x1b, 0x4f, 0xba, 0xb0, 0x2a, 0xa7, 0xd5, 0xb1, 0xe5, 0xfe, 0x1c, 0x5a, 0x35,
	0xa4, 0xe8, 0x72, 0x57, 0xee, 0xb7, 0xe7, 0xfb, 0xff, 0x56, 0x12, 0x1b, 0xb1, 0x6f, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x29, 0x9f, 0x88, 0xdc, 0x77, 0x04, 0x00, 0x00,
}
