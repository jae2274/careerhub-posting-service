package apirepo

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_grpc"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobPostingRepo interface {
	GetJobPostings(ctx context.Context, page, size int32, query *restapi_grpc.QueryReq) ([]jobposting.JobPostingInfo, error)
	GetJobPostingDetail(ctx context.Context, site, postingId string) (*jobposting.JobPostingInfo, error)
	GetJobPostingsById(ctx context.Context, jobPostingIds []*restapi_grpc.JobPostingIdReq) ([]jobposting.JobPostingInfo, error)
}

type JobPostingRepoImpl struct {
	col *mongo.Collection
}

func NewJobPostingRepo(db *mongo.Database) JobPostingRepo {
	jobPostingCollection := db.Collection((&jobposting.JobPostingInfo{}).Collection())
	return &JobPostingRepoImpl{
		col: jobPostingCollection,
	}
}

func (repo *JobPostingRepoImpl) GetJobPostings(ctx context.Context, page, size int32, query *restapi_grpc.QueryReq) ([]jobposting.JobPostingInfo, error) {

	opt := options.Find().SetSkip(int64(page * size)).SetLimit(int64(size)).SetSort(bson.M{jobposting.CreatedAtField: -1})

	filter := createFilter(query)
	cursor, err := repo.col.Find(ctx, filter, opt)
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

func createFilter(query *restapi_grpc.QueryReq) bson.M {
	filter := bson.M{}

	if len(query.Categories) > 0 {
		categories := make([]bson.M, len(query.Categories))
		for i, category := range query.Categories {
			categories[i] = bson.M{jobposting.SiteField: category.Site, jobposting.JobCategoryField: category.CategoryName}
		}
		filter["$or"] = categories
	}

	if len(query.SkillNames) > 0 {
		and := make([]bson.M, len(query.SkillNames))
		for i, skillName := range query.SkillNames {
			and[i] = bson.M{
				jobposting.RequiredSkillField: bson.M{
					"$elemMatch": bson.M{
						jobposting.SkillNameField: bson.M{"$in": skillName.Or},
						jobposting.SkillFromField: bson.M{"$in": []jobposting.SkillFrom{jobposting.Origin, jobposting.FromTitle, jobposting.FromMainTask, jobposting.FromQualifications}},
					},
				},
			}
		}
		filter["$and"] = and
	}

	if query.MinCareer != nil {
		filter[jobposting.MinCareerField] = bson.M{"$gte": *query.MinCareer}
	}

	if query.MaxCareer != nil {
		filter[jobposting.MaxCareerField] = bson.M{"$lte": *query.MaxCareer}
	}

	filter[jobposting.StatusField] = jobposting.HIRING

	return filter
}

func (repo *JobPostingRepoImpl) GetJobPostingDetail(ctx context.Context, site, postingId string) (*jobposting.JobPostingInfo, error) {
	var jobPosting jobposting.JobPostingInfo
	err := repo.col.FindOne(ctx, bson.M{jobposting.SiteField: site, jobposting.PostingIdField: postingId}).Decode(&jobPosting)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, terr.Wrap(err)
	}

	return &jobPosting, nil
}

func (repo *JobPostingRepoImpl) GetJobPostingsById(ctx context.Context, jobPostingIds []*restapi_grpc.JobPostingIdReq) ([]jobposting.JobPostingInfo, error) {
	var ors bson.A
	for _, jobPostingId := range jobPostingIds {
		ors = append(ors, bson.M{jobposting.SiteField: jobPostingId.Site, jobposting.PostingIdField: jobPostingId.PostingId})
	}
	filter := bson.M{
		"$or": ors,
	}
	cur, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	var jobPostings []jobposting.JobPostingInfo
	err = cur.All(ctx, &jobPostings)
	if err != nil {
		return nil, terr.Wrap(err)
	}

	return jobPostings, nil
}
