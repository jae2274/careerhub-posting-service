package bgrepo

import (
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestJobPostingRepo(t *testing.T) {
	sampleJobPostings := samples()

	t.Run("FindAll from empty DB", func(t *testing.T) {
		jpRepo := tinit.InitBgJobPostingRepo(t)

		jobPostings, err := jpRepo.FindAll()
		if err != nil {
			require.NoError(t, err)
		}

		require.Equal(t, 0, len(jobPostings))
	})

	t.Run("Save and FindAll", func(t *testing.T) {
		jpRepo := tinit.InitBgJobPostingRepo(t)

		_, err := jpRepo.Save(sampleJobPostings[0])
		require.NoError(t, err)
		_, err = jpRepo.Save(sampleJobPostings[1])
		require.NoError(t, err)

		jobPostings, err := jpRepo.FindAll()
		require.NoError(t, err)

		require.Equal(t, 2, len(jobPostings))
		jobPostings[0].ID = primitive.ObjectID{} // ignore ID
		require.Equal(t, *sampleJobPostings[0], *(jobPostings[0]))
		jobPostings[1].ID = primitive.ObjectID{} // ignore ID
		require.Equal(t, *sampleJobPostings[1], *(jobPostings[1]))
	})
}

func samples() []*jobposting.JobPostingInfo {
	sampleJobPosting := &jobposting.JobPostingInfo{
		Site:        "sampleSite",
		PostingId:   "samplePostingId",
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
		RequiredSkill: []string{"sampleRequiredSkill"},
		Tags:          []string{"sampleTags"},
		RequiredCareer: jobposting.Career{
			Min: ptr.P(int32(1)),
			Max: ptr.P(int32(3)),
		},
		PublishedAt: ptr.P(int64(1234567890)),
		ClosedAt:    ptr.P(int64(1234567890)),
		Address:     []string{"sampleAddress"},
	}

	sampleJobPosting2 := &jobposting.JobPostingInfo{
		Site:        "sampleSite2",
		PostingId:   "samplePostingId2",
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
		RequiredSkill: []string{"sampleRequiredSkill2"},
		Tags:          []string{"sampleTags2"},
		RequiredCareer: jobposting.Career{
			Min: ptr.P(int32(2)),
			Max: ptr.P(int32(4)),
		},
		PublishedAt: ptr.P(int64(1234567891)),
		ClosedAt:    ptr.P(int64(1234567891)),
		Address:     []string{"sampleAddress2"},
	}

	return []*jobposting.JobPostingInfo{sampleJobPosting, sampleJobPosting2}
}
