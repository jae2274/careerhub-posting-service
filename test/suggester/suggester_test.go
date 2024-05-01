package suggester

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/suggester/suggester_grpc"
	"github.com/jae2274/careerhub-posting-service/test/testutils"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestSuggesterGrpc(t *testing.T) {
	cancelFunc := tinit.RunTestApp(t)
	defer cancelFunc()

	t.Run("return empty", func(t *testing.T) {
		client := initSuggesterClient(t)

		ctx := context.Background()
		postings, err := client.GetPostings(ctx, &suggester_grpc.GetPostingsRequest{
			MinUnixMilli: 0,
			MaxUnixMilli: time.Now().UnixMilli(),
		})
		require.NoError(t, err)
		require.Len(t, postings.Postings, 0)
	})

	t.Run("return postings", func(t *testing.T) {
		ctx := context.Background()
		db := tinit.InitDB(t)
		postingRepo := rpcRepo.NewJobPostingRepo(db)

		//given
		saveWithDelay := func(jp *jobposting.JobPostingInfo) {
			time.Sleep(100 * time.Millisecond)
			defer time.Sleep(100 * time.Millisecond)
			isSuccess, err := postingRepo.SaveHiring(ctx, jp)
			require.NoError(t, err)
			require.True(t, isSuccess)
		}

		willSavedPostings := []*jobposting.JobPostingInfo{newJobPosting(1), newJobPosting(2), newJobPosting(3), newJobPosting(4), newJobPosting(5)}

		saveWithDelay(willSavedPostings[0])

		minUnixMilli := time.Now().UnixMilli()
		for _, jp := range willSavedPostings[1 : len(willSavedPostings)-1] {
			saveWithDelay(jp)
		}
		maxUnixMilli := time.Now().UnixMilli()

		saveWithDelay(willSavedPostings[len(willSavedPostings)-1])

		//when
		client := initSuggesterClient(t)

		postings, err := client.GetPostings(ctx, &suggester_grpc.GetPostingsRequest{
			MinUnixMilli: minUnixMilli,
			MaxUnixMilli: maxUnixMilli,
		})

		//then
		require.NoError(t, err)
		require.Len(t, postings.Postings, len(willSavedPostings)-2)
		for i, saved := range willSavedPostings[1 : len(willSavedPostings)-1] {
			require.Equal(t, saved.JobPostingId.Site, postings.Postings[i].Site)
			require.Equal(t, saved.JobPostingId.PostingId, postings.Postings[i].PostingId)
			require.Equal(t, saved.MainContent.Title, postings.Postings[i].Title)
			require.Equal(t, saved.CompanyId, postings.Postings[i].CompanyId)
		}
	})
}

func initSuggesterClient(t *testing.T) suggester_grpc.PostingClient {
	envVars := tinit.InitEnvVars(t)
	conn := tinit.InitGrpcClient(t, envVars.SuggesterGrpcPort)

	return suggester_grpc.NewPostingClient(conn)
}

func newJobPosting(number int) *jobposting.JobPostingInfo {
	attachN := func(s string, number int) string {
		return fmt.Sprintf("%s%d", s, number)
	}
	return testutils.JobPosting(
		attachN("site", number),
		attachN("postingId", number),
		[]string{attachN("category", number), attachN("category", number+1)},
		nil,
		nil,
		[]jobposting.RequiredSkill{
			{SkillFrom: jobposting.FromMainTask, SkillName: attachN("skill", number)},
			{SkillFrom: jobposting.FromMainTask, SkillName: attachN("skill", number+1)},
		},
	)
}
