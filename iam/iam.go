package users

import (
	"aws-go-automations/utils"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/iam"
)

func GetAllUsers(svc *iam.IAM) {
	utils.DeleteFile("iam/iam-all-users.json")

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

	writeUsersToCsv(result)
	//WriteUsersJsonToFile(result)
}

func writeUsersToCsv(result *iam.ListUsersOutput) {

	wsData := [][]string{}

	csvFile, err := os.Create("outputs/iam/iam-all-users.csv")
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	err = csvwriter.Write([]string{"userid", "username", "created"})
	if err != nil {
		log.Fatalf("Failed adding headers: %s", err)
	}

	for _, user := range result.Users {
		wsData = append(wsData, []string{*user.UserId, *user.UserName, user.CreateDate.String()})
		// fmt.Printf("Id: %s\nUser: %s\nCreated: %s\n-----------------------\n", *user.UserId, *user.UserName, user.CreateDate)
	}

	for _, wrRow := range wsData {
		fmt.Println(wrRow)
		_ = csvwriter.Write(wrRow)
	}

	csvwriter.Flush()
	csvFile.Close()

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
