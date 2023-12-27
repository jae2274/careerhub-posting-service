package bgapp

import (
	"fmt"
	"log"
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

func NewBackgroundApp(jpRepo *bgrepo.JobPostingRepo, companyQueue *queue.CompanyQueue, jobPostingQueue *queue.JobPostingQueue, closedQueue *queue.ClosedJobPostingQueue) *BackgroundApp {
	return &BackgroundApp{
		jpRepo:          jpRepo,
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

func (err ReceiptError) Error() string {
	return fmt.Sprintf("%s\treceipeHandle: %s", err.error.Error(), *err.ReceiptHandle)
}

func NewReceiptError(receiptHandle *string, err error) *ReceiptError {
	if err == nil {
		return nil
	}

	return &ReceiptError{
		ReceiptHandle: receiptHandle,
		error:         err,
	}
}

func (app *BackgroundApp) Run(quitChan <-chan QuitSignal) (<-chan WithRecipt[ProcessedSignal], <-chan error) {
	errChan := make(chan error, 100)
	msgChan := appfunc.NewJobPostingChannel(time.Minute, app.jobPostingQueue, errChan)
	step1 := pipe.NewStep(nil, func(msg *queue.Message) (WithRecipt[*message_v1.JobPostingInfo], *ReceiptError) {
		jobPostingMsg, err := appfunc.UnmarshalJobPosting(msg.Body)
		return NewWithRecipt(jobPostingMsg, msg.ReceiptHandle), NewReceiptError(msg.ReceiptHandle, err)
	})

	step2 := pipe.NewStep(nil, func(msg WithRecipt[*message_v1.JobPostingInfo]) (WithRecipt[bool], *ReceiptError) {
		ok, err := appfunc.SaveJobPosting(app.jpRepo, msg.Data)
		return NewWithRecipt(ok, msg.ReceiptHandle), NewReceiptError(msg.ReceiptHandle, err)
	})

	deleteMsgStep := pipe.NewStep(nil, func(msg WithRecipt[bool]) (WithRecipt[ProcessedSignal], *ReceiptError) {
		err := appfunc.DeleteJobPosting(app.jobPostingQueue, msg.ReceiptHandle)
		return NewWithRecipt(ProcessedSignal{}, msg.ReceiptHandle), NewReceiptError(msg.ReceiptHandle, err)
	})

	receiptErrChan := make(chan *ReceiptError, 100)
	processedChan := pipe.Pipeline3(msgChan, receiptErrChan, quitChan, step1, step2, deleteMsgStep)

	pipe.Transform(receiptErrChan, errChan, quitChan, nil, func(recpErr *ReceiptError) (WithRecipt[ProcessedSignal], error) { //이곳에 도달하는 에러는 그대로 다시 errChan으로 전달된다.
		//TODO: 데드 큐에 넣기
		log.Println("ReceiptError: ", recpErr.Error())
		err := appfunc.DeleteJobPosting(app.jobPostingQueue, recpErr.ReceiptHandle)
		if err != nil {
			return NewWithRecipt(ProcessedSignal{}, recpErr.ReceiptHandle), NewReceiptError(recpErr.ReceiptHandle, err)
		}

		return NewWithRecipt(ProcessedSignal{}, recpErr.ReceiptHandle), recpErr
	})

	return processedChan, errChan
}
