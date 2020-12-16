package apiserver

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/omekov/sample/configs"
	"github.com/omekov/sample/internal/apiserver/handlers"
	"github.com/omekov/sample/internal/apiserver/stores"
	"github.com/omekov/sample/internal/apiserver/stores/caches/redisdb"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
	"github.com/omekov/sample/internal/apiserver/stores/mongodb"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

// Run ...
func Run() {
	if err := app(); err != nil {
		log.Fatal(err)
	}
}

// mongoDBConnect -
func mongoDBConnect(cnf *configs.Mongo) (mongodb.Database, error) {
	dbClient, err := mongodb.NewClient(cnf)
	if err != nil {
		return nil, err
	}
	if err = dbClient.Connect(); err != nil {
		return nil, err
	}
	if err = dbClient.Ping(context.TODO()); err != nil {
		return nil, err
	}
	log.Infof("MongoDB PING - %s\n", "PONG")
	db := mongodb.NewDatabase(cnf, dbClient)
	return db, nil
}

// redisConnect - методе подключаемся к redis и кидаем пинг
func redisConnect(cnf *configs.Redis) (*redis.Client, error) {
	return redisdb.NewClient(cnf)
}

func app() error {
	env := configs.NewENV()
	redisClient, err := redisConnect(env.Caches.Redis)
	if err != nil {
		return errors.Wrap(err, "Redis")
	}
	mongoDB, err := mongoDBConnect(env.Databases.Mongo)
	if err != nil {
		return errors.Wrap(err, "Mongodb")
	}
	server := handlers.Server{
		Config: handlers.ConfigureRouter(env.Port),
		Logger: log.New(),
		Store: &stores.Store{
			Databases: stores.Databases{
				Mongo: mongodb.Repository{
					DB: mongoDB,
				},
			},
			JWT: &jwt.Config{
				AccessTokenSecret:  []byte(env.JWT.Access),
				RefreshTokenSecret: []byte(env.JWT.Refresh),
			},
			Caches: stores.Caches{
				RedisClient: redisdb.Config{
					Client: redisClient,
				},
			},
		},
	}
	server.Run()
	return nil
}
