// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/v1/mesh.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	core1 "github.com/solo-io/mesh-projects/pkg/api/v1/core"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

//
//Meshes represent a currently registered service mesh.
type Mesh struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by mesh-discovery during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// Types that are valid to be assigned to MeshType:
	//	*Mesh_Istio
	//	*Mesh_AwsAppMesh
	//	*Mesh_Linkerd
	MeshType isMesh_MeshType `protobuf_oneof:"mesh_type"`
	// mtls config specifies configuration options for enabling mutual
	// tls between pods in this mesh
	MtlsConfig *MtlsConfig `protobuf:"bytes,2,opt,name=mtls_config,json=mtlsConfig,proto3" json:"mtls_config,omitempty"`
	// configuration for propagating stats and metrics from
	// mesh controllers and sidecars to a centralized datastore
	// such as prometheus
	MonitoringConfig *MonitoringConfig `protobuf:"bytes,3,opt,name=monitoring_config,json=monitoringConfig,proto3" json:"monitoring_config,omitempty"`
	// object which represents the data mesh discovery finds about a given mesh
	DiscoveryMetadata *DiscoveryMetadata `protobuf:"bytes,6,opt,name=discovery_metadata,json=discoveryMetadata,proto3" json:"discovery_metadata,omitempty"`
	// whether or not to use SMI to configure this mesh
	SmiEnabled bool `protobuf:"varint,7,opt,name=smi_enabled,json=smiEnabled,proto3" json:"smi_enabled,omitempty"`
	// reference to the EntryPoints to this mesh
	EntryPoint           *core1.ClusterResourceRef `protobuf:"bytes,8,opt,name=entry_point,json=entryPoint,proto3" json:"entry_point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Mesh) Reset()         { *m = Mesh{} }
func (m *Mesh) String() string { return proto.CompactTextString(m) }
func (*Mesh) ProtoMessage()    {}
func (*Mesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_0569f8953a7ffa98, []int{0}
}
func (m *Mesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mesh.Unmarshal(m, b)
}
func (m *Mesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mesh.Marshal(b, m, deterministic)
}
func (m *Mesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mesh.Merge(m, src)
}
func (m *Mesh) XXX_Size() int {
	return xxx_messageInfo_Mesh.Size(m)
}
func (m *Mesh) XXX_DiscardUnknown() {
	xxx_messageInfo_Mesh.DiscardUnknown(m)
}

var xxx_messageInfo_Mesh proto.InternalMessageInfo

type isMesh_MeshType interface {
	isMesh_MeshType()
	Equal(interface{}) bool
}

type Mesh_Istio struct {
	Istio *IstioMesh `protobuf:"bytes,1,opt,name=istio,proto3,oneof" json:"istio,omitempty"`
}
type Mesh_AwsAppMesh struct {
	AwsAppMesh *AwsAppMesh `protobuf:"bytes,4,opt,name=aws_app_mesh,json=awsAppMesh,proto3,oneof" json:"aws_app_mesh,omitempty"`
}
type Mesh_Linkerd struct {
	Linkerd *LinkerdMesh `protobuf:"bytes,5,opt,name=linkerd,proto3,oneof" json:"linkerd,omitempty"`
}

func (*Mesh_Istio) isMesh_MeshType()      {}
func (*Mesh_AwsAppMesh) isMesh_MeshType() {}
func (*Mesh_Linkerd) isMesh_MeshType()    {}

func (m *Mesh) GetMeshType() isMesh_MeshType {
	if m != nil {
		return m.MeshType
	}
	return nil
}

func (m *Mesh) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Mesh) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Mesh) GetIstio() *IstioMesh {
	if x, ok := m.GetMeshType().(*Mesh_Istio); ok {
		return x.Istio
	}
	return nil
}

func (m *Mesh) GetAwsAppMesh() *AwsAppMesh {
	if x, ok := m.GetMeshType().(*Mesh_AwsAppMesh); ok {
		return x.AwsAppMesh
	}
	return nil
}

func (m *Mesh) GetLinkerd() *LinkerdMesh {
	if x, ok := m.GetMeshType().(*Mesh_Linkerd); ok {
		return x.Linkerd
	}
	return nil
}

func (m *Mesh) GetMtlsConfig() *MtlsConfig {
	if m != nil {
		return m.MtlsConfig
	}
	return nil
}

func (m *Mesh) GetMonitoringConfig() *MonitoringConfig {
	if m != nil {
		return m.MonitoringConfig
	}
	return nil
}

func (m *Mesh) GetDiscoveryMetadata() *DiscoveryMetadata {
	if m != nil {
		return m.DiscoveryMetadata
	}
	return nil
}

func (m *Mesh) GetSmiEnabled() bool {
	if m != nil {
		return m.SmiEnabled
	}
	return false
}

func (m *Mesh) GetEntryPoint() *core1.ClusterResourceRef {
	if m != nil {
		return m.EntryPoint
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Mesh) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Mesh_Istio)(nil),
		(*Mesh_AwsAppMesh)(nil),
		(*Mesh_Linkerd)(nil),
	}
}

// Generic discovery data shared between different meshes
type DiscoveryMetadata struct {
	// The name of the cluster the mesh is installed to
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Whether or not auto-injection is enabled for a given mesh
	EnableAutoInject bool `protobuf:"varint,2,opt,name=enable_auto_inject,json=enableAutoInject,proto3" json:"enable_auto_inject,omitempty"`
	// upstreams which point to injected pods in the mesh
	Upstreams []*core.ResourceRef `protobuf:"bytes,5,rep,name=upstreams,proto3" json:"upstreams,omitempty"`
	// discovered mtls config of the given mesh
	MtlsConfig           *MtlsConfig `protobuf:"bytes,6,opt,name=mtls_config,json=mtlsConfig,proto3" json:"mtls_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DiscoveryMetadata) Reset()         { *m = DiscoveryMetadata{} }
