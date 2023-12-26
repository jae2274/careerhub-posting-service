package vars

import (
	"fmt"
	"os"
)

type Vars struct {
	SqsEndpoint     *string
	JobPostingQueue string
	ClosedQueue     string
	CompanyQueue    string
	MongoUri        string
	DbName          string
}

type ErrNotExistedVar struct {
	VarName string
}

func NotExistedVar(varName string) *ErrNotExistedVar {
	return &ErrNotExistedVar{VarName: varName}
}

func (e *ErrNotExistedVar) Error() string {
	return fmt.Sprintf("%s is not existed", e.VarName)
}

func Variables() (*Vars, error) {
	mongoUri, err := getFromEnv("MONGO_URI")
	if err != nil {
		return nil, err
	}

	dbName, err := getFromEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	sqsEndpoint := getFromEnvPtr("SQS_ENDPOINT")

	companyQueue, err := getFromEnv("COMPANY_QUEUE")
	if err != nil {
		return nil, err
	}

	jobPostingQueue, err := getFromEnv("JOB_POSTING_QUEUE")
	if err != nil {
		return nil, err
	}

	closedQueue, err := getFromEnv("CLOSED_QUEUE")
	if err != nil {
		return nil, err
	}

	return &Vars{
		SqsEndpoint:     sqsEndpoint,
		JobPostingQueue: jobPostingQueue,
		ClosedQueue:     closedQueue,
		CompanyQueue:    companyQueue,
		MongoUri:        mongoUri,
		DbName:          dbName,
	}, nil
}

func getFromEnv(envVar string) (string, error) {
	ev := os.Getenv(envVar)

	if ev == "" {
		return "", fmt.Errorf("%s is not existed", envVar)
	}

	return ev, nil
}

func getFromEnvPtr(envVar string) *string {
	ev := os.Getenv(envVar)

	if ev == "" {
		return nil
	}

	return &ev
}
