// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: src/calculator/v1/service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ItemType int32

const (
	ItemType_APPLE     ItemType = 0
	ItemType_BANANA    ItemType = 1
	ItemType_CUCUMBER  ItemType = 2
	ItemType_COURGETTE ItemType = 3
	ItemType_CABBAGE   ItemType = 4
	ItemType_TURNIP    ItemType = 5
	ItemType_MUSHROOMS ItemType = 6
	ItemType_KALE      ItemType = 7
	ItemType_SPINACH   ItemType = 8
	ItemType_SUEDE     ItemType = 9
	ItemType_CARROT    ItemType = 10
	ItemType_PEAS      ItemType = 11
	ItemType_CORN      ItemType = 12
	ItemType_RADDISH   ItemType = 13
	ItemType_AUGERGINE ItemType = 14
	ItemType_ONION     ItemType = 16
	ItemType_POTATO    ItemType = 17
	ItemType_BEANS     ItemType = 18
	ItemType_AVOCADO   ItemType = 19
	ItemType_BROCCOLI  ItemType = 20
)

// Enum value maps for ItemType.
var (
	ItemType_name = map[int32]string{
		0:  "APPLE",
		1:  "BANANA",
		2:  "CUCUMBER",
		3:  "COURGETTE",
		4:  "CABBAGE",
		5:  "TURNIP",
		6:  "MUSHROOMS",
		7:  "KALE",
		8:  "SPINACH",
		9:  "SUEDE",
		10: "CARROT",
		11: "PEAS",
		12: "CORN",
		13: "RADDISH",
		14: "AUGERGINE",
		16: "ONION",
		17: "POTATO",
		18: "BEANS",
		19: "AVOCADO",
		20: "BROCCOLI",
	}
	ItemType_value = map[string]int32{
		"APPLE":     0,
		"BANANA":    1,
		"CUCUMBER":  2,
		"COURGETTE": 3,
		"CABBAGE":   4,
		"TURNIP":    5,
		"MUSHROOMS": 6,
		"KALE":      7,
		"SPINACH":   8,
		"SUEDE":     9,
		"CARROT":    10,
		"PEAS":      11,
		"CORN":      12,
		"RADDISH":   13,
		"AUGERGINE": 14,
		"ONION":     16,
		"POTATO":    17,
		"BEANS":     18,
		"AVOCADO":   19,
		"BROCCOLI":  20,
	}
)

func (x ItemType) Enum() *ItemType {
	p := new(ItemType)
	*p = x
	return p
}

func (x ItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_calculator_v1_service_proto_enumTypes[0].Descriptor()
}

func (ItemType) Type() protoreflect.EnumType {
	return &file_src_calculator_v1_service_proto_enumTypes[0]
}

func (x ItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemType.Descriptor instead.
func (ItemType) EnumDescriptor() ([]byte, []int) {
	return file_src_calculator_v1_service_proto_rawDescGZIP(), []int{0}
}

type CalculationItemOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item  ItemType `protobuf:"varint,1,opt,name=item,proto3,enum=calculator.v1.ItemType" json:"item,omitempty"`
	Count int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CalculationItemOutput) Reset() {
	*x = CalculationItemOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculator_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculationItemOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculationItemOutput) ProtoMessage() {}

func (x *CalculationItemOutput) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculator_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculationItemOutput.ProtoReflect.Descriptor instead.
func (*CalculationItemOutput) Descriptor() ([]byte, []int) {
	return file_src_calculator_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *CalculationItemOutput) GetItem() ItemType {
	if x != nil {
		return x.Item
	}
	return ItemType_APPLE
}

func (x *CalculationItemOutput) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CalculationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item        []ItemType `protobuf:"varint,1,rep,packed,name=item,proto3,enum=calculator.v1.ItemType" json:"item,omitempty"`
	BasketOwner string     `protobuf:"bytes,2,opt,name=basket_owner,json=basketOwner,proto3" json:"basket_owner,omitempty"`
}

func (x *CalculationInput) Reset() {
	*x = CalculationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculator_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculationInput) ProtoMessage() {}

func (x *CalculationInput) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculator_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculationInput.ProtoReflect.Descriptor instead.
func (*CalculationInput) Descriptor() ([]byte, []int) {
	return file_src_calculator_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *CalculationInput) GetItem() []ItemType {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *CalculationInput) GetBasketOwner() string {
	if x != nil {
		return x.BasketOwner
	}
	return ""
}

type CalculationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output      []*CalculationItemOutput `protobuf:"bytes,1,rep,name=output,proto3" json:"output,omitempty"`
	BasketOwner string                   `protobuf:"bytes,2,opt,name=basket_owner,json=basketOwner,proto3" json:"basket_owner,omitempty"`
}

func (x *CalculationResult) Reset() {
	*x = CalculationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_calculator_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculationResult) ProtoMessage() {}

