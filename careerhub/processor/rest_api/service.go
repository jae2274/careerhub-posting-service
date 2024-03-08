package restapi

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/apirepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/dto"
)

type RestApiService interface {
	GetJobPostings(ctx context.Context, req *dto.GetJobPostingsRequest) ([]*dto.JobPostingRes, error)
	GetJobPostingDetail(ctx context.Context, site, postingId string) (*dto.JobPostingDetailRes, error)
	GetAllCategories(ctx context.Context) (*dto.CategoriesRes, error)
	GetAllSkills(ctx context.Context) ([]string, error)
}

type RestApiServiceImpl struct {
	jobPostingRepo apirepo.JobPostingRepo
	categoryRepo   apirepo.CategoryRepo
	skillRepo      apirepo.SkillNameRepo
}

func NewRestApiService(jobPostingRepo apirepo.JobPostingRepo, categoryRepo apirepo.CategoryRepo, skillRepo apirepo.SkillNameRepo) RestApiService {
	return &RestApiServiceImpl{
		jobPostingRepo: jobPostingRepo,
		categoryRepo:   categoryRepo,
		skillRepo:      skillRepo,
	}
}

func (service *RestApiServiceImpl) GetJobPostings(ctx context.Context, req *dto.GetJobPostingsRequest) ([]*dto.JobPostingRes, error) {
	jobPostings, err := service.jobPostingRepo.GetJobPostings(ctx, req.Page, req.Size, req.QueryReq)
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
			Categories:  jobPosting.JobCategory,
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

func (service *RestApiServiceImpl) GetAllCategories(ctx context.Context) (*dto.CategoriesRes, error) {
	categories, err := service.categoryRepo.GetAllCategories(ctx)
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[string][]string)
	for _, category := range categories {
		if _, ok := categoryMap[category.Site]; !ok {
			categoryMap[category.Site] = make([]string, 0)
		}

		categoryMap[category.Site] = append(categoryMap[category.Site], category.Name)
	}

	categoriesBySite := make([]dto.CategoryRes, len(categoryMap))
	i := 0
	for site, categories := range categoryMap {
		categoriesBySite[i] = dto.CategoryRes{
			Site:       site,
			Categories: categories,
		}
		i++
	}

	return &dto.CategoriesRes{
		CategoriesBySite: categoriesBySite,
	}, nil
}

func (service *RestApiServiceImpl) GetAllSkills(ctx context.Context) ([]string, error) {
	return service.skillRepo.GetAllSkills(ctx)
}
