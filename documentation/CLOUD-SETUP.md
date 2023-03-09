# Vending Machine POC
### Cloud Setup

Simple overview of Cloud Setup pre-requisite for Vending Machine

## Description

Creation of the Management Account with the Cloud Provider is a manual Process.
The user Sign-up with AWS and OCI 

## Starting Position
Creation of the Management Account with the Cloud Provider is a manual Process.
The user Sign-up with AWS and OCI with email address and Credit Card Details

* Root Organization and Management Account is created (AWS)
* Primary Tenancy is created ( OCI )

## Follow-up Steps

* To Safe guard the Management Account. For one time setup, User Sign-in with Management Account
* Create IAM User For Example: "management-iam" with IAMFullAcess Policy, AWSOrganizationReadOnlyAccess AWS Managed Policy
* Create additional two custom Policy "custom-CreateAccount" and "GrantAccessToOrganizationAccountAccessRole"
* For the "management-iam" create Security Credentials AWS Access Key and Secret. IAM >> Users >> "management-iam" >> Security Credentials
* Store the Access Key/Secret in any convinient and secure form to be used programatically

### Policy Json for "custom-CreateAccount"
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "CreateOrganizationAccounts",
            "Effect": "Allow",
            "Action": [
                "organizations:CreateAccount"
            ],
            "Resource": "*"
        }
    ]
}
```



### Policy Json for "GrantAccessToOrganizationAccountAccessRole"

This is a dyanamic policy. For every Child Account Created, additional entry of Resource need to be done programatically

```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": "sts:AssumeRole",
            "Resource": "arn:aws:iam::<<CHILD-ACCOUNT>:role/OrganizationAccountAccessRole"
        }
    ]
}
```

https://github.com/dbader/readme-template 