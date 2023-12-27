package appfunc

import (
	"time"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/bgrepo"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/queue"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/background/queue/message_v1"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/jobposting"
	"google.golang.org/protobuf/proto"
)

func NewJobPostingChannel(delayTime time.Duration, jpQueue *queue.JobPostingQueue, errChan chan<- error) <-chan *queue.Message {
	msgsChan := make(chan *queue.Message)

	go func() {
		for {
			msgs, err := jpQueue.Recv()

			if err != nil {
				errChan <- err
				continue
			}

			if len(msgs) > 0 {
				for _, msg := range msgs {
					msgsChan <- msg
				}
			} else {
				time.Sleep(delayTime)
			}

		}
	}()

	return msgsChan
}

func UnmarshalJobPosting(bMsg []byte) (*message_v1.JobPostingInfo, error) {
	var jobPosting message_v1.JobPostingInfo
	err := proto.Unmarshal(bMsg, &jobPosting)
	if err != nil {
		return nil, err
	}

	return &jobPosting, nil
}

func SaveJobPosting(bgJpRepo *bgrepo.JobPostingRepo, msg *message_v1.JobPostingInfo) (bool, error) {
	jobPosting := jobposting.JobPostingInfo{
		Site:        msg.Site,
		PostingId:   msg.PostingId,
		CompanyId:   msg.CompanyId,
		CompanyName: msg.CompanyName,
		JobCategory: msg.JobCategory,
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
		RequiredSkill: msg.RequiredSkill,
		Tags:          msg.Tags,
		RequiredCareer: jobposting.Career{
			Min: msg.RequiredCareer.Min,
			Max: msg.RequiredCareer.Max,
		},
		PublishedAt: msg.PublishedAt,
		ClosedAt:    msg.ClosedAt,
		Address:     msg.Address,
	}

	return bgJpRepo.Save(&jobPosting)
}

func DeleteJobPosting(jpQueue *queue.JobPostingQueue, recipeHandle *string) error {
	return jpQueue.Delete(recipeHandle)
}
