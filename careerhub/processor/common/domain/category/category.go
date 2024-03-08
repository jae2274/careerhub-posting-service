package category

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	IdField       = "_id"
	SiteField     = "site"
	NameField     = "name"
	PriorityField = "priority"
)

type Category struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Site       string             `bson:"site"`
	Name       string             `bson:"name"`
	Priority   int                `bson:"priority"` //값이 클수록 노출 우선순위가 높음
	InsertedAt time.Time          `bson:"insertedAt"`
	UpdatedAt  time.Time          `bson:"updatedAt"`
}

func (*Category) Collection() string {
	return "category"
}

func (*Category) IndexModels() map[string]*mongo.IndexModel {
	keyName := fmt.Sprintf("%s_1_%s_1", SiteField, NameField)
	return map[string]*mongo.IndexModel{
		keyName: {
			Keys: bson.D{
				{Key: SiteField, Value: 1},
				{Key: NameField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
