// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: lessons_tasks.proto

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LessonsService_CreateLesson_FullMethodName       = "/classroom.v1.LessonsService/CreateLesson"
	LessonsService_GetLessonsByCourse_FullMethodName = "/classroom.v1.LessonsService/GetLessonsByCourse"
	LessonsService_DeleteLesson_FullMethodName       = "/classroom.v1.LessonsService/DeleteLesson"
)

// LessonsServiceClient is the client API for LessonsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для работы с уроками
type LessonsServiceClient interface {
	CreateLesson(ctx context.Context, in *LessonCreateRequest, opts ...grpc.CallOption) (*LessonResponse, error)
	GetLessonsByCourse(ctx context.Context, in *CourseIdRequest, opts ...grpc.CallOption) (*LessonsListResponse, error)
	DeleteLesson(ctx context.Context, in *LessonIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type lessonsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLessonsServiceClient(cc grpc.ClientConnInterface) LessonsServiceClient {
	return &lessonsServiceClient{cc}
}

func (c *lessonsServiceClient) CreateLesson(ctx context.Context, in *LessonCreateRequest, opts ...grpc.CallOption) (*LessonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LessonResponse)
	err := c.cc.Invoke(ctx, LessonsService_CreateLesson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonsServiceClient) GetLessonsByCourse(ctx context.Context, in *CourseIdRequest, opts ...grpc.CallOption) (*LessonsListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LessonsListResponse)
	err := c.cc.Invoke(ctx, LessonsService_GetLessonsByCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lessonsServiceClient) DeleteLesson(ctx context.Context, in *LessonIdRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, LessonsService_DeleteLesson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LessonsServiceServer is the server API for LessonsService service.
// All implementations must embed UnimplementedLessonsServiceServer
// for forward compatibility.
//
// Сервис для работы с уроками
type LessonsServiceServer interface {
	CreateLesson(context.Context, *LessonCreateRequest) (*LessonResponse, error)
	GetLessonsByCourse(context.Context, *CourseIdRequest) (*LessonsListResponse, error)
	DeleteLesson(context.Context, *LessonIdRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedLessonsServiceServer()
}

// UnimplementedLessonsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLessonsServiceServer struct{}

func (UnimplementedLessonsServiceServer) CreateLesson(context.Context, *LessonCreateRequest) (*LessonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLesson not implemented")
}
func (UnimplementedLessonsServiceServer) GetLessonsByCourse(context.Context, *CourseIdRequest) (*LessonsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLessonsByCourse not implemented")
}
func (UnimplementedLessonsServiceServer) DeleteLesson(context.Context, *LessonIdRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLesson not implemented")
}
func (UnimplementedLessonsServiceServer) mustEmbedUnimplementedLessonsServiceServer() {}
func (UnimplementedLessonsServiceServer) testEmbeddedByValue()                        {}

// UnsafeLessonsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LessonsServiceServer will
// result in compilation errors.
type UnsafeLessonsServiceServer interface {
	mustEmbedUnimplementedLessonsServiceServer()
}

func RegisterLessonsServiceServer(s grpc.ServiceRegistrar, srv LessonsServiceServer) {
	// If the following call pancis, it indicates UnimplementedLessonsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LessonsService_ServiceDesc, srv)
}

func _LessonsService_CreateLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LessonCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonsServiceServer).CreateLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LessonsService_CreateLesson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonsServiceServer).CreateLesson(ctx, req.(*LessonCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonsService_GetLessonsByCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonsServiceServer).GetLessonsByCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LessonsService_GetLessonsByCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonsServiceServer).GetLessonsByCourse(ctx, req.(*CourseIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LessonsService_DeleteLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LessonIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LessonsServiceServer).DeleteLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LessonsService_DeleteLesson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LessonsServiceServer).DeleteLesson(ctx, req.(*LessonIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LessonsService_ServiceDesc is the grpc.ServiceDesc for LessonsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LessonsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "classroom.v1.LessonsService",
	HandlerType: (*LessonsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLesson",
			Handler:    _LessonsService_CreateLesson_Handler,
		},
		{
			MethodName: "GetLessonsByCourse",
			Handler:    _LessonsService_GetLessonsByCourse_Handler,
		},
		{
			MethodName: "DeleteLesson",
			Handler:    _LessonsService_DeleteLesson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lessons_tasks.proto",
}

const (
	TasksService_CreateTask_FullMethodName              = "/classroom.v1.TasksService/CreateTask"
	TasksService_GetTasksByCourse_FullMethodName        = "/classroom.v1.TasksService/GetTasksByCourse"
	TasksService_CompleteTask_FullMethodName            = "/classroom.v1.TasksService/CompleteTask"
	TasksService_GetCompletedTasksByUser_FullMethodName = "/classroom.v1.TasksService/GetCompletedTasksByUser"
	TasksService_SubmitTaskAnswer_FullMethodName        = "/classroom.v1.TasksService/SubmitTaskAnswer"
	TasksService_GetTaskAnswers_FullMethodName          = "/classroom.v1.TasksService/GetTaskAnswers"
)

// TasksServiceClient is the client API for TasksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для работы с заданиями
type TasksServiceClient interface {
	CreateTask(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	GetTasksByCourse(ctx context.Context, in *CourseIdRequest, opts ...grpc.CallOption) (*TasksListResponse, error)
	CompleteTask(ctx context.Context, in *UserTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error)
	GetCompletedTasksByUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserTasksListResponse, error)
	// Отправка ответа на задание
	SubmitTaskAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error)
	// Получение ответов ученика на задание
	GetTaskAnswers(ctx context.Context, in *TaskAnswersRequest, opts ...grpc.CallOption) (*TaskAnswersResponse, error)
}

type tasksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksServiceClient(cc grpc.ClientConnInterface) TasksServiceClient {
	return &tasksServiceClient{cc}
}

func (c *tasksServiceClient) CreateTask(ctx context.Context, in *TaskCreateRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, TasksService_CreateTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) GetTasksByCourse(ctx context.Context, in *CourseIdRequest, opts ...grpc.CallOption) (*TasksListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TasksListResponse)
	err := c.cc.Invoke(ctx, TasksService_GetTasksByCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) CompleteTask(ctx context.Context, in *UserTaskRequest, opts ...grpc.CallOption) (*CompleteTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompleteTaskResponse)
	err := c.cc.Invoke(ctx, TasksService_CompleteTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) GetCompletedTasksByUser(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*UserTasksListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserTasksListResponse)
	err := c.cc.Invoke(ctx, TasksService_GetCompletedTasksByUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) SubmitTaskAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitAnswerResponse)
	err := c.cc.Invoke(ctx, TasksService_SubmitTaskAnswer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) GetTaskAnswers(ctx context.Context, in *TaskAnswersRequest, opts ...grpc.CallOption) (*TaskAnswersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TaskAnswersResponse)
	err := c.cc.Invoke(ctx, TasksService_GetTaskAnswers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServiceServer is the server API for TasksService service.
// All implementations must embed UnimplementedTasksServiceServer
// for forward compatibility.
//
// Сервис для работы с заданиями
type TasksServiceServer interface {
	CreateTask(context.Context, *TaskCreateRequest) (*TaskResponse, error)
	GetTasksByCourse(context.Context, *CourseIdRequest) (*TasksListResponse, error)
	CompleteTask(context.Context, *UserTaskRequest) (*CompleteTaskResponse, error)
	GetCompletedTasksByUser(context.Context, *UserIdRequest) (*UserTasksListResponse, error)
	// Отправка ответа на задание
	SubmitTaskAnswer(context.Context, *SubmitAnswerRequest) (*SubmitAnswerResponse, error)
	// Получение ответов ученика на задание
	GetTaskAnswers(context.Context, *TaskAnswersRequest) (*TaskAnswersResponse, error)
	mustEmbedUnimplementedTasksServiceServer()
}

// UnimplementedTasksServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTasksServiceServer struct{}

func (UnimplementedTasksServiceServer) CreateTask(context.Context, *TaskCreateRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTasksServiceServer) GetTasksByCourse(context.Context, *CourseIdRequest) (*TasksListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasksByCourse not implemented")
}
func (UnimplementedTasksServiceServer) CompleteTask(context.Context, *UserTaskRequest) (*CompleteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteTask not implemented")
}
func (UnimplementedTasksServiceServer) GetCompletedTasksByUser(context.Context, *UserIdRequest) (*UserTasksListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompletedTasksByUser not implemented")
}
func (UnimplementedTasksServiceServer) SubmitTaskAnswer(context.Context, *SubmitAnswerRequest) (*SubmitAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTaskAnswer not implemented")
}
func (UnimplementedTasksServiceServer) GetTaskAnswers(context.Context, *TaskAnswersRequest) (*TaskAnswersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskAnswers not implemented")
}
func (UnimplementedTasksServiceServer) mustEmbedUnimplementedTasksServiceServer() {}
func (UnimplementedTasksServiceServer) testEmbeddedByValue()                      {}

// UnsafeTasksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksServiceServer will
// result in compilation errors.
type UnsafeTasksServiceServer interface {
	mustEmbedUnimplementedTasksServiceServer()
}

func RegisterTasksServiceServer(s grpc.ServiceRegistrar, srv TasksServiceServer) {
	// If the following call pancis, it indicates UnimplementedTasksServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TasksService_ServiceDesc, srv)
}

func _TasksService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TasksService_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).CreateTask(ctx, req.(*TaskCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_GetTasksByCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).GetTasksByCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TasksService_GetTasksByCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).GetTasksByCourse(ctx, req.(*CourseIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_CompleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).CompleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TasksService_CompleteTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).CompleteTask(ctx, req.(*UserTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_GetCompletedTasksByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).GetCompletedTasksByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TasksService_GetCompletedTasksByUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).GetCompletedTasksByUser(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_SubmitTaskAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).SubmitTaskAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TasksService_SubmitTaskAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).SubmitTaskAnswer(ctx, req.(*SubmitAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_GetTaskAnswers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskAnswersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).GetTaskAnswers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TasksService_GetTaskAnswers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).GetTaskAnswers(ctx, req.(*TaskAnswersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TasksService_ServiceDesc is the grpc.ServiceDesc for TasksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TasksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "classroom.v1.TasksService",
	HandlerType: (*TasksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TasksService_CreateTask_Handler,
		},
		{
			MethodName: "GetTasksByCourse",
			Handler:    _TasksService_GetTasksByCourse_Handler,
		},
		{
			MethodName: "CompleteTask",
			Handler:    _TasksService_CompleteTask_Handler,
		},
		{
			MethodName: "GetCompletedTasksByUser",
			Handler:    _TasksService_GetCompletedTasksByUser_Handler,
		},
		{
			MethodName: "SubmitTaskAnswer",
			Handler:    _TasksService_SubmitTaskAnswer_Handler,
		},
		{
			MethodName: "GetTaskAnswers",
			Handler:    _TasksService_GetTaskAnswers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lessons_tasks.proto",
}
