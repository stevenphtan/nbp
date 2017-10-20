// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/v1beta1/write.proto

package firestore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A value that is calculated by the server.
type DocumentTransform_FieldTransform_ServerValue int32

const (
	// Unspecified. This value must not be used.
	DocumentTransform_FieldTransform_SERVER_VALUE_UNSPECIFIED DocumentTransform_FieldTransform_ServerValue = 0
	// The time at which the server processed the request.
	DocumentTransform_FieldTransform_REQUEST_TIME DocumentTransform_FieldTransform_ServerValue = 1
)

var DocumentTransform_FieldTransform_ServerValue_name = map[int32]string{
	0: "SERVER_VALUE_UNSPECIFIED",
	1: "REQUEST_TIME",
}
var DocumentTransform_FieldTransform_ServerValue_value = map[string]int32{
	"SERVER_VALUE_UNSPECIFIED": 0,
	"REQUEST_TIME":             1,
}

func (x DocumentTransform_FieldTransform_ServerValue) String() string {
	return proto.EnumName(DocumentTransform_FieldTransform_ServerValue_name, int32(x))
}
func (DocumentTransform_FieldTransform_ServerValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{1, 0, 0}
}

// A write on a document.
type Write struct {
	// The operation to execute.
	//
	// Types that are valid to be assigned to Operation:
	//	*Write_Update
	//	*Write_Delete
	//	*Write_Transform
	Operation isWrite_Operation `protobuf_oneof:"operation"`
	// The fields to update in this write.
	//
	// This field can be set only when the operation is `update`.
	// None of the field paths in the mask may contain a reserved name.
	// If the document exists on the server and has fields not referenced in the
	// mask, they are left unchanged.
	// Fields referenced in the mask, but not present in the input document, are
	// deleted from the document on the server.
	// The field paths in this mask must not contain a reserved field name.
	UpdateMask *DocumentMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
	// An optional precondition on the document.
	//
	// The write will fail if this is set and not met by the target document.
	CurrentDocument *Precondition `protobuf:"bytes,4,opt,name=current_document,json=currentDocument" json:"current_document,omitempty"`
}

func (m *Write) Reset()                    { *m = Write{} }
func (m *Write) String() string            { return proto.CompactTextString(m) }
func (*Write) ProtoMessage()               {}
func (*Write) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type isWrite_Operation interface {
	isWrite_Operation()
}

type Write_Update struct {
	Update *Document `protobuf:"bytes,1,opt,name=update,oneof"`
}
type Write_Delete struct {
	Delete string `protobuf:"bytes,2,opt,name=delete,oneof"`
}
type Write_Transform struct {
	Transform *DocumentTransform `protobuf:"bytes,6,opt,name=transform,oneof"`
}

func (*Write_Update) isWrite_Operation()    {}
func (*Write_Delete) isWrite_Operation()    {}
func (*Write_Transform) isWrite_Operation() {}

func (m *Write) GetOperation() isWrite_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *Write) GetUpdate() *Document {
	if x, ok := m.GetOperation().(*Write_Update); ok {
		return x.Update
	}
	return nil
}

func (m *Write) GetDelete() string {
	if x, ok := m.GetOperation().(*Write_Delete); ok {
		return x.Delete
	}
	return ""
}

func (m *Write) GetTransform() *DocumentTransform {
	if x, ok := m.GetOperation().(*Write_Transform); ok {
		return x.Transform
	}
	return nil
}

func (m *Write) GetUpdateMask() *DocumentMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *Write) GetCurrentDocument() *Precondition {
	if m != nil {
		return m.CurrentDocument
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Write) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Write_OneofMarshaler, _Write_OneofUnmarshaler, _Write_OneofSizer, []interface{}{
		(*Write_Update)(nil),
		(*Write_Delete)(nil),
		(*Write_Transform)(nil),
	}
}

func _Write_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Write)
	// operation
	switch x := m.Operation.(type) {
	case *Write_Update:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *Write_Delete:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Delete)
	case *Write_Transform:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transform); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Write.Operation has unexpected type %T", x)
	}
	return nil
}

func _Write_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Write)
	switch tag {
	case 1: // operation.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Document)
		err := b.DecodeMessage(msg)
		m.Operation = &Write_Update{msg}
		return true, err
	case 2: // operation.delete
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Operation = &Write_Delete{x}
		return true, err
	case 6: // operation.transform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DocumentTransform)
		err := b.DecodeMessage(msg)
		m.Operation = &Write_Transform{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Write_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Write)
	// operation
	switch x := m.Operation.(type) {
	case *Write_Update:
		s := proto.Size(x.Update)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Write_Delete:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Delete)))
		n += len(x.Delete)
	case *Write_Transform:
		s := proto.Size(x.Transform)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A transformation of a document.
