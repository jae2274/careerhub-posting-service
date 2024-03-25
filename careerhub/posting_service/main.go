package main

import (
	"context"
	"os"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/category"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/company"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/mongocfg"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/vars"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/logger"
	providergrpc "github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc"
	restapi "github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api"
	scannergrpc "github.com/jae2274/careerhub-posting-service/careerhub/posting_service/scanner_grpc"

	"github.com/jae2274/goutils/llog"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	app     = "posting-service"
	service = "careerhub"

	ctxKeyTraceID = "trace_id"
)

func main() {
	ctx := context.Background()
	envVars, err := vars.Variables()
	checkErr(ctx, err)

	err = initLogger(ctx, envVars.PostLogUrl)
	checkErr(ctx, err)

	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName, envVars.DBUser)
	checkErr(ctx, err)

	collections, err := initCollections(db)
	checkErr(ctx, err)

	runErr := make(chan error)
	go func() {
		err := providergrpc.Run(ctx, envVars.ProviderGrpcPort, collections)
		runErr <- err
	}()

	go func() {
		err := scannergrpc.Run(ctx, envVars.ScannerGrpcPort, collections)
		runErr <- err
	}()

	go func() {
		err := restapi.Run(ctx, envVars.RestApiPort, envVars.RootPath, collections)
		runErr <- err
	}()

	err = <-runErr
	checkErr(ctx, err)
}

func initLogger(ctx context.Context, postUrl string) error {
	llog.SetMetadata("service", service)
	llog.SetMetadata("app", app)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	llog.SetMetadata("hostname", hostname)

	appLogger, err := logger.NewAppLogger(ctx, postUrl)
	if err != nil {
		return err
	}

	llog.SetDefaultLLoger(appLogger)

	return nil
}

func initCollections(db *mongo.Database) (map[string]*mongo.Collection, error) {
	collections := make(map[string]*mongo.Collection)

	jobPostingModel := &jobposting.JobPostingInfo{}
	jobPostingCollection := db.Collection(jobPostingModel.Collection())
	err := mongocfg.CheckIndexViaCollection(jobPostingCollection, jobPostingModel.IndexModels())
	if err != nil {
		return nil, err
	}
	collections[jobPostingModel.Collection()] = jobPostingCollection

	companyModel := &company.Company{}
	companyCollection := db.Collection(companyModel.Collection())
	err = mongocfg.CheckIndexViaCollection(companyCollection, companyModel.IndexModels())
	if err != nil {
		return nil, err
	}
	collections[companyModel.Collection()] = companyCollection

	skillModel := &skill.Skill{}
	skillCollection := db.Collection(skillModel.Collection())
	err = mongocfg.CheckIndexViaCollection(skillCollection, skillModel.IndexModels())
	if err != nil {
		return nil, err
	}
	collections[skillModel.Collection()] = skillCollection

	skillNameModel := &skill.SkillName{}
	skillNameCollection := db.Collection(skillNameModel.Collection())
	err = mongocfg.CheckIndexViaCollection(skillNameCollection, skillNameModel.IndexModels())
	if err != nil {
		return nil, err
	}
	collections[skillNameModel.Collection()] = skillNameCollection

	categoryModel := &category.Category{}
	categoryCollection := db.Collection(categoryModel.Collection())
	err = mongocfg.CheckIndexViaCollection(categoryCollection, categoryModel.IndexModels())
	if err != nil {
		return nil, err
	}
	collections[categoryModel.Collection()] = categoryCollection

	return collections, nil
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}
