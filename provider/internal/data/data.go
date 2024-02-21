package data

import (
	"context"

	"github.com/castbox/go-guru/pkg/goguru/conf"
	"github.com/castbox/go-guru/pkg/util/redis"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	MongoClient   *mongo.Client
	MongoDatabase *mongo.Database
	RedisClient   *redis.Cache
}

// NewData .
func NewData(ctx context.Context, c *conf.Data, logger log.Logger, mongoClient *mongo.Client, redisClient *redis.Cache) (*Data, func(), error) {
	helper := log.NewHelper(logger).WithContext(ctx)
	cleanup := func() {
		if mongoClient != nil {
			helper.Info("closing the data resources")
			err := mongoClient.Disconnect(ctx)
			if err != nil {
				helper.Errorf("Error closing mongodb connection: %+v", err)
			}
		}
		if redisClient != nil {
			_ = redisClient.Client.Close()
		}
	}
	return &Data{
		MongoClient:   mongoClient,
		MongoDatabase: mongoClient.Database(c.Database.DatabaseName),
		RedisClient:   redisClient,
	}, cleanup, nil
}
