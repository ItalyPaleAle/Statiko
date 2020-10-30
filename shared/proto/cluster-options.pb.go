//
//Copyright © 2020 Alessandro Segala (@ItalyPaleAle)
//
//This program is free software: you can redistribute it and/or modify
//it under the terms of the GNU Affero General Public License as published
//by the Free Software Foundation, version 3 of the License.
//
//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU Affero General Public License for more details.
//
//You should have received a copy of the GNU Affero General Public License
//along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: cluster-options.proto

package proto

import (
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

// Type of codesign key
type ClusterOptions_Codesign_KeyType int32

const (
	// Null: there's no codesign key
	ClusterOptions_Codesign_NULL ClusterOptions_Codesign_KeyType = 0
	// RSA public key
	ClusterOptions_Codesign_RSA ClusterOptions_Codesign_KeyType = 1
)

// Enum value maps for ClusterOptions_Codesign_KeyType.
var (
	ClusterOptions_Codesign_KeyType_name = map[int32]string{
		0: "NULL",
		1: "RSA",
	}
	ClusterOptions_Codesign_KeyType_value = map[string]int32{
		"NULL": 0,
		"RSA":  1,
	}
)

func (x ClusterOptions_Codesign_KeyType) Enum() *ClusterOptions_Codesign_KeyType {
	p := new(ClusterOptions_Codesign_KeyType)
	*p = x
	return p
}

func (x ClusterOptions_Codesign_KeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClusterOptions_Codesign_KeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_cluster_options_proto_enumTypes[0].Descriptor()
}

func (ClusterOptions_Codesign_KeyType) Type() protoreflect.EnumType {
	return &file_cluster_options_proto_enumTypes[0]
}

func (x ClusterOptions_Codesign_KeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClusterOptions_Codesign_KeyType.Descriptor instead.
func (ClusterOptions_Codesign_KeyType) EnumDescriptor() ([]byte, []int) {
	return file_cluster_options_proto_rawDescGZIP(), []int{0, 0, 0}
}

// Message with the cluster options
type ClusterOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Codesign options
	Codesign *ClusterOptions_Codesign `protobuf:"bytes,1,opt,name=codesign,proto3" json:"codesign,omitempty"`
	// Manifest file name
	ManifestFile string `protobuf:"bytes,2,opt,name=manifest_file,json=manifestFile,proto3" json:"manifest_file,omitempty"`
}

func (x *ClusterOptions) Reset() {
	*x = ClusterOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterOptions) ProtoMessage() {}

func (x *ClusterOptions) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterOptions.ProtoReflect.Descriptor instead.
func (*ClusterOptions) Descriptor() ([]byte, []int) {
	return file_cluster_options_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterOptions) GetCodesign() *ClusterOptions_Codesign {
	if x != nil {
		return x.Codesign
	}
	return nil
}

func (x *ClusterOptions) GetManifestFile() string {
	if x != nil {
		return x.ManifestFile
	}
	return ""
}

// Codesign options message
type ClusterOptions_Codesign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requires code signing
	RequireCodesign bool `protobuf:"varint,1,opt,name=require_codesign,json=requireCodesign,proto3" json:"require_codesign,omitempty"`
	// Key type
	Type ClusterOptions_Codesign_KeyType `protobuf:"varint,2,opt,name=type,proto3,enum=statiko.ClusterOptions_Codesign_KeyType" json:"type,omitempty"`
	// RSA key when key type is RSA
	RsaKey *ClusterOptions_Codesign_RSAKey `protobuf:"bytes,10,opt,name=rsa_key,json=rsaKey,proto3" json:"rsa_key,omitempty"`
}

func (x *ClusterOptions_Codesign) Reset() {
	*x = ClusterOptions_Codesign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterOptions_Codesign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterOptions_Codesign) ProtoMessage() {}

func (x *ClusterOptions_Codesign) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterOptions_Codesign.ProtoReflect.Descriptor instead.
func (*ClusterOptions_Codesign) Descriptor() ([]byte, []int) {
	return file_cluster_options_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ClusterOptions_Codesign) GetRequireCodesign() bool {
	if x != nil {
		return x.RequireCodesign
	}
	return false
}

