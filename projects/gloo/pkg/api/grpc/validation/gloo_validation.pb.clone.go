// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/grpc/validation/gloo_validation.proto

package validation

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
func (m *GlooValidationServiceRequest) Clone() proto.Message {
	var target *GlooValidationServiceRequest
	if m == nil {
		return target
	}
	target = &GlooValidationServiceRequest{}

	if h, ok := interface{}(m.GetProxy()).(clone.Cloner); ok {
		target.Proxy = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.Proxy)
	} else {
		target.Proxy = proto.Clone(m.GetProxy()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.Proxy)
	}

	switch m.Resources.(type) {

	case *GlooValidationServiceRequest_ModifiedResources:

		if h, ok := interface{}(m.GetModifiedResources()).(clone.Cloner); ok {
			target.Resources = &GlooValidationServiceRequest_ModifiedResources{
				ModifiedResources: h.Clone().(*ModifiedResources),
			}
		} else {
			target.Resources = &GlooValidationServiceRequest_ModifiedResources{
				ModifiedResources: proto.Clone(m.GetModifiedResources()).(*ModifiedResources),
			}
		}

	case *GlooValidationServiceRequest_DeletedResources:

		if h, ok := interface{}(m.GetDeletedResources()).(clone.Cloner); ok {
			target.Resources = &GlooValidationServiceRequest_DeletedResources{
				DeletedResources: h.Clone().(*DeletedResources),
			}
		} else {
			target.Resources = &GlooValidationServiceRequest_DeletedResources{
				DeletedResources: proto.Clone(m.GetDeletedResources()).(*DeletedResources),
			}
		}

	}

	return target
}

// Clone function
func (m *GlooValidationServiceResponse) Clone() proto.Message {
	var target *GlooValidationServiceResponse
	if m == nil {
		return target
	}
	target = &GlooValidationServiceResponse{}

	if m.GetValidationReports() != nil {
		target.ValidationReports = make([]*ValidationReport, len(m.GetValidationReports()))
		for idx, v := range m.GetValidationReports() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ValidationReports[idx] = h.Clone().(*ValidationReport)
			} else {
				target.ValidationReports[idx] = proto.Clone(v).(*ValidationReport)
			}

		}
	}

	return target
}

// Clone function
func (m *ModifiedResources) Clone() proto.Message {
	var target *ModifiedResources
	if m == nil {
		return target
	}
	target = &ModifiedResources{}

	if m.GetUpstreams() != nil {
		target.Upstreams = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.Upstream, len(m.GetUpstreams()))
		for idx, v := range m.GetUpstreams() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Upstreams[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.Upstream)
			} else {
				target.Upstreams[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.Upstream)
			}

		}
	}

	return target
}

// Clone function
func (m *DeletedResources) Clone() proto.Message {
	var target *DeletedResources
	if m == nil {
		return target
	}
	target = &DeletedResources{}

	if m.GetUpstreamRefs() != nil {
		target.UpstreamRefs = make([]*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef, len(m.GetUpstreamRefs()))
		for idx, v := range m.GetUpstreamRefs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.UpstreamRefs[idx] = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			} else {
				target.UpstreamRefs[idx] = proto.Clone(v).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
			}

		}
	}

	return target
}

// Clone function
func (m *ValidationReport) Clone() proto.Message {
	var target *ValidationReport
	if m == nil {
		return target
	}
	target = &ValidationReport{}

	if h, ok := interface{}(m.GetProxyReport()).(clone.Cloner); ok {
		target.ProxyReport = h.Clone().(*ProxyReport)
	} else {
		target.ProxyReport = proto.Clone(m.GetProxyReport()).(*ProxyReport)
	}

	if m.GetUpstreamReports() != nil {
		target.UpstreamReports = make([]*ResourceReport, len(m.GetUpstreamReports()))
		for idx, v := range m.GetUpstreamReports() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.UpstreamReports[idx] = h.Clone().(*ResourceReport)
			} else {
				target.UpstreamReports[idx] = proto.Clone(v).(*ResourceReport)
			}

		}
	}

	if h, ok := interface{}(m.GetProxy()).(clone.Cloner); ok {
		target.Proxy = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.Proxy)
	} else {
		target.Proxy = proto.Clone(m.GetProxy()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1.Proxy)
	}

	return target
}

