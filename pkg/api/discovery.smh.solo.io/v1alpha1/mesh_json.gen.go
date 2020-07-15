// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/discovery/v1alpha1/mesh.proto

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	_ "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for MeshSpec
func (this *MeshSpec) MarshalJSON() ([]byte, error) {
	str, err := MeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshSpec
func (this *MeshSpec) UnmarshalJSON(b []byte) error {
	return MeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshSpec_AppliedVirtualMesh
func (this *MeshSpec_AppliedVirtualMesh) MarshalJSON() ([]byte, error) {
	str, err := MeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshSpec_AppliedVirtualMesh
func (this *MeshSpec_AppliedVirtualMesh) UnmarshalJSON(b []byte) error {
	return MeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshSpec_Istio
func (this *MeshSpec_Istio) MarshalJSON() ([]byte, error) {
	str, err := MeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshSpec_Istio
func (this *MeshSpec_Istio) UnmarshalJSON(b []byte) error {
	return MeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshSpec_Istio_CitadelInfo
func (this *MeshSpec_Istio_CitadelInfo) MarshalJSON() ([]byte, error) {
	str, err := MeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshSpec_Istio_CitadelInfo
func (this *MeshSpec_Istio_CitadelInfo) UnmarshalJSON(b []byte) error {
	return MeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshSpec_AwsAppMesh
func (this *MeshSpec_AwsAppMesh) MarshalJSON() ([]byte, error) {
	str, err := MeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshSpec_AwsAppMesh
func (this *MeshSpec_AwsAppMesh) UnmarshalJSON(b []byte) error {
	return MeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshSpec_LinkerdMesh
func (this *MeshSpec_LinkerdMesh) MarshalJSON() ([]byte, error) {
	str, err := MeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshSpec_LinkerdMesh
func (this *MeshSpec_LinkerdMesh) UnmarshalJSON(b []byte) error {
	return MeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshSpec_ConsulConnectMesh
func (this *MeshSpec_ConsulConnectMesh) MarshalJSON() ([]byte, error) {
	str, err := MeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshSpec_ConsulConnectMesh
func (this *MeshSpec_ConsulConnectMesh) UnmarshalJSON(b []byte) error {
	return MeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshSpec_MeshInstallation
func (this *MeshSpec_MeshInstallation) MarshalJSON() ([]byte, error) {
	str, err := MeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshSpec_MeshInstallation
func (this *MeshSpec_MeshInstallation) UnmarshalJSON(b []byte) error {
	return MeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for MeshStatus
func (this *MeshStatus) MarshalJSON() ([]byte, error) {
	str, err := MeshMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for MeshStatus
func (this *MeshStatus) UnmarshalJSON(b []byte) error {
	return MeshUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	MeshMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	MeshUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
