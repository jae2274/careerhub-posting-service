package tinit

import (
	"testing"

	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/scanner_grpc/repo"
)

func InitScannerSkillNameRepo(t *testing.T) repo.SkillNameRepo {
	db := InitDB(t)

	skillNameCollection := db.Collection((&skill.SkillName{}).Collection())
	skillNameRepo := repo.NewSkillNameRepo(skillNameCollection)

	return skillNameRepo
}

func InitScannerJobPostingRepo(t *testing.T) repo.JobPostingRepo {
	db := InitDB(t)

	jobPostingCollection := db.Collection((&jobposting.JobPostingInfo{}).Collection())
	jobPostingRepo := repo.NewJobPostingRepo(jobPostingCollection)

	return jobPostingRepo
}
