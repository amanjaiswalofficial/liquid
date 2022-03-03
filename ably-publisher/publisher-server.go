package main

import (
	"log"
	"os"
	"time"
	"fmt"
	"context"
	"sync"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/ably/ably-go/ably"
)

var ctx = context.Background()

var wg = &sync.WaitGroup()

func main(){

}

func getRedis() *redis.Client {
	var (
		host = getEnv("REDIS_HOST", "localhost")
		port = string(getEnv("REDIS_PORT", "6379"))
		password = getEnv("REDIS_PASSWORD", "")
	)

	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		Password: password,
		DB: 	0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func getEnv(envName, valueDefault string) string {
	value := os.GetEnv(envName)
	if value == "" {
		return valueDefault
	}

	return value
}

func getAblyChannel() *ably.RealtimeChannel {
	ablyClient, err := ably.NewRealtime(
		ably.WithKey(getEnv("ABLY_KEY", "No key specified"))
		// ably.WithEchoMessages(false)
	)
	if err != nil {
		panic(err)
	}

	return ablyClient.Channels.Get(getEnv("CHANNEL_NAME", "trades"))
}
