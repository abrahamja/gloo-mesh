// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/certificates/pod_bounce_directive.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
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
//When certificates are issued, pods may need to be bounced (restarted) to ensure they pick up the
//new certificates. If so, the certificate Issuer will create a PodBounceDirective containing the namespaces and labels
//of the pods that need to be bounced in order to pick up the new certs.
type PodBounceDirectiveSpec struct {
	// A list of k8s pods to bounce (delete and cause a restart)
	// when the certificate is issued.
	// This will include the control plane pods as well as any pods
	// which share a data plane with the target mesh.
	PodsToBounce         []*PodBounceDirectiveSpec_PodSelector `protobuf:"bytes,6,rep,name=pods_to_bounce,json=podsToBounce,proto3" json:"pods_to_bounce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PodBounceDirectiveSpec) Reset()         { *m = PodBounceDirectiveSpec{} }
func (m *PodBounceDirectiveSpec) String() string { return proto.CompactTextString(m) }
func (*PodBounceDirectiveSpec) ProtoMessage()    {}
func (*PodBounceDirectiveSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f91a49be4602a86, []int{0}
}
func (m *PodBounceDirectiveSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodBounceDirectiveSpec.Unmarshal(m, b)
}
func (m *PodBounceDirectiveSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodBounceDirectiveSpec.Marshal(b, m, deterministic)
}
func (m *PodBounceDirectiveSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodBounceDirectiveSpec.Merge(m, src)
}
func (m *PodBounceDirectiveSpec) XXX_Size() int {
	return xxx_messageInfo_PodBounceDirectiveSpec.Size(m)
}
func (m *PodBounceDirectiveSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PodBounceDirectiveSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PodBounceDirectiveSpec proto.InternalMessageInfo

func (m *PodBounceDirectiveSpec) GetPodsToBounce() []*PodBounceDirectiveSpec_PodSelector {
	if m != nil {
		return m.PodsToBounce
	}
	return nil
}

// Pods that will be restarted.
type PodBounceDirectiveSpec_PodSelector struct {
	// The namespace in which the pods live.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Any labels shared by the pods.
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PodBounceDirectiveSpec_PodSelector) Reset()         { *m = PodBounceDirectiveSpec_PodSelector{} }
func (m *PodBounceDirectiveSpec_PodSelector) String() string { return proto.CompactTextString(m) }
func (*PodBounceDirectiveSpec_PodSelector) ProtoMessage()    {}
func (*PodBounceDirectiveSpec_PodSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f91a49be4602a86, []int{0, 0}
}
func (m *PodBounceDirectiveSpec_PodSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodBounceDirectiveSpec_PodSelector.Unmarshal(m, b)
}
func (m *PodBounceDirectiveSpec_PodSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodBounceDirectiveSpec_PodSelector.Marshal(b, m, deterministic)
}
func (m *PodBounceDirectiveSpec_PodSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodBounceDirectiveSpec_PodSelector.Merge(m, src)
}
func (m *PodBounceDirectiveSpec_PodSelector) XXX_Size() int {
	return xxx_messageInfo_PodBounceDirectiveSpec_PodSelector.Size(m)
}
func (m *PodBounceDirectiveSpec_PodSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_PodBounceDirectiveSpec_PodSelector.DiscardUnknown(m)
}

var xxx_messageInfo_PodBounceDirectiveSpec_PodSelector proto.InternalMessageInfo

func (m *PodBounceDirectiveSpec_PodSelector) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PodBounceDirectiveSpec_PodSelector) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*PodBounceDirectiveSpec)(nil), "certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec")
	proto.RegisterType((*PodBounceDirectiveSpec_PodSelector)(nil), "certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector")
	proto.RegisterMapType((map[string]string)(nil), "certificates.mesh.gloo.solo.io.PodBounceDirectiveSpec.PodSelector.LabelsEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-mesh/api/certificates/pod_bounce_directive.proto", fileDescriptor_6f91a49be4602a86)
}

var fileDescriptor_6f91a49be4602a86 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x91, 0x4c, 0x0d, 0x5e, 0x97, 0x52, 0x84, 0x29, 0x42, 0x2d, 0xae, 0xe9, 0xc9, 0x17,
	0xef, 0x62, 0xf7, 0xd2, 0xf6, 0x68, 0x5a, 0xe8, 0xc1, 0x84, 0x20, 0xe7, 0x94, 0x8b, 0x59, 0xad,
	0xc6, 0xd2, 0xe2, 0x95, 0x66, 0xd1, 0xae, 0x04, 0x7e, 0xa3, 0x3c, 0x4d, 0x2e, 0x79, 0x83, 0x3c,
	0x49, 0xd0, 0x4a, 0x21, 0x26, 0x0e, 0xbe, 0xe4, 0xa4, 0xd9, 0x7f, 0xfe, 0xf9, 0xf8, 0x07, 0x0d,
	0xf9, 0x9f, 0x49, 0x9b, 0xd7, 0x09, 0x15, 0x58, 0x30, 0x83, 0x0a, 0x17, 0x12, 0x59, 0xa6, 0x10,
	0x17, 0x05, 0x98, 0x9c, 0x71, 0x2d, 0x99, 0x80, 0xca, 0xca, 0xbd, 0x14, 0xdc, 0x82, 0x61, 0x1a,
	0xd3, 0x5d, 0x82, 0x75, 0x29, 0x60, 0x97, 0xca, 0x0a, 0x84, 0x95, 0x0d, 0x50, 0x5d, 0xa1, 0xc5,
	0x60, 0x7a, 0x6a, 0xa4, 0xed, 0x38, 0x6d, 0x41, 0xb4, 0xa5, 0x52, 0x89, 0xd1, 0x57, 0x73, 0x68,
	0x56, 0x1d, 0x11, 0x2b, 0x60, 0xcd, 0xd2, 0x7d, 0xbb, 0xe1, 0xe8, 0x7b, 0x86, 0x98, 0x29, 0x60,
	0xee, 0x95, 0xd4, 0x7b, 0x66, 0x65, 0x01, 0xc6, 0xf2, 0x42, 0xf7, 0x86, 0xe9, 0x6b, 0x43, 0x5a,
	0x57, 0xdc, 0x4a, 0x2c, 0xfb, 0xfe, 0x24, 0xc3, 0x0c, 0x5d, 0xc9, 0xda, 0xaa, 0x53, 0x7f, 0x3c,
	0xf8, 0xe4, 0xcb, 0x35, 0xa6, 0x6b, 0x97, 0xf8, 0xef, 0x73, 0xe0, 0xad, 0x06, 0x11, 0xe4, 0xe4,
	0x93, 0xc6, 0xd4, 0xec, 0x2c, 0xf6, 0x0b, 0x85, 0xc3, 0xd9, 0x60, 0x3e, 0x5e, 0xad, 0xe9, 0xe5,
	0x3d, 0xe8, 0xdb, 0xbc, 0x56, 0xde, 0x82, 0x02, 0x61, 0xb1, 0x8a, 0x3f, 0xb6, 0xe4, 0x1b, 0xec,
	0x6c, 0xd1, 0xbd, 0x47, 0xc6, 0x27, 0xdd, 0xe0, 0x1b, 0x19, 0x95, 0xbc, 0x00, 0xa3, 0xb9, 0x80,
	0xd0, 0x9b, 0x79, 0xf3, 0x51, 0xfc, 0x22, 0x04, 0x7b, 0x32, 0x54, 0x3c, 0x01, 0x65, 0x42, 0xdf,
	0xe5, 0xb9, 0x7a, 0x7f, 0x1e, 0xba, 0x71, 0xc0, 0x7f, 0xa5, 0xad, 0x8e, 0x71, 0x4f, 0x8f, 0x7e,
	0x93, 0xf1, 0x89, 0x1c, 0x7c, 0x26, 0x83, 0x03, 0x1c, 0xfb, 0x38, 0x6d, 0x19, 0x4c, 0xc8, 0x87,
	0x86, 0xab, 0x1a, 0x42, 0xdf, 0x69, 0xdd, 0xe3, 0x8f, 0xff, 0xcb, 0x5b, 0xc7, 0x77, 0x8f, 0x53,
	0xef, 0x76, 0x73, 0xf1, 0x72, 0xf4, 0x21, 0x3b, 0xbb, 0x9e, 0xf3, 0xf0, 0xac, 0x59, 0x72, 0xa5,
	0x73, 0xbe, 0x4a, 0x86, 0xee, 0x87, 0xfd, 0x7c, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x33, 0x14, 0xf1,
	0xd4, 0x90, 0x02, 0x00, 0x00,
}

func (this *PodBounceDirectiveSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PodBounceDirectiveSpec)
	if !ok {
		that2, ok := that.(PodBounceDirectiveSpec)
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
	if len(this.PodsToBounce) != len(that1.PodsToBounce) {
		return false
	}
	for i := range this.PodsToBounce {
		if !this.PodsToBounce[i].Equal(that1.PodsToBounce[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PodBounceDirectiveSpec_PodSelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PodBounceDirectiveSpec_PodSelector)
	if !ok {
		that2, ok := that.(PodBounceDirectiveSpec_PodSelector)
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
	if this.Namespace != that1.Namespace {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
