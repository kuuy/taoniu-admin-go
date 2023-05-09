// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: live.proto

package live

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LiveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    string               `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Open      float32              `protobuf:"fixed32,2,opt,name=open,proto3" json:"open,omitempty"`
	Price     float32              `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	High      float32              `protobuf:"fixed32,4,opt,name=high,proto3" json:"high,omitempty"`
	Low       float32              `protobuf:"fixed32,5,opt,name=low,proto3" json:"low,omitempty"`
	Volume    float64              `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
	Quota     float64              `protobuf:"fixed64,7,opt,name=quota,proto3" json:"quota,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *LiveInfo) Reset() {
	*x = LiveInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveInfo) ProtoMessage() {}

func (x *LiveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_live_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveInfo.ProtoReflect.Descriptor instead.
func (*LiveInfo) Descriptor() ([]byte, []int) {
	return file_live_proto_rawDescGZIP(), []int{0}
}

func (x *LiveInfo) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *LiveInfo) GetOpen() float32 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *LiveInfo) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *LiveInfo) GetHigh() float32 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *LiveInfo) GetLow() float32 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *LiveInfo) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *LiveInfo) GetQuota() float64 {
	if x != nil {
		return x.Quota
	}
	return 0
}

func (x *LiveInfo) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type PagenateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol   string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PagenateRequest) Reset() {
	*x = PagenateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagenateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagenateRequest) ProtoMessage() {}

func (x *PagenateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_live_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagenateRequest.ProtoReflect.Descriptor instead.
func (*PagenateRequest) Descriptor() ([]byte, []int) {
	return file_live_proto_rawDescGZIP(), []int{1}
}

func (x *PagenateRequest) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *PagenateRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PagenateRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PagenateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*LiveInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *PagenateReply) Reset() {
	*x = PagenateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagenateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagenateReply) ProtoMessage() {}

func (x *PagenateReply) ProtoReflect() protoreflect.Message {
	mi := &file_live_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagenateReply.ProtoReflect.Descriptor instead.
func (*PagenateReply) Descriptor() ([]byte, []int) {
	return file_live_proto_rawDescGZIP(), []int{2}
}

func (x *PagenateReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PagenateReply) GetData() []*LiveInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_live_proto protoreflect.FileDescriptor

var file_live_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x74, 0x61,
	0x6f, 0x6e, 0x69, 0x75, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x73, 0x70, 0x6f, 0x74, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x6c, 0x69, 0x76,
	0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xda, 0x01, 0x0a, 0x08, 0x4c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x71, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x59, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x65, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x78, 0x0a, 0x0d, 0x50, 0x61,
	0x67, 0x65, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x51, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x74, 0x61, 0x6f, 0x6e, 0x69, 0x75, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0x9f, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x96, 0x01,
	0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x44, 0x2e, 0x74, 0x61, 0x6f,
	0x6e, 0x69, 0x75, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73,
	0x70, 0x6f, 0x74, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x42, 0x2e, 0x74, 0x61, 0x6f, 0x6e, 0x69, 0x75, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x73, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x6e, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x74, 0x61, 0x6f, 0x6e, 0x69, 0x75,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x70, 0x6f, 0x74,
	0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_live_proto_rawDescOnce sync.Once
	file_live_proto_rawDescData = file_live_proto_rawDesc
)

func file_live_proto_rawDescGZIP() []byte {
	file_live_proto_rawDescOnce.Do(func() {
		file_live_proto_rawDescData = protoimpl.X.CompressGZIP(file_live_proto_rawDescData)
	})
	return file_live_proto_rawDescData
}

var file_live_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_live_proto_goTypes = []interface{}{
	(*LiveInfo)(nil),            // 0: taoniu.local.cryptos.grpc.binance.spot.markets.live.LiveInfo
	(*PagenateRequest)(nil),     // 1: taoniu.local.cryptos.grpc.binance.spot.markets.live.PagenateRequest
	(*PagenateReply)(nil),       // 2: taoniu.local.cryptos.grpc.binance.spot.markets.live.PagenateReply
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_live_proto_depIdxs = []int32{
	3, // 0: taoniu.local.cryptos.grpc.binance.spot.markets.live.LiveInfo.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: taoniu.local.cryptos.grpc.binance.spot.markets.live.PagenateReply.data:type_name -> taoniu.local.cryptos.grpc.binance.spot.markets.live.LiveInfo
	1, // 2: taoniu.local.cryptos.grpc.binance.spot.markets.live.Live.Pagenate:input_type -> taoniu.local.cryptos.grpc.binance.spot.markets.live.PagenateRequest
	2, // 3: taoniu.local.cryptos.grpc.binance.spot.markets.live.Live.Pagenate:output_type -> taoniu.local.cryptos.grpc.binance.spot.markets.live.PagenateReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_live_proto_init() }
func file_live_proto_init() {
	if File_live_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_live_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveInfo); i {
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
		file_live_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagenateRequest); i {
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
		file_live_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagenateReply); i {
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
			RawDescriptor: file_live_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_live_proto_goTypes,
		DependencyIndexes: file_live_proto_depIdxs,
		MessageInfos:      file_live_proto_msgTypes,
	}.Build()
	File_live_proto = out.File
	file_live_proto_rawDesc = nil
	file_live_proto_goTypes = nil
	file_live_proto_depIdxs = nil
}
