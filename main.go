package main

import (
	servers "aws-go-automations/ec2"
	users "aws-go-automations/iam"
	systems "aws-go-automations/ssm"
	ws "aws-go-automations/workspaces"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

func main() {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(goDotEnvVariable("REGION"))},
		Profile: goDotEnvVariable("PROFILE"),
	})

	if err != nil {
		fmt.Printf("Unable to establish a session: %v", err)
	}

	// ------ IAM
	svcIAM := iam.New(sess)
	users.GetAllUsers(svcIAM)

	// ------ Workspaces
	svcWS := workspaces.New(sess)
	ws.GetAllWorkspaces(svcWS)
	ws.GetWorkspaceById(svcWS, []string{"workspaceID1", "workspaceID2", "workspaceID3"})

	// ------ SSM
	svcSSM := ssm.New(sess)
	systems.GetAllManagedInstances(svcSSM)

	// ------EC2
	svcEC2 := ec2.New(sess)
	servers.GetAllInstances(svcEC2)
}
