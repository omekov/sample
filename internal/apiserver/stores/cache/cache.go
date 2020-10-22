package cache

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/omekov/sample/internal/apiserver/models"
)

const (
	prefix = "session_customer_id_"
)

// Config ...
type Config struct {
	Client *redis.Client
}

// NewClient ...
func NewClient(cnf *models.RedisConfig) (*redis.Client, error) {
	// config := &tls.Config{}
	options := redis.Options{
		Addr:     cnf.URL,
		Password: cnf.Password,
		// DB:       0,
	}
	client := redis.NewClient(&options)
	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Redis Ping - %s", pong)
	return client, nil
}

// SetCustomerIDAndRefreshToken ...
func (c *Config) SetCustomerIDAndRefreshToken(customer *models.Customer, token string) error {
	if err := c.Client.Set(prefix+customer.ID.String(), token, 5).Err(); err != nil {
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
