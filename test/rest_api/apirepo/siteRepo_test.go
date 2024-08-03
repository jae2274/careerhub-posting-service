package apirepo

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/site"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestSiteRepo(t *testing.T) {
	t.Run("return isExist false if site not exist", func(t *testing.T) {
		mainCtx := context.Background()
		siteRepo := apirepo.NewSiteRepo(tinit.InitDB(t))

		siteOpt, err := siteRepo.FindBySiteName(mainCtx, "non-exist")
		require.NoError(t, err)
		require.False(t, siteOpt.IsPresent())
	})

	t.Run("return isExist true if site exist", func(t *testing.T) {
		mainCtx := context.Background()
		db := tinit.InitDB(t)
		siteRepo := apirepo.NewSiteRepo(db)

		sampleSite := &site.Site{
			SiteName:         "test",
			PostingUrlFormat: "https://test.com/%s",
		}
		_, err := db.Collection((&site.Site{}).Collection()).InsertOne(mainCtx, sampleSite)
		require.NoError(t, err)

		siteOpt, err := siteRepo.FindBySiteName(mainCtx, sampleSite.SiteName)
		require.NoError(t, err)
		require.True(t, siteOpt.IsPresent())

		siteObj := siteOpt.Get()
		require.Equal(t, sampleSite.SiteName, siteObj.SiteName)
		require.Equal(t, sampleSite.PostingUrlFormat, siteObj.PostingUrlFormat)
	})

}
