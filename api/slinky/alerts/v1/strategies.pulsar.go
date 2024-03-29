// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package alertsv1

import (
	_ "cosmossdk.io/api/amino"
	abci "cosmossdk.io/api/tendermint/abci"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ValidatorAlertIncentive              protoreflect.MessageDescriptor
	fd_ValidatorAlertIncentive_validator    protoreflect.FieldDescriptor
	fd_ValidatorAlertIncentive_alert_signer protoreflect.FieldDescriptor
	fd_ValidatorAlertIncentive_alert_height protoreflect.FieldDescriptor
)

func init() {
	file_slinky_alerts_v1_strategies_proto_init()
	md_ValidatorAlertIncentive = File_slinky_alerts_v1_strategies_proto.Messages().ByName("ValidatorAlertIncentive")
	fd_ValidatorAlertIncentive_validator = md_ValidatorAlertIncentive.Fields().ByName("validator")
	fd_ValidatorAlertIncentive_alert_signer = md_ValidatorAlertIncentive.Fields().ByName("alert_signer")
	fd_ValidatorAlertIncentive_alert_height = md_ValidatorAlertIncentive.Fields().ByName("alert_height")
}

var _ protoreflect.Message = (*fastReflection_ValidatorAlertIncentive)(nil)

type fastReflection_ValidatorAlertIncentive ValidatorAlertIncentive

func (x *ValidatorAlertIncentive) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ValidatorAlertIncentive)(x)
}

