package stores

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/omekov/sample/configs"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
	"github.com/omekov/sample/internal/apiserver/stores/mongodb"
	"github.com/omekov/sample/internal/apiserver/stores/rabbitmq"
	"github.com/omekov/sample/internal/apiserver/stores/redisdb"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

const KeepAlivePollPeriod = 3

// Store ...
type Store struct {
	Mongo        mongodb.Database
	JWT          *jwt.Config
	Redis        redisdb.Config
	RabbitMQ     rabbitmq.Config
	CustomerRepo CustomerRepository
}

// mongoDBConnect -
func mongoDBConnect(ctx context.Context, cnf *configs.Mongo) (*mongodb.Database, error) {
	dbClient, err := mongodb.NewClient(cnf)
	if err != nil {
		return nil, err
	}
	if err = dbClient.Connect(); err != nil {
		return nil, err
	}
	// defer dbClient.DisConnect(ctx)
	if err := dbClient.Ping(ctx); err != nil {
		return nil, err
	}
	log.Infof("MongoDB PING - %s", "PONG")
	db := mongodb.NewDatabase(cnf.Name, dbClient)
	return &db, nil
}

// redisConnect - методе подключаемся к redis и кидаем пинг
func redisConnect(cnf *configs.Redis) (*redis.Client, error) {
	return redisdb.NewClient(cnf)
}

func rabbitMQConnect(cnf *configs.RabbitMQ) (*amqp.Connection, error) {
	return amqp.Dial(cnf.URI)
}

// New ...
func (store *Store) New(ctx context.Context, env *configs.ENV) (*Store, error) {
	redisClient, err := redisConnect(env.Caches.Redis)
	if err != nil {
		return nil, errors.Wrap(err, "Redisdb")
	}
	mongoDB, err := mongoDBConnect(ctx, env.Databases.Mongo)
	if err != nil {
		go store.keepAliveMongo(ctx, env.Databases.Mongo)
		return nil, errors.Wrap(err, "Mongodb")
	}
	rabbitMQ, err := rabbitMQConnect(env.Queues.RabbitMQ)
	if err != nil {
		return nil, errors.Wrap(err, "RabbitMQ")
	}
	store = &Store{
		Mongo: *mongoDB,
		JWT: &jwt.Config{
			AccessTokenSecret:  []byte(env.JWT.Access),
			RefreshTokenSecret: []byte(env.JWT.Refresh),
		},
		Redis: redisdb.Config{
			Client: redisClient,
		},
		RabbitMQ: rabbitmq.Config{
			Conn:          rabbitMQ,
			ExchangeName:  env.Queues.RabbitMQ.Exchange,
			QueueName:     env.Queues.RabbitMQ.Queue,
			VHost:         env.Queues.RabbitMQ.VHost,
			PrefetchCount: env.Queues.RabbitMQ.PrefetchCount,
		},
		CustomerRepo: mongodb.NewCustomerRepository(*mongoDB),
	}
	return store, nil
}

func (store *Store) keepAliveMongo(ctx context.Context, cnf *configs.Mongo) {
	var err error
	for {
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false

		if store.Mongo != nil {
			lostConnect = true
		} else if err = store.Mongo.Client().Ping(ctx); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		log.Debug("[store.KeepAliveMongo] Lost MongoDB connection. Restoring...")
		_, err = mongoDBConnect(ctx, cnf)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Debug("[store.KeepAlivePg] PostgreSQL reconnected")
	}

}
