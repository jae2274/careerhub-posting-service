package queue

import (
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/queue/message_v1"
	"google.golang.org/protobuf/proto"
)

type CompanyQueue struct {
	queue Queue
}

func NewCompanyQueue(queue Queue) *CompanyQueue {
	return &CompanyQueue{
		queue: queue,
	}
}

func (cq *CompanyQueue) Recv() ([]*message_v1.Company, error) {
	bMsgs, err := cq.queue.Recv()
	if err != nil {
		return nil, err
	}

	var companies []*message_v1.Company
	for _, bMsg := range bMsgs {
		var company message_v1.Company
		err := proto.Unmarshal(bMsg, &company)
		if err != nil {
			return nil, err
		}
		companies = append(companies, &company)
	}

	return companies, nil
}

type JobPostingQueue struct {
	queue Queue
}

func NewJobPostingQueue(queue Queue) *JobPostingQueue {
	return &JobPostingQueue{
		queue: queue,
	}
}

func (jpq *JobPostingQueue) Recv() ([]*message_v1.JobPostingInfo, error) {
	bMsgs, err := jpq.queue.Recv()
	if err != nil {
		return nil, err
	}

	var jobPostings []*message_v1.JobPostingInfo
	for _, bMsg := range bMsgs {
		var jobPosting message_v1.JobPostingInfo
		err := proto.Unmarshal(bMsg, &jobPosting)
		if err != nil {
			return nil, err
		}
		jobPostings = append(jobPostings, &jobPosting)
	}

	return jobPostings, nil
}

type ClosedJobPostingQueue struct {
	queue Queue
}

func NewClosedJobPostingQueue(queue Queue) *ClosedJobPostingQueue {
	return &ClosedJobPostingQueue{
		queue: queue,
	}
}

func (cjpq *ClosedJobPostingQueue) Recv() ([]*message_v1.ClosedJobPostings, error) {
	bMsgs, err := cjpq.queue.Recv()
	if err != nil {
		return nil, err
	}

	var closedJobPostings []*message_v1.ClosedJobPostings
	for _, bMsg := range bMsgs {
		var closedJobPosting message_v1.ClosedJobPostings
		err := proto.Unmarshal(bMsg, &closedJobPosting)
		if err != nil {
			return nil, err
		}
		closedJobPostings = append(closedJobPostings, &closedJobPosting)
	}

	return closedJobPostings, nil
}
