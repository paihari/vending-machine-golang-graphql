package main

import (
	"context"
	
	"fmt"
	
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	

)


func main() {

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