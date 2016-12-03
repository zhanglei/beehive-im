// Code generated by protoc-gen-go.
// source: mesg.proto
// DO NOT EDIT!

/*
Package mesg is a generated protocol buffer package.

It is generated from these files:
	mesg.proto

It has these top-level messages:
	MesgOnlineReq
	MesgOnlineAck
	MesgJoinReq
	MesgJoinAck
	MesgUnjoinReq
	MesgUnjoinAck
	MesgLsnRpt
	MesgFrwdRpt
	MesgRoomMsg
	MesgRoomAck
	MesgPrvtMsg
	MesgPrvtAck
*/
package mesg

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

// 运营商ID
type OpId int32

const (
	OpId_OP_UNKNOWN   OpId = 0
	OpId_OP_CN_BGP    OpId = 1
	OpId_OP_CN_UNICOM OpId = 2
	OpId_OP_CN_TELCOM OpId = 3
)

var OpId_name = map[int32]string{
	0: "OP_UNKNOWN",
	1: "OP_CN_BGP",
	2: "OP_CN_UNICOM",
	3: "OP_CN_TELCOM",
}
var OpId_value = map[string]int32{
	"OP_UNKNOWN":   0,
	"OP_CN_BGP":    1,
	"OP_CN_UNICOM": 2,
	"OP_CN_TELCOM": 3,
}

func (x OpId) Enum() *OpId {
	p := new(OpId)
	*p = x
	return p
}
func (x OpId) String() string {
	return proto.EnumName(OpId_name, int32(x))
}
func (x *OpId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OpId_value, data, "OpId")
	if err != nil {
		return err
	}
	*x = OpId(value)
	return nil
}
func (OpId) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MesgOnlineReq struct {
	Uid              *uint64 `protobuf:"varint,1,req,name=Uid" json:"Uid,omitempty"`
	Sid              *uint64 `protobuf:"varint,2,req,name=Sid" json:"Sid,omitempty"`
	Token            *string `protobuf:"bytes,3,req,name=Token" json:"Token,omitempty"`
	App              *string `protobuf:"bytes,4,req,name=App" json:"App,omitempty"`
	Version          *string `protobuf:"bytes,5,req,name=Version" json:"Version,omitempty"`
	Terminal         *uint32 `protobuf:"varint,6,opt,name=Terminal" json:"Terminal,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgOnlineReq) Reset()                    { *m = MesgOnlineReq{} }
func (m *MesgOnlineReq) String() string            { return proto.CompactTextString(m) }
func (*MesgOnlineReq) ProtoMessage()               {}
func (*MesgOnlineReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MesgOnlineReq) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *MesgOnlineReq) GetSid() uint64 {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return 0
}

func (m *MesgOnlineReq) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *MesgOnlineReq) GetApp() string {
	if m != nil && m.App != nil {
		return *m.App
	}
	return ""
}

func (m *MesgOnlineReq) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *MesgOnlineReq) GetTerminal() uint32 {
	if m != nil && m.Terminal != nil {
		return *m.Terminal
	}
	return 0
}

type MesgOnlineAck struct {
	Uid              *uint64 `protobuf:"varint,1,req,name=Uid" json:"Uid,omitempty"`
	Sid              *uint64 `protobuf:"varint,2,req,name=Sid" json:"Sid,omitempty"`
	App              *string `protobuf:"bytes,3,req,name=App" json:"App,omitempty"`
	Version          *string `protobuf:"bytes,4,req,name=Version" json:"Version,omitempty"`
	Terminal         *uint32 `protobuf:"varint,5,opt,name=Terminal" json:"Terminal,omitempty"`
	ErrNum           *uint32 `protobuf:"varint,6,opt,name=ErrNum" json:"ErrNum,omitempty"`
	ErrMsg           *string `protobuf:"bytes,7,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgOnlineAck) Reset()                    { *m = MesgOnlineAck{} }
func (m *MesgOnlineAck) String() string            { return proto.CompactTextString(m) }
func (*MesgOnlineAck) ProtoMessage()               {}
func (*MesgOnlineAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *MesgOnlineAck) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *MesgOnlineAck) GetSid() uint64 {
	if m != nil && m.Sid != nil {
		return *m.Sid
	}
	return 0
}

func (m *MesgOnlineAck) GetApp() string {
	if m != nil && m.App != nil {
		return *m.App
	}
	return ""
}

