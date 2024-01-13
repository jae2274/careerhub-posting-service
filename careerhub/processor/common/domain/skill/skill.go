package skill

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SkillName_NameField = "skillNames.name"
)

type Skill struct {
	ID          string       `bson:"id"`
	DefaultName string       `bson:"defaultName"`
	SkillNames  []*SkillName `bson:"skillNames"`
}

type SkillName struct {
	Name string `bson:"name"`
}

func (*Skill) Collection() string {
	return "skill"
}

func (*Skill) IndexModels() map[string]*mongo.IndexModel {
	keyName := fmt.Sprintf("%s_1", SkillName_NameField)
	return map[string]*mongo.IndexModel{
		keyName: {
			Keys: bson.D{
				{Key: SkillName_NameField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
