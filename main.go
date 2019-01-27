package main

import (
	"fmt"
	"os"

	aws "github.com/aws/aws-sdk-go/aws"
	session "github.com/aws/aws-sdk-go/aws/session"
	getopt "github.com/pborman/getopt"
	lib "github.com/warrensbox/aws-find/lib"
)

func main() {

	helpFlag := getopt.BoolLong("help", 'h', "displays help message")
	awsRegion := getopt.StringLong("region", 'r', "us-east-1", "AWS Region")
	awsEnv := getopt.StringLong("env", 'e', "", "AWS Environment")

	getopt.Parse()

	config := &aws.Config{Region: aws.String(*awsRegion)}

	awsSession, err := session.NewSession(config)
	lib.CheckError("Can't create aws session", err)

	ch := make(chan string)

	if *helpFlag {
		usageMessage()
	} else if *awsEnv != "" {
		go lib.FindEC2(awsSession, *awsEnv, ch)
		<-ch
	}

}

func usageMessage() {
	fmt.Print("\n\n")
	getopt.PrintUsage(os.Stderr)
	fmt.Println("Pass path to csv file to -f or --file; example ./main -f test.csv")
}
