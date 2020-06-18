// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/security/v1alpha1/certificates.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types1 "github.com/gogo/protobuf/types"
	types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
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

type VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus int32

const (
	// have a default value which represents not being set as proto enums require a default 0th value
	VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_PENDING  VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus = 0
	VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_APPROVED VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus = 1
	VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_DENIED   VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus = 2
)

var VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus_name = map[int32]string{
	0: "PENDING",
	1: "APPROVED",
	2: "DENIED",
}

var VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus_value = map[string]int32{
	"PENDING":  0,
	"APPROVED": 1,
	"DENIED":   2,
}

func (x VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus) String() string {
	return proto.EnumName(VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus_name, int32(x))
}

func (VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_09ccef6a2ee48f77, []int{1, 1, 0}
}

type VirtualMeshCertificateSigningRequestSpec struct {
	// Base64-encoded PKCS#10 CSR data
	CsrData    []byte                                               `protobuf:"bytes,1,opt,name=csr_data,json=csrData,proto3" json:"csr_data,omitempty"`
	CertConfig *VirtualMeshCertificateSigningRequestSpec_CertConfig `protobuf:"bytes,2,opt,name=cert_config,json=certConfig,proto3" json:"cert_config,omitempty"`
	//
	//Reference to the virtual mesh which this CSR corresponds to. This is important as it allows the virtual mesh
	//operator to know which trust bundle to use when signing the new certificates.
	//
	//When the CSR is first created by the Virtual Mesh operator, this data will be added by it. However, during a cert
	//rotation scenario this is not possible. Therefore, the csr-agent will write this data to the secret so that
	//it can be retrieved when the cert is going to expire.
	//TODO: Decide how exactly we want to store this data.
	VirtualMeshRef       *types.ResourceRef `protobuf:"bytes,3,opt,name=virtual_mesh_ref,json=virtualMeshRef,proto3" json:"virtual_mesh_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VirtualMeshCertificateSigningRequestSpec) Reset() {
	*m = VirtualMeshCertificateSigningRequestSpec{}
}
func (m *VirtualMeshCertificateSigningRequestSpec) String() string { return proto.CompactTextString(m) }
func (*VirtualMeshCertificateSigningRequestSpec) ProtoMessage()    {}
func (*VirtualMeshCertificateSigningRequestSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ccef6a2ee48f77, []int{0}
}
func (m *VirtualMeshCertificateSigningRequestSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec.Unmarshal(m, b)
}
func (m *VirtualMeshCertificateSigningRequestSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec.Marshal(b, m, deterministic)
}
func (m *VirtualMeshCertificateSigningRequestSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec.Merge(m, src)
}
func (m *VirtualMeshCertificateSigningRequestSpec) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec.Size(m)
}
func (m *VirtualMeshCertificateSigningRequestSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec proto.InternalMessageInfo

func (m *VirtualMeshCertificateSigningRequestSpec) GetCsrData() []byte {
	if m != nil {
		return m.CsrData
	}
	return nil
}

func (m *VirtualMeshCertificateSigningRequestSpec) GetCertConfig() *VirtualMeshCertificateSigningRequestSpec_CertConfig {
	if m != nil {
		return m.CertConfig
	}
	return nil
}

func (m *VirtualMeshCertificateSigningRequestSpec) GetVirtualMeshRef() *types.ResourceRef {
	if m != nil {
		return m.VirtualMeshRef
	}
	return nil
}

type VirtualMeshCertificateSigningRequestSpec_CertConfig struct {
	//
	//list of hostnames and IPs to generate a certificate for.
	//This can also be set to the identity running the workload,
	//like kubernetes service account.
	//
	//Generally for an Istio CA this will take the form `spiffe://cluster.local/ns/istio-system/sa/citadel`.
	//
	//"cluster.local" may be replaced by the root of trust domain for the mesh.
	Hosts []string `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Organization for this certificate.
	Org string `protobuf:"bytes,3,opt,name=org,proto3" json:"org,omitempty"`
	//
	//In the future, the type of mesh, and level of trust will need to be specified here,
	//but for the time being we are only supporting shared trust in Istio.
	MeshType             types.MeshType `protobuf:"varint,5,opt,name=mesh_type,json=meshType,proto3,enum=core.smh.solo.io.MeshType" json:"mesh_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) Reset() {
	*m = VirtualMeshCertificateSigningRequestSpec_CertConfig{}
}
func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshCertificateSigningRequestSpec_CertConfig) ProtoMessage() {}
func (*VirtualMeshCertificateSigningRequestSpec_CertConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ccef6a2ee48f77, []int{0, 0}
}
func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec_CertConfig.Unmarshal(m, b)
}
func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec_CertConfig.Marshal(b, m, deterministic)
}
func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec_CertConfig.Merge(m, src)
}
func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec_CertConfig.Size(m)
}
func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec_CertConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshCertificateSigningRequestSpec_CertConfig proto.InternalMessageInfo

