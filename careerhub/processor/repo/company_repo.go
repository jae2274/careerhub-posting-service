package repo

import (
	"context"
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepo struct {
	col *mongo.Collection
}

func NewCompanyRepo(col *mongo.Collection) *CompanyRepo {
	return &CompanyRepo{
		col: col,
	}
}

func (cRepo *CompanyRepo) Save(ctx context.Context, company *company.SiteCompany) (bool, error) {
	company.InsertedAt = time.Now()
	company.UpdatedAt = time.Now()

	_, err := cRepo.col.InsertOne(ctx, company)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) { // Ignore duplicate key error
			return false, nil
		}
		return false, err
	}

	return true, nil
}
