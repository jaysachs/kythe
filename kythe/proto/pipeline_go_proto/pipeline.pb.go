// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: kythe/proto/pipeline.proto

package pipeline_go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common_go_proto "kythe.io/kythe/proto/common_go_proto"
	schema_go_proto "kythe.io/kythe/proto/schema_go_proto"
	serving_go_proto "kythe.io/kythe/proto/serving_go_proto"
	storage_go_proto "kythe.io/kythe/proto/storage_go_proto"
	xref_serving_go_proto "kythe.io/kythe/proto/xref_serving_go_proto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source *storage_go_proto.VName `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Types that are assignable to Kind:
	//	*Reference_KytheKind
	//	*Reference_GenericKind
	Kind   isReference_Kind                 `protobuf_oneof:"kind"`
	Anchor *serving_go_proto.ExpandedAnchor `protobuf:"bytes,4,opt,name=anchor,proto3" json:"anchor,omitempty"`
	Scope  *storage_go_proto.VName          `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_kythe_proto_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *Reference) GetSource() *storage_go_proto.VName {
	if x != nil {
		return x.Source
	}
	return nil
}

func (m *Reference) GetKind() isReference_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Reference) GetKytheKind() schema_go_proto.EdgeKind {
	if x, ok := x.GetKind().(*Reference_KytheKind); ok {
		return x.KytheKind
	}
	return schema_go_proto.EdgeKind(0)
}

func (x *Reference) GetGenericKind() string {
	if x, ok := x.GetKind().(*Reference_GenericKind); ok {
		return x.GenericKind
	}
	return ""
}

func (x *Reference) GetAnchor() *serving_go_proto.ExpandedAnchor {
	if x != nil {
		return x.Anchor
	}
	return nil
}

func (x *Reference) GetScope() *storage_go_proto.VName {
	if x != nil {
		return x.Scope
	}
	return nil
}

type isReference_Kind interface {
	isReference_Kind()
}

type Reference_KytheKind struct {
	KytheKind schema_go_proto.EdgeKind `protobuf:"varint,2,opt,name=kythe_kind,json=kytheKind,proto3,enum=kythe.proto.schema.EdgeKind,oneof"`
}

type Reference_GenericKind struct {
	GenericKind string `protobuf:"bytes,3,opt,name=generic_kind,json=genericKind,proto3,oneof"`
}

func (*Reference_KytheKind) isReference_Kind() {}

func (*Reference_GenericKind) isReference_Kind() {}

type DecorationPiece struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileVName *storage_go_proto.VName `protobuf:"bytes,1,opt,name=file_v_name,json=fileVName,proto3" json:"file_v_name,omitempty"`
	// Types that are assignable to Piece:
	//	*DecorationPiece_File
	//	*DecorationPiece_Reference
	//	*DecorationPiece_Node
	//	*DecorationPiece_Definition_
	//	*DecorationPiece_Diagnostic
	//	*DecorationPiece_TargetOverride
	Piece isDecorationPiece_Piece `protobuf_oneof:"piece"`
}

func (x *DecorationPiece) Reset() {
	*x = DecorationPiece{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecorationPiece) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecorationPiece) ProtoMessage() {}

func (x *DecorationPiece) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecorationPiece.ProtoReflect.Descriptor instead.
func (*DecorationPiece) Descriptor() ([]byte, []int) {
	return file_kythe_proto_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *DecorationPiece) GetFileVName() *storage_go_proto.VName {
	if x != nil {
		return x.FileVName
	}
	return nil
}

func (m *DecorationPiece) GetPiece() isDecorationPiece_Piece {
	if m != nil {
		return m.Piece
	}
	return nil
}

func (x *DecorationPiece) GetFile() *serving_go_proto.File {
	if x, ok := x.GetPiece().(*DecorationPiece_File); ok {
		return x.File
	}
	return nil
}

func (x *DecorationPiece) GetReference() *Reference {
	if x, ok := x.GetPiece().(*DecorationPiece_Reference); ok {
		return x.Reference
	}
	return nil
}

func (x *DecorationPiece) GetNode() *schema_go_proto.Node {
	if x, ok := x.GetPiece().(*DecorationPiece_Node); ok {
		return x.Node
	}
	return nil
}

func (x *DecorationPiece) GetDefinition() *DecorationPiece_Definition {
	if x, ok := x.GetPiece().(*DecorationPiece_Definition_); ok {
		return x.Definition
	}
	return nil
}

func (x *DecorationPiece) GetDiagnostic() *common_go_proto.Diagnostic {
	if x, ok := x.GetPiece().(*DecorationPiece_Diagnostic); ok {
		return x.Diagnostic
	}
	return nil
}

func (x *DecorationPiece) GetTargetOverride() *xref_serving_go_proto.FileDecorations_TargetOverride {
	if x, ok := x.GetPiece().(*DecorationPiece_TargetOverride); ok {
		return x.TargetOverride
	}
	return nil
}

type isDecorationPiece_Piece interface {
	isDecorationPiece_Piece()
}

type DecorationPiece_File struct {
	File *serving_go_proto.File `protobuf:"bytes,2,opt,name=file,proto3,oneof"`
}

type DecorationPiece_Reference struct {
	Reference *Reference `protobuf:"bytes,3,opt,name=reference,proto3,oneof"`
}

type DecorationPiece_Node struct {
	Node *schema_go_proto.Node `protobuf:"bytes,4,opt,name=node,proto3,oneof"`
}

type DecorationPiece_Definition_ struct {
	Definition *DecorationPiece_Definition `protobuf:"bytes,5,opt,name=definition,proto3,oneof"`
}

type DecorationPiece_Diagnostic struct {
	Diagnostic *common_go_proto.Diagnostic `protobuf:"bytes,6,opt,name=diagnostic,proto3,oneof"`
}

type DecorationPiece_TargetOverride struct {
	TargetOverride *xref_serving_go_proto.FileDecorations_TargetOverride `protobuf:"bytes,7,opt,name=target_override,json=targetOverride,proto3,oneof"`
}

func (*DecorationPiece_File) isDecorationPiece_Piece() {}

func (*DecorationPiece_Reference) isDecorationPiece_Piece() {}

func (*DecorationPiece_Node) isDecorationPiece_Piece() {}

func (*DecorationPiece_Definition_) isDecorationPiece_Piece() {}

func (*DecorationPiece_Diagnostic) isDecorationPiece_Piece() {}

func (*DecorationPiece_TargetOverride) isDecorationPiece_Piece() {}

type DecorationPiece_Definition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node       *storage_go_proto.VName          `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Definition *serving_go_proto.ExpandedAnchor `protobuf:"bytes,2,opt,name=definition,proto3" json:"definition,omitempty"`
}

