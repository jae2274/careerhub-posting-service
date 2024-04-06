package restapi

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_grpc"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_server"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/utils"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, apiGrpcPort int, db *mongo.Database) error {
	jobPostingRepo := apirepo.NewJobPostingRepo(db)
	categoryRepo := apirepo.NewCategoryRepo(db)
	skillRepo := apirepo.NewSkillRepo(db)

	restApiService := restapi_server.NewRestApiService(jobPostingRepo, categoryRepo, skillRepo)

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
