// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.1
// source: google/showcase/v1beta1/sequence.proto

package genproto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Sequence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Sequence of responses to return in order for each attempt. If empty, the
	// default response is an immediate OK.
	Responses []*Sequence_Response `protobuf:"bytes,2,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *Sequence) Reset() {
	*x = Sequence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sequence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sequence) ProtoMessage() {}

func (x *Sequence) ProtoReflect() protoreflect.Message {
	mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sequence.ProtoReflect.Descriptor instead.
func (*Sequence) Descriptor() ([]byte, []int) {
	return file_google_showcase_v1beta1_sequence_proto_rawDescGZIP(), []int{0}
}

func (x *Sequence) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sequence) GetResponses() []*Sequence_Response {
	if x != nil {
		return x.Responses
	}
	return nil
}

type SequenceReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The set of RPC attempts received by the server for a Sequence.
	Attempts []*SequenceReport_Attempt `protobuf:"bytes,2,rep,name=attempts,proto3" json:"attempts,omitempty"`
}

func (x *SequenceReport) Reset() {
	*x = SequenceReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SequenceReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SequenceReport) ProtoMessage() {}

func (x *SequenceReport) ProtoReflect() protoreflect.Message {
	mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SequenceReport.ProtoReflect.Descriptor instead.
func (*SequenceReport) Descriptor() ([]byte, []int) {
	return file_google_showcase_v1beta1_sequence_proto_rawDescGZIP(), []int{1}
}

func (x *SequenceReport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SequenceReport) GetAttempts() []*SequenceReport_Attempt {
	if x != nil {
		return x.Attempts
	}
	return nil
}

type CreateSequenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sequence *Sequence `protobuf:"bytes,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *CreateSequenceRequest) Reset() {
	*x = CreateSequenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSequenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSequenceRequest) ProtoMessage() {}

func (x *CreateSequenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSequenceRequest.ProtoReflect.Descriptor instead.
func (*CreateSequenceRequest) Descriptor() ([]byte, []int) {
	return file_google_showcase_v1beta1_sequence_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSequenceRequest) GetSequence() *Sequence {
	if x != nil {
		return x.Sequence
	}
	return nil
}

type AttemptSequenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AttemptSequenceRequest) Reset() {
	*x = AttemptSequenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttemptSequenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttemptSequenceRequest) ProtoMessage() {}

