// Code generated by protoc-gen-go.
// source: pods.proto
// DO NOT EDIT!

/*
Package pods is a generated protocol buffer package.

It is generated from these files:
	pods.proto

It has these top-level messages:
	CreatePodRequest
	CreatePodResponse
	DeletePodRequest
	DeletePodResponse
	ListPodsRequest
	ListPodsResponse
	StdinStreamRequest
	StdoutStreamResponse
	Pod
	PodSpec
	Container
	Mount
	PodStatus
	ContainerStatus
	ResourceMetadata
*/
package pods

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreatePodRequest struct {
	Pod *Pod `protobuf:"bytes,1,opt,name=pod" json:"pod,omitempty"`
	Tty bool `protobuf:"varint,2,opt,name=tty" json:"tty,omitempty"`
}

func (m *CreatePodRequest) Reset()                    { *m = CreatePodRequest{} }
func (m *CreatePodRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePodRequest) ProtoMessage()               {}
func (*CreatePodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreatePodRequest) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

func (m *CreatePodRequest) GetTty() bool {
	if m != nil {
		return m.Tty
	}
	return false
}

type CreatePodResponse struct {
	Pod *Pod `protobuf:"bytes,1,opt,name=pod" json:"pod,omitempty"`
}

func (m *CreatePodResponse) Reset()                    { *m = CreatePodResponse{} }
func (m *CreatePodResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePodResponse) ProtoMessage()               {}
func (*CreatePodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreatePodResponse) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

type DeletePodRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *DeletePodRequest) Reset()                    { *m = DeletePodRequest{} }
func (m *DeletePodRequest) String() string            { return proto.CompactTextString(m) }
func (*DeletePodRequest) ProtoMessage()               {}
func (*DeletePodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeletePodRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeletePodRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DeletePodResponse struct {
	Pod *Pod `protobuf:"bytes,1,opt,name=pod" json:"pod,omitempty"`
}

func (m *DeletePodResponse) Reset()                    { *m = DeletePodResponse{} }
func (m *DeletePodResponse) String() string            { return proto.CompactTextString(m) }
func (*DeletePodResponse) ProtoMessage()               {}
func (*DeletePodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeletePodResponse) GetPod() *Pod {
	if m != nil {
		return m.Pod
	}
	return nil
}

type ListPodsRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ListPodsRequest) Reset()                    { *m = ListPodsRequest{} }
func (m *ListPodsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPodsRequest) ProtoMessage()               {}
func (*ListPodsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListPodsRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ListPodsResponse struct {
	Pods []*Pod `protobuf:"bytes,1,rep,name=pods" json:"pods,omitempty"`
}

func (m *ListPodsResponse) Reset()                    { *m = ListPodsResponse{} }
func (m *ListPodsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPodsResponse) ProtoMessage()               {}
func (*ListPodsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListPodsResponse) GetPods() []*Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type StdinStreamRequest struct {
	Input []byte `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
}

func (m *StdinStreamRequest) Reset()                    { *m = StdinStreamRequest{} }
func (m *StdinStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StdinStreamRequest) ProtoMessage()               {}
func (*StdinStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StdinStreamRequest) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

type StdoutStreamResponse struct {
	Output []byte `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	// Is this stderr(=true) or stdout(=false)
	Stderr bool `protobuf:"varint,2,opt,name=stderr" json:"stderr,omitempty"`
}

func (m *StdoutStreamResponse) Reset()                    { *m = StdoutStreamResponse{} }
func (m *StdoutStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*StdoutStreamResponse) ProtoMessage()               {}
func (*StdoutStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *StdoutStreamResponse) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *StdoutStreamResponse) GetStderr() bool {
	if m != nil {
		return m.Stderr
	}
	return false
}

type Pod struct {
	Metadata *ResourceMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Spec     *PodSpec          `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status   *PodStatus        `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *Pod) Reset()                    { *m = Pod{} }
