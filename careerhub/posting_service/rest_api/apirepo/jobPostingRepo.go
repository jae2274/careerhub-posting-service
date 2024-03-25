package apirepo

import (
	"context"

	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/rest_api/dto"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobPostingRepo interface {
	GetJobPostings(ctx context.Context, page, size int, query dto.QueryReq) ([]jobposting.JobPostingInfo, error)
	GetJobPostingDetail(ctx context.Context, site, postingId string) (*jobposting.JobPostingInfo, error)
}

type JobPostingRepoImpl struct {
	col *mongo.Collection
}

func NewJobPostingRepo(jobPostingCollection *mongo.Collection) JobPostingRepo {
	return &JobPostingRepoImpl{
		col: jobPostingCollection,
	}
}

func (repo *JobPostingRepoImpl) GetJobPostings(ctx context.Context, page, size int, query dto.QueryReq) ([]jobposting.JobPostingInfo, error) {

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

func createFilter(query dto.QueryReq) bson.M {
	filter := bson.M{}

	if len(query.Categories) > 0 {
		categories := make([]bson.M, len(query.Categories))
		for i, category := range query.Categories {
			categories[i] = bson.M{jobposting.SiteField: category.Site, jobposting.JobCategoryField: category.CategoryName}
		}
		filter["$or"] = categories
	}

	if len(query.SkillNames) > 0 {
		fromPrefferedSkills := make([]bson.M, len(query.SkillNames))
		for i, skillName := range query.SkillNames {
			fromPrefferedSkills[i] = bson.M{jobposting.RequiresSkill_SkillNameField: skillName, jobposting.RequiredSkill_SkillFromField: jobposting.FromPreferred}
		}
		filter["$nor"] = fromPrefferedSkills

		filter[jobposting.RequiresSkill_SkillNameField] = bson.M{"$all": query.SkillNames}
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
