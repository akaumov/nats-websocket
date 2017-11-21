// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/packets.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/packets.proto

It has these top-level messages:
	InputMessage
	SendMessage
	SendMessageToAllUserDevices
	SendMessageToDevice
	Response
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InputMessage_Type int32

const (
	InputMessage_TEXT   InputMessage_Type = 0
	InputMessage_BINARY InputMessage_Type = 1
)

var InputMessage_Type_name = map[int32]string{
	0: "TEXT",
	1: "BINARY",
}
var InputMessage_Type_value = map[string]int32{
	"TEXT":   0,
	"BINARY": 1,
}

func (x InputMessage_Type) String() string {
	return proto.EnumName(InputMessage_Type_name, int32(x))
}
func (InputMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type InputMessage struct {
	Type      InputMessage_Type `protobuf:"varint,1,opt,name=type,enum=pb.InputMessage_Type" json:"type,omitempty"`
	InputTime int64             `protobuf:"varint,2,opt,name=inputTime" json:"inputTime,omitempty"`
	UserId    string            `protobuf:"bytes,3,opt,name=userId" json:"userId,omitempty"`
	DeviceId  string            `protobuf:"bytes,4,opt,name=deviceId" json:"deviceId,omitempty"`
	Body      []byte            `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *InputMessage) Reset()                    { *m = InputMessage{} }
func (m *InputMessage) String() string            { return proto.CompactTextString(m) }
func (*InputMessage) ProtoMessage()               {}
func (*InputMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InputMessage) GetType() InputMessage_Type {
	if m != nil {
		return m.Type
	}
	return InputMessage_TEXT
}

func (m *InputMessage) GetInputTime() int64 {
	if m != nil {
		return m.InputTime
	}
	return 0
}

func (m *InputMessage) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *InputMessage) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *InputMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type SendMessage struct {
	// Types that are valid to be assigned to MessageType:
	//	*SendMessage_SenToAllUserDevices
	//	*SendMessage_SenToDevice
	MessageType isSendMessage_MessageType `protobuf_oneof:"message_type"`
}

func (m *SendMessage) Reset()                    { *m = SendMessage{} }
func (m *SendMessage) String() string            { return proto.CompactTextString(m) }
func (*SendMessage) ProtoMessage()               {}
func (*SendMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isSendMessage_MessageType interface {
	isSendMessage_MessageType()
}

type SendMessage_SenToAllUserDevices struct {
	SenToAllUserDevices *SendMessageToAllUserDevices `protobuf:"bytes,1,opt,name=senToAllUserDevices,oneof"`
}
type SendMessage_SenToDevice struct {
	SenToDevice *SendMessageToDevice `protobuf:"bytes,2,opt,name=senToDevice,oneof"`
}

func (*SendMessage_SenToAllUserDevices) isSendMessage_MessageType() {}
func (*SendMessage_SenToDevice) isSendMessage_MessageType()         {}

func (m *SendMessage) GetMessageType() isSendMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *SendMessage) GetSenToAllUserDevices() *SendMessageToAllUserDevices {
	if x, ok := m.GetMessageType().(*SendMessage_SenToAllUserDevices); ok {
		return x.SenToAllUserDevices
	}
	return nil
}

func (m *SendMessage) GetSenToDevice() *SendMessageToDevice {
	if x, ok := m.GetMessageType().(*SendMessage_SenToDevice); ok {
		return x.SenToDevice
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SendMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SendMessage_OneofMarshaler, _SendMessage_OneofUnmarshaler, _SendMessage_OneofSizer, []interface{}{
		(*SendMessage_SenToAllUserDevices)(nil),
		(*SendMessage_SenToDevice)(nil),
	}
}

func _SendMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SendMessage)
	// message_type
	switch x := m.MessageType.(type) {
	case *SendMessage_SenToAllUserDevices:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SenToAllUserDevices); err != nil {
			return err
		}
	case *SendMessage_SenToDevice:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SenToDevice); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SendMessage.MessageType has unexpected type %T", x)
	}
	return nil
}

func _SendMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SendMessage)
	switch tag {
	case 1: // message_type.senToAllUserDevices
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SendMessageToAllUserDevices)
		err := b.DecodeMessage(msg)
		m.MessageType = &SendMessage_SenToAllUserDevices{msg}
		return true, err
	case 2: // message_type.senToDevice
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SendMessageToDevice)
		err := b.DecodeMessage(msg)
		m.MessageType = &SendMessage_SenToDevice{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SendMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SendMessage)
	// message_type
	switch x := m.MessageType.(type) {
	case *SendMessage_SenToAllUserDevices:
		s := proto.Size(x.SenToAllUserDevices)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SendMessage_SenToDevice:
		s := proto.Size(x.SenToDevice)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SendMessageToAllUserDevices struct {
	UserId  string   `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	Exclude []string `protobuf:"bytes,2,rep,name=exclude" json:"exclude,omitempty"`
}

func (m *SendMessageToAllUserDevices) Reset()                    { *m = SendMessageToAllUserDevices{} }
func (m *SendMessageToAllUserDevices) String() string            { return proto.CompactTextString(m) }
func (*SendMessageToAllUserDevices) ProtoMessage()               {}
func (*SendMessageToAllUserDevices) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SendMessageToAllUserDevices) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SendMessageToAllUserDevices) GetExclude() []string {
	if m != nil {
		return m.Exclude
	}
	return nil
}

type SendMessageToDevice struct {
	DeviceId string `protobuf:"bytes,2,opt,name=deviceId" json:"deviceId,omitempty"`
}