func (x *CalculationResult) ProtoReflect() protoreflect.Message {
	mi := &file_src_calculator_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculationResult.ProtoReflect.Descriptor instead.
func (*CalculationResult) Descriptor() ([]byte, []int) {
	return file_src_calculator_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *CalculationResult) GetOutput() []*CalculationItemOutput {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *CalculationResult) GetBasketOwner() string {
	if x != nil {
		return x.BasketOwner
	}
	return ""
}

var File_src_calculator_v1_service_proto protoreflect.FileDescriptor

var file_src_calculator_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0x5a, 0x0a, 0x15, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x74, 0x65, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62, 0x0a, 0x10,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x22, 0x74, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x2a, 0x81, 0x02, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x50, 0x50, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x42, 0x41, 0x4e, 0x41, 0x4e, 0x41, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55,
	0x43, 0x55, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x55, 0x52,
	0x47, 0x45, 0x54, 0x54, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41, 0x42, 0x42, 0x41,
	0x47, 0x45, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x55, 0x52, 0x4e, 0x49, 0x50, 0x10, 0x05,
	0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x55, 0x53, 0x48, 0x52, 0x4f, 0x4f, 0x4d, 0x53, 0x10, 0x06, 0x12,
	0x08, 0x0a, 0x04, 0x4b, 0x41, 0x4c, 0x45, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x50, 0x49,
	0x4e, 0x41, 0x43, 0x48, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x55, 0x45, 0x44, 0x45, 0x10,
	0x09, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x41, 0x52, 0x52, 0x4f, 0x54, 0x10, 0x0a, 0x12, 0x08, 0x0a,
	0x04, 0x50, 0x45, 0x41, 0x53, 0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x52, 0x4e, 0x10,
	0x0c, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x41, 0x44, 0x44, 0x49, 0x53, 0x48, 0x10, 0x0d, 0x12, 0x0d,
	0x0a, 0x09, 0x41, 0x55, 0x47, 0x45, 0x52, 0x47, 0x49, 0x4e, 0x45, 0x10, 0x0e, 0x12, 0x09, 0x0a,
	0x05, 0x4f, 0x4e, 0x49, 0x4f, 0x4e, 0x10, 0x10, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x4f, 0x54, 0x41,
	0x54, 0x4f, 0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x45, 0x41, 0x4e, 0x53, 0x10, 0x12, 0x12,
	0x0b, 0x0a, 0x07, 0x41, 0x56, 0x4f, 0x43, 0x41, 0x44, 0x4f, 0x10, 0x13, 0x12, 0x0c, 0x0a, 0x08,
	0x42, 0x52, 0x4f, 0x43, 0x43, 0x4f, 0x4c, 0x49, 0x10, 0x14, 0x32, 0x67, 0x0a, 0x0a, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x59, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x20, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_calculator_v1_service_proto_rawDescOnce sync.Once
	file_src_calculator_v1_service_proto_rawDescData = file_src_calculator_v1_service_proto_rawDesc
)

func file_src_calculator_v1_service_proto_rawDescGZIP() []byte {
	file_src_calculator_v1_service_proto_rawDescOnce.Do(func() {
		file_src_calculator_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_calculator_v1_service_proto_rawDescData)
	})
	return file_src_calculator_v1_service_proto_rawDescData
}

var file_src_calculator_v1_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_calculator_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_calculator_v1_service_proto_goTypes = []interface{}{
	(ItemType)(0),                 // 0: calculator.v1.ItemType
	(*CalculationItemOutput)(nil), // 1: calculator.v1.CalculationItemOutput
	(*CalculationInput)(nil),      // 2: calculator.v1.CalculationInput
	(*CalculationResult)(nil),     // 3: calculator.v1.CalculationResult
}
var file_src_calculator_v1_service_proto_depIdxs = []int32{
	0, // 0: calculator.v1.CalculationItemOutput.item:type_name -> calculator.v1.ItemType
	0, // 1: calculator.v1.CalculationInput.item:type_name -> calculator.v1.ItemType
	1, // 2: calculator.v1.CalculationResult.output:type_name -> calculator.v1.CalculationItemOutput
	2, // 3: calculator.v1.Calculator.CalculateDataPoint:input_type -> calculator.v1.CalculationInput
	3, // 4: calculator.v1.Calculator.CalculateDataPoint:output_type -> calculator.v1.CalculationResult
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_src_calculator_v1_service_proto_init() }
func file_src_calculator_v1_service_proto_init() {
	if File_src_calculator_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_calculator_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculationItemOutput); i {
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
		file_src_calculator_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculationInput); i {
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
		file_src_calculator_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculationResult); i {
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
			RawDescriptor: file_src_calculator_v1_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_calculator_v1_service_proto_goTypes,
		DependencyIndexes: file_src_calculator_v1_service_proto_depIdxs,
		EnumInfos:         file_src_calculator_v1_service_proto_enumTypes,
		MessageInfos:      file_src_calculator_v1_service_proto_msgTypes,
	}.Build()
	File_src_calculator_v1_service_proto = out.File
	file_src_calculator_v1_service_proto_rawDesc = nil
	file_src_calculator_v1_service_proto_goTypes = nil
	file_src_calculator_v1_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorClient interface {
	CalculateDataPoint(ctx context.Context, in *CalculationInput, opts ...grpc.CallOption) (*CalculationResult, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) CalculateDataPoint(ctx context.Context, in *CalculationInput, opts ...grpc.CallOption) (*CalculationResult, error) {
	out := new(CalculationResult)
	err := c.cc.Invoke(ctx, "/calculator.v1.Calculator/CalculateDataPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServer is the server API for Calculator service.
type CalculatorServer interface {
	CalculateDataPoint(context.Context, *CalculationInput) (*CalculationResult, error)
}

// UnimplementedCalculatorServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (*UnimplementedCalculatorServer) CalculateDataPoint(context.Context, *CalculationInput) (*CalculationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateDataPoint not implemented")
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_CalculateDataPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).CalculateDataPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.v1.Calculator/CalculateDataPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).CalculateDataPoint(ctx, req.(*CalculationInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.v1.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateDataPoint",
			Handler:    _Calculator_CalculateDataPoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/calculator/v1/service.proto",
}
