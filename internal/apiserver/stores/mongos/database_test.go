package mongos_test

import (
	"os"
	"regexp"
	"errors"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/omekov/sample/internal/apiserver"
	"github.com/omekov/sample/internal/apiserver/stores/mongos"
	"github.com/omekov/sample/internal/apiserver/stores/mongos/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDatabaseHelper struct {
	mock.Mock
}

func (_m *MockDatabaseHelper) Client() mongos.CustomerRepository {
	ret := _m.Called()

	var r0 mongos.CustomerRepository
	if rf, ok := ret.Get(0).(func() mongos.CustomerRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongos.CustomerRepository)
		}
	}
	return r0
}

func (_m *MockDatabaseHelper) Collection(name string) mongos.Collection {
	ret := _m.Called(name)
	var r0 mongos.Collection
	if rf, ok := ret.Get(0).(func(string) mongos.Collection); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongos.Collection)
		}
	}
	return r0
}

func TestNewDatabase(t *testing.T) {
	re := regexp.MustCompile("^(.*pre-sample)")
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	if err := godotenv.Load(string(rootPath) + `/.env.prod`); err != nil {
		log.Fatal("Error loading .env file ", err)
	}
	conf := apiserver.GetConfig()
	dbClient, err := mongos.NewClient(conf)
	assert.NoError(t, err)
	db := mongos.NewDatabase(conf, dbClient)
	assert.NotEmpty(t, db)
}

func TestStartSession(t *testing.T) {
	var db mongos.Database
	var client mongos.Client
	db = &mocks.DatabaseHelper{}
	client = &mocks.ClientHelper{}
	client.(*mocks.ClientHelper).On("StartSession").Return(nil, errors.New("mocked-error"))
	// db.(*MockDatabaseHelper).On("Client").Return(client)
	_, err := db.Client().StartSession()
	assert.EqualError(t, err, "mocked-error")
}
