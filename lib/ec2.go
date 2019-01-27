package lib

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jedib0t/go-pretty/table"
)

func FindEC2(sess *session.Session, envName string, ec2chan chan string) {

	//fmt.Println(sess)

	svc := ec2.New(sess)
	//svc := ec2.New(sess, "us-east-1")

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:environment"),
				Values: []*string{
					aws.String(strings.Join([]string{"*", envName, "*"}, "")),
				},
			},
		},
	}

	resp, err := svc.DescribeInstances(params)

	CheckError("Can't find instance in region", err)

	var instances [][]string
	reservations := resp.Reservations

	//fmt.Println(reservations)
	var count int
	for _, instance := range reservations {

		for _, w := range instance.Instances {

			count++
			var arr []string
			arr = append(arr, fmt.Sprintf("%d", count))
			arr = append(arr, findTags(w.Tags, "Name"))      //Name
			arr = append(arr, *w.State.Name)                 //State
			arr = append(arr, *w.PrivateIpAddress)           //Private IP
			arr = append(arr, *w.InstanceType)               //instance type
			arr = append(arr, *w.Placement.AvailabilityZone) //AZ zone
			profile := w.IamInstanceProfile                  //IAMRole
			if profile != nil {
				temp := *profile.Arn
				arr = append(arr, temp[strings.Index(temp, "/")+1:])
			} else {
				arr = append(arr, "NONE")
			}

			instances = append(instances, arr)
		}

	}

	ec2chan <- printInstance(instances)
}

func findTags(tags []*ec2.Tag, tagValue string) string {
	for i := range tags {
		if strings.EqualFold(*tags[i].Key, tagValue) {
			return *tags[i].Value
		}
	}
	return "NONE"
}

func printInstance(instances [][]string) string {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "State", "Private IP", "Type", "AZ Zone", "Role"})
	// t.AppendRows([]table.Row{
	// 	{1, "Arya", "Stark", 3000},
	// 	{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	// })

	for _, profile := range instances {
		//fmt.Println(test)
		t.AppendRow([]interface{}{profile[0], profile[1], profile[2], profile[3], profile[4], profile[5], profile[6]})
	}
	// fmt.Println([]interface{}{300, "Tyrion", "Lannister", 5000})
	// t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	//t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.SetStyle(table.StyleLight)
	t.Render()
	return "GOOD"
}
