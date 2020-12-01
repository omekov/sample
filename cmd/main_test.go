package main

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/omekov/sample/config"
	"github.com/stretchr/testify/assert"
)

func TestApp_ENV(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file ", err)
	}
	envs := []string{
		config.PORT,
		config.MONGOURI,
		config.MONGOUSERNAME,
		config.MONGOPASSWORD,
		config.MONGONAME,
		config.MONGOCUSTOMERSCOLLECTION,
		config.ACCESSTOKENSECRET,
		config.REFRESHTOKENSECRET,
	}
	for _, e := range envs {
		if config.IsReadyENV(e) == "" {
			t.Errorf("is not ENV - %s", e)
		}
	}
}

func TestGetConfig(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file ", err)
	}
	conf := config.GetMongoConfig()
	assert.NotEmpty(t, conf)
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
