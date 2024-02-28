package apirepo

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobPostingRepo interface {
	GetJobPostings(ctx context.Context, page, size int) ([]jobposting.JobPostingInfo, error)
}

type JobPostingRepoImpl struct {
	col *mongo.Collection
}

func NewJobPostingRepo(jobPostingCollection *mongo.Collection) JobPostingRepo {
	return &JobPostingRepoImpl{
		col: jobPostingCollection,
	}
}

func (repo *JobPostingRepoImpl) GetJobPostings(ctx context.Context, page, size int) ([]jobposting.JobPostingInfo, error) {

	opt := options.Find().SetSkip(int64(page * size)).SetLimit(int64(size)).SetSort(bson.M{jobposting.CreatedAtField: -1})

	cursor, err := repo.col.Find(ctx, bson.D{}, opt)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	var jobPostings []jobposting.JobPostingInfo
	err = cursor.All(ctx, &jobPostings)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return jobPostings, nil
}
