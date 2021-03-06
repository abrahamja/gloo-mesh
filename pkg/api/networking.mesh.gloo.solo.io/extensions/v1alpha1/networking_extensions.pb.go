// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/networking/extensions/v1alpha1/networking_extensions.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/discovery.mesh.gloo.solo.io/v1alpha2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1alpha3 "istio.io/api/networking/v1alpha3"
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

// the parameters provided to the Extensions server when requesting patches
type ExtensionPatchRequest struct {
	// the set of discovery objects provided as inputs for the Gloo Mesh translation
	Inputs *DiscoverySnapshot `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// the base set of output objects translated by Gloo Mesh.
	// these may have been operated upon by a previous Extension server if multiple servers
	// have been configured.
	Outputs              []*GeneratedObject `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExtensionPatchRequest) Reset()         { *m = ExtensionPatchRequest{} }
func (m *ExtensionPatchRequest) String() string { return proto.CompactTextString(m) }
func (*ExtensionPatchRequest) ProtoMessage()    {}
func (*ExtensionPatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{0}
}
func (m *ExtensionPatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionPatchRequest.Unmarshal(m, b)
}
func (m *ExtensionPatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionPatchRequest.Marshal(b, m, deterministic)
}
func (m *ExtensionPatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionPatchRequest.Merge(m, src)
}
func (m *ExtensionPatchRequest) XXX_Size() int {
	return xxx_messageInfo_ExtensionPatchRequest.Size(m)
}
func (m *ExtensionPatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionPatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionPatchRequest proto.InternalMessageInfo

func (m *ExtensionPatchRequest) GetInputs() *DiscoverySnapshot {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *ExtensionPatchRequest) GetOutputs() []*GeneratedObject {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// the set of patches the server wishes to apply to the Gloo Mesh Networking outputs.
// Any objects provided here will be inserted into the final Gloo Mesh snapshot.
// If an object already exists in the snapshot, it will be overridden by the version provided here.
// If multiple extensions servers are configured, this response may be
// operated upon by Extension patches provided by subsequent servers.
type ExtensionPatchResponse struct {
	// the set of modified/added output objects desired by the Extension server.
	PatchedOutputs       []*GeneratedObject `protobuf:"bytes,1,rep,name=patched_outputs,json=patchedOutputs,proto3" json:"patched_outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ExtensionPatchResponse) Reset()         { *m = ExtensionPatchResponse{} }
func (m *ExtensionPatchResponse) String() string { return proto.CompactTextString(m) }
func (*ExtensionPatchResponse) ProtoMessage()    {}
func (*ExtensionPatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{1}
}
func (m *ExtensionPatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionPatchResponse.Unmarshal(m, b)
}
func (m *ExtensionPatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionPatchResponse.Marshal(b, m, deterministic)
}
func (m *ExtensionPatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionPatchResponse.Merge(m, src)
}
func (m *ExtensionPatchResponse) XXX_Size() int {
	return xxx_messageInfo_ExtensionPatchResponse.Size(m)
}
func (m *ExtensionPatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionPatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionPatchResponse proto.InternalMessageInfo

func (m *ExtensionPatchResponse) GetPatchedOutputs() []*GeneratedObject {
	if m != nil {
		return m.PatchedOutputs
	}
	return nil
}

// a Protobuf representation of the set of Discovery objects used to produce the Networking outputs.
type DiscoverySnapshot struct {
	// all meshes in the discovery snapshot
	Meshes []*MeshObject `protobuf:"bytes,1,rep,name=meshes,proto3" json:"meshes,omitempty"`
	// all traffic targets in the discovery snapshot
	TrafficTargets []*TrafficTargetObject `protobuf:"bytes,2,rep,name=traffic_targets,json=trafficTargets,proto3" json:"traffic_targets,omitempty"`
	// all workloads in the discovery snapshot
	Workloads            []*WorkloadObject `protobuf:"bytes,3,rep,name=workloads,proto3" json:"workloads,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DiscoverySnapshot) Reset()         { *m = DiscoverySnapshot{} }
func (m *DiscoverySnapshot) String() string { return proto.CompactTextString(m) }
func (*DiscoverySnapshot) ProtoMessage()    {}
func (*DiscoverySnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{2}
}
func (m *DiscoverySnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverySnapshot.Unmarshal(m, b)
}
func (m *DiscoverySnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverySnapshot.Marshal(b, m, deterministic)
}
func (m *DiscoverySnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverySnapshot.Merge(m, src)
}
func (m *DiscoverySnapshot) XXX_Size() int {
	return xxx_messageInfo_DiscoverySnapshot.Size(m)
}
func (m *DiscoverySnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverySnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverySnapshot proto.InternalMessageInfo

func (m *DiscoverySnapshot) GetMeshes() []*MeshObject {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func (m *DiscoverySnapshot) GetTrafficTargets() []*TrafficTargetObject {
	if m != nil {
		return m.TrafficTargets
	}
	return nil
}

func (m *DiscoverySnapshot) GetWorkloads() []*WorkloadObject {
	if m != nil {
		return m.Workloads
	}
	return nil
}

// a proto-serializable representation of a TrafficTarget object
type TrafficTargetObject struct {
	// metadata of the object
	Metadata *ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// the spec of the object
	Spec *v1alpha2.TrafficTargetSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// the status of the object
	Status               *v1alpha2.TrafficTargetStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TrafficTargetObject) Reset()         { *m = TrafficTargetObject{} }
func (m *TrafficTargetObject) String() string { return proto.CompactTextString(m) }
func (*TrafficTargetObject) ProtoMessage()    {}
func (*TrafficTargetObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{3}
}
func (m *TrafficTargetObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrafficTargetObject.Unmarshal(m, b)
}
func (m *TrafficTargetObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrafficTargetObject.Marshal(b, m, deterministic)
}
func (m *TrafficTargetObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrafficTargetObject.Merge(m, src)
}
func (m *TrafficTargetObject) XXX_Size() int {
	return xxx_messageInfo_TrafficTargetObject.Size(m)
}
func (m *TrafficTargetObject) XXX_DiscardUnknown() {
	xxx_messageInfo_TrafficTargetObject.DiscardUnknown(m)
}

var xxx_messageInfo_TrafficTargetObject proto.InternalMessageInfo

func (m *TrafficTargetObject) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TrafficTargetObject) GetSpec() *v1alpha2.TrafficTargetSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TrafficTargetObject) GetStatus() *v1alpha2.TrafficTargetStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// a proto-serializable representation of a Workload object
type WorkloadObject struct {
	// metadata of the object
	Metadata *ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// the spec of the object
	Spec *v1alpha2.WorkloadSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// the status of the object
	Status               *v1alpha2.WorkloadStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WorkloadObject) Reset()         { *m = WorkloadObject{} }
func (m *WorkloadObject) String() string { return proto.CompactTextString(m) }
func (*WorkloadObject) ProtoMessage()    {}
func (*WorkloadObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{4}
}
func (m *WorkloadObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkloadObject.Unmarshal(m, b)
}
func (m *WorkloadObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkloadObject.Marshal(b, m, deterministic)
}
func (m *WorkloadObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadObject.Merge(m, src)
}
func (m *WorkloadObject) XXX_Size() int {
	return xxx_messageInfo_WorkloadObject.Size(m)
}
func (m *WorkloadObject) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadObject.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadObject proto.InternalMessageInfo

func (m *WorkloadObject) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *WorkloadObject) GetSpec() *v1alpha2.WorkloadSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *WorkloadObject) GetStatus() *v1alpha2.WorkloadStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// a proto-serializable representation of a Mesh object
type MeshObject struct {
	// metadata of the object
	Metadata *ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// the spec of the object
	Spec *v1alpha2.MeshSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// the status of the object
	Status               *v1alpha2.MeshStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MeshObject) Reset()         { *m = MeshObject{} }
func (m *MeshObject) String() string { return proto.CompactTextString(m) }
func (*MeshObject) ProtoMessage()    {}
func (*MeshObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{5}
}
func (m *MeshObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshObject.Unmarshal(m, b)
}
func (m *MeshObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshObject.Marshal(b, m, deterministic)
}
func (m *MeshObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshObject.Merge(m, src)
}
func (m *MeshObject) XXX_Size() int {
	return xxx_messageInfo_MeshObject.Size(m)
}
func (m *MeshObject) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshObject.DiscardUnknown(m)
}

