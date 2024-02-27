package providergrpc

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/skill"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/gServer"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/provider_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/rpcRepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/rpcService"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, grpcPort int, collections map[string]*mongo.Collection) error {
	llog.Info(ctx, "Starting data processor...")

	jobPostingCollection := collections[(&jobposting.JobPostingInfo{}).Collection()]
	companyCollection := collections[(&company.Company{}).Collection()]
	skillCollection := collections[(&skill.Skill{}).Collection()]
	skillNameCollection := collections[(&skill.SkillName{}).Collection()]

	jobPostingRepo := rpcRepo.NewJobPostingRepo(jobPostingCollection)
	companyRepo := rpcRepo.NewCompanyRepo(companyCollection)
	skillRepo := rpcRepo.NewSkillRepo(skillCollection)
	skillNameRepo := rpcRepo.NewSkillNameRepo(skillNameCollection)

	providerGrpcServer := gServer.NewDataProcessorServer(
		rpcService.NewJobPostingService(jobPostingRepo),
		rpcService.NewCompanyService(companyRepo),
		rpcService.NewSkillService(skillRepo, skillNameRepo),
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Start gRPC server").Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer()
	provider_grpc.RegisterProviderGrpcServer(grpcServer, providerGrpcServer) //client가 사용할 수 있도록 등록

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
