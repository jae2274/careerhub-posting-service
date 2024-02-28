package tinit

import (
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/apirepo"
)

func InitRestApiJobPostingRepo(t *testing.T) apirepo.JobPostingRepo {
	db := InitDB(t)

	jobpostingCollection := db.Collection((&jobposting.JobPostingInfo{}).Collection())
	jobpostingRepo := apirepo.NewJobPostingRepo(jobpostingCollection)

	return jobpostingRepo
}
