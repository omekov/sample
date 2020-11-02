package models

// MongoConfig ...
type MongoConfig struct {
	Username     string
	Password     string
	DatabaseName string
	URL          string
}

// RedisConfig ...
type RedisConfig struct {
	Password string
	URL      string
}

// Error ...
type Error struct {
	Error string `json:"error,omitempty" example:"error"`
}

// Token ...
type Token struct {
	AccessToken  string `json:"accessToken,omitempty" example:"jwt-token"`
	Refreshtoken string `json:"refreshToken,omitempty" example:"jwt-token"`
}
