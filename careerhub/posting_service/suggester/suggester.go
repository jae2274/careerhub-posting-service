package suggester

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/suggester/repo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/suggester/suggester_grpc"
	suggesterserver "github.com/jae2274/careerhub-posting-service/careerhub/posting_service/suggester/suggester_server"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/utils"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, grpcPort int, db *mongo.Database) error {

	suggesterGrpcServer := suggesterserver.NewSuggesterService(repo.NewPostingRepo(db))

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Starting suggester grpc server...").Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer(utils.Middlewares()...)
	suggester_grpc.RegisterPostingServer(grpcServer, suggesterGrpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
