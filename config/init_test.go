package config_test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/omekov/sample/config"
	"github.com/stretchr/testify/assert"
)

func TestGetMongoConfig(t *testing.T) {
	godotenv.Load()
	conf := config.GetMongoConfig()

	assert.NotEmpty(t, conf)
}

func TestGetRedisConfig(t *testing.T) {
	godotenv.Load()
	conf := config.GetRedisConfig()

	assert.NotEmpty(t, conf)
}
