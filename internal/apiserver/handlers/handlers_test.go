package handlers_test

import (
	"net/http"
	"reflect"
	"testing"
)

func TestServer_AuthenticateUser(t *testing.T) {

}

func TestServer_signIn(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		s    *Server
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.signIn(tt.args.w, tt.args.r)
		})
	}
}

func TestServer_signUp(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		s    *Server
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.signUp(tt.args.w, tt.args.r)
		})
	}
}

func TestServer_whoami(t *testing.T) {
	tests := []struct {
		name string
		s    *Server
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.whoami(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.whoami() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_refreshToken(t *testing.T) {
	tests := []struct {
		name string
		s    *Server
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.refreshToken(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.refreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