func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) GetOrg() string {
	if m != nil {
		return m.Org
	}
	return ""
}

func (m *VirtualMeshCertificateSigningRequestSpec_CertConfig) GetMeshType() types.MeshType {
	if m != nil {
		return m.MeshType
	}
	return types.MeshType_ISTIO1_5
}

type VirtualMeshCertificateSigningRequestStatus struct {
	// Response from the certificate authority
	Response *VirtualMeshCertificateSigningRequestStatus_Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	// Workflow for approving Certificate Signing Requests
	ThirdPartyApproval   *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow `protobuf:"bytes,2,opt,name=third_party_approval,json=thirdPartyApproval,proto3" json:"third_party_approval,omitempty"`
	ComputedStatus       *types.Status                                                          `protobuf:"bytes,3,opt,name=computed_status,json=computedStatus,proto3" json:"computed_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                               `json:"-"`
	XXX_unrecognized     []byte                                                                 `json:"-"`
	XXX_sizecache        int32                                                                  `json:"-"`
}

func (m *VirtualMeshCertificateSigningRequestStatus) Reset() {
	*m = VirtualMeshCertificateSigningRequestStatus{}
}
func (m *VirtualMeshCertificateSigningRequestStatus) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshCertificateSigningRequestStatus) ProtoMessage() {}
func (*VirtualMeshCertificateSigningRequestStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ccef6a2ee48f77, []int{1}
}
func (m *VirtualMeshCertificateSigningRequestStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus.Unmarshal(m, b)
}
func (m *VirtualMeshCertificateSigningRequestStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus.Marshal(b, m, deterministic)
}
func (m *VirtualMeshCertificateSigningRequestStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus.Merge(m, src)
}
func (m *VirtualMeshCertificateSigningRequestStatus) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus.Size(m)
}
func (m *VirtualMeshCertificateSigningRequestStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus proto.InternalMessageInfo

func (m *VirtualMeshCertificateSigningRequestStatus) GetResponse() *VirtualMeshCertificateSigningRequestStatus_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *VirtualMeshCertificateSigningRequestStatus) GetThirdPartyApproval() *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow {
	if m != nil {
		return m.ThirdPartyApproval
	}
	return nil
}

func (m *VirtualMeshCertificateSigningRequestStatus) GetComputedStatus() *types.Status {
	if m != nil {
		return m.ComputedStatus
	}
	return nil
}

type VirtualMeshCertificateSigningRequestStatus_Response struct {
	// If request was approved, the controller will place the issued certificate here.
	CaCertificate []byte `protobuf:"bytes,1,opt,name=ca_certificate,json=caCertificate,proto3" json:"ca_certificate,omitempty"`
	// root cert shared by all clusters, safe to send over the wire
	RootCertificate      []byte   `protobuf:"bytes,2,opt,name=root_certificate,json=rootCertificate,proto3" json:"root_certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VirtualMeshCertificateSigningRequestStatus_Response) Reset() {
	*m = VirtualMeshCertificateSigningRequestStatus_Response{}
}
func (m *VirtualMeshCertificateSigningRequestStatus_Response) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshCertificateSigningRequestStatus_Response) ProtoMessage() {}
func (*VirtualMeshCertificateSigningRequestStatus_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ccef6a2ee48f77, []int{1, 0}
}
func (m *VirtualMeshCertificateSigningRequestStatus_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_Response.Unmarshal(m, b)
}
func (m *VirtualMeshCertificateSigningRequestStatus_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_Response.Marshal(b, m, deterministic)
}
func (m *VirtualMeshCertificateSigningRequestStatus_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_Response.Merge(m, src)
}
func (m *VirtualMeshCertificateSigningRequestStatus_Response) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_Response.Size(m)
}
func (m *VirtualMeshCertificateSigningRequestStatus_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_Response.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_Response proto.InternalMessageInfo

func (m *VirtualMeshCertificateSigningRequestStatus_Response) GetCaCertificate() []byte {
	if m != nil {
		return m.CaCertificate
	}
	return nil
}

func (m *VirtualMeshCertificateSigningRequestStatus_Response) GetRootCertificate() []byte {
	if m != nil {
		return m.RootCertificate
	}
	return nil
}

type VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow struct {
	// time when the status was last updated
	LastUpdatedTime *types1.Timestamp `protobuf:"bytes,2,opt,name=last_updated_time,json=lastUpdatedTime,proto3" json:"last_updated_time,omitempty"`
	// a user readable message regarding the status of the CSR
	Message              string                                                                               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	ApprovalStatus       VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus `protobuf:"varint,4,opt,name=approval_status,json=approvalStatus,proto3,enum=security.smh.solo.io.VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus" json:"approval_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                             `json:"-"`
	XXX_unrecognized     []byte                                                                               `json:"-"`
	XXX_sizecache        int32                                                                                `json:"-"`
}

func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) Reset() {
	*m = VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow{}
}
func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) String() string {
	return proto.CompactTextString(m)
}
func (*VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) ProtoMessage() {}
func (*VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_09ccef6a2ee48f77, []int{1, 1}
}
func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow.Unmarshal(m, b)
}
func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow.Marshal(b, m, deterministic)
}
func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow.Merge(m, src)
}
func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) XXX_Size() int {
	return xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow.Size(m)
}
func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow proto.InternalMessageInfo

func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) GetLastUpdatedTime() *types1.Timestamp {
	if m != nil {
		return m.LastUpdatedTime
	}
	return nil
}

func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) GetApprovalStatus() VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus {
	if m != nil {
		return m.ApprovalStatus
	}
	return VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_PENDING
}

func init() {
	proto.RegisterEnum("security.smh.solo.io.VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus", VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus_name, VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_ApprovalStatus_value)
	proto.RegisterType((*VirtualMeshCertificateSigningRequestSpec)(nil), "security.smh.solo.io.VirtualMeshCertificateSigningRequestSpec")
	proto.RegisterType((*VirtualMeshCertificateSigningRequestSpec_CertConfig)(nil), "security.smh.solo.io.VirtualMeshCertificateSigningRequestSpec.CertConfig")
	proto.RegisterType((*VirtualMeshCertificateSigningRequestStatus)(nil), "security.smh.solo.io.VirtualMeshCertificateSigningRequestStatus")
	proto.RegisterType((*VirtualMeshCertificateSigningRequestStatus_Response)(nil), "security.smh.solo.io.VirtualMeshCertificateSigningRequestStatus.Response")
	proto.RegisterType((*VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow)(nil), "security.smh.solo.io.VirtualMeshCertificateSigningRequestStatus.ThirdPartyApprovalWorkflow")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/security/v1alpha1/certificates.proto", fileDescriptor_09ccef6a2ee48f77)
}

var fileDescriptor_09ccef6a2ee48f77 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x9c, 0x7e, 0x6d, 0xd3, 0x6d, 0x49, 0xc2, 0x2a, 0x17, 0xc1, 0x12, 0x25, 0xaa, 0x84,
	0x14, 0x10, 0xb1, 0xd5, 0x72, 0xd1, 0xeb, 0xd2, 0x84, 0xaa, 0x12, 0xb4, 0xd1, 0xf6, 0x07, 0x09,
	0x90, 0xac, 0xcd, 0x66, 0x6c, 0x6f, 0x6b, 0x67, 0xcd, 0xee, 0x3a, 0x28, 0x0f, 0xc1, 0x15, 0x3c,
	0x04, 0x2f, 0xc3, 0x4b, 0x70, 0xc7, 0x5b, 0xa0, 0xf5, 0x4f, 0x93, 0xd0, 0x22, 0x45, 0xaa, 0xb8,
	0x9b, 0xd9, 0x3d, 0x7b, 0x8e, 0xe7, 0xcc, 0x8c, 0xd1, 0x9b, 0x80, 0xeb, 0x30, 0x1d, 0x3a, 0x4c,
	0xc4, 0xae, 0x12, 0x91, 0xe8, 0x72, 0xe1, 0x2a, 0x90, 0x13, 0xce, 0xa0, 0x1b, 0x83, 0x0a, 0xbb,
	0x61, 0x3a, 0x74, 0x69, 0xc2, 0x5d, 0x05, 0x2c, 0x95, 0x5c, 0x4f, 0xdd, 0xc9, 0x2e, 0x8d, 0x92,
	0x90, 0xee, 0xba, 0x0c, 0xa4, 0xe6, 0x3e, 0x67, 0x54, 0x83, 0x72, 0x12, 0x29, 0xb4, 0xc0, 0xcd,
	0x12, 0xe5, 0xa8, 0x38, 0x74, 0x0c, 0x9f, 0xc3, 0x85, 0xfd, 0xe2, 0x4e, 0x42, 0x26, 0x24, 0xcc,
	0xc8, 0x24, 0xf8, 0x39, 0x87, 0xed, 0x2e, 0x81, 0x56, 0x9a, 0xea, 0xb4, 0x10, 0xb5, 0xbb, 0x4b,
	0x3c, 0x30, 0x57, 0x05, 0xfc, 0x49, 0x20, 0x44, 0x10, 0x81, 0x9b, 0x65, 0xc3, 0xd4, 0x77, 0x35,
	0x8f, 0x41, 0x69, 0x1a, 0x27, 0x05, 0x60, 0xfb, 0x4f, 0xc0, 0x28, 0x95, 0x54, 0x73, 0x31, 0x2e,
	0xee, 0x9b, 0x81, 0x08, 0x44, 0x16, 0xba, 0x26, 0xca, 0x4f, 0x77, 0x7e, 0x55, 0x50, 0xe7, 0x92,
	0x4b, 0x9d, 0xd2, 0xe8, 0x2d, 0xa8, 0xf0, 0x70, 0x66, 0xce, 0x19, 0x0f, 0xc6, 0x7c, 0x1c, 0x10,
	0xf8, 0x94, 0x82, 0xd2, 0x67, 0x09, 0x30, 0xfc, 0x08, 0x55, 0x99, 0x92, 0xde, 0x88, 0x6a, 0xda,
	0xb2, 0xda, 0x56, 0x67, 0x8b, 0xac, 0x33, 0x25, 0x7b, 0x54, 0x53, 0x7c, 0x85, 0x36, 0x8d, 0xb1,
	0x1e, 0x13, 0x63, 0x9f, 0x07, 0xad, 0x4a, 0xdb, 0xea, 0x6c, 0xee, 0x1d, 0x3b, 0x77, 0x19, 0xeb,
	0x2c, 0xab, 0xe7, 0x98, 0xdb, 0xc3, 0x8c, 0x90, 0x20, 0x76, 0x13, 0xe3, 0x23, 0xd4, 0x98, 0xe4,
	0x14, 0x9e, 0x31, 0xc8, 0x93, 0xe0, 0xb7, 0x56, 0x32, 0xc1, 0xc7, 0x8e, 0xf1, 0x6f, 0x41, 0x8c,
	0x80, 0x12, 0xa9, 0x64, 0x40, 0xc0, 0x27, 0xb5, 0xc9, 0x4c, 0x99, 0x80, 0x6f, 0xc7, 0x08, 0xcd,
	0x24, 0x70, 0x13, 0xad, 0x86, 0x42, 0x69, 0xd5, 0xaa, 0xb4, 0x57, 0x3a, 0x1b, 0x24, 0x4f, 0x70,
	0x03, 0xad, 0x08, 0x19, 0x64, 0xfc, 0x1b, 0xc4, 0x84, 0x78, 0x1f, 0x6d, 0x64, 0xb2, 0x7a, 0x9a,
	0x40, 0x6b, 0xb5, 0x6d, 0x75, 0x6a, 0x7b, 0xf6, 0x6d, 0x5d, 0xa3, 0x71, 0x3e, 0x4d, 0x80, 0x54,
	0xe3, 0x22, 0xda, 0xf9, 0xb6, 0x86, 0x9e, 0x2f, 0x55, 0x7b, 0x36, 0x26, 0x18, 0x50, 0x55, 0x82,
	0x4a, 0xc4, 0x58, 0x41, 0xe6, 0xf6, 0xfd, 0xfc, 0xcc, 0x47, 0x8f, 0x14, 0x84, 0xe4, 0x86, 0x1a,
	0x7f, 0xb1, 0x50, 0x53, 0x87, 0x5c, 0x8e, 0xbc, 0x84, 0x4a, 0x3d, 0xf5, 0x68, 0x92, 0x48, 0x31,
	0xa1, 0x51, 0xd1, 0xc3, 0x0f, 0xf7, 0xd6, 0x3c, 0x37, 0xe4, 0x03, 0xc3, 0x7d, 0x50, 0x50, 0xbf,
	0x13, 0xf2, 0xda, 0x8f, 0xc4, 0x67, 0x82, 0xf5, 0xad, 0x3b, 0x7c, 0x80, 0xea, 0x4c, 0xc4, 0x49,
	0xaa, 0x61, 0xe4, 0xe5, 0x0b, 0x53, 0x34, 0xb7, 0x75, 0xdb, 0xe4, 0x5c, 0x81, 0xd4, 0xca, 0x07,
	0x79, 0x6e, 0x7f, 0x44, 0xd5, 0xb2, 0x50, 0xfc, 0x14, 0xd5, 0x18, 0xf5, 0xe6, 0x96, 0xbe, 0x98,
	0xdc, 0x07, 0x8c, 0xce, 0x7d, 0x38, 0x7e, 0x86, 0x1a, 0x52, 0x08, 0xbd, 0x00, 0xac, 0x64, 0xc0,
	0xba, 0x39, 0x9f, 0x83, 0xda, 0x3f, 0x2a, 0xc8, 0xfe, 0x7b, 0x4d, 0xf8, 0x35, 0x7a, 0x18, 0x51,
	0xa5, 0xbd, 0x34, 0x19, 0x51, 0x53, 0x83, 0xd9, 0xd3, 0xc2, 0x4b, 0xdb, 0xc9, 0x77, 0xd4, 0x29,
	0x77, 0xd4, 0x39, 0x2f, 0x97, 0x98, 0xd4, 0xcd, 0xa3, 0x8b, 0xfc, 0x8d, 0x39, 0xc5, 0x2d, 0xb4,
	0x1e, 0x83, 0x52, 0x34, 0x80, 0x62, 0xf8, 0xca, 0x14, 0x7f, 0xb5, 0x50, 0xbd, 0xec, 0x52, 0x69,
	0xd1, 0xff, 0xd9, 0x1c, 0x5e, 0xfd, 0xc3, 0x66, 0x39, 0xe5, 0x41, 0x69, 0x3a, 0x5d, 0xc8, 0x77,
	0xf6, 0x51, 0x6d, 0x11, 0x81, 0x37, 0xd1, 0xfa, 0xa0, 0x7f, 0xd2, 0x3b, 0x3e, 0x39, 0x6a, 0xfc,
	0x87, 0xb7, 0x50, 0xf5, 0x60, 0x30, 0x20, 0xa7, 0x97, 0xfd, 0x5e, 0xc3, 0xc2, 0x08, 0xad, 0xf5,
	0xfa, 0x27, 0xc7, 0xfd, 0x5e, 0xa3, 0xf2, 0xea, 0xe2, 0xfb, 0xcf, 0x6d, 0xeb, 0xfd, 0xe9, 0x32,
	0x7f, 0xf4, 0xe4, 0x3a, 0x58, 0xf8, 0xab, 0xcf, 0x57, 0x39, 0xfb, 0x6b, 0x9a, 0xc5, 0x54, 0xc3,
	0xb5, 0xcc, 0xe4, 0x97, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x40, 0x06, 0xc0, 0x47, 0x2b, 0x06,
	0x00, 0x00,
}

func (this *VirtualMeshCertificateSigningRequestSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshCertificateSigningRequestSpec)
	if !ok {
		that2, ok := that.(VirtualMeshCertificateSigningRequestSpec)
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
	if !bytes.Equal(this.CsrData, that1.CsrData) {
		return false
	}
	if !this.CertConfig.Equal(that1.CertConfig) {
		return false
	}
	if !this.VirtualMeshRef.Equal(that1.VirtualMeshRef) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshCertificateSigningRequestSpec_CertConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshCertificateSigningRequestSpec_CertConfig)
	if !ok {
		that2, ok := that.(VirtualMeshCertificateSigningRequestSpec_CertConfig)
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
	if len(this.Hosts) != len(that1.Hosts) {
		return false
	}
	for i := range this.Hosts {
		if this.Hosts[i] != that1.Hosts[i] {
			return false
		}
	}
	if this.Org != that1.Org {
		return false
	}
	if this.MeshType != that1.MeshType {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshCertificateSigningRequestStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshCertificateSigningRequestStatus)
	if !ok {
		that2, ok := that.(VirtualMeshCertificateSigningRequestStatus)
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
	if !this.Response.Equal(that1.Response) {
		return false
	}
	if !this.ThirdPartyApproval.Equal(that1.ThirdPartyApproval) {
		return false
	}
	if !this.ComputedStatus.Equal(that1.ComputedStatus) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshCertificateSigningRequestStatus_Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshCertificateSigningRequestStatus_Response)
	if !ok {
		that2, ok := that.(VirtualMeshCertificateSigningRequestStatus_Response)
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
	if !bytes.Equal(this.CaCertificate, that1.CaCertificate) {
		return false
	}
	if !bytes.Equal(this.RootCertificate, that1.RootCertificate) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow)
	if !ok {
		that2, ok := that.(VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow)
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
	if !this.LastUpdatedTime.Equal(that1.LastUpdatedTime) {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if this.ApprovalStatus != that1.ApprovalStatus {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
