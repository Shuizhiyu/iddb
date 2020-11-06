// Copyright 2018 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/devtools/containeranalysis/v1beta1/package/package.proto

package _package

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Instruction set architectures supported by various package managers.
type Architecture int32

const (
	// Unknown architecture.
	Architecture_ARCHITECTURE_UNSPECIFIED Architecture = 0
	// X86 architecture.
	Architecture_X86 Architecture = 1
	// X64 architecture.
	Architecture_X64 Architecture = 2
)

// Enum value maps for Architecture.
var (
	Architecture_name = map[int32]string{
		0: "ARCHITECTURE_UNSPECIFIED",
		1: "X86",
		2: "X64",
	}
	Architecture_value = map[string]int32{
		"ARCHITECTURE_UNSPECIFIED": 0,
		"X86":                      1,
		"X64":                      2,
	}
)

func (x Architecture) Enum() *Architecture {
	p := new(Architecture)
	*p = x
	return p
}

func (x Architecture) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Architecture) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_enumTypes[0].Descriptor()
}

func (Architecture) Type() protoreflect.EnumType {
	return &file_google_devtools_containeranalysis_v1beta1_package_package_proto_enumTypes[0]
}

func (x Architecture) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Architecture.Descriptor instead.
func (Architecture) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescGZIP(), []int{0}
}

// Whether this is an ordinary package version or a sentinel MIN/MAX version.
type Version_VersionKind int32

const (
	// Unknown.
	Version_VERSION_KIND_UNSPECIFIED Version_VersionKind = 0
	// A standard package version.
	Version_NORMAL Version_VersionKind = 1
	// A special version representing negative infinity.
	Version_MINIMUM Version_VersionKind = 2
	// A special version representing positive infinity.
	Version_MAXIMUM Version_VersionKind = 3
)

// Enum value maps for Version_VersionKind.
var (
	Version_VersionKind_name = map[int32]string{
		0: "VERSION_KIND_UNSPECIFIED",
		1: "NORMAL",
		2: "MINIMUM",
		3: "MAXIMUM",
	}
	Version_VersionKind_value = map[string]int32{
		"VERSION_KIND_UNSPECIFIED": 0,
		"NORMAL":                   1,
		"MINIMUM":                  2,
		"MAXIMUM":                  3,
	}
)

func (x Version_VersionKind) Enum() *Version_VersionKind {
	p := new(Version_VersionKind)
	*p = x
	return p
}

func (x Version_VersionKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version_VersionKind) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_enumTypes[1].Descriptor()
}

func (Version_VersionKind) Type() protoreflect.EnumType {
	return &file_google_devtools_containeranalysis_v1beta1_package_package_proto_enumTypes[1]
}

func (x Version_VersionKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version_VersionKind.Descriptor instead.
func (Version_VersionKind) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescGZIP(), []int{5, 0}
}

// This represents a particular channel of distribution for a given package.
// E.g., Debian's jessie-backports dpkg mirror.
type Distribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The cpe_uri in [CPE format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// The CPU architecture for which packages in this distribution channel were
	// built.
	Architecture Architecture `protobuf:"varint,2,opt,name=architecture,proto3,enum=grafeas.v1beta1.package.Architecture" json:"architecture,omitempty"`
	// The latest available version of this package in this distribution channel.
	LatestVersion *Version `protobuf:"bytes,3,opt,name=latest_version,json=latestVersion,proto3" json:"latest_version,omitempty"`
	// A freeform string denoting the maintainer of this package.
	Maintainer string `protobuf:"bytes,4,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
	// The distribution channel-specific homepage for this package.
	Url string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	// The distribution channel-specific description of this package.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Distribution) Reset() {
	*x = Distribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Distribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Distribution) ProtoMessage() {}

func (x *Distribution) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Distribution.ProtoReflect.Descriptor instead.
func (*Distribution) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescGZIP(), []int{0}
}

func (x *Distribution) GetCpeUri() string {
	if x != nil {
		return x.CpeUri
	}
	return ""
}

func (x *Distribution) GetArchitecture() Architecture {
	if x != nil {
		return x.Architecture
	}
	return Architecture_ARCHITECTURE_UNSPECIFIED
}

func (x *Distribution) GetLatestVersion() *Version {
	if x != nil {
		return x.LatestVersion
	}
	return nil
}

func (x *Distribution) GetMaintainer() string {
	if x != nil {
		return x.Maintainer
	}
	return ""
}

func (x *Distribution) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Distribution) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// An occurrence of a particular package installation found within a system's
// filesystem. E.g., glibc was found in `/var/lib/dpkg/status`.
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The CPE URI in [CPE format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	CpeUri string `protobuf:"bytes,1,opt,name=cpe_uri,json=cpeUri,proto3" json:"cpe_uri,omitempty"`
	// The version installed at this location.
	Version *Version `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// The path from which we gathered that this package/version is installed.
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetCpeUri() string {
	if x != nil {
		return x.CpeUri
	}
	return ""
}

func (x *Location) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *Location) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// This represents a particular package that is distributed over various
// channels. E.g., glibc (aka libc6) is distributed by many, at various
// versions.
type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The name of the package.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The various channels by which a package is distributed.
	Distribution []*Distribution `protobuf:"bytes,10,rep,name=distribution,proto3" json:"distribution,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescGZIP(), []int{2}
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Package) GetDistribution() []*Distribution {
	if x != nil {
		return x.Distribution
	}
	return nil
}

// Details of a package occurrence.
type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Where the package was installed.
	Installation *Installation `protobuf:"bytes,1,opt,name=installation,proto3" json:"installation,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescGZIP(), []int{3}
}

func (x *Details) GetInstallation() *Installation {
	if x != nil {
		return x.Installation
	}
	return nil
}

