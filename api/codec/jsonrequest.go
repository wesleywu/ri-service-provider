package codec

import (
	"encoding/json"
	"reflect"

	goguruTypes "github.com/castbox/go-guru/pkg/goguru/types"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Name is the name registered for the json codec.
const Name = "json"

var (
	// MarshalOptions is a configurable JSON format marshaller.
	MarshalOptions protojson.MarshalOptions
	// UnmarshalOptions is a configurable JSON format parser.
	UnmarshalOptions protojson.UnmarshalOptions
)

func init() {
	types := &protoregistry.Types{}
	// register types from wrapperspb package
	doubleValue := wrapperspb.DoubleValue{}
	_ = types.RegisterMessage(doubleValue.ProtoReflect().Type())
	floatValue := wrapperspb.FloatValue{}
	_ = types.RegisterMessage(floatValue.ProtoReflect().Type())
	int64Value := wrapperspb.Int64Value{}
	_ = types.RegisterMessage(int64Value.ProtoReflect().Type())
	uint64Value := wrapperspb.UInt64Value{}
	_ = types.RegisterMessage(uint64Value.ProtoReflect().Type())
	int32Value := wrapperspb.Int32Value{}
	_ = types.RegisterMessage(int32Value.ProtoReflect().Type())
	uint32Value := wrapperspb.UInt32Value{}
	_ = types.RegisterMessage(uint32Value.ProtoReflect().Type())
	boolValue := wrapperspb.BoolValue{}
	_ = types.RegisterMessage(boolValue.ProtoReflect().Type())
	stringValue := wrapperspb.StringValue{}
	_ = types.RegisterMessage(stringValue.ProtoReflect().Type())
	bytesValue := wrapperspb.BytesValue{}
	_ = types.RegisterMessage(bytesValue.ProtoReflect().Type())

	// register types from goguruTypes package
	doubleSlice := goguruTypes.DoubleSlice{}
	_ = types.RegisterMessage(doubleSlice.ProtoReflect().Type())
	floatSlice := goguruTypes.FloatSlice{}
	_ = types.RegisterMessage(floatSlice.ProtoReflect().Type())
	int64Slice := goguruTypes.Int64Slice{}
	_ = types.RegisterMessage(int64Slice.ProtoReflect().Type())
	uint64Slice := goguruTypes.UInt64Slice{}
	_ = types.RegisterMessage(uint64Slice.ProtoReflect().Type())
	int32Slice := goguruTypes.Int32Slice{}
	_ = types.RegisterMessage(int32Slice.ProtoReflect().Type())
	uint32Slice := goguruTypes.UInt32Slice{}
	_ = types.RegisterMessage(uint32Slice.ProtoReflect().Type())
	boolSlice := goguruTypes.BoolSlice{}
	_ = types.RegisterMessage(boolSlice.ProtoReflect().Type())
	stringSlice := goguruTypes.StringSlice{}
	_ = types.RegisterMessage(stringSlice.ProtoReflect().Type())
	timestampSlice := goguruTypes.TimestampSlice{}
	_ = types.RegisterMessage(timestampSlice.ProtoReflect().Type())
	condition := goguruTypes.Condition{}
	_ = types.RegisterMessage(condition.ProtoReflect().Type())

	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
		Resolver:       types,
	}
	MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
		Resolver:        types,
	}
}

// JsonCodec is a Codec implementation with json.
type JsonCodec struct{}

func (JsonCodec) Marshal(v interface{}) ([]byte, error) {
	switch m := v.(type) {
	case json.Marshaler:
		return m.MarshalJSON()
	case proto.Message:
		return MarshalOptions.Marshal(m)
	default:
		return json.Marshal(m)
	}
}

func (JsonCodec) Unmarshal(data []byte, v interface{}) error {
	switch m := v.(type) {
	case json.Unmarshaler:
		return m.UnmarshalJSON(data)
	case proto.Message:
		return UnmarshalOptions.Unmarshal(data, m)
	default:
		rv := reflect.ValueOf(v)
		for rv := rv; rv.Kind() == reflect.Ptr; {
			if rv.IsNil() {
				rv.Set(reflect.New(rv.Type().Elem()))
			}
			rv = rv.Elem()
		}
		if m, ok := reflect.Indirect(rv).Interface().(proto.Message); ok {
			return UnmarshalOptions.Unmarshal(data, m)
		}
		return json.Unmarshal(data, m)
	}
}

func (JsonCodec) Name() string {
	return Name
}
