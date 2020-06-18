// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/networking/v1alpha1/traffic_policy.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for TrafficPolicySpec
func (this *TrafficPolicySpec) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec
func (this *TrafficPolicySpec) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_RetryPolicy
func (this *TrafficPolicySpec_RetryPolicy) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_RetryPolicy
func (this *TrafficPolicySpec_RetryPolicy) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_MultiDestination
func (this *TrafficPolicySpec_MultiDestination) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_MultiDestination
func (this *TrafficPolicySpec_MultiDestination) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_MultiDestination_WeightedDestination
func (this *TrafficPolicySpec_MultiDestination_WeightedDestination) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_MultiDestination_WeightedDestination
func (this *TrafficPolicySpec_MultiDestination_WeightedDestination) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_FaultInjection
func (this *TrafficPolicySpec_FaultInjection) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_FaultInjection
func (this *TrafficPolicySpec_FaultInjection) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_FaultInjection_Delay
func (this *TrafficPolicySpec_FaultInjection_Delay) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_FaultInjection_Delay
func (this *TrafficPolicySpec_FaultInjection_Delay) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_FaultInjection_Abort
func (this *TrafficPolicySpec_FaultInjection_Abort) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_FaultInjection_Abort
func (this *TrafficPolicySpec_FaultInjection_Abort) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_HeaderManipulation
func (this *TrafficPolicySpec_HeaderManipulation) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_HeaderManipulation
func (this *TrafficPolicySpec_HeaderManipulation) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_CorsPolicy
func (this *TrafficPolicySpec_CorsPolicy) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_CorsPolicy
func (this *TrafficPolicySpec_CorsPolicy) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_HttpMatcher
func (this *TrafficPolicySpec_HttpMatcher) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_HttpMatcher
func (this *TrafficPolicySpec_HttpMatcher) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_StringMatch
func (this *TrafficPolicySpec_StringMatch) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_StringMatch
func (this *TrafficPolicySpec_StringMatch) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_HeaderMatcher
func (this *TrafficPolicySpec_HeaderMatcher) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_HeaderMatcher
func (this *TrafficPolicySpec_HeaderMatcher) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_QueryParameterMatcher
func (this *TrafficPolicySpec_QueryParameterMatcher) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_QueryParameterMatcher
func (this *TrafficPolicySpec_QueryParameterMatcher) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_Mirror
func (this *TrafficPolicySpec_Mirror) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_Mirror
func (this *TrafficPolicySpec_Mirror) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicySpec_HttpMethod
func (this *TrafficPolicySpec_HttpMethod) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicySpec_HttpMethod
func (this *TrafficPolicySpec_HttpMethod) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicyStatus
func (this *TrafficPolicyStatus) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicyStatus
func (this *TrafficPolicyStatus) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicyStatus_TranslatorError
func (this *TrafficPolicyStatus_TranslatorError) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicyStatus_TranslatorError
func (this *TrafficPolicyStatus_TranslatorError) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicyStatus_ConflictError
func (this *TrafficPolicyStatus_ConflictError) MarshalJSON() ([]byte, error) {
	str, err := TrafficPolicyMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicyStatus_ConflictError
func (this *TrafficPolicyStatus_ConflictError) UnmarshalJSON(b []byte) error {
	return TrafficPolicyUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	TrafficPolicyMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	TrafficPolicyUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
