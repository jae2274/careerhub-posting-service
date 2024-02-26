package gServer

import (
	"context"
	"io"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/scanner_grpc/repo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/scanner_grpc/scanner_grpc"
)

type ScannerServer struct {
	skillNameRepo  repo.SkillNameRepo
	jobPostingRepo repo.JobPostingRepo
	scanner_grpc.UnimplementedScannerServer
}

func (ss *ScannerServer) GetJobPostings(request *scanner_grpc.ScanComplete, sendSteam scanner_grpc.Scanner_GetJobPostingsServer) error {
	ctx := sendSteam.Context()
	jobPostingChan, err := ss.jobPostingRepo.GetJobPostings(ctx, request.IsScanComplete)
	if err != nil {
		return err
	}

	for jobPosting := range jobPostingChan {
		requiredSkills := make([]string, len(jobPosting.RequiredSkill))
		for i, requiredSkill := range jobPosting.RequiredSkill {
			requiredSkills[i] = requiredSkill.SkillName
		}

		err := sendSteam.Send(&scanner_grpc.JobPostingInfo{
			Site:           jobPosting.JobPostingId.Site,
			PostingId:      jobPosting.JobPostingId.PostingId,
			Title:          jobPosting.MainContent.Title,
			Qualifications: jobPosting.MainContent.Qualifications,
			Preferred:      jobPosting.MainContent.Preferred,
			RequiredSkill:  requiredSkills,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
func (ss *ScannerServer) GetSkills(ctx context.Context, request *scanner_grpc.ScanComplete) (*scanner_grpc.Skills, error) {
	skillNames, err := ss.skillNameRepo.GetSkills(ctx, request.IsScanComplete)
	if err != nil {
		return nil, err
	}

	return &scanner_grpc.Skills{SkillNames: skillNames}, nil
}
func (scanner *ScannerServer) SetRequiredSkills(recvStream scanner_grpc.Scanner_SetRequiredSkillsServer) error {
	ctx := recvStream.Context()
	for {
		jobPosting, err := recvStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		err = scanner.jobPostingRepo.AddRequiredSkills(ctx, jobposting.JobPostingId{Site: jobPosting.Site, PostingId: jobPosting.PostingId}, jobPosting.RequiredSkill)
		if err != nil {
			return err
		}
	}

	return recvStream.SendAndClose(&scanner_grpc.BoolResponse{Success: true})
}

func (scanner *ScannerServer) SetScanComplete(ctx context.Context, skills *scanner_grpc.Skills) (*scanner_grpc.BoolResponse, error) {
	err := scanner.skillNameRepo.SetScanComplete(ctx, skills.SkillNames)
	if err != nil {
		return nil, err
	}

	return &scanner_grpc.BoolResponse{Success: true}, nil
}
