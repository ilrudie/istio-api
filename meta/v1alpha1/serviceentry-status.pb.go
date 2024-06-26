// Copyright 2019 Istio Authors
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
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: meta/v1alpha1/serviceentry-status.proto

// IMPORTANT!!
// IstioServiceEntryStatus IS LARGELY A COPY OF IstioStatus AND SHOULD BE KEPT IN SYNC
// IMPORTANT!!

// $title: Istio ServiceEntry Status
// $description: Status field for Istio ServiceEntry.
// $location: https://istio.io/docs/reference/config/meta/v1beta1/istio-serviceentry-status.html

package v1alpha1

import (
	_ "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1alpha1 "istio.io/api/analysis/v1alpha1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IstioServiceEntryStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current service state of pod.
	// More info: https://istio.io/docs/reference/config/config-status/
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []*IstioCondition `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// Includes any errors or warnings detected by Istio's analyzers.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	ValidationMessages []*v1alpha1.AnalysisMessageBase `protobuf:"bytes,2,rep,name=validation_messages,json=validationMessages,proto3" json:"validation_messages,omitempty"`
	// Resource Generation to which the Reconciled Condition refers.
	// When this value is not equal to the object's metadata generation, reconciled condition  calculation for the current
	// generation is still in progress.  See https://istio.io/latest/docs/reference/config/config-status/ for more info.
	// +optional
	ObservedGeneration int64 `protobuf:"varint,3,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// List of addresses which were automatically allocated by istiod
	// +optional
	Addresses []*IstioServiceEntryAddress `protobuf:"bytes,10,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *IstioServiceEntryStatus) Reset() {
	*x = IstioServiceEntryStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_v1alpha1_serviceentry_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioServiceEntryStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioServiceEntryStatus) ProtoMessage() {}

func (x *IstioServiceEntryStatus) ProtoReflect() protoreflect.Message {
	mi := &file_meta_v1alpha1_serviceentry_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioServiceEntryStatus.ProtoReflect.Descriptor instead.
func (*IstioServiceEntryStatus) Descriptor() ([]byte, []int) {
	return file_meta_v1alpha1_serviceentry_status_proto_rawDescGZIP(), []int{0}
}

func (x *IstioServiceEntryStatus) GetConditions() []*IstioCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *IstioServiceEntryStatus) GetValidationMessages() []*v1alpha1.AnalysisMessageBase {
	if x != nil {
		return x.ValidationMessages
	}
	return nil
}

func (x *IstioServiceEntryStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *IstioServiceEntryStatus) GetAddresses() []*IstioServiceEntryAddress {
	if x != nil {
		return x.Addresses
	}
	return nil
}

// minor abstraction to allow for adding hostnames if relevant
type IstioServiceEntryAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value is the address (192.168.0.2)
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IstioServiceEntryAddress) Reset() {
	*x = IstioServiceEntryAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_v1alpha1_serviceentry_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioServiceEntryAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioServiceEntryAddress) ProtoMessage() {}

func (x *IstioServiceEntryAddress) ProtoReflect() protoreflect.Message {
	mi := &file_meta_v1alpha1_serviceentry_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioServiceEntryAddress.ProtoReflect.Descriptor instead.
func (*IstioServiceEntryAddress) Descriptor() ([]byte, []int) {
	return file_meta_v1alpha1_serviceentry_status_proto_rawDescGZIP(), []int{1}
}

func (x *IstioServiceEntryAddress) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_meta_v1alpha1_serviceentry_status_proto protoreflect.FileDescriptor

var file_meta_v1alpha1_serviceentry_status_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2d, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a,
	0x17, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5d, 0x0a,
	0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x61, 0x73, 0x65, 0x52, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x30, 0x0a, 0x18, 0x49, 0x73,
	0x74, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1c, 0x5a, 0x1a,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_meta_v1alpha1_serviceentry_status_proto_rawDescOnce sync.Once
	file_meta_v1alpha1_serviceentry_status_proto_rawDescData = file_meta_v1alpha1_serviceentry_status_proto_rawDesc
)

func file_meta_v1alpha1_serviceentry_status_proto_rawDescGZIP() []byte {
	file_meta_v1alpha1_serviceentry_status_proto_rawDescOnce.Do(func() {
		file_meta_v1alpha1_serviceentry_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_v1alpha1_serviceentry_status_proto_rawDescData)
	})
	return file_meta_v1alpha1_serviceentry_status_proto_rawDescData
}

var file_meta_v1alpha1_serviceentry_status_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_meta_v1alpha1_serviceentry_status_proto_goTypes = []interface{}{
	(*IstioServiceEntryStatus)(nil),      // 0: istio.meta.v1alpha1.IstioServiceEntryStatus
	(*IstioServiceEntryAddress)(nil),     // 1: istio.meta.v1alpha1.IstioServiceEntryAddress
	(*IstioCondition)(nil),               // 2: istio.meta.v1alpha1.IstioCondition
	(*v1alpha1.AnalysisMessageBase)(nil), // 3: istio.analysis.v1alpha1.AnalysisMessageBase
}
var file_meta_v1alpha1_serviceentry_status_proto_depIdxs = []int32{
	2, // 0: istio.meta.v1alpha1.IstioServiceEntryStatus.conditions:type_name -> istio.meta.v1alpha1.IstioCondition
	3, // 1: istio.meta.v1alpha1.IstioServiceEntryStatus.validation_messages:type_name -> istio.analysis.v1alpha1.AnalysisMessageBase
	1, // 2: istio.meta.v1alpha1.IstioServiceEntryStatus.addresses:type_name -> istio.meta.v1alpha1.IstioServiceEntryAddress
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_meta_v1alpha1_serviceentry_status_proto_init() }
func file_meta_v1alpha1_serviceentry_status_proto_init() {
	if File_meta_v1alpha1_serviceentry_status_proto != nil {
		return
	}
	file_meta_v1alpha1_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meta_v1alpha1_serviceentry_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioServiceEntryStatus); i {
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
		file_meta_v1alpha1_serviceentry_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioServiceEntryAddress); i {
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
			RawDescriptor: file_meta_v1alpha1_serviceentry_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_v1alpha1_serviceentry_status_proto_goTypes,
		DependencyIndexes: file_meta_v1alpha1_serviceentry_status_proto_depIdxs,
		MessageInfos:      file_meta_v1alpha1_serviceentry_status_proto_msgTypes,
	}.Build()
	File_meta_v1alpha1_serviceentry_status_proto = out.File
	file_meta_v1alpha1_serviceentry_status_proto_rawDesc = nil
	file_meta_v1alpha1_serviceentry_status_proto_goTypes = nil
	file_meta_v1alpha1_serviceentry_status_proto_depIdxs = nil
}
