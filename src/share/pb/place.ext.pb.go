// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: place.ext.proto

package pb

import (
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

type Place struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Id          int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	N           string `protobuf:"bytes,3,opt,name=n,proto3" json:"n,omitempty"`
	PinyinFull  string `protobuf:"bytes,4,opt,name=pinyinFull,proto3" json:"pinyinFull,omitempty"`
	PinyinShort string `protobuf:"bytes,5,opt,name=pinyinShort,proto3" json:"pinyinShort,omitempty"`
}

func (x *Place) Reset() {
	*x = Place{}
	if protoimpl.UnsafeEnabled {
		mi := &file_place_ext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Place) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Place) ProtoMessage() {}

func (x *Place) ProtoReflect() protoreflect.Message {
	mi := &file_place_ext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Place.ProtoReflect.Descriptor instead.
func (*Place) Descriptor() ([]byte, []int) {
	return file_place_ext_proto_rawDescGZIP(), []int{0}
}

func (x *Place) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Place) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Place) GetN() string {
	if x != nil {
		return x.N
	}
	return ""
}

func (x *Place) GetPinyinFull() string {
	if x != nil {
		return x.PinyinFull
	}
	return ""
}

func (x *Place) GetPinyinShort() string {
	if x != nil {
		return x.PinyinShort
	}
	return ""
}

type HotCitiesByCinemaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HotCitiesByCinemaReq) Reset() {
	*x = HotCitiesByCinemaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_place_ext_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotCitiesByCinemaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotCitiesByCinemaReq) ProtoMessage() {}

func (x *HotCitiesByCinemaReq) ProtoReflect() protoreflect.Message {
	mi := &file_place_ext_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotCitiesByCinemaReq.ProtoReflect.Descriptor instead.
func (*HotCitiesByCinemaReq) Descriptor() ([]byte, []int) {
	return file_place_ext_proto_rawDescGZIP(), []int{1}
}

type HotCitiesByCinemaRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P []*Place `protobuf:"bytes,1,rep,name=p,proto3" json:"p,omitempty"`
}

func (x *HotCitiesByCinemaRep) Reset() {
	*x = HotCitiesByCinemaRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_place_ext_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HotCitiesByCinemaRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HotCitiesByCinemaRep) ProtoMessage() {}

func (x *HotCitiesByCinemaRep) ProtoReflect() protoreflect.Message {
	mi := &file_place_ext_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HotCitiesByCinemaRep.ProtoReflect.Descriptor instead.
func (*HotCitiesByCinemaRep) Descriptor() ([]byte, []int) {
	return file_place_ext_proto_rawDescGZIP(), []int{2}
}

func (x *HotCitiesByCinemaRep) GetP() []*Place {
	if x != nil {
		return x.P
	}
	return nil
}

var File_place_ext_proto protoreflect.FileDescriptor

var file_place_ext_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7d, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x46, 0x75, 0x6c, 0x6c, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x69, 0x6e, 0x79, 0x69, 0x6e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x22, 0x16, 0x0a, 0x14, 0x48, 0x6f, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43,
	0x69, 0x6e, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x22, 0x2c, 0x0a, 0x14, 0x48, 0x6f, 0x74, 0x43,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x70,
	0x12, 0x14, 0x0a, 0x01, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x52, 0x01, 0x70, 0x32, 0x54, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x78, 0x74, 0x12, 0x41, 0x0a, 0x11, 0x48, 0x6f, 0x74,
	0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x12, 0x15,
	0x2e, 0x48, 0x6f, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x69, 0x6e, 0x65,
	0x6d, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x48, 0x6f, 0x74, 0x43, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x42, 0x79, 0x43, 0x69, 0x6e, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x70, 0x42, 0x14, 0x5a, 0x12,
	0x2e, 0x2e, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2f, 0x70, 0x62, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_place_ext_proto_rawDescOnce sync.Once
	file_place_ext_proto_rawDescData = file_place_ext_proto_rawDesc
)

func file_place_ext_proto_rawDescGZIP() []byte {
	file_place_ext_proto_rawDescOnce.Do(func() {
		file_place_ext_proto_rawDescData = protoimpl.X.CompressGZIP(file_place_ext_proto_rawDescData)
	})
	return file_place_ext_proto_rawDescData
}

var file_place_ext_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_place_ext_proto_goTypes = []interface{}{
	(*Place)(nil),                // 0: Place
	(*HotCitiesByCinemaReq)(nil), // 1: HotCitiesByCinemaReq
	(*HotCitiesByCinemaRep)(nil), // 2: HotCitiesByCinemaRep
}
var file_place_ext_proto_depIdxs = []int32{
	0, // 0: HotCitiesByCinemaRep.p:type_name -> Place
	1, // 1: PlaceServiceExt.HotCitiesByCinema:input_type -> HotCitiesByCinemaReq
	2, // 2: PlaceServiceExt.HotCitiesByCinema:output_type -> HotCitiesByCinemaRep
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_place_ext_proto_init() }
func file_place_ext_proto_init() {
	if File_place_ext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_place_ext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Place); i {
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
		file_place_ext_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotCitiesByCinemaReq); i {
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
		file_place_ext_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HotCitiesByCinemaRep); i {
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
			RawDescriptor: file_place_ext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_place_ext_proto_goTypes,
		DependencyIndexes: file_place_ext_proto_depIdxs,
		MessageInfos:      file_place_ext_proto_msgTypes,
	}.Build()
	File_place_ext_proto = out.File
	file_place_ext_proto_rawDesc = nil
	file_place_ext_proto_goTypes = nil
	file_place_ext_proto_depIdxs = nil
}