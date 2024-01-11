package svc

import (
	"context"
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	grpc "github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/processor_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/repo"
)

// server is used to implement helloworld.GreeterServer.
type DataProcessorServer struct {
	jpRepo      *repo.JobPostingRepo
	companyRepo *repo.CompanyRepo
	grpc.UnimplementedDataProcessorServer
}

func NewDataProcessorServer(jpRepo *repo.JobPostingRepo, companyRepo *repo.CompanyRepo) *DataProcessorServer {
	return &DataProcessorServer{jpRepo: jpRepo, companyRepo: companyRepo}
}

func (sv *DataProcessorServer) CloseJobPostings(ctx context.Context, gJpId *grpc.JobPostings) (*grpc.BoolResponse, error) {
	jpIds := make([]*jobposting.JobPostingId, len(gJpId.JobPostingIds))

	for i, gJpId := range gJpId.JobPostingIds {
		jpIds[i] = &jobposting.JobPostingId{
			Site:      gJpId.Site,
			PostingId: gJpId.PostingId,
		}
	}

	err := sv.jpRepo.CloseAll(ctx, jpIds)

	return &grpc.BoolResponse{Success: err == nil}, err
}

func (sv *DataProcessorServer) RegisterJobPostingInfo(ctx context.Context, msg *grpc.JobPostingInfo) (*grpc.BoolResponse, error) {
	var publishedAt *time.Time = nil
	if msg.PublishedAt != nil {
		temp := time.Unix(*msg.PublishedAt, 0)
		publishedAt = &temp
	}

	var closedAt *time.Time = nil
	if msg.ClosedAt != nil {
		temp := time.Unix(*msg.ClosedAt, 0)
		closedAt = &temp
	}

	createdAt := time.Unix(msg.CreatedAt, 0)
	jobPosting := jobposting.JobPostingInfo{
		JobPostingId: jobposting.JobPostingId{
			Site:      msg.Site,
			PostingId: msg.PostingId,
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

	return &grpc.BoolResponse{Success: result}, err
}

func (sv *DataProcessorServer) RegisterCompany(ctx context.Context, gCompany *grpc.Company) (*grpc.BoolResponse, error) {
	siteCompany := &company.SiteCompany{
		Site:          gCompany.Site,
		CompanyId:     gCompany.CompanyId,
		Name:          gCompany.Name,
		CompanyUrl:    gCompany.CompanyUrl,
		CompanyImages: gCompany.CompanyImages,
		Description:   gCompany.Description,
		CompanyLogo:   gCompany.CompanyLogo,
		CreatedAt:     time.Unix(gCompany.CreatedAt, 0),
	}

	existedCompanyId, err := sv.companyRepo.FindIDByName(ctx, gCompany.Name)

	if err != nil {
		return &grpc.BoolResponse{Success: false}, err
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

	return &grpc.BoolResponse{Success: result}, err
}
