package gworm

import (
	"context"
	"reflect"
	"strings"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/wesleywu/gowing/protobuf/gwtypes"
	"github.com/wesleywu/ri-service-provider/gwerror"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	ConditionQueryPrefix = "condition{"
	ConditionQuerySuffix = "}"
	TagNameMulti         = "multi"
	TagNameWildcard      = "wildcard"
)

type FilterRequest struct {
	PropertyFilters []*PropertyFilter
	Filters         *bson.D
}

type PropertyFilter struct {
	Property string               `json:"property"`
	Value    interface{}          `json:"value"`
	Operator gwtypes.OperatorType `json:"operator"`
	Multi    gwtypes.MultiType    `json:"multi"`
	Wildcard gwtypes.WildcardType `json:"wildcard"`
}

func WhereEq(key string, value interface{}) *bson.E {
	return &bson.E{
		Key:   key,
		Value: value,
	}
}

func WhereNotEq(key string, value interface{}) *bson.E {
	return &bson.E{
		Key:   key,
		Value: bson.M{"$ne": value},
	}
}

func WhereGT(key string, value interface{}) *bson.E {
	return &bson.E{
		Key:   key,
		Value: bson.M{"$gt": value},
	}
}

func WhereGTE(key string, value interface{}) *bson.E {
	return &bson.E{
		Key:   key,
		Value: bson.M{"$gte": value},
	}
}

func WhereLT(key string, value interface{}) *bson.E {
	return &bson.E{
		Key:   key,
		Value: bson.M{"$lt": value},
	}
}

func WhereLTE(key string, value interface{}) *bson.E {
	return &bson.E{
		Key:   key,
		Value: bson.M{"$lte": value},
	}
}

func WhereIn(key string, value []interface{}) *bson.E {
	return &bson.E{
		key,
		bson.M{
			"$in": bson.A(value),
		},
	}
}

func WhereNotIn(key string, value []interface{}) *bson.E {
	return &bson.E{
		key,
		bson.M{
			"$nin": bson.A(value),
		},
	}
}

func WhereBetween(key string, min, max interface{}) *bson.E {
	return &bson.E{
		Key:   key,
		Value: bson.M{"$gte": min, "$lte": max},
	}
}

func WhereNotBetween(key string, min, max interface{}) *bson.E {
	return &bson.E{
		Key: key,
		Value: bson.D{
			{"$or",
				bson.A{
					bson.D{{key, bson.D{{"$gt", max}}}},
					bson.D{{key, bson.D{{"$lt", min}}}},
				},
			},
		},
	}
}

func WhereLike(key string, like string) *bson.E {
	return &bson.E{
		Key:   key,
		Value: bson.M{"$regex": like, "$options": "im"},
	}
}

func WhereNotLike(key string, like string) *bson.E {
	return &bson.E{
		Key: key,
		Value: bson.M{
			"$not": bson.M{"$regex": like, "$options": "im"},
		},
	}
}

func WhereNull(key string) *bson.E {
	return &bson.E{
		Key:   key,
		Value: nil,
	}
}

func WhereNotNull(key string) *bson.E {
	return &bson.E{
		Key:   key,
		Value: bson.M{"$ne": nil},
	}
}

// WherePri does the same logic as Model.Where except that if the parameter `where`
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given `where` parameter as "123", the
// WherePri function treats the condition as "id=123", but Model.Where treats the condition
// as string "123".
func WherePri(args []string) *bson.E {
	lenArgs := len(args)
	switch lenArgs {
	case 0:
		return nil
	case 1:
		return &bson.E{Key: "_id", Value: args[0]}
	default:
		return &bson.E{Key: "_id", Value: bson.E{Key: "$in", Value: args}}
	}
}

func (fr *FilterRequest) addPropertyFilter(f *PropertyFilter) *FilterRequest {
	fr.PropertyFilters = append(fr.PropertyFilters, f)
	return fr
}

