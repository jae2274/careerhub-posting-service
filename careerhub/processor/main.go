package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/skill"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/mongocfg"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/vars"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/gServer"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/processor_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/rpcRepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/rpcService"
	"google.golang.org/grpc"
)

func main() {
	log.Default().Println("Starting data processor...")
	vars, err := vars.Variables()
	checkErr(err)

	db, err := mongocfg.NewDatabase(vars.MongoUri, vars.DbName)
	checkErr(err)

	jobPostingModel := &jobposting.JobPostingInfo{}
	jobPostingCollection := db.Collection(jobPostingModel.Collection())
	err = mongocfg.CheckIndexViaCollection(jobPostingCollection, jobPostingModel.IndexModels())
	checkErr(err)
	jobPostingRepo := rpcRepo.NewJobPostingRepo(jobPostingCollection)

	companyModel := &company.Company{}
	companyCollection := db.Collection(companyModel.Collection())
	err = mongocfg.CheckIndexViaCollection(companyCollection, companyModel.IndexModels())
	checkErr(err)
	companyRepo := rpcRepo.NewCompanyRepo(companyCollection)

	skillModel := &skill.Skill{}
	skillCollection := db.Collection(skillModel.Collection())
	err = mongocfg.CheckIndexViaCollection(skillCollection, skillModel.IndexModels())
	checkErr(err)
	skillRepo := rpcRepo.NewSkillRepo(skillCollection)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", vars.GRPC_PORT))
	checkErr(err)

	grpcServer := grpc.NewServer()
	dataProcessorServer := gServer.NewDataProcessorServer(
		rpcService.NewJobPostingService(jobPostingRepo),
		rpcService.NewCompanyService(companyRepo),
		rpcService.NewSkillService(skillRepo),
	)

	processor_grpc.RegisterDataProcessorServer(grpcServer, dataProcessorServer) //client가 사용할 수 있도록 등록

	log.Printf("gRPC server is running on port %d...", vars.GRPC_PORT)
	err = grpcServer.Serve(listener)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
