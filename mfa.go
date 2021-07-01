package main

import (
	"aws-go-automations/utils"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	sts "github.com/aws/aws-sdk-go/service/sts"
)

var mfaDeviceArn string = goDotEnvVariable("MFA_DEVICE_ARN")
var pathToCredentialsFile string = goDotEnvVariable("PATH_TO_CREDENTIALS_FILE")

// 1H = 3600, 2H = 7200, 3H = 10800, 4H = 14400, 5H = 18000, 6H = 21600, 7H = 25200
var tokenDurationInSeconds int64 = 21600

func authenticateWithMFA(svc *sts.STS) {

	utils.ReplaceWordInFile(".env", "mfa", "default")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your mfa token: ")
	scanner.Scan()

	params := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(tokenDurationInSeconds),
		SerialNumber:    aws.String(mfaDeviceArn),
		TokenCode:       aws.String(scanner.Text()),
	}

	result, err := svc.GetSessionToken(params)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeRegionDisabledException:
				fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	key := result.Credentials.AccessKeyId
	secret := result.Credentials.SecretAccessKey
	sessToken := result.Credentials.SessionToken

	removeOldMFACreds(pathToCredentialsFile)
	writeCredentialsToFile(pathToCredentialsFile, *key, *secret, *sessToken)
}

func removeOldMFACreds(file string) {
	// open original file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// create temp file
	tmp, err := ioutil.TempFile("", "replace-*")
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.Close()

	// replace while copying from f to tmp
	if err := replace(f, tmp); err != nil {
		log.Fatal(err)
	}

	// make sure the tmp file was successfully written to
	if err := tmp.Close(); err != nil {
		log.Fatal(err)
	}

	// close the file we're reading from
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// overwrite the original file with the temp file
	if err := os.Rename(tmp.Name(), file); err != nil {
		log.Fatal(err)
	}
}

func replace(r io.Reader, w io.Writer) error {

	eraseFlag := 0

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()

		if line == "[mfa]" {
			eraseFlag = 1
		}

		if eraseFlag == 1 {
			continue
		}

		if _, err := io.WriteString(w, line+"\n"); err != nil {
			return err
		}
	}

	return sc.Err()
}

func writeCredentialsToFile(file string, key string, secret string, token string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Could not open credentials file")
	}

	_, err = f.WriteString("[mfa]\naws_access_key_id: " + key + "\naws_secret_access_key: " + secret + "\naws_session_token: " + token + "\n")
	if err != nil {
		fmt.Println("Credentials file could not be written to. Error: ", err)
		os.Exit(1)
	}

	utils.ReplaceWordInFile(".env", "default", "mfa")

	f.Close()
}
