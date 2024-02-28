package tinit

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/company"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/skill"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/rpcRepo"
)

func InitProviderJobPostingRepo(t *testing.T) *rpcRepo.JobPostingRepo {
	db := InitDB(t)

	jobpostingCollection := db.Collection((&jobposting.JobPostingInfo{}).Collection())
	jobpostingRepo := rpcRepo.NewJobPostingRepo(jobpostingCollection)

	return jobpostingRepo
}

func InitProviderCompanyRepo(t *testing.T) *rpcRepo.CompanyRepo {
	db := InitDB(t)

	companyCollection := db.Collection((&company.Company{}).Collection())
	companyRepo := rpcRepo.NewCompanyRepo(companyCollection)

	return companyRepo
}

func InitProviderSkillRepo(t *testing.T) *rpcRepo.SkillRepo {
	db := InitDB(t)

	skillCollection := db.Collection((&skill.Skill{}).Collection())
	skillRepo := rpcRepo.NewSkillRepo(skillCollection)

	return skillRepo
}

func InitProviderSkillNameRepo(t *testing.T) *rpcRepo.SkillNameRepo {
	db := InitDB(t)

	skillNameCollection := db.Collection((&skill.SkillName{}).Collection())
	skillNameRepo := rpcRepo.NewSkillNameRepo(skillNameCollection)

	return skillNameRepo
}

func checkError(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
		t.Error(err)
		t.FailNow()
	}
}
