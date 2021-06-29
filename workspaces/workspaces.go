package workspaces

import (
	"aws-go-automations/utils"
	"encoding/csv"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

var writeHeaders bool = true

func GetAllWorkspaces(svc *workspaces.WorkSpaces) {

	utils.DeleteFile("outputs/workspaces/all-workspaces.csv")

	headers := []string{"workspaceid", "username", "state", "computername"}

	err := svc.DescribeWorkspacesPages(nil,
		func(page *workspaces.DescribeWorkspacesOutput, lastPage bool) bool {
			writeWorkspacesToCsv(page, headers, writeHeaders, "all-workspaces")
			writeHeaders = false
			return !lastPage
		})
	if err != nil {
		log.Fatal(err)
	}
}

func GetWorkspaceById(svc *workspaces.WorkSpaces, workspaceIds []string) {

	utils.DeleteFile("outputs/workspaces/workspaces.csv")

	headers := []string{"workspaceid", "username", "state", "computername"}

	params := &workspaces.DescribeWorkspacesInput{
		WorkspaceIds: aws.StringSlice(workspaceIds),
	}

	ws, err := svc.DescribeWorkspaces(params)
	if err != nil {
		log.Fatal(err)
	}
	writeWorkspacesToCsv(ws, headers, writeHeaders, "workspaces")
}

func writeWorkspacesToCsv(result *workspaces.DescribeWorkspacesOutput, headers []string, writeHeaders bool, csvName string) {

	wsData := [][]string{}
	csvFile, err := os.OpenFile("outputs/workspaces/"+csvName+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	if writeHeaders {
		err = csvwriter.Write(headers)
	}

	if err != nil {
		log.Fatalf("Failed adding headers: %s", err)
	}

	for _, workspace := range result.Workspaces {
		wsData = append(wsData, []string{*workspace.WorkspaceId, *workspace.UserName, *workspace.State, *workspace.ComputerName})
	}

	for _, wrRow := range wsData {
		_ = csvwriter.Write(wrRow)
	}

	csvwriter.Flush()
	csvFile.Close()

}
