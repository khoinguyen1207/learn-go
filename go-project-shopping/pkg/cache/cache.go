package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheService struct {
	rdb *redis.Client
}

func NewCacheService(rdb *redis.Client) CacheService {
	return &cacheService{
		rdb: rdb,
	}
}

func (cs *cacheService) Get(ctx context.Context, key string, dest any) error {
	data, err := cs.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return err
	}

	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(data), dest)
}

func (cs *cacheService) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return cs.rdb.Set(ctx, key, data, ttl).Err()
}

func (cs *cacheService) Clear(ctx context.Context, pattern string) error {
	cursor := uint64(0)

	for {
		keys, nextCursor, err := cs.rdb.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			cs.rdb.Del(ctx, keys...)
		}

		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return nil
}

func (cs *cacheService) Delete(ctx context.Context, key string) error {
	return cs.rdb.Del(ctx, key).Err()
}

func (cs *cacheService) Exists(ctx context.Context, key string) (bool, error) {
	count, err := cs.rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
