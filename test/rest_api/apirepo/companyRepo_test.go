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

	t.Run("return isExisted true if company exists", func(t *testing.T) {
		mainCtx := context.TODO()
		db := tinit.InitDB(t)
		companyRepo := apirepo.NewCompanyRepo(db)

		providerCompanySvc := rpcService.NewCompanyService(rpcRepo.NewCompanyRepo(db))

		_, err := providerCompanySvc.RegisterCompany(mainCtx, sampleCompany)
		require.NoError(t, err)
		companyOpt, err := companyRepo.FindByCompanySiteID(mainCtx, sampleCompany.Site, sampleCompany.CompanyId)

		require.NoError(t, err)
		require.True(t, companyOpt.IsPresent())

		require.Equal(t, sampleCompany.Name, companyOpt.GetPtr().CompanyName)
		require.Equal(t, sampleCompany.CompanyUrl, companyOpt.GetPtr().CompanyUrl)
		require.Equal(t, sampleCompany.CompanyLogo, companyOpt.GetPtr().CompanyLogo)
	})

	t.Run("return empty if nothing registered", func(t *testing.T) {
		mainCtx := context.TODO()
		companyRepo := apirepo.NewCompanyRepo(tinit.InitDB(t))

		companies, err := companyRepo.FindByPrefixCompanyName(mainCtx, "not_exist", 100)

		require.NoError(t, err)
		require.Empty(t, companies)
	})

	sampleCompanies := []*provider_grpc.Company{
		{
			Site:       "jumpit",
			CompanyId:  "jumpit_company1",
			Name:       "helloWorld",
			CompanyUrl: ptr.P("https://www.gogule.com"),
			CompanyImages: []string{
				"https://www.gogule.com/images/1.jpg",
			},
			Description: "gogule is a company by jumpit",
			CompanyLogo: "https://www.gogule.com/logo.jpg",
			CreatedAt:   time.Now().Unix(),
		},
		{
			Site:       "wanted",
			CompanyId:  "wanted_company1",
			Name:       "helloWorld",
			CompanyUrl: ptr.P("https://www.gogule.com"),
			CompanyImages: []string{
				"https://www.gogule.com/images/1.jpg",
			},
			Description: "gogule is a company by jumpit",
			CompanyLogo: "https://www.gogule.com/logo.jpg",
			CreatedAt:   time.Now().Unix(),
		},
		{
			Site:       "wanted",
			CompanyId:  "wanted_company2",
			Name:       "helloCompany",
			CompanyUrl: ptr.P("https://www.gogule2.com"),
			CompanyImages: []string{
				"https://www.gogule2.com/images/1.jpg",
			},
			Description: "gogule2 is a company by jumpit",
			CompanyLogo: "https://www.gogule2.com/logo.jpg",
			CreatedAt:   time.Now().Unix(),
		},
		{
			Site:       "jobkorea",
			CompanyId:  "jobkorea_company3",
			Name:       "hiCompany",
			CompanyUrl: ptr.P("https://www.gogule2.com"),
			CompanyImages: []string{
				"https://www.gogule2.com/images/1.jpg",
			},
			Description: "gogule2 is a company by jumpit",
			CompanyLogo: "https://www.gogule2.com/logo.jpg",
			CreatedAt:   time.Now().Unix(),
		},
	}

	t.Run("return companies if keyword not matched", func(t *testing.T) {
		mainCtx := context.TODO()
		db := tinit.InitDB(t)
		companyRepo := apirepo.NewCompanyRepo(db)

		providerCompanySvc := rpcService.NewCompanyService(rpcRepo.NewCompanyRepo(db))

		for _, company := range sampleCompanies {
			isRegistered, err := providerCompanySvc.RegisterCompany(mainCtx, company)
			require.NoError(t, err)
			require.True(t, isRegistered)
		}

		companies, err := companyRepo.FindByPrefixCompanyName(mainCtx, "not_exist", 100)

		require.NoError(t, err)
		require.Len(t, companies, 0)
	})

	t.Run("return companies if keyword matched", func(t *testing.T) {
		mainCtx := context.TODO()
		db := tinit.InitDB(t)
		companyRepo := apirepo.NewCompanyRepo(db)

		providerCompanySvc := rpcService.NewCompanyService(rpcRepo.NewCompanyRepo(db))

		for _, company := range sampleCompanies {
			isRegistered, err := providerCompanySvc.RegisterCompany(mainCtx, company)
			require.NoError(t, err)
			require.True(t, isRegistered)
		}

		companies, err := companyRepo.FindByPrefixCompanyName(mainCtx, "hello", 100)

		require.NoError(t, err)
		require.Len(t, companies, 2)

		require.Equal(t, "helloWorld", companies[0].DefaultName)
		require.Len(t, companies[0].SiteCompanies, 2)

		require.Equal(t, "jumpit", companies[0].SiteCompanies[0].Site)
		require.Equal(t, "jumpit_company1", companies[0].SiteCompanies[0].CompanyId)
		require.Equal(t, "helloWorld", companies[0].SiteCompanies[0].CompanyName)

		require.Equal(t, "wanted", companies[0].SiteCompanies[1].Site)
		require.Equal(t, "wanted_company1", companies[0].SiteCompanies[1].CompanyId)
		require.Equal(t, "helloWorld", companies[0].SiteCompanies[1].CompanyName)

		require.Equal(t, "helloCompany", companies[1].DefaultName)
		require.Len(t, companies[1].SiteCompanies, 1)
		require.Equal(t, "wanted", companies[1].SiteCompanies[0].Site)
		require.Equal(t, "wanted_company2", companies[1].SiteCompanies[0].CompanyId)
		require.Equal(t, "helloCompany", companies[1].SiteCompanies[0].CompanyName)

		companies, err = companyRepo.FindByPrefixCompanyName(mainCtx, "company", 100)

		require.NoError(t, err)
		require.Len(t, companies, 2)

		require.Equal(t, "helloCompany", companies[0].DefaultName)
		require.Len(t, companies[0].SiteCompanies, 1)
		require.Equal(t, "wanted", companies[0].SiteCompanies[0].Site)
		require.Equal(t, "wanted_company2", companies[0].SiteCompanies[0].CompanyId)
		require.Equal(t, "helloCompany", companies[0].SiteCompanies[0].CompanyName)

		require.Equal(t, "hiCompany", companies[1].DefaultName)
		require.Len(t, companies[1].SiteCompanies, 1)
		require.Equal(t, "jobkorea", companies[1].SiteCompanies[0].Site)
		require.Equal(t, "jobkorea_company3", companies[1].SiteCompanies[0].CompanyId)
		require.Equal(t, "hiCompany", companies[1].SiteCompanies[0].CompanyName)
	})

	t.Run("return companies if keyword matched with limit", func(t *testing.T) {
		mainCtx := context.TODO()
		db := tinit.InitDB(t)
		companyRepo := apirepo.NewCompanyRepo(db)

		providerCompanySvc := rpcService.NewCompanyService(rpcRepo.NewCompanyRepo(db))

		for _, company := range sampleCompanies {
			isRegistered, err := providerCompanySvc.RegisterCompany(mainCtx, company)
			require.NoError(t, err)
			require.True(t, isRegistered)
		}

		companies, err := companyRepo.FindByPrefixCompanyName(mainCtx, "hello", 1)

		require.NoError(t, err)
		require.Len(t, companies, 1)

		require.Equal(t, "helloWorld", companies[0].DefaultName)
		require.Len(t, companies[0].SiteCompanies, 2)
		require.Equal(t, "jumpit", companies[0].SiteCompanies[0].Site)
		require.Equal(t, "jumpit_company1", companies[0].SiteCompanies[0].CompanyId)
		require.Equal(t, "wanted", companies[0].SiteCompanies[1].Site)
		require.Equal(t, "wanted_company1", companies[0].SiteCompanies[1].CompanyId)
	})
}
