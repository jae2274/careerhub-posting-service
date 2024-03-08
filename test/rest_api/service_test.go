package restapi

import (
	"context"
	"slices"
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	restapi "github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/dto"
	"github.com/jae2274/Careerhub-dataProcessor/test/testutils"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestRestApiService(t *testing.T) {
	t.Run("GetJobPostings", func(t *testing.T) {
		providerRepo := tinit.InitProviderJobPostingRepo(t)
		jobPostingRepo := tinit.InitRestApiJobPostingRepo(t)
		categoryRepo := tinit.InitRestApiCategoryRepo(t)

		restApiService := restapi.NewRestApiService(jobPostingRepo, categoryRepo)

		mainCtx := context.Background()
		savedJobPostings := []*jobposting.JobPostingInfo{
			testutils.CreateJobPosting("jumpit", "1", []string{"java", "python", "go"}),
			testutils.CreateJobPosting("jumpit", "2", []string{"javascript", "react"}),
			testutils.CreateJobPosting("jumpit", "3", []string{"aws", "gcp", "azure"}),
		}

		for _, jobPosting := range savedJobPostings {
			isSuccess, err := providerRepo.Save(mainCtx, jobPosting)
			require.NoError(t, err)
			require.True(t, isSuccess)
		}

		// Test GetJobPostings
		req := &dto.GetJobPostingsRequest{Page: 0, Size: 10}
		jobPostings, err := restApiService.GetJobPostings(mainCtx, req)
		require.NoError(t, err)
		require.Len(t, jobPostings, 3)
		slices.Reverse(jobPostings) //조회는 createdAt 기준으로 내림차순 정렬되어야 한다.

		for index, jobPosting := range jobPostings {
			require.Equal(t, savedJobPostings[index].JobPostingId.Site, jobPosting.Site)
			require.Equal(t, savedJobPostings[index].JobPostingId.PostingId, jobPosting.PostingId)
			require.Equal(t, savedJobPostings[index].MainContent.Title, jobPosting.Title)

			savedRequiredSkill := make([]string, len(savedJobPostings[index].RequiredSkill))
			for i, skill := range savedJobPostings[index].RequiredSkill {
				savedRequiredSkill[i] = skill.SkillName
			}
			require.Equal(t, savedRequiredSkill, jobPosting.Skills)
		}
	})

	t.Run("GetAllCategories", func(t *testing.T) {
		provCategoryRepo := tinit.InitProviderCategoryRepo(t)
		restApiCategoryRepo := tinit.InitRestApiCategoryRepo(t)
		restApiService := restapi.NewRestApiService(nil, restApiCategoryRepo)

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
