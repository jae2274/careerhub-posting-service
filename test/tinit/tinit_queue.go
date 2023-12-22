package tinit

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/awscfg"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/queue"
	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/vars"
	"github.com/jae2274/goutils/terr"
)

func InitSQS(t *testing.T, queueName string) queue.Queue {
	variables, err := vars.Variables()
	checkError(t, err)

	sqsEndpoint := variables.SqsEndpoint
	cfg, err := awscfg.Config()
	checkError(t, err)

	sqsClient := queue.NewClient(cfg, sqsEndpoint)

	truncateSQS(t, sqsClient, &queueName)

	queue, err := queue.NewSQS(cfg, variables.SqsEndpoint, queueName)
	checkError(t, err)

	return queue
}

func truncateSQS(t *testing.T, sqsClient *sqs.Client, queueName *string) *string {
	queueUrl := getQueueUrl(t, sqsClient, queueName)
	if queueUrl != nil {
		deleteQueue(t, sqsClient, queueUrl)
	}
	createQueue(t, sqsClient, queueName)
	return getQueueUrl(t, sqsClient, queueName)
}

func getQueueUrl(t *testing.T, sqsClient *sqs.Client, queueName *string) *string {
	result, err := sqsClient.GetQueueUrl(
		context.Background(),
		&sqs.GetQueueUrlInput{
			QueueName: queueName,
		},
	)

	var notExisted *types.QueueDoesNotExist
	if errors.As(terr.UnWrap(err), &notExisted) {
		return nil
	}
	checkError(t, err)
	return result.QueueUrl
}

func deleteQueue(t *testing.T, sqsClient *sqs.Client, queueUrl *string) {
	_, err := sqsClient.DeleteQueue(context.Background(), &sqs.DeleteQueueInput{
		QueueUrl: queueUrl,
	})
	checkError(t, err)
}

func createQueue(t *testing.T, sqsClient *sqs.Client, queueName *string) {
	_, err := sqsClient.CreateQueue(context.Background(), &sqs.CreateQueueInput{
		QueueName: queueName,
	})
	checkError(t, err)
}

func checkError(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d\n", file, line)
		t.Error(err)
		t.FailNow()
	}
}