func (x *ValidatorAlertIncentive) slowProtoReflect() protoreflect.Message {
	mi := &file_slinky_alerts_v1_strategies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ValidatorAlertIncentive_messageType fastReflection_ValidatorAlertIncentive_messageType
var _ protoreflect.MessageType = fastReflection_ValidatorAlertIncentive_messageType{}

type fastReflection_ValidatorAlertIncentive_messageType struct{}

func (x fastReflection_ValidatorAlertIncentive_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ValidatorAlertIncentive)(nil)
}
func (x fastReflection_ValidatorAlertIncentive_messageType) New() protoreflect.Message {
	return new(fastReflection_ValidatorAlertIncentive)
}
func (x fastReflection_ValidatorAlertIncentive_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorAlertIncentive
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ValidatorAlertIncentive) Descriptor() protoreflect.MessageDescriptor {
	return md_ValidatorAlertIncentive
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ValidatorAlertIncentive) Type() protoreflect.MessageType {
	return _fastReflection_ValidatorAlertIncentive_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ValidatorAlertIncentive) New() protoreflect.Message {
	return new(fastReflection_ValidatorAlertIncentive)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ValidatorAlertIncentive) Interface() protoreflect.ProtoMessage {
	return (*ValidatorAlertIncentive)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ValidatorAlertIncentive) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Validator != nil {
		value := protoreflect.ValueOfMessage(x.Validator.ProtoReflect())
		if !f(fd_ValidatorAlertIncentive_validator, value) {
			return
		}
	}
	if x.AlertSigner != "" {
		value := protoreflect.ValueOfString(x.AlertSigner)
		if !f(fd_ValidatorAlertIncentive_alert_signer, value) {
			return
		}
	}
	if x.AlertHeight != uint64(0) {
		value := protoreflect.ValueOfUint64(x.AlertHeight)
		if !f(fd_ValidatorAlertIncentive_alert_height, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ValidatorAlertIncentive) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "slinky.alerts.v1.ValidatorAlertIncentive.validator":
		return x.Validator != nil
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_signer":
		return x.AlertSigner != ""
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_height":
		return x.AlertHeight != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: slinky.alerts.v1.ValidatorAlertIncentive"))
		}
		panic(fmt.Errorf("message slinky.alerts.v1.ValidatorAlertIncentive does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAlertIncentive) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "slinky.alerts.v1.ValidatorAlertIncentive.validator":
		x.Validator = nil
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_signer":
		x.AlertSigner = ""
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_height":
		x.AlertHeight = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: slinky.alerts.v1.ValidatorAlertIncentive"))
		}
		panic(fmt.Errorf("message slinky.alerts.v1.ValidatorAlertIncentive does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ValidatorAlertIncentive) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "slinky.alerts.v1.ValidatorAlertIncentive.validator":
		value := x.Validator
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_signer":
		value := x.AlertSigner
		return protoreflect.ValueOfString(value)
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_height":
		value := x.AlertHeight
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: slinky.alerts.v1.ValidatorAlertIncentive"))
		}
		panic(fmt.Errorf("message slinky.alerts.v1.ValidatorAlertIncentive does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAlertIncentive) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "slinky.alerts.v1.ValidatorAlertIncentive.validator":
		x.Validator = value.Message().Interface().(*abci.Validator)
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_signer":
		x.AlertSigner = value.Interface().(string)
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_height":
		x.AlertHeight = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: slinky.alerts.v1.ValidatorAlertIncentive"))
		}
		panic(fmt.Errorf("message slinky.alerts.v1.ValidatorAlertIncentive does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAlertIncentive) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "slinky.alerts.v1.ValidatorAlertIncentive.validator":
		if x.Validator == nil {
			x.Validator = new(abci.Validator)
		}
		return protoreflect.ValueOfMessage(x.Validator.ProtoReflect())
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_signer":
		panic(fmt.Errorf("field alert_signer of message slinky.alerts.v1.ValidatorAlertIncentive is not mutable"))
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_height":
		panic(fmt.Errorf("field alert_height of message slinky.alerts.v1.ValidatorAlertIncentive is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: slinky.alerts.v1.ValidatorAlertIncentive"))
		}
		panic(fmt.Errorf("message slinky.alerts.v1.ValidatorAlertIncentive does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ValidatorAlertIncentive) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "slinky.alerts.v1.ValidatorAlertIncentive.validator":
		m := new(abci.Validator)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_signer":
		return protoreflect.ValueOfString("")
	case "slinky.alerts.v1.ValidatorAlertIncentive.alert_height":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: slinky.alerts.v1.ValidatorAlertIncentive"))
		}
		panic(fmt.Errorf("message slinky.alerts.v1.ValidatorAlertIncentive does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ValidatorAlertIncentive) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in slinky.alerts.v1.ValidatorAlertIncentive", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ValidatorAlertIncentive) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ValidatorAlertIncentive) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ValidatorAlertIncentive) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ValidatorAlertIncentive) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ValidatorAlertIncentive)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Validator != nil {
			l = options.Size(x.Validator)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.AlertSigner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AlertHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.AlertHeight))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorAlertIncentive)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.AlertHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.AlertHeight))
			i--
			dAtA[i] = 0x18
		}
		if len(x.AlertSigner) > 0 {
			i -= len(x.AlertSigner)
			copy(dAtA[i:], x.AlertSigner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AlertSigner)))
			i--
			dAtA[i] = 0x12
		}
		if x.Validator != nil {
			encoded, err := options.Marshal(x.Validator)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ValidatorAlertIncentive)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorAlertIncentive: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ValidatorAlertIncentive: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Validator == nil {
					x.Validator = &abci.Validator{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Validator); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AlertSigner", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AlertSigner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AlertHeight", wireType)
				}
				x.AlertHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.AlertHeight |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: slinky/alerts/v1/strategies.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ValidatorAlertIncentive defines the incentive strategy to be executed for a
