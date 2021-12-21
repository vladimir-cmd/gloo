// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/circuit_breaker.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_golang_protobuf_ptypes_wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *CircuitBreakerConfig) Clone() proto.Message {
	var target *CircuitBreakerConfig
	if m == nil {
		return target
	}
	target = &CircuitBreakerConfig{}

	if h, ok := interface{}(m.GetMaxConnections()).(clone.Cloner); ok {
		target.MaxConnections = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxConnections = proto.Clone(m.GetMaxConnections()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetMaxPendingRequests()).(clone.Cloner); ok {
		target.MaxPendingRequests = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxPendingRequests = proto.Clone(m.GetMaxPendingRequests()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetMaxRequests()).(clone.Cloner); ok {
		target.MaxRequests = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxRequests = proto.Clone(m.GetMaxRequests()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	if h, ok := interface{}(m.GetMaxRetries()).(clone.Cloner); ok {
		target.MaxRetries = h.Clone().(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	} else {
		target.MaxRetries = proto.Clone(m.GetMaxRetries()).(*github_com_golang_protobuf_ptypes_wrappers.UInt32Value)
	}

	return target
}
