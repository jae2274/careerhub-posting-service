package testutils

import (
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/goutils/ptr"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetIgnoreJobPostingFields(jobPostings []*jobposting.JobPostingInfo) {
	for _, jobPosting := range jobPostings {
		jobPosting.ID = primitive.ObjectID{} // ignore ID
		jobPosting.InsertedAt = time.Unix(jobPosting.InsertedAt.Unix(), 0)
		jobPosting.UpdatedAt = time.Unix(jobPosting.UpdatedAt.Unix(), 0)
		jobPosting.CreatedAt = time.Unix(jobPosting.CreatedAt.Unix(), 0)

		if jobPosting.PublishedAt != nil {
			*jobPosting.PublishedAt = time.Unix(jobPosting.PublishedAt.Unix(), 0)
		}
		if jobPosting.ClosedAt != nil {
			*jobPosting.ClosedAt = time.Unix(jobPosting.ClosedAt.Unix(), 0)
		}
	}
}

func CreateJobPosting(site string, postingId string, requiredSkill []string) *jobposting.JobPostingInfo {
	requiredSkillStruct := make([]jobposting.RequiredSkill, len(requiredSkill))
	for i, skill := range requiredSkill {
		requiredSkillStruct[i] = jobposting.RequiredSkill{
			SkillFrom: jobposting.Origin,
			SkillName: skill,
		}
	}

	time.Sleep(1 * time.Millisecond) //createdAt의 차이를 위해
	return &jobposting.JobPostingInfo{
		JobPostingId: jobposting.JobPostingId{
			Site:      site,
			PostingId: postingId,
		},
		Status:      jobposting.HIRING,
		CompanyId:   "companyId",
		CompanyName: "companyName",
		JobCategory: []string{"jobCategory"},
		MainContent: jobposting.MainContent{
			PostUrl:        "postUrl",
			Title:          "title",
			Intro:          "intro",
			MainTask:       "mainTask",
			Qualifications: "qualifications",
			Preferred:      "preferred",
			Benefits:       "benefits",
			RecruitProcess: nil,
		},
		RequiredSkill: requiredSkillStruct,
		Tags:          []string{"tags"},
		RequiredCareer: jobposting.Career{
			Min: ptr.P(int32(1)),
			Max: ptr.P(int32(3)),
		},
		PublishedAt:    nil,
		ClosedAt:       nil,
		Address:        []string{"address"},
		IsScanComplete: false,
		CreatedAt:      time.Now(),
	}
}