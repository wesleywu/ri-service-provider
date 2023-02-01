package gwcache

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/proto"
)

func deserialize(data []byte, m proto.Message) error {
	if g.IsEmpty(data) {
		return ErrEmptyCachedValue
	}
	return proto.Unmarshal(data, m)
}

func serialize(value proto.Message) ([]byte, error) {
	if value == nil {
		return nil, nil
	}
	return proto.Marshal(value)
}
