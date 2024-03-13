package restapi

import (
	"context"
	"testing"

	restapi "github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestRestApiService(t *testing.T) {

	t.Run("GetAllCategories", func(t *testing.T) {
		provCategoryRepo := tinit.InitProviderCategoryRepo(t)
		jobPostingRepo := tinit.InitRestApiJobPostingRepo(t)
		restApiCategoryRepo := tinit.InitRestApiCategoryRepo(t)
		skillNameRepo := tinit.InitRestApiSkillNameRepo(t)
		skillRepo := tinit.InitRestApiSkillRepo(t)

		restApiService := restapi.NewRestApiService(jobPostingRepo, restApiCategoryRepo, skillNameRepo, skillRepo)

		mainCtx := context.Background()

		willSavedCategories := map[string][]string{
			"jumpit": {"서버/백엔드", "프론트"},
			"wanted": {"백엔드 개발자", "프론트 개발자"},
		}

		for site, categories := range willSavedCategories {
			err := provCategoryRepo.SaveCategories(mainCtx, site, categories)
			require.NoError(t, err)
		}

		// Test GetAllCategories

		categories, err := restApiService.GetAllCategories(mainCtx)
		require.NoError(t, err)

		require.Len(t, categories.CategoriesBySite, len(willSavedCategories))

		for _, categoryBySite := range categories.CategoriesBySite {
			require.Equal(t, willSavedCategories[categoryBySite.Site], categoryBySite.Categories)
		}

	})
}
