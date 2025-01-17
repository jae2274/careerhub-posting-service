package vars

import (
	"fmt"
	"os"
	"strconv"
)

type DBUser struct {
	Username string
	Password string
}

type Vars struct {
	MongoUri          string
	DbName            string
	DBUser            *DBUser
	ProviderGrpcPort  int
	ScannerGrpcPort   int
	RestApiGrpcPort   int
	SuggesterGrpcPort int
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

	dbUsername := getFromEnvPtr("DB_USERNAME")
	dbPassword := getFromEnvPtr("DB_PASSWORD")

	var dbUser *DBUser
	if dbUsername != nil && dbPassword != nil {
		dbUser = &DBUser{
			Username: *dbUsername,
			Password: *dbPassword,
		}
	}

	dbName, err := getFromEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	providerGrpcPort, err := getFromEnv("PROVIDER_GRPC_PORT")
	if err != nil {
		return nil, err
	}

	scannerGrpcPort, err := getFromEnv("SCANNER_GRPC_PORT")
	if err != nil {
		return nil, err
	}

	restApiGrpcPort, err := getFromEnv("RESTAPI_GRPC_PORT")
	if err != nil {
		return nil, err
	}

	suggesterGrpcPort, err := getFromEnv("SUGGESTER_GRPC_PORT")
	if err != nil {
		return nil, err
	}

	grpcPortInt, err := strconv.ParseInt(providerGrpcPort, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("GRPC_PORT is not integer.\tGRPC_PORT: %s", providerGrpcPort)
	}

	scannerGrpcPortInt, err := strconv.ParseInt(scannerGrpcPort, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("SCANNER_GRPC_PORT is not integer.\tSCANNER_GRPC_PORT: %s", scannerGrpcPort)
	}

	restApiGrpcPortInt, err := strconv.ParseInt(restApiGrpcPort, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("RESTAPI_GRPC_PORT is not integer.\tREST_API_PORT: %s", restApiGrpcPort)
	}

	suggesterGrpcPortInt, err := strconv.ParseInt(suggesterGrpcPort, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("SUGGESTER_GRPC_PORT is not integer.\tSUGGESTER_GRPC_PORT: %s", suggesterGrpcPort)
	}

	return &Vars{
		MongoUri:          mongoUri,
		DBUser:            dbUser,
		DbName:            dbName,
		ProviderGrpcPort:  int(grpcPortInt),
		ScannerGrpcPort:   int(scannerGrpcPortInt),
		RestApiGrpcPort:   int(restApiGrpcPortInt),
		SuggesterGrpcPort: int(suggesterGrpcPortInt),
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