func (m *DiscoveryMetadata) String() string { return proto.CompactTextString(m) }
func (*DiscoveryMetadata) ProtoMessage()    {}
func (*DiscoveryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0569f8953a7ffa98, []int{1}
}
func (m *DiscoveryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryMetadata.Unmarshal(m, b)
}
func (m *DiscoveryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryMetadata.Marshal(b, m, deterministic)
}
func (m *DiscoveryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryMetadata.Merge(m, src)
}
func (m *DiscoveryMetadata) XXX_Size() int {
	return xxx_messageInfo_DiscoveryMetadata.Size(m)
}
func (m *DiscoveryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryMetadata proto.InternalMessageInfo

func (m *DiscoveryMetadata) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *DiscoveryMetadata) GetEnableAutoInject() bool {
	if m != nil {
		return m.EnableAutoInject
	}
	return false
}

func (m *DiscoveryMetadata) GetUpstreams() []*core.ResourceRef {
	if m != nil {
		return m.Upstreams
	}
	return nil
}

func (m *DiscoveryMetadata) GetMtlsConfig() *MtlsConfig {
	if m != nil {
		return m.MtlsConfig
	}
	return nil
}

// Mesh object representing an installed Istio control plane
type IstioMesh struct {
	// where the istio control plane has been installed
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// version of istio which has been installed
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IstioMesh) Reset()         { *m = IstioMesh{} }
func (m *IstioMesh) String() string { return proto.CompactTextString(m) }
func (*IstioMesh) ProtoMessage()    {}
func (*IstioMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_0569f8953a7ffa98, []int{2}
}
func (m *IstioMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IstioMesh.Unmarshal(m, b)
}
func (m *IstioMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IstioMesh.Marshal(b, m, deterministic)
}
func (m *IstioMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IstioMesh.Merge(m, src)
}
func (m *IstioMesh) XXX_Size() int {
	return xxx_messageInfo_IstioMesh.Size(m)
}
func (m *IstioMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_IstioMesh.DiscardUnknown(m)
}

var xxx_messageInfo_IstioMesh proto.InternalMessageInfo

