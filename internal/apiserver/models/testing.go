package models

import "testing"

// TestCredential ...
func TestCredential(t *testing.T) *Credential {
	return &Credential{
		Username: "admin@example.org",
		Password: "123123",
	}
}
