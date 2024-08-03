package tinit

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/category"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/company"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/site"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/mongocfg"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/vars"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(t *testing.T) *mongo.Database {
	envVars, err := vars.Variables()
	checkError(t, err)

	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName, envVars.DBUser)
	checkError(t, err)

	initCollection(t, db, &jobposting.JobPostingInfo{})
	initCollection(t, db, &company.Company{})
	initCollection(t, db, &skill.Skill{})
	initCollection(t, db, &skill.SkillName{})
	initCollection(t, db, &category.Category{})
	initCollection(t, db, &site.Site{})

	return db
}

func initCollection(t *testing.T, db *mongo.Database, model mongocfg.MongoDBModel) {
	col := db.Collection(model.Collection())
	err := col.Drop(context.TODO())
	checkError(t, err)
	createIndexes(t, col, model.IndexModels())
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

func checkError(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
		t.Error(err)
		t.FailNow()
	}
}
