package repo

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
)

type JobPostingRepo interface {
	GetJobPostings(ctx context.Context, isScanComplete bool) (<-chan jobposting.JobPostingInfo, error)
	AddRequiredSkills(ctx context.Context, jobPostingId jobposting.JobPostingId, requiredSkills []string) error
}
