package awscfg

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/jae2274/goutils/terr"
)

func Config() (*aws.Config, error) {
	awsConfig, err := config.LoadDefaultConfig(context.Background())
	return &awsConfig, terr.Wrap(err)
}
