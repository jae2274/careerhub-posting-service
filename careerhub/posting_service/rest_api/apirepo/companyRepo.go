package apirepo

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/company"
	"github.com/jae2274/goutils/optional"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CompanyRepo interface {
	FindByCompanySiteID(ctx context.Context, site, id string) (optional.Optional[SiteCompanySummary], error)
	FindByPrefixCompanyName(ctx context.Context, prefixKeyword string, limit int64) ([]*CompanySummary, error)
}

type CompanyRepoImpl struct {
	col *mongo.Collection
}

func NewCompanyRepo(db *mongo.Database) CompanyRepo {
	col := db.Collection((&company.Company{}).Collection())
	return &CompanyRepoImpl{
		col: col,
	}
}

type SiteCompanySummary struct {
	CompanyName string
	CompanyUrl  *string
	CompanyLogo string
}

func (repo *CompanyRepoImpl) FindByCompanySiteID(ctx context.Context, site, siteCompanyId string) (optional.Optional[SiteCompanySummary], error) {
	option := options.FindOne().SetProjection(bson.D{
		{company.SiteCompanies_NameField, 1},
		{company.SiteCompanies_SiteField, 1},
		{company.SiteCompanies_CompanyIdField, 1},
		{company.SiteCompanies_CompanyUrlField, 1},
		{company.SiteCompanies_CompanyLogoField, 1},
	})
	result := repo.col.FindOne(ctx, bson.M{company.SiteCompanies_SiteField: site, company.SiteCompanies_CompanyIdField: siteCompanyId}, option)

	if resultErr := result.Err(); resultErr != nil {
		if resultErr == mongo.ErrNoDocuments {
			return optional.NewEmptyOptional[SiteCompanySummary](), nil
		}

		return optional.NewEmptyOptional[SiteCompanySummary](), resultErr
	}

	var company *company.Company
	if err := result.Decode(&company); err != nil {
		return optional.NewEmptyOptional[SiteCompanySummary](), err
	}

	for _, siteCompany := range company.SiteCompanies {
		if siteCompany.Site == site && siteCompany.CompanyId == siteCompanyId {
			return optional.NewOptional(&SiteCompanySummary{
				CompanyName: siteCompany.Name,
				CompanyUrl:  siteCompany.CompanyUrl,
				CompanyLogo: siteCompany.CompanyLogo,
			}), nil
		}
	}

	return optional.NewEmptyOptional[SiteCompanySummary](), nil
}

func (repo *CompanyRepoImpl) FindByPrefixCompanyName(ctx context.Context, prefixKeyword string, limit int64) ([]*CompanySummary, error) {
	option := options.Find().SetProjection(bson.D{
		{company.DefaultNameField, 1},
		{company.SiteCompanies_SiteField, 1},
		{company.SiteCompanies_CompanyIdField, 1},
		{company.SiteCompanies_NameField, 1},
	}).SetLimit(limit)
	cursor, err := repo.col.Find(ctx, bson.M{company.DefaultNameField: bson.M{"$regex": prefixKeyword, "$options": "i"}}, option)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var companies []*CompanySummary
	for cursor.Next(ctx) {
		var company *company.Company
		if err := cursor.Decode(&company); err != nil {
			return nil, err
		}

		siteCompanies := make([]SiteCompany, 0, len(company.SiteCompanies))
		for _, siteCompany := range company.SiteCompanies {
			siteCompanies = append(siteCompanies, SiteCompany{
				Site:        siteCompany.Site,
				CompanyId:   siteCompany.CompanyId,
				CompanyName: siteCompany.Name,
			})
		}

		companies = append(companies, &CompanySummary{
			DefaultName:   company.DefaultName,
			SiteCompanies: siteCompanies,
		})
	}

	return companies, nil
	// return []*CompanySummary{}, nil
}

type CompanySummary struct {
	DefaultName   string
	SiteCompanies []SiteCompany
}

type SiteCompany struct {
	Site        string
	CompanyId   string
	CompanyName string
}
