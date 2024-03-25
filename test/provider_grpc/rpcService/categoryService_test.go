package rpcService

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/category"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcService"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestCategoryService(t *testing.T) {
	t.Run("RegisterCategories", func(t *testing.T) {
		categoryRepo := tinit.InitProviderCategoryRepo(t)
		ctx := context.TODO()
		type SiteCategory struct {
			Site       string
			Categories []string
		}

		willSavedCategories := []SiteCategory{
			{Site: "jumpit", Categories: []string{"서버/백엔드", "프론트엔드"}},
			{Site: "wanted", Categories: []string{"서버 개발자", "웹 개발자"}},
			{Site: "jumpit", Categories: []string{"서버/백엔드", "DevOps"}},
			{Site: "wanted", Categories: []string{"프론트엔드 개발자", "웹 개발자", "프론트엔드"}},
		}

		categoryService := rpcService.NewCategoryService(categoryRepo)

		for _, siteCategory := range willSavedCategories {
			err := categoryService.RegisterCategories(ctx, siteCategory.Site, siteCategory.Categories)
			require.NoError(t, err)
		}

		categories, err := categoryRepo.FindAll(ctx)
		require.NoError(t, err)

		expectedCategories := []category.Category{ //site와 name이 같은 것은 중복으로 처리되어 저장되지 않음
			{Site: "jumpit", Name: "서버/백엔드"},
			{Site: "jumpit", Name: "프론트엔드"},
			{Site: "wanted", Name: "서버 개발자"},
			{Site: "wanted", Name: "웹 개발자"},
			{Site: "jumpit", Name: "DevOps"},
			{Site: "wanted", Name: "프론트엔드 개발자"},
			{Site: "wanted", Name: "프론트엔드"},
		}

		require.Equal(t, len(expectedCategories), len(categories))
		for i, category := range categories {
			require.Equal(t, expectedCategories[i].Site, category.Site)
			require.Equal(t, expectedCategories[i].Name, category.Name)
		}
	})
}
