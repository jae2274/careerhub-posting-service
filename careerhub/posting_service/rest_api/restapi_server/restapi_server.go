package restapi_server

import (
	"context"
	"fmt"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/site"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_grpc"
	"github.com/jae2274/goutils/cchan/async"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/optional"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RestApiService struct {
	jobPostingRepo apirepo.JobPostingRepo
	categoryRepo   apirepo.CategoryRepo
	skillRepo      apirepo.SkillRepo
	companyRepo    apirepo.CompanyRepo
	siteRepo       apirepo.SiteRepo
	restapi_grpc.UnimplementedRestApiGrpcServer
}

func NewRestApiService(jobPostingRepo apirepo.JobPostingRepo, categoryRepo apirepo.CategoryRepo, skillRepo apirepo.SkillRepo, companyRepo apirepo.CompanyRepo, siteRepo apirepo.SiteRepo) *RestApiService {
	return &RestApiService{
		jobPostingRepo: jobPostingRepo,
		categoryRepo:   categoryRepo,
		skillRepo:      skillRepo,
		companyRepo:    companyRepo,
		siteRepo:       siteRepo,
	}
}

func (service *RestApiService) JobPostings(ctx context.Context, req *restapi_grpc.JobPostingsRequest) (*restapi_grpc.JobPostingsResponse, error) {
	jobPostingRes, err := service.jobPostingRepo.GetJobPostings(ctx, req.Page, req.Size, req.QueryReq)
	if err != nil {
		return nil, err
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

func (service *RestApiService) JobPostingDetail(ctx context.Context, req *restapi_grpc.JobPostingDetailRequest) (*restapi_grpc.JobPostingDetailResponse, error) {
	jobPosting, err := service.jobPostingRepo.GetJobPostingDetail(ctx, req.Site, req.PostingId)
	if err != nil {
		return nil, err
	}

	skills := make([]string, len(jobPosting.RequiredSkill))
	for i, skill := range jobPosting.RequiredSkill {
		skills[i] = skill.SkillName
	}

	response := &restapi_grpc.JobPostingDetailResponse{
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
	}

	companySummaryChan := async.ExecAsync(func() (optional.Optional[apirepo.SiteCompanySummary], error) {
		return service.companyRepo.FindByCompanySiteID(ctx, jobPosting.JobPostingId.Site, jobPosting.CompanyId)
	})

	siteChan := async.ExecAsync(func() (optional.Optional[site.Site], error) {
		return service.siteRepo.FindBySiteName(ctx, jobPosting.JobPostingId.Site)
	})

	attachCompanyInfo(ctx, response, <-companySummaryChan)
	attachSiteInfo(ctx, response, <-siteChan)

	return response, nil
}

func attachCompanyInfo(ctx context.Context, response *restapi_grpc.JobPostingDetailResponse, companyOptResult async.Result[optional.Optional[apirepo.SiteCompanySummary]]) {
	if companyOptResult.Err != nil {
		llog.LogErr(ctx, companyOptResult.Err)
		return //에러 발생 시 회사 정보는 무시하고 진행
	}

	if companyOptResult.Value.IsPresent() {
		response.CompanyUrl = companyOptResult.Value.GetPtr().CompanyUrl
		response.CompanyLogo = companyOptResult.Value.GetPtr().CompanyLogo
	} else {
		llog.LogErr(ctx, fmt.Errorf("company not found. site: %s, company_id: %s", response.Site, response.CompanyId)) //회사 정보가 없을 경우 로그 남김
	}
}

func attachSiteInfo(ctx context.Context, response *restapi_grpc.JobPostingDetailResponse, siteOptResult async.Result[optional.Optional[site.Site]]) {
	if siteOptResult.Err != nil {
		llog.LogErr(ctx, siteOptResult.Err)
		return //에러 발생 시 사이트 정보는 무시하고 진행
	}

	if siteOptResult.Value.IsPresent() {
		response.PostUrl = fmt.Sprintf(siteOptResult.Value.GetPtr().PostingUrlFormat, response.PostingId)
	} else {
		llog.LogErr(ctx, fmt.Errorf("site not found. site: %s", response.Site)) //사이트 정보가 없을 경우 로그 남김
	}
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

	jobPostingRes, err := service.jobPostingRepo.GetJobPostingsById(ctx, in.JobPostingIds)
	if err != nil {
		return nil, err
	}

	return &restapi_grpc.JobPostingsResponse{JobPostings: jobPostingRes}, nil
}

func (service *RestApiService) Companies(ctx context.Context, req *restapi_grpc.CompaniesRequest) (*restapi_grpc.CompaniesResponse, error) {
	companies, err := service.companyRepo.FindByPrefixCompanyName(ctx, req.PrefixKeyword)
	if err != nil {
		return nil, err
	}

	companyResList := make([]*restapi_grpc.CompanyRes, 0, len(companies))
	for _, company := range companies {
		siteCompanies := make([]*restapi_grpc.SiteCompanyRes, 0, len(company.SiteCompanies))
		for _, siteCompany := range company.SiteCompanies {
			siteCompanies = append(siteCompanies, &restapi_grpc.SiteCompanyRes{
				Site:      siteCompany.Site,
				CompanyId: siteCompany.CompanyId,
			})
		}

		companyResList = append(companyResList, &restapi_grpc.CompanyRes{
			DefaultName:   company.DefaultName,
			SiteCompanies: siteCompanies,
		})
	}

	return &restapi_grpc.CompaniesResponse{
		Companies: companyResList,
	}, nil
}
