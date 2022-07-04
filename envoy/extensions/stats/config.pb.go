// Copyright 2019 Istio Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: envoy/extensions/stats/config.proto

// $title: Stats Config
// $description: Configuration for Stats Filter.
// $location: https://istio.io/docs/reference/config/proxy_extensions/stats.html
// $weight: 20

package stats

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MetricType int32

const (
	MetricType_COUNTER   MetricType = 0
	MetricType_GAUGE     MetricType = 1
	MetricType_HISTOGRAM MetricType = 2
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0: "COUNTER",
		1: "GAUGE",
		2: "HISTOGRAM",
	}
	MetricType_value = map[string]int32{
		"COUNTER":   0,
		"GAUGE":     1,
		"HISTOGRAM": 2,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_stats_config_proto_enumTypes[0].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_stats_config_proto_enumTypes[0]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_stats_config_proto_rawDescGZIP(), []int{0}
}

// Metric instance configuration overrides.
// The metric value and the metric type are optional and permit changing the
// reported value for an existing metric.
// The standard metrics are optimized and reported through a "fast-path".
// The customizations allow full configurability, at the cost of a "slower"
// path.
type MetricConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Optional) Collection of tag names and tag expressions to include in the
	// metric. Conflicts are resolved by the tag name by overriding previously
	// supplied values.
	Dimensions map[string]string `protobuf:"bytes,1,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// (Optional) Metric name to restrict the override to a metric. If not
	// specified, applies to all.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// (Optional) A list of tags to remove.
	TagsToRemove []string `protobuf:"bytes,3,rep,name=tags_to_remove,json=tagsToRemove,proto3" json:"tags_to_remove,omitempty"`
	// NOT IMPLEMENTED. (Optional) Conditional enabling the override.
	Match string `protobuf:"bytes,4,opt,name=match,proto3" json:"match,omitempty"`
	// (Optional) If this is set to true, the metric(s) selected by this
	// configuration will not be generated or reported.
	Drop bool `protobuf:"varint,5,opt,name=drop,proto3" json:"drop,omitempty"`
}

func (x *MetricConfig) Reset() {
	*x = MetricConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_stats_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricConfig) ProtoMessage() {}

func (x *MetricConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_stats_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricConfig.ProtoReflect.Descriptor instead.
func (*MetricConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_stats_config_proto_rawDescGZIP(), []int{0}
}

func (x *MetricConfig) GetDimensions() map[string]string {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

func (x *MetricConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricConfig) GetTagsToRemove() []string {
	if x != nil {
		return x.TagsToRemove
	}
	return nil
}

func (x *MetricConfig) GetMatch() string {
	if x != nil {
		return x.Match
	}
	return ""
}

func (x *MetricConfig) GetDrop() bool {
	if x != nil {
		return x.Drop
	}
	return false
}

type MetricDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metric name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Metric value expression.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// NOT IMPLEMENTED (Optional) Metric type.
	Type MetricType `protobuf:"varint,3,opt,name=type,proto3,enum=stats.MetricType" json:"type,omitempty"`
}

func (x *MetricDefinition) Reset() {
	*x = MetricDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_stats_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricDefinition) ProtoMessage() {}

func (x *MetricDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_stats_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricDefinition.ProtoReflect.Descriptor instead.
func (*MetricDefinition) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_stats_config_proto_rawDescGZIP(), []int{1}
}

func (x *MetricDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MetricDefinition) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *MetricDefinition) GetType() MetricType {
	if x != nil {
		return x.Type
	}
	return MetricType_COUNTER
}

type PluginConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// next id: 7
	// The following settings should be rarely used.
	// Enable debug for this filter.
	// DEPRECATED.
	Debug bool `protobuf:"varint,1,opt,name=debug,proto3" json:"debug,omitempty"`
	// maximum size of the peer metadata cache.
	// A long lived proxy that connects with many transient peers can build up a
	// large cache. To turn off the cache, set this field to a negative value.
	// DEPRECATED.
	MaxPeerCacheSize int32 `protobuf:"varint,2,opt,name=max_peer_cache_size,json=maxPeerCacheSize,proto3" json:"max_peer_cache_size,omitempty"`
	// prefix to add to stats emitted by the plugin.
	// DEPRECATED.
	StatPrefix string `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"` // default: "istio_"
	// Stats api squashes dimensions in a single string.
	// The squashed string is parsed at prometheus scrape time to recover
	// dimensions. The following 2 fields set the field and value separators {key:
	// value} -->  key{value_separator}value{field_separator}
	FieldSeparator string `protobuf:"bytes,4,opt,name=field_separator,json=fieldSeparator,proto3" json:"field_separator,omitempty"` // default: ";;"
	ValueSeparator string `protobuf:"bytes,5,opt,name=value_separator,json=valueSeparator,proto3" json:"value_separator,omitempty"` // default: "=="
	// Optional: Disable using host header as a fallback if destination service is
	// not available from the controlplane. Disable the fallback if the host
	// header originates outsides the mesh, like at ingress.
	DisableHostHeaderFallback bool `protobuf:"varint,6,opt,name=disable_host_header_fallback,json=disableHostHeaderFallback,proto3" json:"disable_host_header_fallback,omitempty"`
	// Optional. Allows configuration of the time between calls out to for TCP
	// metrics reporting. The default duration is `15s`.
	TcpReportingDuration *duration.Duration `protobuf:"bytes,7,opt,name=tcp_reporting_duration,json=tcpReportingDuration,proto3" json:"tcp_reporting_duration,omitempty"`
	// Metric overrides.
	Metrics []*MetricConfig `protobuf:"bytes,8,rep,name=metrics,proto3" json:"metrics,omitempty"`
	// Metric definitions.
	Definitions []*MetricDefinition `protobuf:"bytes,9,rep,name=definitions,proto3" json:"definitions,omitempty"`
}

