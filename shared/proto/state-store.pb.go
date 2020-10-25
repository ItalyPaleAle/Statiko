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
// source: state-store.proto

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

// Certificate type
type TLSCertificate_Type int32

const (
	// Null value (should not be used)
	TLSCertificate_NULL TLSCertificate_Type = 0
	// Imported external certificate
	TLSCertificate_IMPORTED TLSCertificate_Type = 1
	// Self-signed certificate
	TLSCertificate_SELF_SIGNED TLSCertificate_Type = 2
	// ACME (Let's Encrypt)
	TLSCertificate_ACME TLSCertificate_Type = 16
)

// Enum value maps for TLSCertificate_Type.
var (
	TLSCertificate_Type_name = map[int32]string{
		0:  "NULL",
		1:  "IMPORTED",
		2:  "SELF_SIGNED",
		16: "ACME",
	}
	TLSCertificate_Type_value = map[string]int32{
		"NULL":        0,
		"IMPORTED":    1,
		"SELF_SIGNED": 2,
		"ACME":        16,
	}
)

func (x TLSCertificate_Type) Enum() *TLSCertificate_Type {
	p := new(TLSCertificate_Type)
	*p = x
	return p
}

func (x TLSCertificate_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TLSCertificate_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_state_store_proto_enumTypes[0].Descriptor()
}

func (TLSCertificate_Type) Type() protoreflect.EnumType {
	return &file_state_store_proto_enumTypes[0]
}

func (x TLSCertificate_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TLSCertificate_Type.Descriptor instead.
func (TLSCertificate_Type) EnumDescriptor() ([]byte, []int) {
	return file_state_store_proto_rawDescGZIP(), []int{1, 0}
}

// Represents the node's state
type StateStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// State version
	Version uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// List of sites
	Sites []*Site `protobuf:"bytes,2,rep,name=sites,proto3" json:"sites,omitempty"`
	// List of TLS certificates, mapped by their ID
	Certificates map[string]*TLSCertificate `protobuf:"bytes,3,rep,name=certificates,proto3" json:"certificates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Other secrets (encrypted)
	Secrets map[string][]byte `protobuf:"bytes,4,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// DH parameters
	DhParams *DHParams `protobuf:"bytes,5,opt,name=dh_params,json=dhParams,proto3" json:"dh_params,omitempty"`
}

func (x *StateStore) Reset() {
	*x = StateStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateStore) ProtoMessage() {}

func (x *StateStore) ProtoReflect() protoreflect.Message {
	mi := &file_state_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateStore.ProtoReflect.Descriptor instead.
func (*StateStore) Descriptor() ([]byte, []int) {
	return file_state_store_proto_rawDescGZIP(), []int{0}
}

func (x *StateStore) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *StateStore) GetSites() []*Site {
	if x != nil {
		return x.Sites
	}
	return nil
}

func (x *StateStore) GetCertificates() map[string]*TLSCertificate {
	if x != nil {
		return x.Certificates
	}
	return nil
}

func (x *StateStore) GetSecrets() map[string][]byte {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *StateStore) GetDhParams() *DHParams {
	if x != nil {
		return x.DhParams
	}
	return nil
}

// Configuration for a TLS certificate
type TLSCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type TLSCertificate_Type `protobuf:"varint,1,opt,name=type,proto3,enum=statiko.TLSCertificate_Type" json:"type,omitempty"`
	// Certificate key, PEM-encoded and encrypted (used by certain types only)
	Key []byte `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	// Certificate, PEM-encoded (used by certain types only)
	Certificate []byte `protobuf:"bytes,6,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// Auto-generated information
	XDomains []string `protobuf:"bytes,2000,rep,name=_domains,json=Domains,proto3" json:"_domains,omitempty"`
	XNbf     int64    `protobuf:"varint,2001,opt,name=_nbf,json=Nbf,proto3" json:"_nbf,omitempty"`
	XExp     int64    `protobuf:"varint,2002,opt,name=_exp,json=Exp,proto3" json:"_exp,omitempty"`
}

