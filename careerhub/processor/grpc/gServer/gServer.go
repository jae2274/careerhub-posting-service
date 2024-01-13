package gServer

import (
	"context"
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/processor_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/rpcRepo"
)

// server is used to implement helloworld.GreeterServer.
type DataProcessorServer struct {
	jpRepo      *rpcRepo.JobPostingRepo
	companyRepo *rpcRepo.CompanyRepo
	processor_grpc.UnimplementedDataProcessorServer
}

func NewDataProcessorServer(jpRepo *rpcRepo.JobPostingRepo, companyRepo *rpcRepo.CompanyRepo) *DataProcessorServer {
	return &DataProcessorServer{jpRepo: jpRepo, companyRepo: companyRepo}
}

func (sv *DataProcessorServer) CloseJobPostings(ctx context.Context, gJpId *processor_grpc.JobPostings) (*processor_grpc.BoolResponse, error) {
	jpIds := make([]*jobposting.JobPostingId, len(gJpId.JobPostingIds))

	for i, gJpId := range gJpId.JobPostingIds {
		jpIds[i] = &jobposting.JobPostingId{
			Site:      gJpId.Site,
			PostingId: gJpId.PostingId,
		}
	}

	err := sv.jpRepo.CloseAll(ctx, jpIds)

	return &processor_grpc.BoolResponse{Success: err == nil}, err
}

func (sv *DataProcessorServer) RegisterJobPostingInfo(ctx context.Context, msg *processor_grpc.JobPostingInfo) (*processor_grpc.BoolResponse, error) {
	publishedAt := UnixMilliToTimePtr(msg.PublishedAt)
	closedAt := UnixMilliToTimePtr(msg.ClosedAt)
	createdAt := UnixMilliToTime(msg.CreatedAt)

	jobPosting := jobposting.JobPostingInfo{
		JobPostingId: jobposting.JobPostingId{
			Site:      msg.JobPostingId.Site,
			PostingId: msg.JobPostingId.PostingId,
		},
		Status:      jobposting.HIRING,
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
		PublishedAt: publishedAt,
		ClosedAt:    closedAt,
		Address:     msg.Address,
		CreatedAt:   createdAt,
	}

	result, err := sv.jpRepo.Save(ctx, &jobPosting)

	return &processor_grpc.BoolResponse{Success: result}, err
}

func (sv *DataProcessorServer) RegisterCompany(ctx context.Context, gCompany *processor_grpc.Company) (*processor_grpc.BoolResponse, error) {
	siteCompany := &company.SiteCompany{
		Site:          gCompany.Site,
		CompanyId:     gCompany.CompanyId,
		Name:          gCompany.Name,
		CompanyUrl:    gCompany.CompanyUrl,
		CompanyImages: gCompany.CompanyImages,
		Description:   gCompany.Description,
		CompanyLogo:   gCompany.CompanyLogo,
		CreatedAt:     UnixMilliToTime(gCompany.CreatedAt),
	}

	existedCompanyId, err := sv.companyRepo.FindIDByName(ctx, gCompany.Name)

	if err != nil {
		return &processor_grpc.BoolResponse{Success: false}, err
	}

	var result bool
	if existedCompanyId != nil {
		result, err = sv.companyRepo.AppendSiteCompany(ctx, *existedCompanyId, siteCompany)
	} else {
		company := &company.Company{
			DefaultName:   gCompany.Name,
			SiteCompanies: []*company.SiteCompany{siteCompany},
		}

		result, err = sv.companyRepo.InsertCompany(ctx, company)
	}

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
