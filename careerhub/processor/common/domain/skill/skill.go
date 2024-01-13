package skill

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	IdField             = "_id"
	DefaultNameField    = "defaultName"
	SkillName_NameField = "skillNames.name"
)

type Skill struct {
	ID          string       `bson:"id"`
	DefaultName string       `bson:"defaultName"`
	SkillNames  []*SkillName `bson:"skillNames"`
	InsertedAt  time.Time    `bson:"insertedAt"`
	UpdatedAt   time.Time    `bson:"updatedAt"`
}

type SkillName struct {
	Name       string    `bson:"name"`
	InsertedAt time.Time `bson:"insertedAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
}

func (*Skill) Collection() string {
	return "skill"
}

func (*Skill) IndexModels() map[string]*mongo.IndexModel {
	skillNameIndex := fmt.Sprintf("%s_1", SkillName_NameField)
	defaultnameIndex := fmt.Sprintf("%s_1", DefaultNameField)
	return map[string]*mongo.IndexModel{
		defaultnameIndex: {
			Keys: bson.D{
				{Key: DefaultNameField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
		skillNameIndex: {
			Keys: bson.D{
				{Key: SkillName_NameField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