func (x *DecorationPiece_Definition) Reset() {
	*x = DecorationPiece_Definition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kythe_proto_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecorationPiece_Definition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecorationPiece_Definition) ProtoMessage() {}

func (x *DecorationPiece_Definition) ProtoReflect() protoreflect.Message {
	mi := &file_kythe_proto_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecorationPiece_Definition.ProtoReflect.Descriptor instead.
func (*DecorationPiece_Definition) Descriptor() ([]byte, []int) {
	return file_kythe_proto_pipeline_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DecorationPiece_Definition) GetNode() *storage_go_proto.VName {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *DecorationPiece_Definition) GetDefinition() *serving_go_proto.ExpandedAnchor {
	if x != nil {
		return x.Definition
	}
	return nil
}

var File_kythe_proto_pipeline_proto protoreflect.FileDescriptor

var file_kythe_proto_pipeline_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x1a, 0x18, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x72, 0x65, 0x66, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a,
	0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x79, 0x74,
	0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x5f,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6b, 0x79, 0x74,
	0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x45, 0x64, 0x67, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x09, 0x6b, 0x79, 0x74, 0x68,
	0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x06, 0x61, 0x6e,
	0x63, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x79, 0x74,
	0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52,
	0x06, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0xe7, 0x04, 0x0a, 0x0f, 0x44, 0x65,
	0x63, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x12, 0x32, 0x0a,
	0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x76, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x56, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2f, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44,
	0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0a, 0x64, 0x69, 0x61, 0x67, 0x6e,
	0x6f, 0x73, 0x74, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x44, 0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x48, 0x00, 0x52, 0x0a, 0x64,
	0x69, 0x61, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x12, 0x64, 0x0a, 0x0f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x78, 0x72, 0x65, 0x66, 0x73, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x48, 0x00, 0x52,
	0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x1a,
	0x79, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x79,
	0x74, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x79, 0x74, 0x68,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x52, 0x0a,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x70, 0x69,
	0x65, 0x63, 0x65, 0x42, 0x34, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x6b, 0x79, 0x74, 0x68, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x11, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_kythe_proto_pipeline_proto_rawDescOnce sync.Once
	file_kythe_proto_pipeline_proto_rawDescData = file_kythe_proto_pipeline_proto_rawDesc
)