type DocumentTransform struct {
	// The name of the document to transform.
	Document string `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	// The list of transformations to apply to the fields of the document, in
	// order.
	FieldTransforms []*DocumentTransform_FieldTransform `protobuf:"bytes,2,rep,name=field_transforms,json=fieldTransforms" json:"field_transforms,omitempty"`
}

func (m *DocumentTransform) Reset()                    { *m = DocumentTransform{} }
func (m *DocumentTransform) String() string            { return proto.CompactTextString(m) }
func (*DocumentTransform) ProtoMessage()               {}
func (*DocumentTransform) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *DocumentTransform) GetDocument() string {
	if m != nil {
		return m.Document
	}
	return ""
}

func (m *DocumentTransform) GetFieldTransforms() []*DocumentTransform_FieldTransform {
	if m != nil {
		return m.FieldTransforms
	}
	return nil
}

// A transformation of a field of the document.
type DocumentTransform_FieldTransform struct {
	// The path of the field. See [Document.fields][google.firestore.v1beta1.Document.fields] for the field path syntax
	// reference.
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath" json:"field_path,omitempty"`
	// The transformation to apply on the field.
	//
	// Types that are valid to be assigned to TransformType:
	//	*DocumentTransform_FieldTransform_SetToServerValue
	TransformType isDocumentTransform_FieldTransform_TransformType `protobuf_oneof:"transform_type"`
}

func (m *DocumentTransform_FieldTransform) Reset()         { *m = DocumentTransform_FieldTransform{} }
func (m *DocumentTransform_FieldTransform) String() string { return proto.CompactTextString(m) }
func (*DocumentTransform_FieldTransform) ProtoMessage()    {}
func (*DocumentTransform_FieldTransform) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{1, 0}
}

type isDocumentTransform_FieldTransform_TransformType interface {
	isDocumentTransform_FieldTransform_TransformType()
}

type DocumentTransform_FieldTransform_SetToServerValue struct {
	SetToServerValue DocumentTransform_FieldTransform_ServerValue `protobuf:"varint,2,opt,name=set_to_server_value,json=setToServerValue,enum=google.firestore.v1beta1.DocumentTransform_FieldTransform_ServerValue,oneof"`
}

func (*DocumentTransform_FieldTransform_SetToServerValue) isDocumentTransform_FieldTransform_TransformType() {
}

func (m *DocumentTransform_FieldTransform) GetTransformType() isDocumentTransform_FieldTransform_TransformType {
	if m != nil {
		return m.TransformType
	}
	return nil
}

func (m *DocumentTransform_FieldTransform) GetFieldPath() string {
	if m != nil {
		return m.FieldPath
	}
	return ""
}

func (m *DocumentTransform_FieldTransform) GetSetToServerValue() DocumentTransform_FieldTransform_ServerValue {
	if x, ok := m.GetTransformType().(*DocumentTransform_FieldTransform_SetToServerValue); ok {
		return x.SetToServerValue
	}
	return DocumentTransform_FieldTransform_SERVER_VALUE_UNSPECIFIED
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DocumentTransform_FieldTransform) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DocumentTransform_FieldTransform_OneofMarshaler, _DocumentTransform_FieldTransform_OneofUnmarshaler, _DocumentTransform_FieldTransform_OneofSizer, []interface{}{
		(*DocumentTransform_FieldTransform_SetToServerValue)(nil),
	}
}

func _DocumentTransform_FieldTransform_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DocumentTransform_FieldTransform)
	// transform_type
	switch x := m.TransformType.(type) {
	case *DocumentTransform_FieldTransform_SetToServerValue:
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.SetToServerValue))
	case nil:
	default:
		return fmt.Errorf("DocumentTransform_FieldTransform.TransformType has unexpected type %T", x)
	}
	return nil
}

func _DocumentTransform_FieldTransform_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DocumentTransform_FieldTransform)
	switch tag {
	case 2: // transform_type.set_to_server_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.TransformType = &DocumentTransform_FieldTransform_SetToServerValue{DocumentTransform_FieldTransform_ServerValue(x)}
		return true, err
	default:
		return false, nil
	}
}

