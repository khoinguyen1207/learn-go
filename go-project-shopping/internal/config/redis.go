package config

import (
	"context"
	"log"
	"project-shopping/pkg/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Address,
		Username:     cfg.Redis.Username,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     20,               // Maximum number of connections in the pool
		MinIdleConns: 5,                // Minimum number of idle connections
		DialTimeout:  10 * time.Second, // Connection timeout
		ReadTimeout:  3 * time.Second,  // Read timeout
		WriteTimeout: 3 * time.Second,  // Write timeout
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Failed to connect to Redis")
	}

	log.Println("✅ Connected to Redis successfully!")

	return client
}

func InitRedisCluster(cfg *Config) *redis.ClusterClient {
	if cfg.Redis.RedisNodes == nil || len(cfg.Redis.RedisNodes) == 0 {
		logger.Log.Fatal().Msg("❌ No Redis cluster nodes provided in configuration")
	}
	log.Printf("Connecting to Redis Cluster nodes: %v", cfg.Redis.RedisNodes)

	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        cfg.Redis.RedisNodes,
		Username:     cfg.Redis.Username,
		Password:     cfg.Redis.Password,
		PoolSize:     20,               // Maximum number of connections in the pool
		MinIdleConns: 5,                // Minimum number of idle connections
		DialTimeout:  10 * time.Second, // Connection timeout
		ReadTimeout:  3 * time.Second,  // Read timeout
		WriteTimeout: 3 * time.Second,  // Write timeout
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := clusterClient.Ping(ctx).Err(); err != nil {
		logger.Log.Fatal().Err(err).Msg("❌ Failed to connect to Redis Cluster")
	}

	log.Println("✅ Connected to Redis Cluster successfully!")

	return clusterClient
}
