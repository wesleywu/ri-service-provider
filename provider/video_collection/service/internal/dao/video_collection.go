package dao

import (
	"github.com/WesleyWu/ri-service-provider/provider/video_collection/service/internal/dao/internal"
)

// videoCollectionDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type videoCollectionDao struct {
	*internal.VideoCollectionDaoMongodb
}

var (
	// VideoCollection is globally public accessible object for table tools_gen_table operations.
	VideoCollection = videoCollectionDao{
		internal.NewVideoCollectionDaoMongodb(),
	}
)
