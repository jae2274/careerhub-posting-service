package providergrpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/provider_grpc"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_grpc"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestRestapiGRPC(t *testing.T) {
	cancelFunc := tinit.RunTestApp(t)
	defer cancelFunc()

	t.Run("return empty with empty request", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitRestapiGrpcClient(t)

		res, err := client.JobPostingsById(ctx, &restapi_grpc.JobPostingsByIdRequest{
			JobPostingIds: []*restapi_grpc.JobPostingIdReq{},
		})
		require.NoError(t, err)

		require.Len(t, res.JobPostings, 0)
	})

	t.Run("return empty with not existed id", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)
		client := tinit.InitRestapiGrpcClient(t)

		res, err := client.JobPostingsById(ctx, &restapi_grpc.JobPostingsByIdRequest{
			JobPostingIds: []*restapi_grpc.JobPostingIdReq{
				{
					Site:      "site",
					PostingId: "not_existed_id",
				},
			},
		})
		require.NoError(t, err)

		require.Len(t, res.JobPostings, 0)
	})

	t.Run("return posting with existed id", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)

		providerClient := tinit.InitProviderGrpcClient(t)
		req := newJobPostingInfo("site", "posting_id", 1)
		isSuccess, err := providerClient.RegisterJobPostingInfo(ctx, req)
		require.NoError(t, err)
		require.True(t, isSuccess.Success)

		client := tinit.InitRestapiGrpcClient(t)
		res, err := client.JobPostingsById(ctx, &restapi_grpc.JobPostingsByIdRequest{
			JobPostingIds: []*restapi_grpc.JobPostingIdReq{
				{Site: req.JobPostingId.Site, PostingId: req.JobPostingId.PostingId},
			},
		})
		require.NoError(t, err)
		require.Len(t, res.JobPostings, 1)
		assertEqual(t, req, res.JobPostings[0], jobposting.HIRING)
	})

	t.Run("return postings with multiple existed ids", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)

		providerClient := tinit.InitProviderGrpcClient(t)
		req1 := newJobPostingInfo("site", "posting_id", 1)
		req2 := newJobPostingInfo("site", "posting_id2", 2)
		req3 := newJobPostingInfo("site", "posting_id3", 3)
		reqs := []*provider_grpc.JobPostingInfo{req1, req2, req3}
		for _, req := range reqs {
			isSuccess, err := providerClient.RegisterJobPostingInfo(ctx, req)
			require.NoError(t, err)
			require.True(t, isSuccess.Success)
		}

		client := tinit.InitRestapiGrpcClient(t)
		res, err := client.JobPostingsById(ctx, &restapi_grpc.JobPostingsByIdRequest{
			JobPostingIds: []*restapi_grpc.JobPostingIdReq{
				{Site: req1.JobPostingId.Site, PostingId: req1.JobPostingId.PostingId},
				{Site: req2.JobPostingId.Site, PostingId: req2.JobPostingId.PostingId},
				{Site: "non_existed_site", PostingId: "non_existed_id"},
			},
		})

		require.NoError(t, err)
		require.Len(t, res.JobPostings, 2)
		assertEqual(t, req1, res.JobPostings[0], jobposting.HIRING)
		assertEqual(t, req2, res.JobPostings[1], jobposting.HIRING)
	})

	t.Run("returm closed posting", func(t *testing.T) {
		ctx := context.Background()
		tinit.InitDB(t)

		providerClient := tinit.InitProviderGrpcClient(t)
		req := newJobPostingInfo("site", "posting_id", 1)
		isSuccess, err := providerClient.RegisterJobPostingInfo(ctx, req)
		require.NoError(t, err)
		require.True(t, isSuccess.Success)

		isSuccess, err = providerClient.CloseJobPostings(ctx, &provider_grpc.JobPostings{
			JobPostingIds: []*provider_grpc.JobPostingId{
				{Site: req.JobPostingId.Site, PostingId: req.JobPostingId.PostingId},
			},
		})
		require.NoError(t, err)
		require.True(t, isSuccess.Success)

		client := tinit.InitRestapiGrpcClient(t)
		res, err := client.JobPostingsById(ctx, &restapi_grpc.JobPostingsByIdRequest{
			JobPostingIds: []*restapi_grpc.JobPostingIdReq{
				{Site: req.JobPostingId.Site, PostingId: req.JobPostingId.PostingId},
			},
		})
		require.NoError(t, err)
		require.Len(t, res.JobPostings, 1)
		assertEqual(t, req, res.JobPostings[0], jobposting.CLOSED)
	})
}

func newJobPostingInfo(site, postingId string, number int) *provider_grpc.JobPostingInfo {
	attachN := func(s string, num int) string {
		return fmt.Sprintf("%s%d", s, num)
	}
	return &provider_grpc.JobPostingInfo{
		JobPostingId: &provider_grpc.JobPostingId{
			Site:      site,
			PostingId: postingId,
		},
		CompanyId:   attachN("company_id", number),
		CompanyName: attachN("company_name", number),
		JobCategory: []string{attachN("category", number)},
		MainContent: &provider_grpc.MainContent{
			PostUrl:        attachN("post_url", number),
			Title:          attachN("title", number),
			Intro:          attachN("intro", number),
			MainTask:       attachN("main_task", number),
			Qualifications: attachN("qualifications", number),
			Preferred:      attachN("preferred", number),
			Benefits:       attachN("benefits", number),
			RecruitProcess: ptr.P(attachN("recruit_process", number)),
		},
		RequiredSkill: []string{attachN("required_skill", number)},
		Tags:          []string{attachN("tag", number)},
		RequiredCareer: &provider_grpc.Career{
			Min: ptr.P(int32(1)),
			Max: ptr.P(int32(2)),
		},
		PublishedAt:   ptr.P(time.Now().UnixMilli()),
		ClosedAt:      ptr.P(time.Now().UnixMilli()),
		Address:       []string{attachN("address", number)},
		CreatedAt:     time.Now().UnixMilli(),
		ImageUrl:      ptr.P(attachN("image_url", number)),
		CompanyImages: []string{attachN("company_image", number)},
	}
}

func assertEqual(t *testing.T, expected *provider_grpc.JobPostingInfo, actual *restapi_grpc.JobPostingRes, expectedStatus jobposting.Status) {
	require.Equal(t, expected.JobPostingId.Site, actual.Site)
	require.Equal(t, expected.JobPostingId.PostingId, actual.PostingId)
	require.Equal(t, expected.MainContent.Title, actual.Title)
	require.Equal(t, expected.CompanyName, actual.CompanyName)
	require.Equal(t, expected.RequiredSkill, actual.Skills)
	require.Equal(t, expected.JobCategory, actual.Categories)
	require.Equal(t, expected.ImageUrl, actual.ImageUrl)
	require.Equal(t, expected.Address, actual.Addresses)
	require.Equal(t, expected.RequiredCareer.Min, actual.MinCareer)
	require.Equal(t, expected.RequiredCareer.Max, actual.MaxCareer)
	require.Equal(t, string(expectedStatus), actual.Status)
}
