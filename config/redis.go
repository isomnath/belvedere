package config

type RedisConfig struct {
	host     string
	port     int
	username string
	password string
}

func redisConfig() *RedisConfig {
	return &RedisConfig{
		host:     getString(redisHost, false),
		port:     getInt(redisPort, false),
		username: getString(redisUsername, false),
		password: getString(redisPassword, false),
	}
}

func (redis *RedisConfig) Host() string {
	return redis.host
}

func (redis *RedisConfig) Port() int {
	return redis.port
}

func (redis *RedisConfig) Username() string {
	return redis.username
}

func (redis *RedisConfig) Password() string {
	return redis.password
}
