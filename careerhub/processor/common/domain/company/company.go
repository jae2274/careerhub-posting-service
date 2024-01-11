package company

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	IdField                      = "_id"
	DefaultNameField             = "defaultName"
	SiteCompaniesField           = "siteCompanies"
	SiteCompanies_SiteField      = "siteCompanies.site"
	SiteCompanies_CompanyIdField = "siteCompanies.companyId"
)

type Company struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	DefaultName   string             `bson:"defaultName"`
	SiteCompanies []*SiteCompany     `bson:"siteCompanies"`
	InsertedAt    time.Time          `bson:"insertedAt"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
}

type SiteCompany struct {
	Site          string    `bson:"site"`
	CompanyId     string    `bson:"companyId"`
	Name          string    `bson:"name"`
	CompanyUrl    *string   `bson:"companyUrl"`
	CompanyImages []string  `bson:"companyImages"`
	Description   string    `bson:"description"`
	CompanyLogo   string    `bson:"companyLogo"`
	CreatedAt     time.Time `bson:"createdAt"`
	InsertedAt    time.Time `bson:"insertedAt"`
	UpdatedAt     time.Time `bson:"updatedAt"`
}

func (*Company) Collection() string {
	return "company"
}

func (*Company) IndexModels() map[string]*mongo.IndexModel {
	keyName := fmt.Sprintf("%s_1_%s_1", SiteCompanies_SiteField, SiteCompanies_CompanyIdField)
	return map[string]*mongo.IndexModel{
		keyName: {
			Keys: bson.D{
				{Key: SiteCompanies_SiteField, Value: 1},
				{Key: SiteCompanies_CompanyIdField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}