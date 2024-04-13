package repo

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostingRepo interface {
	GetPostings(ctx context.Context, minInsertedAt time.Time, maxInsertedAt time.Time) ([]*jobposting.JobPostingInfo, error)
}

type PostingRepoImpl struct {
	col *mongo.Collection
}

func NewPostingRepo(db *mongo.Database) PostingRepo {
	col := db.Collection((&jobposting.JobPostingInfo{}).Collection())
	return &PostingRepoImpl{col: col}
}

func (r *PostingRepoImpl) GetPostings(ctx context.Context, minInsertedAt time.Time, maxInsertedAt time.Time) ([]*jobposting.JobPostingInfo, error) {
	filter := bson.M{jobposting.InsertedAtField: bson.M{"$gte": minInsertedAt, "$lt": maxInsertedAt}}

	opts := options.Find().SetProjection(
		bson.M{
			jobposting.SiteField:                    1,
			jobposting.PostingIdField:               1,
			jobposting.MainContent_TitleField:       1,
			jobposting.CompanyIdField:               1,
			jobposting.CompanyNameField:             1,
			jobposting.JobCategoryField:             1,
			jobposting.RequiresSkill_SkillNameField: 1,
			jobposting.MinCareerField:               1,
			jobposting.MaxCareerField:               1,
		},
	)
	cursor, err := r.col.Find(ctx, filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []*jobposting.JobPostingInfo{}, nil
		}
	}

	var postings []*jobposting.JobPostingInfo
	if err = cursor.All(ctx, &postings); err != nil {
		return nil, err
	}

	return postings, nil
}