func (m *MesgOnlineAck) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *MesgOnlineAck) GetTerminal() uint32 {
	if m != nil && m.Terminal != nil {
		return *m.Terminal
	}
	return 0
}

func (m *MesgOnlineAck) GetErrNum() uint32 {
	if m != nil && m.ErrNum != nil {
		return *m.ErrNum
	}
	return 0
}

func (m *MesgOnlineAck) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

type MesgJoinReq struct {
	Uid              *uint64 `protobuf:"varint,1,req,name=Uid" json:"Uid,omitempty"`
	Rid              *uint64 `protobuf:"varint,2,req,name=Rid" json:"Rid,omitempty"`
	Token            *string `protobuf:"bytes,3,req,name=Token" json:"Token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgJoinReq) Reset()                    { *m = MesgJoinReq{} }
func (m *MesgJoinReq) String() string            { return proto.CompactTextString(m) }
func (*MesgJoinReq) ProtoMessage()               {}
func (*MesgJoinReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MesgJoinReq) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *MesgJoinReq) GetRid() uint64 {
	if m != nil && m.Rid != nil {
		return *m.Rid
	}
	return 0
}

func (m *MesgJoinReq) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type MesgJoinAck struct {
	Uid              *uint64 `protobuf:"varint,1,req,name=Uid" json:"Uid,omitempty"`
	Rid              *uint64 `protobuf:"varint,2,req,name=Rid" json:"Rid,omitempty"`
	Gid              *uint32 `protobuf:"varint,3,req,name=Gid" json:"Gid,omitempty"`
	ErrNum           *uint32 `protobuf:"varint,4,opt,name=ErrNum" json:"ErrNum,omitempty"`
	ErrMsg           *string `protobuf:"bytes,5,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgJoinAck) Reset()                    { *m = MesgJoinAck{} }
func (m *MesgJoinAck) String() string            { return proto.CompactTextString(m) }
func (*MesgJoinAck) ProtoMessage()               {}
func (*MesgJoinAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MesgJoinAck) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *MesgJoinAck) GetRid() uint64 {
	if m != nil && m.Rid != nil {
		return *m.Rid
	}
	return 0
}

func (m *MesgJoinAck) GetGid() uint32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *MesgJoinAck) GetErrNum() uint32 {
	if m != nil && m.ErrNum != nil {
		return *m.ErrNum
	}
	return 0
}

func (m *MesgJoinAck) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

type MesgUnjoinReq struct {
	Uid              *uint64 `protobuf:"varint,1,req,name=Uid" json:"Uid,omitempty"`
	Rid              *uint64 `protobuf:"varint,2,req,name=Rid" json:"Rid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgUnjoinReq) Reset()                    { *m = MesgUnjoinReq{} }
func (m *MesgUnjoinReq) String() string            { return proto.CompactTextString(m) }
func (*MesgUnjoinReq) ProtoMessage()               {}
func (*MesgUnjoinReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MesgUnjoinReq) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *MesgUnjoinReq) GetRid() uint64 {
	if m != nil && m.Rid != nil {
		return *m.Rid
	}
	return 0
}

type MesgUnjoinAck struct {
	Uid              *uint64 `protobuf:"varint,1,req,name=Uid" json:"Uid,omitempty"`
	Rid              *uint64 `protobuf:"varint,2,req,name=Rid" json:"Rid,omitempty"`
	ErrNum           *uint32 `protobuf:"varint,3,opt,name=ErrNum" json:"ErrNum,omitempty"`
	ErrMsg           *string `protobuf:"bytes,4,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgUnjoinAck) Reset()                    { *m = MesgUnjoinAck{} }
func (m *MesgUnjoinAck) String() string            { return proto.CompactTextString(m) }
func (*MesgUnjoinAck) ProtoMessage()               {}
func (*MesgUnjoinAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MesgUnjoinAck) GetUid() uint64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *MesgUnjoinAck) GetRid() uint64 {
	if m != nil && m.Rid != nil {
		return *m.Rid
	}
	return 0
}

func (m *MesgUnjoinAck) GetErrNum() uint32 {
	if m != nil && m.ErrNum != nil {
		return *m.ErrNum
	}
	return 0
}

func (m *MesgUnjoinAck) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

