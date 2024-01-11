package tinit

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/mongocfg"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/vars"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/repo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(t *testing.T) *mongo.Database {
	envVars, err := vars.Variables()
	checkError(t, err)

	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName)
	checkError(t, err)

	jpModel := &jobposting.JobPostingInfo{}
	jpCol := db.Collection(jpModel.Collection())
	err = jpCol.Drop(context.TODO())
	checkError(t, err)
	createIndexes(t, jpCol, jpModel.IndexModels())

	companyModel := &company.Company{}
	companyCol := db.Collection(companyModel.Collection())
	err = companyCol.Drop(context.TODO())
	checkError(t, err)
	createIndexes(t, companyCol, companyModel.IndexModels())

	return db
}

func createIndexes(t *testing.T, col *mongo.Collection, indexModels map[string]*mongo.IndexModel) {
	for indexName, indexModel := range indexModels {
		if indexModel.Options == nil {
			indexModel.Options = options.Index()
		}
		indexModel.Options.Name = &indexName

		_, err := col.Indexes().CreateOne(context.TODO(), *indexModel, nil)
		checkError(t, err)
	}
}

func InitJobPostingRepo(t *testing.T) *repo.JobPostingRepo {
	db := InitDB(t)

	jobpostingCollection := db.Collection((&jobposting.JobPostingInfo{}).Collection())
	jobpostingRepo := repo.NewJobPostingRepo(jobpostingCollection)

	return jobpostingRepo
}

func InitCompanyRepo(t *testing.T) *repo.CompanyRepo {
	db := InitDB(t)

	companyCollection := db.Collection((&company.Company{}).Collection())
	companyRepo := repo.NewCompanyRepo(companyCollection)

	return companyRepo
}

func checkError(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
		t.Error(err)
		t.FailNow()
	}
}
