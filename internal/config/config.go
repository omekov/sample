package config

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var once sync.Once

type ENV struct {
	Name      string
	Mode      string
	Port      int
	LogLevel  string
	LogType   string
	Databases *Databases
	Caches    *Caches
	Queues    *Queues
	JWT       *JWT
}

type Databases struct {
	Mongo      *Mongo
	PostgreSQL PostgreSQL
	// MySQL MySQL
}

type PostgreSQL struct {
	URL string
}

type Caches struct {
	Redis *Redis
	// MemCached MemCached
}

type Queues struct {
	RabbitMQ *RabbitMQ
	// Kafka Kafka
}

type Mongo struct {
	URI      string
	Name     string
	Username string
	Password string
}

type Redis struct {
	URL      string
	Password string
}

type RabbitMQ struct {
	URI                 string
	VHost               string
	ExchangeName        string
	QueueName           string
	PrefetchCount       int
	KeepAlivePollPeriod int
}

type JWT struct {
	AccessSecret  string
	RefreshSecret string
	LifeTimeAccess int
	LifeTimeRefresh int
}

// Get ...
func Get() *ENV {
	once.Do(func() {
		pathENV := "./deployments/.env"
		if _, err := os.Stat(pathENV); !os.IsNotExist(err) {
			if err := godotenv.Load(pathENV); err != nil {
				log.Fatal("Error loading .env file")
			}
		}
	})
	return &ENV{
		Name:     GetENVString("APP_NAME", "sample"),
		Mode:     GetENVString("APP_MODE", "development"),
		Port:     GetENVInt("PORT", 9090),
		LogLevel: GetENVString("LOG_LEVEL", "DEBUG"),
		LogType:  GetENVString("LOG_TYPE", "stdout"),
		Databases: &Databases{
			Mongo: &Mongo{
				URI:      GetENVString("MONGO_URI"),
				Username: GetENVString("MONGO_USERNAME"),
				Password: GetENVString("MONGO_PASSWORD"),
				Name:     GetENVString("MONGO_NAME"),
			},
		},
		Caches: &Caches{
			Redis: &Redis{
				URL:      GetENVString("REDIS_URI"),
				Password: GetENVString("REDIS_PASSWORD"),
			},
		},
		Queues: &Queues{
			RabbitMQ: &RabbitMQ{
				URI:           GetENVString("RABBITMQ_URI"),
				VHost:         GetENVString("RABBITMQ_VHOST", ""),
				ExchangeName:  GetENVString("RABBITMQ_EXCHANGE"),
				QueueName:     GetENVString("RABBITMQ_QUEUE"),
				PrefetchCount: GetENVInt("RABBITMQ_PREFETCH_COUNT", 5),
			},
		},
		JWT: &JWT{
			AccessSecret:  GetENVString("ACCESS_TOKEN_SECRET", "ACCESS_TOKEN_SECRET"),
			RefreshSecret: GetENVString("REFRESH_TOKEN_SECRET", "REFRESH_TOKEN_SECRET"),
		},
	}
}

// GetENVString ...
func GetENVString(ev string, defVal ...string) string {
	v := os.Getenv(ev)
	if v == "" {
		if len(defVal) == 0 {
			log.Fatalf("Not exists require env variable %s", ev)
		}
		v = defVal[0]
	}
	return v
}

func GetENVBool(ev string, defVal ...bool) bool {
	v, err := strconv.ParseBool(os.Getenv(ev))
	if err != nil {
		if len(defVal) == 0 {
			log.Fatalf("Not exists require env variable %s", ev)
		}
		v = defVal[0]
	}
	return v
}

func GetENVInt(ev string, defVal ...int) int {
	v, err := strconv.ParseInt(os.Getenv(ev), 10, 32)
	if err != nil {
		if len(defVal) == 0 {
			log.Fatalf("Not exists require env variable %s", ev)
		}
		return defVal[0]
	}
	return int(v)
}

func GetENVSliceOfString(ev string, defVal ...[]string) []string {
	v := os.Getenv(ev)
	if v == "" {
		if len(defVal) == 0 {
			log.Fatalf("Not exists require env variable %s", ev)
		}
		return defVal[0]
	}
	var ret []string
	err := json.Unmarshal([]byte(v), &ret)
	if err != nil {
		log.Fatalf("Error unmarshal slice (%s) %v", v, err)
	}
	return ret
}
