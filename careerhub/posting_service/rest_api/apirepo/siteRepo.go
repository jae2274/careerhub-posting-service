package apirepo

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/site"
	"github.com/jae2274/goutils/optional"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SiteRepo interface {
	FindBySiteName(ctx context.Context, siteName string) (optional.Optional[site.Site], error)
}

type siteRepoImpl struct {
	col *mongo.Collection
}

func NewSiteRepo(db *mongo.Database) SiteRepo {
	return &siteRepoImpl{
		col: db.Collection((&site.Site{}).Collection()),
	}
}

func (repo *siteRepoImpl) FindBySiteName(ctx context.Context, siteName string) (optional.Optional[site.Site], error) {
	result := repo.col.FindOne(ctx, bson.M{site.SiteNameField: siteName})

	if resultErr := result.Err(); resultErr != nil {
		if resultErr == mongo.ErrNoDocuments {
			return optional.NewEmptyOptional[site.Site](), nil
		}

		return optional.NewEmptyOptional[site.Site](), resultErr
	}

	var siteObj *site.Site
	if err := result.Decode(&siteObj); err != nil {
		return optional.NewEmptyOptional[site.Site](), err
	}

	return optional.NewOptional(siteObj), nil
}
