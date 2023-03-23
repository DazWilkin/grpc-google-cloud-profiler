// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/devtools/cloudprofiler/v2/profiler.proto

package cloudprofiler

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProfileType is type of profiling data.
// NOTE: the enumeration member names are used (in lowercase) as unique string
// identifiers of profile types, so they must not be renamed.
type ProfileType int32

const (
	// Unspecified profile type.
	ProfileType_PROFILE_TYPE_UNSPECIFIED ProfileType = 0
	// Thread CPU time sampling.
	ProfileType_CPU ProfileType = 1
	// Wallclock time sampling. More expensive as stops all threads.
	ProfileType_WALL ProfileType = 2
	// In-use heap profile. Represents a snapshot of the allocations that are
	// live at the time of the profiling.
	ProfileType_HEAP ProfileType = 3
	// Single-shot collection of all thread stacks.
	ProfileType_THREADS ProfileType = 4
	// Synchronization contention profile.
	ProfileType_CONTENTION ProfileType = 5
	// Peak heap profile.
	ProfileType_PEAK_HEAP ProfileType = 6
	// Heap allocation profile. It represents the aggregation of all allocations
	// made over the duration of the profile. All allocations are included,
	// including those that might have been freed by the end of the profiling
	// interval. The profile is in particular useful for garbage collecting
	// languages to understand which parts of the code create most of the garbage
	// collection pressure to see if those can be optimized.
	ProfileType_HEAP_ALLOC ProfileType = 7
)

// Enum value maps for ProfileType.
var (
	ProfileType_name = map[int32]string{
		0: "PROFILE_TYPE_UNSPECIFIED",
		1: "CPU",
		2: "WALL",
		3: "HEAP",
		4: "THREADS",
		5: "CONTENTION",
		6: "PEAK_HEAP",
		7: "HEAP_ALLOC",
	}
	ProfileType_value = map[string]int32{
		"PROFILE_TYPE_UNSPECIFIED": 0,
		"CPU":                      1,
		"WALL":                     2,
		"HEAP":                     3,
		"THREADS":                  4,
		"CONTENTION":               5,
		"PEAK_HEAP":                6,
		"HEAP_ALLOC":               7,
	}
)

func (x ProfileType) Enum() *ProfileType {
	p := new(ProfileType)
	*p = x
	return p
}

func (x ProfileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProfileType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_cloudprofiler_v2_profiler_proto_enumTypes[0].Descriptor()
}

func (ProfileType) Type() protoreflect.EnumType {
	return &file_google_devtools_cloudprofiler_v2_profiler_proto_enumTypes[0]
}

func (x ProfileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProfileType.Descriptor instead.
func (ProfileType) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescGZIP(), []int{0}
}

// CreateProfileRequest describes a profile resource online creation request.
// The deployment field must be populated. The profile_type specifies the list
// of profile types supported by the agent. The creation call will hang until a
// profile of one of these types needs to be collected.
type CreateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Parent project to create the profile in.
	Parent string `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	// Deployment details.
	Deployment *Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	// One or more profile types that the agent is capable of providing.
	ProfileType []ProfileType `protobuf:"varint,2,rep,packed,name=profile_type,json=profileType,proto3,enum=google.devtools.cloudprofiler.v2.ProfileType" json:"profile_type,omitempty"`
}

func (x *CreateProfileRequest) Reset() {
	*x = CreateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileRequest) ProtoMessage() {}

func (x *CreateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProfileRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateProfileRequest) GetDeployment() *Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *CreateProfileRequest) GetProfileType() []ProfileType {
	if x != nil {
		return x.ProfileType
	}
	return nil
}

// CreateOfflineProfileRequest describes a profile resource offline creation
// request.
type CreateOfflineProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Parent project to create the profile in.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Contents of the profile to create.
	Profile *Profile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *CreateOfflineProfileRequest) Reset() {
	*x = CreateOfflineProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOfflineProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOfflineProfileRequest) ProtoMessage() {}

func (x *CreateOfflineProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOfflineProfileRequest.ProtoReflect.Descriptor instead.
func (*CreateOfflineProfileRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOfflineProfileRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateOfflineProfileRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

// UpdateProfileRequest contains the profile to update.
type UpdateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Profile to update.
	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	// Field mask used to specify the fields to be overwritten. Currently only
	// profile_bytes and labels fields are supported by UpdateProfile, so only
	// those fields can be specified in the mask. When no mask is provided, all
	// fields are overwritten.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProfileRequest) GetProfile() *Profile {
	if x != nil {
		return x.Profile
	}
	return nil
}

func (x *UpdateProfileRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Profile resource.
type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Opaque, server-assigned, unique ID for this profile.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of profile.
	// For offline mode, this must be specified when creating the profile. For
	// online mode it is assigned and returned by the server.
	ProfileType ProfileType `protobuf:"varint,2,opt,name=profile_type,json=profileType,proto3,enum=google.devtools.cloudprofiler.v2.ProfileType" json:"profile_type,omitempty"`
	// Deployment this profile corresponds to.
	Deployment *Deployment `protobuf:"bytes,3,opt,name=deployment,proto3" json:"deployment,omitempty"`
	// Duration of the profiling session.
	// Input (for the offline mode) or output (for the online mode).
	// The field represents requested profiling duration. It may slightly differ
	// from the effective profiling duration, which is recorded in the profile
	// data, in case the profiling can't be stopped immediately (e.g. in case
	// stopping the profiling is handled asynchronously).
	Duration *durationpb.Duration `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	// Input only. Profile bytes, as a gzip compressed serialized proto, the
	// format is https://github.com/google/pprof/blob/master/proto/profile.proto.
	ProfileBytes []byte `protobuf:"bytes,5,opt,name=profile_bytes,json=profileBytes,proto3" json:"profile_bytes,omitempty"`
	// Input only. Labels associated to this specific profile. These labels will
	// get merged with the deployment labels for the final data set. See
	// documentation on deployment labels for validation rules and limits.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescGZIP(), []int{3}
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetProfileType() ProfileType {
	if x != nil {
		return x.ProfileType
	}
	return ProfileType_PROFILE_TYPE_UNSPECIFIED
}

