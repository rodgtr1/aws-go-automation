package main

import (
	// systems "aws-go-automations/ssm"
	// users "aws-go-automations/iam"
	// servers "aws-go-automations/ec2"
	"fmt"

	ws "aws-go-automations/workspaces"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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
	// svcIAM := iam.New(sess)
	// users.GetAllUsers(svcIAM)

	// ------ Workspaces
	svcWS := workspaces.New(sess)
	ws.GetAllWorkspaces(svcWS)
	// ws.GetWorkspaceById(svcWS, []string{"ws-00000001", "ws-00000002", "ws-00000003"})

	// ------ SSM
	// svcSSM := ssm.New(sess)
	// systems.GetAllManagedInstances(svcSSM)
	// systems.GetManagedInstancesById(svcSSM, []string{"mi-00000000000001", "mi-0000000000002", "i-00000000000003"})

	// ------EC2
	// svcEC2 := ec2.New(sess)
	// servers.GetAllInstances(svcEC2)
}
