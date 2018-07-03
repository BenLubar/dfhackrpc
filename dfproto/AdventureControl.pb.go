// Code generated by protoc-gen-go. DO NOT EDIT.
// source: AdventureControl.proto

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

type AdvmodeMenu int32

const (
	AdvmodeMenu_Default                 AdvmodeMenu = 0
	AdvmodeMenu_Look                    AdvmodeMenu = 1
	AdvmodeMenu_ConversationAddress     AdvmodeMenu = 2
	AdvmodeMenu_ConversationSelect      AdvmodeMenu = 3
	AdvmodeMenu_ConversationSpeak       AdvmodeMenu = 4
	AdvmodeMenu_Inventory               AdvmodeMenu = 5
	AdvmodeMenu_Drop                    AdvmodeMenu = 6
	AdvmodeMenu_ThrowItem               AdvmodeMenu = 7
	AdvmodeMenu_Wear                    AdvmodeMenu = 8
	AdvmodeMenu_Remove                  AdvmodeMenu = 9
	AdvmodeMenu_Interact                AdvmodeMenu = 10
	AdvmodeMenu_Put                     AdvmodeMenu = 11
	AdvmodeMenu_PutContainer            AdvmodeMenu = 12
	AdvmodeMenu_Eat                     AdvmodeMenu = 13
	AdvmodeMenu_ThrowAim                AdvmodeMenu = 14
	AdvmodeMenu_Fire                    AdvmodeMenu = 15
	AdvmodeMenu_Get                     AdvmodeMenu = 16
	AdvmodeMenu_Unk17                   AdvmodeMenu = 17
	AdvmodeMenu_CombatPrefs             AdvmodeMenu = 18
	AdvmodeMenu_Companions              AdvmodeMenu = 19
	AdvmodeMenu_MovementPrefs           AdvmodeMenu = 20
	AdvmodeMenu_SpeedPrefs              AdvmodeMenu = 21
	AdvmodeMenu_InteractAction          AdvmodeMenu = 22
	AdvmodeMenu_MoveCarefully           AdvmodeMenu = 23
	AdvmodeMenu_Announcements           AdvmodeMenu = 24
	AdvmodeMenu_UseBuilding             AdvmodeMenu = 25
	AdvmodeMenu_Travel                  AdvmodeMenu = 26
	AdvmodeMenu_Unk27                   AdvmodeMenu = 27
	AdvmodeMenu_Unk28                   AdvmodeMenu = 28
	AdvmodeMenu_SleepConfirm            AdvmodeMenu = 29
	AdvmodeMenu_SelectInteractionTarget AdvmodeMenu = 30
	AdvmodeMenu_Unk31                   AdvmodeMenu = 31
	AdvmodeMenu_Unk32                   AdvmodeMenu = 32
	AdvmodeMenu_FallAction              AdvmodeMenu = 33
	AdvmodeMenu_ViewTracks              AdvmodeMenu = 34
	AdvmodeMenu_Jump                    AdvmodeMenu = 35
	AdvmodeMenu_Unk36                   AdvmodeMenu = 36
	AdvmodeMenu_AttackConfirm           AdvmodeMenu = 37
	AdvmodeMenu_AttackType              AdvmodeMenu = 38
	AdvmodeMenu_AttackBodypart          AdvmodeMenu = 39
	AdvmodeMenu_AttackStrike            AdvmodeMenu = 40
	AdvmodeMenu_Unk41                   AdvmodeMenu = 41
	AdvmodeMenu_Unk42                   AdvmodeMenu = 42
	AdvmodeMenu_DodgeDirection          AdvmodeMenu = 43
	AdvmodeMenu_Unk44                   AdvmodeMenu = 44
	AdvmodeMenu_Unk45                   AdvmodeMenu = 45
	AdvmodeMenu_Build                   AdvmodeMenu = 46
)

