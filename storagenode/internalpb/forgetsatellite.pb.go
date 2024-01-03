// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: forgetsatellite.proto

package internalpb

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type InitForgetSatelliteRequest struct {
	// satellite_id is the satellite to forget.
	SatelliteId          NodeID   `protobuf:"bytes,1,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	ForceCleanup         bool     `protobuf:"varint,2,opt,name=force_cleanup,json=forceCleanup,proto3" json:"force_cleanup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitForgetSatelliteRequest) Reset()         { *m = InitForgetSatelliteRequest{} }
func (m *InitForgetSatelliteRequest) String() string { return proto.CompactTextString(m) }
func (*InitForgetSatelliteRequest) ProtoMessage()    {}
func (*InitForgetSatelliteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ee03a9da2a32ad8, []int{0}
}
func (m *InitForgetSatelliteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitForgetSatelliteRequest.Unmarshal(m, b)
}
func (m *InitForgetSatelliteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitForgetSatelliteRequest.Marshal(b, m, deterministic)
}
func (m *InitForgetSatelliteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitForgetSatelliteRequest.Merge(m, src)
}
func (m *InitForgetSatelliteRequest) XXX_Size() int {
	return xxx_messageInfo_InitForgetSatelliteRequest.Size(m)
}
func (m *InitForgetSatelliteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitForgetSatelliteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitForgetSatelliteRequest proto.InternalMessageInfo

func (m *InitForgetSatelliteRequest) GetForceCleanup() bool {
	if m != nil {
		return m.ForceCleanup
	}
	return false
}

type InitForgetSatelliteResponse struct {
	SatelliteId          NodeID   `protobuf:"bytes,1,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	InProgress           bool     `protobuf:"varint,2,opt,name=in_progress,json=inProgress,proto3" json:"in_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitForgetSatelliteResponse) Reset()         { *m = InitForgetSatelliteResponse{} }
func (m *InitForgetSatelliteResponse) String() string { return proto.CompactTextString(m) }
func (*InitForgetSatelliteResponse) ProtoMessage()    {}
func (*InitForgetSatelliteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ee03a9da2a32ad8, []int{1}
}
func (m *InitForgetSatelliteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitForgetSatelliteResponse.Unmarshal(m, b)
}
func (m *InitForgetSatelliteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitForgetSatelliteResponse.Marshal(b, m, deterministic)
}
func (m *InitForgetSatelliteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitForgetSatelliteResponse.Merge(m, src)
}
func (m *InitForgetSatelliteResponse) XXX_Size() int {
	return xxx_messageInfo_InitForgetSatelliteResponse.Size(m)
}
func (m *InitForgetSatelliteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitForgetSatelliteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitForgetSatelliteResponse proto.InternalMessageInfo

func (m *InitForgetSatelliteResponse) GetInProgress() bool {
	if m != nil {
		return m.InProgress
	}
	return false
}

type GetUntrustedSatellitesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUntrustedSatellitesRequest) Reset()         { *m = GetUntrustedSatellitesRequest{} }
func (m *GetUntrustedSatellitesRequest) String() string { return proto.CompactTextString(m) }
func (*GetUntrustedSatellitesRequest) ProtoMessage()    {}
func (*GetUntrustedSatellitesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ee03a9da2a32ad8, []int{2}
}
func (m *GetUntrustedSatellitesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUntrustedSatellitesRequest.Unmarshal(m, b)
}
func (m *GetUntrustedSatellitesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUntrustedSatellitesRequest.Marshal(b, m, deterministic)
}
func (m *GetUntrustedSatellitesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUntrustedSatellitesRequest.Merge(m, src)
}
func (m *GetUntrustedSatellitesRequest) XXX_Size() int {
	return xxx_messageInfo_GetUntrustedSatellitesRequest.Size(m)
}
func (m *GetUntrustedSatellitesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUntrustedSatellitesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUntrustedSatellitesRequest proto.InternalMessageInfo

type GetUntrustedSatellitesResponse struct {
	// satellite_ids is the list of satellite ids that are not trusted.
	SatelliteIds         []NodeID `protobuf:"bytes,1,rep,name=satellite_ids,json=satelliteIds,proto3,customtype=NodeID" json:"satellite_ids"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUntrustedSatellitesResponse) Reset()         { *m = GetUntrustedSatellitesResponse{} }
func (m *GetUntrustedSatellitesResponse) String() string { return proto.CompactTextString(m) }
func (*GetUntrustedSatellitesResponse) ProtoMessage()    {}
func (*GetUntrustedSatellitesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ee03a9da2a32ad8, []int{3}
}
func (m *GetUntrustedSatellitesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUntrustedSatellitesResponse.Unmarshal(m, b)
}
func (m *GetUntrustedSatellitesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUntrustedSatellitesResponse.Marshal(b, m, deterministic)
}
func (m *GetUntrustedSatellitesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUntrustedSatellitesResponse.Merge(m, src)
}
func (m *GetUntrustedSatellitesResponse) XXX_Size() int {
	return xxx_messageInfo_GetUntrustedSatellitesResponse.Size(m)
}
func (m *GetUntrustedSatellitesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUntrustedSatellitesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUntrustedSatellitesResponse proto.InternalMessageInfo

type ForgetSatelliteStatusRequest struct {
	// satellite_id is the satellite to forget.
	SatelliteId          NodeID   `protobuf:"bytes,1,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgetSatelliteStatusRequest) Reset()         { *m = ForgetSatelliteStatusRequest{} }
func (m *ForgetSatelliteStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ForgetSatelliteStatusRequest) ProtoMessage()    {}
func (*ForgetSatelliteStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ee03a9da2a32ad8, []int{4}
}
func (m *ForgetSatelliteStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgetSatelliteStatusRequest.Unmarshal(m, b)
}
func (m *ForgetSatelliteStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgetSatelliteStatusRequest.Marshal(b, m, deterministic)
}
func (m *ForgetSatelliteStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgetSatelliteStatusRequest.Merge(m, src)
}
func (m *ForgetSatelliteStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ForgetSatelliteStatusRequest.Size(m)
}
func (m *ForgetSatelliteStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgetSatelliteStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForgetSatelliteStatusRequest proto.InternalMessageInfo

type ForgetSatelliteStatusResponse struct {
	// satellite_id is the satellite to forget.
	SatelliteId NodeID `protobuf:"bytes,1,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	// in_progress is true if the forget satellite operation is in progress.
	InProgress           bool     `protobuf:"varint,2,opt,name=in_progress,json=inProgress,proto3" json:"in_progress,omitempty"`
	Successful           bool     `protobuf:"varint,3,opt,name=successful,proto3" json:"successful,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForgetSatelliteStatusResponse) Reset()         { *m = ForgetSatelliteStatusResponse{} }
func (m *ForgetSatelliteStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ForgetSatelliteStatusResponse) ProtoMessage()    {}
func (*ForgetSatelliteStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ee03a9da2a32ad8, []int{5}
}
func (m *ForgetSatelliteStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForgetSatelliteStatusResponse.Unmarshal(m, b)
}
func (m *ForgetSatelliteStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForgetSatelliteStatusResponse.Marshal(b, m, deterministic)
}
func (m *ForgetSatelliteStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForgetSatelliteStatusResponse.Merge(m, src)
}
func (m *ForgetSatelliteStatusResponse) XXX_Size() int {
	return xxx_messageInfo_ForgetSatelliteStatusResponse.Size(m)
}
func (m *ForgetSatelliteStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ForgetSatelliteStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ForgetSatelliteStatusResponse proto.InternalMessageInfo

func (m *ForgetSatelliteStatusResponse) GetInProgress() bool {
	if m != nil {
		return m.InProgress
	}
	return false
}

func (m *ForgetSatelliteStatusResponse) GetSuccessful() bool {
	if m != nil {
		return m.Successful
	}
	return false
}

func init() {
	proto.RegisterType((*InitForgetSatelliteRequest)(nil), "storagenode.forgetsatellite.InitForgetSatelliteRequest")
	proto.RegisterType((*InitForgetSatelliteResponse)(nil), "storagenode.forgetsatellite.InitForgetSatelliteResponse")
	proto.RegisterType((*GetUntrustedSatellitesRequest)(nil), "storagenode.forgetsatellite.GetUntrustedSatellitesRequest")
	proto.RegisterType((*GetUntrustedSatellitesResponse)(nil), "storagenode.forgetsatellite.GetUntrustedSatellitesResponse")
	proto.RegisterType((*ForgetSatelliteStatusRequest)(nil), "storagenode.forgetsatellite.ForgetSatelliteStatusRequest")
	proto.RegisterType((*ForgetSatelliteStatusResponse)(nil), "storagenode.forgetsatellite.ForgetSatelliteStatusResponse")
}

func init() { proto.RegisterFile("forgetsatellite.proto", fileDescriptor_5ee03a9da2a32ad8) }

var fileDescriptor_5ee03a9da2a32ad8 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6a, 0xe2, 0x50,
	0x14, 0x9d, 0x37, 0x82, 0xcc, 0x5c, 0x75, 0x16, 0x4f, 0x1c, 0x42, 0x1c, 0x8d, 0x64, 0x18, 0xc6,
	0x55, 0xa4, 0x75, 0xd1, 0xd6, 0xee, 0x6c, 0x69, 0x71, 0x53, 0x5a, 0xc5, 0x4d, 0x37, 0x12, 0x93,
	0x6b, 0x48, 0x09, 0xef, 0xc5, 0xf7, 0x5e, 0xfe, 0xa1, 0xab, 0x42, 0xe9, 0x37, 0x15, 0xfa, 0x0d,
	0x5d, 0xf8, 0x2d, 0xc5, 0x24, 0xb5, 0x22, 0x31, 0xd0, 0x40, 0x77, 0xe1, 0xe6, 0x9c, 0x73, 0xcf,
	0xbd, 0xf7, 0xf0, 0xa0, 0xb1, 0xe0, 0xc2, 0x43, 0x25, 0x6d, 0x85, 0x41, 0xe0, 0x2b, 0xb4, 0x42,
	0xc1, 0x15, 0xa7, 0x4d, 0xa9, 0xb8, 0xb0, 0x3d, 0x64, 0xdc, 0x45, 0x6b, 0x07, 0xa2, 0x83, 0xc7,
	0x3d, 0x9e, 0x00, 0x4d, 0x05, 0xfa, 0x88, 0xf9, 0xea, 0x22, 0x86, 0x4c, 0xde, 0x21, 0x63, 0x5c,
	0x46, 0x28, 0x15, 0x3d, 0x80, 0xea, 0x86, 0x36, 0xf3, 0x5d, 0x8d, 0x74, 0x48, 0xf7, 0xe7, 0xf0,
	0xd7, 0xcb, 0xca, 0xf8, 0xf6, 0xba, 0x32, 0xca, 0x57, 0xdc, 0xc5, 0xd1, 0xf9, 0xb8, 0xb2, 0xc1,
	0x8c, 0x5c, 0xfa, 0x17, 0x6a, 0x0b, 0x2e, 0x1c, 0x9c, 0x39, 0x01, 0xda, 0x2c, 0x0a, 0xb5, 0xef,
	0x1d, 0xd2, 0xfd, 0x31, 0xae, 0xc6, 0xc5, 0xb3, 0xa4, 0x66, 0x2e, 0xa1, 0x99, 0xd9, 0x55, 0x86,
	0x9c, 0x49, 0x2c, 0xd2, 0xd6, 0x80, 0x8a, 0xcf, 0x66, 0xa1, 0xe0, 0x9e, 0x40, 0x29, 0xd3, 0xa6,
	0xe0, 0xb3, 0xeb, 0xb4, 0x62, 0x1a, 0xd0, 0xba, 0x44, 0x35, 0x65, 0x4a, 0x44, 0x52, 0xa1, 0xbb,
	0x69, 0x2a, 0xd3, 0x59, 0xcd, 0x29, 0xb4, 0xf7, 0x01, 0x52, 0x5b, 0x7d, 0xa8, 0x6d, 0xdb, 0x92,
	0x1a, 0xe9, 0x94, 0x32, 0x7c, 0x55, 0xb7, 0x7c, 0x49, 0xf3, 0x06, 0xfe, 0xec, 0x8c, 0x39, 0x51,
	0xb6, 0x8a, 0x64, 0xf1, 0x15, 0x9b, 0x4f, 0x04, 0x5a, 0x7b, 0x34, 0xbf, 0x6e, 0x81, 0xb4, 0x0d,
	0x20, 0x23, 0xc7, 0x41, 0x29, 0x17, 0x51, 0xa0, 0x95, 0x92, 0xff, 0x1f, 0x95, 0xc3, 0xe7, 0x12,
	0xd4, 0xd7, 0xc2, 0x3b, 0xce, 0xe8, 0x3d, 0x81, 0x7a, 0xc6, 0xb1, 0xe9, 0x91, 0x95, 0x93, 0x51,
	0x6b, 0x7f, 0x28, 0xf5, 0xe3, 0xcf, 0x13, 0xd3, 0xb5, 0x3c, 0x12, 0xf8, 0x9d, 0x7d, 0x63, 0x3a,
	0xc8, 0x15, 0xcd, 0x4d, 0x8e, 0x7e, 0x5a, 0x88, 0x9b, 0x7a, 0x7a, 0x20, 0xd0, 0xc8, 0x3c, 0x26,
	0x3d, 0xc9, 0x95, 0xcd, 0x0b, 0x95, 0x3e, 0x28, 0x42, 0x4d, 0x0c, 0x0d, 0xff, 0xdf, 0xfe, 0x5b,
	0x93, 0xef, 0x2c, 0x9f, 0xf7, 0xe2, 0x8f, 0xde, 0x96, 0x56, 0xcf, 0x67, 0x0a, 0x05, 0xb3, 0x83,
	0x70, 0x3e, 0x2f, 0xc7, 0x2f, 0x48, 0xff, 0x2d, 0x00, 0x00, 0xff, 0xff, 0xea, 0x72, 0x60, 0x50,
	0x83, 0x04, 0x00, 0x00,
}
