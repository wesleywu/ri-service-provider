package codec

import (
	"encoding/json"
	"reflect"

	"github.com/wesleywu/gowing/protobuf/gwtypes"
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
	s1 := wrapperspb.StringValue{}
	_ = types.RegisterMessage(s1.ProtoReflect().Type())
	b1 := wrapperspb.BoolValue{}
	_ = types.RegisterMessage(b1.ProtoReflect().Type())
	i1 := wrapperspb.Int32Value{}
	_ = types.RegisterMessage(i1.ProtoReflect().Type())
	s := &gwtypes.StringSlice{}
	_ = types.RegisterMessage(s.ProtoReflect().Type())
	b := &gwtypes.BoolSlice{}
	_ = types.RegisterMessage(b.ProtoReflect().Type())
	c := &gwtypes.Condition{}
	_ = types.RegisterMessage(c.ProtoReflect().Type())
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
