package models

// Error ...
type Error struct {
	Error string `json:"error,omitempty" example:"error"`
}

// Token ...
type Token struct {
	AccessToken  string `json:"accessToken,omitempty" example:"jwt-token"`
	Refreshtoken string `json:"refreshToken,omitempty" example:"jwt-token"`
}

// ServerStatus ...
type ServerStatus struct {
	ShutdownStatus string `json:"shutdownStatus,omitempty"`
}

// ServerConfig ...
type ServerConfig struct {
	PORT string
}
