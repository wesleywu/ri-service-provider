package cache

import (
	"context"
	"errors"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

type Storage interface {
	Initialized() bool
	Get(ctx context.Context, key string) ([]byte, error)
	GetString(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, content []byte, duration time.Duration) error
	SetString(ctx context.Context, key string, content string, duration time.Duration) error
	Delete(ctx context.Context, keys []string) error
	SAdd(ctx context.Context, setKey, value string) error
	SMembers(ctx context.Context, setKey string) ([]string, error)
}

var (
	RedisClient            *redis.Client
	RedisLocker            *redislock.Client
	LockTimeout            time.Duration
	storage                Storage
	ErrNilResult           = redis.Nil
	ErrCacheNotInitialized = errors.New("cache: not initialized")
	ErrEmptyCacheKey       = errors.New("cache: cache key is empty")
	ErrLockTimeout         = errors.New("cache: lock timeout")
	ErrNotFound            = errors.New("cache: not found")
)

func init() {
	ctx := gctx.New()
	lockTimeoutSeconds := g.Cfg().MustGet(ctx, "redis.default.lockTimeoutSeconds", 3).Int()
	LockTimeout = time.Duration(lockTimeoutSeconds) * time.Second
	address := g.Cfg().MustGet(ctx, "redis.default.host", "127.0.0.1:6379").String()
	db := g.Cfg().MustGet(ctx, "redis.default.db", 0).Int()
	password := g.Cfg().MustGet(ctx, "redis.default.password", "").String()
	RedisClient = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     address,
		DB:       db,
		Password: password,
	})
	info := RedisClient.Info(ctx)
	cacheInitialized := false
	if info.Err() != nil {
		g.Log().Warningf(ctx, "Error initialize redis: %s", info.Err().Error())
	} else {
		cacheInitialized = true
	}
	if RedisClient != nil {
		RedisLocker = redislock.New(RedisClient)
	}
	storage = &redisStorage{
		initialized: cacheInitialized,
		Client:      RedisClient,
	}
}

type redisStorage struct {
	initialized bool
	Client      *redis.Client
}

func (s *redisStorage) Initialized() bool {
	return s.initialized
}

func (s *redisStorage) Get(ctx context.Context, key string) ([]byte, error) {
	ctx, span := gtrace.NewSpan(ctx, "Cache Get")
	defer span.End()
	return s.Client.Get(ctx, key).Bytes()
}

func (s *redisStorage) GetString(ctx context.Context, key string) (string, error) {
	ctx, span := gtrace.NewSpan(ctx, "Cache GetString")
	defer span.End()
	return s.Client.Get(ctx, key).Result()
}

func (s *redisStorage) Set(ctx context.Context, key string, content []byte, duration time.Duration) error {
	ctx, span := gtrace.NewSpan(ctx, "Cache Set")
	defer span.End()
	return s.Client.Set(ctx, key, content, duration).Err()
}

func (s *redisStorage) SetString(ctx context.Context, key string, content string, duration time.Duration) error {
	ctx, span := gtrace.NewSpan(ctx, "Cache SetString")
	defer span.End()
	return s.Client.Set(ctx, key, content, duration).Err()
}

func (s *redisStorage) Delete(ctx context.Context, keys []string) error {
	ctx, span := gtrace.NewSpan(ctx, "Cache Delete")
	defer span.End()
	return s.Client.Del(ctx, keys...).Err()
}

func (s *redisStorage) SAdd(ctx context.Context, setKey, value string) error {
	ctx, span := gtrace.NewSpan(ctx, "Cache SAdd")
	defer span.End()
	return s.Client.SAdd(ctx, setKey, value).Err()
}

func (s *redisStorage) SMembers(ctx context.Context, setKey string) ([]string, error) {
	ctx, span := gtrace.NewSpan(ctx, "Cache SMembers")
	defer span.End()
	return s.Client.SMembers(ctx, setKey).Result()
}