func (fr *FilterRequest) GetFilters() (*bson.D, error) {
	if fr.Filters != nil {
		return fr.Filters, nil
	}
	var (
		filters bson.D
		filter  *bson.E
		err     error
	)
	filters = bson.D{}
	for _, pf := range fr.PropertyFilters {
		filter, err = pf.getFilter()
		if err != nil {
			return nil, err
		}
		if filter == nil {
			continue
		}
		filters = append(filters, *filter)
	}
	fr.Filters = &filters
	return fr.Filters, nil
}

func (pf *PropertyFilter) getFilter() (*bson.E, error) {
	if pf == nil {
		return nil, nil
	}
	if pf.Value == nil {
		return nil, nil
	}
	property := pf.Property
	switch pf.Operator {
	case gwtypes.OperatorType_EQ:
		switch pf.Multi {
		case gwtypes.MultiType_Exact:
			return WhereEq(property, pf.Value), nil
		case gwtypes.MultiType_Between:
			valueSlice := gconv.SliceAny(pf.Value)
			valueLen := len(valueSlice)
			if valueLen == 0 {
				return nil, nil
			} else if valueLen == 1 {
				return WhereEq(property, valueSlice[0]), nil
			} else if valueLen == 2 {
				return WhereBetween(property, valueSlice[0], valueSlice[1]), nil
			} else {
				return nil, gwerror.NewBadRequestErrorf("column %s requires between query but given %d values", property, valueLen)
			}
		case gwtypes.MultiType_NotBetween:
			valueSlice := gconv.SliceAny(pf.Value)
			valueLen := len(valueSlice)
			if valueLen == 0 {
				return nil, nil
			} else if valueLen == 1 {
				return WhereNotEq(property, valueSlice[0]), nil
			} else if valueLen == 2 {
				return WhereNotBetween(property, valueSlice[0], valueSlice[1]), nil
			} else {
				return nil, gwerror.NewBadRequestErrorf("column %s requires between query but given %d values", property, valueLen)
			}
		case gwtypes.MultiType_In:
			valueSlice := gconv.SliceAny(pf.Value)
			valueLen := len(valueSlice)
			if valueLen == 0 {
				return nil, nil
			} else if valueLen == 1 {
				return WhereEq(property, valueSlice[0]), nil
			} else {
				return WhereIn(property, valueSlice), nil
			}
		case gwtypes.MultiType_NotIn:
			valueSlice := gconv.SliceAny(pf.Value)
			valueLen := len(valueSlice)
			if valueLen == 0 {
				return nil, nil
			} else if valueLen == 1 {
				return WhereNotEq(property, valueSlice[0]), nil
			} else {
				return WhereNotIn(property, valueSlice), nil
			}
		}
	case gwtypes.OperatorType_NE:
		return WhereNotEq(property, pf.Value), nil
	case gwtypes.OperatorType_GT:
		return WhereGT(property, pf.Value), nil
	case gwtypes.OperatorType_GTE:
		return WhereGTE(property, pf.Value), nil
	case gwtypes.OperatorType_LT:
		return WhereLT(property, pf.Value), nil
	case gwtypes.OperatorType_LTE:
		return WhereLTE(property, pf.Value), nil
	case gwtypes.OperatorType_Like:
		valueStr := gconv.String(pf.Value)
		if g.IsEmpty(valueStr) {
			return nil, nil
		}
		valueStr = decorateValueStrForWildcard(valueStr, pf.Wildcard)
		return WhereLike(property, valueStr), nil
	case gwtypes.OperatorType_NotLike:
		valueStr := gconv.String(pf.Value)
		if g.IsEmpty(valueStr) {
			return nil, nil
		}
		valueStr = decorateValueStrForWildcard(valueStr, pf.Wildcard)
		return WhereNotLike(property, valueStr), nil
	case gwtypes.OperatorType_Null:
		return WhereNotEq(property, pf.Value), nil
	case gwtypes.OperatorType_NotNull:
		return WhereNotEq(property, pf.Value), nil
	}
	return nil, nil
}

