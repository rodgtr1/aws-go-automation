package workspaces

import (
	"aws-go-automations/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

func GetAllWorkspaces(svc *workspaces.WorkSpaces) {
	utils.DeleteFile("json-outputs/all-workspaces.json")

	err := svc.DescribeWorkspacesPages(nil,
		func(page *workspaces.DescribeWorkspacesOutput, lastPage bool) bool {
			for _, w := range page.Workspaces {
				fmt.Printf("Username: %s\nState: %s\nWorkspaceId: %s\nComputerName: %s\n-----------------\n", *w.UserName, *w.State, *w.WorkspaceId, *w.ComputerName)
			}
			//writeWorkspacesJsonToFile(page)
			return !lastPage
		})
	if err != nil {
		log.Fatal(err)
	}
}

func GetWorkspaceById(svc *workspaces.WorkSpaces, workspaceIds []string) {

	params := &workspaces.DescribeWorkspacesInput{
		WorkspaceIds: aws.StringSlice(workspaceIds),
	}

	ws, err := svc.DescribeWorkspaces(params)
	if err != nil {
		log.Fatal(err)
	}
	for _, w := range ws.Workspaces {
		fmt.Printf("Username: %s\nState: %s\nWorkspaceId: %s\nComputerName: %s\n-----------------\n", *w.UserName, *w.State, *w.WorkspaceId, *w.ComputerName)
	}
}

func writeWorkspacesJsonToFile(result *workspaces.DescribeWorkspacesOutput) {

	f, err := os.OpenFile("json-outputs/all-workspaces.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	enc.Encode(result.Workspaces)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
