package rpcRepo

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
// 	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
// 	"github.com/jae2274/goutils/ptr"
// 	"github.com/stretchr/testify/require"
// )

// func TestCompanyRepo(t *testing.T) {
// 	t.Run("FindAll from empty DB", func(t *testing.T) {
// 		companyRepo := tinit.InitCompanyRepo(t)

// 		findedCompany, err := companyRepo.FindByName(context.TODO(), "NonExistentCompany")
// 		require.NoError(t, err)
// 		require.Nil(t, findedCompany)
// 	})

// 	t.Run("InsertCompany", func(t *testing.T) {
// 		companyRepo := tinit.InitCompanyRepo(t)

// 		sampleCompanies := sampleSiteCompanies()
// 		for _, siteCompany := range sampleCompanies {

// 			_, err := companyRepo.InsertCompany(context.Background(), &company.Company{
// 				DefaultName: siteCompany.Name,
// 				SiteCompanies: []*company.SiteCompany{
// 					siteCompany,
// 				},
// 			})
// 			require.NoError(t, err)
// 		}
// 		companies, err := companyRepo.FindAll()
// 		require.NoError(t, err)
// 		require.Equal(t, 2, len(companies))

// 		targetCompanyName := "name1"
// 		findedCompanyID, err := companyRepo.FindByName(context.Background(), targetCompanyName)
// 		require.NoError(t, err)
// 		require.NotNil(t, findedCompanyID)

// 		var isFinded bool = false
// 		for _, company := range companies {
// 			if company.DefaultName == targetCompanyName {
// 				setIgnoreCompanyFields(company)
// 				setIgnoreSiteCompanyFields(sampleCompanies...)
// 				require.Equal(t, *sampleCompanies[0], *company.SiteCompanies[0])
// 				require.Equal(t, company.ID, *findedCompanyID)
// 				isFinded = true
// 			}
// 		}
// 		if !isFinded {
// 			require.Fail(t, fmt.Sprintf("%s not found", targetCompanyName))
// 		}
// 	})
// }

// func sampleSiteCompanies() []*company.SiteCompany {
// 	return []*company.SiteCompany{
// 		{
// 			Site:       "site1",
// 			CompanyId:  "company1",
// 			Name:       "name1",
// 			CompanyUrl: ptr.P("companyUrl1"),
// 			CompanyImages: []string{
// 				"companyImage1",
// 				"companyImage2",
// 			},
// 			Description: "description1",
// 			CompanyLogo: "companyLogo1",
// 			CreatedAt:   time.Now(),
// 			InsertedAt:  time.Now(),
// 			UpdatedAt:   time.Now(),
// 		},
// 		{
// 			Site:      "site2",
// 			CompanyId: "company2",
// 			Name:      "name2",
// 			CompanyImages: []string{
// 				"companyImage3",
// 				"companyImage4",
// 			},
// 			Description: "description2",
// 			CompanyLogo: "companyLogo2",
// 			CreatedAt:   time.Now(),
// 			InsertedAt:  time.Now(),
// 			UpdatedAt:   time.Now(),
// 		},
// 	}
// }

// func setIgnoreCompanyFields(companies ...*company.Company) {
// 	for _, compny := range companies {
// 		compny.InsertedAt = time.Unix(compny.InsertedAt.Unix(), 0)
// 		compny.UpdatedAt = time.Unix(compny.UpdatedAt.Unix(), 0)

// 		for _, siteCompany := range compny.SiteCompanies {
// 			setIgnoreSiteCompanyFields(siteCompany)
// 		}
// 	}
// }

// func setIgnoreSiteCompanyFields(siteCompanies ...*company.SiteCompany) {
// 	for _, siteCompany := range siteCompanies {
// 		siteCompany.InsertedAt = time.Unix(siteCompany.InsertedAt.Unix(), 0)
// 		siteCompany.UpdatedAt = time.Unix(siteCompany.UpdatedAt.Unix(), 0)
// 		siteCompany.CreatedAt = time.Unix(siteCompany.CreatedAt.Unix(), 0)
// 	}
// }
