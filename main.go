package main

import (
	"fmt"

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
	// getAllUsers(svcIAM)

	// ------ Workspaces
	svcWS := workspaces.New(sess)
	getAllWorkspaces(svcWS)
	// getWorkspaceById(svcWS, []string{"workspaceID1", "workspaceID2", "workspaceID3"})

	// ------ SSM
	// svcSSM := ssm.New(sess)
	// getAllManagedInstances(svcSSM)

	// ------EC2
	// svcEC2 := ec2.New(sess)
	// getAllInstances(svcEC2)
}
