package rpcRepo

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobPostingRepo struct {
	col *mongo.Collection
}

func NewJobPostingRepo(db *mongo.Database) *JobPostingRepo {
	return &JobPostingRepo{
		col: db.Collection((&jobposting.JobPostingInfo{}).Collection()),
	}
}

func (jpRepo *JobPostingRepo) SaveHiring(ctx context.Context, jobPosting *jobposting.JobPostingInfo) (bool, error) {
	jobPosting.Status = jobposting.HIRING
	jobPosting.IsScanComplete = false
	now := time.Now()
	jobPosting.UpdatedAt = now
	jobPosting.InsertedAt = time.Time{}

	_, err := jpRepo.col.UpdateOne(ctx, bson.M{jobposting.SiteField: jobPosting.JobPostingId.Site, jobposting.PostingIdField: jobPosting.JobPostingId.PostingId},
		bson.M{
			"$setOnInsert": bson.M{
				jobposting.InsertedAtField: now,
			},
			"$set": jobPosting,
		},
		options.Update().SetUpsert(true))

	if err != nil {
		return false, err
	}

	return true, nil
}

func (jpRepo *JobPostingRepo) GetAllHiring(ctx context.Context, site string) ([]*jobposting.JobPostingId, error) {
	var jobPostings []*jobposting.JobPostingInfo

	cursor, err := jpRepo.col.Find(ctx, bson.M{jobposting.StatusField: jobposting.HIRING, jobposting.SiteField: site})
	if err != nil {
		if mongo.ErrNilDocument == err {
			return []*jobposting.JobPostingId{}, nil
		}
		return nil, terr.Wrap(err)
	}

	if err := cursor.All(context.Background(), &jobPostings); err != nil {
		return nil, terr.Wrap(err)
	}

	jobPostingIds := make([]*jobposting.JobPostingId, len(jobPostings))
	for i, jobPosting := range jobPostings {
		jobPostingIds[i] = &jobPosting.JobPostingId
	}

	return jobPostingIds, nil
}

func (jpRepo *JobPostingRepo) FindAll() ([]*jobposting.JobPostingInfo, error) {
	var jobPostings []*jobposting.JobPostingInfo

	cursor, err := jpRepo.col.Find(context.Background(), bson.D{})
	if err != nil {
		if mongo.ErrNilDocument == err {
			return []*jobposting.JobPostingInfo{}, nil
		}
		return nil, err
	}

	if err := cursor.All(context.Background(), &jobPostings); err != nil {
		return nil, err
	}

	return jobPostings, nil
}

func (jpRepo *JobPostingRepo) CloseAll(ctx context.Context, jobPostingIds []*jobposting.JobPostingId) error {
	if len(jobPostingIds) == 0 {
		return nil
	}

	_, err := jpRepo.col.UpdateMany(ctx, bson.M{
		"jobPostingId": bson.M{
			"$in": jobPostingIds,
		},
	}, bson.M{
		"$set": bson.M{
			jobposting.StatusField: jobposting.CLOSED,
		},
	})

	return err
}
