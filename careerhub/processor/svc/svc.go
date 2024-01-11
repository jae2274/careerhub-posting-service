package svc

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	grpc "github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/processor_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// server is used to implement helloworld.GreeterServer.
type DataProcessorServer struct {
	jpRepo *repo.JobPostingRepo
	grpc.UnimplementedDataProcessorServer
}

func NewDataProcessorServer(jpRepo *repo.JobPostingRepo) *DataProcessorServer {
	return &DataProcessorServer{jpRepo: jpRepo}
}

func (sv *DataProcessorServer) CloseJobPostings(context.Context, *grpc.JobPostings) (*grpc.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseJobPostings not implemented")
}
func (sv *DataProcessorServer) RegisterJobPostingInfo(ctx context.Context, msg *grpc.JobPostingInfo) (*grpc.BoolResponse, error) {
	jobPosting := jobposting.JobPostingInfo{
		Site:        msg.Site,
		PostingId:   msg.PostingId,
		CompanyId:   msg.CompanyId,
		CompanyName: msg.CompanyName,
		JobCategory: msg.JobCategory,
		MainContent: jobposting.MainContent{
			PostUrl:        msg.MainContent.PostUrl,
			Title:          msg.MainContent.Title,
			Intro:          msg.MainContent.Intro,
			MainTask:       msg.MainContent.MainTask,
			Qualifications: msg.MainContent.Qualifications,
			Preferred:      msg.MainContent.Preferred,
			Benefits:       msg.MainContent.Benefits,
			RecruitProcess: msg.MainContent.RecruitProcess,
		},
		RequiredSkill: msg.RequiredSkill,
		Tags:          msg.Tags,
		RequiredCareer: jobposting.Career{
			Min: msg.RequiredCareer.Min,
			Max: msg.RequiredCareer.Max,
		},
		PublishedAt: msg.PublishedAt,
		ClosedAt:    msg.ClosedAt,
		Address:     msg.Address,
	}

	result, err := sv.jpRepo.Save(&jobPosting)

	return &grpc.BoolResponse{Success: result}, err
}

func (sv *DataProcessorServer) RegisterCompany(context.Context, *grpc.Company) (*grpc.BoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCompany not implemented")
}