var xxx_messageInfo_MeshObject proto.InternalMessageInfo

func (m *MeshObject) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MeshObject) GetSpec() *v1alpha2.MeshSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *MeshObject) GetStatus() *v1alpha2.MeshStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// a generated object can be of any output type supported by Gloo Mesh.
// the content of the type field should be used to determine
// the type of the output object.
// TODO(ilackarms): consider parameterizing Gloo Mesh to allow excluding GeneratedObjects from patch requests in the case where an implementer only performs additions (no updates required).
type GeneratedObject struct {
	// metadata of the object
	Metadata *ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// type of the object.
	//
	// Types that are valid to be assigned to Type:
	//	*GeneratedObject_DestinationRule
	//	*GeneratedObject_EnvoyFilter
	//	*GeneratedObject_ServiceEntry
	//	*GeneratedObject_VirtualService
	Type                 isGeneratedObject_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GeneratedObject) Reset()         { *m = GeneratedObject{} }
func (m *GeneratedObject) String() string { return proto.CompactTextString(m) }
func (*GeneratedObject) ProtoMessage()    {}
func (*GeneratedObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{6}
}
func (m *GeneratedObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneratedObject.Unmarshal(m, b)
}
func (m *GeneratedObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneratedObject.Marshal(b, m, deterministic)
}
func (m *GeneratedObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneratedObject.Merge(m, src)
}
func (m *GeneratedObject) XXX_Size() int {
	return xxx_messageInfo_GeneratedObject.Size(m)
}
func (m *GeneratedObject) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneratedObject.DiscardUnknown(m)
}

