package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func getAllInstances(svc *ec2.EC2) {
	deleteFile("json-outputs/all-ec2-instances.json")
	err := svc.DescribeInstancesPages(nil,
		func(page *ec2.DescribeInstancesOutput, lastPage bool) bool {
			writeInstancesJsonToFile(page)
			return !lastPage
		})
	if err != nil {
		log.Fatal(err)
	}
}

func writeInstancesJsonToFile(result *ec2.DescribeInstancesOutput) {

	f, err := os.OpenFile("json-outputs/all-ec2-instances.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	enc.Encode(result.Reservations)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