func (x *Profile) GetDeployment() *Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *Profile) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Profile) GetProfileBytes() []byte {
	if x != nil {
		return x.ProfileBytes
	}
	return nil
}

func (x *Profile) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Deployment contains the deployment identification information.
type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project ID is the ID of a cloud project.
	// Validation regex: `^[a-z][-a-z0-9:.]{4,61}[a-z0-9]$`.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Target is the service name used to group related deployments:
	// * Service name for App Engine Flex / Standard.
	// * Cluster and container name for GKE.
	// * User-specified string for direct Compute Engine profiling (e.g. Java).
	// * Job name for Dataflow.
	// Validation regex: `^[a-z0-9]([-a-z0-9_.]{0,253}[a-z0-9])?$`.
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// Labels identify the deployment within the user universe and same target.
	// Validation regex for label names: `^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// Value for an individual label must be <= 512 bytes, the total
	// size of all label names and values must be <= 1024 bytes.
	//
	// Label named "language" can be used to record the programming language of
	// the profiled deployment. The standard choices for the value include "java",
	// "go", "python", "ruby", "nodejs", "php", "dotnet".
	//
	// For deployments running on Google Cloud Platform, "zone" or "region" label
	// should be present describing the deployment location. An example of a zone
	// is "us-central1-a", an example of a region is "us-central1" or
	// "us-central".
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescGZIP(), []int{4}
}

func (x *Deployment) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Deployment) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Deployment) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_google_devtools_cloudprofiler_v2_profiler_proto protoreflect.FileDescriptor

var file_google_devtools_cloudprofiler_v2_profiler_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f,
	0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x48, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x30, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x0a, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x32, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x50, 0x0a, 0x0c, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x1b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xfa, 0x41, 0x2d,
	0x0a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x84, 0x04, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x0a,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x28, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x04, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x04, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x50, 0xea, 0x41, 0x4d, 0x0a,
	0x24, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x25, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x7d, 0x22, 0xd0, 0x01, 0x0a,
	0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x50, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a,
	0x84, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x43, 0x50, 0x55, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x4c, 0x4c, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x50, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x48,
	0x52, 0x45, 0x41, 0x44, 0x53, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x54, 0x45,
	0x4e, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x45, 0x41, 0x4b, 0x5f,
	0x48, 0x45, 0x41, 0x50, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x48, 0x45, 0x41, 0x50, 0x5f, 0x41,
	0x4c, 0x4c, 0x4f, 0x43, 0x10, 0x07, 0x32, 0xfe, 0x05, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x36, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xd2, 0x01, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22,
	0x50, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x22, 0x2e, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x3a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0xda, 0x41, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0xc3, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x4f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x32, 0x28,
	0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0xda, 0x41, 0x13, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x1a, 0xad, 0x01, 0xca, 0x41, 0x1c, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x8a, 0x01, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77,
	0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0xda, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x42, 0x0d, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f,
	0x76, 0x32, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0xaa, 0x02, 0x18, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x18, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a,
	0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescOnce sync.Once
	file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescData = file_google_devtools_cloudprofiler_v2_profiler_proto_rawDesc
)

func file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescGZIP() []byte {
	file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescOnce.Do(func() {
		file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescData)
	})
	return file_google_devtools_cloudprofiler_v2_profiler_proto_rawDescData
}

