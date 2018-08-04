// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_ModifyOrder/Trd_ModifyOrder.proto

package Trd_ModifyOrder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import Common "limgo/futupb/Common"
import Trd_Common "limgo/futupb/Trd_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type C2S struct {
	PacketID      *Common.PacketID      `protobuf:"bytes,1,req,name=packetID" json:"packetID,omitempty"`
	Header        *Trd_Common.TrdHeader `protobuf:"bytes,2,req,name=header" json:"header,omitempty"`
	OrderID       *uint64               `protobuf:"varint,3,req,name=orderID" json:"orderID,omitempty"`
	ModifyOrderOp *int32                `protobuf:"varint,4,req,name=modifyOrderOp" json:"modifyOrderOp,omitempty"`
	ForAll        *bool                 `protobuf:"varint,5,opt,name=forAll" json:"forAll,omitempty"`
	// 下面的字段仅在modifyOrderOp为ModifyOrderOp_Normal有效
	Qty   *float64 `protobuf:"fixed64,8,opt,name=qty" json:"qty,omitempty"`
	Price *float64 `protobuf:"fixed64,9,opt,name=price" json:"price,omitempty"`
	// 以下为调整价格使用，都传才有效，对港、A股有意义，因为港股有价位，A股2位精度，美股可不传
	AdjustPrice          *bool    `protobuf:"varint,10,opt,name=adjustPrice" json:"adjustPrice,omitempty"`
	AdjustSideAndLimit   *float64 `protobuf:"fixed64,11,opt,name=adjustSideAndLimit" json:"adjustSideAndLimit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_ModifyOrder_511b294f45632e93, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (dst *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(dst, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetPacketID() *Common.PacketID {
	if m != nil {
		return m.PacketID
	}
	return nil
}

func (m *C2S) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *C2S) GetOrderID() uint64 {
	if m != nil && m.OrderID != nil {
		return *m.OrderID
	}
	return 0
}

func (m *C2S) GetModifyOrderOp() int32 {
	if m != nil && m.ModifyOrderOp != nil {
		return *m.ModifyOrderOp
	}
	return 0
}

func (m *C2S) GetForAll() bool {
	if m != nil && m.ForAll != nil {
		return *m.ForAll
	}
	return false
}

func (m *C2S) GetQty() float64 {
	if m != nil && m.Qty != nil {
		return *m.Qty
	}
	return 0
}

func (m *C2S) GetPrice() float64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *C2S) GetAdjustPrice() bool {
	if m != nil && m.AdjustPrice != nil {
		return *m.AdjustPrice
	}
	return false
}

func (m *C2S) GetAdjustSideAndLimit() float64 {
	if m != nil && m.AdjustSideAndLimit != nil {
		return *m.AdjustSideAndLimit
	}
	return 0
}

type S2C struct {
	Header               *Trd_Common.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	OrderID              *uint64               `protobuf:"varint,2,req,name=orderID" json:"orderID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_ModifyOrder_511b294f45632e93, []int{1}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *S2C) GetOrderID() uint64 {
	if m != nil && m.OrderID != nil {
		return *m.OrderID
	}
	return 0
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_ModifyOrder_511b294f45632e93, []int{2}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_ModifyOrder_511b294f45632e93, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Trd_ModifyOrder.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_ModifyOrder.S2C")
	proto.RegisterType((*Request)(nil), "Trd_ModifyOrder.Request")
	proto.RegisterType((*Response)(nil), "Trd_ModifyOrder.Response")
}

func init() {
	proto.RegisterFile("Trd_ModifyOrder/Trd_ModifyOrder.proto", fileDescriptor_Trd_ModifyOrder_511b294f45632e93)
}