func ExtractFilters(ctx context.Context, req interface{}, columnMap map[string]string) (fr FilterRequest, err error) {
	var f *PropertyFilter
	p := reflect.TypeOf(req)
	if p.Kind() != reflect.Ptr { // 要求传入值必须是个指针
		err = gwerror.NewBadRequestErrorf(req, "服务函数的输入参数必须是结构体指针")
		return
	}
	t := p.Elem()
	//g.Log().Debugf(ctx, "kind of input parameter is %s", t.Name())

	queryValue := reflect.ValueOf(req).Elem()

	// 循环结构体的字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		// Only do converting to public attributes.
		if !field.IsExported() {
			continue
		}
		fieldType := field.Type
		//g.Log().Debugf(ctx, "kind of field \"%s\" is %s", fieldName, field.Type.Kind().String())
		switch fieldType.Kind() {
		case reflect.Ptr:
			columnName, exists := columnMap[fieldName]
			if !exists {
				continue
			}
			//g.Log().Debugf(ctx, "kind of element of field %s is %s", fieldName, fieldElemType.Kind().String())
			anyValue := queryValue.Field(i).Interface().(*anypb.Any)
			if anyValue != nil {
				f, err = unwrapAnyFilter(columnName, field.Tag, anyValue)
				if err != nil {
					return
				}
				if f != nil {
					fr.addPropertyFilter(f)
				}
			}
		case reflect.Struct:
			structValue := queryValue.Field(i)
			g.Log().Debugf(ctx, "value of field %s is %x", fieldName, structValue)
			for si := 0; si < fieldType.NumField(); si++ {
				innerField := fieldType.Field(si)
				if innerField.Type.Kind() != reflect.Interface { // 仅处理类型为 interface{} 的字段
					continue
				}
				columnName, exists := columnMap[innerField.Name] // 仅处理在表字段定义中有的字段
				if !exists {
					continue
				}
				fieldValue := structValue.Field(si).Interface()
				if fieldValue == nil { // 不出来值为nil的字段
					continue
				}
				g.Log().Debugf(ctx, "inner field %s kind:%si, column:%s, value:%s", innerField.Name, innerField.Type.Kind().String(), columnName, fieldValue)
				f, err = parsePropertyFilter(ctx, req, columnName, innerField.Tag, fieldValue)
				if err != nil {
					return
				}
				fr.addPropertyFilter(f)
			}
		case reflect.Interface:
			columnName, exists := columnMap[fieldName]
			if !exists {
				continue
			}
			fieldValue := queryValue.Field(i).Interface()
			if fieldValue == nil {
				continue
			}
			f, err = parsePropertyFilter(ctx, req, columnName, field.Tag, fieldValue)
			if err != nil {
				return
			}
			fr.addPropertyFilter(f)
		default:
			continue
		}
	}
	_, err = fr.GetFilters()
	return
}

