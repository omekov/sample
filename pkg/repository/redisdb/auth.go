package redisdb

import (
	"strconv"

	"github.com/omekov/sample/pkg/domain"
)

const (
	PREFIX = "session_customer_id_"
)

type Auth struct {
	redis *RedisConfig
}

func NewAuth() *Auth {
	return &Auth{}
}

// SetUserIDAndRefreshToken ...
func (c *Auth) SetUserIDAndRefreshToken(user domain.User, token string) error {
	return c.redis.client.Set(PREFIX+strconv.Itoa(int(user.ID)), token, 0).Err()
}

// GetRefreshToken ...
func (c *Auth) GetRefreshToken(user domain.User) (string, error) {
	return c.redis.client.Get(PREFIX + strconv.Itoa(int(user.ID))).Result()
}

func (c *Auth) SetBlackListCustomer() error {
	return nil
}
