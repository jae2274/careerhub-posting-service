package apirepo

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/company"
	"github.com/jae2274/goutils/optional"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepo interface {
	FindByCompanySiteID(ctx context.Context, site, id string) (optional.Optional[CompanySummary], error)
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

type CompanySummary struct {
	CompanyName string
	CompanyUrl  *string
	CompanyLogo string
}

func (repo *CompanyRepoImpl) FindByCompanySiteID(ctx context.Context, site, siteCompanyId string) (optional.Optional[CompanySummary], error) {
	result := repo.col.FindOne(ctx, bson.M{company.SiteCompanies_SiteField: site, company.SiteCompanies_CompanyIdField: siteCompanyId})

	if resultErr := result.Err(); resultErr != nil {
		if resultErr == mongo.ErrNoDocuments {
			return optional.NewEmptyOptional[CompanySummary](), nil
		}

		return optional.NewEmptyOptional[CompanySummary](), resultErr
	}

	var company *company.Company
	if err := result.Decode(&company); err != nil {
		return optional.NewEmptyOptional[CompanySummary](), err
	}

	for _, siteCompany := range company.SiteCompanies {
		if siteCompany.Site == site && siteCompany.CompanyId == siteCompanyId {
			return optional.NewOptional(&CompanySummary{
				CompanyName: siteCompany.Name,
				CompanyUrl:  siteCompany.CompanyUrl,
				CompanyLogo: siteCompany.CompanyLogo,
			}), nil
		}
	}

	return optional.NewEmptyOptional[CompanySummary](), nil
}
