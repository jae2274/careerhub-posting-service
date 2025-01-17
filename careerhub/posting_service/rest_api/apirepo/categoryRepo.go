package apirepo

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/category"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryRepo interface {
	GetAllCategories(ctx context.Context) ([]*category.Category, error)
}

type CategoryRepoImpl struct {
	col *mongo.Collection
}

func NewCategoryRepo(db *mongo.Database) CategoryRepo {
	categoryCollection := db.Collection((&category.Category{}).Collection())
	return &CategoryRepoImpl{
		col: categoryCollection,
	}
}

func (repo *CategoryRepoImpl) GetAllCategories(ctx context.Context) ([]*category.Category, error) {
	options := options.Find().SetSort(bson.M{category.PriorityField: -1})
	cursor, err := repo.col.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}

	var categories []*category.Category
	err = cursor.All(ctx, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
