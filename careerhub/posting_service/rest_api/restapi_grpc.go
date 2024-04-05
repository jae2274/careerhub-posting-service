package restapi

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/category"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_grpc"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_server"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/utils"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, apiGrpcPort int, collections map[string]*mongo.Collection) error {
	jobPostingRepo := apirepo.NewJobPostingRepo(collections[(&jobposting.JobPostingInfo{}).Collection()])
	categoryRepo := apirepo.NewCategoryRepo(collections[(&category.Category{}).Collection()])
	skillNameRepo := apirepo.NewSkillNameRepo(collections[(&skill.SkillName{}).Collection()])

	restApiService := restapi_server.NewRestApiService(jobPostingRepo, categoryRepo, skillNameRepo)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", apiGrpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Starting restAPI grpc server").Level(llog.INFO).Data("port", apiGrpcPort).Log(ctx)

	grpcServer := grpc.NewServer(utils.Middlewares()...)
	restapi_grpc.RegisterRestApiGrpcServer(grpcServer, restApiService)

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