// Clone function
func (m *ResourceReport) Clone() proto.Message {
	var target *ResourceReport
	if m == nil {
		return target
	}
	target = &ResourceReport{}

	if h, ok := interface{}(m.GetResourceRef()).(clone.Cloner); ok {
		target.ResourceRef = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	} else {
		target.ResourceRef = proto.Clone(m.GetResourceRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef)
	}

	if m.GetWarnings() != nil {
		target.Warnings = make([]string, len(m.GetWarnings()))
		for idx, v := range m.GetWarnings() {

			target.Warnings[idx] = v

		}
	}

	if m.GetErrors() != nil {
		target.Errors = make([]string, len(m.GetErrors()))
		for idx, v := range m.GetErrors() {

			target.Errors[idx] = v

		}
	}

	return target
}

// Clone function
func (m *NotifyOnResyncRequest) Clone() proto.Message {
	var target *NotifyOnResyncRequest
	if m == nil {
		return target
	}
	target = &NotifyOnResyncRequest{}

	return target
}

// Clone function
func (m *NotifyOnResyncResponse) Clone() proto.Message {
	var target *NotifyOnResyncResponse
	if m == nil {
		return target
	}
	target = &NotifyOnResyncResponse{}

	return target
}

// Clone function
func (m *ProxyReport) Clone() proto.Message {
	var target *ProxyReport
	if m == nil {
		return target
	}
	target = &ProxyReport{}

	if m.GetListenerReports() != nil {
		target.ListenerReports = make([]*ListenerReport, len(m.GetListenerReports()))
		for idx, v := range m.GetListenerReports() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.ListenerReports[idx] = h.Clone().(*ListenerReport)
			} else {
				target.ListenerReports[idx] = proto.Clone(v).(*ListenerReport)
			}

		}
	}

	return target
}

// Clone function
func (m *ListenerReport) Clone() proto.Message {
	var target *ListenerReport
	if m == nil {
		return target
	}
	target = &ListenerReport{}

	if m.GetErrors() != nil {
		target.Errors = make([]*ListenerReport_Error, len(m.GetErrors()))
		for idx, v := range m.GetErrors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Errors[idx] = h.Clone().(*ListenerReport_Error)
			} else {
				target.Errors[idx] = proto.Clone(v).(*ListenerReport_Error)
			}

		}
	}

	switch m.ListenerTypeReport.(type) {

	case *ListenerReport_HttpListenerReport:

		if h, ok := interface{}(m.GetHttpListenerReport()).(clone.Cloner); ok {
			target.ListenerTypeReport = &ListenerReport_HttpListenerReport{
				HttpListenerReport: h.Clone().(*HttpListenerReport),
			}
		} else {
			target.ListenerTypeReport = &ListenerReport_HttpListenerReport{
				HttpListenerReport: proto.Clone(m.GetHttpListenerReport()).(*HttpListenerReport),
			}
		}

	case *ListenerReport_TcpListenerReport:

		if h, ok := interface{}(m.GetTcpListenerReport()).(clone.Cloner); ok {
			target.ListenerTypeReport = &ListenerReport_TcpListenerReport{
				TcpListenerReport: h.Clone().(*TcpListenerReport),
			}
		} else {
			target.ListenerTypeReport = &ListenerReport_TcpListenerReport{
				TcpListenerReport: proto.Clone(m.GetTcpListenerReport()).(*TcpListenerReport),
			}
		}

	}

	return target
}

// Clone function
func (m *HttpListenerReport) Clone() proto.Message {
	var target *HttpListenerReport
	if m == nil {
		return target
	}
	target = &HttpListenerReport{}

	if m.GetErrors() != nil {
		target.Errors = make([]*HttpListenerReport_Error, len(m.GetErrors()))
		for idx, v := range m.GetErrors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Errors[idx] = h.Clone().(*HttpListenerReport_Error)
			} else {
				target.Errors[idx] = proto.Clone(v).(*HttpListenerReport_Error)
			}

		}
	}

	if m.GetVirtualHostReports() != nil {
		target.VirtualHostReports = make([]*VirtualHostReport, len(m.GetVirtualHostReports()))
		for idx, v := range m.GetVirtualHostReports() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.VirtualHostReports[idx] = h.Clone().(*VirtualHostReport)
			} else {
				target.VirtualHostReports[idx] = proto.Clone(v).(*VirtualHostReport)
			}

		}
	}

	return target
}

// Clone function
func (m *VirtualHostReport) Clone() proto.Message {
	var target *VirtualHostReport
	if m == nil {
		return target
	}
	target = &VirtualHostReport{}

	if m.GetErrors() != nil {
		target.Errors = make([]*VirtualHostReport_Error, len(m.GetErrors()))
		for idx, v := range m.GetErrors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Errors[idx] = h.Clone().(*VirtualHostReport_Error)
			} else {
				target.Errors[idx] = proto.Clone(v).(*VirtualHostReport_Error)
			}

		}
	}

	if m.GetRouteReports() != nil {
		target.RouteReports = make([]*RouteReport, len(m.GetRouteReports()))
		for idx, v := range m.GetRouteReports() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.RouteReports[idx] = h.Clone().(*RouteReport)
			} else {
				target.RouteReports[idx] = proto.Clone(v).(*RouteReport)
			}

		}
	}

	return target
}