func (x *TLSCertificate) Reset() {
	*x = TLSCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_state_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSCertificate) ProtoMessage() {}

func (x *TLSCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_state_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSCertificate.ProtoReflect.Descriptor instead.
func (*TLSCertificate) Descriptor() ([]byte, []int) {
	return file_state_store_proto_rawDescGZIP(), []int{1}
}

func (x *TLSCertificate) GetType() TLSCertificate_Type {
	if x != nil {
		return x.Type
	}
	return TLSCertificate_NULL
}

func (x *TLSCertificate) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *TLSCertificate) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *TLSCertificate) GetXDomains() []string {
	if x != nil {
		return x.XDomains
	}
	return nil
}

func (x *TLSCertificate) GetXNbf() int64 {
	if x != nil {
		return x.XNbf
	}
	return 0
}

func (x *TLSCertificate) GetXExp() int64 {
	if x != nil {
		return x.XExp
	}
	return 0
}

var File_state_store_proto protoreflect.FileDescriptor

var file_state_store_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x1a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x03, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x69, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2e, 0x53, 0x69, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x69, 0x74, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x0c, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x12,
	0x2e, 0x0a, 0x09, 0x64, 0x68, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2e, 0x44, 0x48, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x08, 0x64, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x58, 0x0a, 0x11, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2e,
	0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf5, 0x01, 0x0a, 0x0e, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f,
	0x2e, 0x54, 0x4c, 0x53, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0xd0, 0x0f, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x5f, 0x6e,
	0x62, 0x66, 0x18, 0xd1, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4e, 0x62, 0x66, 0x12, 0x12,
	0x0a, 0x04, 0x5f, 0x65, 0x78, 0x70, 0x18, 0xd2, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x45,
	0x78, 0x70, 0x22, 0x39, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55,
	0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x4c, 0x46, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x43, 0x4d, 0x45, 0x10, 0x10, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6b, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6b, 0x6f, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_state_store_proto_rawDescOnce sync.Once
	file_state_store_proto_rawDescData = file_state_store_proto_rawDesc
)

func file_state_store_proto_rawDescGZIP() []byte {
	file_state_store_proto_rawDescOnce.Do(func() {
		file_state_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_state_store_proto_rawDescData)
	})
	return file_state_store_proto_rawDescData
}

var file_state_store_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_state_store_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_state_store_proto_goTypes = []interface{}{
	(TLSCertificate_Type)(0), // 0: statiko.TLSCertificate.Type
	(*StateStore)(nil),       // 1: statiko.StateStore
	(*TLSCertificate)(nil),   // 2: statiko.TLSCertificate
	nil,                      // 3: statiko.StateStore.CertificatesEntry
	nil,                      // 4: statiko.StateStore.SecretsEntry
	(*Site)(nil),             // 5: statiko.Site
	(*DHParams)(nil),         // 6: statiko.DHParams
}
var file_state_store_proto_depIdxs = []int32{
	5, // 0: statiko.StateStore.sites:type_name -> statiko.Site
	3, // 1: statiko.StateStore.certificates:type_name -> statiko.StateStore.CertificatesEntry
	4, // 2: statiko.StateStore.secrets:type_name -> statiko.StateStore.SecretsEntry
	6, // 3: statiko.StateStore.dh_params:type_name -> statiko.DHParams
	0, // 4: statiko.TLSCertificate.type:type_name -> statiko.TLSCertificate.Type
	2, // 5: statiko.StateStore.CertificatesEntry.value:type_name -> statiko.TLSCertificate
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_state_store_proto_init() }
func file_state_store_proto_init() {
	if File_state_store_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_state_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateStore); i {
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
		file_state_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSCertificate); i {
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
			RawDescriptor: file_state_store_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_state_store_proto_goTypes,
		DependencyIndexes: file_state_store_proto_depIdxs,
		EnumInfos:         file_state_store_proto_enumTypes,
		MessageInfos:      file_state_store_proto_msgTypes,
	}.Build()
	File_state_store_proto = out.File
	file_state_store_proto_rawDesc = nil
	file_state_store_proto_goTypes = nil
	file_state_store_proto_depIdxs = nil
}