var file_google_devtools_cloudprofiler_v2_profiler_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_google_devtools_cloudprofiler_v2_profiler_proto_goTypes = []interface{}{
	(ProfileType)(0),                    // 0: google.devtools.cloudprofiler.v2.ProfileType
	(*CreateProfileRequest)(nil),        // 1: google.devtools.cloudprofiler.v2.CreateProfileRequest
	(*CreateOfflineProfileRequest)(nil), // 2: google.devtools.cloudprofiler.v2.CreateOfflineProfileRequest
	(*UpdateProfileRequest)(nil),        // 3: google.devtools.cloudprofiler.v2.UpdateProfileRequest
	(*Profile)(nil),                     // 4: google.devtools.cloudprofiler.v2.Profile
	(*Deployment)(nil),                  // 5: google.devtools.cloudprofiler.v2.Deployment
	nil,                                 // 6: google.devtools.cloudprofiler.v2.Profile.LabelsEntry
	nil,                                 // 7: google.devtools.cloudprofiler.v2.Deployment.LabelsEntry
	(*fieldmaskpb.FieldMask)(nil),       // 8: google.protobuf.FieldMask
	(*durationpb.Duration)(nil),         // 9: google.protobuf.Duration
}
var file_google_devtools_cloudprofiler_v2_profiler_proto_depIdxs = []int32{
	5,  // 0: google.devtools.cloudprofiler.v2.CreateProfileRequest.deployment:type_name -> google.devtools.cloudprofiler.v2.Deployment
	0,  // 1: google.devtools.cloudprofiler.v2.CreateProfileRequest.profile_type:type_name -> google.devtools.cloudprofiler.v2.ProfileType
	4,  // 2: google.devtools.cloudprofiler.v2.CreateOfflineProfileRequest.profile:type_name -> google.devtools.cloudprofiler.v2.Profile
	4,  // 3: google.devtools.cloudprofiler.v2.UpdateProfileRequest.profile:type_name -> google.devtools.cloudprofiler.v2.Profile
	8,  // 4: google.devtools.cloudprofiler.v2.UpdateProfileRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 5: google.devtools.cloudprofiler.v2.Profile.profile_type:type_name -> google.devtools.cloudprofiler.v2.ProfileType
	5,  // 6: google.devtools.cloudprofiler.v2.Profile.deployment:type_name -> google.devtools.cloudprofiler.v2.Deployment
	9,  // 7: google.devtools.cloudprofiler.v2.Profile.duration:type_name -> google.protobuf.Duration
	6,  // 8: google.devtools.cloudprofiler.v2.Profile.labels:type_name -> google.devtools.cloudprofiler.v2.Profile.LabelsEntry
	7,  // 9: google.devtools.cloudprofiler.v2.Deployment.labels:type_name -> google.devtools.cloudprofiler.v2.Deployment.LabelsEntry
	1,  // 10: google.devtools.cloudprofiler.v2.ProfilerService.CreateProfile:input_type -> google.devtools.cloudprofiler.v2.CreateProfileRequest
	2,  // 11: google.devtools.cloudprofiler.v2.ProfilerService.CreateOfflineProfile:input_type -> google.devtools.cloudprofiler.v2.CreateOfflineProfileRequest
	3,  // 12: google.devtools.cloudprofiler.v2.ProfilerService.UpdateProfile:input_type -> google.devtools.cloudprofiler.v2.UpdateProfileRequest
	4,  // 13: google.devtools.cloudprofiler.v2.ProfilerService.CreateProfile:output_type -> google.devtools.cloudprofiler.v2.Profile
	4,  // 14: google.devtools.cloudprofiler.v2.ProfilerService.CreateOfflineProfile:output_type -> google.devtools.cloudprofiler.v2.Profile
	4,  // 15: google.devtools.cloudprofiler.v2.ProfilerService.UpdateProfile:output_type -> google.devtools.cloudprofiler.v2.Profile
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_devtools_cloudprofiler_v2_profiler_proto_init() }
func file_google_devtools_cloudprofiler_v2_profiler_proto_init() {
	if File_google_devtools_cloudprofiler_v2_profiler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileRequest); i {
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
		file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOfflineProfileRequest); i {
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
		file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileRequest); i {
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
		file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
			RawDescriptor: file_google_devtools_cloudprofiler_v2_profiler_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_devtools_cloudprofiler_v2_profiler_proto_goTypes,
		DependencyIndexes: file_google_devtools_cloudprofiler_v2_profiler_proto_depIdxs,
		EnumInfos:         file_google_devtools_cloudprofiler_v2_profiler_proto_enumTypes,
		MessageInfos:      file_google_devtools_cloudprofiler_v2_profiler_proto_msgTypes,
	}.Build()
	File_google_devtools_cloudprofiler_v2_profiler_proto = out.File
	file_google_devtools_cloudprofiler_v2_profiler_proto_rawDesc = nil
	file_google_devtools_cloudprofiler_v2_profiler_proto_goTypes = nil
	file_google_devtools_cloudprofiler_v2_profiler_proto_depIdxs = nil
}
