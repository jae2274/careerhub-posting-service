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

	// dialer := &net.Dialer{
	// 	Timeout:   30 * time.Second,
	// 	KeepAlive: 300 * time.Second,
	// }
	clientOptions := options.Client().ApplyURI(uri)
	//해당 시간이 지나면 커넥션을 닫는다. 설정하지 않을 시 broken connection pipe의 에러 원인으로 추정된다.
	//시간을 수정하며 동작에 어떤 영향을 미치는지 파악하자
	clientOptions.SetMaxConnIdleTime(10 * time.Second)
	// clientOptions.SetDialer(dialer)

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

func InitCollections(db *mongo.Database, models ...MongoDBModel) error {
	errs := []error{}

	for _, model := range models {
		err := initCollection(db, model)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return terr.Wrap(errors.Join(errs...))
	}

	return nil
}

func initCollection(db *mongo.Database, model MongoDBModel) error {
	col := db.Collection(model.Collection())
	err := CheckIndexViaCollection(col, model)
	if err != nil {
		return err
	}

	return nil
}
