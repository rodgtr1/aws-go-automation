package main

import (
	// systems "aws-go-automations/ssm"
	// users "aws-go-automations/iam"
	// servers "aws-go-automations/ec2"
	ws "aws-go-automations/workspaces"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

func main() {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(goDotEnvVariable("REGION"))},
		Profile:           goDotEnvVariable("PROFILE"),
		SharedConfigFiles: []string{goDotEnvVariable("PATH_TO_CREDENTIALS_FILE")},
	})

	if err != nil {
		fmt.Printf("Unable to establish a session: %v", err)
	}

	// ------ MFA AUTH
	// svcSTS := sts.New(sess)
	// authenticateWithMFA(svcSTS)

	// ------ IAM
	// svcIAM := iam.New(sess)
	// users.GetAllUsers(svcIAM)

	// ------ Workspaces
	svcWS := workspaces.New(sess)
	ws.GetAllWorkspaces(svcWS)
	// ws.GetWorkspaceById(svcWS, []string{"ws-123456789", "ws-123456788", "123456777"})

	// ------ SSM
	// svcSSM := ssm.New(sess)
	// systems.GetAllManagedInstances(svcSSM)
	// systems.GetManagedInstancesById(svcSSM, []string{"mi-00000000000000001", "mi-00000000000000002"})
	// systems.DeregisterInstance(svcSSM, "mi-00000000000001")

	// ------EC2
	// svcEC2 := ec2.New(sess)
	// servers.GetAllInstances(svcEC2)
}
