package mongocfg

import (
	"context"
	"errors"
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

func InitCollections(db *mongo.Database, models ...MongoDBModel) (map[string]*mongo.Collection, error) {
	collections := make(map[string]*mongo.Collection)

	errs := []error{}

	for _, model := range models {
		err := initCollection(db, collections, model)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, terr.Wrap(errors.Join(errs...))
	}

	return collections, nil
}

func initCollection(db *mongo.Database, collections map[string]*mongo.Collection, model MongoDBModel) error {
	col := db.Collection(model.Collection())
	err := CheckIndexViaCollection(col, model)
	if err != nil {
		return err
	}
	collections[model.Collection()] = col

	return nil
}
