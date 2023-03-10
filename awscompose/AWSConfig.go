package awscompose

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

)

func GetAwsCredenctialConfig() (cfg aws.Config, err error){

	awsManagementAccessKey := os.Getenv("AWS_MANAGEMENT_ACCESS_KEY_ID")
	if awsManagementAccessKey == "" {
		return  aws.Config{}, fmt.Errorf("Error: AWS_MANAGEMENT_ACCESS_KEY_ID missing in .env %w", err)
	}

	awsManagementSecretKey := os.Getenv("AWS_MANAGEMENT_ACCESS_SECRET")

	if awsManagementSecretKey == "" {
		return  aws.Config{}, fmt.Errorf("Error: AWS_MANAGEMENT_ACCESS_SECRET missing in .env %w", err)

	}


	awsRegion := os.Getenv("AWS_REGION")

	if awsRegion == "" {
		return  aws.Config{}, fmt.Errorf("Error: AWS_MANAGEMENT_SECRET missing in .env %w", err)

	}
	return  config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(awsRegion),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				awsManagementAccessKey,
				awsManagementSecretKey,
				"",
			)),
		)
}