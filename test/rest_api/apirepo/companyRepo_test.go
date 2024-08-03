package apirepo

import (
	"context"
	"testing"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/provider_grpc"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcService"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestCompanyRepo(t *testing.T) {
	t.Run("return isExisted false if company not exists", func(t *testing.T) {
		mainCtx := context.TODO()
		companyRepo := apirepo.NewCompanyRepo(tinit.InitDB(t))

		companyOpt, err := companyRepo.FindByCompanySiteID(mainCtx, "jumpit", "1234")

		require.NoError(t, err)
		require.False(t, companyOpt.IsPresent())
	})

	t.Run("return isExisted true if company exists", func(t *testing.T) {
		mainCtx := context.TODO()
		db := tinit.InitDB(t)
		companyRepo := apirepo.NewCompanyRepo(db)

		providerCompanySvc := rpcService.NewCompanyService(rpcRepo.NewCompanyRepo(db))
		sampleCompany := &provider_grpc.Company{
			Site:       "jumpit",
			CompanyId:  "jumpit_company1",
			Name:       "gogule",
			CompanyUrl: ptr.P("https://www.gogule.com"),
			CompanyImages: []string{
				"https://www.gogule.com/images/1.jpg",
				"https://www.gogule.com/images/2.jpg",
			},
			Description: "gogule is a company by jumpit",
			CompanyLogo: "https://www.gogule.com/logo.jpg",
			CreatedAt:   time.Now().Unix(),
		}

		_, err := providerCompanySvc.RegisterCompany(mainCtx, sampleCompany)
		require.NoError(t, err)
		companyOpt, err := companyRepo.FindByCompanySiteID(mainCtx, sampleCompany.Site, sampleCompany.CompanyId)

		require.NoError(t, err)
		require.True(t, companyOpt.IsPresent())

		require.Equal(t, sampleCompany.Name, companyOpt.GetPtr().CompanyName)
		require.Equal(t, sampleCompany.CompanyUrl, companyOpt.GetPtr().CompanyUrl)
		require.Equal(t, sampleCompany.CompanyLogo, companyOpt.GetPtr().CompanyLogo)
	})
}
