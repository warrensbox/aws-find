package lib

import (
	"fmt"
	"log"

	session "github.com/aws/aws-sdk-go/aws/session"
)

//Constructor : struct for session
type Constructor struct {
	TagName  string
	TagValue string
	Session  *session.Session
}

//NewConstructor :validate object
func NewConstructor(attr *Constructor) (*Constructor, error) {

	if attr.Session == nil {
		log.Panic("Unable to obtain AWS credentials")
	} else {

	}
	if attr.TagName == "" {
		log.Println("Empty Name")
		attr.TagName = fmt.Sprintf("tag:%s", "environment")
	} else {
		log.Println("Empty Name")
		attr.TagName = fmt.Sprintf("tag:%s", attr.TagName)
	}
	if attr.TagValue == "" {
		log.Println("Empty Value")
	}

	return attr, nil
}
