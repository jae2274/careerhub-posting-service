package rpcService

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/utils"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/provider_grpc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/rpcRepo"
)

type JobPostingService struct {
	jpRepo *rpcRepo.JobPostingRepo
}

func NewJobPostingService(jobPostingRepo *rpcRepo.JobPostingRepo) *JobPostingService {
	return &JobPostingService{
		jpRepo: jobPostingRepo,
	}
}

func (sv *JobPostingService) RegisterJobPostingInfo(ctx context.Context, msg *provider_grpc.JobPostingInfo) (bool, error) {
	publishedAt := utils.UnixMilliToTimePtr(msg.PublishedAt)
	closedAt := utils.UnixMilliToTimePtr(msg.ClosedAt)
	createdAt := utils.UnixMilliToTime(msg.CreatedAt)

	requiredSkills := make([]jobposting.RequiredSkill, len(msg.RequiredSkill))
	for i, skill := range msg.RequiredSkill {
		requiredSkills[i] = jobposting.RequiredSkill{
			SkillFrom: jobposting.Origin,
			SkillName: skill,
		}
	}

	jobPosting := jobposting.JobPostingInfo{
		JobPostingId: jobposting.JobPostingId{
			Site:      msg.JobPostingId.Site,
			PostingId: msg.JobPostingId.PostingId,
		},
		Status:      jobposting.HIRING,
		CompanyId:   msg.CompanyId,
		CompanyName: msg.CompanyName,
		JobCategory: msg.JobCategory,
		ImageUrl:    msg.ImageUrl,
		MainContent: jobposting.MainContent{
			PostUrl:        msg.MainContent.PostUrl,
			Title:          msg.MainContent.Title,
			Intro:          msg.MainContent.Intro,
			MainTask:       msg.MainContent.MainTask,
			Qualifications: msg.MainContent.Qualifications,
			Preferred:      msg.MainContent.Preferred,
			Benefits:       msg.MainContent.Benefits,
			RecruitProcess: msg.MainContent.RecruitProcess,
		},
		RequiredSkill: requiredSkills,
		Tags:          msg.Tags,
		RequiredCareer: jobposting.Career{
			Min: msg.RequiredCareer.Min,
			Max: msg.RequiredCareer.Max,
		},
		PublishedAt: publishedAt,
		ClosedAt:    closedAt,
		Address:     msg.Address,
		CreatedAt:   createdAt,
	}

	return sv.jpRepo.Save(ctx, &jobPosting)
}

func (sv *JobPostingService) CloseJobPostings(ctx context.Context, gJpId *provider_grpc.JobPostings) error {
	jpIds := make([]*jobposting.JobPostingId, len(gJpId.JobPostingIds))

	for i, gJpId := range gJpId.JobPostingIds {
		jpIds[i] = &jobposting.JobPostingId{
			Site:      gJpId.Site,
			PostingId: gJpId.PostingId,
		}
	}

	return sv.jpRepo.CloseAll(ctx, jpIds)
}
