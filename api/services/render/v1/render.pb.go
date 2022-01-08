// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: git.underland.io/ehazlett/finca/api/services/render/v1/render.proto

package render

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueueJobRequest struct {
	// Types that are valid to be assigned to Data:
	//	*QueueJobRequest_Request
	//	*QueueJobRequest_ChunkData
	Data                 isQueueJobRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *QueueJobRequest) Reset()         { *m = QueueJobRequest{} }
func (m *QueueJobRequest) String() string { return proto.CompactTextString(m) }
func (*QueueJobRequest) ProtoMessage()    {}
func (*QueueJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{0}
}
func (m *QueueJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueJobRequest.Unmarshal(m, b)
}
func (m *QueueJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueJobRequest.Marshal(b, m, deterministic)
}
func (m *QueueJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueJobRequest.Merge(m, src)
}
func (m *QueueJobRequest) XXX_Size() int {
	return xxx_messageInfo_QueueJobRequest.Size(m)
}
func (m *QueueJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueueJobRequest proto.InternalMessageInfo

type isQueueJobRequest_Data interface {
	isQueueJobRequest_Data()
}

type QueueJobRequest_Request struct {
	Request *JobRequest `protobuf:"bytes,1,opt,name=request,proto3,oneof" json:"request,omitempty"`
}
type QueueJobRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof" json:"chunk_data,omitempty"`
}

func (*QueueJobRequest_Request) isQueueJobRequest_Data()   {}
func (*QueueJobRequest_ChunkData) isQueueJobRequest_Data() {}

func (m *QueueJobRequest) GetData() isQueueJobRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *QueueJobRequest) GetRequest() *JobRequest {
	if x, ok := m.GetData().(*QueueJobRequest_Request); ok {
		return x.Request
	}
	return nil
}

func (m *QueueJobRequest) GetChunkData() []byte {
	if x, ok := m.GetData().(*QueueJobRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*QueueJobRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*QueueJobRequest_Request)(nil),
		(*QueueJobRequest_ChunkData)(nil),
	}
}

type JobRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ResolutionX          int64    `protobuf:"varint,2,opt,name=resolution_x,json=resolutionX,proto3" json:"resolution_x,omitempty"`
	ResolutionY          int64    `protobuf:"varint,3,opt,name=resolution_y,json=resolutionY,proto3" json:"resolution_y,omitempty"`
	ResolutionScale      int64    `protobuf:"varint,4,opt,name=resolution_scale,json=resolutionScale,proto3" json:"resolution_scale,omitempty"`
	RenderSamples        int64    `protobuf:"varint,5,opt,name=render_samples,json=renderSamples,proto3" json:"render_samples,omitempty"`
	RenderStartFrame     int64    `protobuf:"varint,6,opt,name=render_start_frame,json=renderStartFrame,proto3" json:"render_start_frame,omitempty"`
	RenderEndFrame       int64    `protobuf:"varint,7,opt,name=render_end_frame,json=renderEndFrame,proto3" json:"render_end_frame,omitempty"`
	RenderUseGPU         bool     `protobuf:"varint,8,opt,name=render_use_gpu,json=renderUseGpu,proto3" json:"render_use_gpu,omitempty"`
	RenderPriority       int64    `protobuf:"varint,9,opt,name=render_priority,json=renderPriority,proto3" json:"render_priority,omitempty"`
	CPU                  int64    `protobuf:"varint,10,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory               int64    `protobuf:"varint,11,opt,name=memory,proto3" json:"memory,omitempty"`
	RenderSlices         int64    `protobuf:"varint,12,opt,name=render_slices,json=renderSlices,proto3" json:"render_slices,omitempty"`
	ContentType          string   `protobuf:"bytes,13,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobRequest) Reset()         { *m = JobRequest{} }
