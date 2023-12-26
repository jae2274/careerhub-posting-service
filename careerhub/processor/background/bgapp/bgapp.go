package bgapp

import (
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/bgapp/appfunc"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/bgrepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/queue"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/queue/message_v1"
	"github.com/jae2274/goutils/cchan/pipe"
)

type QuitSignal struct{}
type ProcessedSignal struct{}

type BackgroundApp struct {
	jpRepo          *bgrepo.JobPostingRepo
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

type WithRecipt[DATA any] struct {
	Data          DATA
	ReceiptHandle *string
}

func NewWithRecipt[DATA any](data DATA, receiptHandle *string) WithRecipt[DATA] {
	return WithRecipt[DATA]{
		Data:          data,
		ReceiptHandle: receiptHandle,
	}
}

type ReceiptError struct {
	ReceiptHandle *string
	error
}

func NewReceiptError(receiptHandle *string, err error) ReceiptError {
	return ReceiptError{
		ReceiptHandle: receiptHandle,
		error:         err,
	}
}

func (app *BackgroundApp) Run(quitChan <-chan QuitSignal) error {
	msgChan, recvErrChan := appfunc.NewJobPostingChannel(time.Minute, app.jobPostingQueue)
	step1 := pipe.NewStep(nil, func(msg *queue.Message) (WithRecipt[*message_v1.JobPostingInfo], ReceiptError) {
		jobPostingMsg, err := appfunc.UnmarshalJobPosting(msg.Body)
		return NewWithRecipt(jobPostingMsg, msg.ReceiptHandle), NewReceiptError(msg.ReceiptHandle, err)
	})

	step2 := pipe.NewStep(nil, func(msg WithRecipt[*message_v1.JobPostingInfo]) (WithRecipt[bool], ReceiptError) {
		ok, err := appfunc.SaveJobPosting(app.jpRepo, msg.Data)
		return NewWithRecipt(ok, msg.ReceiptHandle), NewReceiptError(msg.ReceiptHandle, err)
	})

	step3 := pipe.NewStep(nil, func(msg WithRecipt[bool]) (ProcessedSignal, ReceiptError) {
		err := appfunc.DeleteJobPosting(app.jobPostingQueue, msg.ReceiptHandle)
		return ProcessedSignal{}, NewReceiptError(msg.ReceiptHandle, err)
	})

	processedChan, receiptErrChan := pipe.Pipeline3(msgChan, quitChan, 100, step1, step2, step3)

}
