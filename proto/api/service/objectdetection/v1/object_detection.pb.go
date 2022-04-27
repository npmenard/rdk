// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/service/objectdetection/v1/object_detection.proto

package v1

import (
	_ "go.viam.com/rdk/proto/api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DetectorNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DetectorNamesRequest) Reset() {
	*x = DetectorNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectorNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectorNamesRequest) ProtoMessage() {}

func (x *DetectorNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectorNamesRequest.ProtoReflect.Descriptor instead.
func (*DetectorNamesRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescGZIP(), []int{0}
}

type DetectorNamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// detectors in the registry
	DetectorNames []string `protobuf:"bytes,1,rep,name=detector_names,json=detectorNames,proto3" json:"detector_names,omitempty"`
}

func (x *DetectorNamesResponse) Reset() {
	*x = DetectorNamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectorNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectorNamesResponse) ProtoMessage() {}

func (x *DetectorNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectorNamesResponse.ProtoReflect.Descriptor instead.
func (*DetectorNamesResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescGZIP(), []int{1}
}

func (x *DetectorNamesResponse) GetDetectorNames() []string {
	if x != nil {
		return x.DetectorNames
	}
	return nil
}

type AddDetectorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetectorName       string           `protobuf:"bytes,1,opt,name=detector_name,json=detectorName,proto3" json:"detector_name,omitempty"`
	DetectorModelType  string           `protobuf:"bytes,2,opt,name=detector_model_type,json=detectorModelType,proto3" json:"detector_model_type,omitempty"`
	DetectorParameters *structpb.Struct `protobuf:"bytes,3,opt,name=detector_parameters,json=detectorParameters,proto3" json:"detector_parameters,omitempty"`
}

func (x *AddDetectorRequest) Reset() {
	*x = AddDetectorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDetectorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDetectorRequest) ProtoMessage() {}

func (x *AddDetectorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDetectorRequest.ProtoReflect.Descriptor instead.
func (*AddDetectorRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescGZIP(), []int{2}
}

func (x *AddDetectorRequest) GetDetectorName() string {
	if x != nil {
		return x.DetectorName
	}
	return ""
}

func (x *AddDetectorRequest) GetDetectorModelType() string {
	if x != nil {
		return x.DetectorModelType
	}
	return ""
}

func (x *AddDetectorRequest) GetDetectorParameters() *structpb.Struct {
	if x != nil {
		return x.DetectorParameters
	}
	return nil
}

type AddDetectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddDetectorResponse) Reset() {
	*x = AddDetectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDetectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDetectorResponse) ProtoMessage() {}

