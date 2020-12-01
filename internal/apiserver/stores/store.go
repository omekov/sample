package stores

import (
	"github.com/omekov/sample/internal/apiserver/stores/cache"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
	"github.com/omekov/sample/internal/apiserver/stores/mongos"
)

// Store ...
type Store struct {
	Databases Databases
	JWT       *jwt.Config
	Cachies   Cachies
	Queues    Queues
	Websocket Websocket
}

// Databases ...
type Databases struct {
	MongoDB mongos.MongoDBRepositories
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
type Cachies struct {
	MemCache    string
	RedisClient cache.Config
}

// Websocket ...
type Websocket struct{}