func file_kythe_proto_pipeline_proto_rawDescGZIP() []byte {
	file_kythe_proto_pipeline_proto_rawDescOnce.Do(func() {
		file_kythe_proto_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_kythe_proto_pipeline_proto_rawDescData)
	})
	return file_kythe_proto_pipeline_proto_rawDescData
}

var file_kythe_proto_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kythe_proto_pipeline_proto_goTypes = []interface{}{
	(*Reference)(nil),                                            // 0: kythe.proto.pipeline.Reference
	(*DecorationPiece)(nil),                                      // 1: kythe.proto.pipeline.DecorationPiece
	(*DecorationPiece_Definition)(nil),                           // 2: kythe.proto.pipeline.DecorationPiece.Definition
	(*storage_go_proto.VName)(nil),                               // 3: kythe.proto.VName
	(schema_go_proto.EdgeKind)(0),                                // 4: kythe.proto.schema.EdgeKind
	(*serving_go_proto.ExpandedAnchor)(nil),                      // 5: kythe.proto.serving.ExpandedAnchor
	(*serving_go_proto.File)(nil),                                // 6: kythe.proto.serving.File
	(*schema_go_proto.Node)(nil),                                 // 7: kythe.proto.schema.Node
	(*common_go_proto.Diagnostic)(nil),                           // 8: kythe.proto.common.Diagnostic
	(*xref_serving_go_proto.FileDecorations_TargetOverride)(nil), // 9: kythe.proto.serving.xrefs.FileDecorations.TargetOverride
}
var file_kythe_proto_pipeline_proto_depIdxs = []int32{
	3,  // 0: kythe.proto.pipeline.Reference.source:type_name -> kythe.proto.VName
	4,  // 1: kythe.proto.pipeline.Reference.kythe_kind:type_name -> kythe.proto.schema.EdgeKind
	5,  // 2: kythe.proto.pipeline.Reference.anchor:type_name -> kythe.proto.serving.ExpandedAnchor
	3,  // 3: kythe.proto.pipeline.Reference.scope:type_name -> kythe.proto.VName
	3,  // 4: kythe.proto.pipeline.DecorationPiece.file_v_name:type_name -> kythe.proto.VName
	6,  // 5: kythe.proto.pipeline.DecorationPiece.file:type_name -> kythe.proto.serving.File
	0,  // 6: kythe.proto.pipeline.DecorationPiece.reference:type_name -> kythe.proto.pipeline.Reference
	7,  // 7: kythe.proto.pipeline.DecorationPiece.node:type_name -> kythe.proto.schema.Node
	2,  // 8: kythe.proto.pipeline.DecorationPiece.definition:type_name -> kythe.proto.pipeline.DecorationPiece.Definition
	8,  // 9: kythe.proto.pipeline.DecorationPiece.diagnostic:type_name -> kythe.proto.common.Diagnostic
	9,  // 10: kythe.proto.pipeline.DecorationPiece.target_override:type_name -> kythe.proto.serving.xrefs.FileDecorations.TargetOverride
	3,  // 11: kythe.proto.pipeline.DecorationPiece.Definition.node:type_name -> kythe.proto.VName
	5,  // 12: kythe.proto.pipeline.DecorationPiece.Definition.definition:type_name -> kythe.proto.serving.ExpandedAnchor
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_kythe_proto_pipeline_proto_init() }
func file_kythe_proto_pipeline_proto_init() {
	if File_kythe_proto_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kythe_proto_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
		file_kythe_proto_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecorationPiece); i {
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
		file_kythe_proto_pipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecorationPiece_Definition); i {
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
	file_kythe_proto_pipeline_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Reference_KytheKind)(nil),
		(*Reference_GenericKind)(nil),
	}
	file_kythe_proto_pipeline_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DecorationPiece_File)(nil),
		(*DecorationPiece_Reference)(nil),
		(*DecorationPiece_Node)(nil),
		(*DecorationPiece_Definition_)(nil),
		(*DecorationPiece_Diagnostic)(nil),
		(*DecorationPiece_TargetOverride)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kythe_proto_pipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kythe_proto_pipeline_proto_goTypes,
		DependencyIndexes: file_kythe_proto_pipeline_proto_depIdxs,
		MessageInfos:      file_kythe_proto_pipeline_proto_msgTypes,
	}.Build()
	File_kythe_proto_pipeline_proto = out.File
	file_kythe_proto_pipeline_proto_rawDesc = nil
	file_kythe_proto_pipeline_proto_goTypes = nil
	file_kythe_proto_pipeline_proto_depIdxs = nil
}
