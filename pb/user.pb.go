// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Timestamp from public import google/protobuf/timestamp.proto
type Timestamp = timestamp.Timestamp

// Empty from public import google/protobuf/empty.proto
type Empty = empty.Empty

type User struct {
	Username             string               `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type RegisterUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterUserRequest) Reset()         { *m = RegisterUserRequest{} }
func (m *RegisterUserRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterUserRequest) ProtoMessage()    {}
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *RegisterUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterUserRequest.Unmarshal(m, b)
}
func (m *RegisterUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterUserRequest.Marshal(b, m, deterministic)
}
func (m *RegisterUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterUserRequest.Merge(m, src)
}
func (m *RegisterUserRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterUserRequest.Size(m)
}
func (m *RegisterUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterUserRequest proto.InternalMessageInfo

func (m *RegisterUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserResponse struct {
	Username             string               `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type UserStreamRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserStreamRequest) Reset()         { *m = UserStreamRequest{} }
func (m *UserStreamRequest) String() string { return proto.CompactTextString(m) }
func (*UserStreamRequest) ProtoMessage()    {}
func (*UserStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *UserStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserStreamRequest.Unmarshal(m, b)
}
func (m *UserStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserStreamRequest.Marshal(b, m, deterministic)
}
func (m *UserStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserStreamRequest.Merge(m, src)
}
func (m *UserStreamRequest) XXX_Size() int {
	return xxx_messageInfo_UserStreamRequest.Size(m)
}
func (m *UserStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserStreamRequest proto.InternalMessageInfo

func (m *UserStreamRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserListResponse struct {
	User                 []*User  `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*RegisterUserRequest)(nil), "user.RegisterUserRequest")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
	proto.RegisterType((*UserStreamRequest)(nil), "user.UserStreamRequest")
	proto.RegisterType((*UserListResponse)(nil), "user.UserListResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4f, 0x4f, 0x3a, 0x31,
	0x10, 0xa5, 0x3f, 0xf8, 0x99, 0x65, 0x30, 0x46, 0xab, 0xc1, 0xb5, 0x26, 0x4a, 0x7a, 0xe2, 0xb4,
	0x24, 0xeb, 0xc9, 0x68, 0x62, 0x3c, 0x78, 0xf3, 0x40, 0x56, 0x3d, 0x9b, 0x02, 0x23, 0xd9, 0x84,
	0xd2, 0xb5, 0x2d, 0x1a, 0xe3, 0x27, 0xf4, 0x5b, 0x99, 0xb6, 0x14, 0x50, 0x39, 0xf8, 0xe7, 0xd6,
	0x99, 0x79, 0x33, 0xef, 0xbd, 0x99, 0x02, 0xcc, 0x0c, 0xea, 0xac, 0xd2, 0xca, 0x2a, 0xda, 0x70,
	0x6f, 0x76, 0x3c, 0x56, 0x6a, 0x3c, 0xc1, 0x9e, 0xcf, 0x0d, 0x66, 0x0f, 0x3d, 0x5b, 0x4a, 0x34,
	0x56, 0xc8, 0x2a, 0xc0, 0xd8, 0xe1, 0x67, 0x00, 0xca, 0xca, 0xbe, 0x84, 0x22, 0x37, 0xd0, 0xb8,
	0x33, 0xa8, 0x29, 0x83, 0xc4, 0x4d, 0x9b, 0x0a, 0x89, 0x29, 0xe9, 0x90, 0x6e, 0xb3, 0x58, 0xc4,
	0x74, 0x0f, 0xfe, 0xa3, 0x14, 0xe5, 0x24, 0xfd, 0xe7, 0x0b, 0x21, 0xa0, 0xa7, 0x00, 0x43, 0x8d,
	0xc2, 0xe2, 0xe8, 0x5e, 0xd8, 0xb4, 0xd1, 0x21, 0xdd, 0x56, 0xce, 0xb2, 0xc0, 0x95, 0x45, 0xae,
	0xec, 0x36, 0x8a, 0x29, 0x9a, 0x73, 0xf4, 0xa5, 0xe5, 0x43, 0xd8, 0x2d, 0x70, 0x5c, 0x1a, 0x8b,
	0xda, 0x91, 0x17, 0xf8, 0x38, 0x43, 0x63, 0x7f, 0xa1, 0x81, 0x41, 0x52, 0x09, 0x63, 0x9e, 0x95,
	0x1e, 0xa5, 0xf5, 0xd0, 0x11, 0x63, 0xfe, 0x0a, 0x9b, 0x61, 0xb8, 0xa9, 0xd4, 0xd4, 0xe0, 0x9f,
	0x1d, 0xd6, 0x7f, 0xe2, 0xb0, 0x07, 0x3b, 0x8e, 0xfc, 0xc6, 0x6a, 0x14, 0xf2, 0x1b, 0xfe, 0x78,
	0x0e, 0xdb, 0xae, 0xe1, 0xba, 0x34, 0x76, 0xa1, 0xf8, 0x08, 0xfc, 0x85, 0x53, 0xd2, 0xa9, 0x77,
	0x5b, 0x39, 0x64, 0xfe, 0xf4, 0xde, 0x93, 0xcf, 0xe7, 0x6f, 0x04, 0x5a, 0x9e, 0x05, 0xf5, 0x53,
	0x39, 0x44, 0x7a, 0x06, 0x49, 0x5c, 0x2b, 0x3d, 0x08, 0xe8, 0x35, 0x6b, 0x66, 0x74, 0x65, 0xd0,
	0x9c, 0x8a, 0xd7, 0xe8, 0x39, 0x24, 0x51, 0x00, 0x6d, 0x7f, 0x31, 0x79, 0xe5, 0xbe, 0x0c, 0x6b,
	0x2f, 0x3b, 0x57, 0x85, 0xf2, 0x1a, 0xbd, 0x80, 0xad, 0x98, 0x0d, 0x9e, 0xe9, 0xfe, 0x12, 0xfb,
	0x61, 0x0b, 0xeb, 0xe9, 0xfb, 0xb5, 0x3e, 0x19, 0x6c, 0x78, 0xba, 0x93, 0xf7, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x6d, 0xf6, 0x01, 0x48, 0xe4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	//注册用户
	Register(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	//用户列表(无参数)
	UserList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserListResponse, error)
	//用户流(服务端推送流)
	UserListStream(ctx context.Context, in *UserStreamRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserListStream(ctx context.Context, in *UserStreamRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UserListStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	//注册用户
	Register(context.Context, *RegisterUserRequest) (*UserResponse, error)
	//用户列表(无参数)
	UserList(context.Context, *empty.Empty) (*UserListResponse, error)
	//用户流(服务端推送流)
	UserListStream(context.Context, *UserStreamRequest) (*UserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Register(ctx context.Context, req *RegisterUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUserServiceServer) UserList(ctx context.Context, req *empty.Empty) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (*UnimplementedUserServiceServer) UserListStream(ctx context.Context, req *UserStreamRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserListStream not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserListStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserListStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UserListStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserListStream(ctx, req.(*UserStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _UserService_UserList_Handler,
		},
		{
			MethodName: "UserListStream",
			Handler:    _UserService_UserListStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
