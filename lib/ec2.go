package lib

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jedib0t/go-pretty/table"
)

func (id *Constructor) FindEC2(ch chan string) {

	svc := ec2.New(id.Session)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(id.TagName),
				Values: []*string{
					aws.String(strings.Join([]string{"*", id.TagValue, "*"}, "")),
				},
			},
		},
	}

	resp, err := svc.DescribeInstances(params)
	CheckError("Can't find instance in region", err)

	var instances [][]interface{}
	reservations := resp.Reservations

	//fmt.Println(reservations)
	if len(reservations) == 0 {
		fmt.Println("No instances found with the given tag name and value. Try again...")
		ch <- "No instances found..."
		return
	}

	var numberOfInst int
	for _, instance := range reservations {

		for _, w := range instance.Instances {

			numberOfInst++
			var arr []interface{}
			arr = append(arr, fmt.Sprintf("%d", numberOfInst)) //Number of instances
			arr = append(arr, FindTags(w.Tags, "Name"))        //Name
			arr = append(arr, *w.InstanceId)                   //InstanceID
			arr = append(arr, *w.State.Name)                   //State
			arr = append(arr, *w.PrivateIpAddress)             //Private IP
			arr = append(arr, *w.InstanceType)                 //instance type
			arr = append(arr, *w.Placement.AvailabilityZone)   //AZ zone
			profile := w.IamInstanceProfile                    //IAMRole
			if profile != nil {
				temp := *profile.Arn
				arr = append(arr, temp[strings.Index(temp, "/")+1:])
			} else {
				arr = append(arr, "No Profile")
			}
			instances = append(instances, arr)
		}
	}

	ch <- printInstance(instances)
}

func printInstance(instances [][]interface{}) string {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if len(instances) == 0 {
		t.AppendHeader(table.Row{"Instances"})
		t.AppendRow([]interface{}{"Nothing found"})
		return "No instances found..."
	} else {
		t.AppendHeader(table.Row{"#", "Name", "Instance ID", "State", "Private IP", "Type", "AZ Zone", "Role"})

		for _, profile := range instances {
			t.AppendRow(profile)
		}

		// for _, profile := range instances {
		// 	t.AppendRow([]interface{}{profile[0], profile[1], profile[2], profile[3], profile[4], profile[5], profile[6], profile[7]})
		// }
		t.SetStyle(table.StyleLight)
		t.Render()
		return "Found instances"
	}
}

func FindTags(tags []*ec2.Tag, tagValue string) string {
	for i := range tags {
		if strings.EqualFold(*tags[i].Key, tagValue) {
			return *tags[i].Value
		}
	}
	return "No Tag"
}