func (m *Pod) String() string            { return proto.CompactTextString(m) }
func (*Pod) ProtoMessage()               {}
func (*Pod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Pod) GetMetadata() *ResourceMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Pod) GetSpec() *PodSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Pod) GetStatus() *PodStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type PodSpec struct {
	Containers  []*Container `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
	HostNetwork bool         `protobuf:"varint,2,opt,name=hostNetwork" json:"hostNetwork,omitempty"`
}

func (m *PodSpec) Reset()                    { *m = PodSpec{} }
func (m *PodSpec) String() string            { return proto.CompactTextString(m) }
func (*PodSpec) ProtoMessage()               {}
func (*PodSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *PodSpec) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *PodSpec) GetHostNetwork() bool {
	if m != nil {
		return m.HostNetwork
	}
	return false
}

type Container struct {
	Name       string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Image      string   `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	Tty        bool     `protobuf:"varint,3,opt,name=tty" json:"tty,omitempty"`
	WorkingDir string   `protobuf:"bytes,4,opt,name=workingDir" json:"workingDir,omitempty"`
	Args       []string `protobuf:"bytes,5,rep,name=args" json:"args,omitempty"`
	Env        []string `protobuf:"bytes,6,rep,name=env" json:"env,omitempty"`
	Mounts     []*Mount `protobuf:"bytes,7,rep,name=mounts" json:"mounts,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Container) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetTty() bool {
	if m != nil {
		return m.Tty
	}
	return false
}

func (m *Container) GetWorkingDir() string {
	if m != nil {
		return m.WorkingDir
	}
	return ""
}

func (m *Container) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Container) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Container) GetMounts() []*Mount {
	if m != nil {
		return m.Mounts
	}
	return nil
}

type Mount struct {
	Type        string   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Source      string   `protobuf:"bytes,2,opt,name=source" json:"source,omitempty"`
	Destination string   `protobuf:"bytes,3,opt,name=destination" json:"destination,omitempty"`
	Options     []string `protobuf:"bytes,4,rep,name=options" json:"options,omitempty"`
}

func (m *Mount) Reset()                    { *m = Mount{} }
func (m *Mount) String() string            { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()               {}
func (*Mount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Mount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Mount) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Mount) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *Mount) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

type PodStatus struct {
	ContainerStatuses []*ContainerStatus `protobuf:"bytes,1,rep,name=containerStatuses" json:"containerStatuses,omitempty"`
}

func (m *PodStatus) Reset()                    { *m = PodStatus{} }
func (m *PodStatus) String() string            { return proto.CompactTextString(m) }
func (*PodStatus) ProtoMessage()               {}
func (*PodStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PodStatus) GetContainerStatuses() []*ContainerStatus {
	if m != nil {
		return m.ContainerStatuses
	}
	return nil
}

type ContainerStatus struct {
	ContainerID string `protobuf:"bytes,1,opt,name=containerID" json:"containerID,omitempty"`
	Image       string `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	State       string `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
}

func (m *ContainerStatus) Reset()                    { *m = ContainerStatus{} }
func (m *ContainerStatus) String() string            { return proto.CompactTextString(m) }
func (*ContainerStatus) ProtoMessage()               {}
func (*ContainerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ContainerStatus) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *ContainerStatus) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ContainerStatus) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

// ResourceMetadata is metadata that all resources must have
type ResourceMetadata struct {
	// Name must be unique within a namespace.
	// Cannot be updated.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Namespace defines space within each name must be unique.
	// An empty namespace is equivalent to the default namespace.
	// Cannot be updated.
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ResourceMetadata) Reset()                    { *m = ResourceMetadata{} }
func (m *ResourceMetadata) String() string            { return proto.CompactTextString(m) }
func (*ResourceMetadata) ProtoMessage()               {}
func (*ResourceMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ResourceMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceMetadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*CreatePodRequest)(nil), "cand.services.pods.v1.CreatePodRequest")
	proto.RegisterType((*CreatePodResponse)(nil), "cand.services.pods.v1.CreatePodResponse")
	proto.RegisterType((*DeletePodRequest)(nil), "cand.services.pods.v1.DeletePodRequest")
	proto.RegisterType((*DeletePodResponse)(nil), "cand.services.pods.v1.DeletePodResponse")
	proto.RegisterType((*ListPodsRequest)(nil), "cand.services.pods.v1.ListPodsRequest")
	proto.RegisterType((*ListPodsResponse)(nil), "cand.services.pods.v1.ListPodsResponse")
	proto.RegisterType((*StdinStreamRequest)(nil), "cand.services.pods.v1.StdinStreamRequest")
	proto.RegisterType((*StdoutStreamResponse)(nil), "cand.services.pods.v1.StdoutStreamResponse")
	proto.RegisterType((*Pod)(nil), "cand.services.pods.v1.Pod")
	proto.RegisterType((*PodSpec)(nil), "cand.services.pods.v1.PodSpec")
	proto.RegisterType((*Container)(nil), "cand.services.pods.v1.Container")
	proto.RegisterType((*Mount)(nil), "cand.services.pods.v1.Mount")
	proto.RegisterType((*PodStatus)(nil), "cand.services.pods.v1.PodStatus")
	proto.RegisterType((*ContainerStatus)(nil), "cand.services.pods.v1.ContainerStatus")
	proto.RegisterType((*ResourceMetadata)(nil), "cand.services.pods.v1.ResourceMetadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Pods service

type PodsClient interface {
	Create(ctx context.Context, in *CreatePodRequest, opts ...grpc.CallOption) (*CreatePodResponse, error)
	Delete(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error)
	List(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error)
	Attach(ctx context.Context, opts ...grpc.CallOption) (Pods_AttachClient, error)
}

type podsClient struct {
	cc *grpc.ClientConn
}

func NewPodsClient(cc *grpc.ClientConn) PodsClient {
	return &podsClient{cc}
}

func (c *podsClient) Create(ctx context.Context, in *CreatePodRequest, opts ...grpc.CallOption) (*CreatePodResponse, error) {
	out := new(CreatePodResponse)
	err := grpc.Invoke(ctx, "/cand.services.pods.v1.Pods/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podsClient) Delete(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error) {
	out := new(DeletePodResponse)
	err := grpc.Invoke(ctx, "/cand.services.pods.v1.Pods/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podsClient) List(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error) {
	out := new(ListPodsResponse)
	err := grpc.Invoke(ctx, "/cand.services.pods.v1.Pods/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podsClient) Attach(ctx context.Context, opts ...grpc.CallOption) (Pods_AttachClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Pods_serviceDesc.Streams[0], c.cc, "/cand.services.pods.v1.Pods/Attach", opts...)
	if err != nil {
		return nil, err
	}
	x := &podsAttachClient{stream}
	return x, nil
}

type Pods_AttachClient interface {
	Send(*StdinStreamRequest) error
	Recv() (*StdoutStreamResponse, error)
	grpc.ClientStream
}

type podsAttachClient struct {
	grpc.ClientStream
}

func (x *podsAttachClient) Send(m *StdinStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *podsAttachClient) Recv() (*StdoutStreamResponse, error) {
	m := new(StdoutStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Pods service

type PodsServer interface {
	Create(context.Context, *CreatePodRequest) (*CreatePodResponse, error)
	Delete(context.Context, *DeletePodRequest) (*DeletePodResponse, error)
	List(context.Context, *ListPodsRequest) (*ListPodsResponse, error)
	Attach(Pods_AttachServer) error
}

func RegisterPodsServer(s *grpc.Server, srv PodsServer) {
	s.RegisterService(&_Pods_serviceDesc, srv)
}

func _Pods_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cand.services.pods.v1.Pods/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodsServer).Create(ctx, req.(*CreatePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pods_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cand.services.pods.v1.Pods/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodsServer).Delete(ctx, req.(*DeletePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pods_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cand.services.pods.v1.Pods/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodsServer).List(ctx, req.(*ListPodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pods_Attach_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PodsServer).Attach(&podsAttachServer{stream})
}

type Pods_AttachServer interface {
	Send(*StdoutStreamResponse) error
	Recv() (*StdinStreamRequest, error)
	grpc.ServerStream
}

type podsAttachServer struct {
	grpc.ServerStream
}

func (x *podsAttachServer) Send(m *StdoutStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *podsAttachServer) Recv() (*StdinStreamRequest, error) {
	m := new(StdinStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Pods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cand.services.pods.v1.Pods",
	HandlerType: (*PodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Pods_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Pods_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Pods_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Attach",
			Handler:       _Pods_Attach_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pods.proto",
}

func init() { proto.RegisterFile("pods.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdf, 0x6f, 0xd3, 0x40,
	0x0c, 0x56, 0xd6, 0x34, 0x5b, 0x3d, 0xa4, 0x75, 0xa7, 0x81, 0xa2, 0x6a, 0x8c, 0x2a, 0x0f, 0xac,
	0xfc, 0x50, 0xcb, 0x0a, 0x0f, 0x48, 0xbc, 0xb0, 0xad, 0x42, 0x42, 0x62, 0x68, 0xba, 0x22, 0x21,
	0xb1, 0x07, 0x74, 0x24, 0xa7, 0x2e, 0x1a, 0xcd, 0x85, 0x9c, 0x53, 0xb4, 0xff, 0x8c, 0x77, 0xfe,
	0x30, 0x90, 0xaf, 0x97, 0x34, 0x6b, 0x17, 0x0a, 0x3c, 0xcd, 0xf6, 0x3e, 0x7f, 0xb6, 0xcf, 0xce,
	0x57, 0x80, 0x54, 0x45, 0xba, 0x9f, 0x66, 0x0a, 0x15, 0xbb, 0x1b, 0x8a, 0x24, 0xea, 0x6b, 0x99,
	0xcd, 0xe2, 0x50, 0xea, 0xbe, 0xf9, 0xcf, 0xec, 0x28, 0xe0, 0xd0, 0x3e, 0xcd, 0xa4, 0x40, 0x79,
	0xae, 0x22, 0x2e, 0xbf, 0xe5, 0x52, 0x23, 0x7b, 0x0a, 0x8d, 0x54, 0x45, 0xbe, 0xd3, 0x75, 0x7a,
	0xdb, 0xc3, 0x4e, 0xff, 0xd6, 0xc4, 0x3e, 0xe1, 0x09, 0xc6, 0xda, 0xd0, 0x40, 0xbc, 0xf6, 0x37,
	0xba, 0x4e, 0x6f, 0x8b, 0x93, 0x19, 0x1c, 0xc3, 0x6e, 0x85, 0x53, 0xa7, 0x2a, 0xd1, 0xf2, 0xdf,
	0x48, 0x83, 0x11, 0xb4, 0x47, 0xf2, 0xab, 0xbc, 0xd1, 0xd6, 0x3e, 0xb4, 0x12, 0x31, 0x95, 0x3a,
	0x15, 0xa1, 0x34, 0x3c, 0x2d, 0xbe, 0x08, 0x30, 0x06, 0x2e, 0x39, 0xa6, 0x8f, 0x16, 0x37, 0x36,
	0x35, 0x52, 0x61, 0xf9, 0xaf, 0x46, 0x06, 0xb0, 0xf3, 0x2e, 0xd6, 0x78, 0xae, 0x22, 0xfd, 0x57,
	0x7d, 0x04, 0x27, 0xd0, 0x5e, 0x24, 0xd8, 0x92, 0x7d, 0x70, 0x89, 0xd8, 0x77, 0xba, 0x8d, 0x35,
	0x35, 0x0d, 0x2e, 0x78, 0x0c, 0x6c, 0x8c, 0x51, 0x9c, 0x8c, 0x31, 0x93, 0x62, 0x5a, 0xd4, 0xdd,
	0x83, 0x66, 0x9c, 0xa4, 0x39, 0x9a, 0x9a, 0x77, 0xf8, 0xdc, 0x09, 0xde, 0xc0, 0xde, 0x18, 0x23,
	0x95, 0x63, 0x01, 0xb6, 0x35, 0xef, 0x81, 0xa7, 0x72, 0x5c, 0xc0, 0xad, 0x47, 0x71, 0x8d, 0x91,
	0xcc, 0x32, 0xbb, 0x31, 0xeb, 0x05, 0x3f, 0x1c, 0x68, 0x9c, 0xab, 0x88, 0x9d, 0xc2, 0xd6, 0x54,
	0xa2, 0x88, 0x04, 0x0a, 0xfb, 0x46, 0x87, 0x35, 0xfd, 0x72, 0xa9, 0x55, 0x9e, 0x85, 0xf2, 0xcc,
	0xc2, 0x79, 0x99, 0xc8, 0x86, 0xe0, 0xea, 0x54, 0x86, 0xa6, 0xc4, 0xf6, 0xf0, 0xa0, 0x7e, 0xe0,
	0x71, 0x2a, 0x43, 0x6e, 0xb0, 0xec, 0x25, 0x35, 0x26, 0x30, 0xd7, 0x7e, 0xc3, 0x64, 0x75, 0xff,
	0x90, 0x65, 0x70, 0xdc, 0xe2, 0x83, 0x29, 0x6c, 0x5a, 0x2a, 0xf6, 0x1a, 0x20, 0x54, 0x09, 0x8a,
	0x38, 0x91, 0x59, 0xf1, 0xde, 0x75, 0x44, 0xa7, 0x05, 0x90, 0x57, 0x72, 0x58, 0x17, 0xb6, 0x2f,
	0x95, 0xc6, 0xf7, 0x12, 0xbf, 0xab, 0xec, 0xca, 0x3e, 0x52, 0x35, 0x14, 0xfc, 0x74, 0xa0, 0x55,
	0xe6, 0x96, 0x77, 0xe7, 0x2c, 0xee, 0xce, 0x6c, 0x6a, 0x2a, 0x26, 0xc5, 0x31, 0xce, 0x9d, 0xe2,
	0x43, 0x69, 0x94, 0x1f, 0x0a, 0x3b, 0x00, 0x20, 0xc6, 0x38, 0x99, 0x8c, 0xe2, 0xcc, 0x77, 0x0d,
	0xb8, 0x12, 0x21, 0x6e, 0x91, 0x4d, 0xb4, 0xdf, 0xec, 0x36, 0x88, 0x9b, 0x6c, 0x62, 0x91, 0xc9,
	0xcc, 0xf7, 0x4c, 0x88, 0x4c, 0xf6, 0x02, 0xbc, 0xa9, 0xca, 0x13, 0xd4, 0xfe, 0xa6, 0x99, 0x77,
	0xbf, 0x66, 0xde, 0x33, 0x02, 0x71, 0x8b, 0x0d, 0x14, 0x34, 0x4d, 0x80, 0x8a, 0xe0, 0x75, 0x5a,
	0x0e, 0x40, 0xb6, 0x39, 0x12, 0xb3, 0x5b, 0x3b, 0x81, 0xf5, 0xe8, 0x71, 0x22, 0xa9, 0x31, 0x4e,
	0x04, 0xc6, 0x2a, 0x31, 0xa3, 0xb4, 0x78, 0x35, 0xc4, 0x7c, 0xd8, 0x54, 0x29, 0x59, 0xda, 0x77,
	0x4d, 0x8b, 0x85, 0x1b, 0x08, 0x68, 0x95, 0xab, 0x63, 0x1f, 0x60, 0xb7, 0x7c, 0xf3, 0x79, 0x48,
	0x16, 0xeb, 0x7a, 0xb8, 0x6e, 0x5d, 0x76, 0xfb, 0xab, 0x04, 0xc1, 0x67, 0xd8, 0x59, 0x42, 0x51,
	0xc7, 0x25, 0xee, 0xed, 0xc8, 0x0e, 0x59, 0x0d, 0xd5, 0x2c, 0x6b, 0x0f, 0x9a, 0x74, 0x5d, 0xd2,
	0xce, 0x38, 0x77, 0x48, 0x96, 0x96, 0xaf, 0xfe, 0xd6, 0x03, 0xb8, 0x21, 0x11, 0x1b, 0x4b, 0x12,
	0x31, 0xfc, 0xb5, 0x01, 0x2e, 0xe9, 0x03, 0xbb, 0x00, 0x6f, 0x2e, 0x94, 0xac, 0xee, 0x1b, 0x5b,
	0xd6, 0xe6, 0x4e, 0x6f, 0x3d, 0xd0, 0x0a, 0xc0, 0x05, 0x78, 0x73, 0xf1, 0xab, 0x25, 0x5f, 0x56,
	0xd8, 0x5a, 0xf2, 0x55, 0x11, 0xfd, 0x08, 0x2e, 0xa9, 0x1c, 0xab, 0x5b, 0xd6, 0x92, 0x66, 0x76,
	0x0e, 0xd7, 0xe2, 0x2c, 0x71, 0x04, 0xde, 0x31, 0xa2, 0x08, 0x2f, 0xd9, 0xa3, 0x9a, 0x94, 0x55,
	0x65, 0xec, 0x3c, 0xa9, 0x87, 0xae, 0x08, 0x63, 0xcf, 0x79, 0xe6, 0x9c, 0x3c, 0xf8, 0x74, 0x3f,
	0xbd, 0x9a, 0x0c, 0x44, 0x1a, 0x0f, 0x8a, 0xa4, 0x01, 0x25, 0x0d, 0x66, 0x47, 0xaf, 0xe8, 0xef,
	0x17, 0xcf, 0xfc, 0x68, 0x3e, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x9c, 0x95, 0x8f, 0x42,
	0x07, 0x00, 0x00,
}
