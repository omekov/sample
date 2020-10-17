package models


// MongoConfig ...
type MongoConfig struct {
	Username     string
	Password     string
	DatabaseName string
	URL          string
}


// Error ...
type Error struct {
	Error string `json:"error,omitempty" example:"error"`
}
