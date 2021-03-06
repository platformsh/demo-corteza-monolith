// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

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

type MakeJWTRequest struct {
	UserID               uint64   `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakeJWTRequest) Reset()         { *m = MakeJWTRequest{} }
func (m *MakeJWTRequest) String() string { return proto.CompactTextString(m) }
func (*MakeJWTRequest) ProtoMessage()    {}
func (*MakeJWTRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *MakeJWTRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeJWTRequest.Unmarshal(m, b)
}
func (m *MakeJWTRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeJWTRequest.Marshal(b, m, deterministic)
}
func (m *MakeJWTRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeJWTRequest.Merge(m, src)
}
func (m *MakeJWTRequest) XXX_Size() int {
	return xxx_messageInfo_MakeJWTRequest.Size(m)
}
func (m *MakeJWTRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeJWTRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MakeJWTRequest proto.InternalMessageInfo

func (m *MakeJWTRequest) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type MakeJWTResponse struct {
	JWT                  string   `protobuf:"bytes,1,opt,name=JWT,proto3" json:"JWT,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MakeJWTResponse) Reset()         { *m = MakeJWTResponse{} }
func (m *MakeJWTResponse) String() string { return proto.CompactTextString(m) }
func (*MakeJWTResponse) ProtoMessage()    {}
func (*MakeJWTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *MakeJWTResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MakeJWTResponse.Unmarshal(m, b)
}
func (m *MakeJWTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MakeJWTResponse.Marshal(b, m, deterministic)
}
func (m *MakeJWTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MakeJWTResponse.Merge(m, src)
}
func (m *MakeJWTResponse) XXX_Size() int {
	return xxx_messageInfo_MakeJWTResponse.Size(m)
}
func (m *MakeJWTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MakeJWTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MakeJWTResponse proto.InternalMessageInfo

func (m *MakeJWTResponse) GetJWT() string {
	if m != nil {
		return m.JWT
	}
	return ""
}

type FindByIDRequest struct {
	UserID               uint64   `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByIDRequest) Reset()         { *m = FindByIDRequest{} }
func (m *FindByIDRequest) String() string { return proto.CompactTextString(m) }
func (*FindByIDRequest) ProtoMessage()    {}
func (*FindByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *FindByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIDRequest.Unmarshal(m, b)
}
func (m *FindByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIDRequest.Marshal(b, m, deterministic)
}
func (m *FindByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIDRequest.Merge(m, src)
}
func (m *FindByIDRequest) XXX_Size() int {
	return xxx_messageInfo_FindByIDRequest.Size(m)
}
func (m *FindByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIDRequest proto.InternalMessageInfo

func (m *FindByIDRequest) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type FindByIDResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByIDResponse) Reset()         { *m = FindByIDResponse{} }
func (m *FindByIDResponse) String() string { return proto.CompactTextString(m) }
func (*FindByIDResponse) ProtoMessage()    {}
func (*FindByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *FindByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIDResponse.Unmarshal(m, b)
}
func (m *FindByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIDResponse.Marshal(b, m, deterministic)
}
func (m *FindByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIDResponse.Merge(m, src)
}
func (m *FindByIDResponse) XXX_Size() int {
	return xxx_messageInfo_FindByIDResponse.Size(m)
}
func (m *FindByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIDResponse proto.InternalMessageInfo

func (m *FindByIDResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	ID                   uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Handle               string   `protobuf:"bytes,3,opt,name=handle,proto3" json:"handle,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Kind                 string   `protobuf:"bytes,5,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
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

func (m *User) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func init() {
	proto.RegisterType((*MakeJWTRequest)(nil), "system.MakeJWTRequest")
	proto.RegisterType((*MakeJWTResponse)(nil), "system.MakeJWTResponse")
	proto.RegisterType((*FindByIDRequest)(nil), "system.FindByIDRequest")
	proto.RegisterType((*FindByIDResponse)(nil), "system.FindByIDResponse")
	proto.RegisterType((*User)(nil), "system.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x49, 0x7a, 0x49, 0xf5, 0x29, 0x6d, 0x79, 0x48, 0x3d, 0x3a, 0x95, 0xb8, 0xd4, 0x25,
	0x43, 0x75, 0x12, 0x5c, 0x4a, 0x10, 0x52, 0x70, 0x09, 0x95, 0x82, 0x5b, 0x24, 0x0f, 0x0c, 0x4d,
	0x2e, 0x35, 0x97, 0x0e, 0x5d, 0xfd, 0xcb, 0xe5, 0x5e, 0xae, 0x11, 0xcd, 0xd0, 0x29, 0xef, 0xfb,
	0x7d, 0x9f, 0xbc, 0x5f, 0x07, 0x70, 0xd0, 0x54, 0x87, 0xfb, 0xba, 0x6a, 0x2a, 0xf4, 0xf5, 0x51,
	0x37, 0x54, 0x06, 0x0b, 0x18, 0xbd, 0xa6, 0x3b, 0x5a, 0x6f, 0x37, 0x09, 0x7d, 0x1d, 0x48, 0x37,
	0x38, 0x05, 0xdf, 0x70, 0x71, 0x24, 0x9d, 0xb9, 0xb3, 0x10, 0x89, 0x55, 0xc1, 0x1d, 0x8c, 0x3b,
	0x52, 0xef, 0x2b, 0xa5, 0x09, 0x27, 0x30, 0x58, 0x6f, 0x37, 0xcc, 0x5d, 0x26, 0x26, 0x0c, 0xee,
	0x61, 0xfc, 0x92, 0xab, 0x6c, 0x75, 0x8c, 0xa3, 0x73, 0xf5, 0x1e, 0x61, 0xf2, 0x8b, 0xda, 0x82,
	0x73, 0x10, 0x26, 0xcb, 0xe4, 0xd5, 0xf2, 0x3a, 0x6c, 0x87, 0x0c, 0xdf, 0x34, 0xd5, 0x09, 0x67,
	0x82, 0x02, 0x84, 0x51, 0x38, 0x02, 0xb7, 0xab, 0xe8, 0xc6, 0x11, 0xde, 0x80, 0x47, 0x65, 0x9a,
	0x17, 0xd2, 0xe5, 0x61, 0x5a, 0x61, 0x7a, 0x7f, 0xa6, 0x2a, 0x2b, 0x48, 0x0e, 0xd8, 0xb6, 0x0a,
	0x11, 0x84, 0x4a, 0x4b, 0x92, 0x82, 0x5d, 0x8e, 0x8d, 0xb7, 0xcb, 0x55, 0x26, 0xbd, 0xd6, 0x33,
	0xf1, 0xf2, 0xdb, 0x01, 0xcf, 0xb4, 0xd3, 0xf8, 0x04, 0x43, 0xbb, 0x3d, 0x4e, 0x4f, 0x63, 0xfd,
	0x3d, 0xdc, 0xec, 0xb6, 0xe7, 0xdb, 0xad, 0x9e, 0xe1, 0xe2, 0xb4, 0x29, 0x76, 0xd0, 0xbf, 0x33,
	0xcd, 0x64, 0x3f, 0xd1, 0xfe, 0xbe, 0x1a, 0xbe, 0x7b, 0xfc, 0x66, 0x1f, 0x3e, 0x7f, 0x1e, 0x7e,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x04, 0x91, 0x4b, 0x16, 0xc8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	MakeJWT(ctx context.Context, in *MakeJWTRequest, opts ...grpc.CallOption) (*MakeJWTResponse, error)
	FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDResponse, error)
}

type usersClient struct {
	cc *grpc.ClientConn
}

func NewUsersClient(cc *grpc.ClientConn) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) MakeJWT(ctx context.Context, in *MakeJWTRequest, opts ...grpc.CallOption) (*MakeJWTResponse, error) {
	out := new(MakeJWTResponse)
	err := c.cc.Invoke(ctx, "/system.Users/MakeJWT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDResponse, error) {
	out := new(FindByIDResponse)
	err := c.cc.Invoke(ctx, "/system.Users/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	MakeJWT(context.Context, *MakeJWTRequest) (*MakeJWTResponse, error)
	FindByID(context.Context, *FindByIDRequest) (*FindByIDResponse, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) MakeJWT(ctx context.Context, req *MakeJWTRequest) (*MakeJWTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeJWT not implemented")
}
func (*UnimplementedUsersServer) FindByID(ctx context.Context, req *FindByIDRequest) (*FindByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_MakeJWT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeJWTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).MakeJWT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.Users/MakeJWT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).MakeJWT(ctx, req.(*MakeJWTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.Users/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).FindByID(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "system.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeJWT",
			Handler:    _Users_MakeJWT_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _Users_FindByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
