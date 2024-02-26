package rpcService

import (
	"context"
	"testing"
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/provider_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/rpcService"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestRegisterNCloseJobPostings(t *testing.T) {
	t.Run("RegisterNCloseJobPostings", func(t *testing.T) {
		jobPostingRepo := tinit.InitJobPostingRepo(t)
		jobPostingService := rpcService.NewJobPostingService(jobPostingRepo)

		pbJobPostings := samplePbJobPostings()

		for _, pbJobPosting := range pbJobPostings {
			jobPostingService.RegisterJobPostingInfo(context.TODO(), pbJobPosting)
		}

		savedJobPostings, err := jobPostingRepo.FindAll()
		require.NoError(t, err)
		require.Equal(t, 4, len(savedJobPostings))

		for i, savedJobPosting := range savedJobPostings {
			assertJobPostingEqual(t, pbJobPostings[i], savedJobPosting)
			require.Equal(t, jobposting.HIRING, savedJobPosting.Status)
		}

		// CloseJobPostings
		jobPostingService.CloseJobPostings(context.TODO(), &provider_grpc.JobPostings{
			JobPostingIds: []*provider_grpc.JobPostingId{
				pbJobPostings[0].JobPostingId,
				pbJobPostings[2].JobPostingId,
			},
		})

		savedJobPostings, err = jobPostingRepo.FindAll()
		require.NoError(t, err)
		require.Equal(t, 4, len(savedJobPostings))

		require.Equal(t, jobposting.CLOSED, savedJobPostings[0].Status)
		require.Equal(t, jobposting.HIRING, savedJobPostings[1].Status)
		require.Equal(t, jobposting.CLOSED, savedJobPostings[2].Status)
		require.Equal(t, jobposting.HIRING, savedJobPostings[3].Status)
	})
}

func assertJobPostingEqual(t *testing.T, pbJobPosting *provider_grpc.JobPostingInfo, savedJobPosting *jobposting.JobPostingInfo) {
	require.Equal(t, pbJobPosting.JobPostingId.Site, savedJobPosting.JobPostingId.Site)
	require.Equal(t, pbJobPosting.JobPostingId.PostingId, savedJobPosting.JobPostingId.PostingId)
	require.Equal(t, pbJobPosting.CompanyId, savedJobPosting.CompanyId)
	require.Equal(t, pbJobPosting.CompanyName, savedJobPosting.CompanyName)
	require.Equal(t, pbJobPosting.JobCategory, savedJobPosting.JobCategory)
	require.Equal(t, pbJobPosting.MainContent.PostUrl, savedJobPosting.MainContent.PostUrl)
	require.Equal(t, pbJobPosting.MainContent.Title, savedJobPosting.MainContent.Title)
	require.Equal(t, pbJobPosting.MainContent.Intro, savedJobPosting.MainContent.Intro)
	require.Equal(t, pbJobPosting.MainContent.MainTask, savedJobPosting.MainContent.MainTask)
	require.Equal(t, pbJobPosting.MainContent.Qualifications, savedJobPosting.MainContent.Qualifications)
	require.Equal(t, pbJobPosting.MainContent.Preferred, savedJobPosting.MainContent.Preferred)
	require.Equal(t, pbJobPosting.MainContent.Benefits, savedJobPosting.MainContent.Benefits)
	if pbJobPosting.MainContent.RecruitProcess != nil {
		require.Equal(t, *pbJobPosting.MainContent.RecruitProcess, *savedJobPosting.MainContent.RecruitProcess)
	}

	savedJobPostingRequiredSkills := make([]string, len(savedJobPosting.RequiredSkill))
	for i, skill := range savedJobPosting.RequiredSkill {
		savedJobPostingRequiredSkills[i] = skill.SkillName
	}

	require.Equal(t, pbJobPosting.RequiredSkill, savedJobPostingRequiredSkills)
	require.Equal(t, pbJobPosting.Tags, savedJobPosting.Tags)
	if pbJobPosting.RequiredCareer.Min != nil {
		require.NotNil(t, savedJobPosting.RequiredCareer.Min)
		require.Equal(t, *pbJobPosting.RequiredCareer.Min, *savedJobPosting.RequiredCareer.Min)
	} else {
		require.Nil(t, savedJobPosting.RequiredCareer.Min)
	}

	if pbJobPosting.RequiredCareer.Max != nil {
		require.NotNil(t, savedJobPosting.RequiredCareer.Max)
		require.Equal(t, *pbJobPosting.RequiredCareer.Max, *savedJobPosting.RequiredCareer.Max)
	} else {
		require.Nil(t, savedJobPosting.RequiredCareer.Max)
	}

	if pbJobPosting.PublishedAt != nil {
		require.NotNil(t, savedJobPosting.PublishedAt)
		require.Equal(t, *pbJobPosting.PublishedAt, (*savedJobPosting.PublishedAt).UnixMilli())
	} else {
		require.Nil(t, savedJobPosting.PublishedAt)
	}

	if pbJobPosting.ClosedAt != nil {
		require.NotNil(t, savedJobPosting.ClosedAt)
		require.Equal(t, *pbJobPosting.ClosedAt, (*savedJobPosting.ClosedAt).UnixMilli())
	} else {
		require.Nil(t, savedJobPosting.ClosedAt)
	}
	require.Equal(t, pbJobPosting.Address, savedJobPosting.Address)
	require.Equal(t, pbJobPosting.CreatedAt, (savedJobPosting.CreatedAt).UnixMilli())
}

