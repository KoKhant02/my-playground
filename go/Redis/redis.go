package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type GasPriceEstimate struct {
	PlatformId         string     `json:"platformId"`
	EstimatedGasPrices []GasPrice `json:"estimatedGasPrices"`
}

type GasPrice struct {
	Option            string `json:"option"`
	TotalFee          string `json:"totalFee"`
	MaxPriorityFeeTip string `json:"maxPriorityFeeTip"`
	MaxFee            string `json:"maxFee"`
}

func getRedisKey(ctx context.Context, client *redis.Client, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key does not exist
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func insertRedisKey(ctx context.Context, client *redis.Client, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return client.Set(ctx, key, data, 0).Err()
}

func main() {
	// Directly set the Redis reader endpoint and password
	redisEndpoint := "redis-xx.x2xxx5.xx.0001.xxxx1.cache.amazonaws.com:1111"
	redisPassword := "Xx12345678xX"

	// Setup Redis client
	options := &redis.Options{
		Addr:     redisEndpoint,
		Password: redisPassword,
		DB:       0,
	}

	client := redis.NewClient(options)
	ctx := context.Background()

	// Define gas price estimates
	estimates := []GasPriceEstimate{
		{
			PlatformId: "1", // Sepolia Testnet
			EstimatedGasPrices: []GasPrice{
				{"SLOW", "14181927690", "413659435", "21066061817"},
				{"NORMAL", "14408109969", "639841714", "21292244096"},
				{"FAST", "15317466217", "1549197962", "22201600345"},
			},
		},
		{
			PlatformId: "2", // Polygon Mumbai Testnet
			EstimatedGasPrices: []GasPrice{
				{"SLOW", "78450000101", "33450000071", "100950000116"},
				{"NORMAL", "80637500102", "35637500072", "103137500117"},
				{"FAST", "82825000104", "37825000074", "105325000119"},
			},
		},
	}

	// Insert estimates into Redis
	for _, estimate := range estimates {
		err := insertRedisKey(ctx, client, fmt.Sprintf("gasPrices:%s", estimate.PlatformId), estimate)
		if err != nil {
			log.Fatalf("Error inserting key: %v", err)
		}
	}

	// Get keys and print them
	allKeys, err := client.Keys(ctx, "*").Result()
	if err != nil {
		log.Fatalf("Error fetching keys: %v", err)
	}

	fmt.Printf("All keys in Redis: %v\n", allKeys)

	// Example usage of getRedisKey
	for _, key := range allKeys {
		val, err := getRedisKey(ctx, client, key)
		if err != nil {
			log.Fatalf("Error fetching key: %v", err)
		}
		if val != "" {
			fmt.Printf("Key: %s, Value: %s\n", key, val)
		}
	}
}
