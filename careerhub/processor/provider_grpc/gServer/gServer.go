package gServer

import (
	"context"
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/processor_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/rpcService"
)

// server is used to implement helloworld.GreeterServer.
type DataProcessorServer struct {
	jobPostingService *rpcService.JobPostingService
	companyService    *rpcService.CompanyService
	skillService      *rpcService.SkillService
	processor_grpc.UnimplementedDataProcessorServer
}

func NewDataProcessorServer(jobPostingService *rpcService.JobPostingService, companyService *rpcService.CompanyService, skillService *rpcService.SkillService) *DataProcessorServer {
	return &DataProcessorServer{jobPostingService: jobPostingService, companyService: companyService, skillService: skillService}
}

func (sv *DataProcessorServer) CloseJobPostings(ctx context.Context, gJpId *processor_grpc.JobPostings) (*processor_grpc.BoolResponse, error) {
	err := sv.jobPostingService.CloseJobPostings(ctx, gJpId)

	return &processor_grpc.BoolResponse{Success: err == nil}, err
}

func (sv *DataProcessorServer) RegisterJobPostingInfo(ctx context.Context, jobPostingInfo *processor_grpc.JobPostingInfo) (*processor_grpc.BoolResponse, error) {
	preprocessedSkillNames, err := sv.skillService.RegisterSkill(ctx, jobPostingInfo.RequiredSkill)
	if err != nil {
		return &processor_grpc.BoolResponse{Success: false}, err
	}

	jobPostingInfo.RequiredSkill = preprocessedSkillNames
	result, err := sv.jobPostingService.RegisterJobPostingInfo(ctx, jobPostingInfo)

	return &processor_grpc.BoolResponse{Success: result}, err
}

func (sv *DataProcessorServer) RegisterCompany(ctx context.Context, gCompany *processor_grpc.Company) (*processor_grpc.BoolResponse, error) {
	result, err := sv.companyService.RegisterCompany(ctx, gCompany)

	return &processor_grpc.BoolResponse{Success: result}, err
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