var AdvmodeMenu_name = map[int32]string{
	0:  "Default",
	1:  "Look",
	2:  "ConversationAddress",
	3:  "ConversationSelect",
	4:  "ConversationSpeak",
	5:  "Inventory",
	6:  "Drop",
	7:  "ThrowItem",
	8:  "Wear",
	9:  "Remove",
	10: "Interact",
	11: "Put",
	12: "PutContainer",
	13: "Eat",
	14: "ThrowAim",
	15: "Fire",
	16: "Get",
	17: "Unk17",
	18: "CombatPrefs",
	19: "Companions",
	20: "MovementPrefs",
	21: "SpeedPrefs",
	22: "InteractAction",
	23: "MoveCarefully",
	24: "Announcements",
	25: "UseBuilding",
	26: "Travel",
	27: "Unk27",
	28: "Unk28",
	29: "SleepConfirm",
	30: "SelectInteractionTarget",
	31: "Unk31",
	32: "Unk32",
	33: "FallAction",
	34: "ViewTracks",
	35: "Jump",
	36: "Unk36",
	37: "AttackConfirm",
	38: "AttackType",
	39: "AttackBodypart",
	40: "AttackStrike",
	41: "Unk41",
	42: "Unk42",
	43: "DodgeDirection",
	44: "Unk44",
	45: "Unk45",
	46: "Build",
}
var AdvmodeMenu_value = map[string]int32{
	"Default":                 0,
	"Look":                    1,
	"ConversationAddress":     2,
	"ConversationSelect":      3,
	"ConversationSpeak":       4,
	"Inventory":               5,
	"Drop":                    6,
	"ThrowItem":               7,
	"Wear":                    8,
	"Remove":                  9,
	"Interact":                10,
	"Put":                     11,
	"PutContainer":            12,
	"Eat":                     13,
	"ThrowAim":                14,
	"Fire":                    15,
	"Get":                     16,
	"Unk17":                   17,
	"CombatPrefs":             18,
	"Companions":              19,
	"MovementPrefs":           20,
	"SpeedPrefs":              21,
	"InteractAction":          22,
	"MoveCarefully":           23,
	"Announcements":           24,
	"UseBuilding":             25,
	"Travel":                  26,
	"Unk27":                   27,
	"Unk28":                   28,
	"SleepConfirm":            29,
	"SelectInteractionTarget": 30,
	"Unk31":                   31,
	"Unk32":                   32,
	"FallAction":              33,
	"ViewTracks":              34,
	"Jump":                    35,
	"Unk36":                   36,
	"AttackConfirm":           37,
	"AttackType":              38,
	"AttackBodypart":          39,
	"AttackStrike":            40,
	"Unk41":                   41,
	"Unk42":                   42,
	"DodgeDirection":          43,
	"Unk44":                   44,
	"Unk45":                   45,
	"Build":                   46,
}

func (x AdvmodeMenu) Enum() *AdvmodeMenu {
	p := new(AdvmodeMenu)
	*p = x
	return p
}
func (x AdvmodeMenu) String() string {
	return proto.EnumName(AdvmodeMenu_name, int32(x))
}
func (x *AdvmodeMenu) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AdvmodeMenu_value, data, "AdvmodeMenu")
	if err != nil {
		return err
	}
	*x = AdvmodeMenu(value)
	return nil
}
func (AdvmodeMenu) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_AdventureControl_03cf9a2c25a3c3b9, []int{0}
}

type CarefulMovementType int32

const (
	CarefulMovementType_DEFAULT_MOVEMENT    CarefulMovementType = 0
	CarefulMovementType_RELEASE_ITEM_HOLD   CarefulMovementType = 1
	CarefulMovementType_RELEASE_TILE_HOLD   CarefulMovementType = 2
	CarefulMovementType_ATTACK_CREATURE     CarefulMovementType = 3
	CarefulMovementType_HOLD_TILE           CarefulMovementType = 4
	CarefulMovementType_MOVE                CarefulMovementType = 5
	CarefulMovementType_CLIMB               CarefulMovementType = 6
	CarefulMovementType_HOLD_ITEM           CarefulMovementType = 7
	CarefulMovementType_BUILDING_INTERACT   CarefulMovementType = 8
	CarefulMovementType_ITEM_INTERACT       CarefulMovementType = 9
	CarefulMovementType_ITEM_INTERACT_GUIDE CarefulMovementType = 10
	CarefulMovementType_ITEM_INTERACT_RIDE  CarefulMovementType = 11
	CarefulMovementType_ITEM_INTERACT_PUSH  CarefulMovementType = 12
)

