// Code generated by protoc-gen-go.
// source: bi.proto
// DO NOT EDIT!

/*
Package p is a generated protocol buffer package.

It is generated from these files:
	bi.proto

It has these top-level messages:
	BiLog
	BiResult
	PerformPathIssue
	DeviceInfo
	CommandResult
	CommandExecutePerformance
*/
package p

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type BiLog struct {
	ProjectName string `protobuf:"bytes,1,opt,name=projectName" json:"projectName,omitempty"`
	ActionName  string `protobuf:"bytes,2,opt,name=actionName" json:"actionName,omitempty"`
	Timestamp   int64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Detail      []byte `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (m *BiLog) Reset()                    { *m = BiLog{} }
func (m *BiLog) String() string            { return proto.CompactTextString(m) }
func (*BiLog) ProtoMessage()               {}
func (*BiLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BiResult struct {
	Result bool   `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *BiResult) Reset()                    { *m = BiResult{} }
func (m *BiResult) String() string            { return proto.CompactTextString(m) }
func (*BiResult) ProtoMessage()               {}
func (*BiResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PerformPathIssue struct {
	QueryId             string `protobuf:"bytes,1,opt,name=query_id,json=queryId" json:"query_id,omitempty"`
	CanPerformPathCount int32  `protobuf:"varint,2,opt,name=can_perform_path_count,json=canPerformPathCount" json:"can_perform_path_count,omitempty"`
}

func (m *PerformPathIssue) Reset()                    { *m = PerformPathIssue{} }
func (m *PerformPathIssue) String() string            { return proto.CompactTextString(m) }
func (*PerformPathIssue) ProtoMessage()               {}
func (*PerformPathIssue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DeviceInfo struct {
	DeviceId     string `protobuf:"bytes,1,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Manufacturer string `protobuf:"bytes,2,opt,name=manufacturer" json:"manufacturer,omitempty"`
	Model        string `protobuf:"bytes,3,opt,name=model" json:"model,omitempty"`
	AppVersion   string `protobuf:"bytes,4,opt,name=app_version,json=appVersion" json:"app_version,omitempty"`
	NiVersion    string `protobuf:"bytes,5,opt,name=ni_version,json=niVersion" json:"ni_version,omitempty"`
	OpenAppTime  int64  `protobuf:"varint,6,opt,name=open_app_time,json=openAppTime" json:"open_app_time,omitempty"`
}

func (m *DeviceInfo) Reset()                    { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string            { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()               {}
func (*DeviceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CommandResult struct {
	CommandName       string `protobuf:"bytes,1,opt,name=command_name,json=commandName" json:"command_name,omitempty"`
	QueryId           string `protobuf:"bytes,2,opt,name=query_id,json=queryId" json:"query_id,omitempty"`
	UseSearchResult   bool   `protobuf:"varint,3,opt,name=use_search_result,json=useSearchResult" json:"use_search_result,omitempty"`
	UserChooseTrigger bool   `protobuf:"varint,4,opt,name=user_choose_trigger,json=userChooseTrigger" json:"user_choose_trigger,omitempty"`
	DateTime          int64  `protobuf:"varint,5,opt,name=date_time,json=dateTime" json:"date_time,omitempty"`
}

func (m *CommandResult) Reset()                    { *m = CommandResult{} }
func (m *CommandResult) String() string            { return proto.CompactTextString(m) }
func (*CommandResult) ProtoMessage()               {}
func (*CommandResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type CommandExecutePerformance struct {
	SpeakTime              int64  `protobuf:"varint,1,opt,name=speak_time,json=speakTime" json:"speak_time,omitempty"`
	ReceiveVoiceResultTime int64  `protobuf:"varint,2,opt,name=receive_voice_result_time,json=receiveVoiceResultTime" json:"receive_voice_result_time,omitempty"`
	SendCommandTime        int64  `protobuf:"varint,3,opt,name=send_command_time,json=sendCommandTime" json:"send_command_time,omitempty"`
	ReceiveResultTime      int64  `protobuf:"varint,4,opt,name=receive_result_time,json=receiveResultTime" json:"receive_result_time,omitempty"`
	CommandText            string `protobuf:"bytes,5,opt,name=command_text,json=commandText" json:"command_text,omitempty"`
	QueryId                string `protobuf:"bytes,6,opt,name=query_id,json=queryId" json:"query_id,omitempty"`
	DeviceId               string `protobuf:"bytes,7,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
}

func (m *CommandExecutePerformance) Reset()                    { *m = CommandExecutePerformance{} }
func (m *CommandExecutePerformance) String() string            { return proto.CompactTextString(m) }
func (*CommandExecutePerformance) ProtoMessage()               {}
func (*CommandExecutePerformance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*BiLog)(nil), "p.BiLog")
	proto.RegisterType((*BiResult)(nil), "p.BiResult")
	proto.RegisterType((*PerformPathIssue)(nil), "p.PerformPathIssue")
	proto.RegisterType((*DeviceInfo)(nil), "p.DeviceInfo")
	proto.RegisterType((*CommandResult)(nil), "p.CommandResult")
	proto.RegisterType((*CommandExecutePerformance)(nil), "p.CommandExecutePerformance")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for MisBi service

type MisBiClient interface {
	Bi(ctx context.Context, in *BiLog, opts ...grpc.CallOption) (*BiResult, error)
	BiDeviceInfo(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*BiResult, error)
}

type misBiClient struct {
	cc *grpc.ClientConn
}

func NewMisBiClient(cc *grpc.ClientConn) MisBiClient {
	return &misBiClient{cc}
}

func (c *misBiClient) Bi(ctx context.Context, in *BiLog, opts ...grpc.CallOption) (*BiResult, error) {
	out := new(BiResult)
	err := grpc.Invoke(ctx, "/p.MisBi/Bi", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *misBiClient) BiDeviceInfo(ctx context.Context, in *DeviceInfo, opts ...grpc.CallOption) (*BiResult, error) {
	out := new(BiResult)
	err := grpc.Invoke(ctx, "/p.MisBi/BiDeviceInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MisBi service

type MisBiServer interface {
	Bi(context.Context, *BiLog) (*BiResult, error)
	BiDeviceInfo(context.Context, *DeviceInfo) (*BiResult, error)
}

func RegisterMisBiServer(s *grpc.Server, srv MisBiServer) {
	s.RegisterService(&_MisBi_serviceDesc, srv)
}

func _MisBi_Bi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BiLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MisBiServer).Bi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p.MisBi/Bi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MisBiServer).Bi(ctx, req.(*BiLog))
	}
	return interceptor(ctx, in, info, handler)
}

