package config

import (
	"context"
	"log"
	"os"
	"project-shopping/pkg/logger"
	"strings"
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
	redisNodes := os.Getenv("REDIS_CLUSTER_NODES")
	if redisNodes == "" {
		logger.Log.Fatal().Msg("REDIS_CLUSTER_NODES environment variable is not set")
	}

	redisNodesList := strings.Split(redisNodes, ",")
	log.Printf("Connecting to Redis Cluster nodes: %v", redisNodesList)

	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        redisNodesList,
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
