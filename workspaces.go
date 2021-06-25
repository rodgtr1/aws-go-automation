package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

func getAllWorkspaces(svc *workspaces.WorkSpaces) {
	deleteFile("json-outputs/all-workspaces.json")

	err := svc.DescribeWorkspacesPages(nil,
		func(page *workspaces.DescribeWorkspacesOutput, lastPage bool) bool {
			writeWorkspacesJsonToFile(page)
			return !lastPage
		})
	if err != nil {
		log.Fatal(err)
	}
}

func getWorkspaceById(svc *workspaces.WorkSpaces, workspaceIds []string) {

	params := &workspaces.DescribeWorkspacesInput{
		WorkspaceIds: aws.StringSlice(workspaceIds),
	}

	ws, err := svc.DescribeWorkspaces(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ws)
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
