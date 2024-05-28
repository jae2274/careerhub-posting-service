package restapi_server

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RestApiService struct {
	jobPostingRepo apirepo.JobPostingRepo
	categoryRepo   apirepo.CategoryRepo
	skillRepo      apirepo.SkillRepo
	restapi_grpc.UnimplementedRestApiGrpcServer
}

func NewRestApiService(jobPostingRepo apirepo.JobPostingRepo, categoryRepo apirepo.CategoryRepo, skillRepo apirepo.SkillRepo) *RestApiService {
	return &RestApiService{
		jobPostingRepo: jobPostingRepo,
		categoryRepo:   categoryRepo,
		skillRepo:      skillRepo,
	}
}

func (service *RestApiService) JobPostings(ctx context.Context, req *restapi_grpc.JobPostingsRequest) (*restapi_grpc.JobPostingsResponse, error) {
	jobPostings, err := service.jobPostingRepo.GetJobPostings(ctx, req.Page, req.Size, req.QueryReq)
	if err != nil {
		return nil, err
	}

	jobPostingRes := make([]*restapi_grpc.JobPostingRes, len(jobPostings))
	for i, jobPosting := range jobPostings {
		jobPostingRes[i] = convertJobPostingInfoToGrpc(&jobPosting)
	}

	return &restapi_grpc.JobPostingsResponse{
		JobPostings: jobPostingRes,
	}, nil
}

func (service *RestApiService) CountJobPostings(ctx context.Context, req *restapi_grpc.JobPostingsRequest) (*restapi_grpc.CountJobPostingsResponse, error) {
	count, err := service.jobPostingRepo.CountJobPostings(ctx, req.QueryReq)
	if err != nil {
		return nil, err
	}

	return &restapi_grpc.CountJobPostingsResponse{
		Count: count,
	}, nil
}

func convertJobPostingInfoToGrpc(jobPosting *jobposting.JobPostingInfo) *restapi_grpc.JobPostingRes {
	skills := make([]string, len(jobPosting.RequiredSkill))
	for i, skill := range jobPosting.RequiredSkill {
		skills[i] = skill.SkillName
	}

	return &restapi_grpc.JobPostingRes{
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
		Status:      string(jobPosting.Status),
	}
}

func (service *RestApiService) JobPostingDetail(ctx context.Context, req *restapi_grpc.JobPostingDetailRequest) (*restapi_grpc.JobPostingDetailResponse, error) {
	jobPosting, err := service.jobPostingRepo.GetJobPostingDetail(ctx, req.Site, req.PostingId)
	if err != nil {
		return nil, err
	}

	skills := make([]string, len(jobPosting.RequiredSkill))
	for i, skill := range jobPosting.RequiredSkill {
		skills[i] = skill.SkillName
	}

	return &restapi_grpc.JobPostingDetailResponse{
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
		Status:         string(jobPosting.Status),
	}, nil
}

func (service *RestApiService) Categories(ctx context.Context, _ *emptypb.Empty) (*restapi_grpc.CategoriesResponse, error) {
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

	categoriesBySite := make([]*restapi_grpc.CategoryRes, len(categoryMap))
	i := 0
	for site, categories := range categoryMap {
		categoriesBySite[i] = &restapi_grpc.CategoryRes{
			Site:       site,
			Categories: categories,
		}
		i++
	}

	return &restapi_grpc.CategoriesResponse{
		CategoriesBySite: categoriesBySite,
	}, nil
}

func (service *RestApiService) Skills(ctx context.Context, _ *emptypb.Empty) (*restapi_grpc.SkillsResponse, error) {
	skills, err := service.skillRepo.GetAllSkills(ctx)
	if err != nil {
		return nil, err
	}

	var skillResList []*restapi_grpc.SkillRes
	for _, skill := range skills {
		skillResList = append(skillResList, &restapi_grpc.SkillRes{
			DefaultName: skill.DefaultName,
			SkillNames:  skill.SkillNames,
		})
	}

	return &restapi_grpc.SkillsResponse{
		Skills: skillResList,
	}, nil
}

func (service *RestApiService) JobPostingsById(ctx context.Context, in *restapi_grpc.JobPostingsByIdRequest) (*restapi_grpc.JobPostingsResponse, error) {
	if len(in.JobPostingIds) == 0 {
		return &restapi_grpc.JobPostingsResponse{JobPostings: []*restapi_grpc.JobPostingRes{}}, nil
	}
	postings, err := service.jobPostingRepo.GetJobPostingsById(ctx, in.JobPostingIds)
	if err != nil {
		return nil, err
	}

	jobPostingRes := make([]*restapi_grpc.JobPostingRes, len(postings))
	for i, jobPosting := range postings {
		jobPostingRes[i] = convertJobPostingInfoToGrpc(&jobPosting)
	}

	return &restapi_grpc.JobPostingsResponse{JobPostings: jobPostingRes}, nil
}
