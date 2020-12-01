package mongodb

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
	Customer CustomerRepository
}
