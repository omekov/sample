package redisdb

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/omekov/sample/internal/config"
	"github.com/sirupsen/logrus"
)

// Client ...
type RedisConfig struct {
	cfg    *config.Redis
	logger *logrus.Logger
	client *redis.Client
}

// NewClient ...
func NewRedisConfig(cnf *config.Redis, logger *logrus.Logger) *RedisConfig {
	return &RedisConfig{
		cfg:    cnf,
		logger: logger,
	}
}

func (r *RedisConfig) GetConnection(ctx context.Context) *redis.Client {
	options := redis.Options{
		Addr:     r.cfg.URL,
		Password: r.cfg.Password,
	}
	r.client = redis.NewClient(&options)
	return r.client
}

func (r *RedisConfig) ping() error {
	pong, err := r.client.Ping().Result()
	if err != nil {
		return err
	}
	r.logger.Infof("Redis PING - %s", pong)
	return nil
}
