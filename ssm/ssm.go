package systems

import (
	"aws-go-automations/utils"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var writeHeaders bool = true

func GetAllManagedInstances(svc *ssm.SSM) {
	utils.DeleteFile("outputs/ssm/all-managed-instances.csv")

	headers := []string{"Instance Id", "Ping Status", "Platform Type", "Name"}

	err := svc.DescribeInstanceInformationPages(nil,
		func(page *ssm.DescribeInstanceInformationOutput, lastPage bool) bool {
			writeSSMInstancesToCsv(page, headers, writeHeaders, "all-managed-instances")
			writeHeaders = false
			return !lastPage
		})
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func GetManagedInstancesById(svc *ssm.SSM, instanceIds []string) {

	utils.DeleteFile("outputs/ssn/managed-instances.csv")

	headers := []string{"workspaceid", "username", "state", "computername"}

	params := &ssm.DescribeInstanceInformationInput{
		Filters: []*ssm.InstanceInformationStringFilter{
			{
				Key:    aws.String("InstanceIds"),
				Values: aws.StringSlice(instanceIds),
			},
		},
	}

	ws, err := svc.DescribeInstanceInformation(params)

	if err != nil {
		log.Fatal(err)
	}
	writeSSMInstancesToCsv(ws, headers, writeHeaders, "managed-instances")
}

func writeSSMInstancesToCsv(result *ssm.DescribeInstanceInformationOutput, headers []string, writeHeaders bool, csvName string) {

	wsData := [][]string{}
	csvFile, err := os.OpenFile("outputs/ssm/"+csvName+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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

	for _, instance := range result.InstanceInformationList {
		wsData = append(wsData, []string{*instance.InstanceId, *instance.PingStatus, *instance.PlatformType, *instance.ComputerName})
	}

	for _, wrRow := range wsData {
		_ = csvwriter.Write(wrRow)
	}

	csvwriter.Flush()
	csvFile.Close()

}
