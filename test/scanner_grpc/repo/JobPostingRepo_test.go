package repo

import (
	"context"
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/test/testutils"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestJobPostingRepo(t *testing.T) {
	providerRepo := tinit.InitProviderJobPostingRepo(t)
	scannerRepo := tinit.InitScannerJobPostingRepo(t)

	savedJobPosting1 := testutils.CreateJobPosting("jumpit", "1", []string{"java", "python", "go"})
	savedJobPosting2 := testutils.CreateJobPosting("jumpit", "2", []string{"javascript", "react"})
	savedJobPosting3 := testutils.CreateJobPosting("jumpit", "3", []string{"aws", "gcp", "azure"})
	savedJobPostings := []*jobposting.JobPostingInfo{savedJobPosting1, savedJobPosting2, savedJobPosting3}

	for _, jobPosting := range savedJobPostings {
		isSuccess, err := providerRepo.Save(context.Background(), jobPosting)
		require.NoError(t, err)
		require.True(t, isSuccess)
	}

	jobPostingChan, err := scannerRepo.GetJobPostings(context.Background(), false)
	require.NoError(t, err)
	index := 0
	for jobPosting := range jobPostingChan {
		testutils.SetIgnoreJobPostingFields([]*jobposting.JobPostingInfo{jobPosting, savedJobPostings[index]})
		require.Equal(t, *jobPosting, *savedJobPostings[index])
		index++
	}

	jobPostingChan, err = scannerRepo.GetJobPostings(context.Background(), true)
	require.NoError(t, err)
	_, ok := <-jobPostingChan
	require.False(t, ok)

	err = scannerRepo.AddRequiredSkills(context.Background(), savedJobPosting1.JobPostingId, []string{"kotlin", "swift"})
	require.NoError(t, err)

	jobPostingChan, err = scannerRepo.GetJobPostings(context.Background(), false)
	require.NoError(t, err)
	index = 1
	for jobPosting := range jobPostingChan {
		testutils.SetIgnoreJobPostingFields([]*jobposting.JobPostingInfo{jobPosting, savedJobPostings[index]})
		require.Equal(t, *jobPosting, *savedJobPostings[index])
		index++
	}

	jobPostingChan, err = scannerRepo.GetJobPostings(context.Background(), true)
	require.NoError(t, err)
	jobPosting, ok := <-jobPostingChan
	require.True(t, ok)
	require.Equal(t, savedJobPosting1.JobPostingId, jobPosting.JobPostingId)
	require.True(t, jobPosting.IsScanComplete)
	require.Equal(t, savedJobPosting1.RequiredSkill, jobPosting.RequiredSkill[0:3])
	require.Equal(t, jobposting.RequiredSkill{SkillFrom: jobposting.Scanned, SkillName: "kotlin"}, jobPosting.RequiredSkill[3])
	require.Equal(t, jobposting.RequiredSkill{SkillFrom: jobposting.Scanned, SkillName: "swift"}, jobPosting.RequiredSkill[4])

}