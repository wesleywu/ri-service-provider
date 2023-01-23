package gwtypes

import (
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

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
	valueAny := &DoubleSlice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyFloatSlice(v []float32) *anypb.Any {
	valueAny := &FloatSlice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyInt64Slice(v []int64) *anypb.Any {
	valueAny := &Int64Slice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyUInt64Slice(v []uint64) *anypb.Any {
	valueAny := &UInt64Slice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyInt32Slice(v []int32) *anypb.Any {
	valueAny := &Int32Slice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyUInt32Slice(v []uint32) *anypb.Any {
	valueAny := &UInt32Slice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyBoolSlice(v []bool) *anypb.Any {
	valueAny := &BoolSlice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyStringSlice(v []string) *anypb.Any {
	valueAny := &StringSlice{Value: v}
	result, _ := anypb.New(valueAny)
	return result
}

func AnyCondition(operator OperatorType, multi MultiType, wildcard WildcardType, value *anypb.Any) *anypb.Any {
	valueCondition := &Condition{
		Operator: operator,
		Multi:    multi,
		Wildcard: wildcard,
		Value:    value,
	}
	result, _ := anypb.New(valueCondition)
	return result
}