func (x *AttemptSequenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttemptSequenceRequest.ProtoReflect.Descriptor instead.
func (*AttemptSequenceRequest) Descriptor() ([]byte, []int) {
	return file_google_showcase_v1beta1_sequence_proto_rawDescGZIP(), []int{3}
}

func (x *AttemptSequenceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetSequenceReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSequenceReportRequest) Reset() {
	*x = GetSequenceReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSequenceReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSequenceReportRequest) ProtoMessage() {}

func (x *GetSequenceReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSequenceReportRequest.ProtoReflect.Descriptor instead.
func (*GetSequenceReportRequest) Descriptor() ([]byte, []int) {
	return file_google_showcase_v1beta1_sequence_proto_rawDescGZIP(), []int{4}
}

func (x *GetSequenceReportRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// A server response to an RPC Attempt in a sequence.
type Sequence_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status to return for an individual attempt.
	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// The amount of time to delay sending the response.
	Delay *duration.Duration `protobuf:"bytes,2,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *Sequence_Response) Reset() {
	*x = Sequence_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sequence_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sequence_Response) ProtoMessage() {}

func (x *Sequence_Response) ProtoReflect() protoreflect.Message {
	mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sequence_Response.ProtoReflect.Descriptor instead.
func (*Sequence_Response) Descriptor() ([]byte, []int) {
	return file_google_showcase_v1beta1_sequence_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Sequence_Response) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *Sequence_Response) GetDelay() *duration.Duration {
	if x != nil {
		return x.Delay
	}
	return nil
}

// Contains metrics on individual RPC Attempts in a sequence.
type SequenceReport_Attempt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The attempt number - starting at 0.
	AttemptNumber int32 `protobuf:"varint,1,opt,name=attempt_number,json=attemptNumber,proto3" json:"attempt_number,omitempty"`
	// The deadline dictated by the attempt to the server.
	AttemptDeadline *timestamp.Timestamp `protobuf:"bytes,2,opt,name=attempt_deadline,json=attemptDeadline,proto3" json:"attempt_deadline,omitempty"`
	// The time that the server responded to the RPC attempt. Used for
	// calculating attempt_delay.
	ResponseTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	// The server perceived delay between sending the last response and
	// receiving this attempt. Used for validating attempt delay backoff.
	AttemptDelay *duration.Duration `protobuf:"bytes,4,opt,name=attempt_delay,json=attemptDelay,proto3" json:"attempt_delay,omitempty"`
	// The status returned to the attempt.
	Status *status.Status `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SequenceReport_Attempt) Reset() {
	*x = SequenceReport_Attempt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SequenceReport_Attempt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SequenceReport_Attempt) ProtoMessage() {}

func (x *SequenceReport_Attempt) ProtoReflect() protoreflect.Message {
	mi := &file_google_showcase_v1beta1_sequence_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SequenceReport_Attempt.ProtoReflect.Descriptor instead.
func (*SequenceReport_Attempt) Descriptor() ([]byte, []int) {
	return file_google_showcase_v1beta1_sequence_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SequenceReport_Attempt) GetAttemptNumber() int32 {
	if x != nil {
		return x.AttemptNumber
	}
	return 0
}

func (x *SequenceReport_Attempt) GetAttemptDeadline() *timestamp.Timestamp {
	if x != nil {
		return x.AttemptDeadline
	}
	return nil
}

func (x *SequenceReport_Attempt) GetResponseTime() *timestamp.Timestamp {
	if x != nil {
		return x.ResponseTime
	}
	return nil
}

func (x *SequenceReport_Attempt) GetAttemptDelay() *duration.Duration {
	if x != nil {
		return x.AttemptDelay
	}
	return nil
}

func (x *SequenceReport_Attempt) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_google_showcase_v1beta1_sequence_proto protoreflect.FileDescriptor

var file_google_showcase_v1beta1_sequence_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x08,
	0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x48, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x1a, 0x67, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x3a, 0x3b, 0xea, 0x41, 0x38, 0x0a, 0x20, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x7d, 0x22, 0xef, 0x03, 0x0a, 0x0e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a,
	0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x1a, 0xa4, 0x02, 0x0a, 0x07, 0x41,
	0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x45, 0x0a,
	0x10, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x44, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x3a, 0x50, 0xea, 0x41, 0x4d, 0x0a, 0x26, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x23,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x7d, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0x56, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x08,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x56, 0x0a, 0x16, 0x41,
	0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x28, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61,
	0x73, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x5e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x42, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xfa,
	0x41, 0x28, 0x0a, 0x26, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0xf4, 0x03, 0x0a, 0x0f, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x2f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x3a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0xda, 0x41, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0xaa,
	0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x0f,
	0x41, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x22, 0x1b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x01, 0x2a,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x11, 0xca, 0x41, 0x0e, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x37, 0x34, 0x36, 0x39, 0x42, 0x71, 0x0a, 0x1b, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x67, 0x61, 0x70, 0x69, 0x63, 0x2d, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x53, 0x68, 0x6f, 0x77,
	0x63, 0x61, 0x73, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_showcase_v1beta1_sequence_proto_rawDescOnce sync.Once
	file_google_showcase_v1beta1_sequence_proto_rawDescData = file_google_showcase_v1beta1_sequence_proto_rawDesc
)

func file_google_showcase_v1beta1_sequence_proto_rawDescGZIP() []byte {
	file_google_showcase_v1beta1_sequence_proto_rawDescOnce.Do(func() {
		file_google_showcase_v1beta1_sequence_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_showcase_v1beta1_sequence_proto_rawDescData)
	})
	return file_google_showcase_v1beta1_sequence_proto_rawDescData
}

var file_google_showcase_v1beta1_sequence_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_showcase_v1beta1_sequence_proto_goTypes = []interface{}{
	(*Sequence)(nil),                 // 0: google.showcase.v1beta1.Sequence
	(*SequenceReport)(nil),           // 1: google.showcase.v1beta1.SequenceReport
	(*CreateSequenceRequest)(nil),    // 2: google.showcase.v1beta1.CreateSequenceRequest
	(*AttemptSequenceRequest)(nil),   // 3: google.showcase.v1beta1.AttemptSequenceRequest
	(*GetSequenceReportRequest)(nil), // 4: google.showcase.v1beta1.GetSequenceReportRequest
	(*Sequence_Response)(nil),        // 5: google.showcase.v1beta1.Sequence.Response
	(*SequenceReport_Attempt)(nil),   // 6: google.showcase.v1beta1.SequenceReport.Attempt
	(*status.Status)(nil),            // 7: google.rpc.Status
	(*duration.Duration)(nil),        // 8: google.protobuf.Duration
	(*timestamp.Timestamp)(nil),      // 9: google.protobuf.Timestamp
	(*empty.Empty)(nil),              // 10: google.protobuf.Empty
}
var file_google_showcase_v1beta1_sequence_proto_depIdxs = []int32{
	5,  // 0: google.showcase.v1beta1.Sequence.responses:type_name -> google.showcase.v1beta1.Sequence.Response
	6,  // 1: google.showcase.v1beta1.SequenceReport.attempts:type_name -> google.showcase.v1beta1.SequenceReport.Attempt
	0,  // 2: google.showcase.v1beta1.CreateSequenceRequest.sequence:type_name -> google.showcase.v1beta1.Sequence
	7,  // 3: google.showcase.v1beta1.Sequence.Response.status:type_name -> google.rpc.Status
	8,  // 4: google.showcase.v1beta1.Sequence.Response.delay:type_name -> google.protobuf.Duration
	9,  // 5: google.showcase.v1beta1.SequenceReport.Attempt.attempt_deadline:type_name -> google.protobuf.Timestamp
	9,  // 6: google.showcase.v1beta1.SequenceReport.Attempt.response_time:type_name -> google.protobuf.Timestamp
	8,  // 7: google.showcase.v1beta1.SequenceReport.Attempt.attempt_delay:type_name -> google.protobuf.Duration
	7,  // 8: google.showcase.v1beta1.SequenceReport.Attempt.status:type_name -> google.rpc.Status
	2,  // 9: google.showcase.v1beta1.SequenceService.CreateSequence:input_type -> google.showcase.v1beta1.CreateSequenceRequest
	4,  // 10: google.showcase.v1beta1.SequenceService.GetSequenceReport:input_type -> google.showcase.v1beta1.GetSequenceReportRequest
	3,  // 11: google.showcase.v1beta1.SequenceService.AttemptSequence:input_type -> google.showcase.v1beta1.AttemptSequenceRequest
	0,  // 12: google.showcase.v1beta1.SequenceService.CreateSequence:output_type -> google.showcase.v1beta1.Sequence
	1,  // 13: google.showcase.v1beta1.SequenceService.GetSequenceReport:output_type -> google.showcase.v1beta1.SequenceReport
	10, // 14: google.showcase.v1beta1.SequenceService.AttemptSequence:output_type -> google.protobuf.Empty
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_google_showcase_v1beta1_sequence_proto_init() }
func file_google_showcase_v1beta1_sequence_proto_init() {
	if File_google_showcase_v1beta1_sequence_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_showcase_v1beta1_sequence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sequence); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_showcase_v1beta1_sequence_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SequenceReport); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_showcase_v1beta1_sequence_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSequenceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_showcase_v1beta1_sequence_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttemptSequenceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_showcase_v1beta1_sequence_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSequenceReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_showcase_v1beta1_sequence_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sequence_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_showcase_v1beta1_sequence_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SequenceReport_Attempt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_showcase_v1beta1_sequence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_showcase_v1beta1_sequence_proto_goTypes,
		DependencyIndexes: file_google_showcase_v1beta1_sequence_proto_depIdxs,
		MessageInfos:      file_google_showcase_v1beta1_sequence_proto_msgTypes,
	}.Build()
	File_google_showcase_v1beta1_sequence_proto = out.File
	file_google_showcase_v1beta1_sequence_proto_rawDesc = nil
	file_google_showcase_v1beta1_sequence_proto_goTypes = nil
	file_google_showcase_v1beta1_sequence_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SequenceServiceClient is the client API for SequenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SequenceServiceClient interface {
	CreateSequence(ctx context.Context, in *CreateSequenceRequest, opts ...grpc.CallOption) (*Sequence, error)
	GetSequenceReport(ctx context.Context, in *GetSequenceReportRequest, opts ...grpc.CallOption) (*SequenceReport, error)
	AttemptSequence(ctx context.Context, in *AttemptSequenceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sequenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSequenceServiceClient(cc grpc.ClientConnInterface) SequenceServiceClient {
	return &sequenceServiceClient{cc}
}

func (c *sequenceServiceClient) CreateSequence(ctx context.Context, in *CreateSequenceRequest, opts ...grpc.CallOption) (*Sequence, error) {
	out := new(Sequence)
	err := c.cc.Invoke(ctx, "/google.showcase.v1beta1.SequenceService/CreateSequence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceServiceClient) GetSequenceReport(ctx context.Context, in *GetSequenceReportRequest, opts ...grpc.CallOption) (*SequenceReport, error) {
	out := new(SequenceReport)
	err := c.cc.Invoke(ctx, "/google.showcase.v1beta1.SequenceService/GetSequenceReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceServiceClient) AttemptSequence(ctx context.Context, in *AttemptSequenceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.showcase.v1beta1.SequenceService/AttemptSequence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SequenceServiceServer is the server API for SequenceService service.
type SequenceServiceServer interface {
	CreateSequence(context.Context, *CreateSequenceRequest) (*Sequence, error)
	GetSequenceReport(context.Context, *GetSequenceReportRequest) (*SequenceReport, error)
	AttemptSequence(context.Context, *AttemptSequenceRequest) (*empty.Empty, error)
}

// UnimplementedSequenceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSequenceServiceServer struct {
}

func (*UnimplementedSequenceServiceServer) CreateSequence(context.Context, *CreateSequenceRequest) (*Sequence, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method CreateSequence not implemented")
}
func (*UnimplementedSequenceServiceServer) GetSequenceReport(context.Context, *GetSequenceReportRequest) (*SequenceReport, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetSequenceReport not implemented")
}
func (*UnimplementedSequenceServiceServer) AttemptSequence(context.Context, *AttemptSequenceRequest) (*empty.Empty, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method AttemptSequence not implemented")
}

func RegisterSequenceServiceServer(s *grpc.Server, srv SequenceServiceServer) {
	s.RegisterService(&_SequenceService_serviceDesc, srv)
}

func _SequenceService_CreateSequence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSequenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServiceServer).CreateSequence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1beta1.SequenceService/CreateSequence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServiceServer).CreateSequence(ctx, req.(*CreateSequenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SequenceService_GetSequenceReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSequenceReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServiceServer).GetSequenceReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1beta1.SequenceService/GetSequenceReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServiceServer).GetSequenceReport(ctx, req.(*GetSequenceReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SequenceService_AttemptSequence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttemptSequenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SequenceServiceServer).AttemptSequence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1beta1.SequenceService/AttemptSequence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SequenceServiceServer).AttemptSequence(ctx, req.(*AttemptSequenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SequenceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.showcase.v1beta1.SequenceService",
	HandlerType: (*SequenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSequence",
			Handler:    _SequenceService_CreateSequence_Handler,
		},
		{
			MethodName: "GetSequenceReport",
			Handler:    _SequenceService_GetSequenceReport_Handler,
		},
		{
			MethodName: "AttemptSequence",
			Handler:    _SequenceService_AttemptSequence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/showcase/v1beta1/sequence.proto",
}
