package lib

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func (id *Constructor) FindALB(ch chan string) {

	tagName := id.TagName
	tagValue := id.TagValue

	tag := fmt.Sprintf("tag:%s", "environment") //default tagname

	if tagName != "" {
		tag = fmt.Sprintf("tag:%s", tagName)
	}

	fmt.Println(id.TagName)
	fmt.Println(id.Session)
	fmt.Println("hello")

	svc := ec2.New(id.Session)
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String(tag),
				Values: []*string{
					aws.String(strings.Join([]string{"*", tagValue, "*"}, "")),
				},
			},
		},
	}

	resp, _ := svc.DescribeInstances(params)

	fmt.Println(resp)
}

//testing git hub user
