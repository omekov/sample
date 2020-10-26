package apiserver

import (
	"reflect"
	"testing"

	"github.com/omekov/sample/internal/apiserver/models"
)

func TestIsReadyENV(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsReadyENV(tt.args.key); got != tt.want {
				t.Errorf("IsReadyENV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagAndLoadENV(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FlagAndLoadENV()
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Run()
		})
	}
}

func TestGetMongoConfig(t *testing.T) {
	tests := []struct {
		name string
		want *models.MongoConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMongoConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMongoConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRedisConfig(t *testing.T) {
	tests := []struct {
		name string
		want *models.RedisConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRedisConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newServer(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := newServer(); (err != nil) != tt.wantErr {
				t.Errorf("newServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
