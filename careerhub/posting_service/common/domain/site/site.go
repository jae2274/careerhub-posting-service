package site

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SiteNameField         = "siteName"
	PostingUrlFormatField = "postingUrlFormat"
)

type Site struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	SiteName         string             `bson:"siteName"`
	PostingUrlFormat string             `bson:"postingUrlFormat"`
}

func (*Site) Collection() string {
	return "site"
}

func (*Site) IndexModels() map[string]*mongo.IndexModel {
	keyName := fmt.Sprintf("%s_1", SiteNameField)
	return map[string]*mongo.IndexModel{
		keyName: {
			Keys: bson.D{
				{Key: SiteNameField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
