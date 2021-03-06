/*
 * Copyright 2019 The FATE Authors. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MetricType int32

const (
	MetricType_INTERFACE MetricType = 0
	MetricType_MODEL     MetricType = 1
)

var MetricType_name = map[int32]string{
	0: "INTERFACE",
	1: "MODEL",
}

var MetricType_value = map[string]int32{
	"INTERFACE": 0,
	"MODEL":     1,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}

func (MetricType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b038c4ef64de8386, []int{0}
}

type QueryMetricRequest struct {
	BeginMs              int64      `protobuf:"varint,1,opt,name=beginMs,proto3" json:"beginMs,omitempty"`
	EndMs                int64      `protobuf:"varint,2,opt,name=endMs,proto3" json:"endMs,omitempty"`
	Source               string     `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Type                 MetricType `protobuf:"varint,4,opt,name=type,proto3,enum=com.webank.ai.fate.api.networking.common.MetricType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueryMetricRequest) Reset()         { *m = QueryMetricRequest{} }
func (m *QueryMetricRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMetricRequest) ProtoMessage()    {}
func (*QueryMetricRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b038c4ef64de8386, []int{0}
}

func (m *QueryMetricRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryMetricRequest.Unmarshal(m, b)
}
func (m *QueryMetricRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryMetricRequest.Marshal(b, m, deterministic)
}
func (m *QueryMetricRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMetricRequest.Merge(m, src)
}
func (m *QueryMetricRequest) XXX_Size() int {
	return xxx_messageInfo_QueryMetricRequest.Size(m)
}
func (m *QueryMetricRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMetricRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMetricRequest proto.InternalMessageInfo

func (m *QueryMetricRequest) GetBeginMs() int64 {
	if m != nil {
		return m.BeginMs
	}
	return 0
}

func (m *QueryMetricRequest) GetEndMs() int64 {
	if m != nil {
		return m.EndMs
	}
	return 0
}

func (m *QueryMetricRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *QueryMetricRequest) GetType() MetricType {
	if m != nil {
		return m.Type
	}
	return MetricType_INTERFACE
}

type QueryMetricResponse struct {
	Metrics              []byte   `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryMetricResponse) Reset()         { *m = QueryMetricResponse{} }
func (m *QueryMetricResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMetricResponse) ProtoMessage()    {}
func (*QueryMetricResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b038c4ef64de8386, []int{1}
}

func (m *QueryMetricResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryMetricResponse.Unmarshal(m, b)
}
func (m *QueryMetricResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryMetricResponse.Marshal(b, m, deterministic)
}
func (m *QueryMetricResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMetricResponse.Merge(m, src)
}
func (m *QueryMetricResponse) XXX_Size() int {
	return xxx_messageInfo_QueryMetricResponse.Size(m)
}
func (m *QueryMetricResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMetricResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMetricResponse proto.InternalMessageInfo

func (m *QueryMetricResponse) GetMetrics() []byte {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type UpdateFlowRuleRequest struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	AllowQps             float64  `protobuf:"fixed64,2,opt,name=allowQps,proto3" json:"allowQps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFlowRuleRequest) Reset()         { *m = UpdateFlowRuleRequest{} }
func (m *UpdateFlowRuleRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFlowRuleRequest) ProtoMessage()    {}
func (*UpdateFlowRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b038c4ef64de8386, []int{2}
}

func (m *UpdateFlowRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFlowRuleRequest.Unmarshal(m, b)
}
func (m *UpdateFlowRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFlowRuleRequest.Marshal(b, m, deterministic)
}
func (m *UpdateFlowRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFlowRuleRequest.Merge(m, src)
}
func (m *UpdateFlowRuleRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFlowRuleRequest.Size(m)
}
func (m *UpdateFlowRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFlowRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFlowRuleRequest proto.InternalMessageInfo

func (m *UpdateFlowRuleRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *UpdateFlowRuleRequest) GetAllowQps() float64 {
	if m != nil {
		return m.AllowQps
	}
	return 0
}

type QueryPropsRequest struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryPropsRequest) Reset()         { *m = QueryPropsRequest{} }
func (m *QueryPropsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPropsRequest) ProtoMessage()    {}
func (*QueryPropsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b038c4ef64de8386, []int{3}
}

func (m *QueryPropsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPropsRequest.Unmarshal(m, b)
}
func (m *QueryPropsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPropsRequest.Marshal(b, m, deterministic)
}
func (m *QueryPropsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPropsRequest.Merge(m, src)
}
func (m *QueryPropsRequest) XXX_Size() int {
	return xxx_messageInfo_QueryPropsRequest.Size(m)
}
func (m *QueryPropsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPropsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPropsRequest proto.InternalMessageInfo

func (m *QueryPropsRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type QueryJvmInfoRequest struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryJvmInfoRequest) Reset()         { *m = QueryJvmInfoRequest{} }
func (m *QueryJvmInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryJvmInfoRequest) ProtoMessage()    {}
func (*QueryJvmInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b038c4ef64de8386, []int{4}
}

func (m *QueryJvmInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryJvmInfoRequest.Unmarshal(m, b)
}
func (m *QueryJvmInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryJvmInfoRequest.Marshal(b, m, deterministic)
}
func (m *QueryJvmInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryJvmInfoRequest.Merge(m, src)
}
func (m *QueryJvmInfoRequest) XXX_Size() int {
	return xxx_messageInfo_QueryJvmInfoRequest.Size(m)
}
func (m *QueryJvmInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryJvmInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryJvmInfoRequest proto.InternalMessageInfo

func (m *QueryJvmInfoRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type UpdateServiceRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	RouterMode           string   `protobuf:"bytes,2,opt,name=routerMode,proto3" json:"routerMode,omitempty"`
	Weight               int32    `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Version              int64    `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateServiceRequest) Reset()         { *m = UpdateServiceRequest{} }
func (m *UpdateServiceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateServiceRequest) ProtoMessage()    {}
func (*UpdateServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b038c4ef64de8386, []int{5}
}

func (m *UpdateServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateServiceRequest.Unmarshal(m, b)
}
func (m *UpdateServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateServiceRequest.Marshal(b, m, deterministic)
}
func (m *UpdateServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServiceRequest.Merge(m, src)
}
func (m *UpdateServiceRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateServiceRequest.Size(m)
}
func (m *UpdateServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServiceRequest proto.InternalMessageInfo

func (m *UpdateServiceRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *UpdateServiceRequest) GetRouterMode() string {
	if m != nil {
		return m.RouterMode
	}
	return ""
}

func (m *UpdateServiceRequest) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *UpdateServiceRequest) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type CommonResponse struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b038c4ef64de8386, []int{6}
}

func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *CommonResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CommonResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("com.webank.ai.fate.api.networking.common.MetricType", MetricType_name, MetricType_value)
	proto.RegisterType((*QueryMetricRequest)(nil), "com.webank.ai.fate.api.networking.common.QueryMetricRequest")
	proto.RegisterType((*QueryMetricResponse)(nil), "com.webank.ai.fate.api.networking.common.QueryMetricResponse")
	proto.RegisterType((*UpdateFlowRuleRequest)(nil), "com.webank.ai.fate.api.networking.common.UpdateFlowRuleRequest")
	proto.RegisterType((*QueryPropsRequest)(nil), "com.webank.ai.fate.api.networking.common.QueryPropsRequest")
	proto.RegisterType((*QueryJvmInfoRequest)(nil), "com.webank.ai.fate.api.networking.common.QueryJvmInfoRequest")
	proto.RegisterType((*UpdateServiceRequest)(nil), "com.webank.ai.fate.api.networking.common.UpdateServiceRequest")
	proto.RegisterType((*CommonResponse)(nil), "com.webank.ai.fate.api.networking.common.CommonResponse")
}

func init() { proto.RegisterFile("common-service.proto", fileDescriptor_b038c4ef64de8386) }

var fileDescriptor_b038c4ef64de8386 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x59, 0x92, 0x14, 0x32, 0x4a, 0xa2, 0xb0, 0x04, 0x14, 0xf5, 0x50, 0x45, 0x39, 0xa0,
	0x08, 0xa9, 0x46, 0x2a, 0x1c, 0x90, 0xf8, 0x27, 0x1a, 0x52, 0x51, 0x20, 0xd0, 0x2e, 0xe5, 0x8a,
	0xb4, 0xb1, 0xa7, 0xc1, 0x8a, 0xed, 0x75, 0x77, 0xd7, 0xb1, 0xc2, 0x91, 0x03, 0x08, 0xf1, 0x1e,
	0xbc, 0x07, 0x6f, 0x86, 0xbc, 0xb6, 0x13, 0x1b, 0x90, 0x48, 0x9a, 0x9b, 0xbf, 0x91, 0x67, 0xe6,
	0xb7, 0x33, 0xdf, 0x2e, 0x74, 0x6c, 0xe1, 0xfb, 0x22, 0xd8, 0x57, 0x28, 0xe7, 0xae, 0x8d, 0x56,
	0x28, 0x85, 0x16, 0x74, 0x60, 0x0b, 0xdf, 0x8a, 0x71, 0xc2, 0x83, 0x99, 0xc5, 0x5d, 0xeb, 0x9c,
	0x6b, 0xb4, 0x78, 0xe8, 0x5a, 0x01, 0xea, 0x58, 0xc8, 0x99, 0x1b, 0x4c, 0xad, 0x34, 0xaf, 0xff,
	0x93, 0x00, 0x3d, 0x8d, 0x50, 0x2e, 0xc6, 0xa8, 0xa5, 0x6b, 0x33, 0xbc, 0x88, 0x50, 0x69, 0xda,
	0x85, 0x6b, 0x13, 0x9c, 0xba, 0xc1, 0x58, 0x75, 0x49, 0x8f, 0x0c, 0x2a, 0x2c, 0x97, 0xb4, 0x03,
	0x35, 0x0c, 0x9c, 0xb1, 0xea, 0x5e, 0x35, 0xf1, 0x54, 0xd0, 0xdb, 0xb0, 0xa3, 0x44, 0x24, 0x6d,
	0xec, 0x56, 0x7a, 0x64, 0x50, 0x67, 0x99, 0xa2, 0x2f, 0xa1, 0xaa, 0x17, 0x21, 0x76, 0xab, 0x3d,
	0x32, 0x68, 0x1d, 0x3c, 0xb0, 0xd6, 0xe5, 0xb2, 0x52, 0x9c, 0xb3, 0x45, 0x88, 0xcc, 0x54, 0xe8,
	0xdf, 0x83, 0x9b, 0x25, 0x4e, 0x15, 0x8a, 0x40, 0x61, 0x02, 0xea, 0x9b, 0x48, 0x0a, 0xda, 0x60,
	0xb9, 0xec, 0xbf, 0x86, 0x5b, 0x1f, 0x42, 0x87, 0x6b, 0x3c, 0xf2, 0x44, 0xcc, 0x22, 0x0f, 0xf3,
	0xb3, 0xad, 0x58, 0x49, 0x89, 0x75, 0x17, 0xae, 0x73, 0xcf, 0x13, 0xf1, 0x69, 0x98, 0x1e, 0x8e,
	0xb0, 0xa5, 0xee, 0xef, 0xc3, 0x0d, 0xd3, 0xfd, 0x44, 0x8a, 0x50, 0x15, 0x86, 0x34, 0xc3, 0x45,
	0x2c, 0xa4, 0x93, 0x55, 0xca, 0xe5, 0x12, 0xf6, 0xd5, 0xdc, 0x3f, 0x0e, 0xce, 0xc5, 0xff, 0x13,
	0x3e, 0x43, 0x27, 0x85, 0x7d, 0x9f, 0xee, 0x31, 0xcf, 0x68, 0x43, 0x25, 0x92, 0x5e, 0xf6, 0x77,
	0xf2, 0x49, 0xf7, 0x00, 0xa4, 0x88, 0x34, 0xca, 0xb1, 0x70, 0xd0, 0x70, 0xd6, 0x59, 0x21, 0x92,
	0x9c, 0x2e, 0x46, 0x77, 0xfa, 0x49, 0x9b, 0x4d, 0xd4, 0x58, 0xa6, 0x92, 0xde, 0x73, 0x94, 0xca,
	0x15, 0x81, 0x59, 0x46, 0x85, 0xe5, 0xb2, 0xff, 0x11, 0x5a, 0x43, 0x33, 0xf4, 0xe5, 0x50, 0xf7,
	0x00, 0x94, 0xe6, 0x3a, 0x52, 0xc3, 0xa4, 0x07, 0x31, 0x75, 0x0a, 0x91, 0x74, 0xe8, 0x4a, 0xf1,
	0x69, 0x0e, 0x90, 0x4b, 0x4a, 0xa1, 0xea, 0x70, 0xcd, 0x4d, 0xef, 0x06, 0x33, 0xdf, 0x77, 0xef,
	0x00, 0xac, 0xb6, 0x49, 0x9b, 0x50, 0x3f, 0x7e, 0x7b, 0x36, 0x62, 0x47, 0xcf, 0x87, 0xa3, 0xf6,
	0x15, 0x5a, 0x87, 0xda, 0xf8, 0xdd, 0x8b, 0xd1, 0x9b, 0x36, 0x39, 0xf8, 0x55, 0x83, 0x66, 0x0a,
	0x92, 0x0d, 0x81, 0x7e, 0x25, 0xd0, 0xb8, 0x58, 0x2d, 0x5d, 0xd1, 0xc7, 0xeb, 0x1b, 0xe8, 0x6f,
	0x53, 0xef, 0x3e, 0x5c, 0x3f, 0xfb, 0x8f, 0x81, 0xfc, 0x20, 0xd0, 0x8a, 0x4a, 0x66, 0xa2, 0xcf,
	0xd6, 0x2f, 0xf6, 0x4f, 0x1b, 0x6e, 0x41, 0xf3, 0x85, 0x40, 0xdd, 0x73, 0x95, 0x36, 0x66, 0xa4,
	0x8f, 0x36, 0x9c, 0x49, 0xd1, 0xc2, 0x5b, 0x40, 0x7c, 0xcb, 0x77, 0x93, 0x79, 0x9c, 0x3e, 0xd9,
	0x90, 0xa3, 0x7c, 0x37, 0xb6, 0x20, 0xf9, 0x4e, 0xa0, 0x19, 0x15, 0x2f, 0x0f, 0x7d, 0xba, 0xe9,
	0x6e, 0xca, 0xb7, 0xee, 0xf2, 0x2c, 0x87, 0x9d, 0x43, 0x5a, 0xb2, 0xf0, 0x49, 0xf2, 0x1c, 0x4f,
	0x76, 0xcc, 0xab, 0x7c, 0xff, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0x5d, 0xe1, 0xd0, 0xad,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommonServiceClient is the client API for CommonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommonServiceClient interface {
	QueryMetrics(ctx context.Context, in *QueryMetricRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdateFlowRule(ctx context.Context, in *UpdateFlowRuleRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ListProps(ctx context.Context, in *QueryPropsRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	QueryJvmInfo(ctx context.Context, in *QueryJvmInfoRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type commonServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommonServiceClient(cc *grpc.ClientConn) CommonServiceClient {
	return &commonServiceClient{cc}
}

func (c *commonServiceClient) QueryMetrics(ctx context.Context, in *QueryMetricRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/com.webank.ai.fate.api.networking.common.CommonService/queryMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) UpdateFlowRule(ctx context.Context, in *UpdateFlowRuleRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/com.webank.ai.fate.api.networking.common.CommonService/updateFlowRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) ListProps(ctx context.Context, in *QueryPropsRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/com.webank.ai.fate.api.networking.common.CommonService/listProps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) QueryJvmInfo(ctx context.Context, in *QueryJvmInfoRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/com.webank.ai.fate.api.networking.common.CommonService/queryJvmInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonServiceClient) UpdateService(ctx context.Context, in *UpdateServiceRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/com.webank.ai.fate.api.networking.common.CommonService/updateService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServiceServer is the server API for CommonService service.
type CommonServiceServer interface {
	QueryMetrics(context.Context, *QueryMetricRequest) (*CommonResponse, error)
	UpdateFlowRule(context.Context, *UpdateFlowRuleRequest) (*CommonResponse, error)
	ListProps(context.Context, *QueryPropsRequest) (*CommonResponse, error)
	QueryJvmInfo(context.Context, *QueryJvmInfoRequest) (*CommonResponse, error)
	UpdateService(context.Context, *UpdateServiceRequest) (*CommonResponse, error)
}

// UnimplementedCommonServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommonServiceServer struct {
}

func (*UnimplementedCommonServiceServer) QueryMetrics(ctx context.Context, req *QueryMetricRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMetrics not implemented")
}
func (*UnimplementedCommonServiceServer) UpdateFlowRule(ctx context.Context, req *UpdateFlowRuleRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlowRule not implemented")
}
func (*UnimplementedCommonServiceServer) ListProps(ctx context.Context, req *QueryPropsRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProps not implemented")
}
func (*UnimplementedCommonServiceServer) QueryJvmInfo(ctx context.Context, req *QueryJvmInfoRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryJvmInfo not implemented")
}
func (*UnimplementedCommonServiceServer) UpdateService(ctx context.Context, req *UpdateServiceRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateService not implemented")
}

func RegisterCommonServiceServer(s *grpc.Server, srv CommonServiceServer) {
	s.RegisterService(&_CommonService_serviceDesc, srv)
}

func _CommonService_QueryMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).QueryMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.webank.ai.fate.api.networking.common.CommonService/QueryMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).QueryMetrics(ctx, req.(*QueryMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_UpdateFlowRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlowRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).UpdateFlowRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.webank.ai.fate.api.networking.common.CommonService/UpdateFlowRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).UpdateFlowRule(ctx, req.(*UpdateFlowRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_ListProps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPropsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).ListProps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.webank.ai.fate.api.networking.common.CommonService/ListProps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).ListProps(ctx, req.(*QueryPropsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_QueryJvmInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryJvmInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).QueryJvmInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.webank.ai.fate.api.networking.common.CommonService/QueryJvmInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).QueryJvmInfo(ctx, req.(*QueryJvmInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonService_UpdateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServiceServer).UpdateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.webank.ai.fate.api.networking.common.CommonService/UpdateService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServiceServer).UpdateService(ctx, req.(*UpdateServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommonService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.webank.ai.fate.api.networking.common.CommonService",
	HandlerType: (*CommonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "queryMetrics",
			Handler:    _CommonService_QueryMetrics_Handler,
		},
		{
			MethodName: "updateFlowRule",
			Handler:    _CommonService_UpdateFlowRule_Handler,
		},
		{
			MethodName: "listProps",
			Handler:    _CommonService_ListProps_Handler,
		},
		{
			MethodName: "queryJvmInfo",
			Handler:    _CommonService_QueryJvmInfo_Handler,
		},
		{
			MethodName: "updateService",
			Handler:    _CommonService_UpdateService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common-service.proto",
}
