package handlers

import (
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/omekov/sample/docs"
)

func TestServer_ConfigureRouter(t *testing.T) {
	type args struct {
		PORT string
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want *mux.Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ConfigureRouter(tt.args.PORT); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ConfigureRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
