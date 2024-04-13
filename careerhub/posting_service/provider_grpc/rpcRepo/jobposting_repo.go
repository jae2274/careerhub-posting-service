package rpcRepo

import (
	"context"
	"fmt"

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

func (jpRepo *JobPostingRepo) Save(ctx context.Context, jobPosting *jobposting.JobPostingInfo) (bool, error) {
	// Convert decks to []interface{}
	jobPosting.Status = jobposting.HIRING
	jobPosting.IsScanComplete = false

	result, err := jpRepo.col.UpdateOne(ctx, bson.M{jobposting.SiteField: jobPosting.JobPostingId.Site, jobposting.PostingIdField: jobPosting.JobPostingId.PostingId},
		bson.M{
			"$currentDate": bson.M{
				jobposting.UpdatedAtField: true,
			},
			"$set": jobPosting,
		},
		options.Update().SetUpsert(true))

	if err != nil {
		return false, err
	}

	if result.UpsertedID != nil { //본래 $setOnInsert를 사용하여 insertedAt을 설정하려 하였으나, $setOnInsert는 $currentDate와 다르게 mongodb의 기준으로 시간을 설정하는 방법이 존재하지 않아 아래와 같이 처리함
		//TODO: mongodb 시간 기준으로 insertedAt과 updatedAt을 동시에 설정하는 방법 찾기. time.Now() 금지!
		_, err = jpRepo.col.UpdateOne(ctx, bson.M{jobposting.SiteField: jobPosting.JobPostingId.Site, jobposting.PostingIdField: jobPosting.JobPostingId.PostingId},
			bson.A{
				bson.M{
					"$set": bson.M{
						jobposting.InsertedAtField: fmt.Sprintf("$%s", jobposting.UpdatedAtField),
					},
				},
			},
		)
		if err != nil {
			return false, err
		}
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
