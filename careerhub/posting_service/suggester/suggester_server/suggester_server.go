package suggesterserver

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/jobposting"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/suggester/repo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/suggester/suggester_grpc"
)

type SuggesterService struct {
	postingRepo repo.PostingRepo
	suggester_grpc.UnimplementedPostingServer
}

func NewSuggesterService(postingRepo repo.PostingRepo) *SuggesterService {
	return &SuggesterService{
		postingRepo: postingRepo,
	}
}

func (s *SuggesterService) GetPostings(ctx context.Context, req *suggester_grpc.GetPostingsRequest) (*suggester_grpc.GetPostingsResponse, error) {
	minCreatedAt := time.UnixMilli(req.MinUnixMilli)
	maxCreatedAt := time.UnixMilli(req.MaxUnixMilli)

	postings, err := s.postingRepo.GetPostings(ctx, minCreatedAt, maxCreatedAt)

	if err != nil {
		return nil, err
	}

	grpcPostings := make([]*suggester_grpc.JobPosting, len(postings))
	for i, posting := range postings {
		grpcPostings[i] = ConvertJobPostingToGrpc(posting)
	}

	return &suggester_grpc.GetPostingsResponse{
		Postings: grpcPostings,
	}, nil
}

func ConvertJobPostingToGrpc(posting *jobposting.JobPostingInfo) *suggester_grpc.JobPosting {
	skillNames := make([]string, len(posting.RequiredSkill))
	for i, skill := range posting.RequiredSkill {
		skillNames[i] = skill.SkillName
	}

	return &suggester_grpc.JobPosting{
		Site:        posting.JobPostingId.Site,
		PostingId:   posting.JobPostingId.PostingId,
		Title:       posting.MainContent.Title,
		CompanyId:   posting.CompanyId,
		CompanyName: posting.CompanyName,
		Info: &suggester_grpc.PostingInfo{
			Categories: posting.JobCategory,
			SkillNames: skillNames,
			MinCareer:  posting.RequiredCareer.Min,
			MaxCareer:  posting.RequiredCareer.Max,
		},
		ImageUrl: posting.ImageUrl,
	}
}
