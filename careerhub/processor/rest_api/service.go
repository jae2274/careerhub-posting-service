package restapi

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/apirepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/dto"
)

type RestApiService interface {
	GetJobPostings(ctx context.Context, page, size int) ([]*dto.JobPostingRes, error)
	GetJobPostingDetail(ctx context.Context, site, postingId string) (*dto.JobPostingDetailRes, error)
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
			ImageUrl:    jobPosting.ImageUrl,
		}
	}

	return jobPostingRes, nil
}

func (service *RestApiServiceImpl) GetJobPostingDetail(ctx context.Context, site, postingId string) (*dto.JobPostingDetailRes, error) {
	jobPosting, err := service.jobPostingRepo.GetJobPostingDetail(ctx, site, postingId)
	if err != nil {
		return nil, err
	}

	skills := make([]string, len(jobPosting.RequiredSkill))
	for i, skill := range jobPosting.RequiredSkill {
		skills[i] = skill.SkillName
	}

	return &dto.JobPostingDetailRes{
		Site:           jobPosting.JobPostingId.Site,
		PostingId:      jobPosting.JobPostingId.PostingId,
		Title:          jobPosting.MainContent.Title,
		Skills:         skills,
		MainTask:       jobPosting.MainContent.MainTask,
		Qualifications: jobPosting.MainContent.Qualifications,
		Preferred:      jobPosting.MainContent.Preferred,
		Benefits:       jobPosting.MainContent.Benefits,
		RecruitProcess: jobPosting.MainContent.RecruitProcess,
		CareerMin:      jobPosting.RequiredCareer.Min,
		CareerMax:      jobPosting.RequiredCareer.Max,
		Addresses:      jobPosting.Address,
		CompanyId:      jobPosting.CompanyId,
		CompanyName:    jobPosting.CompanyName,
		CompanyImages:  jobPosting.CompanyImages,
		Intro:          jobPosting.MainContent.Intro,
		Tags:           jobPosting.Tags,
	}, nil
}