func _DocumentTransform_FieldTransform_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DocumentTransform_FieldTransform)
	// transform_type
	switch x := m.TransformType.(type) {
	case *DocumentTransform_FieldTransform_SetToServerValue:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.SetToServerValue))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The result of applying a write.
type WriteResult struct {
	// The last update time of the document after applying the write. Not set
	// after a `delete`.
	//
	// If the write did not actually change the document, this will be the
	// previous update_time.
	UpdateTime *google_protobuf1.Timestamp `protobuf:"bytes,1,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
	// The results of applying each [DocumentTransform.FieldTransform][google.firestore.v1beta1.DocumentTransform.FieldTransform], in the
	// same order.
	TransformResults []*Value `protobuf:"bytes,2,rep,name=transform_results,json=transformResults" json:"transform_results,omitempty"`
}

func (m *WriteResult) Reset()                    { *m = WriteResult{} }
func (m *WriteResult) String() string            { return proto.CompactTextString(m) }
func (*WriteResult) ProtoMessage()               {}
func (*WriteResult) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *WriteResult) GetUpdateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *WriteResult) GetTransformResults() []*Value {
	if m != nil {
		return m.TransformResults
	}
	return nil
}

// A [Document][google.firestore.v1beta1.Document] has changed.
//
// May be the result of multiple [writes][google.firestore.v1beta1.Write], including deletes, that
// ultimately resulted in a new value for the [Document][google.firestore.v1beta1.Document].
//
// Multiple [DocumentChange][google.firestore.v1beta1.DocumentChange] messages may be returned for the same logical
// change, if multiple targets are affected.
type DocumentChange struct {
	// The new state of the [Document][google.firestore.v1beta1.Document].
	//
	// If `mask` is set, contains only fields that were updated or added.
	Document *Document `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	// A set of target IDs of targets that match this document.
	TargetIds []int32 `protobuf:"varint,5,rep,packed,name=target_ids,json=targetIds" json:"target_ids,omitempty"`
	// A set of target IDs for targets that no longer match this document.
	RemovedTargetIds []int32 `protobuf:"varint,6,rep,packed,name=removed_target_ids,json=removedTargetIds" json:"removed_target_ids,omitempty"`
}

func (m *DocumentChange) Reset()                    { *m = DocumentChange{} }
func (m *DocumentChange) String() string            { return proto.CompactTextString(m) }
func (*DocumentChange) ProtoMessage()               {}
func (*DocumentChange) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *DocumentChange) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *DocumentChange) GetTargetIds() []int32 {
	if m != nil {
		return m.TargetIds
	}
	return nil
}

func (m *DocumentChange) GetRemovedTargetIds() []int32 {
	if m != nil {
		return m.RemovedTargetIds
	}
	return nil
}

// A [Document][google.firestore.v1beta1.Document] has been deleted.
//
// May be the result of multiple [writes][google.firestore.v1beta1.Write], including updates, the
// last of which deleted the [Document][google.firestore.v1beta1.Document].
//
// Multiple [DocumentDelete][google.firestore.v1beta1.DocumentDelete] messages may be returned for the same logical
// delete, if multiple targets are affected.
type DocumentDelete struct {
	// The resource name of the [Document][google.firestore.v1beta1.Document] that was deleted.
	Document string `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	// A set of target IDs for targets that previously matched this entity.
	RemovedTargetIds []int32 `protobuf:"varint,6,rep,packed,name=removed_target_ids,json=removedTargetIds" json:"removed_target_ids,omitempty"`
	// The read timestamp at which the delete was observed.
	//
	// Greater or equal to the `commit_time` of the delete.
	ReadTime *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=read_time,json=readTime" json:"read_time,omitempty"`
}

func (m *DocumentDelete) Reset()                    { *m = DocumentDelete{} }
func (m *DocumentDelete) String() string            { return proto.CompactTextString(m) }
func (*DocumentDelete) ProtoMessage()               {}
func (*DocumentDelete) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *DocumentDelete) GetDocument() string {
	if m != nil {
		return m.Document
	}
	return ""
}

func (m *DocumentDelete) GetRemovedTargetIds() []int32 {
	if m != nil {
		return m.RemovedTargetIds
	}
	return nil
}

func (m *DocumentDelete) GetReadTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

