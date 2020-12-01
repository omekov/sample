package stores

import (
	"github.com/omekov/sample/internal/apiserver/stores/cache"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
	"github.com/omekov/sample/internal/apiserver/stores/mongodb"
)

// Store ...
type Store struct {
	Databases Databases
	JWT       *jwt.Config
	Caches    Caches
	Queues    Queues
	Websocket Websocket
}

// Databases ...
type Databases struct {
	MongoDB mongodb.MongoDBRepositories
	// PostgresDB
	// MySQLDB
	// Cassandra
	// ElasticSearch
}

// Queues ...
type Queues struct {
	RabbiMQ string
	Kafka   string
}

// Cachies ...
type Caches struct {
	MemCache    string
	RedisClient cache.Config
}

// Websocket ...
type Websocket struct{}
