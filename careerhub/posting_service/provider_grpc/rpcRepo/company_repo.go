package rpcRepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/company"
	"github.com/jae2274/goutils/llog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CompanyRepo struct {
	col *mongo.Collection
}

func NewCompanyRepo(col *mongo.Collection) *CompanyRepo {
	return &CompanyRepo{
		col: col,
	}
}

func (cRepo *CompanyRepo) FindIDByName(ctx context.Context, companyName string) (*primitive.ObjectID, error) {
	var result struct {
		ID primitive.ObjectID `bson:"_id"`
	}

	opts := options.FindOne().SetProjection(bson.D{{Key: company.IdField, Value: 1}})
	err := cRepo.col.FindOne(ctx, bson.M{company.DefaultNameField: companyName}, opts).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &result.ID, nil
}

func (cRepo *CompanyRepo) InsertCompany(ctx context.Context, company *company.Company) (bool, error) {
	now := time.Now()
	company.InsertedAt = now
	company.UpdatedAt = now

	for _, siteCompany := range company.SiteCompanies {
		siteCompany.InsertedAt = now
		siteCompany.UpdatedAt = now
	}

	result, err := cRepo.col.InsertOne(ctx, company)

	if err != nil {
		return false, err
	}

	if result.InsertedID == nil {
		return false, fmt.Errorf("no document was inserted")
	}

	return true, nil
}

func (cRepo *CompanyRepo) IsExisted(ctx context.Context, site, companyId string) (bool, error) {
	count, err := cRepo.col.CountDocuments(ctx, bson.M{company.SiteCompaniesField: bson.M{"$elemMatch": bson.M{company.SiteField: site, company.CompanyIdField: companyId}}})
	if err != nil {
		return false, err
	}

	if count > 1 {
		llog.Level(llog.ERROR).Data("site", site).Data("companyId", companyId).Log(ctx)
		log.Fatal("company count is greater than 1")
	}

	return count > 0, nil
}

func (cRepo *CompanyRepo) AppendSiteCompany(ctx context.Context, companyId primitive.ObjectID, siteCompany *company.SiteCompany) (bool, error) {
	now := time.Now()
	siteCompany.InsertedAt = now
	siteCompany.UpdatedAt = now

	result, err := cRepo.col.UpdateByID(ctx, companyId, bson.M{
		"$push": bson.M{
			company.SiteCompaniesField: siteCompany,
		},
	})

	if err != nil {
		return false, err
	}

	if result.ModifiedCount == 0 {
		return false, fmt.Errorf("no document was modified. CompanyId: %s", companyId.Hex())
	}

	return true, nil
}

func (cRepo *CompanyRepo) FindAll() ([]*company.Company, error) {
	var companies []*company.Company

	cursor, err := cRepo.col.Find(context.Background(), bson.D{})
	if err != nil {
		if mongo.ErrNilDocument == err {
			return []*company.Company{}, nil
		}
		return nil, err
	}

	if err := cursor.All(context.Background(), &companies); err != nil {
		return nil, err
	}

	return companies, nil
}
