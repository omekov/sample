package handlers

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_responseWriter_WriteHeader(t *testing.T) {
	type args struct {
		statusCode int
	}
	tests := []struct {
		name string
		w    *responseWriter
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.WriteHeader(tt.args.statusCode)
		})
	}
}

func TestServer_authenticateUser(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.authenticateUser(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.authenticateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_logRequest(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.logRequest(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.logRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_setRequestID(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name string
		s    *Server
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.setRequestID(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.setRequestID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_respond(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		r    *http.Request
		code int
		data interface{}
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
			tt.s.respond(tt.args.w, tt.args.r, tt.args.code, tt.args.data)
		})
	}
}

func TestServer_error(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		r    *http.Request
		code int
		err  error
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
			tt.s.error(tt.args.w, tt.args.r, tt.args.code, tt.args.err)
		})
	}
}
