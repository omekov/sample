package mongos

import "github.com/omekov/sample/internal/apiserver/stores/mongos/customer"

// Collections ...
type Collections struct {
	Customer string
	Roles    string
}

// MongoConfig ...
type MongoConfig struct {
	Username     string
	Password     string
	DatabaseName string
	URL          string
	Collections  Collections
}

// MongoDBRepositories ...
type MongoDBRepositories struct {
	Customer customer.CustomerRepository
}
