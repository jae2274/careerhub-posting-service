package skill

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SkillName_IdField             = "_id"
	SkillName_NameField           = "name"
	SkillName_IsScanCompleteField = "isScanComplete"
)

type SkillName struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Name           string             `bson:"name"`
	InsertedAt     time.Time          `bson:"insertedAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
	IsScanComplete bool               `bson:"isScanComplete"`
}

func (*SkillName) Collection() string {
	return "skillName"
}

func (*SkillName) IndexModels() map[string]*mongo.IndexModel {
	nameIndex := fmt.Sprintf("%s_1", SkillName_NameField)
	return map[string]*mongo.IndexModel{
		nameIndex: {
			Keys: bson.D{
				{Key: SkillName_NameField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
