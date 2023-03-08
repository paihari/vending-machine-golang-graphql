package awscompose

import (
    "context"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/iam"
    "github.com/aws/aws-sdk-go-v2/service/sts"
)

func createIAMUserInChildAccount(managementAccessKeyID, managementSecretAccessKey, childAccountID, childRoleARN, childRegion, childUserName string) error {
    // Create a new session using the management account's access key ID and secret access key
    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion(childRegion),
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
            managementAccessKeyID,
            managementSecretAccessKey,
            "",
        )),
    )
    if err != nil {
        return err
    }

    // Assume a role in the child account
    stsClient := sts.NewFromConfig(cfg)
    assumeRoleOutput, err := stsClient.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
        RoleArn:         aws.String(childRoleARN),
        RoleSessionName: aws.String("mysession"),
    })
    if err != nil {
        return err
    }

    // Create a new session using the temporary credentials from the AssumeRole output
    cfg.Credentials = credentials.NewStaticCredentialsProvider(
        *assumeRoleOutput.Credentials.AccessKeyId,
        *assumeRoleOutput.Credentials.SecretAccessKey,
        *assumeRoleOutput.Credentials.SessionToken,
    )

    // Create an IAM client for the child account
    svc := iam.NewFromConfig(cfg)

    // Create a new IAM user in the child account
    _, err = svc.CreateUser(context.TODO(), &iam.CreateUserInput{
        UserName: aws.String(childUserName),
    })
    if err != nil {
        return err
    }

    // Attach a policy to the IAM user
    _, err = svc.AttachUserPolicy(context.TODO(), &iam.AttachUserPolicyInput{
        PolicyArn: aws.String("arn:aws:iam::aws:policy/ReadOnlyAccess"),
        UserName:  aws.String(childUserName),
    })
    if err != nil {
        return err
    }

    return nil
}
