package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/omekov/sample/pkg/contant"
	"github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

type ctxKey int

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)

func (s *Server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			s.error(w, r, http.StatusBadRequest, contant.ErrNotAuthenticated)
			return
		}
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			s.error(w, r, http.StatusBadRequest, contant.ErrNotAuthenticated)
			return
		}

		u, err := s.Store.JWT.GetClaims(splitted[1])
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

// logRequest - middleware для логирование запросов
func (s *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Formatter = &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@t",
				logrus.FieldKeyMsg:   "@m",
				logrus.FieldKeyLevel: "@l",
			},
		}
		// b, _ := ioutil.ReadAll(r.Body)
		logger := s.Logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
			// "request":     string(b),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)
		start := time.Now()
		// для обработки response до и после
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)
		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}
func (s *Server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

// respond - Обработка успешного ответа
func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// error - Обработка ошибочного ответа
func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	logger := s.Logger.WithFields(logrus.Fields{
		"remote_addr": r.RemoteAddr,
		"request_id":  r.Context().Value(ctxKeyRequestID),
	})
	logger.Infof("%s", err.Error())
	if code == http.StatusForbidden {
		s.respond(w, r, code, Error{Error: contant.ErrIncorrectEmailPassword.Error()})
	} else {
		s.respond(w, r, code, Error{Error: err.Error()})
	}
}

func (s *Server) setHeaderAccessControlAllow(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