type MesgLsnRpt struct {
	Nid              *uint64 `protobuf:"varint,1,req,name=nid" json:"nid,omitempty"`
	Op               *OpId   `protobuf:"varint,2,req,name=op,enum=mesg.OpId,def=0" json:"op,omitempty"`
	Ipaddr           *string `protobuf:"bytes,3,req,name=ipaddr" json:"ipaddr,omitempty"`
	Port             *uint32 `protobuf:"varint,4,req,name=port" json:"port,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgLsnRpt) Reset()                    { *m = MesgLsnRpt{} }
func (m *MesgLsnRpt) String() string            { return proto.CompactTextString(m) }
func (*MesgLsnRpt) ProtoMessage()               {}
func (*MesgLsnRpt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

const Default_MesgLsnRpt_Op OpId = OpId_OP_UNKNOWN

func (m *MesgLsnRpt) GetNid() uint64 {
	if m != nil && m.Nid != nil {
		return *m.Nid
	}
	return 0
}

func (m *MesgLsnRpt) GetOp() OpId {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return Default_MesgLsnRpt_Op
}

func (m *MesgLsnRpt) GetIpaddr() string {
	if m != nil && m.Ipaddr != nil {
		return *m.Ipaddr
	}
	return ""
}

func (m *MesgLsnRpt) GetPort() uint32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

type MesgFrwdRpt struct {
	Nid              *uint64 `protobuf:"varint,1,req,name=nid" json:"nid,omitempty"`
	Ipaddr           *string `protobuf:"bytes,2,req,name=ipaddr" json:"ipaddr,omitempty"`
	ForwardPort      *uint32 `protobuf:"varint,3,req,name=forward_port" json:"forward_port,omitempty"`
	BackendPort      *uint32 `protobuf:"varint,4,req,name=backend_port" json:"backend_port,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgFrwdRpt) Reset()                    { *m = MesgFrwdRpt{} }
func (m *MesgFrwdRpt) String() string            { return proto.CompactTextString(m) }
func (*MesgFrwdRpt) ProtoMessage()               {}
func (*MesgFrwdRpt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MesgFrwdRpt) GetNid() uint64 {
	if m != nil && m.Nid != nil {
		return *m.Nid
	}
	return 0
}

func (m *MesgFrwdRpt) GetIpaddr() string {
	if m != nil && m.Ipaddr != nil {
		return *m.Ipaddr
	}
	return ""
}

func (m *MesgFrwdRpt) GetForwardPort() uint32 {
	if m != nil && m.ForwardPort != nil {
		return *m.ForwardPort
	}
	return 0
}

func (m *MesgFrwdRpt) GetBackendPort() uint32 {
	if m != nil && m.BackendPort != nil {
		return *m.BackendPort
	}
	return 0
}

// 聊天室消息
type MesgRoomMsg struct {
	Rid              *uint64 `protobuf:"varint,1,req,name=Rid" json:"Rid,omitempty"`
	Gid              *uint32 `protobuf:"varint,2,req,name=Gid" json:"Gid,omitempty"`
	Level            *uint32 `protobuf:"varint,3,req,name=level" json:"level,omitempty"`
	Data             []byte  `protobuf:"bytes,4,req,name=Data" json:"Data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgRoomMsg) Reset()                    { *m = MesgRoomMsg{} }
func (m *MesgRoomMsg) String() string            { return proto.CompactTextString(m) }
func (*MesgRoomMsg) ProtoMessage()               {}
func (*MesgRoomMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MesgRoomMsg) GetRid() uint64 {
	if m != nil && m.Rid != nil {
		return *m.Rid
	}
	return 0
}

func (m *MesgRoomMsg) GetGid() uint32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *MesgRoomMsg) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *MesgRoomMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 聊天室消息应答
type MesgRoomAck struct {
	ErrNum           *uint32 `protobuf:"varint,1,opt,name=ErrNum" json:"ErrNum,omitempty"`
	ErrMsg           *string `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgRoomAck) Reset()                    { *m = MesgRoomAck{} }
func (m *MesgRoomAck) String() string            { return proto.CompactTextString(m) }
func (*MesgRoomAck) ProtoMessage()               {}
func (*MesgRoomAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MesgRoomAck) GetErrNum() uint32 {
	if m != nil && m.ErrNum != nil {
		return *m.ErrNum
	}
	return 0
}

func (m *MesgRoomAck) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

// 私聊消息
type MesgPrvtMsg struct {
	Rid              *uint64 `protobuf:"varint,1,req,name=Rid" json:"Rid,omitempty"`
	Gid              *uint32 `protobuf:"varint,2,req,name=Gid" json:"Gid,omitempty"`
	Level            *uint32 `protobuf:"varint,3,req,name=level" json:"level,omitempty"`
	Data             []byte  `protobuf:"bytes,4,req,name=Data" json:"Data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgPrvtMsg) Reset()                    { *m = MesgPrvtMsg{} }
func (m *MesgPrvtMsg) String() string            { return proto.CompactTextString(m) }
func (*MesgPrvtMsg) ProtoMessage()               {}
func (*MesgPrvtMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *MesgPrvtMsg) GetRid() uint64 {
	if m != nil && m.Rid != nil {
		return *m.Rid
	}
	return 0
}

func (m *MesgPrvtMsg) GetGid() uint32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *MesgPrvtMsg) GetLevel() uint32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *MesgPrvtMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 私聊消息应答
type MesgPrvtAck struct {
	ErrNum           *uint32 `protobuf:"varint,1,opt,name=ErrNum" json:"ErrNum,omitempty"`
	ErrMsg           *string `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MesgPrvtAck) Reset()                    { *m = MesgPrvtAck{} }
func (m *MesgPrvtAck) String() string            { return proto.CompactTextString(m) }
func (*MesgPrvtAck) ProtoMessage()               {}
func (*MesgPrvtAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MesgPrvtAck) GetErrNum() uint32 {
	if m != nil && m.ErrNum != nil {
		return *m.ErrNum
	}
	return 0
}

func (m *MesgPrvtAck) GetErrMsg() string {
	if m != nil && m.ErrMsg != nil {
		return *m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*MesgOnlineReq)(nil), "mesg.mesg_online_req")
	proto.RegisterType((*MesgOnlineAck)(nil), "mesg.mesg_online_ack")
	proto.RegisterType((*MesgJoinReq)(nil), "mesg.mesg_join_req")
	proto.RegisterType((*MesgJoinAck)(nil), "mesg.mesg_join_ack")
	proto.RegisterType((*MesgUnjoinReq)(nil), "mesg.mesg_unjoin_req")
	proto.RegisterType((*MesgUnjoinAck)(nil), "mesg.mesg_unjoin_ack")
	proto.RegisterType((*MesgLsnRpt)(nil), "mesg.mesg_lsn_rpt")
	proto.RegisterType((*MesgFrwdRpt)(nil), "mesg.mesg_frwd_rpt")
	proto.RegisterType((*MesgRoomMsg)(nil), "mesg.mesg_room_msg")
	proto.RegisterType((*MesgRoomAck)(nil), "mesg.mesg_room_ack")
	proto.RegisterType((*MesgPrvtMsg)(nil), "mesg.mesg_prvt_msg")
	proto.RegisterType((*MesgPrvtAck)(nil), "mesg.mesg_prvt_ack")
	proto.RegisterEnum("mesg.OpId", OpId_name, OpId_value)
}

