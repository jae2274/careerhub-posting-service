package tinit

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
)

func InitProviderJobPostingRepo(t *testing.T) *rpcRepo.JobPostingRepo {
	db := InitDB(t)

	jobpostingRepo := rpcRepo.NewJobPostingRepo(db)

	return jobpostingRepo
}

func InitProviderCompanyRepo(t *testing.T) *rpcRepo.CompanyRepo {
	db := InitDB(t)

	companyRepo := rpcRepo.NewCompanyRepo(db)

	return companyRepo
}

func InitProviderSkillRepo(t *testing.T) *rpcRepo.SkillRepo {
	db := InitDB(t)

	skillRepo := rpcRepo.NewSkillRepo(db)

	return skillRepo
}

func InitProviderSkillNameRepo(t *testing.T) *rpcRepo.SkillNameRepo {
	db := InitDB(t)

	skillNameRepo := rpcRepo.NewSkillNameRepo(db)

	return skillNameRepo
}

func InitProviderCategoryRepo(t *testing.T) *rpcRepo.CategoryRepo {
	db := InitDB(t)

	categoryRepo := rpcRepo.NewCategoryRepo(db)

	return categoryRepo
}

func checkError(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
		t.Error(err)
		t.FailNow()
	}
}
