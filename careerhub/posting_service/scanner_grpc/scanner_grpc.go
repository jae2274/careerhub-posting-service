package scannergrpc

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/scanner_grpc/repo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/scanner_grpc/scanner_grpc"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/scanner_grpc/scanner_server"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, grpcPort int, collections map[string]*mongo.Collection) error {

	jobPostingCollection := collections[(&jobposting.JobPostingInfo{}).Collection()]
	skillNameCollection := collections[(&skill.SkillName{}).Collection()]

	jobPostingRepo := repo.NewJobPostingRepo(jobPostingCollection)
	skillNameRepo := repo.NewSkillNameRepo(skillNameCollection)

	scannerGrpcServer := scanner_server.NewScannerServer(skillNameRepo, jobPostingRepo)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Starting Scanner server...").Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer()
	scanner_grpc.RegisterScannerGrpcServer(grpcServer, scannerGrpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