func parsePropertyFilter(ctx context.Context, req interface{}, columnName string, tag reflect.StructTag, value interface{}) (*PropertyFilter, error) {
	if value == nil { // todo processing: is null/is not null
		return nil, nil
	}
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Ptr:
		if t.Elem() == reflect.TypeOf(PropertyFilter{}) {
			pf := value.(*PropertyFilter)
			pf.Property = columnName
			return pf, nil
		}
		return &PropertyFilter{
			Property: columnName,
			Value:    value,
			Operator: gwtypes.OperatorType_EQ,
			Multi:    gwtypes.MultiType_Exact,
			Wildcard: gwtypes.WildcardType_None,
		}, nil
	case reflect.Slice, reflect.Array:
		valueSlice := gconv.SliceAny(value)
		multiTag, ok := tag.Lookup(TagNameMulti)
		if ok {
			multi, err := gwtypes.ParseMultiType(multiTag)
			if err != nil {
				return nil, err
			}
			switch len(valueSlice) {
			case 0:
				return nil, nil
			case 1:
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice[0],
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			default:
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    multi,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			}
		} else {
			switch len(valueSlice) {
			case 0:
				return nil, nil
			case 1:
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice[0],
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			default:
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_In,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			}
		}
	case reflect.Struct, reflect.Func, reflect.Map, reflect.Chan:
		return nil, gerror.Newf("Query field kind %s is not supported", t.Kind())
	case reflect.String:
		valueString := value.(string)
		if g.IsEmpty(valueString) {
			return nil, nil
		}
		if strings.HasPrefix(valueString, ConditionQueryPrefix) && strings.HasSuffix(valueString, ConditionQuerySuffix) {
			var condition *PropertyFilter
			err := gjson.DecodeTo(valueString[9:], &condition)
			condition.Property = columnName
			if err != nil {
				return nil, gwerror.NewBadRequestErrorf(req, err.Error())
			}
			g.Log().Debugf(ctx, "Query field type is orm.Condition: %s", gjson.MustEncodeString(condition))
			return condition, nil
		}
		wildcardString, ok := tag.Lookup(TagNameWildcard)
		if ok {
			wildcard, err := gwtypes.ParseWildcardType(wildcardString)
			if err != nil {
				return nil, err
			}
			switch wildcard {
			case gwtypes.WildcardType_Contains:
				return &PropertyFilter{
					Property: columnName,
					Value:    decorateValueStrForWildcard(valueString, gwtypes.WildcardType_Contains),
					Operator: gwtypes.OperatorType_Like,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: wildcard,
				}, nil
			case gwtypes.WildcardType_StartsWith:
				return &PropertyFilter{
					Property: columnName,
					Value:    decorateValueStrForWildcard(valueString, gwtypes.WildcardType_StartsWith),
					Operator: gwtypes.OperatorType_Like,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: wildcard,
				}, nil
			case gwtypes.WildcardType_EndsWith:
				return &PropertyFilter{
					Property: columnName,
					Value:    decorateValueStrForWildcard(valueString, gwtypes.WildcardType_EndsWith),
					Operator: gwtypes.OperatorType_Like,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: wildcard,
				}, nil
			default:
				return &PropertyFilter{
					Property: columnName,
					Value:    valueString,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: wildcard,
				}, nil
			}
		} else {
			return &PropertyFilter{
				Property: columnName,
				Value:    valueString,
				Operator: gwtypes.OperatorType_EQ,
				Multi:    gwtypes.MultiType_Exact,
				Wildcard: gwtypes.WildcardType_None,
			}, nil
		}
	default:
		return &PropertyFilter{
			Property: columnName,
			Value:    value,
			Operator: gwtypes.OperatorType_EQ,
			Multi:    gwtypes.MultiType_Exact,
			Wildcard: gwtypes.WildcardType_None,
		}, nil
	}
}

func unwrapAnyFilter(columnName string, tag reflect.StructTag, valueAny *anypb.Any) (pf *PropertyFilter, err error) {
	if valueAny == nil {
		return nil, nil
	}
	v, err := valueAny.UnmarshalNew()
	if err != nil {
		return nil, nil
	}

	switch vt := v.(type) {
	case *gwtypes.BoolSlice:
		return parseFieldSliceFilter(columnName, tag, vt.Value)
	case *gwtypes.DoubleSlice:
		return parseFieldSliceFilter(columnName, tag, vt.Value)
	case *gwtypes.FloatSlice:
		return parseFieldSliceFilter(columnName, tag, vt.Value)
	case *gwtypes.UInt32Slice:
		return parseFieldSliceFilter(columnName, tag, vt.Value)
	case *gwtypes.UInt64Slice:
		return parseFieldSliceFilter(columnName, tag, vt.Value)
	case *gwtypes.Int32Slice:
		return parseFieldSliceFilter(columnName, tag, vt.Value)
	case *gwtypes.Int64Slice:
		return parseFieldSliceFilter(columnName, tag, vt.Value)
	case *gwtypes.StringSlice:
		return parseFieldSliceFilter(columnName, tag, vt.Value)
	case *wrapperspb.BoolValue:
		return parseFieldSingleFilter(columnName, vt.Value)
	case *wrapperspb.BytesValue:
		return parseFieldSingleFilter(columnName, vt.Value)
	case *wrapperspb.DoubleValue:
		return parseFieldSingleFilter(columnName, vt.Value)
	case *wrapperspb.FloatValue:
		return parseFieldSingleFilter(columnName, vt.Value)
	case *wrapperspb.Int32Value:
		return parseFieldSingleFilter(columnName, vt.Value)
	case *wrapperspb.Int64Value:
		return parseFieldSingleFilter(columnName, vt.Value)
	case *wrapperspb.UInt32Value:
		return parseFieldSingleFilter(columnName, vt.Value)
	case *wrapperspb.UInt64Value:
		return parseFieldSingleFilter(columnName, vt.Value)
	case *wrapperspb.StringValue:
		return parseFieldSingleFilter(columnName, vt.Value)
	case *gwtypes.Condition:
		return addConditionFilter(columnName, vt)
	default:
		return nil, gerror.Newf("Unsupported value type: %v", vt)
	}
}

