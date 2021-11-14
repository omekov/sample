package stores

import (
	"context"
	"github.com/jackc/pgx/v4"

	"github.com/omekov/sample/internal/config"
	"github.com/omekov/sample/pkg/repository/postgresql"
	"github.com/sirupsen/logrus"
)

type Store struct {
	// MongoClient     mongodb.Client
	// RedisClient     *redis.Client
	// RabbitMQConn    *amqp.Connection
	PostgresConn *pgx.Conn
}

func NewStore(ctx context.Context, env *config.ENV, logger *logrus.Logger) *Store {

	// redisClient := redisdb.NewRedisConfig(env.Caches.Redis, logger)
	// mongoClient := mongodb.NewClient(env.Databases.Mongo, logger)
	// RabbitMQConn := rabbitmq.NewClient(env.Queues.RabbitMQ.KeepAlivePollPeriod, logger)
	PostgresManager := postgresql.NewManager(env.Databases.PostgreSQL.URL, logger)
	return &Store{
		// MongoClient:    mongoClient.GetConnection(ctx),
		// RedisClient:     redisClient.GetConnection(ctx),
		// RabbitMQConn:    RabbitMQConn.GetConnection(env.Queues.RabbitMQ.URI),
		PostgresConn: PostgresManager.GetConnection(ctx),
	}
}
