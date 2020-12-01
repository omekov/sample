package mongos

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database ...
type Database interface {
	Collection(name string) Collection
	Client() Client
}

// Collection ...
type Collection interface {
	FindOne(context.Context, interface{}) SingleResult
	FindOneAndUpdate(context.Context, interface{}, interface{}) SingleResult
	InsertOne(context.Context, interface{}) (*mongo.InsertOneResult, error)
	DeleteOne(ctx context.Context, filter interface{}) (int64, error)
	UpdateOne(ctx context.Context, id primitive.ObjectID, update interface{}) (*mongo.UpdateResult, error)
}

// SingleResult ...
type SingleResult interface {
	Decode(v interface{}) error
}

// Client ...
type Client interface {
	Database(string) Database
	Connect() error
	StartSession() (mongo.Session, error)
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoSession struct {
	mongo.Session
}

// NewClient ...
func NewClient(cnf *MongoConfig) (Client, error) {
	clientOptions := options.Client().SetAuth(
		options.Credential{
			Username:   cnf.Username,
			Password:   cnf.Password,
			AuthSource: cnf.DatabaseName,
		}).ApplyURI(cnf.URL).SetRetryWrites(false)
	c, err := mongo.NewClient(clientOptions)
	return &mongoClient{cl: c}, err

}

// NewDatabase ...
func NewDatabase(cnf *MongoConfig, client Client) Database {
	return client.Database(cnf.DatabaseName)
}

func (mc *mongoClient) Database(dbName string) Database {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Connect() error {
	// mongo client does not use context on connect method. There is a ticket
	// with a request to deprecate this functionality and another one with
	// explanation why it could be useful in synchronous requests.
	// https://jira.mongodb.org/browse/GODRIVER-1031
	// https://jira.mongodb.org/browse/GODRIVER-979
	return mc.cl.Connect(nil)
}

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, id primitive.ObjectID, update interface{}) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateOne(ctx, bson.M{"_id": id}, update)
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return mc.coll.InsertOne(ctx, document)
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (mc *mongoCollection) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}) SingleResult {
	singleResult := mc.coll.FindOneAndUpdate(ctx, filter, update)
	return &mongoSingleResult{sr: singleResult}
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}
