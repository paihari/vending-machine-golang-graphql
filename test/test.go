package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func createIAMUserInChildAccount(managementAccessKeyID, managementSecretAccessKey, childAccountID, childRegion, childUserName string) error {
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

    // Create an IAM client for the child account
    svc := iam.NewFromConfig(cfg)

    // Create a new IAM user in the child account
    user, err := svc.CreateUser(context.TODO(), &iam.CreateUserInput{
        UserName: aws.String(childUserName),
    })
    if err != nil {
        return err
    }

	fmt.Println("User")
	fmt.Println(user)

    // // Attach a policy to the IAM user
    // _, err = svc.AttachUserPolicy(context.TODO(), &iam.AttachUserPolicyInput{
    //     PolicyArn: aws.String("arn:aws:iam::aws:policy/ReadOnlyAccess"),
    //     UserName:  aws.String(childUserName),
    // })
    // if err != nil {
    //     return err
    // }

    return nil
}

func main() {
  err := createIAMUserInChildAccount("AKIA5CFTTGW5M2ZUHTJX", "j5S4H4x/S5NfXDmqmrMZrwQc1Q6EMw3m2lSijcUx", "161987549706", "us-east-1", "cloud-control-iam-user")
  fmt.Println(err)
}
