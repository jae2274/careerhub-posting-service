package jobposting

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SiteField      = "site"
	PostingIdField = "postingId"
)

type JobPostingInfo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Site           string             `bson:"site"`
	PostingId      string             `bson:"postingId"`
	CompanyId      string             `bson:"companyId"`
	CompanyName    string             `bson:"companyName"`
	JobCategory    []string           `bson:"jobCategory"`
	MainContent    MainContent        `bson:"mainContent"`
	RequiredSkill  []string           `bson:"requiredSkill"`
	Tags           []string           `bson:"tags"`
	RequiredCareer Career             `bson:"requiredCareer"`
	PublishedAt    *int64             `bson:"publishedAt,omitempty"`
	ClosedAt       *int64             `bson:"closedAt,omitempty"`
	Address        []string           `bson:"address"`
}

type MainContent struct {
	PostUrl        string  `bson:"postUrl"`
	Title          string  `bson:"title"`
	Intro          string  `bson:"intro"`
	MainTask       string  `bson:"mainTask"`
	Qualifications string  `bson:"qualifications"`
	Preferred      string  `bson:"preferred"`
	Benefits       string  `bson:"benefits"`
	RecruitProcess *string `bson:"recruitProcess,omitempty"`
}

type Career struct {
	Min *int32 `bson:"min,omitempty"`
	Max *int32 `bson:"max,omitempty"`
}

func (*JobPostingInfo) Collection() string {
	return "deck"
}

func (*JobPostingInfo) IndexModels() map[string]*mongo.IndexModel {
	keyName := fmt.Sprintf("%s_1_%s_1", SiteField, PostingIdField)
	return map[string]*mongo.IndexModel{
		keyName: {
			Keys: bson.D{
				{Key: SiteField, Value: 1},
				{Key: PostingIdField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
