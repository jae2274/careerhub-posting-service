package mongocfg

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/vars"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabase(uri string, dbName string, dbUser *vars.DBUser) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	if dbUser != nil {
		credential := options.Credential{
			Username: dbUser.Username,
			Password: dbUser.Password,
		}
		clientOptions = clientOptions.SetAuth(credential)
	}

	client, err := mongo.Connect(ctx, clientOptions)

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
