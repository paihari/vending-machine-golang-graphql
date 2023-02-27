package awscompose

import (
	"context"
	"errors"
	"fmt"
	"time"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"

)

func AttachFederationTagPolicyToResidentAccount() {

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		fmt.Println("Error loading the Config")
		return
	}

	//emailAddress := "pai2023022403@pai.ch"

	// Create Organizations client
	svc := organizations.NewFromConfig(cfg)

	arn := "arn:aws:organizations::898024814010:policy/o-py48lgs0w1/tag_policy/p-953m46re9l"
	parts := strings.Split(arn, "/")
	policyID := parts[len(parts)-1] // 12345678-abcd-1234-abcd-123456789012
	childAccountId := "161987549706"


	_, err = svc.AttachPolicy(context.TODO(), &organizations.AttachPolicyInput{
		PolicyId: &policyID,
		TargetId: &childAccountId,
	})

	fmt.Println("Policy attached successfully to child account.")

}



func CreateResidentAccount(accountName string, emailAddress string) (string, error) {
	// Load AWS config
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return "", fmt.Errorf("error loading AWS config: %w", err)
	}

	//emailAddress := "pai2023022403@pai.ch"

	// Create Organizations client
	svc := organizations.NewFromConfig(cfg)

	// Set up CreateAccount input parameters
    createAccountResult, err :=  svc.CreateAccount(context.TODO(), &organizations.CreateAccountInput{
        AccountName:  &accountName,
        Email:        &emailAddress,
        IamUserAccessToBilling: "DENY",
    })
	
	// Retrieve the account creation request ID
	createAccountRequestId := createAccountResult.CreateAccountStatus.Id

	// Wait for account creation process to complete

	for {
		describeCreateAccountInput := &organizations.DescribeCreateAccountStatusInput{
			CreateAccountRequestId: createAccountRequestId,
		}
		describeCreateAccountResult, err := svc.DescribeCreateAccountStatus(context.Background(), describeCreateAccountInput)
		if err != nil {
			return "", fmt.Errorf("error describing create account status: %w", err)
		}


		state := describeCreateAccountResult.CreateAccountStatus.State
		
		if state == types.CreateAccountStateSucceeded {
			// Account creation process is complete
			return *describeCreateAccountResult.CreateAccountStatus.AccountId, nil
		} else if state == types.CreateAccountStateFailed {
			// Account creation process has failed
			return "", errors.New(fmt.Sprintf("Account creation failed with error: %s\n", fmt.Sprintf("%v", describeCreateAccountResult.CreateAccountStatus.FailureReason)))
			
		} else {
			// Account creation process is still in progress, wait for some time before checking again
			time.Sleep(30 * time.Second)
		}
	}
}
