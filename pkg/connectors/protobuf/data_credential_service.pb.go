// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.1
// source: data_credential_service.proto

package protobuf

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_data_credential_service_proto protoreflect.FileDescriptor

var file_data_credential_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x1d, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x76, 0x0a, 0x15, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x22, 0x00, 0x42, 0x47, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x6d, 0x65, 0x73,
	0x68, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x62,
	0x6d, 0x2f, 0x74, 0x68, 0x65, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x66, 0x6f, 0x72, 0x2d, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_data_credential_service_proto_goTypes = []interface{}{
	(*DatasetCredentialsRequest)(nil), // 0: connectors.DatasetCredentialsRequest
	(*DatasetCredentials)(nil),        // 1: connectors.DatasetCredentials
}
var file_data_credential_service_proto_depIdxs = []int32{
	0, // 0: connectors.DataCredentialService.GetCredentialsInfo:input_type -> connectors.DatasetCredentialsRequest
	1, // 1: connectors.DataCredentialService.GetCredentialsInfo:output_type -> connectors.DatasetCredentials
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_credential_service_proto_init() }
func file_data_credential_service_proto_init() {
	if File_data_credential_service_proto != nil {
		return
	}
	file_data_credential_request_proto_init()
	file_data_credential_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_credential_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_credential_service_proto_goTypes,
		DependencyIndexes: file_data_credential_service_proto_depIdxs,
	}.Build()
	File_data_credential_service_proto = out.File
	file_data_credential_service_proto_rawDesc = nil
	file_data_credential_service_proto_goTypes = nil
	file_data_credential_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DataCredentialServiceClient is the client API for DataCredentialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataCredentialServiceClient interface {
	GetCredentialsInfo(ctx context.Context, in *DatasetCredentialsRequest, opts ...grpc.CallOption) (*DatasetCredentials, error)
}

type dataCredentialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataCredentialServiceClient(cc grpc.ClientConnInterface) DataCredentialServiceClient {
	return &dataCredentialServiceClient{cc}
}

func (c *dataCredentialServiceClient) GetCredentialsInfo(ctx context.Context, in *DatasetCredentialsRequest, opts ...grpc.CallOption) (*DatasetCredentials, error) {
	out := new(DatasetCredentials)
	err := c.cc.Invoke(ctx, "/connectors.DataCredentialService/GetCredentialsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataCredentialServiceServer is the server API for DataCredentialService service.
type DataCredentialServiceServer interface {
	GetCredentialsInfo(context.Context, *DatasetCredentialsRequest) (*DatasetCredentials, error)
}

// UnimplementedDataCredentialServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataCredentialServiceServer struct {
}

func (*UnimplementedDataCredentialServiceServer) GetCredentialsInfo(context.Context, *DatasetCredentialsRequest) (*DatasetCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredentialsInfo not implemented")
}

func RegisterDataCredentialServiceServer(s *grpc.Server, srv DataCredentialServiceServer) {
	s.RegisterService(&_DataCredentialService_serviceDesc, srv)
}

func _DataCredentialService_GetCredentialsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatasetCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataCredentialServiceServer).GetCredentialsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.DataCredentialService/GetCredentialsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataCredentialServiceServer).GetCredentialsInfo(ctx, req.(*DatasetCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataCredentialService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connectors.DataCredentialService",
	HandlerType: (*DataCredentialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCredentialsInfo",
			Handler:    _DataCredentialService_GetCredentialsInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data_credential_service.proto",
}