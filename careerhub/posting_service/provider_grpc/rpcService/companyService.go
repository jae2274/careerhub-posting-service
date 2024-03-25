package rpcService

import (
	"context"

	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/company"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/utils"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/provider_grpc/provider_grpc"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
)

type CompanyService struct {
	companyRepo *rpcRepo.CompanyRepo
}

func NewCompanyService(companyRepo *rpcRepo.CompanyRepo) *CompanyService {
	return &CompanyService{companyRepo: companyRepo}
}

func (sv *CompanyService) IsCompanyRegistered(ctx context.Context, gCompanyId *provider_grpc.CompanyId) (bool, error) {
	return sv.companyRepo.IsExisted(ctx, gCompanyId.Site, gCompanyId.CompanyId)
}

func (sv *CompanyService) RegisterCompany(ctx context.Context, gCompany *provider_grpc.Company) (bool, error) {
	siteCompany := &company.SiteCompany{
		Site:          gCompany.Site,
		CompanyId:     gCompany.CompanyId,
		Name:          gCompany.Name,
		CompanyUrl:    gCompany.CompanyUrl,
		CompanyImages: gCompany.CompanyImages,
		Description:   gCompany.Description,
		CompanyLogo:   gCompany.CompanyLogo,
		CreatedAt:     utils.UnixMilliToTime(gCompany.CreatedAt),
	}

	existedCompanyId, err := sv.companyRepo.FindIDByName(ctx, gCompany.Name)

	if err != nil {
		return false, err
	}

	if existedCompanyId != nil {
		return sv.companyRepo.AppendSiteCompany(ctx, *existedCompanyId, siteCompany)
	} else {
		company := &company.Company{
			DefaultName:   gCompany.Name,
			SiteCompanies: []*company.SiteCompany{siteCompany},
		}

		return sv.companyRepo.InsertCompany(ctx, company)
	}
}
