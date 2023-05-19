package store

import (
	"fmt"

	"gopkg.in/redis.v5"

	"github.com/isomnath/belvedere/config"
)

type RedisStore struct {
	client *redis.Client
}

var rd *RedisStore

//func initializeSecureDBConnection(config *config.RedisConfig, databaseID int) *redis.Client {
//	return redis.NewClient(
//		&redis.Options{
//			Addr:     fmt.Sprintf("%s:%d", config.Host(), config.Port()),
//			Username: config.Username(),
//			Password: config.Password(),
//			DB:       databaseID,
//		},
//	)
//}

func initializeDBConnection(config *config.RedisConfig, databaseID int) *redis.Client {
	cl := redis.NewClient(
		&redis.Options{
			Addr: fmt.Sprintf("%s:%d", config.Host(), config.Port()),
			DB:   databaseID,
		})
	return cl
}

// RedisConnect - Connects to mongo and initializes the DB client to be reused across the application
func RedisConnect(config *config.RedisConfig, databaseID int) {
	rd = &RedisStore{
		client: initializeDBConnection(config, databaseID),
	}
}

// RedisSecureConnect expects config to hold
//func RedisSecureConnect(config *config.RedisConfig, databaseID int) {
//	rd = &RedisStore{
//		client: initializeSecureDBConnection(config, databaseID),
//	}
//}

// GetRedisClient - Returns the instance of redis client created when RedisConnect was invoked
func GetRedisClient() *redis.Client {
	return rd.client
}
