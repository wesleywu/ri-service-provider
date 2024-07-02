// Code generated by protoc-gen-go-guru. DO NOT EDIT.
// versions:
//  protoc-gen-go-guru v0.2.14
// source: app/episode/service/enum/episode_enum.proto

package enum

import (
	json "encoding/json"
	errors "errors"
	fmt "fmt"

	bson "go.mongodb.org/mongo-driver/bson"
	bsontype "go.mongodb.org/mongo-driver/bson/bsontype"
	bsoncore "go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
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

// MarshalJSON 确保 ContentType 在导出为 Json 时，使用字符串形式编码
func (e ContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ContentType_name[int32(e)])
}

// UnmarshalJSON 确保从 Json 中解码 ContentType 类型时，使用字符串形式解码，并解析为对应的 ContentType
func (e *ContentType) UnmarshalJSON(data []byte) error {
	var (
		enumStr   string
		enumInt32 int32
		ok        bool
	)
	if err := json.Unmarshal(data, &enumStr); err != nil {
		return err
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

// MarshalJSON 确保 FilterType 在导出为 Json 时，使用字符串形式编码
func (e FilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(FilterType_name[int32(e)])
}

// UnmarshalJSON 确保从 Json 中解码 FilterType 类型时，使用字符串形式解码，并解析为对应的 FilterType
func (e *FilterType) UnmarshalJSON(data []byte) error {
	var (
		enumStr   string
		enumInt32 int32
		ok        bool
	)
	if err := json.Unmarshal(data, &enumStr); err != nil {
		return err
	}
	if enumInt32, ok = FilterType_value[enumStr]; !ok {
		return fmt.Errorf("%s is not a valid FilterType", enumStr)
	}
	*e = FilterType(enumInt32)
	return nil
}
