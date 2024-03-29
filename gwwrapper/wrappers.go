package gwwrapper

import (
	"context"
	"reflect"
	"strings"
	"time"

	goguruTypes "github.com/castbox/go-guru/pkg/goguru/types"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	ConditionQueryPrefix = "condition{"
	ConditionQuerySuffix = "}"
	TagNameOperator      = "operator"
	TagNameMulti         = "multi"
	TagNameWildcard      = "wildcard"

	FieldNameOperator = "operator"
	FieldNameMulti    = "multi"
	FieldNameWildcard = "wildcard"
	FieldNameValue    = "value"
)

func WrapDouble(v interface{}) *float64 {
	if v == nil {
		return nil
	}
	if r, ok := v.(float64); ok {
		return &r
	}
	r := gconv.Float64(v)
	return &r
}

func WrapFloat(v interface{}) *float32 {
	if v == nil {
		return nil
	}
	if r, ok := v.(float32); ok {
		return &r
	}
	r := gconv.Float32(v)
	return &r
}

func WrapInt64(v interface{}) *int64 {
	if v == nil {
		return nil
	}
	if r, ok := v.(int64); ok {
		return &r
	}
	r := gconv.Int64(v)
	return &r
}

func WrapUInt64(v interface{}) *uint64 {
	if v == nil {
		return nil
	}
	if r, ok := v.(uint64); ok {
		return &r
	}
	r := gconv.Uint64(v)
	return &r
}

func WrapInt32(v interface{}) *int32 {
	if v == nil {
		return nil
	}
	if r, ok := v.(int32); ok {
		return &r
	}
	r := gconv.Int32(v)
	return &r
}

func WrapUInt32(v interface{}) *uint32 {
	if v == nil {
		return nil
	}
	if r, ok := v.(uint32); ok {
		return &r
	}
	r := gconv.Uint32(v)
	return &r
}

func WrapBool(v interface{}) *bool {
	if v == nil {
		return nil
	}
	if r, ok := v.(bool); ok {
		return &r
	}
	r := gconv.Bool(v)
	return &r
}

func WrapString(v interface{}) *string {
	if v == nil {
		return nil
	}
	if r, ok := v.(string); ok {
		return &r
	}
	r := gconv.String(v)
	return &r
}

func WrapTimestamp(v interface{}) *timestamppb.Timestamp {
	if v == nil {
		return nil
	}
	switch r := v.(type) {
	case time.Time:
		return timestamppb.New(r)
	case *time.Time:
		return timestamppb.New(*r)
	case gtime.Time:
		return timestamppb.New(r.Time)
	case *gtime.Time:
		return timestamppb.New(r.Time)
	default:
		return timestamppb.New(gtime.New(v).Time)
	}
}

