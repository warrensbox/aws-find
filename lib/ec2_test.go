package lib_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	lib "github.com/warrensbox/aws-find/lib"
)

type Tag struct {
	_ struct{} `type:"structure"`

	// The type of tag on which to filter.
	//
	// Key is a required field
	Key *string `locationName:"key" type:"string" required:"true"`

	// A value for a tag key on which to filter.
	//
	// Value is a required field
	Value *string `locationName:"value" type:"string" required:"true"`
}

type CreateTagsInput struct {
	_         struct{}  `type:"structure"`
	DryRun    *bool     `locationName:"dryRun" type:"boolean"`
	Resources []*string `locationName:"ResourceId" type:"list" required:"true"`
	Tags      []*Tag    `locationName:"Tag" locationNameList:"item" type:"list" required:"true"`
}

// TestPrintInstance : print tag values
func TestPrintInstance(t *testing.T) {

	Input := &ec2.CreateTagsInput{
		Resources: []*string{
			aws.String("ami-78a54011"),
		},
		Tags: []*ec2.Tag{
			{
				Key:   aws.String("Stack"),
				Value: aws.String("production"),
			},
		},
	}

	tag := lib.FindTags(Input.Tags, "Stack")

	if tag == "production" {
		t.Logf("Tag found [expected]")
	} else {
		t.Error("No tag found [unexpected]")
	}
}