var xxx_messageInfo_GeneratedObject proto.InternalMessageInfo

type isGeneratedObject_Type interface {
	isGeneratedObject_Type()
}

type GeneratedObject_DestinationRule struct {
	DestinationRule *v1alpha3.DestinationRule `protobuf:"bytes,2,opt,name=destination_rule,json=destinationRule,proto3,oneof" json:"destination_rule,omitempty"`
}
type GeneratedObject_EnvoyFilter struct {
	EnvoyFilter *v1alpha3.EnvoyFilter `protobuf:"bytes,3,opt,name=envoy_filter,json=envoyFilter,proto3,oneof" json:"envoy_filter,omitempty"`
}
type GeneratedObject_ServiceEntry struct {
	ServiceEntry *v1alpha3.ServiceEntry `protobuf:"bytes,4,opt,name=service_entry,json=serviceEntry,proto3,oneof" json:"service_entry,omitempty"`
}
type GeneratedObject_VirtualService struct {
	VirtualService *v1alpha3.VirtualService `protobuf:"bytes,5,opt,name=virtual_service,json=virtualService,proto3,oneof" json:"virtual_service,omitempty"`
}

func (*GeneratedObject_DestinationRule) isGeneratedObject_Type() {}
func (*GeneratedObject_EnvoyFilter) isGeneratedObject_Type()     {}
func (*GeneratedObject_ServiceEntry) isGeneratedObject_Type()    {}
func (*GeneratedObject_VirtualService) isGeneratedObject_Type()  {}

func (m *GeneratedObject) GetType() isGeneratedObject_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *GeneratedObject) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *GeneratedObject) GetDestinationRule() *v1alpha3.DestinationRule {
	if x, ok := m.GetType().(*GeneratedObject_DestinationRule); ok {
		return x.DestinationRule
	}
	return nil
}

func (m *GeneratedObject) GetEnvoyFilter() *v1alpha3.EnvoyFilter {
	if x, ok := m.GetType().(*GeneratedObject_EnvoyFilter); ok {
		return x.EnvoyFilter
	}
	return nil
}

func (m *GeneratedObject) GetServiceEntry() *v1alpha3.ServiceEntry {
	if x, ok := m.GetType().(*GeneratedObject_ServiceEntry); ok {
		return x.ServiceEntry
	}
	return nil
}

func (m *GeneratedObject) GetVirtualService() *v1alpha3.VirtualService {
	if x, ok := m.GetType().(*GeneratedObject_VirtualService); ok {
		return x.VirtualService
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GeneratedObject) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GeneratedObject_DestinationRule)(nil),
		(*GeneratedObject_EnvoyFilter)(nil),
		(*GeneratedObject_ServiceEntry)(nil),
		(*GeneratedObject_VirtualService)(nil),
	}
}

