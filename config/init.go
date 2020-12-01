package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/omekov/sample/internal/apiserver/models"
	"github.com/omekov/sample/internal/apiserver/stores/jwt"
	"github.com/omekov/sample/internal/apiserver/stores/mongodb"
)

const (
	MONGOURI                 = "MONGOURI"
	MONGONAME                = "MONGONAME"
	MONGOUSERNAME            = "MONGOUSERNAME"
	MONGOPASSWORD            = "MONGOPASSWORD"
	MONGOCUSTOMERSCOLLECTION = "MONGOCUSTOMERSCOLLECTION"
	MONGOROLESCOLLECTION     = "MONGOROLESCOLLECTION"
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
func Init(pathname string) {
	if err := godotenv.Load(pathname); err != nil {
		log.Fatal("Error loading .env file ", err)
	}
}

// GetMongoConfig ...
func GetMongoConfig() *mongodb.MongoConfig {
	return &mongodb.MongoConfig{
		Username:     IsReadyENV(MONGOUSERNAME),
		Password:     IsReadyENV(MONGOPASSWORD),
		URL:          IsReadyENV(MONGOURI),
		DatabaseName: IsReadyENV(MONGONAME),
		Collections: mongodb.Collections{
			Customer: IsReadyENV(MONGOCUSTOMERSCOLLECTION),
			Roles:    IsReadyENV(MONGOROLESCOLLECTION),
		},
	}
}

// GetRedisConfig ...
func GetRedisConfig() *models.RedisConfig {
	return &models.RedisConfig{
		Password: IsReadyENV(REDISPASSWORD),
		URL:      IsReadyENV(REDISURI),
	}
}

// GetJWTConfig
func GetJWTConfig() *jwt.Config {
	return &jwt.Config{
		RefreshTokenSecret: []byte(IsReadyENV(REFRESHTOKENSECRET)),
		AccessTokenSecret:  []byte(IsReadyENV(ACCESSTOKENSECRET)),
	}
}
