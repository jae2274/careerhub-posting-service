package apirepo_test

import (
	"context"
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/dto"
	"github.com/jae2274/Careerhub-dataProcessor/test/testutils"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestJobPostingRepo(t *testing.T) {

	t.Run("Test Queries", func(t *testing.T) {
		ctx := context.Background()
		forSaveRepo := tinit.InitProviderJobPostingRepo(t)
		jobPostingRepo := tinit.InitRestApiJobPostingRepo(t)

		testSamples := []*TestSample{
			{
				Site:           "jumpit",
				PostingId:      "1",
				RequiredSkills: testutils.RequiredSkills(jobposting.Origin, "java", "python", "go"),
			},
			{
				Site:           "jumpit",
				PostingId:      "2",
				RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java", "python"), testutils.RequiredSkills(jobposting.FromTitle, "go")...),
			},
			{
				Site:           "jumpit",
				PostingId:      "3",
				RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java", "python"), testutils.RequiredSkills(jobposting.FromMainTask, "go")...),
			},
			{
				Site:           "jumpit",
				PostingId:      "4",
				RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java", "python"), testutils.RequiredSkills(jobposting.FromQualifications, "go")...),
			},
			{
				Site:           "jumpit",
				PostingId:      "5",
				RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java", "python"), testutils.RequiredSkills(jobposting.FromPreferred, "go")...),
			},
			{
				Site:           "jumpit",
				PostingId:      "6",
				RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java"), testutils.RequiredSkills(jobposting.FromPreferred, "python", "go")...),
			},
			{
				Site:           "jumpit",
				PostingId:      "7",
				RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "python", "gcp"), testutils.RequiredSkills(jobposting.FromQualifications, "k8s")...),
			},
		}

		for _, sample := range testSamples {
			success, err := forSaveRepo.Save(ctx, testutils.JobPosting(sample.Site, sample.PostingId, sample.RequiredSkills))
			require.NoError(t, err)
			require.True(t, success)
		}

		testCases := []TestCase{ //first in last out, 먼저 저장된 데이터가 나중에 나옴
			{"Exclude FromPreferred", dto.QueryReq{SkillNames: []string{"go"}}, []TestResult{{"jumpit", "4"}, {"jumpit", "3"}, {"jumpit", "2"}, {"jumpit", "1"}}},
			{"Skill has \"AND\" conditions", dto.QueryReq{SkillNames: []string{"python", "go"}}, []TestResult{{"jumpit", "4"}, {"jumpit", "3"}, {"jumpit", "2"}, {"jumpit", "1"}}},
		}

		for _, testCase := range testCases {
			t.Run(testCase.TestName, func(t *testing.T) {
				results, err := jobPostingRepo.GetJobPostings(ctx, 0, 100, testCase.Query)
				require.NoError(t, err)
				require.Len(t, results, len(testCase.ExpectedResults))

				for i, expected := range testCase.ExpectedResults {
					require.Equal(t, expected.Site, results[i].JobPostingId.Site)
					require.Equal(t, expected.PostingId, results[i].JobPostingId.PostingId)
				}
			})

		}
	})
}

type TestSample struct {
	Site           string
	PostingId      string
	RequiredSkills []jobposting.RequiredSkill
}

type TestResult struct {
	Site      string
	PostingId string
}
type TestCase struct {
	TestName        string
	Query           dto.QueryReq
	ExpectedResults []TestResult
}

// func initJobPostings() []*jobposting.JobPostingInfo {

// }
