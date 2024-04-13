package apirepo_test

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_grpc"
	"github.com/jae2274/careerhub-posting-service/test/testutils"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestJobPostingRepo(t *testing.T) {

	t.Run("Test Queries", func(t *testing.T) {
		ctx := context.Background()
		db := tinit.InitDB(t)
		forSaveRepo := rpcRepo.NewJobPostingRepo(db)
		jobPostingRepo := apirepo.NewJobPostingRepo(db)

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
			RequiredSkills: append(testutils.RequiredSkills(jobposting.Origin, "python", "gcp", "golang"), testutils.RequiredSkills(jobposting.FromQualifications, "k8s")...),
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
			{"Exclude FromPreferred", NewQueryReqBuilder().AddSkillNames("go").Build(), []*TestSample{jumpit4, jumpit3, jumpit2, jumpit1}},
			{"Skill has \"OR\" conditions(For example: Different name, same skill)", NewQueryReqBuilder().AddSkillNames("go", "golang").Build(), []*TestSample{wanted1, jumpit4, jumpit3, jumpit2, jumpit1}},
			{"Skill has \"AND\" conditions", NewQueryReqBuilder().AddSkillNames("c++").AddSkillNames("go", "golang").Build(), []*TestSample{jumpit4, jumpit3}},
			{"Category has \"OR\" conditions", NewQueryReqBuilder().AddCategory("jumpit", "backend").AddCategory("jumpit", "frontend").AddCategory("jumpit", "devops").Build(), []*TestSample{jumpit4, jumpit3, jumpit2, jumpit1}},
			{"Category: same name, different site", NewQueryReqBuilder().AddCategory("jumpit", "pm").Build(), []*TestSample{jumpit6, jumpit5}},
			{"Career range contains posting's career range", NewQueryReqBuilder().SetMinCareer(4).SetMaxCareer(8).Build(), []*TestSample{jumpit2}},
			{"MinCareer=nil contains posting's maxCareer", NewQueryReqBuilder().SetMaxCareer(5).Build(), []*TestSample{jumpit6, jumpit1}},
			{"MaxCareer=nil contains posting's minCareer", NewQueryReqBuilder().SetMinCareer(6).Build(), []*TestSample{jumpit5, jumpit3}},
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

type QueryReqBuilder struct {
	categories []*restapi_grpc.CategoryQueryReq
	skillNames []*restapi_grpc.SkillQueryReq
	minCareer  *int32
	maxCareer  *int32
}

func NewQueryReqBuilder() *QueryReqBuilder {
	return &QueryReqBuilder{}
}

func (qb *QueryReqBuilder) AddCategory(site, categoryName string) *QueryReqBuilder {
	qb.categories = append(qb.categories, &restapi_grpc.CategoryQueryReq{Site: site, CategoryName: categoryName})
	return qb
}

func (qb *QueryReqBuilder) AddSkillNames(skillNames ...string) *QueryReqBuilder {
	qb.skillNames = append(qb.skillNames, &restapi_grpc.SkillQueryReq{Or: skillNames})
	return qb
}

func (qb *QueryReqBuilder) SetMinCareer(minCareer int) *QueryReqBuilder {
	qb.minCareer = ptrInt32(minCareer)
	return qb
}

func (qb *QueryReqBuilder) SetMaxCareer(maxCareer int) *QueryReqBuilder {
	qb.maxCareer = ptrInt32(maxCareer)
	return qb
}

func (qb *QueryReqBuilder) Build() *restapi_grpc.QueryReq {
	return &restapi_grpc.QueryReq{
		Categories: qb.categories,
		SkillNames: qb.skillNames,
		MinCareer:  qb.minCareer,
		MaxCareer:  qb.maxCareer,
	}
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