func (m *SendMessageToDevice) Reset()                    { *m = SendMessageToDevice{} }
func (m *SendMessageToDevice) String() string            { return proto.CompactTextString(m) }
func (*SendMessageToDevice) ProtoMessage()               {}
func (*SendMessageToDevice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SendMessageToDevice) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

type Response struct {
	Status int64  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Response) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*InputMessage)(nil), "pb.InputMessage")
	proto.RegisterType((*SendMessage)(nil), "pb.SendMessage")
	proto.RegisterType((*SendMessageToAllUserDevices)(nil), "pb.SendMessageToAllUserDevices")
	proto.RegisterType((*SendMessageToDevice)(nil), "pb.SendMessageToDevice")
	proto.RegisterType((*Response)(nil), "pb.Response")
	proto.RegisterEnum("pb.InputMessage_Type", InputMessage_Type_name, InputMessage_Type_value)
}

func init() { proto.RegisterFile("pb/packets.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xed, 0x24, 0xf9, 0xfa, 0xb5, 0x37, 0xa5, 0x94, 0x29, 0x6a, 0xd0, 0x82, 0x21, 0xab, 0xb8,
	0x89, 0x58, 0xc1, 0x8d, 0xab, 0x16, 0x85, 0x66, 0xa1, 0xc2, 0x34, 0x82, 0xae, 0xa4, 0xe9, 0x5c,
	0xa4, 0xd8, 0x26, 0x43, 0x67, 0x22, 0xf6, 0xa9, 0x7c, 0x02, 0xdf, 0x4d, 0x32, 0xe9, 0x4f, 0xa2,
	0xc5, 0x5d, 0xce, 0x3d, 0x27, 0xf7, 0xe7, 0x9c, 0x81, 0x8e, 0x88, 0xcf, 0xc5, 0x64, 0xfa, 0x86,
	0x4a, 0x06, 0x62, 0x99, 0xaa, 0x94, 0x1a, 0x22, 0xf6, 0xbe, 0x08, 0xb4, 0xc2, 0x44, 0x64, 0xea,
	0x0e, 0xa5, 0x9c, 0xbc, 0x22, 0x3d, 0x03, 0x4b, 0xad, 0x04, 0x3a, 0xc4, 0x25, 0x7e, 0xbb, 0x7f,
	0x10, 0x88, 0x38, 0x28, 0xf3, 0x41, 0xb4, 0x12, 0xc8, 0xb4, 0x84, 0xf6, 0xa0, 0x39, 0xcb, 0xa9,
	0x68, 0xb6, 0x40, 0xc7, 0x70, 0x89, 0x6f, 0xb2, 0x5d, 0x81, 0x1e, 0x42, 0x3d, 0x93, 0xb8, 0x0c,
	0xb9, 0x63, 0xba, 0xc4, 0x6f, 0xb2, 0x35, 0xa2, 0xc7, 0xd0, 0xe0, 0xf8, 0x3e, 0x9b, 0x62, 0xc8,
	0x1d, 0x4b, 0x33, 0x5b, 0x4c, 0x29, 0x58, 0x71, 0xca, 0x57, 0xce, 0x3f, 0x97, 0xf8, 0x2d, 0xa6,
	0xbf, 0xbd, 0x1e, 0x58, 0xf9, 0x4c, 0xda, 0x00, 0x2b, 0xba, 0x7d, 0x8a, 0x3a, 0x35, 0x0a, 0x50,
	0x1f, 0x86, 0xf7, 0x03, 0xf6, 0xdc, 0x21, 0xde, 0x27, 0x01, 0x7b, 0x8c, 0x09, 0xdf, 0xac, 0x3f,
	0x86, 0xae, 0xc4, 0x24, 0x4a, 0x07, 0xf3, 0xf9, 0xa3, 0xc4, 0xe5, 0x8d, 0xee, 0x2c, 0xf5, 0x35,
	0x76, 0xff, 0x34, 0xbf, 0xa6, 0xa4, 0xfe, 0x29, 0x1b, 0xd5, 0xd8, 0xbe, 0xbf, 0xe9, 0x35, 0xd8,
	0xba, 0x5c, 0x60, 0x7d, 0xaa, 0xdd, 0x3f, 0xfa, 0xd5, 0xac, 0xa0, 0x47, 0x35, 0x56, 0x56, 0x0f,
	0xdb, 0xd0, 0x5a, 0x14, 0x8a, 0x97, 0xdc, 0x35, 0xef, 0x01, 0x4e, 0xfe, 0x58, 0xa1, 0x64, 0x1b,
	0xa9, 0xd8, 0xe6, 0xc0, 0x7f, 0xfc, 0x98, 0xce, 0x33, 0x9e, 0xcf, 0x37, 0xfd, 0x26, 0xdb, 0x40,
	0xef, 0x02, 0xba, 0x7b, 0xd6, 0xa8, 0xf8, 0x6c, 0x54, 0x7d, 0xf6, 0xae, 0xa0, 0xc1, 0x50, 0x8a,
	0x34, 0x91, 0x3a, 0x27, 0xa9, 0x26, 0x2a, 0x2b, 0x4c, 0x32, 0xd9, 0x1a, 0x6d, 0xb3, 0x30, 0x76,
	0x59, 0xc4, 0x75, 0xfd, 0x70, 0x2e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xa2, 0xed, 0xf2,
	0x4c, 0x02, 0x00, 0x00,
}