func (x *PluginConfig) Reset() {
	*x = PluginConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_stats_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginConfig) ProtoMessage() {}

func (x *PluginConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_stats_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginConfig.ProtoReflect.Descriptor instead.
func (*PluginConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_stats_config_proto_rawDescGZIP(), []int{2}
}

func (x *PluginConfig) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *PluginConfig) GetMaxPeerCacheSize() int32 {
	if x != nil {
		return x.MaxPeerCacheSize
	}
	return 0
}

func (x *PluginConfig) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *PluginConfig) GetFieldSeparator() string {
	if x != nil {
		return x.FieldSeparator
	}
	return ""
}

func (x *PluginConfig) GetValueSeparator() string {
	if x != nil {
		return x.ValueSeparator
	}
	return ""
}

func (x *PluginConfig) GetDisableHostHeaderFallback() bool {
	if x != nil {
		return x.DisableHostHeaderFallback
	}
	return false
}

func (x *PluginConfig) GetTcpReportingDuration() *duration.Duration {
	if x != nil {
		return x.TcpReportingDuration
	}
	return nil
}

func (x *PluginConfig) GetMetrics() []*MetricConfig {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *PluginConfig) GetDefinitions() []*MetricDefinition {
	if x != nil {
		return x.Definitions
	}
	return nil
}

var File_envoy_extensions_stats_config_proto protoreflect.FileDescriptor

var file_envoy_extensions_stats_config_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a,
	0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x43, 0x0a,
	0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x67, 0x73, 0x5f, 0x74,
	0x6f, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x61, 0x67, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x72, 0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x64, 0x72, 0x6f, 0x70, 0x1a, 0x3d, 0x0a, 0x0f, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x63, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xc2, 0x03, 0x0a, 0x0c, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x12, 0x2d, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x6d, 0x61, 0x78, 0x50, 0x65, 0x65, 0x72, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x70, 0x61, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x53, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x70, 0x61, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x1c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x46, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x4f, 0x0a, 0x16, 0x74, 0x63, 0x70, 0x5f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x14, 0x74, 0x63, 0x70, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a,
	0x33, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x41,
	0x55, 0x47, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52,
	0x41, 0x4d, 0x10, 0x02, 0x42, 0x25, 0x5a, 0x23, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_stats_config_proto_rawDescOnce sync.Once
	file_envoy_extensions_stats_config_proto_rawDescData = file_envoy_extensions_stats_config_proto_rawDesc
)

func file_envoy_extensions_stats_config_proto_rawDescGZIP() []byte {
	file_envoy_extensions_stats_config_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_stats_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_stats_config_proto_rawDescData)
	})
	return file_envoy_extensions_stats_config_proto_rawDescData
}

var file_envoy_extensions_stats_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_stats_config_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_stats_config_proto_goTypes = []interface{}{
	(MetricType)(0),           // 0: stats.MetricType
	(*MetricConfig)(nil),      // 1: stats.MetricConfig
	(*MetricDefinition)(nil),  // 2: stats.MetricDefinition
	(*PluginConfig)(nil),      // 3: stats.PluginConfig
	nil,                       // 4: stats.MetricConfig.DimensionsEntry
	(*duration.Duration)(nil), // 5: google.protobuf.Duration
}
var file_envoy_extensions_stats_config_proto_depIdxs = []int32{
	4, // 0: stats.MetricConfig.dimensions:type_name -> stats.MetricConfig.DimensionsEntry
	0, // 1: stats.MetricDefinition.type:type_name -> stats.MetricType
	5, // 2: stats.PluginConfig.tcp_reporting_duration:type_name -> google.protobuf.Duration
	1, // 3: stats.PluginConfig.metrics:type_name -> stats.MetricConfig
	2, // 4: stats.PluginConfig.definitions:type_name -> stats.MetricDefinition
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_stats_config_proto_init() }
func file_envoy_extensions_stats_config_proto_init() {
	if File_envoy_extensions_stats_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_stats_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_stats_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricDefinition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_stats_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_stats_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_stats_config_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_stats_config_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_stats_config_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_stats_config_proto_msgTypes,
	}.Build()
	File_envoy_extensions_stats_config_proto = out.File
	file_envoy_extensions_stats_config_proto_rawDesc = nil
	file_envoy_extensions_stats_config_proto_goTypes = nil
	file_envoy_extensions_stats_config_proto_depIdxs = nil
}
