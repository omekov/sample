package mongodb

/*
import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"

)

// mongoDBConnect -
func mongoDBConnect(ctx context.Context, cnf *configs.Mongo) (*Database, error) {
	dbClient, err := NewClient(cnf)
	if err != nil {
		return nil, err
	}
	if err = dbClient.Connect(); err != nil {
		return nil, err
	}
	// defer dbClient.DisConnect(ctx)
	if err := dbClient.Ping(ctx); err != nil {
		return nil, err
	}
	log.Infof("MongoDB PING - %s", "PONG")
	db := NewDatabase(cnf.Name, dbClient)
	return &db, nil
}

func (store *Store) keepAliveMongo(ctx context.Context, cnf *configs.Mongo) {
	var err error
	for {
		time.Sleep(time.Second * KeepAlivePollPeriod)
		lostConnect := false

		if store.Mongo != nil {
			lostConnect = true
		} else if err = store.Mongo.Client().Ping(ctx); err != nil {
			lostConnect = true
		}
		if !lostConnect {
			continue
		}
		log.Debug("[store.KeepAliveMongo] Lost MongoDB connection. Restoring...")
		_, err = mongoDBConnect(ctx, cnf)
		if err != nil {
			log.Error(err)
			continue
		}
		log.Debug("[store.KeepAlivePg] PostgreSQL reconnected")
	}

}

*/