var fileDescriptor_Trd_ModifyOrder_511b294f45632e93 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0xab, 0xda, 0x40,
	0x14, 0xc5, 0x99, 0xfc, 0xd1, 0x78, 0xa5, 0x54, 0xa6, 0xb6, 0x0c, 0x16, 0xca, 0x10, 0xda, 0x92,
	0x45, 0x8d, 0x36, 0x74, 0xd5, 0x9d, 0xc4, 0x45, 0x85, 0x5a, 0x65, 0xe2, 0xbe, 0x48, 0x66, 0x6c,
	0xd3, 0x1a, 0x27, 0x4e, 0xc6, 0x85, 0xfb, 0xf7, 0x91, 0xde, 0x07, 0x7c, 0x4c, 0x12, 0x9f, 0x7f,
	0x70, 0xf3, 0x56, 0xb9, 0xe7, 0x9c, 0x5f, 0x2e, 0x87, 0x3b, 0xf0, 0x69, 0xa5, 0xf8, 0xef, 0xb9,
	0xe4, 0xd9, 0xe6, 0xb8, 0x50, 0x5c, 0xa8, 0xd1, 0x8d, 0x0e, 0x0b, 0x25, 0xb5, 0xc4, 0xaf, 0x6f,
	0xec, 0xc1, 0x9b, 0x58, 0xe6, 0xb9, 0xdc, 0x8d, 0xea, 0x4f, 0x4d, 0x0d, 0xde, 0x1b, 0xaa, 0x09,
	0xce, 0x63, 0x1d, 0xfa, 0x8f, 0x16, 0xd8, 0x71, 0x94, 0xe0, 0x2f, 0xe0, 0x15, 0xeb, 0xf4, 0xbf,
	0xd0, 0xb3, 0x29, 0x41, 0xd4, 0x0a, 0xba, 0x51, 0x2f, 0x6c, 0xc0, 0x65, 0xe3, 0xb3, 0x67, 0x02,
	0x0f, 0xa1, 0xf5, 0x57, 0xac, 0xb9, 0x50, 0xc4, 0xaa, 0xd8, 0xb7, 0xe1, 0xc5, 0xe2, 0x95, 0xe2,
	0x3f, 0xaa, 0x90, 0x35, 0x10, 0x26, 0xd0, 0x96, 0xa6, 0xdf, 0x6c, 0x4a, 0x6c, 0x6a, 0x05, 0x0e,
	0x3b, 0x49, 0xfc, 0x11, 0x5e, 0xe5, 0xe7, 0xfe, 0x8b, 0x82, 0x38, 0xd4, 0x0a, 0x5c, 0x76, 0x6d,
	0xe2, 0x77, 0xd0, 0xda, 0x48, 0x35, 0xd9, 0x6e, 0x89, 0x4b, 0x51, 0xe0, 0xb1, 0x46, 0xe1, 0x1e,
	0xd8, 0x7b, 0x7d, 0x24, 0x1e, 0x45, 0x01, 0x62, 0x66, 0xc4, 0x7d, 0x70, 0x0b, 0x95, 0xa5, 0x82,
	0x74, 0x2a, 0xaf, 0x16, 0x98, 0x42, 0x77, 0xcd, 0xff, 0x1d, 0x4a, 0xbd, 0xac, 0x32, 0xa8, 0x96,
	0x5c, 0x5a, 0x38, 0x04, 0x5c, 0xcb, 0x24, 0xe3, 0x62, 0xb2, 0xe3, 0x3f, 0xb3, 0x3c, 0xd3, 0xa4,
	0x5b, 0x2d, 0xb9, 0x93, 0xf8, 0xbf, 0xc0, 0x4e, 0xa2, 0xf8, 0xe2, 0x0e, 0xe8, 0x85, 0x77, 0xb0,
	0xae, 0xee, 0xe0, 0x7f, 0x85, 0x36, 0x13, 0xfb, 0x83, 0x28, 0x35, 0xfe, 0x0c, 0x76, 0x1a, 0x95,
	0xcd, 0xc2, 0x7e, 0x78, 0xfb, 0xf2, 0x71, 0x94, 0x30, 0x03, 0xf8, 0x0f, 0x08, 0x3c, 0x26, 0xca,
	0x42, 0xee, 0x4a, 0x81, 0x3f, 0x40, 0x5b, 0x09, 0xbd, 0x3a, 0x16, 0xa2, 0xfa, 0xd1, 0xfd, 0xee,
	0x0c, 0xbf, 0x8d, 0xc7, 0xec, 0x64, 0x9a, 0x0b, 0x2a, 0xa1, 0xe7, 0xe5, 0x1f, 0x62, 0x51, 0x14,
	0x74, 0x58, 0xa3, 0x4c, 0x23, 0xa1, 0x54, 0x2c, 0xb9, 0x20, 0x36, 0x45, 0x81, 0xcb, 0x4e, 0xd2,
	0xd4, 0x28, 0xa3, 0x94, 0x38, 0x14, 0xdd, 0xad, 0x91, 0x44, 0x31, 0x33, 0xc0, 0x53, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xad, 0x25, 0x20, 0x5e, 0xab, 0x02, 0x00, 0x00,
}
