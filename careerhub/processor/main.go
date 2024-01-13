package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/mongocfg"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/vars"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/gServer"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/processor_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/rpcRepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/rpcService"
	"google.golang.org/grpc"
)

func main() {
	vars, err := vars.Variables()
	checkErr(err)

	db, err := mongocfg.NewDatabase(vars.MongoUri, vars.DbName)
	checkErr(err)

	jobPostingRepo := rpcRepo.NewJobPostingRepo(db.Collection(
		(&jobposting.JobPostingInfo{}).Collection(),
	))

	companyRepo := rpcRepo.NewCompanyRepo(db.Collection(
		(&company.Company{}).Collection(),
	))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", vars.GRPC_PORT))
	checkErr(err)

	grpcServer := grpc.NewServer()
	dataProcessorServer := gServer.NewDataProcessorServer(
		rpcService.NewJobPostingService(jobPostingRepo),
		rpcService.NewCompanyService(companyRepo),
	)

	processor_grpc.RegisterDataProcessorServer(grpcServer, dataProcessorServer) //client가 사용할 수 있도록 등록

	err = grpcServer.Serve(listener)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
