package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/ssm"
)

func getAllManagedInstances(svc *ssm.SSM) {
	deleteFile("json-outputs/all-managed-instances.json")
	pageNum := 0
	err := svc.DescribeInstanceInformationPages(nil,
		func(page *ssm.DescribeInstanceInformationOutput, lastPage bool) bool {
			pageNum++
			writeSSMInstancesJsonToFile(page)
			return pageNum <= 50
		})
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func writeSSMInstancesJsonToFile(result *ssm.DescribeInstanceInformationOutput) {

	f, err := os.OpenFile("json-outputs/all-managed-instances.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	enc.Encode(result.InstanceInformationList)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
