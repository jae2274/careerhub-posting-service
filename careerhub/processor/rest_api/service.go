package restapi

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/apirepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/dto"
)

type RestApiService interface {
	GetJobPostings(ctx context.Context, page, size int) ([]*dto.JobPostingRes, error)
}

type RestApiServiceImpl struct {
	jobPostingRepo apirepo.JobPostingRepo
}

func NewRestApiService(jobPostingRepo apirepo.JobPostingRepo) RestApiService {
	return &RestApiServiceImpl{
		jobPostingRepo: jobPostingRepo,
	}
}

func (service *RestApiServiceImpl) GetJobPostings(ctx context.Context, page, size int) ([]*dto.JobPostingRes, error) {
	jobPostings, err := service.jobPostingRepo.GetJobPostings(ctx, page, size)
	if err != nil {
		return nil, err
	}

	jobPostingRes := make([]*dto.JobPostingRes, len(jobPostings))
	for i, jobPosting := range jobPostings {
		skills := make([]string, len(jobPosting.RequiredSkill))
		for i, skill := range jobPosting.RequiredSkill {
			skills[i] = skill.SkillName
		}

		jobPostingRes[i] = &dto.JobPostingRes{
			Site:        jobPosting.JobPostingId.Site,
			PostingId:   jobPosting.JobPostingId.PostingId,
			Title:       jobPosting.MainContent.Title,
			CompanyName: jobPosting.CompanyName,
			Skills:      skills,
			Addresses:   jobPosting.Address,
			MinCareer:   jobPosting.RequiredCareer.Min,
			MaxCareer:   jobPosting.RequiredCareer.Max,
		}
	}

	return jobPostingRes, nil
}