func AnyDouble(v float64) *anypb.Any {
	valueAny := &wrapperspb.DoubleValue{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyFloat(v float32) *anypb.Any {
	valueAny := &wrapperspb.FloatValue{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyInt64(v int64) *anypb.Any {
	valueAny := &wrapperspb.Int64Value{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyUInt64(v uint64) *anypb.Any {
	valueAny := &wrapperspb.UInt64Value{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyInt(v int) *anypb.Any {
	valueAny := &wrapperspb.Int32Value{Value: int32(v)}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyUInt(v uint) *anypb.Any {
	valueAny := &wrapperspb.UInt32Value{Value: uint32(v)}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyInt32(v int32) *anypb.Any {
	valueAny := &wrapperspb.Int32Value{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyUInt32(v uint32) *anypb.Any {
	valueAny := &wrapperspb.UInt32Value{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyBool(v bool) *anypb.Any {
	valueAny := &wrapperspb.BoolValue{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyString(v string) *anypb.Any {
	valueAny := &wrapperspb.StringValue{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyDoubleSlice(v []float64) *anypb.Any {
	valueAny := &goguruTypes.DoubleSlice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyFloatSlice(v []float32) *anypb.Any {
	valueAny := &goguruTypes.FloatSlice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyIntSlice(v []int) *anypb.Any {
	value := make([]int32, len(v))
	for i, v := range v {
		value[i] = int32(v)
	}
	return AnyInt32Slice(value)
}

func AnyUIntSlice(v []uint) *anypb.Any {
	value := make([]uint32, len(v))
	for i, v := range v {
		value[i] = uint32(v)
	}
	return AnyUInt32Slice(value)
}

func AnyInt64Slice(v []int64) *anypb.Any {
	valueAny := &goguruTypes.Int64Slice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyUInt64Slice(v []uint64) *anypb.Any {
	valueAny := &goguruTypes.UInt64Slice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyInt32Slice(v []int32) *anypb.Any {
	valueAny := &goguruTypes.Int32Slice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyUInt32Slice(v []uint32) *anypb.Any {
	valueAny := &goguruTypes.UInt32Slice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyBoolSlice(v []bool) *anypb.Any {
	valueAny := &goguruTypes.BoolSlice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyStringSlice(v []string) *anypb.Any {
	valueAny := &goguruTypes.StringSlice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyCondition(operator goguruTypes.OperatorType, multi goguruTypes.MultiType, wildcard goguruTypes.WildcardType, value *anypb.Any) *anypb.Any {
	valueCondition := &goguruTypes.Condition{
		Operator: operator,
		Multi:    multi,
		Wildcard: wildcard,
		Value:    value,
	}
	result, _ := anypb.New(valueCondition)
	return result
}

func AnyInterface(ctx context.Context, any interface{}, tag reflect.StructTag) *anypb.Any {
	if any == nil {
		return nil
	}
	switch value := any.(type) {
	case string:
		return AnyString(value)
	case int:
		return AnyInt(value)
	case int8, int16, uint8, uint16:
		g.Log().Warning(ctx, "Cannot convert int8,int16,uint8,uint16 types to *anypb.Any")
		return nil
	case int32:
		return AnyInt32(value)
	case int64:
		return AnyInt64(value)
	case uint:
		return AnyUInt(value)
	case uint32:
		return AnyUInt32(value)
	case uint64:
		return AnyUInt64(value)
	case bool:
		return AnyBool(value)
	case float32:
		return AnyFloat(value)
	case float64:
		return AnyDouble(value)
	default:
		t := reflect.TypeOf(any)
		switch t.Kind() {
		case reflect.Slice, reflect.Array:
			return AnySlice(ctx, any, tag)
		case reflect.Ptr:
			return AnyInterface(ctx, reflect.ValueOf(any).Elem(), tag)
		default:
			g.Log().Warningf(ctx, "Cannot convert type %s to *anypb.Any", t.String())
			return nil
		}
	}
}

func AnySlice(ctx context.Context, any interface{}, tag reflect.StructTag) *anypb.Any {
	switch value := any.(type) {
	case []string:
		return AnyStringSlice(value)
	case []int:
		return AnyIntSlice(value)
	case []int8, int16, []uint8, []uint16:
		g.Log().Warning(ctx, "Cannot convert int8,int16,uint8,uint16 types to *anypb.Any")
		return nil
	case []int32:
		return AnyInt32Slice(value)
	case []int64:
		return AnyInt64Slice(value)
	case []uint:
		return AnyUIntSlice(value)
	case []uint32:
		return AnyUInt32Slice(value)
	case []uint64:
		return AnyUInt64Slice(value)
	case []bool:
		return AnyBoolSlice(value)
	case []float32:
		return AnyFloatSlice(value)
	case []float64:
		return AnyDoubleSlice(value)
	default:
		g.Log().Warningf(ctx, "Cannot convert %s types to *anypb.Any", reflect.TypeOf(any).String())
		return nil
	}
}

func AnySliceCondition(ctx context.Context, any interface{}, tag reflect.StructTag) *anypb.Any {
	operator := goguruTypes.MustParseOperatorType(tag.Get(TagNameOperator), goguruTypes.OperatorType_EQ)
	multi := goguruTypes.MustParseMultiType(tag.Get(TagNameMulti), goguruTypes.MultiType_In)
	wildcard := goguruTypes.MustParseWildcardType(tag.Get(TagNameWildcard), goguruTypes.WildcardType_NoWildcard)
	switch value := any.(type) {
	case []interface{}:
		array := make([]string, len(value))
		for k, v := range value {
			array[k] = gconv.String(v)
		}
		return AnyCondition(operator, multi, wildcard, AnyStringSlice(array))
	case []string:
		return AnyCondition(operator, multi, wildcard, AnyStringSlice(value))
	case []int:
		return AnyCondition(operator, multi, wildcard, AnyIntSlice(value))
	case []int8, []int16, []uint8, []uint16:
		g.Log().Warning(ctx, "Cannot convert []int8,[]int16,[]uint8,[]uint16 types to *anypb.Any")
		return nil
	case []int32:
		return AnyCondition(operator, multi, wildcard, AnyInt32Slice(value))
	case []int64:
		return AnyCondition(operator, multi, wildcard, AnyInt64Slice(value))
	case []uint:
		return AnyCondition(operator, multi, wildcard, AnyUIntSlice(value))
	case []uint32:
		return AnyCondition(operator, multi, wildcard, AnyUInt32Slice(value))
	case []uint64:
		return AnyCondition(operator, multi, wildcard, AnyUInt64Slice(value))
	case []bool:
		return AnyCondition(operator, multi, wildcard, AnyBoolSlice(value))
	case []float32:
		return AnyCondition(operator, multi, wildcard, AnyFloatSlice(value))
	case []float64:
		return AnyCondition(operator, multi, wildcard, AnyDoubleSlice(value))
	default:
		g.Log().Warningf(ctx, "Cannot convert %s types to *anypb.Any", reflect.TypeOf(any).String())
		return nil
	}
}

func AnyInterfaceCondition(ctx context.Context, any interface{}, tag reflect.StructTag) *anypb.Any {
	operator := goguruTypes.MustParseOperatorType(tag.Get(TagNameOperator), goguruTypes.OperatorType_EQ)
	multi := goguruTypes.MustParseMultiType(tag.Get(TagNameMulti), goguruTypes.MultiType_NoMulti)
	wildcard := goguruTypes.MustParseWildcardType(tag.Get(TagNameWildcard), goguruTypes.WildcardType_NoWildcard)
	switch value := any.(type) {
	case nil:
		return nil
	case int:
		return AnyCondition(operator, multi, wildcard, AnyInt(value))
	case int32:
		return AnyCondition(operator, multi, wildcard, AnyInt32(value))
	case int64:
		return AnyCondition(operator, multi, wildcard, AnyInt64(value))
	case uint:
		return AnyCondition(operator, multi, wildcard, AnyUInt(value))
	case uint32:
		return AnyCondition(operator, multi, wildcard, AnyUInt32(value))
	case uint64:
		return AnyCondition(operator, multi, wildcard, AnyUInt64(value))
	case bool:
		return AnyCondition(operator, multi, wildcard, AnyBool(value))
	case float32:
		return AnyCondition(operator, multi, wildcard, AnyFloat(value))
	case float64:
		return AnyCondition(operator, multi, wildcard, AnyDouble(value))
	case string:
		if g.IsEmpty(value) {
			return nil
		}
		if strings.HasPrefix(value, ConditionQueryPrefix) && strings.HasSuffix(value, ConditionQuerySuffix) {
			conditionJson, err := gjson.DecodeToJson(value[9:])
			if err != nil {
				g.Log().Warningf(ctx, "Cannot decode %s to goguruTypes.Condition", value)
				return nil
			}
			operatorValue := goguruTypes.MustParseOperatorType(conditionJson.Get(FieldNameOperator).String(), goguruTypes.OperatorType_EQ)
			multiValue := goguruTypes.MustParseMultiType(conditionJson.Get(FieldNameMulti).String(), goguruTypes.MultiType_NoMulti)
			wildcardValue := goguruTypes.MustParseWildcardType(conditionJson.Get(FieldNameWildcard).String(), goguruTypes.WildcardType_NoWildcard)
			return AnyCondition(operatorValue, multiValue, wildcardValue, AnyInterface(ctx, conditionJson.Get(FieldNameValue).Interface(), tag))
		}
		switch wildcard {
		case goguruTypes.WildcardType_Contains:
			return AnyCondition(operator, multi, wildcard, AnyString("%"+value+"%"))
		case goguruTypes.WildcardType_StartsWith:
			return AnyCondition(operator, multi, wildcard, AnyString(value+"%"))
		case goguruTypes.WildcardType_EndsWith:
			return AnyCondition(operator, multi, wildcard, AnyString("%"+value))
		default:
			return AnyCondition(operator, multi, wildcard, AnyString(value))
		}
	default:
		t := reflect.TypeOf(any)
		switch t.Kind() {
		case reflect.Slice, reflect.Array:
			return AnySliceCondition(ctx, any, tag)
		case reflect.Ptr:
			return AnyInterfaceCondition(ctx, reflect.ValueOf(any).Elem(), tag)
		default:
			g.Log().Warningf(ctx, "Cannot convert type %s to *anypb.Any", t.String())
			return nil
		}
	}
}
