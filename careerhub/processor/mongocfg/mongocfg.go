package mongocfg

import (
	"context"
	"time"

	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabase(uri string, dbName string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, terr.Wrap(err)
	}

	db := client.Database(dbName)
	var result bson.M
	if err := db.RunCommand(ctx, bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, terr.Wrap(err)
	}

	return db, nil
}
