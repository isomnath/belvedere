package config

import (
	"fmt"
	"time"
)

type MongoConfig struct {
	hosts             string
	dbname            string
	username          string
	password          string
	maxPoolSize       int
	socketTimeout     time.Duration
	connectionTimeout time.Duration
}

func mongoConfig() *MongoConfig {
	return &MongoConfig{
		hosts:             getString(mongoHosts, false),
		dbname:            getString(mongoDbName, false),
		username:          getString(mongoUsername, false),
		password:          getString(mongoPassword, false),
		maxPoolSize:       getInt(mongoPoolSize, false),
		socketTimeout:     time.Duration(getInt(mongoSocketTimeout, false)),
		connectionTimeout: time.Duration(getInt(mongoConnectionTimeout, false)),
	}
}

func (mg *MongoConfig) ConnectionURL() string {
	return fmt.Sprintf("mongodb://%s:%s@%s/%s", mg.username, mg.password, mg.hosts, mg.dbname)
}

func (mg *MongoConfig) Hosts() string {
	return mg.hosts
}

func (mg *MongoConfig) DbName() string {
	return mg.dbname
}

func (mg *MongoConfig) Username() string {
	return mg.username
}

func (mg *MongoConfig) Password() string {
	return mg.password
}

func (mg *MongoConfig) PoolSize() int {
	return mg.maxPoolSize
}

func (mg *MongoConfig) SocketTimeout() time.Duration {
	return mg.socketTimeout
}

func (mg *MongoConfig) ConnectionTimeout() time.Duration {
	return mg.connectionTimeout
}
