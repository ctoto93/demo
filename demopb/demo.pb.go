// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: demopb/demo.proto

package demopb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
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

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Credit   int32      `protobuf:"varint,3,opt,name=credit,proto3" json:"credit,omitempty"`
	Students []*Student `protobuf:"bytes,4,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demopb_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_demopb_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_demopb_demo_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetCredit() int32 {
	if x != nil {
		return x.Credit
	}
	return 0
}

func (x *Course) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age     int32     `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Courses []*Course `protobuf:"bytes,4,rep,name=courses,proto3" json:"courses,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demopb_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_demopb_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_demopb_demo_proto_rawDescGZIP(), []int{1}
}

func (x *Student) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Student) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

var File_demopb_demo_proto protoreflect.FileDescriptor

var file_demopb_demo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x2b,
	0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x69, 0x0a, 0x07, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x32, 0xce, 0x03, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x12, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0e, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x0e, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x1a, 0x0e, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x0e, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x1a, 0x0e, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demopb_demo_proto_rawDescOnce sync.Once
	file_demopb_demo_proto_rawDescData = file_demopb_demo_proto_rawDesc
)

func file_demopb_demo_proto_rawDescGZIP() []byte {
	file_demopb_demo_proto_rawDescOnce.Do(func() {
		file_demopb_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demopb_demo_proto_rawDescData)
	})
	return file_demopb_demo_proto_rawDescData
}

var file_demopb_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_demopb_demo_proto_goTypes = []interface{}{
	(*Course)(nil),                 // 0: demopb.Course
	(*Student)(nil),                // 1: demopb.Student
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 3: google.protobuf.Empty
}
var file_demopb_demo_proto_depIdxs = []int32{
	1,  // 0: demopb.Course.students:type_name -> demopb.Student
	0,  // 1: demopb.Student.courses:type_name -> demopb.Course
	2,  // 2: demopb.DemoService.GetStudent:input_type -> google.protobuf.StringValue
	1,  // 3: demopb.DemoService.AddStudent:input_type -> demopb.Student
	1,  // 4: demopb.DemoService.EditStudent:input_type -> demopb.Student
	2,  // 5: demopb.DemoService.DeleteStudent:input_type -> google.protobuf.StringValue
	2,  // 6: demopb.DemoService.GetCourse:input_type -> google.protobuf.StringValue
	0,  // 7: demopb.DemoService.AddCourse:input_type -> demopb.Course
	0,  // 8: demopb.DemoService.EditCourse:input_type -> demopb.Course
	2,  // 9: demopb.DemoService.DeleteCourse:input_type -> google.protobuf.StringValue
	1,  // 10: demopb.DemoService.GetStudent:output_type -> demopb.Student
	1,  // 11: demopb.DemoService.AddStudent:output_type -> demopb.Student
	1,  // 12: demopb.DemoService.EditStudent:output_type -> demopb.Student
	3,  // 13: demopb.DemoService.DeleteStudent:output_type -> google.protobuf.Empty
	0,  // 14: demopb.DemoService.GetCourse:output_type -> demopb.Course
	0,  // 15: demopb.DemoService.AddCourse:output_type -> demopb.Course
	0,  // 16: demopb.DemoService.EditCourse:output_type -> demopb.Course
	3,  // 17: demopb.DemoService.DeleteCourse:output_type -> google.protobuf.Empty
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_demopb_demo_proto_init() }
func file_demopb_demo_proto_init() {
	if File_demopb_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demopb_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_demopb_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
			RawDescriptor: file_demopb_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demopb_demo_proto_goTypes,
		DependencyIndexes: file_demopb_demo_proto_depIdxs,
		MessageInfos:      file_demopb_demo_proto_msgTypes,
	}.Build()
	File_demopb_demo_proto = out.File
	file_demopb_demo_proto_rawDesc = nil
	file_demopb_demo_proto_goTypes = nil
	file_demopb_demo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoServiceClient interface {
	GetStudent(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Student, error)
	AddStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error)
	EditStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error)
	DeleteStudent(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCourse(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Course, error)
	AddCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error)
	EditCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error)
	DeleteCourse(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type demoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDemoServiceClient(cc grpc.ClientConnInterface) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) GetStudent(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/demopb.DemoService/GetStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) AddStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/demopb.DemoService/AddStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) EditStudent(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/demopb.DemoService/EditStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) DeleteStudent(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/demopb.DemoService/DeleteStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) GetCourse(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/demopb.DemoService/GetCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) AddCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/demopb.DemoService/AddCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) EditCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/demopb.DemoService/EditCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *demoServiceClient) DeleteCourse(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/demopb.DemoService/DeleteCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
type DemoServiceServer interface {
	GetStudent(context.Context, *wrapperspb.StringValue) (*Student, error)
	AddStudent(context.Context, *Student) (*Student, error)
	EditStudent(context.Context, *Student) (*Student, error)
	DeleteStudent(context.Context, *wrapperspb.StringValue) (*emptypb.Empty, error)
	GetCourse(context.Context, *wrapperspb.StringValue) (*Course, error)
	AddCourse(context.Context, *Course) (*Course, error)
	EditCourse(context.Context, *Course) (*Course, error)
	DeleteCourse(context.Context, *wrapperspb.StringValue) (*emptypb.Empty, error)
}

// UnimplementedDemoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (*UnimplementedDemoServiceServer) GetStudent(context.Context, *wrapperspb.StringValue) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (*UnimplementedDemoServiceServer) AddStudent(context.Context, *Student) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudent not implemented")
}
func (*UnimplementedDemoServiceServer) EditStudent(context.Context, *Student) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditStudent not implemented")
}
func (*UnimplementedDemoServiceServer) DeleteStudent(context.Context, *wrapperspb.StringValue) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudent not implemented")
}
func (*UnimplementedDemoServiceServer) GetCourse(context.Context, *wrapperspb.StringValue) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (*UnimplementedDemoServiceServer) AddCourse(context.Context, *Course) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCourse not implemented")
}
func (*UnimplementedDemoServiceServer) EditCourse(context.Context, *Course) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCourse not implemented")
}
func (*UnimplementedDemoServiceServer) DeleteCourse(context.Context, *wrapperspb.StringValue) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demopb.DemoService/GetStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).GetStudent(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_AddStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).AddStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demopb.DemoService/AddStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).AddStudent(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_EditStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).EditStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demopb.DemoService/EditStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).EditStudent(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_DeleteStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).DeleteStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demopb.DemoService/DeleteStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).DeleteStudent(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demopb.DemoService/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).GetCourse(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_AddCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).AddCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demopb.DemoService/AddCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).AddCourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_EditCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).EditCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demopb.DemoService/EditCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).EditCourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _DemoService_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demopb.DemoService/DeleteCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).DeleteCourse(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demopb.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudent",
			Handler:    _DemoService_GetStudent_Handler,
		},
		{
			MethodName: "AddStudent",
			Handler:    _DemoService_AddStudent_Handler,
		},
		{
			MethodName: "EditStudent",
			Handler:    _DemoService_EditStudent_Handler,
		},
		{
			MethodName: "DeleteStudent",
			Handler:    _DemoService_DeleteStudent_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _DemoService_GetCourse_Handler,
		},
		{
			MethodName: "AddCourse",
			Handler:    _DemoService_AddCourse_Handler,
		},
		{
			MethodName: "EditCourse",
			Handler:    _DemoService_EditCourse_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _DemoService_DeleteCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demopb/demo.proto",
}
