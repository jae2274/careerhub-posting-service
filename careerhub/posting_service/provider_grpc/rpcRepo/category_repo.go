package rpcRepo

import (
	"context"
	"time"

	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/category"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CategoryRepo struct {
	col *mongo.Collection
}

func NewCategoryRepo(col *mongo.Collection) *CategoryRepo {
	return &CategoryRepo{
		col: col,
	}
}

func (cRepo *CategoryRepo) SaveCategories(ctx context.Context, site string, categoryNames []string) error {
	now := time.Now()

	categorys := make([]any, len(categoryNames))
	for i, categoryName := range categoryNames {
		categorys[i] = &category.Category{
			Name:       categoryName,
			Site:       site,
			InsertedAt: now,
			UpdatedAt:  now,
		}
	}

	opts := options.InsertMany().SetOrdered(false) // 중복되는 데이터가 있어도 에러를 내지 않고 나머지 데이터를 저장하도록 설정
	_, err := cRepo.col.InsertMany(ctx, categorys, opts)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil
		}
		return err
	}

	return nil
}

func (cRepo *CategoryRepo) FindAll(ctx context.Context) ([]*category.Category, error) {
	var companies []*category.Category

	cursor, err := cRepo.col.Find(ctx, bson.D{})
	if err != nil {
		if mongo.ErrNilDocument == err {
			return []*category.Category{}, nil
		}
		return nil, err
	}

	if err := cursor.All(context.Background(), &companies); err != nil {
		return nil, err
	}

	return companies, nil
}
