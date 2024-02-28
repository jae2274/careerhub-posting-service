package restapi

import (
	"context"
	"slices"
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	restapi "github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api"
	"github.com/jae2274/Careerhub-dataProcessor/test/testutils"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestRestApiService(t *testing.T) {
	providerRepo := tinit.InitProviderJobPostingRepo(t)
	restApiRepo := tinit.InitRestApiJobPostingRepo(t)

	restApiService := restapi.NewRestApiService(restApiRepo)

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

	jobPostings, err := restApiService.GetJobPostings(mainCtx, 0, 10)
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
}
