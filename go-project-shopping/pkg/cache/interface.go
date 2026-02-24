package cache

import (
	"context"
	"time"
)

type CacheService interface {
	Get(ctx context.Context, key string, dest any) error
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Clear(ctx context.Context, pattern string) error
	Delete(ctx context.Context, key string) error
}