// A [Document][google.firestore.v1beta1.Document] has been removed from the view of the targets.
//
// Sent if the document is no longer relevant to a target and is out of view.
// Can be sent instead of a DocumentDelete or a DocumentChange if the server
// can not send the new value of the document.
//
// Multiple [DocumentRemove][google.firestore.v1beta1.DocumentRemove] messages may be returned for the same logical
// write or delete, if multiple targets are affected.
type DocumentRemove struct {
	// The resource name of the [Document][google.firestore.v1beta1.Document] that has gone out of view.
	Document string `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
	// A set of target IDs for targets that previously matched this document.
	RemovedTargetIds []int32 `protobuf:"varint,2,rep,packed,name=removed_target_ids,json=removedTargetIds" json:"removed_target_ids,omitempty"`
	// The read timestamp at which the remove was observed.
	//
	// Greater or equal to the `commit_time` of the change/delete/remove.
	ReadTime *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=read_time,json=readTime" json:"read_time,omitempty"`
}

func (m *DocumentRemove) Reset()                    { *m = DocumentRemove{} }
func (m *DocumentRemove) String() string            { return proto.CompactTextString(m) }
func (*DocumentRemove) ProtoMessage()               {}
func (*DocumentRemove) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *DocumentRemove) GetDocument() string {
	if m != nil {
		return m.Document
	}
	return ""
}

func (m *DocumentRemove) GetRemovedTargetIds() []int32 {
	if m != nil {
		return m.RemovedTargetIds
	}
	return nil
}

func (m *DocumentRemove) GetReadTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.ReadTime
	}
	return nil
}

