package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/omekov/sample/internal/apiserver/models"
)

const (
	MONGOURI                 = "MONGOURI"
	MONGONAME                = "MONGONAME"
	MONGOUSERNAME            = "MONGOUSERNAME"
	MONGOPASSWORD            = "MONGOPASSWORD"
	MONGOCUSTOMERSCOLLECTION = "MONGOCUSTOMERSCOLLECTION"
	ACCESSTOKENSECRET        = "ACCESSTOKENSECRET"
	REFRESHTOKENSECRET       = "REFRESHTOKENSECRET"
	PORT                     = "PORT"
	REDISPASSWORD            = "REDISPASSWORD"
	REDISURI                 = "REDISURI"
)

// IsReadyENV ...
func IsReadyENV(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("Error is not env - %s", key)
		return ""
	}
	return os.Getenv(key)
}

// Init ...
func Init() {
	useENV := flag.Bool("env", false, "not load env file")
	flag.Parse()
	if *useENV {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file ", err)
		}
	}
}

// GetMongoConfig ...
func GetMongoConfig() *models.MongoConfig {
	return &models.MongoConfig{
		Username:     IsReadyENV(MONGOUSERNAME),
		Password:     IsReadyENV(MONGOPASSWORD),
		URL:          IsReadyENV(MONGOURI),
		DatabaseName: IsReadyENV(MONGONAME),
	}
}

// GetRedisConfig ...
func GetRedisConfig() *models.RedisConfig {
	return &models.RedisConfig{
		Password: IsReadyENV(REDISPASSWORD),
		URL:      IsReadyENV(REDISURI),
	}
}
