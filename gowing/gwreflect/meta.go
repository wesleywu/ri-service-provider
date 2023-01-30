package gwreflect

import (
	"context"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gstructs"
	"reflect"
)

const (
	metaAttributeName = "Meta"
	metaTypeName      = "gmeta.Meta" // metaTypeName is for type string comparison.
)

func GetMetaField(ctx context.Context, req interface{}) (*reflect.StructField, error) {
	ctx, span := gtrace.NewSpan(ctx, "GetMetaField")
	defer span.End()
	reflectType, err := gstructs.StructType(req)
	if err != nil {
		return nil, err
	}
	metaField, ok := reflectType.FieldByName(metaAttributeName)
	if !ok {
		return nil, nil
	}
	if metaField.Type.String() != metaTypeName {
		return nil, nil
	}
	return &metaField, nil
}