func (m *JobRequest) String() string { return proto.CompactTextString(m) }
func (*JobRequest) ProtoMessage()    {}
func (*JobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{1}
}
func (m *JobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobRequest.Unmarshal(m, b)
}
func (m *JobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobRequest.Marshal(b, m, deterministic)
}
func (m *JobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobRequest.Merge(m, src)
}
func (m *JobRequest) XXX_Size() int {
	return xxx_messageInfo_JobRequest.Size(m)
}
func (m *JobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobRequest proto.InternalMessageInfo

func (m *JobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobRequest) GetResolutionX() int64 {
	if m != nil {
		return m.ResolutionX
	}
	return 0
}

func (m *JobRequest) GetResolutionY() int64 {
	if m != nil {
		return m.ResolutionY
	}
	return 0
}

func (m *JobRequest) GetResolutionScale() int64 {
	if m != nil {
		return m.ResolutionScale
	}
	return 0
}

func (m *JobRequest) GetRenderSamples() int64 {
	if m != nil {
		return m.RenderSamples
	}
	return 0
}

func (m *JobRequest) GetRenderStartFrame() int64 {
	if m != nil {
		return m.RenderStartFrame
	}
	return 0
}

func (m *JobRequest) GetRenderEndFrame() int64 {
	if m != nil {
		return m.RenderEndFrame
	}
	return 0
}

func (m *JobRequest) GetRenderUseGPU() bool {
	if m != nil {
		return m.RenderUseGPU
	}
	return false
}

func (m *JobRequest) GetRenderPriority() int64 {
	if m != nil {
		return m.RenderPriority
	}
	return 0
}

func (m *JobRequest) GetCPU() int64 {
	if m != nil {
		return m.CPU
	}
	return 0
}

func (m *JobRequest) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *JobRequest) GetRenderSlices() int64 {
	if m != nil {
		return m.RenderSlices
	}
	return 0
}

func (m *JobRequest) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

type QueueJobResponse struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueJobResponse) Reset()         { *m = QueueJobResponse{} }
func (m *QueueJobResponse) String() string { return proto.CompactTextString(m) }
func (*QueueJobResponse) ProtoMessage()    {}
func (*QueueJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{2}
}
func (m *QueueJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueJobResponse.Unmarshal(m, b)
}
func (m *QueueJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueJobResponse.Marshal(b, m, deterministic)
}
func (m *QueueJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueJobResponse.Merge(m, src)
}
func (m *QueueJobResponse) XXX_Size() int {
	return xxx_messageInfo_QueueJobResponse.Size(m)
}
func (m *QueueJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueueJobResponse proto.InternalMessageInfo

func (m *QueueJobResponse) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type JobStatus struct {
	Job                  *Job            `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Succeeded            bool            `protobuf:"varint,2,opt,name=succeeded,proto3" json:"succeeded,omitempty"`
	Duration             *types.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Worker               *Worker         `protobuf:"bytes,4,opt,name=worker,proto3" json:"worker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *JobStatus) Reset()         { *m = JobStatus{} }
func (m *JobStatus) String() string { return proto.CompactTextString(m) }
func (*JobStatus) ProtoMessage()    {}
func (*JobStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{3}
}
func (m *JobStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobStatus.Unmarshal(m, b)
}
func (m *JobStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobStatus.Marshal(b, m, deterministic)
}
func (m *JobStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobStatus.Merge(m, src)
}
func (m *JobStatus) XXX_Size() int {
	return xxx_messageInfo_JobStatus.Size(m)
}
func (m *JobStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_JobStatus.DiscardUnknown(m)
}

var xxx_messageInfo_JobStatus proto.InternalMessageInfo

func (m *JobStatus) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *JobStatus) GetSucceeded() bool {
	if m != nil {
		return m.Succeeded
	}
	return false
}

func (m *JobStatus) GetDuration() *types.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *JobStatus) GetWorker() *Worker {
	if m != nil {
		return m.Worker
	}
	return nil
}

