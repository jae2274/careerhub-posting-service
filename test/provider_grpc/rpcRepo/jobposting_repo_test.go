package rpcRepo

import (
	"context"
	"testing"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
	"github.com/jae2274/careerhub-posting-service/test/testutils"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestJobPostingRepo(t *testing.T) {
	sampleJobPostings := samples()

	t.Run("FindAll from empty DB", func(t *testing.T) {
		jpRepo := rpcRepo.NewJobPostingRepo(tinit.InitDB(t))

		jobPostings, err := jpRepo.FindAll()
		if err != nil {
			require.NoError(t, err)
		}

		require.Equal(t, 0, len(jobPostings))
	})

	t.Run("Save and FindAll", func(t *testing.T) {
		jpRepo := rpcRepo.NewJobPostingRepo(tinit.InitDB(t))

		_, err := jpRepo.SaveHiring(context.TODO(), sampleJobPostings[0])
		require.NoError(t, err)
		_, err = jpRepo.SaveHiring(context.TODO(), sampleJobPostings[1])
		require.NoError(t, err)

		jobPostings, err := jpRepo.FindAll()
		require.NoError(t, err)

		require.Equal(t, 2, len(jobPostings))
		testutils.SetIgnoreJobPostingFields(sampleJobPostings)
		testutils.SetIgnoreJobPostingFields(jobPostings)
		require.Equal(t, *sampleJobPostings[0], *(jobPostings[0]))
		require.Equal(t, *sampleJobPostings[1], *(jobPostings[1]))
	})

	t.Run("CloseAll", func(t *testing.T) {
		jpRepo := rpcRepo.NewJobPostingRepo(tinit.InitDB(t))

		_, err := jpRepo.SaveHiring(context.TODO(), sampleJobPostings[0])
		require.NoError(t, err)
		_, err = jpRepo.SaveHiring(context.TODO(), sampleJobPostings[1])
		require.NoError(t, err)

		err = jpRepo.CloseAll(context.TODO(), []*jobposting.JobPostingId{&sampleJobPostings[1].JobPostingId})
		require.NoError(t, err)

		jobPostings, err := jpRepo.FindAll()
		require.NoError(t, err)

		require.Equal(t, 2, len(jobPostings))
		require.Equal(t, jobposting.HIRING, jobPostings[0].Status)
		require.Equal(t, jobposting.CLOSED, jobPostings[1].Status)
	})
}

func samples() []*jobposting.JobPostingInfo {
	sampleJobPosting := &jobposting.JobPostingInfo{
		JobPostingId: jobposting.JobPostingId{
			Site:      "sampleSite",
			PostingId: "samplePostingId",
		},
		Status:      jobposting.HIRING,
		CompanyId:   "sampleCompanyId",
		CompanyName: "sampleCompanyName",
		JobCategory: []string{"sampleJobCategory"},
		MainContent: jobposting.MainContent{
			PostUrl:        "samplePostUrl",
			Title:          "sampleTitle",
			Intro:          "sampleIntro",
			MainTask:       "sampleMainTask",
			Qualifications: "sampleQualifications",
			Preferred:      "samplePreferred",
			Benefits:       "sampleBenefits",
		},
		RequiredSkill: []jobposting.RequiredSkill{
			{
				SkillFrom: jobposting.Origin,
				SkillName: "sampleSkillName",
			},
		},
		Tags: []string{"sampleTags"},
		RequiredCareer: jobposting.Career{
			Min: ptr.P(int32(1)),
			Max: ptr.P(int32(3)),
		},
		PublishedAt: ptr.P(time.Now()),
		ClosedAt:    ptr.P(time.Now()),
		CreatedAt:   time.Now(),
		Address:     []string{"sampleAddress"},
	}

	sampleJobPosting2 := &jobposting.JobPostingInfo{
		JobPostingId: jobposting.JobPostingId{
			Site:      "sampleSite2",
			PostingId: "samplePostingId2",
		},
		Status:      jobposting.HIRING,
		CompanyId:   "sampleCompanyId2",
		CompanyName: "sampleCompanyName2",
		JobCategory: []string{"sampleJobCategory2"},
		MainContent: jobposting.MainContent{
			PostUrl: "samplePostUrl2",

			Title:          "sampleTitle2",
			Intro:          "sampleIntro2",
			MainTask:       "sampleMainTask2",
			Qualifications: "sampleQualifications2",
			Preferred:      "samplePreferred2",
			Benefits:       "sampleBenefits2",
		},
		RequiredSkill: []jobposting.RequiredSkill{
			{
				SkillFrom: jobposting.Origin,
				SkillName: "sampleSkillName2",
			},
		},
		Tags: []string{"sampleTags2"},
		RequiredCareer: jobposting.Career{
			Min: ptr.P(int32(2)),
			Max: ptr.P(int32(4)),
		},
		PublishedAt: ptr.P(time.Now()),
		ClosedAt:    ptr.P(time.Now()),
		CreatedAt:   time.Now(),
		Address:     []string{"sampleAddress2"},
	}

	return []*jobposting.JobPostingInfo{sampleJobPosting, sampleJobPosting2}
}