// ObjectMeta is a simplified clone of the kubernetes ObjectMeta used to represent object metadata
// for K8s objects passed as messages in the NetworkingExtensions API.
type ObjectMeta struct {
	// the kubernetes name of the object
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// the kubernetes namespace of the object
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// the kubernetes clusterName of the object (used internally by Gloo Mesh)
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// the kubernetes labels on the object
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// the kubernetes annotations on the object
	Annotations          map[string]string `protobuf:"bytes,5,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ObjectMeta) Reset()         { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()    {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{7}
}
func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMeta.Unmarshal(m, b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMeta.Marshal(b, m, deterministic)
}
func (m *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(m, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return xxx_messageInfo_ObjectMeta.Size(m)
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *ObjectMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectMeta) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ObjectMeta) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ObjectMeta) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

// request to initiate push notifications
type WatchPushNotificationsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchPushNotificationsRequest) Reset()         { *m = WatchPushNotificationsRequest{} }
func (m *WatchPushNotificationsRequest) String() string { return proto.CompactTextString(m) }
func (*WatchPushNotificationsRequest) ProtoMessage()    {}
func (*WatchPushNotificationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{8}
}
func (m *WatchPushNotificationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchPushNotificationsRequest.Unmarshal(m, b)
}
func (m *WatchPushNotificationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchPushNotificationsRequest.Marshal(b, m, deterministic)
}
func (m *WatchPushNotificationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchPushNotificationsRequest.Merge(m, src)
}
func (m *WatchPushNotificationsRequest) XXX_Size() int {
	return xxx_messageInfo_WatchPushNotificationsRequest.Size(m)
}
func (m *WatchPushNotificationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchPushNotificationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchPushNotificationsRequest proto.InternalMessageInfo

// triggers a resync of Gloo Mesh objects
type PushNotification struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotification) Reset()         { *m = PushNotification{} }
func (m *PushNotification) String() string { return proto.CompactTextString(m) }
func (*PushNotification) ProtoMessage()    {}
func (*PushNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_686b689659d3c356, []int{9}
}
func (m *PushNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotification.Unmarshal(m, b)
}
func (m *PushNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotification.Marshal(b, m, deterministic)
}
func (m *PushNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotification.Merge(m, src)
}
func (m *PushNotification) XXX_Size() int {
	return xxx_messageInfo_PushNotification.Size(m)
}
func (m *PushNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotification.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotification proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtensionPatchRequest)(nil), "extensions.networking.mesh.gloo.solo.io.ExtensionPatchRequest")
	proto.RegisterType((*ExtensionPatchResponse)(nil), "extensions.networking.mesh.gloo.solo.io.ExtensionPatchResponse")
	proto.RegisterType((*DiscoverySnapshot)(nil), "extensions.networking.mesh.gloo.solo.io.DiscoverySnapshot")
	proto.RegisterType((*TrafficTargetObject)(nil), "extensions.networking.mesh.gloo.solo.io.TrafficTargetObject")
	proto.RegisterType((*WorkloadObject)(nil), "extensions.networking.mesh.gloo.solo.io.WorkloadObject")
	proto.RegisterType((*MeshObject)(nil), "extensions.networking.mesh.gloo.solo.io.MeshObject")
	proto.RegisterType((*GeneratedObject)(nil), "extensions.networking.mesh.gloo.solo.io.GeneratedObject")
	proto.RegisterType((*ObjectMeta)(nil), "extensions.networking.mesh.gloo.solo.io.ObjectMeta")
	proto.RegisterMapType((map[string]string)(nil), "extensions.networking.mesh.gloo.solo.io.ObjectMeta.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "extensions.networking.mesh.gloo.solo.io.ObjectMeta.LabelsEntry")
	proto.RegisterType((*WatchPushNotificationsRequest)(nil), "extensions.networking.mesh.gloo.solo.io.WatchPushNotificationsRequest")
	proto.RegisterType((*PushNotification)(nil), "extensions.networking.mesh.gloo.solo.io.PushNotification")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-mesh/api/networking/extensions/v1alpha1/networking_extensions.proto", fileDescriptor_686b689659d3c356)
}

var fileDescriptor_686b689659d3c356 = []byte{
	// 948 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x97, 0xdf, 0x72, 0xdb, 0x44,
	0x1b, 0xc6, 0x2d, 0xdb, 0xf1, 0xf7, 0xe5, 0x75, 0x88, 0xc3, 0xb6, 0x74, 0x3c, 0x1e, 0x18, 0x5a,
	0xcf, 0xd0, 0xa4, 0xed, 0x54, 0x6e, 0x93, 0x61, 0x68, 0x3a, 0xd0, 0x0c, 0x21, 0x69, 0x3d, 0x53,
	0x9a, 0x74, 0x94, 0x14, 0xcf, 0xc0, 0x81, 0x67, 0x23, 0xbf, 0xb6, 0x97, 0x28, 0x5a, 0xa1, 0x5d,
	0xb9, 0x78, 0xb8, 0x11, 0x4e, 0x80, 0xbb, 0xe0, 0x94, 0x8b, 0xe0, 0x88, 0x0b, 0xe0, 0x0e, 0xb8,
	0x00, 0x46, 0xab, 0x55, 0xf4, 0xa7, 0xaa, 0xa3, 0x86, 0x1c, 0x25, 0x5a, 0x3d, 0xef, 0xef, 0xd9,
	0xe7, 0xd5, 0x5a, 0xbb, 0x82, 0xef, 0x26, 0x4c, 0x4e, 0x83, 0x13, 0xd3, 0xe6, 0x67, 0x3d, 0xc1,
	0x1d, 0x7e, 0x9f, 0xf1, 0xde, 0xc4, 0xe1, 0xfc, 0xfe, 0x19, 0x8a, 0x69, 0x8f, 0x7a, 0xac, 0xe7,
	0xa2, 0x7c, 0xcd, 0xfd, 0x53, 0xe6, 0x4e, 0x7a, 0xf8, 0xa3, 0x44, 0x57, 0x30, 0xee, 0x8a, 0xde,
	0xec, 0x21, 0x75, 0xbc, 0x29, 0x7d, 0x98, 0xba, 0x3d, 0x4c, 0x6e, 0x9b, 0x9e, 0xcf, 0x25, 0x27,
	0xeb, 0xa9, 0x91, 0x44, 0x67, 0x86, 0x68, 0x33, 0x34, 0x31, 0x43, 0x47, 0x93, 0xf1, 0xce, 0xbd,
	0xac, 0xe5, 0x88, 0x09, 0x9b, 0xcf, 0xd0, 0x9f, 0xc7, 0x36, 0x9b, 0x3d, 0x55, 0xa3, 0xa8, 0x9d,
	0x4f, 0x2f, 0x14, 0x4b, 0x9f, 0x8e, 0xc7, 0xcc, 0x1e, 0x4a, 0xea, 0x4f, 0x50, 0xea, 0xb2, 0xde,
	0x85, 0x65, 0xe1, 0xfc, 0x1c, 0x4e, 0x47, 0xba, 0xe0, 0x6e, 0x2a, 0xb9, 0xd6, 0x6c, 0xf5, 0x46,
	0x28, 0x24, 0x73, 0xa9, 0x64, 0xdc, 0x1d, 0xfa, 0x81, 0x83, 0x5a, 0x7b, 0xbb, 0x48, 0x8b, 0xee,
	0x8c, 0xcf, 0x87, 0x63, 0xe6, 0x48, 0xf4, 0xb5, 0xee, 0x56, 0x91, 0x6e, 0x42, 0x25, 0xbe, 0xa6,
	0x73, 0x2d, 0x59, 0x2f, 0x92, 0x08, 0xf4, 0x67, 0xcc, 0xc6, 0x21, 0xba, 0xd2, 0x9f, 0x2f, 0x62,
	0x09, 0x36, 0x42, 0x9b, 0xc6, 0x76, 0x77, 0x8a, 0x24, 0x33, 0xe6, 0xcb, 0x80, 0x3a, 0x43, 0xcd,
	0xd4, 0xd2, 0x8d, 0x22, 0x69, 0xdc, 0x91, 0xb4, 0x6f, 0xf7, 0x0f, 0x03, 0x3e, 0xd8, 0x8f, 0x1f,
	0xec, 0x4b, 0x2a, 0xed, 0xa9, 0x85, 0x3f, 0x04, 0x28, 0x24, 0xb1, 0xa0, 0xc1, 0x5c, 0x2f, 0x90,
	0xa2, 0x6d, 0xdc, 0x34, 0x36, 0x9a, 0x9b, 0x8f, 0xcd, 0x92, 0x0b, 0xc0, 0xdc, 0x8b, 0x9f, 0xc6,
	0x91, 0x4b, 0x3d, 0x31, 0xe5, 0xd2, 0xd2, 0x24, 0x62, 0xc1, 0xff, 0x78, 0x20, 0x15, 0xb4, 0x7a,
	0xb3, 0xb6, 0xd1, 0xdc, 0x7c, 0x54, 0x1a, 0xfa, 0x0c, 0x5d, 0xf4, 0xa9, 0xc4, 0xd1, 0xe1, 0xc9,
	0xf7, 0x68, 0x4b, 0x2b, 0x06, 0x75, 0x7f, 0x82, 0x1b, 0xf9, 0x00, 0xc2, 0xe3, 0xae, 0x40, 0x42,
	0xa1, 0xe5, 0x85, 0x03, 0x38, 0x1a, 0xc6, 0xae, 0xc6, 0x7f, 0x74, 0x5d, 0xd5, 0xc0, 0x43, 0x6d,
	0xfe, 0x6b, 0x15, 0xde, 0x7f, 0x23, 0x2e, 0x79, 0x0e, 0x8d, 0x10, 0x85, 0xb1, 0xdf, 0x56, 0x69,
	0xbf, 0x17, 0x28, 0xa6, 0xda, 0x4a, 0x23, 0x08, 0x42, 0x2b, 0xfb, 0x13, 0x88, 0x7b, 0xf7, 0x79,
	0x69, 0xea, 0x71, 0x54, 0x7f, 0xac, 0xca, 0xe3, 0x24, 0x32, 0x3d, 0x28, 0xc8, 0x2b, 0x58, 0x8e,
	0x17, 0x88, 0x68, 0xd7, 0x94, 0xc1, 0x67, 0xa5, 0x0d, 0x06, 0xba, 0x52, 0xb3, 0x13, 0x52, 0xf7,
	0x1f, 0x03, 0xae, 0x15, 0xd8, 0x93, 0x43, 0xf8, 0xff, 0x19, 0x4a, 0x3a, 0xa2, 0x92, 0xea, 0xf5,
	0x55, 0xbe, 0x49, 0x11, 0xe2, 0x05, 0x4a, 0x6a, 0x9d, 0x43, 0xc8, 0x2e, 0xd4, 0x85, 0x87, 0x76,
	0xbb, 0xaa, 0x60, 0xa6, 0x79, 0xfe, 0x4a, 0xb8, 0xa8, 0x1f, 0x47, 0x1e, 0xda, 0x96, 0xaa, 0x25,
	0x7d, 0x68, 0x08, 0x49, 0x65, 0x10, 0x36, 0x20, 0xa4, 0x3c, 0x78, 0x07, 0x8a, 0xaa, 0xb3, 0x74,
	0x7d, 0xf7, 0x6f, 0x03, 0x56, 0xb3, 0x4d, 0xb9, 0xfa, 0xc4, 0x5f, 0x64, 0x12, 0xdf, 0x59, 0x38,
	0xd7, 0x78, 0x2e, 0xa9, 0xb0, 0x5f, 0xe5, 0xc2, 0xde, 0x2b, 0x07, 0xc8, 0xe6, 0xfc, 0xcb, 0x00,
	0x48, 0xd6, 0xec, 0xd5, 0x67, 0xdc, 0xce, 0x64, 0xfc, 0x64, 0xe1, 0x14, 0xc3, 0x79, 0xa4, 0xf2,
	0xed, 0xe4, 0xf2, 0xad, 0x5f, 0x5c, 0x9c, 0xcd, 0xf6, 0x7b, 0x0d, 0x5a, 0xb9, 0xdf, 0xff, 0xd5,
	0x07, 0x1c, 0xc0, 0x5a, 0x7e, 0x17, 0xd2, 0x61, 0xef, 0x9a, 0x4c, 0x48, 0xc6, 0xd3, 0xcc, 0xf8,
	0x55, 0x6e, 0xee, 0x25, 0x25, 0x56, 0xe0, 0x60, 0xbf, 0x62, 0xb5, 0x46, 0xd9, 0x21, 0xf2, 0x1c,
	0x56, 0xd2, 0x5b, 0x96, 0x6e, 0xc2, 0xed, 0x05, 0xd0, 0xfd, 0x50, 0xfe, 0x54, 0xa9, 0xfb, 0x15,
	0xab, 0x89, 0xc9, 0x25, 0x39, 0x80, 0xf7, 0x32, 0x9b, 0x56, 0xbb, 0xae, 0x5b, 0xfa, 0x76, 0xda,
	0x51, 0xa4, 0xdf, 0x0f, 0xe5, 0xfd, 0x8a, 0xb5, 0x22, 0x52, 0xd7, 0xe4, 0x18, 0x5a, 0xb9, 0x8d,
	0xab, 0xbd, 0xa4, 0x57, 0xf1, 0xdb, 0x89, 0xdf, 0x44, 0x15, 0x1a, 0xdc, 0xaf, 0x58, 0xab, 0xb3,
	0xcc, 0xc8, 0x6e, 0x03, 0xea, 0x72, 0xee, 0x61, 0xf7, 0x97, 0x1a, 0x40, 0xd2, 0x6c, 0x42, 0xa0,
	0xee, 0xd2, 0x33, 0x54, 0xcf, 0x6b, 0xd9, 0x52, 0xff, 0x93, 0x0f, 0x61, 0x39, 0xfc, 0x2b, 0x3c,
	0x6a, 0x47, 0xfd, 0x5e, 0xb6, 0x92, 0x01, 0x72, 0x0b, 0x56, 0x6c, 0x27, 0x10, 0x12, 0xfd, 0xa1,
	0xaa, 0xac, 0x29, 0x41, 0x53, 0x8f, 0x1d, 0x84, 0x80, 0x01, 0x34, 0x1c, 0x7a, 0x82, 0x8e, 0x68,
	0xd7, 0xd5, 0xbb, 0x72, 0xe7, 0x12, 0xcb, 0xc0, 0xfc, 0x5a, 0x11, 0x54, 0x4b, 0x2c, 0x8d, 0x23,
	0x63, 0x68, 0x52, 0xd7, 0xe5, 0x52, 0x3d, 0x49, 0xd1, 0x5e, 0x52, 0xf4, 0xbd, 0xcb, 0xd0, 0xbf,
	0x4c, 0x30, 0x91, 0x45, 0x1a, 0xdc, 0xd9, 0x86, 0x66, 0xca, 0x9e, 0xac, 0x41, 0xed, 0x14, 0xe7,
	0xba, 0x47, 0xe1, 0xbf, 0xe4, 0x3a, 0x2c, 0xcd, 0xa8, 0x13, 0xc4, 0xed, 0x89, 0x2e, 0x1e, 0x57,
	0x1f, 0x19, 0x9d, 0x27, 0xb0, 0x96, 0x67, 0xbf, 0x4b, 0x7d, 0xf7, 0x63, 0xf8, 0x68, 0x10, 0x6e,
	0xa3, 0x2f, 0x03, 0x31, 0x3d, 0xe0, 0x92, 0x8d, 0x99, 0x1d, 0xa1, 0xf4, 0xd1, 0xa3, 0x4b, 0x60,
	0x2d, 0x7f, 0x6f, 0xf3, 0xcf, 0x2a, 0x5c, 0x3f, 0x38, 0x4f, 0x7e, 0xbe, 0xe3, 0x0b, 0xf2, 0xb3,
	0x01, 0xd7, 0x9e, 0xa1, 0xcc, 0x9e, 0x01, 0x50, 0x90, 0x27, 0xa5, 0x7b, 0x56, 0x78, 0xfe, 0xe9,
	0xec, 0x5c, 0xba, 0x3e, 0x3a, 0x7e, 0x74, 0x2b, 0xe4, 0x37, 0x03, 0x6e, 0x14, 0x27, 0x25, 0x4f,
	0xcb, 0xef, 0xad, 0x8b, 0x5a, 0xd5, 0xd9, 0x2e, 0xcd, 0xc9, 0x23, 0xba, 0x95, 0x07, 0xc6, 0xee,
	0xe0, 0xdb, 0x57, 0x0b, 0xbf, 0x19, 0xbc, 0xd3, 0x49, 0xee, 0xbb, 0xe1, 0x4d, 0x72, 0xd1, 0x97,
	0xc4, 0x49, 0x43, 0x1d, 0x2f, 0xb7, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x76, 0xa6, 0xd8, 0xc8,
	0x93, 0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkingExtensionsClient is the client API for NetworkingExtensions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkingExtensionsClient interface {
	// GetExtensionPatches fetches a set of patches to the output configuration from the Extensions server.
	// The current discovery snapshot and translated outputs are provided in the ExtensionPatchRequest
	GetExtensionPatches(ctx context.Context, in *ExtensionPatchRequest, opts ...grpc.CallOption) (*ExtensionPatchResponse, error)
	// WatchPushNotifications initiates a streaming connection which allows the NetworkingExtensions server
	// to push notifications to Gloo Mesh telling it to resync its configuration.
	// This allows a NetworkingExtensions server to trigger Gloo Mesh to resync its state for
	// events triggered by objects not watched by Gloo Mesh.
	WatchPushNotifications(ctx context.Context, in *WatchPushNotificationsRequest, opts ...grpc.CallOption) (NetworkingExtensions_WatchPushNotificationsClient, error)
}

type networkingExtensionsClient struct {
	cc *grpc.ClientConn
}

func NewNetworkingExtensionsClient(cc *grpc.ClientConn) NetworkingExtensionsClient {
	return &networkingExtensionsClient{cc}
}

func (c *networkingExtensionsClient) GetExtensionPatches(ctx context.Context, in *ExtensionPatchRequest, opts ...grpc.CallOption) (*ExtensionPatchResponse, error) {
	out := new(ExtensionPatchResponse)
	err := c.cc.Invoke(ctx, "/extensions.networking.mesh.gloo.solo.io.NetworkingExtensions/GetExtensionPatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkingExtensionsClient) WatchPushNotifications(ctx context.Context, in *WatchPushNotificationsRequest, opts ...grpc.CallOption) (NetworkingExtensions_WatchPushNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkingExtensions_serviceDesc.Streams[0], "/extensions.networking.mesh.gloo.solo.io.NetworkingExtensions/WatchPushNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkingExtensionsWatchPushNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetworkingExtensions_WatchPushNotificationsClient interface {
	Recv() (*PushNotification, error)
	grpc.ClientStream
}

type networkingExtensionsWatchPushNotificationsClient struct {
	grpc.ClientStream
}

func (x *networkingExtensionsWatchPushNotificationsClient) Recv() (*PushNotification, error) {
	m := new(PushNotification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetworkingExtensionsServer is the server API for NetworkingExtensions service.
type NetworkingExtensionsServer interface {
	// GetExtensionPatches fetches a set of patches to the output configuration from the Extensions server.
	// The current discovery snapshot and translated outputs are provided in the ExtensionPatchRequest
	GetExtensionPatches(context.Context, *ExtensionPatchRequest) (*ExtensionPatchResponse, error)
	// WatchPushNotifications initiates a streaming connection which allows the NetworkingExtensions server
	// to push notifications to Gloo Mesh telling it to resync its configuration.
	// This allows a NetworkingExtensions server to trigger Gloo Mesh to resync its state for
	// events triggered by objects not watched by Gloo Mesh.
	WatchPushNotifications(*WatchPushNotificationsRequest, NetworkingExtensions_WatchPushNotificationsServer) error
}

// UnimplementedNetworkingExtensionsServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkingExtensionsServer struct {
}

func (*UnimplementedNetworkingExtensionsServer) GetExtensionPatches(ctx context.Context, req *ExtensionPatchRequest) (*ExtensionPatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExtensionPatches not implemented")
}
func (*UnimplementedNetworkingExtensionsServer) WatchPushNotifications(req *WatchPushNotificationsRequest, srv NetworkingExtensions_WatchPushNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchPushNotifications not implemented")
}

func RegisterNetworkingExtensionsServer(s *grpc.Server, srv NetworkingExtensionsServer) {
	s.RegisterService(&_NetworkingExtensions_serviceDesc, srv)
}

func _NetworkingExtensions_GetExtensionPatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtensionPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkingExtensionsServer).GetExtensionPatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/extensions.networking.mesh.gloo.solo.io.NetworkingExtensions/GetExtensionPatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkingExtensionsServer).GetExtensionPatches(ctx, req.(*ExtensionPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkingExtensions_WatchPushNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchPushNotificationsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetworkingExtensionsServer).WatchPushNotifications(m, &networkingExtensionsWatchPushNotificationsServer{stream})
}

type NetworkingExtensions_WatchPushNotificationsServer interface {
	Send(*PushNotification) error
	grpc.ServerStream
}

type networkingExtensionsWatchPushNotificationsServer struct {
	grpc.ServerStream
}

func (x *networkingExtensionsWatchPushNotificationsServer) Send(m *PushNotification) error {
	return x.ServerStream.SendMsg(m)
}

var _NetworkingExtensions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "extensions.networking.mesh.gloo.solo.io.NetworkingExtensions",
	HandlerType: (*NetworkingExtensionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExtensionPatches",
			Handler:    _NetworkingExtensions_GetExtensionPatches_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchPushNotifications",
			Handler:       _NetworkingExtensions_WatchPushNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/solo-io/gloo-mesh/api/networking/extensions/v1alpha1/networking_extensions.proto",
}
