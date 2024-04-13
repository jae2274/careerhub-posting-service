package restapi

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_server"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestRestApiService(t *testing.T) {

	t.Run("GetAllCategories", func(t *testing.T) {
		db := tinit.InitDB(t)
		provCategoryRepo := rpcRepo.NewCategoryRepo(db)
		restApiCategoryRepo := apirepo.NewCategoryRepo(db)
		// skillNameRepo := tinit.InitRestApiSkillNameRepo(t)
		restApiService := restapi_server.NewRestApiService(nil, restApiCategoryRepo, nil)

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

		categories, err := restApiService.Categories(mainCtx, &emptypb.Empty{})
		require.NoError(t, err)

		require.Len(t, categories.CategoriesBySite, len(willSavedCategories))

		for _, categoryBySite := range categories.CategoriesBySite {
			require.Equal(t, willSavedCategories[categoryBySite.Site], categoryBySite.Categories)
		}

	})
}