func parseFieldSingleFilter(columnName string, value interface{}) (pf *PropertyFilter, err error) {
	if value == nil {
		return nil, nil
	}
	return &PropertyFilter{
		Property: columnName,
		Value:    value,
		Operator: gwtypes.OperatorType_EQ,
		Multi:    gwtypes.MultiType_Exact,
		Wildcard: gwtypes.WildcardType_None,
	}, nil
}

func parseFieldSliceFilter[T any](columnName string, tag reflect.StructTag, value []T) (pf *PropertyFilter, err error) {
	if value == nil {
		return nil, nil
	}
	if multiTag, ok := tag.Lookup(TagNameMulti); ok {
		multi, err := gwtypes.ParseMultiType(multiTag)
		if err != nil {
			return nil, err
		}
		switch len(value) {
		case 0:
			return nil, nil
		case 1:
			return &PropertyFilter{
				Property: columnName,
				Value:    value[0],
				Operator: gwtypes.OperatorType_EQ,
				Multi:    gwtypes.MultiType_Exact,
				Wildcard: gwtypes.WildcardType_None,
			}, nil
		case 2:
			if multi == gwtypes.MultiType_Between {
				return &PropertyFilter{
					Property: columnName,
					Value:    value,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_Between,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			} else {
				return &PropertyFilter{
					Property: columnName,
					Value:    value,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_In,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			}
		default:
			return &PropertyFilter{
				Property: columnName,
				Value:    value,
				Operator: gwtypes.OperatorType_EQ,
				Multi:    gwtypes.MultiType_In,
				Wildcard: gwtypes.WildcardType_None,
			}, nil
		}
	} else {
		switch len(value) {
		case 0:
			return nil, nil
		case 1:
			return &PropertyFilter{
				Property: columnName,
				Value:    value[0],
				Operator: gwtypes.OperatorType_EQ,
				Multi:    gwtypes.MultiType_Exact,
				Wildcard: gwtypes.WildcardType_None,
			}, nil
		default:
			return &PropertyFilter{
				Property: columnName,
				Value:    value,
				Operator: gwtypes.OperatorType_EQ,
				Multi:    gwtypes.MultiType_In,
				Wildcard: gwtypes.WildcardType_None,
			}, nil
		}
	}
}

func addConditionFilter(columnName string, condition *gwtypes.Condition) (pf *PropertyFilter, err error) {
	if condition == nil {
		return nil, nil
	}
	// todo 当 condition.Operator 为 Null、NotNull 时，允许 nil 的 Value
	if condition.Value == nil {
		return nil, nil
	}
	v, err := condition.Value.UnmarshalNew()
	if err != nil {
		return nil, nil
	}
	switch vt := v.(type) {
	case *gwtypes.BoolSlice:
		return parseFieldConditionSliceFilter(columnName, condition, v.(*gwtypes.BoolSlice).Value)
	case *gwtypes.DoubleSlice:
		return parseFieldConditionSliceFilter(columnName, condition, v.(*gwtypes.DoubleSlice).Value)
	case *gwtypes.FloatSlice:
		return parseFieldConditionSliceFilter(columnName, condition, v.(*gwtypes.FloatSlice).Value)
	case *gwtypes.UInt32Slice:
		return parseFieldConditionSliceFilter(columnName, condition, v.(*gwtypes.UInt32Slice).Value)
	case *gwtypes.UInt64Slice:
		return parseFieldConditionSliceFilter(columnName, condition, v.(*gwtypes.UInt64Slice).Value)
	case *gwtypes.Int32Slice:
		return parseFieldConditionSliceFilter(columnName, condition, v.(*gwtypes.Int32Slice).Value)
	case *gwtypes.Int64Slice:
		return parseFieldConditionSliceFilter(columnName, condition, v.(*gwtypes.Int64Slice).Value)
	case *gwtypes.StringSlice:
		return parseFieldConditionSliceFilter(columnName, condition, v.(*gwtypes.StringSlice).Value)
	case *wrapperspb.BoolValue:
		return parseFieldConditionSingleFilter(columnName, condition, v.(*wrapperspb.BoolValue).Value)
	case *wrapperspb.BytesValue:
		return parseFieldConditionSingleFilter(columnName, condition, v.(*wrapperspb.BytesValue).Value)
	case *wrapperspb.DoubleValue:
		return parseFieldConditionSingleFilter(columnName, condition, v.(*wrapperspb.DoubleValue).Value)
	case *wrapperspb.FloatValue:
		return parseFieldConditionSingleFilter(columnName, condition, v.(*wrapperspb.FloatValue).Value)
	case *wrapperspb.Int32Value:
		return parseFieldConditionSingleFilter(columnName, condition, v.(*wrapperspb.Int32Value).Value)
	case *wrapperspb.Int64Value:
		return parseFieldConditionSingleFilter(columnName, condition, v.(*wrapperspb.Int64Value).Value)
	case *wrapperspb.UInt32Value:
		return parseFieldConditionSingleFilter(columnName, condition, v.(*wrapperspb.UInt32Value).Value)
	case *wrapperspb.UInt64Value:
		return parseFieldConditionSingleFilter(columnName, condition, v.(*wrapperspb.UInt64Value).Value)
	case *wrapperspb.StringValue:
		return parseFieldConditionSingleFilter(columnName, condition, v.(*wrapperspb.StringValue).Value)
	default:
		return nil, gerror.Newf("不支持的Value类型%v", vt)
	}
}

func parseFieldConditionSliceFilter[T any](columnName string, condition *gwtypes.Condition, valueSlice []T) (*PropertyFilter, error) {
	valueLen := len(valueSlice)
	if valueLen == 0 {
		return nil, nil
	}
	switch condition.Operator {
	case gwtypes.OperatorType_EQ:
		switch condition.Multi {
		case gwtypes.MultiType_Exact:
			if valueLen == 1 {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice[0],
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			} else {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_In,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			}
		case gwtypes.MultiType_Between:
			if valueLen == 1 {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice[0],
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			} else if valueLen == 2 {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_Between,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			} else {
				return nil, gwerror.NewBadRequestErrorf("column %s requires between query but given %d values", columnName, valueLen)
			}
		case gwtypes.MultiType_NotBetween:
			if valueLen == 1 {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice[0],
					Operator: gwtypes.OperatorType_NE,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			} else if valueLen == 2 {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_NotBetween,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			} else {
				return nil, gwerror.NewBadRequestErrorf("column %s requires between query but given %d values", columnName, valueLen)
			}
		case gwtypes.MultiType_In:
			if valueLen == 1 {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice[0],
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			} else {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_In,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			}
		case gwtypes.MultiType_NotIn:
			if valueLen == 1 {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice[0],
					Operator: gwtypes.OperatorType_NE,
					Multi:    gwtypes.MultiType_Exact,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			} else {
				return &PropertyFilter{
					Property: columnName,
					Value:    valueSlice,
					Operator: gwtypes.OperatorType_EQ,
					Multi:    gwtypes.MultiType_NotIn,
					Wildcard: gwtypes.WildcardType_None,
				}, nil
			}
		}
	case gwtypes.OperatorType_NE:
	case gwtypes.OperatorType_GT:
	case gwtypes.OperatorType_GTE:
	case gwtypes.OperatorType_LT:
	case gwtypes.OperatorType_LTE:
	case gwtypes.OperatorType_Like:
	case gwtypes.OperatorType_NotLike:
		return nil, gerror.Newf("Operator值为'%s'，但传入的Value: '%s'不应该是数组", gwtypes.OperatorType_name[int32(condition.Operator)], gconv.String(valueSlice))
	case gwtypes.OperatorType_Null:
	case gwtypes.OperatorType_NotNull:
		return nil, gerror.Newf("Operator值为'%s'，但传入的Value: '%s'应该为nil", gwtypes.OperatorType_name[int32(condition.Operator)], gconv.String(valueSlice))
	default:
		return nil, gerror.Newf("不支持的Operator值，传入Value: '%s'", gwtypes.OperatorType_name[int32(condition.Operator)], gconv.String(valueSlice))
	}
	return nil, nil
}

func parseFieldConditionSingleFilter(columnName string, condition *gwtypes.Condition, value interface{}) (*PropertyFilter, error) {
	if value == nil && condition.Operator != gwtypes.OperatorType_Null && condition.Operator != gwtypes.OperatorType_NotNull {
		return nil, nil
	}
	switch condition.Operator {
	case gwtypes.OperatorType_EQ:
		switch condition.Multi {
		case gwtypes.MultiType_Exact:
			return &PropertyFilter{
				Property: columnName,
				Value:    value,
				Operator: gwtypes.OperatorType_EQ,
				Multi:    gwtypes.MultiType_Exact,
				Wildcard: gwtypes.WildcardType_None,
			}, nil
		case gwtypes.MultiType_Between:
			return nil, gerror.Newf("Multi值为'%s'，但传入的Value: '%s'并非数组", gwtypes.MultiType_name[int32(gwtypes.MultiType_Between)], gconv.String(value))
		case gwtypes.MultiType_NotBetween:
			return nil, gerror.Newf("Multi值为'%s'，但传入的Value: '%s'并非数组", gwtypes.MultiType_name[int32(gwtypes.MultiType_NotBetween)], gconv.String(value))
		case gwtypes.MultiType_In:
			return nil, gerror.Newf("Multi值为'%s'，但传入的Value: '%s'并非数组", gwtypes.MultiType_name[int32(gwtypes.MultiType_In)], gconv.String(value))
		case gwtypes.MultiType_NotIn:
			return nil, gerror.Newf("Multi值为'%s'，但传入的Value: '%s'并非数组", gwtypes.MultiType_name[int32(gwtypes.MultiType_NotIn)], gconv.String(value))
		}
	case gwtypes.OperatorType_NE, gwtypes.OperatorType_GT, gwtypes.OperatorType_GTE, gwtypes.OperatorType_LT, gwtypes.OperatorType_LTE:
		return &PropertyFilter{
			Property: columnName,
			Value:    value,
			Operator: condition.Operator,
			Multi:    gwtypes.MultiType_Exact,
			Wildcard: gwtypes.WildcardType_None,
		}, nil
	case gwtypes.OperatorType_Like, gwtypes.OperatorType_NotLike:
		valueStr := gconv.String(value)
		if g.IsEmpty(valueStr) {
			return nil, nil
		}
		valueStr = decorateValueStrForWildcard(valueStr, condition.Wildcard)
		return &PropertyFilter{
			Property: columnName,
			Value:    valueStr,
			Operator: condition.Operator,
			Multi:    gwtypes.MultiType_Exact,
			Wildcard: condition.Wildcard,
		}, nil
	case gwtypes.OperatorType_Null, gwtypes.OperatorType_NotNull:
		return &PropertyFilter{
			Property: columnName,
			Value:    nil,
			Operator: condition.Operator,
			Multi:    gwtypes.MultiType_Exact,
			Wildcard: gwtypes.WildcardType_None,
		}, nil
	}
	return nil, nil
}

func decorateValueStrForWildcard(valueStr string, wildcardType gwtypes.WildcardType) string {
	switch wildcardType {
	case gwtypes.WildcardType_Contains:
		return valueStr
	case gwtypes.WildcardType_StartsWith:
		return "^" + valueStr
	case gwtypes.WildcardType_EndsWith:
		return valueStr + "$"
	}
	return valueStr
}
