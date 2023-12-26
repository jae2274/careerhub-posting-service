package tinit

import (
	"context"
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/bgrepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/mongocfg"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/vars"
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

func InitBgJobPostingRepo(t *testing.T) *bgrepo.JobPostingRepo {
	db := InitDB(t)

	jobpostingCollection := db.Collection((&jobposting.JobPostingInfo{}).Collection())
	jobpostingRepo := bgrepo.NewJobPostingRepo(jobpostingCollection)

	return jobpostingRepo
}
