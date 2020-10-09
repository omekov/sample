package main

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	api "github.com/omekov/sample/internal/apiserver"
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
