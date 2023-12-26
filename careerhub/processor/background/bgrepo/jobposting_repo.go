package bgrepo

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobPostingRepo struct {
	col *mongo.Collection
}

func NewJobPostingRepo(col *mongo.Collection) *JobPostingRepo {
	return &JobPostingRepo{
		col: col,
	}
}

func (jpRepo *JobPostingRepo) Save(jobPosting *jobposting.JobPostingInfo) (bool, error) {
	// Convert decks to []interface{}

	_, err := jpRepo.col.InsertOne(context.TODO(), jobPosting)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) { // Ignore duplicate key error
			return false, nil
		}
		return false, err
	}

	return true, nil
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
