package bgapp

import "github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/queue"

type BackgroundApp struct {
	companyQueue    *queue.CompanyQueue
	jobPostingQueue *queue.JobPostingQueue
	closedQueue     *queue.ClosedJobPostingQueue
}

func NewBackgroundApp(companyQueue *queue.CompanyQueue, jobPostingQueue *queue.JobPostingQueue, closedQueue *queue.ClosedJobPostingQueue) *BackgroundApp {
	return &BackgroundApp{
		companyQueue:    companyQueue,
		jobPostingQueue: jobPostingQueue,
		closedQueue:     closedQueue,
	}
}
