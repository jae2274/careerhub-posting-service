package rpcService

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
)

type CategoryService struct {
	categoryRepo *rpcRepo.CategoryRepo
}

func NewCategoryService(categoryRepo *rpcRepo.CategoryRepo) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *CategoryService) RegisterCategories(ctx context.Context, site string, categories []string) error {
	if len(categories) == 0 {
		return nil
	}

	return s.categoryRepo.SaveCategories(ctx, site, categories)
}
