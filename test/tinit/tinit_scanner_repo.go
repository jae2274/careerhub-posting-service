package tinit

import (
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/scanner_grpc/repo"
)

func InitScannerSkillNameRepo(t *testing.T) repo.SkillNameRepo {
	db := InitDB(t)

	skillNameRepo := repo.NewSkillNameRepo(db)

	return skillNameRepo
}

func InitScannerJobPostingRepo(t *testing.T) repo.JobPostingRepo {
	db := InitDB(t)

	jobPostingRepo := repo.NewJobPostingRepo(db)

	return jobPostingRepo
}
