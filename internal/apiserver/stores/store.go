package stores

import (
	"github.com/omekov/sample/internal/apiserver/stores/caches/redisdb"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
	"github.com/omekov/sample/internal/apiserver/stores/mongodb"
	"github.com/streadway/amqp"
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
	Mongo mongodb.Repository
	// PostgresDB
	// MariaDB
	// Cassandra
	// ElasticSearch
	// CouchDB
}

// Queues ...
type Queues struct {
	RabbiMQ amqp.Authentication
	Kafka   string
}

// Caches ...
type Caches struct {
	MemCache    string
	RedisClient redisdb.Config
}

// Websocket ...
type Websocket struct{}
