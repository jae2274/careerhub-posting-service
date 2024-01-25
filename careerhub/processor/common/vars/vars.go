package vars

import (
	"fmt"
	"os"
	"strconv"
)

type Vars struct {
	MongoUri   string
	DbName     string
	GRPC_PORT  int
	PostLogUrl string
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

	grpcPort, err := getFromEnv("GRPC_PORT")
	if err != nil {
		return nil, err
	}

	postLogUrl, err := getFromEnv("POST_LOG_URL")
	if err != nil {
		return nil, err
	}

	grpcPortInt, err := strconv.ParseInt(grpcPort, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("GRPC_PORT is not integer.\tGRPC_PORT: %s", grpcPort)
	}

	return &Vars{
		MongoUri:   mongoUri,
		DbName:     dbName,
		GRPC_PORT:  int(grpcPortInt),
		PostLogUrl: postLogUrl,
	}, nil
}

func getFromEnv(envVar string) (string, error) {
	ev := os.Getenv(envVar)

	if ev == "" {
		return "", fmt.Errorf("%s is not existed", envVar)
	}

	return ev, nil
}
