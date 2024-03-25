package restapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/category"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
)

func Run(ctx context.Context, apiPort int, rootPath string, collections map[string]*mongo.Collection) error {
	jobPostingRepo := apirepo.NewJobPostingRepo(collections[(&jobposting.JobPostingInfo{}).Collection()])
	categoryRepo := apirepo.NewCategoryRepo(collections[(&category.Category{}).Collection()])
	skillNameRepo := apirepo.NewSkillNameRepo(collections[(&skill.SkillName{}).Collection()])

	restApiService := NewRestApiService(jobPostingRepo, categoryRepo, skillNameRepo)

	router := mux.NewRouter()
	controller := NewRestApiController(restApiService, router)
	controller.RegisterRoutes(rootPath)

	llog.Msg("Rest API server is running").Level(llog.INFO).Data("apiPort", apiPort).Data("rootPath", rootPath).Log(ctx)
	err := http.ListenAndServe(fmt.Sprintf(":%d", apiPort), router)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
