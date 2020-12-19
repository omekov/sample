package redisdb

import (
	"github.com/go-redis/redis"
	"github.com/omekov/sample/configs"
	"github.com/omekov/sample/internal/apiserver/models"
	log "github.com/sirupsen/logrus"
)

const (
	prefix = "session_customer_id_"
)

// Config ...
type Config struct {
	Client *redis.Client
}

// NewClient ...
func NewClient(cnf *configs.Redis) (*redis.Client, error) {
	options := redis.Options{
		Addr:     cnf.URL,
		Password: cnf.Password,
	}
	client := redis.NewClient(&options)
	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	log.Infof("Redis PING - %s", pong)
	return client, nil
}

// SetCustomerIDAndRefreshToken ...
func (c *Config) SetCustomerIDAndRefreshToken(customer *models.Customer, token string) error {
	if err := c.Client.Set(prefix+customer.ID.String(), token, 0).Err(); err != nil {
		return err
	}
	return nil
}

// GetRefreshToken ...
func (c *Config) GetRefreshToken(customer *models.Customer) (string, error) {
	refToken, err := c.Client.Get(prefix + customer.ID.String()).Result()
	if err != nil {
		return "", err
	}
	return refToken, nil
}