// Clone function
func (m *RouteReport) Clone() proto.Message {
	var target *RouteReport
	if m == nil {
		return target
	}
	target = &RouteReport{}

	if m.GetErrors() != nil {
		target.Errors = make([]*RouteReport_Error, len(m.GetErrors()))
		for idx, v := range m.GetErrors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Errors[idx] = h.Clone().(*RouteReport_Error)
			} else {
				target.Errors[idx] = proto.Clone(v).(*RouteReport_Error)
			}

		}
	}

	if m.GetWarnings() != nil {
		target.Warnings = make([]*RouteReport_Warning, len(m.GetWarnings()))
		for idx, v := range m.GetWarnings() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Warnings[idx] = h.Clone().(*RouteReport_Warning)
			} else {
				target.Warnings[idx] = proto.Clone(v).(*RouteReport_Warning)
			}

		}
	}

	return target
}

// Clone function
func (m *TcpListenerReport) Clone() proto.Message {
	var target *TcpListenerReport
	if m == nil {
		return target
	}
	target = &TcpListenerReport{}

	if m.GetErrors() != nil {
		target.Errors = make([]*TcpListenerReport_Error, len(m.GetErrors()))
		for idx, v := range m.GetErrors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Errors[idx] = h.Clone().(*TcpListenerReport_Error)
			} else {
				target.Errors[idx] = proto.Clone(v).(*TcpListenerReport_Error)
			}

		}
	}

	if m.GetTcpHostReports() != nil {
		target.TcpHostReports = make([]*TcpHostReport, len(m.GetTcpHostReports()))
		for idx, v := range m.GetTcpHostReports() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.TcpHostReports[idx] = h.Clone().(*TcpHostReport)
			} else {
				target.TcpHostReports[idx] = proto.Clone(v).(*TcpHostReport)
			}

		}
	}

	return target
}

// Clone function
func (m *TcpHostReport) Clone() proto.Message {
	var target *TcpHostReport
	if m == nil {
		return target
	}
	target = &TcpHostReport{}

	if m.GetErrors() != nil {
		target.Errors = make([]*TcpHostReport_Error, len(m.GetErrors()))
		for idx, v := range m.GetErrors() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Errors[idx] = h.Clone().(*TcpHostReport_Error)
			} else {
				target.Errors[idx] = proto.Clone(v).(*TcpHostReport_Error)
			}

		}
	}

	return target
}

// Clone function
func (m *ListenerReport_Error) Clone() proto.Message {
	var target *ListenerReport_Error
	if m == nil {
		return target
	}
	target = &ListenerReport_Error{}

	target.Type = m.GetType()

	target.Reason = m.GetReason()

	return target
}

// Clone function
func (m *HttpListenerReport_Error) Clone() proto.Message {
	var target *HttpListenerReport_Error
	if m == nil {
		return target
	}
	target = &HttpListenerReport_Error{}

	target.Type = m.GetType()

	target.Reason = m.GetReason()

	return target
}

// Clone function
func (m *VirtualHostReport_Error) Clone() proto.Message {
	var target *VirtualHostReport_Error
	if m == nil {
		return target
	}
	target = &VirtualHostReport_Error{}

	target.Type = m.GetType()

	target.Reason = m.GetReason()

	return target
}

// Clone function
func (m *RouteReport_Error) Clone() proto.Message {
	var target *RouteReport_Error
	if m == nil {
		return target
	}
	target = &RouteReport_Error{}

	target.Type = m.GetType()

	target.Reason = m.GetReason()

	return target
}

// Clone function
func (m *RouteReport_Warning) Clone() proto.Message {
	var target *RouteReport_Warning
	if m == nil {
		return target
	}
	target = &RouteReport_Warning{}

	target.Type = m.GetType()

	target.Reason = m.GetReason()

	return target
}

// Clone function
func (m *TcpListenerReport_Error) Clone() proto.Message {
	var target *TcpListenerReport_Error
	if m == nil {
		return target
	}
	target = &TcpListenerReport_Error{}

	target.Type = m.GetType()

	target.Reason = m.GetReason()

	return target
}

// Clone function
func (m *TcpHostReport_Error) Clone() proto.Message {
	var target *TcpHostReport_Error
	if m == nil {
		return target
	}
	target = &TcpHostReport_Error{}

	target.Type = m.GetType()

	target.Reason = m.GetReason()

	return target
}
