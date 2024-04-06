package tinit

import (
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/category"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/apirepo"
)

func InitRestApiJobPostingRepo(t *testing.T) apirepo.JobPostingRepo {
	db := InitDB(t)

	jobpostingCollection := db.Collection((&jobposting.JobPostingInfo{}).Collection())
	jobpostingRepo := apirepo.NewJobPostingRepo(jobpostingCollection)

	return jobpostingRepo
}

func InitRestApiCategoryRepo(t *testing.T) apirepo.CategoryRepo {
	db := InitDB(t)

	categoryCollection := db.Collection((&category.Category{}).Collection())
	categoryRepo := apirepo.NewCategoryRepo(categoryCollection)

	return categoryRepo
}

func InitRestApiSkillNameRepo(t *testing.T) apirepo.SkillRepo {
	db := InitDB(t)

	skillNameCollection := db.Collection((&skill.SkillName{}).Collection())
	skillNameRepo := apirepo.NewSkillRepo(skillNameCollection)

	return skillNameRepo
}
