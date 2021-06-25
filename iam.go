package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
)

func getAllUsers(svc *iam.IAM) {
	deleteFile("iam/iam-all-users.json")

	input := &iam.ListUsersInput{}

	result, err := svc.ListUsers(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case iam.ErrCodeServiceFailureException:
				fmt.Println(iam.ErrCodeServiceFailureException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	writeUsersJsonToFile(result)
}

func writeUsersJsonToFile(result *iam.ListUsersOutput) {

	f, err := os.OpenFile("json-outputs/iam-all-users.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	enc.Encode(result)

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
