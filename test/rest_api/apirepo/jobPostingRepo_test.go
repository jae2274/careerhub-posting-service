package apirepo_test

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_grpc"
	"github.com/jae2274/careerhub-posting-service/test/testutils"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestJobPostingRepo(t *testing.T) {

	t.Run("Test Queries", func(t *testing.T) {
		ctx := context.Background()
		forSaveRepo := tinit.InitProviderJobPostingRepo(t)
		jobPostingRepo := tinit.InitRestApiJobPostingRepo(t)

		jumpit1 := &TestSample{
			Site:           "jumpit",
			PostingId:      "1",
			Categories:     []string{"backend", "frontend", "devops"},
			MinCareer:      ptr.P(3),
			MaxCareer:      ptr.P(5),
			RequiredSkills: testutils.RequiredSkills(jobposting.Origin, "java", "python", "go"),
		}
		jumpit2 := &TestSample{
			Site:           "jumpit",
			PostingId:      "2",
			Categories:     []string{"backend"},
			MinCareer:      ptr.P(5),
			MaxCareer:      ptr.P(7),
			RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java", "python"), testutils.RequiredSkills(jobposting.FromTitle, "go")...),
		}
		jumpit3 := &TestSample{
			Site:           "jumpit",
			PostingId:      "3",
			Categories:     []string{"frontend"},
			MinCareer:      ptr.P(7),
			MaxCareer:      ptr.P(9),
			RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java", "python"), testutils.RequiredSkills(jobposting.FromMainTask, "go", "c++")...),
		}
		jumpit4 := &TestSample{
			Site:           "jumpit",
			PostingId:      "4",
			Categories:     []string{"devops"},
			MinCareer:      ptr.P(5),
			RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java", "python"), testutils.RequiredSkills(jobposting.FromQualifications, "go", "c++")...),
		}
		jumpit5 := &TestSample{
			Site:           "jumpit",
			PostingId:      "5",
			Categories:     []string{"pm", "cto"},
			MinCareer:      ptr.P(7),
			RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java", "python"), testutils.RequiredSkills(jobposting.FromPreferred, "go")...),
		}
		jumpit6 := &TestSample{
			Site:           "jumpit",
			PostingId:      "6",
			Categories:     []string{"pm"},
			MaxCareer:      ptr.P(3),
			RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "java"), testutils.RequiredSkills(jobposting.FromPreferred, "python", "go")...),
		}
		wanted1 := &TestSample{
			Site:           "wanted",
			PostingId:      "1",
			Categories:     []string{"pm"},
			MaxCareer:      ptr.P(6),
			RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "python", "gcp"), testutils.RequiredSkills(jobposting.FromQualifications, "k8s")...),
		}
		wanted2 := &TestSample{
			Site:      "wanted",
			PostingId: "2",
		}
		willClosed := &TestSample{
			Site:      "wanted",
			PostingId: "3",
		}

		testSamples := []*TestSample{jumpit1, jumpit2, jumpit3, jumpit4, jumpit5, jumpit6, wanted1, wanted2, willClosed}

		for _, sample := range testSamples {
			jp := testutils.JobPosting(sample.Site, sample.PostingId, sample.Categories, sample.MinCareer, sample.MaxCareer, sample.RequiredSkills)
			success, err := forSaveRepo.Save(ctx, jp)
			require.NoError(t, err)
			require.True(t, success)
		}

		require.NoError(t,
			forSaveRepo.CloseAll(ctx, []*jobposting.JobPostingId{{Site: willClosed.Site, PostingId: willClosed.PostingId}}),
		)

		var reversedTestSamples []*TestSample
		for i := len(testSamples) - 1; i >= 0; i-- {
			reversedTestSamples = append(reversedTestSamples, testSamples[i])
		}

		testCases := []TestCase{ //first in last out, 먼저 저장된 데이터가 나중에 나옴
			{"Exclude FromPreferred", &restapi_grpc.QueryReq{SkillNames: []string{"go"}}, []*TestSample{jumpit4, jumpit3, jumpit2, jumpit1}},
			{"Skill has \"AND\" conditions", &restapi_grpc.QueryReq{SkillNames: []string{"c++", "go"}}, []*TestSample{jumpit4, jumpit3}},
			{"Category has \"OR\" conditions", &restapi_grpc.QueryReq{Categories: []*restapi_grpc.CategoryQueryReq{{Site: "jumpit", CategoryName: "backend"}, {Site: "jumpit", CategoryName: "frontend"}, {Site: "jumpit", CategoryName: "devops"}}}, []*TestSample{jumpit4, jumpit3, jumpit2, jumpit1}},
			{"Category: same name, different site", &restapi_grpc.QueryReq{Categories: []*restapi_grpc.CategoryQueryReq{{Site: "jumpit", CategoryName: "pm"}}}, []*TestSample{jumpit6, jumpit5}},
			{"Career range contains posting's career range", &restapi_grpc.QueryReq{MinCareer: ptr.P(int32(4)), MaxCareer: ptr.P(int32(8))}, []*TestSample{jumpit2}},
			{"MinCareer=nil contains posting's maxCareer", &restapi_grpc.QueryReq{MinCareer: nil, MaxCareer: ptr.P(int32(5))}, []*TestSample{jumpit6, jumpit1}},
			{"MaxCareer=nil contains posting's minCareer", &restapi_grpc.QueryReq{MinCareer: ptr.P(int32(6)), MaxCareer: nil}, []*TestSample{jumpit5, jumpit3}},
			{"All jobpostings except closed", &restapi_grpc.QueryReq{}, reversedTestSamples[1:]},
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
	Categories     []string
	MinCareer      *int
	MaxCareer      *int
}

type TestCase struct {
	TestName        string
	Query           *restapi_grpc.QueryReq
	ExpectedResults []*TestSample
}

func ptrInt32(i int) *int32 {
	ptrI32 := int32(i)
	return &ptrI32
}

// func initJobPostings() []*jobposting.JobPostingInfo {

// }
