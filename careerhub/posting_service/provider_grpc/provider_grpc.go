package providergrpc

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/gServer"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/provider_grpc"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcService"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/utils"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, grpcPort int, db *mongo.Database) error {

	jobPostingRepo := rpcRepo.NewJobPostingRepo(db)
	companyRepo := rpcRepo.NewCompanyRepo(db)
	skillRepo := rpcRepo.NewSkillRepo(db)
	skillNameRepo := rpcRepo.NewSkillNameRepo(db)
	categoryRepo := rpcRepo.NewCategoryRepo(db)

	providerGrpcServer := gServer.NewProviderGrpcServer(
		rpcService.NewJobPostingService(jobPostingRepo),
		rpcService.NewCompanyService(companyRepo),
		rpcService.NewSkillService(skillRepo, skillNameRepo),
		rpcService.NewCategoryService(categoryRepo),
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Starting provider grpc server...").Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer(utils.Middlewares()...)

	provider_grpc.RegisterProviderGrpcServer(grpcServer, providerGrpcServer) //client가 사용할 수 있도록 등록

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
