package enum

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

// MarshalBSONValue 确保 ContentType 在往 mongodb 写入数据时，使用字符串形式保存
func (e ContentType) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.TypeString, bsoncore.AppendString(nil, ContentType_name[int32(e)]), nil
}

// UnmarshalBSONValue 确保 ContentType 在从 mongodb 读取数据时，使用字符串形式读取，并解析为对应的 ContentType
func (e *ContentType) UnmarshalBSONValue(t bsontype.Type, b []byte) error {
	var (
		enumStr   string
		enumInt32 int32
		ok        bool
	)
	enumStr, _, ok = bsoncore.ReadString(b)
	if !ok {
		return errors.New("ContentType UnmarshalBSONValue error")
	}

	if enumInt32, ok = ContentType_value[enumStr]; !ok {
		return fmt.Errorf("%s is not a valid ContentType", enumStr)
	}
	*e = ContentType(enumInt32)
	return nil
}

// MarshalBSONValue 确保 FilterType 在往 mongodb 写入数据时，使用字符串形式保存
func (e FilterType) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.TypeString, bsoncore.AppendString(nil, FilterType_name[int32(e)]), nil
}

// UnmarshalBSONValue 确保 FilterType 在从 mongodb 读取数据时，使用字符串形式读取，并解析为对应的 FilterType
func (e *FilterType) UnmarshalBSONValue(t bsontype.Type, b []byte) error {
	var (
		enumStr   string
		enumInt32 int32
		ok        bool
	)
	enumStr, _, ok = bsoncore.ReadString(b)
	if !ok {
		return errors.New("FilterType UnmarshalBSONValue error")
	}

	if enumInt32, ok = FilterType_value[enumStr]; !ok {
		return fmt.Errorf("%s is not a valid FilterType", enumStr)
	}
	*e = FilterType(enumInt32)
	return nil
}
