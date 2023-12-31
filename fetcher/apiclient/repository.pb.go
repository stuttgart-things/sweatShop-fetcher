// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: repository.proto

package apiclient

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

// Repository is a repository holding application configurations
type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Repo contains the URL to the remote repository
	Repo string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	// Username contains the user name used for authenticating at the remote repository
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// Password contains the password or PAT used for authenticating at the remote repository
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// Insecure specifies whether the connection to the repository ignores any errors when verifying TLS certificates or SSH host keys
	Insecure *bool `protobuf:"varint,4,opt,name=insecure,proto3,oneof" json:"insecure,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{0}
}

func (x *Repository) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *Repository) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Repository) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Repository) GetInsecure() bool {
	if x != nil && x.Insecure != nil {
		return *x.Insecure
	}
	return false
}

// ListFilesRequest requests a repository directory structure
type ListFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo     *Repository `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Revision string      `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
	Path     string      `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ListFilesRequest) Reset() {
	*x = ListFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFilesRequest) ProtoMessage() {}

func (x *ListFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFilesRequest.ProtoReflect.Descriptor instead.
func (*ListFilesRequest) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{1}
}

func (x *ListFilesRequest) GetRepo() *Repository {
	if x != nil {
		return x.Repo
	}
	return nil
}

func (x *ListFilesRequest) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *ListFilesRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// FileList returns the contents of the repo of a ListFiles request
type FileList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []string `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *FileList) Reset() {
	*x = FileList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileList) ProtoMessage() {}

func (x *FileList) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileList.ProtoReflect.Descriptor instead.
func (*FileList) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{2}
}

func (x *FileList) GetFiles() []string {
	if x != nil {
		return x.Files
	}
	return nil
}

// TestRepositoryRequest is a query to test repository is valid or not and has valid access.
type TestRepositoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo *Repository `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
}

func (x *TestRepositoryRequest) Reset() {
	*x = TestRepositoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRepositoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRepositoryRequest) ProtoMessage() {}

func (x *TestRepositoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRepositoryRequest.ProtoReflect.Descriptor instead.
func (*TestRepositoryRequest) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{3}
}

func (x *TestRepositoryRequest) GetRepo() *Repository {
	if x != nil {
		return x.Repo
	}
	return nil
}

// TestRepositoryResponse represents the TestRepository response
type TestRepositoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Request to verify the signature when generating the manifests (only for Git repositories)
	VerifiedRepository bool `protobuf:"varint,1,opt,name=verifiedRepository,proto3" json:"verifiedRepository,omitempty"`
}

func (x *TestRepositoryResponse) Reset() {
	*x = TestRepositoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repository_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRepositoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRepositoryResponse) ProtoMessage() {}

func (x *TestRepositoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repository_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRepositoryResponse.ProtoReflect.Descriptor instead.
func (*TestRepositoryResponse) Descriptor() ([]byte, []int) {
	return file_repository_proto_rawDescGZIP(), []int{4}
}

func (x *TestRepositoryResponse) GetVerifiedRepository() bool {
	if x != nil {
		return x.VerifiedRepository
	}
	return false
}

var File_repository_proto protoreflect.FileDescriptor

var file_repository_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x0a,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x69, 0x6e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x65, 0x22, 0x6b, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x72, 0x65, 0x70,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x20, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04,
	0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x04, 0x72, 0x65, 0x70, 0x6f, 0x22, 0x48, 0x0a, 0x16, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x32,
	0xa2, 0x01, 0x0a, 0x0e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x19, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x53, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x1e, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x75, 0x74, 0x74, 0x67, 0x61, 0x72, 0x74, 0x2d, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x2f, 0x73, 0x77, 0x65, 0x61, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x2d, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_repository_proto_rawDescOnce sync.Once
	file_repository_proto_rawDescData = file_repository_proto_rawDesc
)

func file_repository_proto_rawDescGZIP() []byte {
	file_repository_proto_rawDescOnce.Do(func() {
		file_repository_proto_rawDescData = protoimpl.X.CompressGZIP(file_repository_proto_rawDescData)
	})
	return file_repository_proto_rawDescData
}

var file_repository_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_repository_proto_goTypes = []interface{}{
	(*Repository)(nil),             // 0: fetcher.Repository
	(*ListFilesRequest)(nil),       // 1: fetcher.ListFilesRequest
	(*FileList)(nil),               // 2: fetcher.FileList
	(*TestRepositoryRequest)(nil),  // 3: fetcher.TestRepositoryRequest
	(*TestRepositoryResponse)(nil), // 4: fetcher.TestRepositoryResponse
}
var file_repository_proto_depIdxs = []int32{
	0, // 0: fetcher.ListFilesRequest.repo:type_name -> fetcher.Repository
	0, // 1: fetcher.TestRepositoryRequest.repo:type_name -> fetcher.Repository
	1, // 2: fetcher.FetcherService.ListFiles:input_type -> fetcher.ListFilesRequest
	3, // 3: fetcher.FetcherService.TestRepository:input_type -> fetcher.TestRepositoryRequest
	2, // 4: fetcher.FetcherService.ListFiles:output_type -> fetcher.FileList
	4, // 5: fetcher.FetcherService.TestRepository:output_type -> fetcher.TestRepositoryResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_repository_proto_init() }
func file_repository_proto_init() {
	if File_repository_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repository_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repository); i {
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
		file_repository_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFilesRequest); i {
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
		file_repository_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileList); i {
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
		file_repository_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRepositoryRequest); i {
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
		file_repository_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRepositoryResponse); i {
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
	file_repository_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_repository_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repository_proto_goTypes,
		DependencyIndexes: file_repository_proto_depIdxs,
		MessageInfos:      file_repository_proto_msgTypes,
	}.Build()
	File_repository_proto = out.File
	file_repository_proto_rawDesc = nil
	file_repository_proto_goTypes = nil
	file_repository_proto_depIdxs = nil
}
