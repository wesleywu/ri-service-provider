package gwreflect

import (
	"context"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gtag"
	"reflect"
)

func MergeDefaultStructValue(_ context.Context, pointer interface{}) error {
	defaultValueTags := []string{gtag.DefaultShort, gtag.Default}
	tagFields, err := gstructs.TagFields(pointer, defaultValueTags)
	if err != nil {
		return err
	}
	if len(tagFields) > 0 {
		for _, field := range tagFields {
			if field.Value.IsZero() {
				//fieldValue := reflect.ValueOf(field.TagValue)
				tagValueConverted := gconv.Convert(field.TagValue, field.Type().String())
				field.Value.Set(reflect.ValueOf(tagValueConverted))
			}
		}
	}
	return nil
}
