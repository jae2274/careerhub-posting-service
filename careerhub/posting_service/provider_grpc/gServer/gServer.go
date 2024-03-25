package gServer

import (
	"context"
	"time"

	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/provider_grpc/provider_grpc"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcService"
)

// server is used to implement helloworld.GreeterServer.
type ProviderGrpcServer struct {
	jobPostingService *rpcService.JobPostingService
	companyService    *rpcService.CompanyService
	skillService      *rpcService.SkillService
	categoryService   *rpcService.CategoryService
	provider_grpc.UnimplementedProviderGrpcServer
}

func NewProviderGrpcServer(jobPostingService *rpcService.JobPostingService, companyService *rpcService.CompanyService, skillService *rpcService.SkillService, categoryService *rpcService.CategoryService) *ProviderGrpcServer {
	return &ProviderGrpcServer{jobPostingService: jobPostingService, companyService: companyService, skillService: skillService, categoryService: categoryService}
}

func (sv *ProviderGrpcServer) IsCompanyRegistered(ctx context.Context, in *provider_grpc.CompanyId) (*provider_grpc.BoolResponse, error) {
	result, err := sv.companyService.IsCompanyRegistered(ctx, in)

	return &provider_grpc.BoolResponse{Success: result}, err
}

func (sv *ProviderGrpcServer) GetAllHiring(ctx context.Context, in *provider_grpc.Site) (*provider_grpc.JobPostings, error) {
	result, err := sv.jobPostingService.GetAllHiring(ctx, in.Site)

	return result, err
}

func (sv *ProviderGrpcServer) CloseJobPostings(ctx context.Context, gJpId *provider_grpc.JobPostings) (*provider_grpc.BoolResponse, error) {
	err := sv.jobPostingService.CloseJobPostings(ctx, gJpId)

	return &provider_grpc.BoolResponse{Success: err == nil}, err
}

func (sv *ProviderGrpcServer) RegisterJobPostingInfo(ctx context.Context, jobPostingInfo *provider_grpc.JobPostingInfo) (*provider_grpc.BoolResponse, error) {
	err := sv.categoryService.RegisterCategories(ctx, jobPostingInfo.JobPostingId.Site, jobPostingInfo.JobCategory)
	if err != nil {
		return &provider_grpc.BoolResponse{Success: false}, err
	}

	preprocessedSkillNames, err := sv.skillService.RegisterSkill(ctx, jobPostingInfo.RequiredSkill)
	if err != nil {
		return &provider_grpc.BoolResponse{Success: false}, err
	}

	jobPostingInfo.RequiredSkill = preprocessedSkillNames
	result, err := sv.jobPostingService.RegisterJobPostingInfo(ctx, jobPostingInfo)

	return &provider_grpc.BoolResponse{Success: result}, err
}

func (sv *ProviderGrpcServer) RegisterCompany(ctx context.Context, gCompany *provider_grpc.Company) (*provider_grpc.BoolResponse, error) {
	result, err := sv.companyService.RegisterCompany(ctx, gCompany)

	return &provider_grpc.BoolResponse{Success: result}, err
}

func UnixMilliToTime(unixMilli int64) time.Time {
	seconds := unixMilli / 1000
	nanoseconds := (unixMilli % 1000) * 1e6
	return time.Unix(seconds, nanoseconds)
}

func UnixMilliToTimePtr(unixMilli *int64) *time.Time {
	if unixMilli == nil {
		return nil
	}
	result := UnixMilliToTime(*unixMilli)
	return &result
}