var CarefulMovementType_name = map[int32]string{
	0:  "DEFAULT_MOVEMENT",
	1:  "RELEASE_ITEM_HOLD",
	2:  "RELEASE_TILE_HOLD",
	3:  "ATTACK_CREATURE",
	4:  "HOLD_TILE",
	5:  "MOVE",
	6:  "CLIMB",
	7:  "HOLD_ITEM",
	8:  "BUILDING_INTERACT",
	9:  "ITEM_INTERACT",
	10: "ITEM_INTERACT_GUIDE",
	11: "ITEM_INTERACT_RIDE",
	12: "ITEM_INTERACT_PUSH",
}
var CarefulMovementType_value = map[string]int32{
	"DEFAULT_MOVEMENT":    0,
	"RELEASE_ITEM_HOLD":   1,
	"RELEASE_TILE_HOLD":   2,
	"ATTACK_CREATURE":     3,
	"HOLD_TILE":           4,
	"MOVE":                5,
	"CLIMB":               6,
	"HOLD_ITEM":           7,
	"BUILDING_INTERACT":   8,
	"ITEM_INTERACT":       9,
	"ITEM_INTERACT_GUIDE": 10,
	"ITEM_INTERACT_RIDE":  11,
	"ITEM_INTERACT_PUSH":  12,
}

func (x CarefulMovementType) Enum() *CarefulMovementType {
	p := new(CarefulMovementType)
	*p = x
	return p
}
func (x CarefulMovementType) String() string {
	return proto.EnumName(CarefulMovementType_name, int32(x))
}
func (x *CarefulMovementType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CarefulMovementType_value, data, "CarefulMovementType")
	if err != nil {
		return err
	}
	*x = CarefulMovementType(value)
	return nil
}
func (CarefulMovementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_AdventureControl_03cf9a2c25a3c3b9, []int{1}
}

type MiscMoveType int32

const (
	MiscMoveType_SET_CLIMB  MiscMoveType = 0
	MiscMoveType_SET_STAND  MiscMoveType = 1
	MiscMoveType_SET_CANCEL MiscMoveType = 2
)

var MiscMoveType_name = map[int32]string{
	0: "SET_CLIMB",
	1: "SET_STAND",
	2: "SET_CANCEL",
}
var MiscMoveType_value = map[string]int32{
	"SET_CLIMB":  0,
	"SET_STAND":  1,
	"SET_CANCEL": 2,
}

func (x MiscMoveType) Enum() *MiscMoveType {
	p := new(MiscMoveType)
	*p = x
	return p
}
func (x MiscMoveType) String() string {
	return proto.EnumName(MiscMoveType_name, int32(x))
}
func (x *MiscMoveType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MiscMoveType_value, data, "MiscMoveType")
	if err != nil {
		return err
	}
	*x = MiscMoveType(value)
	return nil
}
func (MiscMoveType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_AdventureControl_03cf9a2c25a3c3b9, []int{2}
}

