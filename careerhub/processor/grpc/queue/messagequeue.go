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

func (cq *CompanyQueue) Send(message *message_v1.Company) error {
	b, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	return cq.queue.Send(b)
}

type JobPostingQueue struct {
	queue Queue
}

func NewJobPostingQueue(queue Queue) *JobPostingQueue {
	return &JobPostingQueue{
		queue: queue,
	}
}

func (jpq *JobPostingQueue) Send(message *message_v1.JobPostingInfo) error {
	b, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	return jpq.queue.Send(b)
}

type ClosedJobPostingQueue struct {
	queue Queue
}

func NewClosedJobPostingQueue(queue Queue) *ClosedJobPostingQueue {
	return &ClosedJobPostingQueue{
		queue: queue,
	}
}

func (cjpq *ClosedJobPostingQueue) Send(message *message_v1.ClosedJobPostings) error {
	b, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	return cjpq.queue.Send(b)
}