func samplePbJobPostings() []*provider_grpc.JobPostingInfo {
	return []*provider_grpc.JobPostingInfo{
		{
			JobPostingId: &provider_grpc.JobPostingId{
				Site:      "jumpit",
				PostingId: "jumpit_job1",
			},
			CompanyId:   "jumpit_company1",
			CompanyName: "gogule job1",
			JobCategory: []string{"IT", "WEB"},
			MainContent: &provider_grpc.MainContent{
				PostUrl:        "https://www.gogule.com/job1",
				Title:          "gogule job1",
				Intro:          "gogule intro is a job by jumpit",
				MainTask:       "gogule maintask is a job by jumpit",
				Qualifications: "gogule qualifications is a job by jumpit",
				Preferred:      "gogule preferred is a job by jumpit",
				Benefits:       "gogule benefits is a job by jumpit",
				RecruitProcess: ptr.P("gogule recruitprocess is a job by jumpit"),
			},
			RequiredSkill: []string{"golang", "python"},
			Tags:          []string{"간식", "칼퇴"},
			RequiredCareer: &provider_grpc.Career{
				Min: ptr.P(int32(1)),
				Max: ptr.P(int32(3)),
			},
			PublishedAt: ptr.P(time.Now().UnixMilli()),
			ClosedAt:    ptr.P(time.Now().UnixMilli()),
			Address:     []string{"서울시 강남구", "서울시 강북구"},
			CreatedAt:   time.Now().UnixMilli(),
		},
		{
			JobPostingId: &provider_grpc.JobPostingId{
				Site:      "jumpit",
				PostingId: "jumpit_job2",
			},
			CompanyId:   "jumpit_company1",
			CompanyName: "gogule job1",
			JobCategory: []string{"MARKETING", "BUSENESS"},
			MainContent: &provider_grpc.MainContent{
				PostUrl:        "https://www.gogule.com/job2",
				Title:          "gogule job2",
				Intro:          "gogule intro2 is a job by jumpit",
				MainTask:       "gogule maintask2 is a job by jumpit",
				Qualifications: "gogule qualifications2 is a job by jumpit",
				Preferred:      "gogule preferred2 is a job by jumpit",
				Benefits:       "gogule benefits2 is a job by jumpit",
				RecruitProcess: ptr.P("gogule recruitprocess2 is a job by jumpit"),
			},
			RequiredSkill: []string{"excel", "powerpoint"},
			Tags:          []string{"연봉1%", "스톡옵션"},
			RequiredCareer: &provider_grpc.Career{
				Min: ptr.P(int32(5)),
			},
			PublishedAt: ptr.P(time.Now().UnixMilli()),
			ClosedAt:    ptr.P(time.Now().UnixMilli()),
			Address:     []string{"미국", "뉴욕"},
			CreatedAt:   time.Now().UnixMilli(),
		},
		{
			JobPostingId: &provider_grpc.JobPostingId{
				Site:      "wanted",
				PostingId: "wanted_job1",
			},
			CompanyId:   "wanted_company1",
			CompanyName: "fadeout job1",
			JobCategory: []string{"ENGINEER"},
			MainContent: &provider_grpc.MainContent{
				PostUrl:        "https://www.fadeout.com/job1",
				Title:          "fadeout job1",
				Intro:          "fadeout intro is a job by wanted",
				MainTask:       "fadeout maintask is a job by wanted",
				Qualifications: "fadeout qualifications is a job by wanted",
				Preferred:      "fadeout preferred is a job by wanted",
				Benefits:       "fadeout benefits is a job by wanted",
				RecruitProcess: ptr.P("fadeout recruitprocess is a job by wanted"),
			},
			RequiredSkill: []string{"linux", "c++"},
			Tags:          []string{"점심지급"},
			RequiredCareer: &provider_grpc.Career{
				Max: ptr.P(int32(5)),
			},
			PublishedAt: ptr.P(time.Now().UnixMilli()),
			Address:     []string{"일본", "도쿄"},
			CreatedAt:   time.Now().UnixMilli(),
		},
		{
			JobPostingId: &provider_grpc.JobPostingId{
				Site:      "wanted",
				PostingId: "wanted_job2",
			},
			CompanyId:   "wanted_company2",
			CompanyName: "applepie job1",
			JobCategory: []string{"DESIGN"},
			MainContent: &provider_grpc.MainContent{
				PostUrl:        "https://www.applepie.com/job1",
				Title:          "applepie job1",
				Intro:          "applepie intro is a job by wanted",
				MainTask:       "applepie maintask is a job by wanted",
				Qualifications: "applepie qualifications is a job by wanted",
				Preferred:      "applepie preferred is a job by wanted",
				Benefits:       "applepie benefits is a job by wanted",
				RecruitProcess: ptr.P("applepie recruitprocess is a job by wanted"),
			},
			RequiredSkill:  []string{"photoshop", "illustrator"},
			Tags:           []string{"퇴근시간자유"},
			RequiredCareer: &provider_grpc.Career{},
			ClosedAt:       ptr.P(time.Now().UnixMilli()),
			Address:        []string{"중국", "북경"},
			CreatedAt:      time.Now().UnixMilli(),
		},
	}

}