func (x *AddDetectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDetectorResponse.ProtoReflect.Descriptor instead.
func (*AddDetectorResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescGZIP(), []int{3}
}

func (x *AddDetectorResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DetectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of camera source to use as input
	CameraName string `protobuf:"bytes,1,opt,name=camera_name,json=cameraName,proto3" json:"camera_name,omitempty"`
	// name of the registered detector to use
	DetectorName string `protobuf:"bytes,2,opt,name=detector_name,json=detectorName,proto3" json:"detector_name,omitempty"`
}

func (x *DetectRequest) Reset() {
	*x = DetectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectRequest) ProtoMessage() {}

func (x *DetectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectRequest.ProtoReflect.Descriptor instead.
func (*DetectRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescGZIP(), []int{4}
}

func (x *DetectRequest) GetCameraName() string {
	if x != nil {
		return x.CameraName
	}
	return ""
}

func (x *DetectRequest) GetDetectorName() string {
	if x != nil {
		return x.DetectorName
	}
	return ""
}

type DetectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the bounding boxes and labels
	Detections []*Detection `protobuf:"bytes,1,rep,name=detections,proto3" json:"detections,omitempty"`
}

func (x *DetectResponse) Reset() {
	*x = DetectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectResponse) ProtoMessage() {}

func (x *DetectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectResponse.ProtoReflect.Descriptor instead.
func (*DetectResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescGZIP(), []int{5}
}

func (x *DetectResponse) GetDetections() []*Detection {
	if x != nil {
		return x.Detections
	}
	return nil
}

type Detection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the four corners of the box
	XMin int64 `protobuf:"varint,1,opt,name=x_min,json=xMin,proto3" json:"x_min,omitempty"`
	YMin int64 `protobuf:"varint,2,opt,name=y_min,json=yMin,proto3" json:"y_min,omitempty"`
	XMax int64 `protobuf:"varint,3,opt,name=x_max,json=xMax,proto3" json:"x_max,omitempty"`
	YMax int64 `protobuf:"varint,4,opt,name=y_max,json=yMax,proto3" json:"y_max,omitempty"`
	// the confidence of the detection
	Confidence float64 `protobuf:"fixed64,5,opt,name=confidence,proto3" json:"confidence,omitempty"`
	// label associated with the detected object
	ClassName string `protobuf:"bytes,6,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
}

func (x *Detection) Reset() {
	*x = Detection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detection) ProtoMessage() {}

func (x *Detection) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detection.ProtoReflect.Descriptor instead.
func (*Detection) Descriptor() ([]byte, []int) {
	return file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescGZIP(), []int{6}
}

func (x *Detection) GetXMin() int64 {
	if x != nil {
		return x.XMin
	}
	return 0
}

func (x *Detection) GetYMin() int64 {
	if x != nil {
		return x.YMin
	}
	return 0
}

func (x *Detection) GetXMax() int64 {
	if x != nil {
		return x.XMax
	}
	return 0
}

func (x *Detection) GetYMax() int64 {
	if x != nil {
		return x.YMax
	}
	return 0
}

func (x *Detection) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *Detection) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

var File_proto_api_service_objectdetection_v1_object_detection_proto protoreflect.FileDescriptor

var file_proto_api_service_objectdetection_v1_object_detection_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x15, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x12, 0x41, 0x64,
	0x64, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x13, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x12, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x22,
	0x2f, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x55, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x65, 0x72, 0x61, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x09, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x78, 0x5f, 0x6d, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x78, 0x4d, 0x69, 0x6e, 0x12, 0x13, 0x0a,
	0x05, 0x79, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x4d,
	0x69, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x78, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x78, 0x4d, 0x61, 0x78, 0x12, 0x13, 0x0a, 0x05, 0x79, 0x5f, 0x6d, 0x61, 0x78,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x4d, 0x61, 0x78, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xce, 0x04, 0x0a, 0x16,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc6, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x3a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f, 0x76, 0x69, 0x61, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0xbe, 0x01, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x38, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x22, 0x32, 0x2f, 0x76,
	0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x64, 0x64, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0xa9, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x12, 0x33, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22, 0x2c,
	0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x42, 0x69, 0x0a, 0x31,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x72, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescOnce sync.Once
	file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescData = file_proto_api_service_objectdetection_v1_object_detection_proto_rawDesc
)

func file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescGZIP() []byte {
	file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescOnce.Do(func() {
		file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescData)
	})
	return file_proto_api_service_objectdetection_v1_object_detection_proto_rawDescData
}

var file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_api_service_objectdetection_v1_object_detection_proto_goTypes = []interface{}{
	(*DetectorNamesRequest)(nil),  // 0: proto.api.service.objectdetection.v1.DetectorNamesRequest
	(*DetectorNamesResponse)(nil), // 1: proto.api.service.objectdetection.v1.DetectorNamesResponse
	(*AddDetectorRequest)(nil),    // 2: proto.api.service.objectdetection.v1.AddDetectorRequest
	(*AddDetectorResponse)(nil),   // 3: proto.api.service.objectdetection.v1.AddDetectorResponse
	(*DetectRequest)(nil),         // 4: proto.api.service.objectdetection.v1.DetectRequest
	(*DetectResponse)(nil),        // 5: proto.api.service.objectdetection.v1.DetectResponse
	(*Detection)(nil),             // 6: proto.api.service.objectdetection.v1.Detection
	(*structpb.Struct)(nil),       // 7: google.protobuf.Struct
}
var file_proto_api_service_objectdetection_v1_object_detection_proto_depIdxs = []int32{
	7, // 0: proto.api.service.objectdetection.v1.AddDetectorRequest.detector_parameters:type_name -> google.protobuf.Struct
	6, // 1: proto.api.service.objectdetection.v1.DetectResponse.detections:type_name -> proto.api.service.objectdetection.v1.Detection
	0, // 2: proto.api.service.objectdetection.v1.ObjectDetectionService.DetectorNames:input_type -> proto.api.service.objectdetection.v1.DetectorNamesRequest
	2, // 3: proto.api.service.objectdetection.v1.ObjectDetectionService.AddDetector:input_type -> proto.api.service.objectdetection.v1.AddDetectorRequest
	4, // 4: proto.api.service.objectdetection.v1.ObjectDetectionService.Detect:input_type -> proto.api.service.objectdetection.v1.DetectRequest
	1, // 5: proto.api.service.objectdetection.v1.ObjectDetectionService.DetectorNames:output_type -> proto.api.service.objectdetection.v1.DetectorNamesResponse
	3, // 6: proto.api.service.objectdetection.v1.ObjectDetectionService.AddDetector:output_type -> proto.api.service.objectdetection.v1.AddDetectorResponse
	5, // 7: proto.api.service.objectdetection.v1.ObjectDetectionService.Detect:output_type -> proto.api.service.objectdetection.v1.DetectResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_api_service_objectdetection_v1_object_detection_proto_init() }
func file_proto_api_service_objectdetection_v1_object_detection_proto_init() {
	if File_proto_api_service_objectdetection_v1_object_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectorNamesRequest); i {
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
		file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectorNamesResponse); i {
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
		file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDetectorRequest); i {
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
		file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDetectorResponse); i {
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
		file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectRequest); i {
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
		file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectResponse); i {
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
		file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detection); i {
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
			RawDescriptor: file_proto_api_service_objectdetection_v1_object_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_service_objectdetection_v1_object_detection_proto_goTypes,
		DependencyIndexes: file_proto_api_service_objectdetection_v1_object_detection_proto_depIdxs,
		MessageInfos:      file_proto_api_service_objectdetection_v1_object_detection_proto_msgTypes,
	}.Build()
	File_proto_api_service_objectdetection_v1_object_detection_proto = out.File
	file_proto_api_service_objectdetection_v1_object_detection_proto_rawDesc = nil
	file_proto_api_service_objectdetection_v1_object_detection_proto_goTypes = nil
	file_proto_api_service_objectdetection_v1_object_detection_proto_depIdxs = nil
}