// validator that has been confirmed to have at fault for an x/alerts alert.
// This strategy is expected to slash half of the validator's stake.
type ValidatorAlertIncentive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The validator that has been confirmed to have been at fault for an alert.
	Validator *abci.Validator `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	// AlertSigner is the signer of the alert referenced by the conclusion that
	// created this incentive.
	AlertSigner string `protobuf:"bytes,2,opt,name=alert_signer,json=alertSigner,proto3" json:"alert_signer,omitempty"`
	// AlertHeight is the height at which the infraction occurred
	AlertHeight uint64 `protobuf:"varint,3,opt,name=alert_height,json=alertHeight,proto3" json:"alert_height,omitempty"`
}

func (x *ValidatorAlertIncentive) Reset() {
	*x = ValidatorAlertIncentive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slinky_alerts_v1_strategies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorAlertIncentive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorAlertIncentive) ProtoMessage() {}

// Deprecated: Use ValidatorAlertIncentive.ProtoReflect.Descriptor instead.
func (*ValidatorAlertIncentive) Descriptor() ([]byte, []int) {
	return file_slinky_alerts_v1_strategies_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorAlertIncentive) GetValidator() *abci.Validator {
	if x != nil {
		return x.Validator
	}
	return nil
}

func (x *ValidatorAlertIncentive) GetAlertSigner() string {
	if x != nil {
		return x.AlertSigner
	}
	return ""
}

func (x *ValidatorAlertIncentive) GetAlertHeight() uint64 {
	if x != nil {
		return x.AlertHeight
	}
	return 0
}

var File_slinky_alerts_v1_strategies_proto protoreflect.FileDescriptor

var file_slinky_alerts_v1_strategies_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x79, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x79, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69,
	0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x61, 0x62, 0x63, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x02, 0x0a, 0x17, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2e, 0x61, 0x62, 0x63, 0x69, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xd2, 0xb4, 0x2d, 0x14, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x52, 0x0b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x3a, 0x4e, 0xca, 0xb4, 0x2d, 0x1e, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x79, 0x2e, 0x69,
	0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x63,
	0x65, 0x6e, 0x74, 0x69, 0x76, 0x65, 0x8a, 0xe7, 0xb0, 0x2a, 0x27, 0x73, 0x6c, 0x69, 0x6e, 0x6b,
	0x79, 0x2f, 0x78, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x69,
	0x76, 0x65, 0x42, 0xb5, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6c, 0x69, 0x6e, 0x6b,
	0x79, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x6c, 0x69, 0x6e, 0x6b, 0x79, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x41, 0x58,
	0xaa, 0x02, 0x10, 0x53, 0x6c, 0x69, 0x6e, 0x6b, 0x79, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x53, 0x6c, 0x69, 0x6e, 0x6b, 0x79, 0x5c, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x53, 0x6c, 0x69, 0x6e, 0x6b, 0x79, 0x5c,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x53, 0x6c, 0x69, 0x6e, 0x6b, 0x79, 0x3a, 0x3a,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_slinky_alerts_v1_strategies_proto_rawDescOnce sync.Once
	file_slinky_alerts_v1_strategies_proto_rawDescData = file_slinky_alerts_v1_strategies_proto_rawDesc
)

func file_slinky_alerts_v1_strategies_proto_rawDescGZIP() []byte {
	file_slinky_alerts_v1_strategies_proto_rawDescOnce.Do(func() {
		file_slinky_alerts_v1_strategies_proto_rawDescData = protoimpl.X.CompressGZIP(file_slinky_alerts_v1_strategies_proto_rawDescData)
	})
	return file_slinky_alerts_v1_strategies_proto_rawDescData
}

var file_slinky_alerts_v1_strategies_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_slinky_alerts_v1_strategies_proto_goTypes = []interface{}{
	(*ValidatorAlertIncentive)(nil), // 0: slinky.alerts.v1.ValidatorAlertIncentive
	(*abci.Validator)(nil),          // 1: tendermint.abci.Validator
}
var file_slinky_alerts_v1_strategies_proto_depIdxs = []int32{
	1, // 0: slinky.alerts.v1.ValidatorAlertIncentive.validator:type_name -> tendermint.abci.Validator
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_slinky_alerts_v1_strategies_proto_init() }
func file_slinky_alerts_v1_strategies_proto_init() {
	if File_slinky_alerts_v1_strategies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slinky_alerts_v1_strategies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorAlertIncentive); i {
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
			RawDescriptor: file_slinky_alerts_v1_strategies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_slinky_alerts_v1_strategies_proto_goTypes,
		DependencyIndexes: file_slinky_alerts_v1_strategies_proto_depIdxs,
		MessageInfos:      file_slinky_alerts_v1_strategies_proto_msgTypes,
	}.Build()
	File_slinky_alerts_v1_strategies_proto = out.File
	file_slinky_alerts_v1_strategies_proto_rawDesc = nil
	file_slinky_alerts_v1_strategies_proto_goTypes = nil
	file_slinky_alerts_v1_strategies_proto_depIdxs = nil
}