func init() { proto.RegisterFile("mesg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xe2, 0x6c, 0xf4, 0x91, 0x6c, 0x51, 0xc5, 0x21, 0xc7, 0xc9, 0xa7, 0x09, 0xa4,
	0x21, 0x71, 0x42, 0xdc, 0x60, 0x8c, 0xa9, 0x1a, 0x4b, 0xa6, 0xd1, 0x0e, 0x2e, 0x28, 0x0a, 0xc4,
	0xab, 0xb2, 0x25, 0xb6, 0x71, 0xb2, 0xed, 0xc0, 0x3f, 0x8f, 0xfd, 0xe2, 0x91, 0x2c, 0x2d, 0x85,
	0xc3, 0x4e, 0xed, 0xf7, 0xda, 0xf7, 0xfb, 0x3e, 0x7f, 0x4e, 0x00, 0x6a, 0xd6, 0x2c, 0x0f, 0xa4,
	0x12, 0xad, 0x98, 0x12, 0xf3, 0x9d, 0x5e, 0xc1, 0xae, 0xf9, 0xcc, 0x04, 0xaf, 0x4a, 0xce, 0x32,
	0xc5, 0x7e, 0x4e, 0x9f, 0x81, 0xb7, 0x28, 0x8b, 0xd8, 0xd9, 0x73, 0xf7, 0x89, 0x11, 0x9f, 0xb5,
	0x70, 0x51, 0x84, 0xe0, 0xcf, 0xc5, 0x35, 0xe3, 0xb1, 0xa7, 0xe5, 0xc4, 0xfc, 0xf6, 0x4e, 0xca,
	0x98, 0xa0, 0xd8, 0x85, 0xed, 0x0b, 0xa6, 0x9a, 0x52, 0xf0, 0xd8, 0xc7, 0x41, 0x04, 0x4f, 0xe7,
	0x4c, 0xd5, 0x25, 0xcf, 0xab, 0x78, 0x6b, 0xcf, 0xd9, 0x0f, 0xe9, 0xaf, 0x87, 0x5e, 0xf9, 0x8f,
	0xeb, 0x0d, 0x5e, 0x16, 0xee, 0x8d, 0xe1, 0x64, 0x05, 0xee, 0x1b, 0xf8, 0x74, 0x07, 0xb6, 0x8e,
	0x94, 0x4a, 0x6e, 0xea, 0xce, 0xcc, 0xea, 0xd3, 0x66, 0x19, 0x6f, 0x6b, 0x3d, 0xa1, 0x6f, 0x20,
	0x44, 0xf3, 0x2b, 0x51, 0xf2, 0xb5, 0xc7, 0x3c, 0xff, 0xcb, 0x31, 0xe9, 0xc5, 0x70, 0x73, 0x5d,
	0xe8, 0xf3, 0x61, 0xe8, 0x63, 0x2d, 0xcc, 0xde, 0x30, 0x11, 0x19, 0x25, 0xf2, 0x31, 0xd1, 0x4b,
	0x5b, 0xc7, 0x0d, 0xff, 0x77, 0x26, 0x7a, 0xf2, 0xf0, 0xcf, 0x9b, 0x63, 0xf4, 0xce, 0xde, 0xc8,
	0x99, 0xa0, 0xf3, 0x37, 0x08, 0x10, 0x56, 0x35, 0xda, 0x56, 0xb6, 0x66, 0x99, 0xff, 0x21, 0x51,
	0x70, 0x85, 0x44, 0xd0, 0xce, 0x6b, 0x38, 0xc0, 0x07, 0x26, 0x95, 0xb3, 0xe2, 0x2d, 0xa4, 0x67,
	0xd9, 0x22, 0x39, 0x49, 0xd2, 0x2f, 0x89, 0x01, 0x96, 0x32, 0x2f, 0x0a, 0x65, 0xef, 0x27, 0x00,
	0x22, 0x85, 0x6a, 0xf1, 0x72, 0x42, 0xfa, 0xd5, 0x16, 0x76, 0xa9, 0xee, 0x8a, 0x55, 0x7e, 0xbf,
	0xeb, 0xe2, 0xee, 0x73, 0x08, 0x2e, 0x85, 0xba, 0xcb, 0x55, 0x91, 0x21, 0xa3, 0x2b, 0x4f, 0x4f,
	0xbf, 0xeb, 0x43, 0x32, 0x6e, 0xa7, 0x1d, 0xf9, 0xa3, 0x25, 0x2b, 0x21, 0xea, 0xac, 0x6e, 0x96,
	0xf7, 0xc7, 0x76, 0x86, 0xed, 0xbb, 0x08, 0xd0, 0x97, 0x58, 0xb1, 0x5b, 0x56, 0x59, 0x9e, 0x4e,
	0xf8, 0x21, 0x6f, 0x73, 0xe4, 0x04, 0xf4, 0xd5, 0x90, 0x63, 0xba, 0xec, 0x1b, 0x73, 0x46, 0x8d,
	0xb9, 0xd8, 0xd8, 0xbd, 0xb1, 0x54, 0xb7, 0xed, 0x63, 0x18, 0x23, 0xe7, 0x3f, 0x8c, 0x5f, 0xcc,
	0x80, 0x98, 0xf6, 0xf5, 0x7c, 0xd0, 0x7f, 0xf4, 0x44, 0xbb, 0x4c, 0xb4, 0x3e, 0x4c, 0xb2, 0xf7,
	0xc7, 0x67, 0x91, 0xa3, 0xdf, 0x87, 0xa0, 0x93, 0x8b, 0x64, 0x76, 0x98, 0x9e, 0x46, 0x6e, 0x3f,
	0x99, 0x1f, 0x7d, 0x32, 0x13, 0xef, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x8d, 0x56, 0xbc,
	0xfd, 0x03, 0x00, 0x00,
}
