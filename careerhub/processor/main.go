package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/skill"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/mongocfg"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/vars"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/gServer"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/processor_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/rpcRepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/rpcService"
	"github.com/jae2274/goutils/llog"
	"google.golang.org/grpc"
)

const (
	app     = "data-processor"
	service = "careerhub"

	ctxKeyTraceID = "trace_id"
)

func main() {
	ctx := context.Background()

	initLogger(ctx)
	listener, jobPostingRepo, companyRepo, skillRepo := initApp(ctx)

	dataProcessorServer := gServer.NewDataProcessorServer(
		rpcService.NewJobPostingService(jobPostingRepo),
		rpcService.NewCompanyService(companyRepo),
		rpcService.NewSkillService(skillRepo),
	)

	grpcServer := grpc.NewServer()
	processor_grpc.RegisterDataProcessorServer(grpcServer, dataProcessorServer) //client가 사용할 수 있도록 등록

	err := grpcServer.Serve(listener)
	checkErr(ctx, err)
}

func initLogger(ctx context.Context) {
	llog.SetMetadata("service", service)
	llog.SetMetadata("app", app)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	checkErr(ctx, err)

	llog.SetMetadata("hostname", hostname)
}

func initApp(ctx context.Context) (net.Listener, *rpcRepo.JobPostingRepo, *rpcRepo.CompanyRepo, *rpcRepo.SkillRepo) {
	llog.Info(ctx, "Starting data processor...")

	vars, err := vars.Variables()
	checkErr(ctx, err)

	db, err := mongocfg.NewDatabase(vars.MongoUri, vars.DbName)
	checkErr(ctx, err)

	jobPostingModel := &jobposting.JobPostingInfo{}
	jobPostingCollection := db.Collection(jobPostingModel.Collection())
	err = mongocfg.CheckIndexViaCollection(jobPostingCollection, jobPostingModel.IndexModels())
	checkErr(ctx, err)
	jobPostingRepo := rpcRepo.NewJobPostingRepo(jobPostingCollection)

	companyModel := &company.Company{}
	companyCollection := db.Collection(companyModel.Collection())
	err = mongocfg.CheckIndexViaCollection(companyCollection, companyModel.IndexModels())
	checkErr(ctx, err)
	companyRepo := rpcRepo.NewCompanyRepo(companyCollection)

	skillModel := &skill.Skill{}
	skillCollection := db.Collection(skillModel.Collection())
	err = mongocfg.CheckIndexViaCollection(skillCollection, skillModel.IndexModels())
	checkErr(ctx, err)

	skillRepo := rpcRepo.NewSkillRepo(skillCollection)
	checkErr(ctx, err)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", vars.GRPC_PORT))
	checkErr(ctx, err)

	llog.Msg("Start gRPC server").Data("port", vars.GRPC_PORT).Log(context.Background())

	return listener, jobPostingRepo, companyRepo, skillRepo
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}
