package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/omekov/sample/internal/apiserver/models"
)

func TestHadler_SignIn(t *testing.T) {
	c := models.TestCredential(t)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"username": c.Username,
				"password": c.Password,
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid email",
			payload: map[string]string{
				"username": "invalid",
				"password": c.Password,
			},
			expectedCode: http.StatusForbidden,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"username": c.Username,
				"password": "inv",
			},
			expectedCode: http.StatusForbidden,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			buf := &bytes.Buffer{}
			json.NewEncoder(buf).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/signin", buf)
			handlers.Ser
		})
	}
}
