// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: github.com/google/cloudprober/probes/udp/proto/config.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProbeConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Port to send UDP Ping to (UDP Echo).  If running with the UDP server that
	// comes with cloudprober, it should be same as
	// ProberConfig.udp_echo_server_port.
	Port *int32 `protobuf:"varint,3,opt,name=port,def=31122" json:"port,omitempty"`
	// Number of sending side ports to use.
	NumTxPorts *int32 `protobuf:"varint,4,opt,name=num_tx_ports,json=numTxPorts,def=16" json:"num_tx_ports,omitempty"`
	// message max to account for MTU.
	MaxLength *int32 `protobuf:"varint,5,opt,name=max_length,json=maxLength,def=1300" json:"max_length,omitempty"`
	// Payload size
	PayloadSize *int32 `protobuf:"varint,6,opt,name=payload_size,json=payloadSize" json:"payload_size,omitempty"`
	// Changes the exported monitoring streams to be per port:
	// 1. Changes the streams names to total-per-port, success-per-port etc.
	// 2. Adds src_port and dst_port as stream labels.
	// Note that the field name is experimental and may change in the future.
	ExportMetricsByPort *bool `protobuf:"varint,7,opt,name=export_metrics_by_port,json=exportMetricsByPort,def=0" json:"export_metrics_by_port,omitempty"`
	// Whether to use all transmit ports per probe, per target.
	// Default is to probe each target once per probe and round-robin through the
	// source ports.
	// Setting this field to true changes the behavior to send traffic from all
	// ports to all targets in each probe.
	// For example, if num_tx_ports is set to 16, in every probe cycle, we'll send
	// 16 packets to every target (1 per tx port).
	// Note that setting this field to true will increase the probe traffic.
	UseAllTxPortsPerProbe *bool `protobuf:"varint,8,opt,name=use_all_tx_ports_per_probe,json=useAllTxPortsPerProbe,def=0" json:"use_all_tx_ports_per_probe,omitempty"`
}

// Default values for ProbeConf fields.
const (
	Default_ProbeConf_Port                  = int32(31122)
	Default_ProbeConf_NumTxPorts            = int32(16)
	Default_ProbeConf_MaxLength             = int32(1300)
	Default_ProbeConf_ExportMetricsByPort   = bool(false)
	Default_ProbeConf_UseAllTxPortsPerProbe = bool(false)
)

func (x *ProbeConf) Reset() {
	*x = ProbeConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_google_cloudprober_probes_udp_proto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeConf) ProtoMessage() {}

func (x *ProbeConf) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_google_cloudprober_probes_udp_proto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeConf.ProtoReflect.Descriptor instead.
func (*ProbeConf) Descriptor() ([]byte, []int) {
	return file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDescGZIP(), []int{0}
}

func (x *ProbeConf) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return Default_ProbeConf_Port
}

func (x *ProbeConf) GetNumTxPorts() int32 {
	if x != nil && x.NumTxPorts != nil {
		return *x.NumTxPorts
	}
	return Default_ProbeConf_NumTxPorts
}

func (x *ProbeConf) GetMaxLength() int32 {
	if x != nil && x.MaxLength != nil {
		return *x.MaxLength
	}
	return Default_ProbeConf_MaxLength
}

func (x *ProbeConf) GetPayloadSize() int32 {
	if x != nil && x.PayloadSize != nil {
		return *x.PayloadSize
	}
	return 0
}

func (x *ProbeConf) GetExportMetricsByPort() bool {
	if x != nil && x.ExportMetricsByPort != nil {
		return *x.ExportMetricsByPort
	}
	return Default_ProbeConf_ExportMetricsByPort
}

func (x *ProbeConf) GetUseAllTxPortsPerProbe() bool {
	if x != nil && x.UseAllTxPortsPerProbe != nil {
		return *x.UseAllTxPortsPerProbe
	}
	return Default_ProbeConf_UseAllTxPortsPerProbe
}

var File_github_com_google_cloudprober_probes_udp_proto_config_proto protoreflect.FileDescriptor

var file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x73, 0x2e, 0x75, 0x64, 0x70, 0x22, 0x92, 0x02, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x12, 0x19, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x3a, 0x05, 0x33, 0x31, 0x31, 0x32, 0x32, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x24,
	0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x78, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x3a, 0x02, 0x31, 0x36, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x54, 0x78, 0x50,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x04, 0x31, 0x33, 0x30, 0x30, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3a, 0x0a, 0x16,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x62,
	0x79, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61,
	0x6c, 0x73, 0x65, 0x52, 0x13, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x42, 0x79, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x40, 0x0a, 0x1a, 0x75, 0x73, 0x65, 0x5f,
	0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x78, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61,
	0x6c, 0x73, 0x65, 0x52, 0x15, 0x75, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x54, 0x78, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x50, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDescOnce sync.Once
	file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDescData = file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDesc
)

func file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDescGZIP() []byte {
	file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDescOnce.Do(func() {
		file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDescData)
	})
	return file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDescData
}

var file_github_com_google_cloudprober_probes_udp_proto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_google_cloudprober_probes_udp_proto_config_proto_goTypes = []interface{}{
	(*ProbeConf)(nil), // 0: cloudprober.probes.udp.ProbeConf
}
var file_github_com_google_cloudprober_probes_udp_proto_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_google_cloudprober_probes_udp_proto_config_proto_init() }
func file_github_com_google_cloudprober_probes_udp_proto_config_proto_init() {
	if File_github_com_google_cloudprober_probes_udp_proto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_google_cloudprober_probes_udp_proto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeConf); i {
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
			RawDescriptor: file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_google_cloudprober_probes_udp_proto_config_proto_goTypes,
		DependencyIndexes: file_github_com_google_cloudprober_probes_udp_proto_config_proto_depIdxs,
		MessageInfos:      file_github_com_google_cloudprober_probes_udp_proto_config_proto_msgTypes,
	}.Build()
	File_github_com_google_cloudprober_probes_udp_proto_config_proto = out.File
	file_github_com_google_cloudprober_probes_udp_proto_config_proto_rawDesc = nil
	file_github_com_google_cloudprober_probes_udp_proto_config_proto_goTypes = nil
	file_github_com_google_cloudprober_probes_udp_proto_config_proto_depIdxs = nil
}
