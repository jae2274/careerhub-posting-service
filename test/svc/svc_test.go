package svc

import (
	"context"
	"testing"
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/gServer"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/processor_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/jae2274/goutils/ptr"
	"github.com/stretchr/testify/require"
)

func TestRegisterCompany(t *testing.T) {
	t.Run("RegisterCompany", func(t *testing.T) {
		companyRepo := tinit.InitCompanyRepo(t)
		grpcServer := gServer.NewDataProcessorServer(tinit.InitJobPostingRepo(t), companyRepo)

		pbCompanies := samplePbCompany()

		for _, pbCompany := range pbCompanies {
			grpcServer.RegisterCompany(context.TODO(), pbCompany)
		}

		savedCompanies, err := companyRepo.FindAll()
		require.NoError(t, err)

		require.Equal(t, 3, len(savedCompanies))
		savedCompaniesMap := make(map[string]*company.Company)
		for _, savedCompany := range savedCompanies {
			savedCompaniesMap[savedCompany.DefaultName] = savedCompany
		}

		findedGogule, ok := savedCompaniesMap["gogule"]
		require.True(t, ok)
		require.Len(t, findedGogule.SiteCompanies, 2)
		require.Equal(t, pbCompanies[0].Description, findedGogule.SiteCompanies[0].Description)
		require.Equal(t, pbCompanies[2].Description, findedGogule.SiteCompanies[1].Description)

		findedApplepie, ok := savedCompaniesMap["applepie"]
		require.True(t, ok)
		require.Len(t, findedApplepie.SiteCompanies, 1)
		require.Equal(t, pbCompanies[1].Description, findedApplepie.SiteCompanies[0].Description)

		findedFaceboot, ok := savedCompaniesMap["faceboot"]
		require.True(t, ok)
		require.Len(t, findedFaceboot.SiteCompanies, 1)
		require.Equal(t, pbCompanies[3].Description, findedFaceboot.SiteCompanies[0].Description)
	})
}

func samplePbCompany() []*processor_grpc.Company {
	return []*processor_grpc.Company{
		{
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
		},
		{
			Site:       "jumpit",
			CompanyId:  "jumpit_company2",
			Name:       "applepie",
			CompanyUrl: ptr.P("https://www.applepie.com"),
			CompanyImages: []string{
				"https://www.applepie.com/images/1.jpg",
				"https://www.applepie.com/images/2.jpg",
			},
			Description: "applepie is a company by jumpit",
			CompanyLogo: "https://www.applepie.com/logo.jpg",
			CreatedAt:   time.Now().Unix(),
		},
		{
			Site:       "wanted",
			CompanyId:  "wanted_company1",
			Name:       "gogule",
			CompanyUrl: ptr.P("https://www.gogule.com"),
			CompanyImages: []string{
				"https://www.gogule.com/images/1.jpg",
				"https://www.gogule.com/images/2.jpg",
			},
			Description: "gogule is a company by wanted",
			CompanyLogo: "https://www.gogule.com/logo.jpg",
			CreatedAt:   time.Now().Unix(),
		},
		{
			Site:       "wanted",
			CompanyId:  "wanted_company2",
			Name:       "faceboot",
			CompanyUrl: ptr.P("https://www.faceboot.com"),
			CompanyImages: []string{
				"https://www.faceboot.com/images/1.jpg",
				"https://www.faceboot.com/images/2.jpg",
			},
			Description: "faceboot is a company by wanted",
			CompanyLogo: "https://www.faceboot.com/logo.jpg",
			CreatedAt:   time.Now().Unix(),
		},
	}
}