// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rename.proto

package dfproto

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

// RPC RenameSquad : RenameSquadIn -> EmptyMessage
type RenameSquadIn struct {
	SquadId              *int32   `protobuf:"varint,1,req,name=squad_id,json=squadId" json:"squad_id,omitempty"`
	Nickname             *string  `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Alias                *string  `protobuf:"bytes,3,opt,name=alias" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameSquadIn) Reset()         { *m = RenameSquadIn{} }
func (m *RenameSquadIn) String() string { return proto.CompactTextString(m) }
func (*RenameSquadIn) ProtoMessage()    {}
func (*RenameSquadIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_rename_27a8cee6274d0ab0, []int{0}
}
func (m *RenameSquadIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameSquadIn.Unmarshal(m, b)
}
func (m *RenameSquadIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameSquadIn.Marshal(b, m, deterministic)
}
func (dst *RenameSquadIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameSquadIn.Merge(dst, src)
}
func (m *RenameSquadIn) XXX_Size() int {
	return xxx_messageInfo_RenameSquadIn.Size(m)
}
func (m *RenameSquadIn) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameSquadIn.DiscardUnknown(m)
}

var xxx_messageInfo_RenameSquadIn proto.InternalMessageInfo

func (m *RenameSquadIn) GetSquadId() int32 {
	if m != nil && m.SquadId != nil {
		return *m.SquadId
	}
	return 0
}

func (m *RenameSquadIn) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *RenameSquadIn) GetAlias() string {
	if m != nil && m.Alias != nil {
		return *m.Alias
	}
	return ""
}

// RPC RenameUnit : RenameUnitIn -> EmptyMessage
type RenameUnitIn struct {
	UnitId               *int32   `protobuf:"varint,1,req,name=unit_id,json=unitId" json:"unit_id,omitempty"`
	Nickname             *string  `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	Profession           *string  `protobuf:"bytes,3,opt,name=profession" json:"profession,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameUnitIn) Reset()         { *m = RenameUnitIn{} }
func (m *RenameUnitIn) String() string { return proto.CompactTextString(m) }
func (*RenameUnitIn) ProtoMessage()    {}
func (*RenameUnitIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_rename_27a8cee6274d0ab0, []int{1}
}
func (m *RenameUnitIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameUnitIn.Unmarshal(m, b)
}
func (m *RenameUnitIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameUnitIn.Marshal(b, m, deterministic)
}
func (dst *RenameUnitIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameUnitIn.Merge(dst, src)
}
func (m *RenameUnitIn) XXX_Size() int {
	return xxx_messageInfo_RenameUnitIn.Size(m)
}
func (m *RenameUnitIn) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameUnitIn.DiscardUnknown(m)
}

var xxx_messageInfo_RenameUnitIn proto.InternalMessageInfo

func (m *RenameUnitIn) GetUnitId() int32 {
	if m != nil && m.UnitId != nil {
		return *m.UnitId
	}
	return 0
}

func (m *RenameUnitIn) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *RenameUnitIn) GetProfession() string {
	if m != nil && m.Profession != nil {
		return *m.Profession
	}
	return ""
}

// RPC RenameBuilding : RenameBuildingIn -> EmptyMessage
type RenameBuildingIn struct {
	BuildingId           *int32   `protobuf:"varint,1,req,name=building_id,json=buildingId" json:"building_id,omitempty"`
	Name                 *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameBuildingIn) Reset()         { *m = RenameBuildingIn{} }
func (m *RenameBuildingIn) String() string { return proto.CompactTextString(m) }
func (*RenameBuildingIn) ProtoMessage()    {}
func (*RenameBuildingIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_rename_27a8cee6274d0ab0, []int{2}
}
func (m *RenameBuildingIn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameBuildingIn.Unmarshal(m, b)
}
func (m *RenameBuildingIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameBuildingIn.Marshal(b, m, deterministic)
}
func (dst *RenameBuildingIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameBuildingIn.Merge(dst, src)
}
func (m *RenameBuildingIn) XXX_Size() int {
	return xxx_messageInfo_RenameBuildingIn.Size(m)
}
func (m *RenameBuildingIn) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameBuildingIn.DiscardUnknown(m)
}

var xxx_messageInfo_RenameBuildingIn proto.InternalMessageInfo

func (m *RenameBuildingIn) GetBuildingId() int32 {
	if m != nil && m.BuildingId != nil {
		return *m.BuildingId
	}
	return 0
}

func (m *RenameBuildingIn) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*RenameSquadIn)(nil), "dfproto.RenameSquadIn")
	proto.RegisterType((*RenameUnitIn)(nil), "dfproto.RenameUnitIn")
	proto.RegisterType((*RenameBuildingIn)(nil), "dfproto.RenameBuildingIn")
}

func init() { proto.RegisterFile("rename.proto", fileDescriptor_rename_27a8cee6274d0ab0) }

var fileDescriptor_rename_27a8cee6274d0ab0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8d, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0x03, 0x88, 0xe0, 0x88, 0x89, 0x99, 0x98, 0x88, 0x1e, 0x94, 0x70, 0xe2, 0xe4, 0x8f,
	0xe0, 0xa2, 0x7b, 0xc5, 0xf4, 0xd6, 0xa4, 0xd9, 0xb2, 0xd0, 0x4c, 0x4a, 0x67, 0x29, 0x0b, 0xff,
	0xbf, 0x61, 0xa1, 0x85, 0x5b, 0x6f, 0xef, 0x9b, 0x37, 0xf9, 0x1e, 0x44, 0x5d, 0xc5, 0xf2, 0x54,
	0xfd, 0xb4, 0x9d, 0xee, 0x35, 0x06, 0xaa, 0xb6, 0x21, 0xdd, 0xc2, 0x4b, 0x61, 0x8b, 0xff, 0xf3,
	0x20, 0x95, 0x60, 0xfc, 0x80, 0xd0, 0x8c, 0x71, 0x47, 0x2a, 0x76, 0x12, 0x37, 0xf3, 0x8b, 0xc0,
	0xb2, 0x50, 0xf8, 0x09, 0x21, 0x53, 0x79, 0x1c, 0xbf, 0x63, 0x37, 0x71, 0xb2, 0xa7, 0xe2, 0xc6,
	0xf8, 0x06, 0xbe, 0x6c, 0x48, 0x9a, 0xd8, 0xb3, 0xc5, 0x04, 0x69, 0x09, 0xd1, 0x64, 0xdf, 0x30,
	0xf5, 0x82, 0xf1, 0x1d, 0x82, 0x81, 0xa9, 0x5f, 0xdc, 0x8f, 0x23, 0xde, 0x51, 0x7f, 0x01, 0xb4,
	0x9d, 0xae, 0x2b, 0x63, 0x48, 0xf3, 0xec, 0x5f, 0x5d, 0xd2, 0x5f, 0x78, 0x9d, 0x46, 0xf2, 0x81,
	0x1a, 0x45, 0x7c, 0x10, 0x8c, 0xdf, 0xf0, 0xbc, 0x9f, 0x69, 0x19, 0x83, 0xeb, 0x49, 0x28, 0x44,
	0x78, 0x58, 0x8d, 0xd9, 0x9c, 0xbb, 0x7f, 0xde, 0x25, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xe1, 0xdf,
	0xbd, 0x27, 0x01, 0x00, 0x00,
}