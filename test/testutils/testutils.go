package testutils

import (
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
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

func JobPosting(site string, postingId string, categories []string, minCareer *int, maxCareer *int, requiredSkills []jobposting.RequiredSkill) *jobposting.JobPostingInfo {
	var min, max *int32
	if minCareer != nil {
		min = ptr.P(int32(*minCareer))
	}
	if maxCareer != nil {
		max = ptr.P(int32(*maxCareer))
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
		JobCategory: categories,
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
		RequiredSkill: requiredSkills,
		Tags:          []string{"tags"},
		RequiredCareer: jobposting.Career{
			Min: min,
			Max: max,
		},
		PublishedAt:    nil,
		ClosedAt:       nil,
		Address:        []string{"address"},
		IsScanComplete: false,
		CreatedAt:      time.Now(),
	}
}

func RequiredSkills(skillFrom jobposting.SkillFrom, skillNames ...string) []jobposting.RequiredSkill {
	requiredSkills := make([]jobposting.RequiredSkill, len(skillNames))
	for i, name := range skillNames {
		requiredSkills[i] = jobposting.RequiredSkill{
			SkillName: name,
			SkillFrom: skillFrom,
		}
	}
	return requiredSkills
}
