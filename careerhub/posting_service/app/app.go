package app

import (
	"context"
	"os"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/category"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/company"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/mongocfg"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/vars"
	providergrpc "github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc"
	restapi "github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api"
	scannergrpc "github.com/jae2274/careerhub-posting-service/careerhub/posting_service/scanner_grpc"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/mw"
)

const (
	app     = "posting-service"
	service = "careerhub"

	ctxKeyTraceID = string(mw.CtxKeyTraceID)
)

func Run(ctx context.Context) {

	err := initLogger(ctx)
	checkErr(ctx, err)
	llog.Info(ctx, "Start Application")

	envVars, err := vars.Variables()
	checkErr(ctx, err)

	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName, envVars.DBUser)
	checkErr(ctx, err)

	err = mongocfg.InitCollections(db, &jobposting.JobPostingInfo{}, &company.Company{}, &skill.Skill{}, &skill.SkillName{}, &category.Category{})
	checkErr(ctx, err)

	runErr := make(chan error)
	go func() {
		err := providergrpc.Run(ctx, envVars.ProviderGrpcPort, db)
		runErr <- err
	}()

	go func() {
		err := scannergrpc.Run(ctx, envVars.ScannerGrpcPort, db)
		runErr <- err
	}()

	go func() {
		err := restapi.Run(ctx, envVars.RestApiGrpcPort, db)
		runErr <- err
	}()

	select {
	case <-ctx.Done():
		llog.Info(ctx, "Finish Application")
	case err := <-runErr:
		checkErr(ctx, err)
	}
}

func initLogger(ctx context.Context) error {
	llog.SetMetadata("service", service)
	llog.SetMetadata("app", app)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	llog.SetMetadata("hostname", hostname)

	return nil
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}
