package main

/*
* Version 0.1.0
* Compatible with Mac OS X and linux OS ONLY
 */

/*** OPERATION WORKFLOW ***/
/*
* 1- SSM library gets AWS credentials from host machine
* 2- Makes API call to aws
* 3- Prints desired in table as output
 */

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/warrensbox/aws-find/lib"
	"github.com/warrensbox/aws-find/modal"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var version = "0.1.0\n"

var (
	versionFlag *bool
	helpFlag    *bool
	awsRegion   *string
	tagName     *string
	tagValue    *string
	services    *string
)

func init() {

	const (
		cmdDesc         = "Look the IP addresses and instance IDs by their tag name in AWS"
		versionFlagDesc = "Displays the version of aws-find"
		servicesArgDesc = "Provide service(s) to be queried. Ex: ec2,alb,rds,sg. Default: ec2"
		awsRegionDesc   = "Provide AWS Region. Default is us-east-1"
		tagNameDesc     = "Provide AWS Tag name"
		tagValueDesc    = "Provide AWS Tag value"
	)

	versionFlag = kingpin.Flag("version", versionFlagDesc).Short('v').Bool()
	awsRegion = kingpin.Flag("region", awsRegionDesc).Short('r').String()
	services = kingpin.Arg("services", servicesArgDesc).String()
	tagName = kingpin.Flag("tag", tagNameDesc).Short('T').String()
	tagValue = kingpin.Flag("value", tagValueDesc).Short('V').String()

}

func main() {

	kingpin.Parse()

	config := &aws.Config{Region: aws.String(*awsRegion)}

	session := session.Must(session.NewSession(config))

	t := &modal.InstanceProfile{
		TagName:  *tagName,
		TagValue: *tagValue,
	}

	ch := make(chan string)

	if *versionFlag {
		fmt.Printf("\nVersion: %v\n", version)
	} else if *services != "" {
		fmt.Println("Printing services")

	} else {
		go lib.FindEC2(session, t, ch)
		<-ch
	}

}
