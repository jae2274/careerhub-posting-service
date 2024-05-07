package tinit

import (
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/provider_grpc"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/rest_api/restapi_grpc"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/suggester/suggester_grpc"
)

func InitSuggesterClient(t *testing.T) suggester_grpc.PostingClient {
	envVars := InitEnvVars(t)
	conn := InitGrpcClient(t, envVars.SuggesterGrpcPort)

	return suggester_grpc.NewPostingClient(conn)
}

func InitRestapiGrpcClient(t *testing.T) restapi_grpc.RestApiGrpcClient {
	envVars := InitEnvVars(t)
	conn := InitGrpcClient(t, envVars.RestApiGrpcPort)

	return restapi_grpc.NewRestApiGrpcClient(conn)
}

func InitProviderGrpcClient(t *testing.T) provider_grpc.ProviderGrpcClient {
	envVars := InitEnvVars(t)
	conn := InitGrpcClient(t, envVars.ProviderGrpcPort)

	return provider_grpc.NewProviderGrpcClient(conn)
}