type Worker struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	CPUs                 uint32   `protobuf:"varint,3,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Memory               int64    `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	GPUs                 []string `protobuf:"bytes,5,rep,name=gpus,proto3" json:"gpus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Worker) Reset()         { *m = Worker{} }
func (m *Worker) String() string { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()    {}
func (*Worker) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{4}
}
func (m *Worker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Worker.Unmarshal(m, b)
}
func (m *Worker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Worker.Marshal(b, m, deterministic)
}
func (m *Worker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Worker.Merge(m, src)
}
func (m *Worker) XXX_Size() int {
	return xxx_messageInfo_Worker.Size(m)
}
func (m *Worker) XXX_DiscardUnknown() {
	xxx_messageInfo_Worker.DiscardUnknown(m)
}

var xxx_messageInfo_Worker proto.InternalMessageInfo

func (m *Worker) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Worker) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Worker) GetCPUs() uint32 {
	if m != nil {
		return m.CPUs
	}
	return 0
}

func (m *Worker) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *Worker) GetGPUs() []string {
	if m != nil {
		return m.GPUs
	}
	return nil
}

type Job struct {
	ID                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Request              *JobRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	JobSource            string      `protobuf:"bytes,3,opt,name=job_source,json=jobSource,proto3" json:"job_source,omitempty"`
	RenderSliceIndex     int64       `protobuf:"varint,4,opt,name=render_slice_index,json=renderSliceIndex,proto3" json:"render_slice_index,omitempty"`
	RenderSliceMinX      float32     `protobuf:"fixed32,5,opt,name=render_slice_min_x,json=renderSliceMinX,proto3" json:"render_slice_min_x,omitempty"`
	RenderSliceMaxX      float32     `protobuf:"fixed32,6,opt,name=render_slice_max_x,json=renderSliceMaxX,proto3" json:"render_slice_max_x,omitempty"`
	RenderSliceMinY      float32     `protobuf:"fixed32,7,opt,name=render_slice_min_y,json=renderSliceMinY,proto3" json:"render_slice_min_y,omitempty"`
	RenderSliceMaxY      float32     `protobuf:"fixed32,8,opt,name=render_slice_max_y,json=renderSliceMaxY,proto3" json:"render_slice_max_y,omitempty"`
	RenderFrame          int64       `protobuf:"varint,9,opt,name=render_frame,json=renderFrame,proto3" json:"render_frame,omitempty"`
	OutputDir            string      `protobuf:"bytes,10,opt,name=output_dir,json=outputDir,proto3" json:"output_dir,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{5}
}
func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Job) GetRequest() *JobRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Job) GetJobSource() string {
	if m != nil {
		return m.JobSource
	}
	return ""
}

func (m *Job) GetRenderSliceIndex() int64 {
	if m != nil {
		return m.RenderSliceIndex
	}
	return 0
}

func (m *Job) GetRenderSliceMinX() float32 {
	if m != nil {
		return m.RenderSliceMinX
	}
	return 0
}

func (m *Job) GetRenderSliceMaxX() float32 {
	if m != nil {
		return m.RenderSliceMaxX
	}
	return 0
}

func (m *Job) GetRenderSliceMinY() float32 {
	if m != nil {
		return m.RenderSliceMinY
	}
	return 0
}

func (m *Job) GetRenderSliceMaxY() float32 {
	if m != nil {
		return m.RenderSliceMaxY
	}
	return 0
}

func (m *Job) GetRenderFrame() int64 {
	if m != nil {
		return m.RenderFrame
	}
	return 0
}

func (m *Job) GetOutputDir() string {
	if m != nil {
		return m.OutputDir
	}
	return ""
}

type ListJobsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsRequest) Reset()         { *m = ListJobsRequest{} }
func (m *ListJobsRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobsRequest) ProtoMessage()    {}
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{6}
}
func (m *ListJobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsRequest.Unmarshal(m, b)
}
func (m *ListJobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsRequest.Marshal(b, m, deterministic)
}
func (m *ListJobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsRequest.Merge(m, src)
}
func (m *ListJobsRequest) XXX_Size() int {
	return xxx_messageInfo_ListJobsRequest.Size(m)
}
func (m *ListJobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsRequest proto.InternalMessageInfo

type ListJobsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsResponse) Reset()         { *m = ListJobsResponse{} }
func (m *ListJobsResponse) String() string { return proto.CompactTextString(m) }
func (*ListJobsResponse) ProtoMessage()    {}
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{7}
}
func (m *ListJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsResponse.Unmarshal(m, b)
}
func (m *ListJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsResponse.Marshal(b, m, deterministic)
}
func (m *ListJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsResponse.Merge(m, src)
}
func (m *ListJobsResponse) XXX_Size() int {
	return xxx_messageInfo_ListJobsResponse.Size(m)
}
func (m *ListJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsResponse proto.InternalMessageInfo

type ListWorkersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListWorkersRequest) Reset()         { *m = ListWorkersRequest{} }
func (m *ListWorkersRequest) String() string { return proto.CompactTextString(m) }
func (*ListWorkersRequest) ProtoMessage()    {}
func (*ListWorkersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{8}
}
func (m *ListWorkersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkersRequest.Unmarshal(m, b)
}
func (m *ListWorkersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkersRequest.Marshal(b, m, deterministic)
}
func (m *ListWorkersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkersRequest.Merge(m, src)
}
func (m *ListWorkersRequest) XXX_Size() int {
	return xxx_messageInfo_ListWorkersRequest.Size(m)
}
func (m *ListWorkersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkersRequest proto.InternalMessageInfo

type ListWorkersResponse struct {
	Workers              []*Worker `protobuf:"bytes,1,rep,name=workers,proto3" json:"workers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListWorkersResponse) Reset()         { *m = ListWorkersResponse{} }
func (m *ListWorkersResponse) String() string { return proto.CompactTextString(m) }
func (*ListWorkersResponse) ProtoMessage()    {}
func (*ListWorkersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faa69ac579737354, []int{9}
}
func (m *ListWorkersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListWorkersResponse.Unmarshal(m, b)
}
func (m *ListWorkersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListWorkersResponse.Marshal(b, m, deterministic)
}
func (m *ListWorkersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListWorkersResponse.Merge(m, src)
}
func (m *ListWorkersResponse) XXX_Size() int {
	return xxx_messageInfo_ListWorkersResponse.Size(m)
}
func (m *ListWorkersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListWorkersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListWorkersResponse proto.InternalMessageInfo

func (m *ListWorkersResponse) GetWorkers() []*Worker {
	if m != nil {
		return m.Workers
	}
	return nil
}

func init() {
	proto.RegisterType((*QueueJobRequest)(nil), "finca.services.render.v1.QueueJobRequest")
	proto.RegisterType((*JobRequest)(nil), "finca.services.render.v1.JobRequest")
	proto.RegisterType((*QueueJobResponse)(nil), "finca.services.render.v1.QueueJobResponse")
	proto.RegisterType((*JobStatus)(nil), "finca.services.render.v1.JobStatus")
	proto.RegisterType((*Worker)(nil), "finca.services.render.v1.Worker")
	proto.RegisterType((*Job)(nil), "finca.services.render.v1.Job")
	proto.RegisterType((*ListJobsRequest)(nil), "finca.services.render.v1.ListJobsRequest")
	proto.RegisterType((*ListJobsResponse)(nil), "finca.services.render.v1.ListJobsResponse")
	proto.RegisterType((*ListWorkersRequest)(nil), "finca.services.render.v1.ListWorkersRequest")
	proto.RegisterType((*ListWorkersResponse)(nil), "finca.services.render.v1.ListWorkersResponse")
}

func init() {
	proto.RegisterFile("git.underland.io/ehazlett/finca/api/services/render/v1/render.proto", fileDescriptor_faa69ac579737354)
}

var fileDescriptor_faa69ac579737354 = []byte{
	// 915 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0x37, 0x71, 0x36, 0x89, 0x4f, 0x52, 0x12, 0x86, 0xd5, 0xca, 0x5b, 0xb5, 0x24, 0x18,
	0x10, 0x59, 0x58, 0x6c, 0xb6, 0x08, 0x84, 0x40, 0x7c, 0x28, 0x0d, 0xec, 0xb6, 0x02, 0xa9, 0xeb,
	0x10, 0xd1, 0x72, 0x63, 0xf9, 0x63, 0x9a, 0x75, 0x37, 0xf1, 0x98, 0xf9, 0x28, 0x09, 0xe2, 0x09,
	0x78, 0x0a, 0x1e, 0x84, 0x37, 0xe0, 0x1d, 0x7a, 0xd1, 0x6b, 0x1e, 0x02, 0xcd, 0x8c, 0x9d, 0xa4,
	0xe9, 0x87, 0xa2, 0xbd, 0x1b, 0xff, 0xe7, 0x77, 0xe6, 0x9c, 0x39, 0x73, 0xce, 0x49, 0x60, 0x7f,
	0x9c, 0x70, 0x47, 0xa4, 0x31, 0xa6, 0x93, 0x20, 0x8d, 0x9d, 0x84, 0xb8, 0xf8, 0x65, 0xf0, 0xc7,
	0x04, 0x73, 0xee, 0x9e, 0x26, 0x69, 0x14, 0xb8, 0x41, 0x96, 0xb8, 0x0c, 0xd3, 0xf3, 0x24, 0xc2,
	0xcc, 0xa5, 0x58, 0x92, 0xee, 0xf9, 0xd3, 0x7c, 0xe5, 0x64, 0x94, 0x70, 0x82, 0x2c, 0x85, 0x3a,
	0x05, 0xe6, 0xe4, 0x9b, 0xe7, 0x4f, 0xb7, 0x1f, 0x8c, 0xc9, 0x98, 0x28, 0xc8, 0x95, 0x2b, 0xcd,
	0x6f, 0xbf, 0x3d, 0x26, 0x64, 0x3c, 0xc1, 0xae, 0xfa, 0x0a, 0xc5, 0xa9, 0x1b, 0x0b, 0x1a, 0xf0,
	0x84, 0xa4, 0x7a, 0xdf, 0xfe, 0x13, 0x5a, 0x2f, 0x04, 0x16, 0xf8, 0x90, 0x84, 0x1e, 0xfe, 0x4d,
	0x60, 0xc6, 0xd1, 0x77, 0x50, 0xa3, 0x7a, 0x69, 0x95, 0xba, 0xa5, 0x5e, 0x63, 0xef, 0x3d, 0xe7,
	0x36, 0xa7, 0xce, 0xd2, 0xec, 0xf9, 0x3d, 0xaf, 0x30, 0x43, 0x1d, 0x80, 0xe8, 0xa5, 0x48, 0x5f,
	0xf9, 0x71, 0xc0, 0x03, 0xab, 0xdc, 0x2d, 0xf5, 0x9a, 0xcf, 0xef, 0x79, 0xa6, 0xd2, 0x06, 0x01,
	0x0f, 0xfa, 0x55, 0xa8, 0xc8, 0x2d, 0xfb, 0x3f, 0x03, 0x60, 0xc5, 0x33, 0x82, 0x4a, 0x1a, 0x4c,
	0xb1, 0x72, 0x6b, 0x7a, 0x6a, 0x8d, 0xde, 0x81, 0x26, 0xc5, 0x8c, 0x4c, 0x84, 0x0c, 0xda, 0x9f,
	0xa9, 0xd3, 0x0c, 0xaf, 0xb1, 0xd4, 0x8e, 0xd7, 0x90, 0xb9, 0x65, 0xac, 0x23, 0x27, 0xe8, 0x31,
	0xb4, 0x57, 0x10, 0x16, 0x05, 0x13, 0x6c, 0x55, 0x14, 0xd6, 0x5a, 0xea, 0x43, 0x29, 0xa3, 0xf7,
	0xe1, 0x0d, 0x7d, 0x3f, 0x9f, 0x05, 0xd3, 0x6c, 0x82, 0x99, 0x75, 0x5f, 0x81, 0x5b, 0x5a, 0x1d,
	0x6a, 0x11, 0x3d, 0x01, 0x54, 0x60, 0x3c, 0xa0, 0xdc, 0x3f, 0xa5, 0x32, 0xf2, 0xaa, 0x42, 0xdb,
	0x39, 0x2a, 0x37, 0x7e, 0x90, 0x3a, 0xea, 0x41, 0xae, 0xf9, 0x38, 0x8d, 0x73, 0xb6, 0xa6, 0xd8,
	0xdc, 0xd9, 0xf7, 0x69, 0xac, 0xc9, 0xcf, 0x17, 0xee, 0x05, 0xc3, 0xfe, 0x38, 0x13, 0x56, 0xbd,
	0x5b, 0xea, 0xd5, 0xfb, 0xed, 0xcb, 0x8b, 0x4e, 0xd3, 0x53, 0x3b, 0x23, 0x86, 0x9f, 0x1d, 0x8d,
	0xbc, 0x26, 0x5d, 0x7c, 0x65, 0x02, 0x7d, 0x00, 0xad, 0xdc, 0x2e, 0xa3, 0x09, 0xa1, 0x09, 0x9f,
	0x5b, 0xe6, 0xaa, 0x83, 0xa3, 0x5c, 0x45, 0x8f, 0xc0, 0x88, 0x32, 0x61, 0x81, 0xdc, 0xec, 0xd7,
	0x2e, 0x2f, 0x3a, 0xc6, 0xfe, 0xd1, 0xc8, 0x93, 0x1a, 0x7a, 0x08, 0xd5, 0x29, 0x9e, 0x12, 0x3a,
	0xb7, 0x1a, 0xca, 0x34, 0xff, 0x42, 0xef, 0xc2, 0x56, 0x71, 0xd7, 0x89, 0x2c, 0x00, 0xab, 0xa9,
	0xb6, 0xf3, 0x00, 0x86, 0x4a, 0x93, 0xaf, 0x10, 0x91, 0x94, 0xe3, 0x94, 0xfb, 0x7c, 0x9e, 0x61,
	0x6b, 0x4b, 0x3d, 0x62, 0x23, 0xd7, 0x7e, 0x9e, 0x67, 0xd8, 0xfe, 0x04, 0xda, 0xcb, 0x62, 0x63,
	0x19, 0x49, 0x19, 0x46, 0x3b, 0x50, 0x11, 0x22, 0x89, 0xf5, 0x9b, 0xf7, 0xeb, 0x97, 0x17, 0x9d,
	0xca, 0x68, 0x74, 0x30, 0xf0, 0x94, 0x6a, 0xff, 0x5b, 0x02, 0xf3, 0x90, 0x84, 0x43, 0x1e, 0x70,
	0xc1, 0x90, 0x0b, 0xc6, 0x19, 0x09, 0xf3, 0xaa, 0xdc, 0xbd, 0xbb, 0x2a, 0x25, 0x89, 0x76, 0xc0,
	0x64, 0x22, 0x8a, 0x30, 0x8e, 0x71, 0xac, 0x2a, 0xa7, 0xee, 0x2d, 0x05, 0xf4, 0x19, 0xd4, 0x8b,
	0x6e, 0x50, 0x35, 0xd3, 0xd8, 0x7b, 0xe4, 0xe8, 0x76, 0x71, 0x8a, 0x76, 0x71, 0x06, 0x39, 0xe0,
	0x2d, 0x50, 0xf4, 0x05, 0x54, 0x7f, 0x27, 0xf4, 0x15, 0xa6, 0xaa, 0x82, 0x1a, 0x7b, 0xdd, 0xdb,
	0x03, 0xf9, 0x45, 0x71, 0x5e, 0xce, 0xdb, 0x7f, 0x95, 0xa0, 0xaa, 0xa5, 0x1b, 0x4b, 0xdd, 0x82,
	0xda, 0x39, 0xa6, 0x4c, 0x86, 0x53, 0x56, 0x72, 0xf1, 0x29, 0x93, 0x14, 0x65, 0x82, 0xa9, 0x28,
	0xb7, 0x74, 0x92, 0xf6, 0x8f, 0x46, 0xcc, 0x53, 0xea, 0xca, 0xb3, 0x55, 0xae, 0x3c, 0xdb, 0x0e,
	0x54, 0xc6, 0xd2, 0xea, 0x7e, 0xd7, 0x28, 0x52, 0xfb, 0x4c, 0x59, 0x49, 0xd5, 0xfe, 0xdb, 0x00,
	0xe3, 0x90, 0x84, 0xe8, 0x21, 0x94, 0x17, 0xe9, 0xaf, 0x5e, 0x5e, 0x74, 0xca, 0x07, 0x03, 0xaf,
	0x9c, 0xc4, 0xe8, 0x9b, 0xe5, 0x18, 0x28, 0x6f, 0x3e, 0x06, 0x96, 0x43, 0x60, 0x17, 0xe0, 0x8c,
	0x84, 0x3e, 0x23, 0x82, 0x46, 0x58, 0x45, 0x6e, 0x7a, 0xe6, 0x19, 0x09, 0x87, 0x4a, 0x58, 0xed,
	0x1f, 0x59, 0x3f, 0x7e, 0x92, 0xc6, 0x78, 0x96, 0x5f, 0xa0, 0xbd, 0x52, 0x58, 0x07, 0x52, 0x47,
	0x1f, 0xad, 0xd1, 0xd3, 0x44, 0xce, 0x02, 0xd9, 0x98, 0x65, 0xaf, 0xb5, 0x42, 0xff, 0x94, 0xa4,
	0xc7, 0xd7, 0xe1, 0x60, 0xe6, 0xcf, 0x54, 0x6b, 0xae, 0xc1, 0xc1, 0xec, 0xf8, 0xc6, 0x93, 0xe7,
	0xaa, 0x37, 0xaf, 0x9d, 0x7c, 0x72, 0xe3, 0xc9, 0x73, 0xd5, 0xa0, 0xd7, 0x4e, 0x3e, 0xd1, 0x63,
	0x49, 0xc1, 0xba, 0xdf, 0xcd, 0x62, 0x2c, 0x49, 0x4d, 0x37, 0xfb, 0x2e, 0x00, 0x11, 0x3c, 0x13,
	0xdc, 0x8f, 0x13, 0xaa, 0x5a, 0xd2, 0xf4, 0x4c, 0xad, 0x0c, 0x12, 0x6a, 0xbf, 0x09, 0xad, 0x1f,
	0x13, 0xc6, 0x0f, 0x49, 0xc8, 0xf2, 0xf4, 0xda, 0x08, 0xda, 0x4b, 0x49, 0xb7, 0x90, 0xfd, 0x00,
	0x90, 0xd4, 0x74, 0x65, 0x2d, 0xc8, 0x17, 0xf0, 0xd6, 0x15, 0x35, 0xef, 0xb7, 0x2f, 0xa1, 0xa6,
	0xab, 0x91, 0x59, 0xa5, 0xae, 0xb1, 0x51, 0xf9, 0x16, 0x06, 0x7b, 0xff, 0x94, 0xa1, 0xaa, 0x47,
	0x10, 0x8a, 0xa0, 0x5e, 0xb4, 0x32, 0x7a, 0x7c, 0xfb, 0x09, 0x6b, 0xbf, 0x2d, 0xdb, 0x1f, 0x6e,
	0x82, 0xea, 0x48, 0x7b, 0x25, 0x14, 0x40, 0xbd, 0xb8, 0xec, 0x5d, 0x4e, 0xd6, 0x72, 0x74, 0x97,
	0x93, 0xf5, 0xdc, 0xa1, 0x33, 0x68, 0xac, 0x64, 0x09, 0x3d, 0xb9, 0xdb, 0xf4, 0x6a, 0x8a, 0xb7,
	0x3f, 0xde, 0x90, 0xd6, 0xbe, 0xfa, 0xdf, 0xfe, 0xfa, 0xf5, 0xeb, 0xfd, 0x05, 0xf8, 0x4a, 0xaf,
	0xc2, 0xaa, 0x1a, 0x4b, 0x9f, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x47, 0x81, 0x6e, 0x4a,
	0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RenderClient is the client API for Render service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RenderClient interface {
	QueueJob(ctx context.Context, opts ...grpc.CallOption) (Render_QueueJobClient, error)
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error)
	ListWorkers(ctx context.Context, in *ListWorkersRequest, opts ...grpc.CallOption) (*ListWorkersResponse, error)
}

type renderClient struct {
	cc *grpc.ClientConn
}

func NewRenderClient(cc *grpc.ClientConn) RenderClient {
	return &renderClient{cc}
}

func (c *renderClient) QueueJob(ctx context.Context, opts ...grpc.CallOption) (Render_QueueJobClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Render_serviceDesc.Streams[0], "/finca.services.render.v1.Render/QueueJob", opts...)
	if err != nil {
		return nil, err
	}
	x := &renderQueueJobClient{stream}
	return x, nil
}

type Render_QueueJobClient interface {
	Send(*QueueJobRequest) error
	CloseAndRecv() (*QueueJobResponse, error)
	grpc.ClientStream
}

type renderQueueJobClient struct {
	grpc.ClientStream
}

func (x *renderQueueJobClient) Send(m *QueueJobRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *renderQueueJobClient) CloseAndRecv() (*QueueJobResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(QueueJobResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *renderClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error) {
	out := new(ListJobsResponse)
	err := c.cc.Invoke(ctx, "/finca.services.render.v1.Render/ListJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *renderClient) ListWorkers(ctx context.Context, in *ListWorkersRequest, opts ...grpc.CallOption) (*ListWorkersResponse, error) {
	out := new(ListWorkersResponse)
	err := c.cc.Invoke(ctx, "/finca.services.render.v1.Render/ListWorkers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RenderServer is the server API for Render service.
type RenderServer interface {
	QueueJob(Render_QueueJobServer) error
	ListJobs(context.Context, *ListJobsRequest) (*ListJobsResponse, error)
	ListWorkers(context.Context, *ListWorkersRequest) (*ListWorkersResponse, error)
}

// UnimplementedRenderServer can be embedded to have forward compatible implementations.
type UnimplementedRenderServer struct {
}

func (*UnimplementedRenderServer) QueueJob(srv Render_QueueJobServer) error {
	return status.Errorf(codes.Unimplemented, "method QueueJob not implemented")
}
func (*UnimplementedRenderServer) ListJobs(ctx context.Context, req *ListJobsRequest) (*ListJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJobs not implemented")
}
func (*UnimplementedRenderServer) ListWorkers(ctx context.Context, req *ListWorkersRequest) (*ListWorkersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkers not implemented")
}

func RegisterRenderServer(s *grpc.Server, srv RenderServer) {
	s.RegisterService(&_Render_serviceDesc, srv)
}

func _Render_QueueJob_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RenderServer).QueueJob(&renderQueueJobServer{stream})
}

type Render_QueueJobServer interface {
	SendAndClose(*QueueJobResponse) error
	Recv() (*QueueJobRequest, error)
	grpc.ServerStream
}

type renderQueueJobServer struct {
	grpc.ServerStream
}

func (x *renderQueueJobServer) SendAndClose(m *QueueJobResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *renderQueueJobServer) Recv() (*QueueJobRequest, error) {
	m := new(QueueJobRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Render_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RenderServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finca.services.render.v1.Render/ListJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RenderServer).ListJobs(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Render_ListWorkers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RenderServer).ListWorkers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/finca.services.render.v1.Render/ListWorkers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RenderServer).ListWorkers(ctx, req.(*ListWorkersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Render_serviceDesc = grpc.ServiceDesc{
	ServiceName: "finca.services.render.v1.Render",
	HandlerType: (*RenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListJobs",
			Handler:    _Render_ListJobs_Handler,
		},
		{
			MethodName: "ListWorkers",
			Handler:    _Render_ListWorkers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "QueueJob",
			Handler:       _Render_QueueJob_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "git.underland.io/ehazlett/finca/api/services/render/v1/render.proto",
}
