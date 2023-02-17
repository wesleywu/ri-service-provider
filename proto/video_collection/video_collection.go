package video_collection

import (
	"github.com/WesleyWu/gowing/util/gwwrapper"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reflect"
)

func (x *VideoCollectionItem) UnmarshalValue(value interface{}) error {
	var (
		id          *string
		name        *string
		contentType *int32
		filterType  *int32
		count       *uint32
		isOnline    *bool
		createdAt   *timestamppb.Timestamp
		updatedAt   *timestamppb.Timestamp
	)
	if record, ok := value.(gdb.Record); ok {
		if idVar := record["id"]; !idVar.IsNil() {
			id = gwwrapper.WrapString(idVar.String())
		}
		if nameVar := record["name"]; !nameVar.IsNil() {
			name = gwwrapper.WrapString(nameVar.String())
		}
		if contentTypeVar := record["content_type"]; !contentTypeVar.IsNil() {
			contentType = gwwrapper.WrapInt32(contentTypeVar.Int32())
		}
		if filterTypeVar := record["filter_type"]; !filterTypeVar.IsNil() {
			filterType = gwwrapper.WrapInt32(filterTypeVar.Int32())
		}
		if countVar := record["count"]; !countVar.IsNil() {
			count = gwwrapper.WrapUInt32(countVar.Uint32())
		}
		if isOnlineVar := record["is_online"]; !isOnlineVar.IsNil() {
			isOnline = gwwrapper.WrapBool(isOnlineVar.Bool())
		}
		if createdAtVar := record["created_at"]; !createdAtVar.IsNil() {
			createdAt = gwwrapper.WrapTimestamp(createdAtVar.Time())
		}
		if updatedAtVar := record["updated_at"]; !updatedAtVar.IsNil() {
			updatedAt = gwwrapper.WrapTimestamp(updatedAtVar.Time())
		}
		*x = VideoCollectionItem{
			Id:          id,
			Name:        name,
			ContentType: contentType,
			FilterType:  filterType,
			Count:       count,
			IsOnline:    isOnline,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
		}
		return nil
	}
	return gerror.Newf(`unsupported value type for UnmarshalValue: %v`, reflect.TypeOf(value))
}