// This represents how a particular software package may be installed on a
// system.
type Installation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The name of the installed package.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. All of the places within the filesystem versions of this package
	// have been found.
	Location []*Location `protobuf:"bytes,2,rep,name=location,proto3" json:"location,omitempty"`
}

func (x *Installation) Reset() {
	*x = Installation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Installation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Installation) ProtoMessage() {}

func (x *Installation) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Installation.ProtoReflect.Descriptor instead.
func (*Installation) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescGZIP(), []int{4}
}

func (x *Installation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Installation) GetLocation() []*Location {
	if x != nil {
		return x.Location
	}
	return nil
}

// Version contains structured information about the version of a package.
type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to correct mistakes in the version numbering scheme.
	Epoch int32 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Required only when version kind is NORMAL. The main part of the version
	// name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The iteration of the package build from the above version.
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// Required. Distinguishes between sentinel MIN/MAX versions and normal
	// versions.
	Kind Version_VersionKind `protobuf:"varint,4,opt,name=kind,proto3,enum=grafeas.v1beta1.package.Version_VersionKind" json:"kind,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescGZIP(), []int{5}
}

func (x *Version) GetEpoch() int32 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *Version) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Version) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *Version) GetKind() Version_VersionKind {
	if x != nil {
		return x.Kind
	}
	return Version_VERSION_KIND_UNSPECIFIED
}

var File_google_devtools_containeranalysis_v1beta1_package_package_proto protoreflect.FileDescriptor

var file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x8f, 0x02, 0x0a, 0x0c, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x70, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x70,
	0x65, 0x55, 0x72, 0x69, 0x12, 0x49, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x72, 0x61,
	0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x47, 0x0a, 0x0e, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61,
	0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x70, 0x65, 0x5f,
	0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x70, 0x65, 0x55, 0x72,
	0x69, 0x12, 0x3a, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x68, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x49, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x07, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x49, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x61, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61,
	0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe4, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x69,
	0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x51, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x18, 0x56, 0x45, 0x52, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x4e, 0x49, 0x4d, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x0b,
	0x0a, 0x07, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x55, 0x4d, 0x10, 0x03, 0x2a, 0x3e, 0x0a, 0x0c, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x41,
	0x52, 0x43, 0x48, 0x49, 0x54, 0x45, 0x43, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x38, 0x36,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x36, 0x34, 0x10, 0x02, 0x42, 0x7a, 0x0a, 0x16, 0x69,
	0x6f, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x70, 0x6b, 0x67, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x3b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0xa2, 0x02, 0x03, 0x47, 0x52, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescOnce sync.Once
	file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescData = file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDesc
)

func file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescGZIP() []byte {
	file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescOnce.Do(func() {
		file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescData)
	})
	return file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDescData
}

var file_google_devtools_containeranalysis_v1beta1_package_package_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_devtools_containeranalysis_v1beta1_package_package_proto_goTypes = []interface{}{
	(Architecture)(0),        // 0: grafeas.v1beta1.package.Architecture
	(Version_VersionKind)(0), // 1: grafeas.v1beta1.package.Version.VersionKind
	(*Distribution)(nil),     // 2: grafeas.v1beta1.package.Distribution
	(*Location)(nil),         // 3: grafeas.v1beta1.package.Location
	(*Package)(nil),          // 4: grafeas.v1beta1.package.Package
	(*Details)(nil),          // 5: grafeas.v1beta1.package.Details
	(*Installation)(nil),     // 6: grafeas.v1beta1.package.Installation
	(*Version)(nil),          // 7: grafeas.v1beta1.package.Version
}
var file_google_devtools_containeranalysis_v1beta1_package_package_proto_depIdxs = []int32{
	0, // 0: grafeas.v1beta1.package.Distribution.architecture:type_name -> grafeas.v1beta1.package.Architecture
	7, // 1: grafeas.v1beta1.package.Distribution.latest_version:type_name -> grafeas.v1beta1.package.Version
	7, // 2: grafeas.v1beta1.package.Location.version:type_name -> grafeas.v1beta1.package.Version
	2, // 3: grafeas.v1beta1.package.Package.distribution:type_name -> grafeas.v1beta1.package.Distribution
	6, // 4: grafeas.v1beta1.package.Details.installation:type_name -> grafeas.v1beta1.package.Installation
	3, // 5: grafeas.v1beta1.package.Installation.location:type_name -> grafeas.v1beta1.package.Location
	1, // 6: grafeas.v1beta1.package.Version.kind:type_name -> grafeas.v1beta1.package.Version.VersionKind
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_devtools_containeranalysis_v1beta1_package_package_proto_init() }
func file_google_devtools_containeranalysis_v1beta1_package_package_proto_init() {
	if File_google_devtools_containeranalysis_v1beta1_package_package_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Distribution); i {
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
		file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
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
		file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Installation); i {
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
		file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
			RawDescriptor: file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_containeranalysis_v1beta1_package_package_proto_goTypes,
		DependencyIndexes: file_google_devtools_containeranalysis_v1beta1_package_package_proto_depIdxs,
		EnumInfos:         file_google_devtools_containeranalysis_v1beta1_package_package_proto_enumTypes,
		MessageInfos:      file_google_devtools_containeranalysis_v1beta1_package_package_proto_msgTypes,
	}.Build()
	File_google_devtools_containeranalysis_v1beta1_package_package_proto = out.File
	file_google_devtools_containeranalysis_v1beta1_package_package_proto_rawDesc = nil
	file_google_devtools_containeranalysis_v1beta1_package_package_proto_goTypes = nil
	file_google_devtools_containeranalysis_v1beta1_package_package_proto_depIdxs = nil
}
