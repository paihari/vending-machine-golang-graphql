package awscompose

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
    "github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func CreateIAMUserInChildAccount(childAccountID, childUserName, tag  string) (string, error) {
    
    cfg, err := GetAwsCredenctialConfig()

    if err != nil {
        fmt.Println(err)
        return "", err
    }

    // Assume a role in the child account
    stsClient := sts.NewFromConfig(cfg)

    childRoleARN := "arn:aws:iam::" + childAccountID + ":role/OrganizationAccountAccessRole" 


    assumeRoleOutput, err := stsClient.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
        RoleArn:         aws.String(childRoleARN),
        RoleSessionName: aws.String("mysession"),
    })
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    // Create a new session using the temporary credentials from the AssumeRole output
    cfg.Credentials = credentials.NewStaticCredentialsProvider(
        *assumeRoleOutput.Credentials.AccessKeyId,
        *assumeRoleOutput.Credentials.SecretAccessKey,
        *assumeRoleOutput.Credentials.SessionToken,
    )

    // Create an IAM client for the child account
    svc := iam.NewFromConfig(cfg)


    tags := []types.Tag{
        {
            Key:   aws.String("Environment"),
            Value: aws.String("Production"),
        },
        {
            Key:   aws.String("Team"),
            Value: aws.String("Engineering"),
        },
    }

    // Create a new IAM user in the child account
    user, err := svc.CreateUser(context.TODO(), &iam.CreateUserInput{
        UserName: aws.String(childUserName),
        Tags:     tags,
    })
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    // Attach a policy to the IAM user
    _, err = svc.AttachUserPolicy(context.TODO(), &iam.AttachUserPolicyInput{
        PolicyArn: aws.String("arn:aws:iam::aws:policy/ReadOnlyAccess"),
        UserName:  aws.String(childUserName),
    })
    if err != nil {
        fmt.Println(err)
        return "", err
    }

    return *user.User.Arn, nil
}
