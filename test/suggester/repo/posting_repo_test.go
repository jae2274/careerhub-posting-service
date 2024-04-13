package repo

import (
	"context"
	"testing"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/suggester/repo"
	suggesterserver "github.com/jae2274/careerhub-posting-service/careerhub/posting_service/suggester/suggester_server"
	"github.com/jae2274/careerhub-posting-service/test/testutils"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestSuggesterPostingRepo(t *testing.T) {
	//존재하지 않는 경우
	//존재하는 경우

	t.Run("empty jobPostings", func(t *testing.T) {
		ctx := context.Background()
		db := tinit.InitDB(t)
		postingRepo := repo.NewPostingRepo(db)

		now := time.Now()
		postingRepo.GetPostings(ctx, now, now.Add(time.Hour))
	})

	t.Run("test minInsertedAt/maxInsertedAt", func(t *testing.T) {
		for i := 0; i < 10; i++ { //millisecond 단위의 조회로 인해 잘못 조회되는 경우가 있을 수 있으므로 여러번 반복

			ctx := context.Background()
			db := tinit.InitDB(t)
			sugggesterPostingRepo := repo.NewPostingRepo(db)
			providerPostingRepo := rpcRepo.NewJobPostingRepo(db)

			willSavedPostings := []*jobposting.JobPostingInfo{
				testutils.JobPosting("jumpit", "1", []string{"backend"}, ptr.P(3), ptr.P(5), testutils.RequiredSkills(jobposting.Origin, "java", "python", "go")),
				testutils.JobPosting("jumpit", "2", []string{"frontend"}, nil, ptr.P(5), testutils.RequiredSkills(jobposting.Origin, "javascript", "react")),
				testutils.JobPosting("jumpit", "3", []string{"devops"}, ptr.P(3), nil, testutils.RequiredSkills(jobposting.Origin, "aws", "gcp", "azure")),
				testutils.JobPosting("wanted", "4", []string{"devops", "backend"}, nil, nil, testutils.RequiredSkills(jobposting.Origin, "terraform", "k8s")),
				testutils.JobPosting("wanted", "5", []string{"embedded"}, nil, nil, testutils.RequiredSkills(jobposting.Origin, "c", "c++")),
			}

			//db 저장
			allRangeStart := time.Now()
			save(t, ctx, providerPostingRepo, willSavedPostings[0])

			time.Sleep(time.Millisecond * 50) //mongodb에 저장되며 milliseconds의 일부 단위가 소실되므로 이를 고려하여 시간을 늘려줌
			minCreated := time.Now()
			time.Sleep(time.Millisecond * 50)
			for _, jobPosting := range willSavedPostings[1 : len(willSavedPostings)-1] {
				save(t, ctx, providerPostingRepo, jobPosting)
			}
			time.Sleep(time.Millisecond * 50)
			maxCreated := time.Now()
			time.Sleep(time.Millisecond * 50)

			save(t, ctx, providerPostingRepo, willSavedPostings[len(willSavedPostings)-1])
			allRangeEnd := time.Now()

			//일부 jobPosting 조회
			postings, err := sugggesterPostingRepo.GetPostings(ctx, minCreated, maxCreated)
			require.NoError(t, err)

			require.Len(t, postings, len(willSavedPostings)-2, "minCreated, maxCreated: ", minCreated, maxCreated)
			for i, saved := range willSavedPostings[1 : len(willSavedPostings)-1] {
				assertEqualJobPosting(t, saved, postings[i])
			}

			//모든 jobPosting 조회
			postings, err = sugggesterPostingRepo.GetPostings(ctx, allRangeStart, allRangeEnd)
			require.NoError(t, err)

			require.Len(t, postings, len(willSavedPostings))
			for i, saved := range willSavedPostings {
				assertEqualJobPosting(t, saved, postings[i])
			}
		}
	})

	t.Run("Already saved(close->hiring)", func(t *testing.T) { //이미 저장되었다 close->hiring으로 전환된 경우, InsertedAt은 처음 저장된 시간으로 유지되어야 함
		ctx := context.Background()
		db := tinit.InitDB(t)
		sugggesterPostingRepo := repo.NewPostingRepo(db)
		providerPostingRepo := rpcRepo.NewJobPostingRepo(db)

		willSaved := testutils.JobPosting("jumpit", "1", []string{"backend"}, ptr.P(3), ptr.P(5), testutils.RequiredSkills(jobposting.Origin, "java", "python", "go"))
		save(t, ctx, providerPostingRepo, willSaved)
		providerPostingRepo.CloseAll(context.Background(), []*jobposting.JobPostingId{&willSaved.JobPostingId})

		minInsertedAt := time.Now()
		time.Sleep(time.Millisecond * 100)
		save(t, ctx, providerPostingRepo, willSaved)
		maxInsertedAt := time.Now()

		postings, err := sugggesterPostingRepo.GetPostings(ctx, minInsertedAt, maxInsertedAt)
		require.NoError(t, err)
		require.Len(t, postings, 0)
	})
}

func save(t *testing.T, ctx context.Context, postingRepo *rpcRepo.JobPostingRepo, jp *jobposting.JobPostingInfo) {
	_, err := postingRepo.Save(ctx, jp)
	require.NoError(t, err)
}

func assertEqualJobPosting(t *testing.T, expected *jobposting.JobPostingInfo, actual *jobposting.JobPostingInfo) {
	grpcPosting := suggesterserver.ConvertJobPostingToGrpc(actual)

	require.Equal(t, expected.JobPostingId.Site, grpcPosting.Site)
	require.Equal(t, expected.JobPostingId.PostingId, grpcPosting.PostingId)
	require.Equal(t, expected.MainContent.Title, grpcPosting.Title)
	require.Equal(t, expected.CompanyId, grpcPosting.CompanyId)
	require.Equal(t, expected.CompanyName, grpcPosting.CompanyName)

	require.Equal(t, expected.JobCategory, grpcPosting.Info.Categories)
	require.Len(t, grpcPosting.Info.SkillNames, len(expected.RequiredSkill))
	for i, skill := range expected.RequiredSkill {
		require.Equal(t, skill.SkillName, grpcPosting.Info.SkillNames[i])
	}
	require.Equal(t, expected.RequiredCareer.Min, grpcPosting.Info.MinCareer)
	require.Equal(t, expected.RequiredCareer.Max, grpcPosting.Info.MaxCareer)

	require.Equal(t, expected.ImageUrl, grpcPosting.ImageUrl)

}
