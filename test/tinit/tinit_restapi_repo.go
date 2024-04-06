package tinit

import (
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
)

func InitRestApiJobPostingRepo(t *testing.T) apirepo.JobPostingRepo {
	db := InitDB(t)

	jobpostingRepo := apirepo.NewJobPostingRepo(db)

	return jobpostingRepo
}

func InitRestApiCategoryRepo(t *testing.T) apirepo.CategoryRepo {
	db := InitDB(t)

	categoryRepo := apirepo.NewCategoryRepo(db)

	return categoryRepo
}

func InitRestApiSkillNameRepo(t *testing.T) apirepo.SkillRepo {
	db := InitDB(t)

	skillNameRepo := apirepo.NewSkillRepo(db)

	return skillNameRepo
}
