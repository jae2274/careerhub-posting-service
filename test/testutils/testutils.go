package testutils

import (
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetIgnoreJobPostingFields(jobPostings []*jobposting.JobPostingInfo) {
	for _, jobPosting := range jobPostings {
		jobPosting.ID = primitive.ObjectID{} // ignore ID
		jobPosting.InsertedAt = time.Unix(jobPosting.InsertedAt.Unix(), 0)
		jobPosting.UpdatedAt = time.Unix(jobPosting.UpdatedAt.Unix(), 0)
		jobPosting.CreatedAt = time.Unix(jobPosting.CreatedAt.Unix(), 0)

		if jobPosting.PublishedAt != nil {
			*jobPosting.PublishedAt = time.Unix(jobPosting.PublishedAt.Unix(), 0)
		}
		if jobPosting.ClosedAt != nil {
			*jobPosting.ClosedAt = time.Unix(jobPosting.ClosedAt.Unix(), 0)
		}
	}
}
