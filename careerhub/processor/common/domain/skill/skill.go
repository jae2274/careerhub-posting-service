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
	// Skill_IdField          = "_id"
	Skill_DefaultNameField = "defaultName"
	Skill_SkillNamesField  = "skillNames"
)

type Skill struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	DefaultName string             `bson:"defaultName"`
	SkillNames  []string           `bson:"skillNames"`
	InsertedAt  time.Time          `bson:"insertedAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

func (*Skill) Collection() string {
	return "skill"
}

func (*Skill) IndexModels() map[string]*mongo.IndexModel {
	skillNamesIndex := fmt.Sprintf("%s_1", Skill_SkillNamesField)
	defaultnameIndex := fmt.Sprintf("%s_1", Skill_DefaultNameField)
	return map[string]*mongo.IndexModel{
		defaultnameIndex: {
			Keys: bson.D{
				{Key: Skill_DefaultNameField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
		skillNamesIndex: {
			Keys: bson.D{
				{Key: Skill_SkillNamesField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
