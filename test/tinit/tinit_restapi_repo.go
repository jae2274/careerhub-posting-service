package tinit

import (
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/category"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/skill"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/apirepo"
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

func InitRestApiSkillNameRepo(t *testing.T) apirepo.SkillNameRepo {
	db := InitDB(t)

	skillNameCollection := db.Collection((&skill.SkillName{}).Collection())
	skillNameRepo := apirepo.NewSkillNameRepo(skillNameCollection)

	return skillNameRepo
}
