package customers

import "go.mongodb.org/mongo-driver/mongo"

type Customers struct {
	collection mongo.Collection
}