// A digest of all the documents that match a given target.
type ExistenceFilter struct {
	// The target ID to which this filter applies.
	TargetId int32 `protobuf:"varint,1,opt,name=target_id,json=targetId" json:"target_id,omitempty"`
	// The total count of documents that match [target_id][google.firestore.v1beta1.ExistenceFilter.target_id].
	//
	// If different from the count of documents in the client that match, the
	// client must manually determine which documents no longer match the target.
	Count int32 `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *ExistenceFilter) Reset()                    { *m = ExistenceFilter{} }
func (m *ExistenceFilter) String() string            { return proto.CompactTextString(m) }
func (*ExistenceFilter) ProtoMessage()               {}
func (*ExistenceFilter) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *ExistenceFilter) GetTargetId() int32 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *ExistenceFilter) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Write)(nil), "google.firestore.v1beta1.Write")
	proto.RegisterType((*DocumentTransform)(nil), "google.firestore.v1beta1.DocumentTransform")
	proto.RegisterType((*DocumentTransform_FieldTransform)(nil), "google.firestore.v1beta1.DocumentTransform.FieldTransform")
	proto.RegisterType((*WriteResult)(nil), "google.firestore.v1beta1.WriteResult")
	proto.RegisterType((*DocumentChange)(nil), "google.firestore.v1beta1.DocumentChange")
	proto.RegisterType((*DocumentDelete)(nil), "google.firestore.v1beta1.DocumentDelete")
	proto.RegisterType((*DocumentRemove)(nil), "google.firestore.v1beta1.DocumentRemove")
	proto.RegisterType((*ExistenceFilter)(nil), "google.firestore.v1beta1.ExistenceFilter")
	proto.RegisterEnum("google.firestore.v1beta1.DocumentTransform_FieldTransform_ServerValue", DocumentTransform_FieldTransform_ServerValue_name, DocumentTransform_FieldTransform_ServerValue_value)
}

func init() { proto.RegisterFile("google/firestore/v1beta1/write.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0x8e, 0x93, 0x26, 0xaa, 0x4f, 0x50, 0xea, 0x3b, 0xb0, 0xb0, 0x42, 0x2f, 0x37, 0x8a, 0xf8,
	0x89, 0x04, 0x72, 0xd4, 0xb2, 0x40, 0xa2, 0x80, 0xd4, 0x24, 0x4e, 0x1b, 0xd1, 0xa2, 0x74, 0x92,
	0x06, 0x89, 0x8d, 0x35, 0x8d, 0x27, 0xae, 0x55, 0xdb, 0x63, 0xcd, 0x8c, 0x53, 0x78, 0x0d, 0x58,
	0xc0, 0x86, 0x0d, 0x4b, 0xde, 0x87, 0x87, 0x61, 0x87, 0x3c, 0xfe, 0x69, 0x4a, 0x15, 0x42, 0xab,
	0xbb, 0xf3, 0x39, 0xf3, 0x9d, 0xef, 0x7c, 0xe7, 0x4f, 0x86, 0x0f, 0x3d, 0xc6, 0xbc, 0x80, 0xf6,
	0x57, 0x3e, 0xa7, 0x42, 0x32, 0x4e, 0xfb, 0xeb, 0xa3, 0x1b, 0x2a, 0xc9, 0x51, 0xff, 0x9e, 0xfb,
	0x92, 0x5a, 0x31, 0x67, 0x92, 0x21, 0x33, 0x43, 0x59, 0x25, 0xca, 0xca, 0x51, 0xed, 0xc3, 0x3c,
	0x9e, 0xc4, 0x7e, 0x9f, 0x44, 0x11, 0x93, 0x44, 0xfa, 0x2c, 0x12, 0x59, 0x5c, 0xfb, 0xa3, 0xad,
	0xec, 0x4b, 0x16, 0x86, 0x2c, 0xca, 0x61, 0x9f, 0x6c, 0x85, 0xb9, 0x6c, 0x99, 0x84, 0x34, 0x92,
	0x39, 0xf0, 0x4d, 0x0e, 0x54, 0xd6, 0x4d, 0xb2, 0xea, 0x4b, 0x3f, 0xa4, 0x42, 0x92, 0x30, 0xce,
	0x00, 0xdd, 0xbf, 0xaa, 0x50, 0xff, 0x3e, 0x15, 0x8e, 0xbe, 0x82, 0x46, 0x12, 0xbb, 0x44, 0x52,
	0x53, 0xeb, 0x68, 0xbd, 0xe6, 0x71, 0xd7, 0xda, 0x56, 0x83, 0x35, 0xca, 0x93, 0x9c, 0x57, 0x70,
	0x1e, 0x83, 0x4c, 0x68, 0xb8, 0x34, 0xa0, 0x92, 0x9a, 0xd5, 0x8e, 0xd6, 0xd3, 0xd3, 0x97, 0xcc,
	0x46, 0xdf, 0x82, 0x2e, 0x39, 0x89, 0xc4, 0x8a, 0xf1, 0xd0, 0x6c, 0x28, 0xea, 0x4f, 0x77, 0x53,
	0xcf, 0x8b, 0x90, 0xf3, 0x0a, 0x7e, 0x88, 0x47, 0x67, 0xd0, 0xcc, 0x12, 0x3a, 0x21, 0x11, 0x77,
	0x66, 0x4d, 0xd1, 0x7d, 0xbc, 0x9b, 0xee, 0x92, 0x88, 0x3b, 0x0c, 0x59, 0x68, 0xfa, 0x8d, 0xae,
	0xc0, 0x58, 0x26, 0x9c, 0xd3, 0x48, 0x3a, 0x45, 0xcb, 0xcc, 0xbd, 0x5d, 0x6c, 0x53, 0x4e, 0x97,
	0x2c, 0x72, 0xfd, 0x74, 0x62, 0xf8, 0x20, 0x8f, 0x2f, 0x52, 0x0c, 0x9a, 0xa0, 0xb3, 0x98, 0x72,
	0x35, 0xcf, 0xee, 0xcf, 0x35, 0x78, 0xf5, 0xa4, 0x16, 0xd4, 0x86, 0xfd, 0x32, 0x5b, 0xda, 0x65,
	0x1d, 0x97, 0x36, 0xa2, 0x60, 0xac, 0x7c, 0x1a, 0xb8, 0x4e, 0x59, 0xad, 0x30, 0xab, 0x9d, 0x5a,
	0xaf, 0x79, 0xfc, 0xe5, 0x33, 0xda, 0x65, 0x8d, 0x53, 0x8e, 0xd2, 0xc4, 0x07, 0xab, 0x47, 0xb6,
	0x68, 0xff, 0xad, 0x41, 0xeb, 0x31, 0x06, 0xbd, 0x06, 0xc8, 0x32, 0xc7, 0x44, 0xde, 0xe6, 0xba,
	0x74, 0xe5, 0x99, 0x12, 0x79, 0x8b, 0xee, 0xe1, 0x5d, 0x41, 0xa5, 0x23, 0x99, 0x23, 0x28, 0x5f,
	0x53, 0xee, 0xac, 0x49, 0x90, 0x64, 0x73, 0x6e, 0x1d, 0x8f, 0x5f, 0xae, 0xcd, 0x9a, 0x29, 0xba,
	0x45, 0xca, 0x76, 0x5e, 0xc1, 0x86, 0xa0, 0x72, 0xce, 0x36, 0x7c, 0xdd, 0xaf, 0xa1, 0xb9, 0x61,
	0xa2, 0x43, 0x30, 0x67, 0x36, 0x5e, 0xd8, 0xd8, 0x59, 0x9c, 0x5e, 0x5c, 0xdb, 0xce, 0xf5, 0x77,
	0xb3, 0xa9, 0x3d, 0x9c, 0x8c, 0x27, 0xf6, 0xc8, 0xa8, 0x20, 0x03, 0xde, 0xc1, 0xf6, 0xd5, 0xb5,
	0x3d, 0x9b, 0x3b, 0xf3, 0xc9, 0xa5, 0x6d, 0x68, 0x03, 0x03, 0x5a, 0x65, 0x2b, 0x1d, 0xf9, 0x53,
	0x4c, 0xbb, 0xbf, 0x69, 0xd0, 0x54, 0xcb, 0x8e, 0xa9, 0x48, 0x02, 0x89, 0x4e, 0xca, 0x6d, 0x4a,
	0xcf, 0x22, 0xdf, 0xfb, 0x76, 0x51, 0x51, 0x71, 0x33, 0xd6, 0xbc, 0xb8, 0x99, 0x62, 0x83, 0x52,
	0x07, 0xba, 0x80, 0x57, 0x0f, 0xf4, 0x5c, 0x11, 0x16, 0x03, 0x7b, 0xb3, 0xbd, 0x29, 0xaa, 0x14,
	0x6c, 0x94, 0x91, 0x99, 0x12, 0xd1, 0xfd, 0x5d, 0x83, 0x56, 0xd1, 0xb0, 0xe1, 0x2d, 0x89, 0x3c,
	0x8a, 0xbe, 0xf9, 0xd7, 0xb2, 0xfc, 0xaf, 0x93, 0xdc, 0x58, 0xa8, 0xd7, 0x00, 0x92, 0x70, 0x8f,
	0x4a, 0xc7, 0x77, 0x85, 0x59, 0xef, 0xd4, 0x7a, 0x75, 0xac, 0x67, 0x9e, 0x89, 0x2b, 0xd0, 0x67,
	0x80, 0x38, 0x0d, 0xd9, 0x9a, 0xba, 0xce, 0x06, 0xac, 0xa1, 0x60, 0x46, 0xfe, 0x32, 0x2f, 0xd0,
	0xdd, 0x5f, 0x36, 0xf4, 0x8d, 0xb2, 0xc3, 0xfe, 0xaf, 0x65, 0x7e, 0x16, 0x39, 0xfa, 0x02, 0x74,
	0x4e, 0x89, 0x9b, 0x4d, 0x61, 0x6f, 0xe7, 0x14, 0xf6, 0x53, 0x70, 0x6a, 0x3e, 0x52, 0x85, 0x15,
	0xeb, 0x0b, 0x54, 0x55, 0xdf, 0xb6, 0xaa, 0x11, 0x1c, 0xd8, 0x3f, 0xfa, 0x42, 0xd2, 0x68, 0x49,
	0xc7, 0x7e, 0x20, 0x29, 0x47, 0xef, 0x83, 0x5e, 0x66, 0x54, 0xb2, 0xea, 0x78, 0xbf, 0x18, 0x05,
	0x7a, 0x0f, 0xea, 0x4b, 0x96, 0x44, 0x52, 0x9d, 0x54, 0x1d, 0x67, 0xc6, 0xe0, 0x57, 0x0d, 0x0e,
	0x97, 0x2c, 0xdc, 0x3a, 0xf2, 0x01, 0xa8, 0x55, 0x9e, 0xa6, 0x4a, 0xa6, 0xda, 0x0f, 0xa7, 0x39,
	0xce, 0x63, 0x01, 0x89, 0x3c, 0x8b, 0x71, 0xaf, 0xef, 0xd1, 0x48, 0xe9, 0xec, 0x67, 0x4f, 0x24,
	0xf6, 0xc5, 0xd3, 0x3f, 0xc6, 0x49, 0xe9, 0xf9, 0xa3, 0xba, 0x77, 0x36, 0x1c, 0xcf, 0xfe, 0xac,
	0x7e, 0x70, 0x96, 0x51, 0x0d, 0x03, 0x96, 0xb8, 0xd6, 0xb8, 0x4c, 0xbc, 0x38, 0x1a, 0xa4, 0x11,
	0x37, 0x0d, 0xc5, 0xfa, 0xf9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70, 0x36, 0x10, 0xb5, 0x0b,
	0x07, 0x00, 0x00,
}