type MoveCommandParams struct {
	Direction            *Coord   `protobuf:"bytes,1,opt,name=direction" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoveCommandParams) Reset()         { *m = MoveCommandParams{} }
func (m *MoveCommandParams) String() string { return proto.CompactTextString(m) }
func (*MoveCommandParams) ProtoMessage()    {}
func (*MoveCommandParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_AdventureControl_03cf9a2c25a3c3b9, []int{0}
}
func (m *MoveCommandParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveCommandParams.Unmarshal(m, b)
}
func (m *MoveCommandParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveCommandParams.Marshal(b, m, deterministic)
}
func (dst *MoveCommandParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveCommandParams.Merge(dst, src)
}
func (m *MoveCommandParams) XXX_Size() int {
	return xxx_messageInfo_MoveCommandParams.Size(m)
}
func (m *MoveCommandParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveCommandParams.DiscardUnknown(m)
}

var xxx_messageInfo_MoveCommandParams proto.InternalMessageInfo

func (m *MoveCommandParams) GetDirection() *Coord {
	if m != nil {
		return m.Direction
	}
	return nil
}

type MovementOption struct {
	Dest                 *Coord               `protobuf:"bytes,1,opt,name=dest" json:"dest,omitempty"`
	Source               *Coord               `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Grab                 *Coord               `protobuf:"bytes,3,opt,name=grab" json:"grab,omitempty"`
	MovementType         *CarefulMovementType `protobuf:"varint,4,opt,name=movement_type,json=movementType,enum=AdventureControl.CarefulMovementType" json:"movement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MovementOption) Reset()         { *m = MovementOption{} }
func (m *MovementOption) String() string { return proto.CompactTextString(m) }
func (*MovementOption) ProtoMessage()    {}
func (*MovementOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_AdventureControl_03cf9a2c25a3c3b9, []int{1}
}
func (m *MovementOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovementOption.Unmarshal(m, b)
}
func (m *MovementOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovementOption.Marshal(b, m, deterministic)
}
func (dst *MovementOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovementOption.Merge(dst, src)
}
func (m *MovementOption) XXX_Size() int {
	return xxx_messageInfo_MovementOption.Size(m)
}
func (m *MovementOption) XXX_DiscardUnknown() {
	xxx_messageInfo_MovementOption.DiscardUnknown(m)
}

var xxx_messageInfo_MovementOption proto.InternalMessageInfo

func (m *MovementOption) GetDest() *Coord {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *MovementOption) GetSource() *Coord {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *MovementOption) GetGrab() *Coord {
	if m != nil {
		return m.Grab
	}
	return nil
}

func (m *MovementOption) GetMovementType() CarefulMovementType {
	if m != nil && m.MovementType != nil {
		return *m.MovementType
	}
	return CarefulMovementType_DEFAULT_MOVEMENT
}

type MenuContents struct {
	CurrentMenu          *AdvmodeMenu      `protobuf:"varint,1,opt,name=current_menu,json=currentMenu,enum=AdventureControl.AdvmodeMenu" json:"current_menu,omitempty"`
	Movements            []*MovementOption `protobuf:"bytes,2,rep,name=movements" json:"movements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MenuContents) Reset()         { *m = MenuContents{} }
func (m *MenuContents) String() string { return proto.CompactTextString(m) }
func (*MenuContents) ProtoMessage()    {}
func (*MenuContents) Descriptor() ([]byte, []int) {
	return fileDescriptor_AdventureControl_03cf9a2c25a3c3b9, []int{2}
}
func (m *MenuContents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MenuContents.Unmarshal(m, b)
}
func (m *MenuContents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MenuContents.Marshal(b, m, deterministic)
}
func (dst *MenuContents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MenuContents.Merge(dst, src)
}
func (m *MenuContents) XXX_Size() int {
	return xxx_messageInfo_MenuContents.Size(m)
}
func (m *MenuContents) XXX_DiscardUnknown() {
	xxx_messageInfo_MenuContents.DiscardUnknown(m)
}

var xxx_messageInfo_MenuContents proto.InternalMessageInfo

func (m *MenuContents) GetCurrentMenu() AdvmodeMenu {
	if m != nil && m.CurrentMenu != nil {
		return *m.CurrentMenu
	}
	return AdvmodeMenu_Default
}

func (m *MenuContents) GetMovements() []*MovementOption {
	if m != nil {
		return m.Movements
	}
	return nil
}

type MiscMoveParams struct {
	Type                 *MiscMoveType `protobuf:"varint,1,opt,name=type,enum=AdventureControl.MiscMoveType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MiscMoveParams) Reset()         { *m = MiscMoveParams{} }
func (m *MiscMoveParams) String() string { return proto.CompactTextString(m) }
func (*MiscMoveParams) ProtoMessage()    {}
func (*MiscMoveParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_AdventureControl_03cf9a2c25a3c3b9, []int{3}
}
func (m *MiscMoveParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MiscMoveParams.Unmarshal(m, b)
}
func (m *MiscMoveParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MiscMoveParams.Marshal(b, m, deterministic)
}
func (dst *MiscMoveParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiscMoveParams.Merge(dst, src)
}
func (m *MiscMoveParams) XXX_Size() int {
	return xxx_messageInfo_MiscMoveParams.Size(m)
}
func (m *MiscMoveParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MiscMoveParams.DiscardUnknown(m)
}

var xxx_messageInfo_MiscMoveParams proto.InternalMessageInfo

func (m *MiscMoveParams) GetType() MiscMoveType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MiscMoveType_SET_CLIMB
}

func init() {
	proto.RegisterType((*MoveCommandParams)(nil), "AdventureControl.MoveCommandParams")
	proto.RegisterType((*MovementOption)(nil), "AdventureControl.MovementOption")
	proto.RegisterType((*MenuContents)(nil), "AdventureControl.MenuContents")
	proto.RegisterType((*MiscMoveParams)(nil), "AdventureControl.MiscMoveParams")
	proto.RegisterEnum("AdventureControl.AdvmodeMenu", AdvmodeMenu_name, AdvmodeMenu_value)
	proto.RegisterEnum("AdventureControl.CarefulMovementType", CarefulMovementType_name, CarefulMovementType_value)
	proto.RegisterEnum("AdventureControl.MiscMoveType", MiscMoveType_name, MiscMoveType_value)
}

func init() {
	proto.RegisterFile("AdventureControl.proto", fileDescriptor_AdventureControl_03cf9a2c25a3c3b9)
}

var fileDescriptor_AdventureControl_03cf9a2c25a3c3b9 = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x6d, 0x73, 0xdb, 0x44,
	0x10, 0x8e, 0x5f, 0xf2, 0xe2, 0xb3, 0xe3, 0x6c, 0x2e, 0x6d, 0x62, 0x12, 0x5a, 0x42, 0xa0, 0x10,
	0x02, 0x84, 0x89, 0x5b, 0x28, 0xcc, 0x30, 0x0c, 0x8a, 0xac, 0x24, 0x2e, 0xb6, 0x93, 0x91, 0xe5,
	0xf2, 0xd1, 0x73, 0xb5, 0x36, 0x41, 0x63, 0xeb, 0x4e, 0x73, 0x3a, 0xb9, 0x93, 0x7f, 0xc1, 0x07,
	0xfe, 0x13, 0x3f, 0x88, 0x3f, 0xc0, 0xac, 0x5e, 0xa8, 0x43, 0x32, 0xd3, 0x7e, 0xd3, 0x3d, 0xfb,
	0x3c, 0xbb, 0xcf, 0xee, 0xed, 0x89, 0x6d, 0x5b, 0xfe, 0x1c, 0xa5, 0x49, 0x34, 0xda, 0x4a, 0x1a,
	0xad, 0x66, 0xc7, 0x91, 0x56, 0x46, 0x71, 0xf8, 0x3f, 0xbe, 0xbb, 0xeb, 0x62, 0xa8, 0x0c, 0x9e,
	0x29, 0x6d, 0x34, 0xc6, 0xb1, 0x8b, 0xc2, 0x47, 0x9d, 0xb1, 0x0f, 0x06, 0x6c, 0xb3, 0xaf, 0xe6,
	0x68, 0xab, 0x30, 0x14, 0xd2, 0xbf, 0x12, 0x5a, 0x84, 0x31, 0xff, 0x89, 0xd5, 0xfc, 0x40, 0xe3,
	0xc4, 0x04, 0x4a, 0xb6, 0x4a, 0xfb, 0xa5, 0xc3, 0x7a, 0x7b, 0xef, 0xf8, 0xc1, 0x24, 0xb6, 0x52,
	0xda, 0x77, 0xdf, 0xb1, 0x0f, 0xfe, 0x29, 0xb1, 0x26, 0x25, 0x0c, 0x51, 0x9a, 0xcb, 0x88, 0x20,
	0xfe, 0x1d, 0xab, 0xfa, 0x18, 0x9b, 0x0f, 0x49, 0x94, 0x12, 0xf9, 0x73, 0xb6, 0x12, 0xab, 0x44,
	0x4f, 0xb0, 0x55, 0x7e, 0xbf, 0x24, 0xa7, 0x52, 0x95, 0x1b, 0x2d, 0xde, 0xb4, 0x2a, 0x1f, 0x50,
	0x85, 0x88, 0xfc, 0x15, 0x5b, 0x0f, 0x73, 0xa3, 0x63, 0x73, 0x1b, 0x61, 0xab, 0xba, 0x5f, 0x3a,
	0x6c, 0xb6, 0x9f, 0x1d, 0xdf, 0x9b, 0xab, 0x2d, 0x34, 0x5e, 0x27, 0xb3, 0xa2, 0x2d, 0xef, 0x36,
	0x42, 0xb7, 0x11, 0x2e, 0x9c, 0x0e, 0xfe, 0x2c, 0xb1, 0x46, 0x1f, 0x65, 0x42, 0x0a, 0x94, 0x26,
	0xe6, 0xbf, 0xb2, 0xc6, 0x24, 0xd1, 0x9a, 0x72, 0x87, 0x28, 0x93, 0xb4, 0xf7, 0x66, 0xfb, 0xc9,
	0xfd, 0xdc, 0x96, 0x3f, 0x0f, 0x95, 0x8f, 0x24, 0x76, 0xeb, 0xb9, 0x84, 0x0e, 0xfc, 0x17, 0x56,
	0x2b, 0x4a, 0xc4, 0xad, 0xf2, 0x7e, 0xe5, 0xb0, 0xde, 0xde, 0xbf, 0x2f, 0xbf, 0x3b, 0x6a, 0xf7,
	0x9d, 0xe4, 0xa0, 0xc3, 0x9a, 0xfd, 0x20, 0x9e, 0x10, 0x21, 0xbf, 0xd5, 0x36, 0xab, 0xa6, 0x7d,
	0x66, 0x5e, 0x9e, 0x3e, 0x90, 0x2c, 0xe7, 0xa7, 0x0d, 0xa6, 0xdc, 0xa3, 0xbf, 0x97, 0x59, 0x7d,
	0xc1, 0x22, 0xaf, 0xb3, 0xd5, 0x0e, 0x5e, 0x8b, 0x64, 0x66, 0x60, 0x89, 0xaf, 0xb1, 0x6a, 0x4f,
	0xa9, 0x29, 0x94, 0xf8, 0x0e, 0xdb, 0xb2, 0x95, 0x9c, 0xa3, 0x8e, 0x05, 0xf9, 0xb0, 0x7c, 0x9f,
	0x86, 0x0e, 0x65, 0xbe, 0xcd, 0xf8, 0x62, 0x60, 0x88, 0x33, 0x9c, 0x18, 0xa8, 0xf0, 0xc7, 0x6c,
	0xf3, 0x0e, 0x1e, 0xa1, 0x98, 0x42, 0x95, 0xaf, 0xb3, 0x5a, 0x57, 0x92, 0x2b, 0xa5, 0x6f, 0x61,
	0x99, 0x0a, 0x74, 0xb4, 0x8a, 0x60, 0x85, 0x02, 0xde, 0x1f, 0x5a, 0xbd, 0xed, 0x1a, 0x0c, 0x61,
	0x95, 0x02, 0xbf, 0xa3, 0xd0, 0xb0, 0xc6, 0x19, 0x5b, 0xa1, 0x9b, 0x9e, 0x23, 0xd4, 0x78, 0x83,
	0xad, 0x75, 0xa5, 0x41, 0x2d, 0x26, 0x06, 0x18, 0x5f, 0x65, 0x95, 0xab, 0xc4, 0x40, 0x9d, 0x03,
	0x6b, 0x5c, 0x25, 0x86, 0x9a, 0x14, 0x81, 0x44, 0x0d, 0x0d, 0x0a, 0x39, 0xc2, 0xc0, 0x3a, 0x29,
	0xd2, 0xb4, 0x56, 0x10, 0x42, 0x93, 0xb2, 0x9e, 0x05, 0x1a, 0x61, 0x83, 0x08, 0xe7, 0x68, 0x00,
	0x78, 0x8d, 0x2d, 0x8f, 0xe4, 0xf4, 0xe4, 0x25, 0x6c, 0xf2, 0x0d, 0x56, 0xb7, 0x55, 0xf8, 0x46,
	0x98, 0x2b, 0x8d, 0xd7, 0x31, 0x70, 0xde, 0x64, 0xcc, 0x56, 0x61, 0x24, 0x64, 0xa0, 0x64, 0x0c,
	0x5b, 0x7c, 0x93, 0xad, 0x17, 0xd7, 0x91, 0x51, 0x1e, 0x11, 0x65, 0x18, 0x21, 0xfa, 0xd9, 0xf9,
	0x31, 0xe7, 0xac, 0x59, 0x38, 0xb4, 0xd2, 0xf7, 0x02, 0xdb, 0x85, 0x2c, 0x5f, 0xb2, 0xd9, 0x2d,
	0xec, 0x10, 0x64, 0x49, 0xa9, 0x12, 0x39, 0xc9, 0x2e, 0x13, 0x5a, 0x54, 0x7d, 0x14, 0xe3, 0x69,
	0x12, 0xcc, 0xfc, 0x40, 0xde, 0xc0, 0x47, 0xd4, 0xb8, 0xa7, 0xc5, 0x1c, 0x67, 0xb0, 0x9b, 0xbb,
	0x6c, 0xbf, 0x84, 0xbd, 0xe2, 0xf3, 0x47, 0xf8, 0x98, 0xfa, 0x1e, 0xce, 0x10, 0x23, 0x5b, 0xc9,
	0xeb, 0x40, 0x87, 0xf0, 0x84, 0xef, 0xb1, 0x9d, 0xec, 0x06, 0x0a, 0x13, 0x81, 0x92, 0x9e, 0xd0,
	0x37, 0x68, 0xe0, 0x69, 0xae, 0x7c, 0x7e, 0x02, 0x9f, 0x14, 0x9f, 0x6d, 0xd8, 0xa7, 0x0e, 0xce,
	0xc4, 0x6c, 0x96, 0xbb, 0xfd, 0x94, 0xce, 0xaf, 0x03, 0x7c, 0xeb, 0x69, 0x31, 0x99, 0xc6, 0x70,
	0x40, 0x33, 0x7b, 0x95, 0x84, 0x11, 0x7c, 0x56, 0x88, 0x7e, 0x80, 0xcf, 0x53, 0xff, 0xc6, 0x88,
	0xc9, 0xb4, 0x28, 0xfd, 0x8c, 0x74, 0x19, 0x44, 0xcb, 0x05, 0x5f, 0xd0, 0x24, 0xb2, 0xf3, 0xa9,
	0xf2, 0x6f, 0x23, 0xa1, 0x0d, 0x7c, 0x49, 0x86, 0x33, 0x6c, 0x68, 0x74, 0x30, 0x45, 0x38, 0xcc,
	0x73, 0xbe, 0x38, 0x81, 0xaf, 0x8a, 0xcf, 0x36, 0x1c, 0x91, 0xb6, 0xa3, 0xfc, 0x1b, 0xec, 0x14,
	0x7f, 0x1d, 0xf8, 0xba, 0x08, 0xbf, 0x80, 0x6f, 0x8a, 0xcf, 0xef, 0xe1, 0x5b, 0xfa, 0x4c, 0x47,
	0x06, 0xc7, 0x47, 0x7f, 0x95, 0xd9, 0xd6, 0x03, 0x0f, 0x99, 0x3f, 0x62, 0xd0, 0x71, 0xce, 0xac,
	0x51, 0xcf, 0x1b, 0xf7, 0x2f, 0x5f, 0x3b, 0x7d, 0x67, 0xe0, 0xc1, 0x12, 0xed, 0xa7, 0xeb, 0xf4,
	0x1c, 0x6b, 0xe8, 0x8c, 0xbb, 0x9e, 0xd3, 0x1f, 0x5f, 0x5c, 0xf6, 0x3a, 0x50, 0x5a, 0x84, 0xbd,
	0x6e, 0xcf, 0xc9, 0xe0, 0x32, 0xdf, 0x62, 0x1b, 0x96, 0xe7, 0x59, 0xf6, 0x6f, 0x63, 0xdb, 0x75,
	0x2c, 0x6f, 0xe4, 0x3a, 0x50, 0xa1, 0x95, 0xa5, 0x70, 0x4a, 0x84, 0x2a, 0x0d, 0x8a, 0xf2, 0xc3,
	0x32, 0x99, 0xb2, 0x7b, 0xdd, 0xfe, 0x69, 0xb6, 0xd6, 0x29, 0x87, 0x6a, 0xc0, 0x2a, 0xa5, 0x3f,
	0x1d, 0x75, 0x7b, 0x9d, 0xee, 0xe0, 0x7c, 0xdc, 0x1d, 0x78, 0x8e, 0x6b, 0xd9, 0x1e, 0xac, 0xd1,
	0x38, 0x53, 0x13, 0xff, 0x41, 0x35, 0x7a, 0x70, 0x77, 0xa0, 0xf1, 0xf9, 0xa8, 0xdb, 0x71, 0x80,
	0xd1, 0x83, 0xbb, 0x1b, 0x70, 0x09, 0xaf, 0xdf, 0xc7, 0xaf, 0x46, 0xc3, 0x0b, 0x68, 0x1c, 0xfd,
	0xcc, 0x1a, 0x8b, 0xcf, 0x9e, 0x1c, 0x0d, 0x1d, 0x6f, 0x9c, 0x19, 0x5c, 0x2a, 0x8e, 0x43, 0xcf,
	0x1a, 0x50, 0xff, 0xb4, 0xcf, 0x14, 0xb5, 0x06, 0xb6, 0xd3, 0x83, 0xf2, 0x69, 0xf9, 0xa2, 0xf2,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xd2, 0x25, 0x44, 0x88, 0x06, 0x00, 0x00,
}
