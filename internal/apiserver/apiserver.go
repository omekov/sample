package apiserver

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/omekov/sample/config"
	"github.com/omekov/sample/internal/apiserver/handlers"
	"github.com/omekov/sample/internal/apiserver/stores"
	"github.com/omekov/sample/internal/apiserver/stores/cache"
	"github.com/omekov/sample/internal/apiserver/stores/mongodb"

	"github.com/sirupsen/logrus"
)

// Run ...
func Run() {
	if err := app(); err != nil {
		log.Fatal(err)
	}
}

// mongoDBConnect - методе подключаемся к mongodb и создаем ему базу и документы, также возвращаем методы по клиенту
func mongoDBConnect() (mongodb.CustomerRepository, error) {
	dbClient, err := mongodb.NewClient(config.GetMongoConfig())
	if err != nil {
		return nil, err
	}
	if err = dbClient.Connect(); err != nil {
		return nil, err
	}
	db := mongodb.NewDatabase(config.GetMongoConfig(), dbClient)
	customer := mongodb.NewCustomerRepository(
		db,
		config.GetMongoConfig().Collections.Customer,
	)
	return customer, nil
}

// redisConnect - методе подключаемся к redis и кидаем пинг
func redisConnect() (*redis.Client, error) {
	return cache.NewClient(config.GetRedisConfig())
}

func app() error {
	config.Init(".env.prod")
	redisClient, err := redisConnect()
	if err != nil {
		return err
	}
	customer, err := mongoDBConnect()
	if err != nil {
		return err
	}
	server := handlers.Server{
		Config: handlers.ConfigureRouter(),
		Logger: logrus.New(),
		Store: &stores.Store{
			Databases: stores.Databases{
				MongoDB: mongodb.MongoDBRepositories{
					Customer: customer,
				},
			},
			JWT: config.GetJWTConfig(),
			Cachies: stores.Cachies{
				RedisClient: cache.Config{
					Client: redisClient,
				},
			},
		},
	}
	server.Run()
	return nil
}
