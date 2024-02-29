package jobposting

import (
	"fmt"
	"time"

	"github.com/jae2274/goutils/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SiteField           = "jobPostingId.site"
	PostingIdField      = "jobPostingId.postingId"
	StatusField         = "status"
	IsScanCompleteField = "isScanComplete"
	RequiredSkillField  = "requiredSkill"
	CreatedAtField      = "createdAt"
)

type StatusValues struct{}

type Status = enum.Enum[StatusValues]

const (
	HIRING = Status("hiring")
	CLOSED = Status("closed")
)

func (StatusValues) Values() []string {
	return []string{string(HIRING), string(CLOSED)}
}

func (StatusValues) ParseStatus(s string) (Status, error) {
	switch s {
	case string(HIRING):
		return HIRING, nil
	case string(CLOSED):
		return CLOSED, nil
	default:
		return "", fmt.Errorf("invalid status: %s", s)
	}
}

type JobPostingId struct {
	Site      string `bson:"site"`
	PostingId string `bson:"postingId"`
}

type JobPostingInfo struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	JobPostingId   JobPostingId       `bson:"jobPostingId"`
	Status         Status             `bson:"status"`
	CompanyId      string             `bson:"companyId"`
	CompanyName    string             `bson:"companyName"`
	JobCategory    []string           `bson:"jobCategory"`
	ImageUrl       *string            `bson:"imageUrl"`
	MainContent    MainContent        `bson:"mainContent"`
	RequiredSkill  []RequiredSkill    `bson:"requiredSkill"`
	Tags           []string           `bson:"tags"`
	RequiredCareer Career             `bson:"requiredCareer"`
	PublishedAt    *time.Time         `bson:"publishedAt,omitempty"`
	ClosedAt       *time.Time         `bson:"closedAt,omitempty"`
	Address        []string           `bson:"address"`
	CreatedAt      time.Time          `bson:"createdAt"`
	InsertedAt     time.Time          `bson:"insertedAt"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
	IsScanComplete bool               `bson:"isScanComplete"`
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

type SkillFromValues struct{}
type SkillFrom = enum.Enum[SkillFromValues]

const (
	Scanned = SkillFrom("SCANNED")
	Origin  = SkillFrom("ORIGIN")
)

func (SkillFromValues) Values() []string {
	return []string{string(Scanned), string(Origin)}
}

type RequiredSkill struct {
	SkillFrom SkillFrom `bson:"skillFrom"`
	SkillName string    `bson:"skillName"`
}

type Career struct {
	Min *int32 `bson:"min,omitempty"`
	Max *int32 `bson:"max,omitempty"`
}

func (*JobPostingInfo) Collection() string {
	return "jobPostingInfo"
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
