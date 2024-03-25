package repo

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobPostingRepo interface {
	GetJobPostings(ctx context.Context, isScanComplete bool) (<-chan *jobposting.JobPostingInfo, error)
	AddRequiredSkills(ctx context.Context, jobPostingId jobposting.JobPostingId, requiredSkills []jobposting.RequiredSkill) error
}

type JobPostingRepoImpl struct {
	col *mongo.Collection
}

func NewJobPostingRepo(col *mongo.Collection) JobPostingRepo {
	return &JobPostingRepoImpl{col: col}
}

func (r *JobPostingRepoImpl) GetJobPostings(ctx context.Context, isScanComplete bool) (<-chan *jobposting.JobPostingInfo, error) {
	cursor, err := r.col.Find(ctx, bson.M{jobposting.IsScanCompleteField: isScanComplete})
	if err != nil {
		return nil, err
	}

	jobPostingInfoChan := make(chan *jobposting.JobPostingInfo)
	go func() {
		defer close(jobPostingInfoChan)
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var jobPostingInfo jobposting.JobPostingInfo
			if err := cursor.Decode(&jobPostingInfo); err != nil {
				return
			}
			jobPostingInfoChan <- &jobPostingInfo
		}
	}()

	return jobPostingInfoChan, nil
}

func (r *JobPostingRepoImpl) AddRequiredSkills(ctx context.Context, jobPostingId jobposting.JobPostingId, requiredSkills []jobposting.RequiredSkill) error {
	_, err := r.col.UpdateOne(ctx, bson.M{jobposting.SiteField: jobPostingId.Site, jobposting.PostingIdField: jobPostingId.PostingId}, bson.M{"$set": bson.M{jobposting.IsScanCompleteField: true}, "$addToSet": bson.M{jobposting.RequiredSkillField: bson.M{"$each": requiredSkills}}})
	return err
}