func _MisBi_BiDeviceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MisBiServer).BiDeviceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/p.MisBi/BiDeviceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MisBiServer).BiDeviceInfo(ctx, req.(*DeviceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _MisBi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "p.MisBi",
	HandlerType: (*MisBiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bi",
			Handler:    _MisBi_Bi_Handler,
		},
		{
			MethodName: "BiDeviceInfo",
			Handler:    _MisBi_BiDeviceInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x54, 0xcd, 0x72, 0xd3, 0x3c,
	0x14, 0xfd, 0x9c, 0x34, 0xa9, 0x73, 0x93, 0x4e, 0x5b, 0xf5, 0x9b, 0x4e, 0x5a, 0xfe, 0x8a, 0x37,
	0x30, 0x0c, 0xe3, 0x05, 0x65, 0xc3, 0x12, 0x17, 0x16, 0x9d, 0x01, 0x26, 0x63, 0x32, 0xdd, 0x7a,
	0x5c, 0xe7, 0x36, 0x15, 0xc4, 0x96, 0x90, 0xe4, 0x4c, 0x59, 0xf1, 0x42, 0x3c, 0x07, 0xbc, 0x16,
	0xd2, 0x95, 0x42, 0x1d, 0x36, 0x19, 0xe9, 0x9c, 0xfb, 0x7b, 0x8e, 0x62, 0x88, 0xaf, 0x79, 0x2a,
	0x95, 0x30, 0x82, 0x45, 0x32, 0xf9, 0x01, 0x83, 0x8c, 0x7f, 0x10, 0x4b, 0x76, 0x06, 0x63, 0x0b,
	0x7e, 0xc1, 0xca, 0x7c, 0x2a, 0x6b, 0x9c, 0x46, 0x67, 0xd1, 0xf3, 0x51, 0xde, 0x85, 0xd8, 0x63,
	0x80, 0xb2, 0x32, 0x5c, 0x34, 0x14, 0xd0, 0xa3, 0x80, 0x0e, 0xc2, 0x1e, 0xc2, 0xc8, 0xf0, 0x1a,
	0xb5, 0x29, 0x6b, 0x39, 0xed, 0x5b, 0xba, 0x9f, 0xdf, 0x03, 0xec, 0x18, 0x86, 0x0b, 0x34, 0x25,
	0x5f, 0x4d, 0x77, 0x2c, 0x35, 0xc9, 0xc3, 0x2d, 0x79, 0x0d, 0x71, 0xc6, 0x73, 0xd4, 0xed, 0xca,
	0xb8, 0x18, 0x45, 0x27, 0x6a, 0x1f, 0xe7, 0xe1, 0xc6, 0x0e, 0xa0, 0x5f, 0xeb, 0x65, 0x68, 0xe9,
	0x8e, 0xc9, 0x35, 0x1c, 0xcc, 0x50, 0xdd, 0x08, 0x55, 0xcf, 0x4a, 0x73, 0x7b, 0xa9, 0x75, 0x8b,
	0xec, 0x04, 0xe2, 0x6f, 0x2d, 0xaa, 0xef, 0x05, 0x5f, 0x84, 0xf1, 0x77, 0xe9, 0x7e, 0xb9, 0x60,
	0xe7, 0x70, 0x5c, 0x95, 0x4d, 0x21, 0x7d, 0x4a, 0x21, 0x6d, 0x4e, 0x51, 0x89, 0xb6, 0x31, 0x54,
	0x73, 0x90, 0x1f, 0x59, 0xb6, 0x53, 0xef, 0xc2, 0x51, 0xc9, 0xef, 0x08, 0xe0, 0x1d, 0xae, 0x79,
	0x85, 0x97, 0xcd, 0x8d, 0x60, 0x0f, 0x60, 0xb4, 0xa0, 0xdb, 0x7d, 0xfd, 0xd8, 0x03, 0xb6, 0x41,
	0x02, 0x93, 0xba, 0x6c, 0xda, 0x1b, 0x2b, 0x47, 0xab, 0x50, 0x85, 0x51, 0xb7, 0x30, 0xf6, 0x3f,
	0x0c, 0x6a, 0xb1, 0xc0, 0x15, 0x69, 0x33, 0xca, 0xfd, 0x85, 0x3d, 0x81, 0x71, 0x29, 0x65, 0xb1,
	0x46, 0xa5, 0xad, 0x90, 0x24, 0x8e, 0x93, 0x55, 0xca, 0x2b, 0x8f, 0xb0, 0x47, 0x00, 0x0d, 0xff,
	0xcb, 0x0f, 0x88, 0x1f, 0x35, 0x7c, 0x43, 0x27, 0xb0, 0x27, 0x24, 0x36, 0x85, 0x2b, 0xe2, 0xd4,
	0x9e, 0x0e, 0x49, 0xf9, 0xb1, 0x03, 0xdf, 0x4a, 0x39, 0xb7, 0x50, 0xf2, 0x2b, 0x82, 0xbd, 0x0b,
	0x51, 0xdb, 0x69, 0x16, 0x41, 0xe9, 0xa7, 0x30, 0xa9, 0x3c, 0x50, 0x34, 0x1d, 0xbb, 0x03, 0x46,
	0x76, 0x76, 0xe5, 0xec, 0x6d, 0xcb, 0xf9, 0x02, 0x0e, 0x5b, 0x8d, 0x85, 0xc6, 0x52, 0x55, 0xb7,
	0x45, 0xb0, 0xac, 0x4f, 0x96, 0xed, 0x5b, 0xe2, 0x33, 0xe1, 0xa1, 0x53, 0x0a, 0x47, 0x16, 0x52,
	0x45, 0x75, 0x2b, 0x84, 0xcd, 0x31, 0x8a, 0x2f, 0x97, 0x56, 0xa0, 0x1d, 0x8a, 0x76, 0x65, 0xd4,
	0x05, 0x31, 0x73, 0x4f, 0x90, 0xcc, 0xa5, 0x41, 0xbf, 0xcb, 0x80, 0x76, 0x89, 0x1d, 0x40, 0x8b,
	0xfc, 0xec, 0xc1, 0x49, 0x58, 0xe4, 0xfd, 0x1d, 0x56, 0xad, 0xc1, 0xe0, 0x5a, 0xd9, 0x54, 0xe8,
	0x94, 0xd2, 0x12, 0xcb, 0xaf, 0x3e, 0x37, 0xf2, 0x2f, 0x90, 0x10, 0x97, 0xcc, 0xde, 0xc0, 0x89,
	0xc2, 0x0a, 0xf9, 0x1a, 0x8b, 0xb5, 0x70, 0x3e, 0xfa, 0xc1, 0x7d, 0x74, 0x8f, 0xa2, 0x8f, 0x43,
	0xc0, 0x95, 0xe3, 0xfd, 0x02, 0x94, 0x6a, 0x17, 0xd6, 0x68, 0xb5, 0xda, 0x68, 0x46, 0x29, 0xfe,
	0x89, 0xef, 0x3b, 0x22, 0xcc, 0x44, 0xb1, 0x76, 0xe1, 0x4d, 0x9b, 0x6e, 0x83, 0x1d, 0x8a, 0x3e,
	0x0c, 0x54, 0xa7, 0x76, 0xc7, 0x0a, 0x83, 0x77, 0x26, 0x38, 0xbc, 0xb1, 0x62, 0x6e, 0xa1, 0x2d,
	0x2b, 0x86, 0xdb, 0x56, 0x6c, 0xbd, 0xca, 0xdd, 0xed, 0x57, 0xf9, 0x6a, 0x0e, 0x83, 0x8f, 0x5c,
	0x67, 0xdc, 0x2a, 0xd3, 0xb3, 0xbf, 0x71, 0x2a, 0x53, 0xfa, 0xb3, 0x9f, 0x8e, 0xe9, 0xe4, 0x87,
	0x48, 0xfe, 0x63, 0x2f, 0x61, 0x92, 0xf1, 0xce, 0x53, 0xdf, 0xb3, 0xf4, 0xfd, 0xf5, 0x9f, 0xe8,
	0xec, 0x19, 0x9c, 0xda, 0xe1, 0x52, 0xcd, 0x9b, 0x65, 0xbb, 0x2a, 0x95, 0x14, 0xbc, 0x31, 0x3a,
	0xe5, 0x22, 0x5d, 0x2a, 0x59, 0x65, 0xbb, 0x19, 0x9f, 0xb9, 0x8f, 0xcb, 0x2c, 0xba, 0x1e, 0xd2,
	0x57, 0xe6, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x3d, 0x1d, 0xf0, 0x71, 0x04, 0x00,
	0x00,
}