func (m *IstioMesh) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *IstioMesh) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Mesh object representing AWS App Mesh
type AwsAppMesh struct {
	// Reference to the secret that holds the AWS credentials that will be used to access the AWS App Mesh service.
	AwsSecret *core.ResourceRef `protobuf:"bytes,1,opt,name=aws_secret,json=awsSecret,proto3" json:"aws_secret,omitempty"`
	// The AWS region the AWS App Mesh control plane resources (Virtual Nodes, Virtual Routers, etc.) will be created in.
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Determines whether pods will be automatically injected with the AWS App Mesh Envoy sidecar proxy.
	EnableAutoInject bool `protobuf:"varint,3,opt,name=enable_auto_inject,json=enableAutoInject,proto3" json:"enable_auto_inject,omitempty"`
	// Pods matching this selector will be injected with the sidecar proxy at creation time.
	//
	// NOTE: the sidecar injector webhook currently supports only the NamespaceSelector and LabelSelector
	InjectionSelector *PodSelector `protobuf:"bytes,4,opt,name=injection_selector,json=injectionSelector,proto3" json:"injection_selector,omitempty"`
	// If auto-injection is enabled, the value of the pod label with this key will be used to calculate the value of
	// APPMESH_VIRTUAL_NODE_NAME environment variable that is set on the injected sidecar proxy container.
	VirtualNodeLabel string `protobuf:"bytes,5,opt,name=virtual_node_label,json=virtualNodeLabel,proto3" json:"virtual_node_label,omitempty"`
	// Reference to the config map that contains the patch that will be applied to the spec of the pods matching the
	// injection_selector.
	SidecarPatchConfigMap *core.ResourceRef `protobuf:"bytes,6,opt,name=sidecar_patch_config_map,json=sidecarPatchConfigMap,proto3" json:"sidecar_patch_config_map,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *AwsAppMesh) Reset()         { *m = AwsAppMesh{} }
func (m *AwsAppMesh) String() string { return proto.CompactTextString(m) }
func (*AwsAppMesh) ProtoMessage()    {}
func (*AwsAppMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_0569f8953a7ffa98, []int{3}
}
func (m *AwsAppMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsAppMesh.Unmarshal(m, b)
}
func (m *AwsAppMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsAppMesh.Marshal(b, m, deterministic)
}
func (m *AwsAppMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsAppMesh.Merge(m, src)
}
func (m *AwsAppMesh) XXX_Size() int {
	return xxx_messageInfo_AwsAppMesh.Size(m)
}
func (m *AwsAppMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsAppMesh.DiscardUnknown(m)
}

var xxx_messageInfo_AwsAppMesh proto.InternalMessageInfo

func (m *AwsAppMesh) GetAwsSecret() *core.ResourceRef {
	if m != nil {
		return m.AwsSecret
	}
	return nil
}

func (m *AwsAppMesh) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *AwsAppMesh) GetEnableAutoInject() bool {
	if m != nil {
		return m.EnableAutoInject
	}
	return false
}

func (m *AwsAppMesh) GetInjectionSelector() *PodSelector {
	if m != nil {
		return m.InjectionSelector
	}
	return nil
}

func (m *AwsAppMesh) GetVirtualNodeLabel() string {
	if m != nil {
		return m.VirtualNodeLabel
	}
	return ""
}

func (m *AwsAppMesh) GetSidecarPatchConfigMap() *core.ResourceRef {
	if m != nil {
		return m.SidecarPatchConfigMap
	}
	return nil
}

// Mesh object representing an installed Linkerd control plane
type LinkerdMesh struct {
	// where the Linkerd control plane has been installed
	InstallationNamespace string `protobuf:"bytes,1,opt,name=installation_namespace,json=installationNamespace,proto3" json:"installation_namespace,omitempty"`
	// version of istio which has been installed
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkerdMesh) Reset()         { *m = LinkerdMesh{} }
func (m *LinkerdMesh) String() string { return proto.CompactTextString(m) }
func (*LinkerdMesh) ProtoMessage()    {}
func (*LinkerdMesh) Descriptor() ([]byte, []int) {
	return fileDescriptor_0569f8953a7ffa98, []int{4}
}
func (m *LinkerdMesh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkerdMesh.Unmarshal(m, b)
}
func (m *LinkerdMesh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkerdMesh.Marshal(b, m, deterministic)
}
func (m *LinkerdMesh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkerdMesh.Merge(m, src)
}
func (m *LinkerdMesh) XXX_Size() int {
	return xxx_messageInfo_LinkerdMesh.Size(m)
}
func (m *LinkerdMesh) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkerdMesh.DiscardUnknown(m)
}

var xxx_messageInfo_LinkerdMesh proto.InternalMessageInfo

func (m *LinkerdMesh) GetInstallationNamespace() string {
	if m != nil {
		return m.InstallationNamespace
	}
	return ""
}

func (m *LinkerdMesh) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// the encryption configuration that will be applied by the role
type MtlsConfig struct {
	// whether or not mutual TLS should be enabled between pods in this mesh
	MtlsEnabled bool `protobuf:"varint,1,opt,name=mtls_enabled,json=mtlsEnabled,proto3" json:"mtls_enabled,omitempty"`
	// Connection can be either plaintext or TLS, and client cert can be omitted
	// this is true by default for Linkerd
	MtlsPermissive bool `protobuf:"varint,2,opt,name=mtls_permissive,json=mtlsPermissive,proto3" json:"mtls_permissive,omitempty"`
	// if set, rootCertificate will override the root certificate used by the mesh
	// to encrypt mtls connections.
	//
	// The structure of the secret must be a standard kubernetes TLS secret
	// such as can be created via `kubectl create secret tls`
	//
	// if mtlsEnabled is false, this field is ignored
	// If deploying to Consul, Consul Connect requires that the cert and key are generated using ec, not rsa.
	RootCertificate      *core.ResourceRef `protobuf:"bytes,3,opt,name=root_certificate,json=rootCertificate,proto3" json:"root_certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MtlsConfig) Reset()         { *m = MtlsConfig{} }
func (m *MtlsConfig) String() string { return proto.CompactTextString(m) }
func (*MtlsConfig) ProtoMessage()    {}
func (*MtlsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0569f8953a7ffa98, []int{5}
}
func (m *MtlsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MtlsConfig.Unmarshal(m, b)
}
func (m *MtlsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MtlsConfig.Marshal(b, m, deterministic)
}
func (m *MtlsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MtlsConfig.Merge(m, src)
}
func (m *MtlsConfig) XXX_Size() int {
	return xxx_messageInfo_MtlsConfig.Size(m)
}
func (m *MtlsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MtlsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MtlsConfig proto.InternalMessageInfo

func (m *MtlsConfig) GetMtlsEnabled() bool {
	if m != nil {
		return m.MtlsEnabled
	}
	return false
}

func (m *MtlsConfig) GetMtlsPermissive() bool {
	if m != nil {
		return m.MtlsPermissive
	}
	return false
}

func (m *MtlsConfig) GetRootCertificate() *core.ResourceRef {
	if m != nil {
		return m.RootCertificate
	}
	return nil
}

// Contains configuration options for monitoring a mesh
// Currently MonitoringConfig only contains options for configuring
// an in-cluster Prometheus instance to scrape a mesh for metrics
type MonitoringConfig struct {
	// indicates to mesh-discovery that metrics should be propagated to one or more instances of prometheus.
	// add a [`core.solo.io.ResourceRef`](../../../../solo-kit/api/v1/ref.proto.sk#ResourceRef) for each
	// NAMESPACE.NAME of the configmap used to configure each prometheus instance.
	// assumes that the configmap contains a key named `prometheus.yml` or `prometheus.yaml` whose value
	// is the prometheus yaml config as an inline string
	PrometheusConfigmaps []core.ResourceRef `protobuf:"bytes,1,rep,name=prometheus_configmaps,json=prometheusConfigmaps,proto3" json:"prometheus_configmaps"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MonitoringConfig) Reset()         { *m = MonitoringConfig{} }
func (m *MonitoringConfig) String() string { return proto.CompactTextString(m) }
func (*MonitoringConfig) ProtoMessage()    {}
func (*MonitoringConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0569f8953a7ffa98, []int{6}
}
func (m *MonitoringConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitoringConfig.Unmarshal(m, b)
}
func (m *MonitoringConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitoringConfig.Marshal(b, m, deterministic)
}
func (m *MonitoringConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringConfig.Merge(m, src)
}
func (m *MonitoringConfig) XXX_Size() int {
	return xxx_messageInfo_MonitoringConfig.Size(m)
}
func (m *MonitoringConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringConfig proto.InternalMessageInfo

func (m *MonitoringConfig) GetPrometheusConfigmaps() []core.ResourceRef {
	if m != nil {
		return m.PrometheusConfigmaps
	}
	return nil
}

type MeshGroup struct {
	// Status indicates the validation status of this resource.
	// Status is read-only by clients, and set by mesh-discovery during validation
	Status core.Status `protobuf:"bytes,100,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,101,opt,name=metadata,proto3" json:"metadata"`
	// User-provided display name for the mesh group.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The meshes contained in this group.
	Meshes               []*core.ResourceRef `protobuf:"bytes,3,rep,name=meshes,proto3" json:"meshes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MeshGroup) Reset()         { *m = MeshGroup{} }
func (m *MeshGroup) String() string { return proto.CompactTextString(m) }
func (*MeshGroup) ProtoMessage()    {}
func (*MeshGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_0569f8953a7ffa98, []int{7}
}
func (m *MeshGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshGroup.Unmarshal(m, b)
}
func (m *MeshGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshGroup.Marshal(b, m, deterministic)
}
func (m *MeshGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshGroup.Merge(m, src)
}
func (m *MeshGroup) XXX_Size() int {
	return xxx_messageInfo_MeshGroup.Size(m)
}
func (m *MeshGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MeshGroup proto.InternalMessageInfo

func (m *MeshGroup) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *MeshGroup) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *MeshGroup) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *MeshGroup) GetMeshes() []*core.ResourceRef {
	if m != nil {
		return m.Meshes
	}
	return nil
}

func init() {
	proto.RegisterType((*Mesh)(nil), "zephyr.solo.io.Mesh")
	proto.RegisterType((*DiscoveryMetadata)(nil), "zephyr.solo.io.DiscoveryMetadata")
	proto.RegisterType((*IstioMesh)(nil), "zephyr.solo.io.IstioMesh")
	proto.RegisterType((*AwsAppMesh)(nil), "zephyr.solo.io.AwsAppMesh")
	proto.RegisterType((*LinkerdMesh)(nil), "zephyr.solo.io.LinkerdMesh")
	proto.RegisterType((*MtlsConfig)(nil), "zephyr.solo.io.MtlsConfig")
	proto.RegisterType((*MonitoringConfig)(nil), "zephyr.solo.io.MonitoringConfig")
	proto.RegisterType((*MeshGroup)(nil), "zephyr.solo.io.MeshGroup")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/v1/mesh.proto", fileDescriptor_0569f8953a7ffa98)
}

var fileDescriptor_0569f8953a7ffa98 = []byte{
	// 969 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4f, 0x4f, 0xdc, 0x46,
	0x14, 0x8f, 0x61, 0x59, 0xd8, 0xb7, 0x28, 0x81, 0x29, 0x20, 0x87, 0x4a, 0x01, 0xf6, 0x92, 0x48,
	0x4d, 0x76, 0x45, 0xaa, 0x36, 0x11, 0x95, 0x2a, 0x01, 0xa9, 0x92, 0x54, 0x21, 0x42, 0xa6, 0xea,
	0xa1, 0xaa, 0x6a, 0x0d, 0xf6, 0x5b, 0xef, 0x14, 0xdb, 0x33, 0x9a, 0x19, 0x2f, 0xda, 0x9e, 0x2a,
	0x3e, 0x48, 0x8f, 0x55, 0x3f, 0x4a, 0xaf, 0xbd, 0xf6, 0x90, 0x43, 0xbf, 0xc1, 0xf6, 0x03, 0x54,
	0xd5, 0x8c, 0xc7, 0x5e, 0x58, 0x42, 0x00, 0xa9, 0x87, 0x9e, 0xec, 0x79, 0xef, 0xf7, 0x7b, 0xf3,
	0xfe, 0x0f, 0x3c, 0x4d, 0x98, 0x1e, 0x14, 0xc7, 0xdd, 0x88, 0x67, 0x3d, 0xc5, 0x53, 0xfe, 0x84,
	0xf1, 0x5e, 0x86, 0x6a, 0xf0, 0x44, 0x48, 0xfe, 0x23, 0x46, 0x5a, 0xf5, 0xa8, 0x60, 0xbd, 0xe1,
	0xb6, 0x15, 0x76, 0x85, 0xe4, 0x9a, 0x93, 0xbb, 0x3f, 0xa1, 0x18, 0x8c, 0x64, 0xd7, 0xe0, 0xbb,
	0x8c, 0xaf, 0xaf, 0x24, 0x3c, 0xe1, 0x56, 0xd5, 0x33, 0x7f, 0x25, 0x6a, 0xfd, 0x41, 0xc2, 0x79,
	0x92, 0x62, 0xcf, 0x9e, 0x8e, 0x8b, 0x7e, 0x2f, 0x2e, 0x24, 0xd5, 0x8c, 0xe7, 0x57, 0xe9, 0x4f,
	0x25, 0x15, 0x02, 0xa5, 0x72, 0xfa, 0xed, 0xf7, 0x78, 0x66, 0xbf, 0x27, 0x4c, 0x4f, 0x9c, 0xd2,
	0x34, 0xa6, 0x9a, 0x3a, 0x4a, 0xef, 0x06, 0x14, 0xa5, 0xa9, 0x2e, 0xaa, 0x3b, 0x1e, 0xdf, 0x80,
	0x20, 0xb1, 0x7f, 0x0b, 0x8f, 0xaa, 0xb3, 0xa3, 0x7c, 0x7e, 0xd3, 0xf4, 0x2a, 0x4c, 0x31, 0xd2,
	0x5c, 0xde, 0x96, 0x17, 0x71, 0x89, 0x13, 0x17, 0x3b, 0xbf, 0xce, 0x41, 0xe3, 0x00, 0xd5, 0x80,
	0xbc, 0x84, 0x66, 0x19, 0xa9, 0x1f, 0x6f, 0x7a, 0x8f, 0xda, 0x4f, 0x57, 0xba, 0x06, 0x59, 0x95,
	0xac, 0x7b, 0x64, 0x75, 0x7b, 0xf7, 0x7f, 0x7f, 0xb7, 0x71, 0xe7, 0xef, 0x77, 0x1b, 0xcb, 0x1a,
	0x95, 0x8e, 0x59, 0xbf, 0xbf, 0xd3, 0x61, 0x49, 0xce, 0x25, 0x76, 0x02, 0x47, 0x27, 0xcf, 0x61,
	0xa1, 0xca, 0xb2, 0x8f, 0xd6, 0xd4, 0xda, 0x45, 0x53, 0x07, 0x4e, 0xbb, 0xd7, 0x30, 0xc6, 0x82,
	0x1a, 0x4d, 0xb6, 0x61, 0x8e, 0x29, 0xcd, 0xb8, 0xef, 0x59, 0xda, 0xfd, 0xee, 0xc5, 0xb6, 0xe9,
	0xbe, 0x36, 0x4a, 0xe3, 0xec, 0xab, 0x3b, 0x41, 0x89, 0x24, 0x5f, 0xc2, 0x22, 0x3d, 0x55, 0x21,
	0x15, 0x22, 0x34, 0xd1, 0xfa, 0x0d, 0xcb, 0x5c, 0x9f, 0x66, 0xee, 0x9e, 0xaa, 0x5d, 0x21, 0x1c,
	0x15, 0x68, 0x7d, 0x22, 0xcf, 0x60, 0x3e, 0x65, 0xf9, 0x09, 0xca, 0xd8, 0x9f, 0xb3, 0xd4, 0x8f,
	0xa7, 0xa9, 0x6f, 0x4a, 0xb5, 0xe3, 0x56, 0x68, 0xf2, 0x05, 0xb4, 0x33, 0x9d, 0xaa, 0x30, 0xe2,
	0x79, 0x9f, 0x25, 0xfe, 0xcc, 0xfb, 0xef, 0x3d, 0xd0, 0xa9, 0xda, 0xb7, 0x88, 0x00, 0xb2, 0xfa,
	0x9f, 0x1c, 0xc0, 0x72, 0xc6, 0x73, 0xa6, 0xb9, 0x64, 0x79, 0x52, 0x99, 0x98, 0xb5, 0x26, 0x36,
	0x2f, 0x99, 0xa8, 0x81, 0xce, 0xd0, 0x52, 0x36, 0x25, 0x21, 0xdf, 0x02, 0x89, 0x99, 0x8a, 0xf8,
	0x10, 0xe5, 0x28, 0xac, 0x73, 0xdf, 0xb4, 0xf6, 0xb6, 0xa6, 0xed, 0xbd, 0xa8, 0x90, 0x93, 0x32,
	0xfc, 0x3c, 0x6e, 0x78, 0xc1, 0x72, 0x3c, 0xad, 0x20, 0x1b, 0xd0, 0x56, 0x19, 0x0b, 0x31, 0xa7,
	0xc7, 0x29, 0xc6, 0xfe, 0xfc, 0xa6, 0xf7, 0x68, 0x21, 0x00, 0x95, 0xb1, 0xaf, 0x4a, 0x09, 0x79,
	0x05, 0x6d, 0xcc, 0xb5, 0x1c, 0x85, 0x82, 0xb3, 0x5c, 0xfb, 0x0b, 0xf6, 0xc6, 0x87, 0x65, 0xb5,
	0xa7, 0xae, 0xdd, 0x4f, 0x0b, 0xa5, 0x51, 0x06, 0xa8, 0x78, 0x21, 0x23, 0x0c, 0xb0, 0x1f, 0x80,
	0xe5, 0x1e, 0x1a, 0xea, 0xce, 0x47, 0x67, 0xe3, 0xc6, 0x2c, 0x78, 0xd9, 0xd9, 0xb8, 0xb1, 0x40,
	0x9a, 0xa6, 0x94, 0xa8, 0xf6, 0xda, 0xd0, 0x32, 0x7f, 0xa1, 0x1e, 0x09, 0xec, 0xfc, 0xe1, 0xc1,
	0xf2, 0x25, 0xdf, 0x89, 0x0f, 0xf3, 0x51, 0x69, 0xd9, 0x36, 0x4d, 0x2b, 0xa8, 0x8e, 0xe4, 0x31,
	0x90, 0xd2, 0xf1, 0x90, 0x16, 0x9a, 0x87, 0x2c, 0x37, 0x53, 0x60, 0xeb, 0xb4, 0x10, 0x2c, 0x95,
	0x9a, 0xdd, 0x42, 0xf3, 0xd7, 0x56, 0x4e, 0x9e, 0x41, 0xab, 0x10, 0x4a, 0x4b, 0xa4, 0x99, 0xf2,
	0xe7, 0x36, 0x67, 0x6d, 0xfb, 0x5d, 0xe8, 0xda, 0xf3, 0x9e, 0x4f, 0xb0, 0xd3, 0x7d, 0xd0, 0xbc,
	0x4d, 0x1f, 0x74, 0xbe, 0x87, 0x56, 0xdd, 0xd3, 0xe4, 0x33, 0x58, 0x63, 0xb9, 0xd2, 0x34, 0x4d,
	0xed, 0xd2, 0x0b, 0x73, 0x9a, 0xa1, 0x12, 0x34, 0x42, 0x17, 0xd9, 0xea, 0x79, 0xed, 0xdb, 0x4a,
	0x69, 0x32, 0x30, 0x44, 0xa9, 0x18, 0xcf, 0x6d, 0x70, 0xad, 0xa0, 0x3a, 0x76, 0xfe, 0x9c, 0x01,
	0x98, 0x34, 0x3e, 0x79, 0x0e, 0xa6, 0xf1, 0x43, 0x85, 0x91, 0x44, 0x5d, 0x8f, 0xd8, 0xd5, 0x31,
	0xd2, 0x53, 0x75, 0x64, 0xb1, 0x64, 0x0d, 0x9a, 0x12, 0x93, 0xc9, 0x0d, 0xee, 0x74, 0x45, 0x8a,
	0x67, 0xaf, 0x48, 0xf1, 0xd7, 0x40, 0x4a, 0x84, 0x09, 0xae, 0xda, 0x5e, 0x6e, 0x60, 0x2f, 0x4d,
	0xdd, 0x21, 0x8f, 0x8f, 0x1c, 0x24, 0x58, 0xae, 0x69, 0x95, 0xc8, 0xdc, 0x3c, 0x64, 0x52, 0x17,
	0x34, 0x0d, 0x73, 0x1e, 0x63, 0x98, 0xd2, 0x63, 0x4c, 0xed, 0x04, 0xb7, 0x82, 0x25, 0xa7, 0x79,
	0xcb, 0x63, 0x7c, 0x63, 0xe4, 0x24, 0x00, 0x5f, 0xb1, 0x18, 0x23, 0x2a, 0x43, 0x41, 0x75, 0x34,
	0x70, 0xc5, 0x0a, 0x33, 0x2a, 0x5c, 0xc1, 0x3e, 0x90, 0x87, 0x55, 0x47, 0x3d, 0x34, 0xcc, 0xb2,
	0x6e, 0x07, 0x54, 0x74, 0x7e, 0x80, 0xf6, 0xb9, 0xcd, 0xf0, 0xdf, 0x17, 0xef, 0x17, 0x0f, 0x60,
	0xd2, 0x35, 0x64, 0x0b, 0x16, 0x6d, 0x9b, 0x55, 0xb3, 0xe8, 0xd9, 0x24, 0xdb, 0xd6, 0xab, 0x86,
	0xf1, 0x21, 0xdc, 0xb3, 0x10, 0x81, 0x32, 0x63, 0x4a, 0xb1, 0x21, 0xba, 0x6e, 0xbf, 0x6b, 0xc4,
	0x87, 0xb5, 0x94, 0xbc, 0x80, 0x25, 0xc9, 0xb9, 0x0e, 0x23, 0x94, 0x9a, 0xf5, 0x59, 0x44, 0x35,
	0xba, 0xe5, 0xf3, 0x81, 0x34, 0xdc, 0x33, 0x94, 0xfd, 0x09, 0xa3, 0x33, 0x80, 0xa5, 0xe9, 0xd5,
	0x44, 0xbe, 0x81, 0x55, 0x21, 0x79, 0x86, 0x7a, 0x80, 0x45, 0x35, 0x12, 0x19, 0x15, 0xca, 0xf7,
	0xae, 0x99, 0x28, 0xf7, 0x14, 0xac, 0x4c, 0xd8, 0xfb, 0x35, 0xb9, 0xf3, 0x8f, 0x07, 0x2d, 0x93,
	0xe4, 0x97, 0x92, 0x17, 0xe2, 0xff, 0xf0, 0x4e, 0x6d, 0xc1, 0x62, 0xcc, 0x94, 0x48, 0xe9, 0xc8,
	0xd6, 0xd9, 0x95, 0xb8, 0xed, 0x64, 0xa6, 0xba, 0x64, 0x1b, 0xdc, 0x12, 0xf3, 0x67, 0xaf, 0x5b,
	0x26, 0x0e, 0xb8, 0xe3, 0x9f, 0x8d, 0x1b, 0x0d, 0x98, 0xc9, 0x92, 0xb3, 0x71, 0x63, 0x91, 0x80,
	0x91, 0x26, 0x26, 0x62, 0xb5, 0xb7, 0xfd, 0xdb, 0x5f, 0x0f, 0xbc, 0xef, 0x3e, 0xb9, 0xf6, 0x85,
	0x17, 0x27, 0x89, 0x7b, 0xe5, 0x8f, 0x9b, 0xf6, 0x75, 0xff, 0xf4, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0xa7, 0xa1, 0xab, 0xae, 0x09, 0x00, 0x00,
}

func (this *Mesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh)
	if !ok {
		that2, ok := that.(Mesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if that1.MeshType == nil {
		if this.MeshType != nil {
			return false
		}
	} else if this.MeshType == nil {
		return false
	} else if !this.MeshType.Equal(that1.MeshType) {
		return false
	}
	if !this.MtlsConfig.Equal(that1.MtlsConfig) {
		return false
	}
	if !this.MonitoringConfig.Equal(that1.MonitoringConfig) {
		return false
	}
	if !this.DiscoveryMetadata.Equal(that1.DiscoveryMetadata) {
		return false
	}
	if this.SmiEnabled != that1.SmiEnabled {
		return false
	}
	if !this.EntryPoint.Equal(that1.EntryPoint) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Mesh_Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_Istio)
	if !ok {
		that2, ok := that.(Mesh_Istio)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Istio.Equal(that1.Istio) {
		return false
	}
	return true
}
func (this *Mesh_AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_AwsAppMesh)
	if !ok {
		that2, ok := that.(Mesh_AwsAppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AwsAppMesh.Equal(that1.AwsAppMesh) {
		return false
	}
	return true
}
func (this *Mesh_Linkerd) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Mesh_Linkerd)
	if !ok {
		that2, ok := that.(Mesh_Linkerd)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Linkerd.Equal(that1.Linkerd) {
		return false
	}
	return true
}
func (this *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Cluster != that1.Cluster {
		return false
	}
	if this.EnableAutoInject != that1.EnableAutoInject {
		return false
	}
	if len(this.Upstreams) != len(that1.Upstreams) {
		return false
	}
	for i := range this.Upstreams {
		if !this.Upstreams[i].Equal(that1.Upstreams[i]) {
			return false
		}
	}
	if !this.MtlsConfig.Equal(that1.MtlsConfig) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IstioMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IstioMesh)
	if !ok {
		that2, ok := that.(IstioMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AwsAppMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsAppMesh)
	if !ok {
		that2, ok := that.(AwsAppMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AwsSecret.Equal(that1.AwsSecret) {
		return false
	}
	if this.Region != that1.Region {
		return false
	}
	if this.EnableAutoInject != that1.EnableAutoInject {
		return false
	}
	if !this.InjectionSelector.Equal(that1.InjectionSelector) {
		return false
	}
	if this.VirtualNodeLabel != that1.VirtualNodeLabel {
		return false
	}
	if !this.SidecarPatchConfigMap.Equal(that1.SidecarPatchConfigMap) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *LinkerdMesh) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LinkerdMesh)
	if !ok {
		that2, ok := that.(LinkerdMesh)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.InstallationNamespace != that1.InstallationNamespace {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MtlsConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MtlsConfig)
	if !ok {
		that2, ok := that.(MtlsConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MtlsEnabled != that1.MtlsEnabled {
		return false
	}
	if this.MtlsPermissive != that1.MtlsPermissive {
		return false
	}
	if !this.RootCertificate.Equal(that1.RootCertificate) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MonitoringConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MonitoringConfig)
	if !ok {
		that2, ok := that.(MonitoringConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.PrometheusConfigmaps) != len(that1.PrometheusConfigmaps) {
		return false
	}
	for i := range this.PrometheusConfigmaps {
		if !this.PrometheusConfigmaps[i].Equal(&that1.PrometheusConfigmaps[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshGroup) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshGroup)
	if !ok {
		that2, ok := that.(MeshGroup)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if this.DisplayName != that1.DisplayName {
		return false
	}
	if len(this.Meshes) != len(that1.Meshes) {
		return false
	}
	for i := range this.Meshes {
		if !this.Meshes[i].Equal(that1.Meshes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