func (x *ClusterOptions_Codesign) GetType() ClusterOptions_Codesign_KeyType {
	if x != nil {
		return x.Type
	}
	return ClusterOptions_Codesign_NULL
}

func (x *ClusterOptions_Codesign) GetRsaKey() *ClusterOptions_Codesign_RSAKey {
	if x != nil {
		return x.RsaKey
	}
	return nil
}

// Arguments for RSA key
type ClusterOptions_Codesign_RSAKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Modulus
	N []byte `protobuf:"bytes,1,opt,name=n,proto3" json:"n,omitempty"`
	// Exponent
	// Note on using 32-bit integers: https://www.imperialviolet.org/2012/03/16/rsae.html
	E uint32 `protobuf:"varint,2,opt,name=e,proto3" json:"e,omitempty"`
}

func (x *ClusterOptions_Codesign_RSAKey) Reset() {
	*x = ClusterOptions_Codesign_RSAKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterOptions_Codesign_RSAKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterOptions_Codesign_RSAKey) ProtoMessage() {}

func (x *ClusterOptions_Codesign_RSAKey) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterOptions_Codesign_RSAKey.ProtoReflect.Descriptor instead.
func (*ClusterOptions_Codesign_RSAKey) Descriptor() ([]byte, []int) {
	return file_cluster_options_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ClusterOptions_Codesign_RSAKey) GetN() []byte {
	if x != nil {
		return x.N
	}
	return nil
}

func (x *ClusterOptions_Codesign_RSAKey) GetE() uint32 {
	if x != nil {
		return x.E
	}
	return 0
}

var File_cluster_options_proto protoreflect.FileDescriptor

var file_cluster_options_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f,
	0x22, 0xef, 0x02, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x08, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0xf9, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x64, 0x65, 0x73,
	0x69, 0x67, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x3c,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x4b,
	0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x07,
	0x72, 0x73, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x2e,
	0x52, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x72, 0x73, 0x61, 0x4b, 0x65, 0x79, 0x1a, 0x24,
	0x0a, 0x06, 0x52, 0x53, 0x41, 0x4b, 0x65, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x01, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x01, 0x65, 0x22, 0x1c, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x53, 0x41,
	0x10, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6b, 0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cluster_options_proto_rawDescOnce sync.Once
	file_cluster_options_proto_rawDescData = file_cluster_options_proto_rawDesc
)

func file_cluster_options_proto_rawDescGZIP() []byte {
	file_cluster_options_proto_rawDescOnce.Do(func() {
		file_cluster_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_cluster_options_proto_rawDescData)
	})
	return file_cluster_options_proto_rawDescData
}

var file_cluster_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cluster_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cluster_options_proto_goTypes = []interface{}{
	(ClusterOptions_Codesign_KeyType)(0),   // 0: statiko.ClusterOptions.Codesign.KeyType
	(*ClusterOptions)(nil),                 // 1: statiko.ClusterOptions
	(*ClusterOptions_Codesign)(nil),        // 2: statiko.ClusterOptions.Codesign
	(*ClusterOptions_Codesign_RSAKey)(nil), // 3: statiko.ClusterOptions.Codesign.RSAKey
}
var file_cluster_options_proto_depIdxs = []int32{
	2, // 0: statiko.ClusterOptions.codesign:type_name -> statiko.ClusterOptions.Codesign
	0, // 1: statiko.ClusterOptions.Codesign.type:type_name -> statiko.ClusterOptions.Codesign.KeyType
	3, // 2: statiko.ClusterOptions.Codesign.rsa_key:type_name -> statiko.ClusterOptions.Codesign.RSAKey
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cluster_options_proto_init() }
func file_cluster_options_proto_init() {
	if File_cluster_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cluster_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterOptions); i {
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
		file_cluster_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterOptions_Codesign); i {
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
		file_cluster_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterOptions_Codesign_RSAKey); i {
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
			RawDescriptor: file_cluster_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cluster_options_proto_goTypes,
		DependencyIndexes: file_cluster_options_proto_depIdxs,
		EnumInfos:         file_cluster_options_proto_enumTypes,
		MessageInfos:      file_cluster_options_proto_msgTypes,
	}.Build()
	File_cluster_options_proto = out.File
	file_cluster_options_proto_rawDesc = nil
	file_cluster_options_proto_goTypes = nil
	file_cluster_options_proto_depIdxs = nil
}