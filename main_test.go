package main

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/omekov/sample/internal/apiserver"
	api "github.com/omekov/sample/internal/apiserver"
	"github.com/stretchr/testify/assert"
)

func TestApp_ENV(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file ", err)
	}
	envs := []string{
		api.PORT,
		api.MONGOURI,
		api.MONGOUSERNAME,
		api.MONGOPASSWORD,
		api.MONGONAME,
		api.MONGOAUTHCOLLECTION,
		api.MONGOPODCASTCOLLECTION,
		api.TOKENSECRET}
	for _, e := range envs {
		if api.IsReadyENV(e) == "" {
			t.Errorf("is not ENV - %s", e)
		}
	}
}

func TestGetConfig(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file ", err)
	}
	conf := apiserver.GetConfig()
	assert.NotEmpty(t, conf)
}
