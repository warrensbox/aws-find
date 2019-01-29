package main_test

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
)

func init() {
	os.Unsetenv("AWS_REGION")
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")

}

// TestMain : remove set
func TestMain(t *testing.T) {

	os.Setenv("AWS_ACCESS_KEY_ID", "XXX")

	_, err := session.NewSession()

	if err == nil {
		t.Logf("Credentials found [expected]")
	} else {
		t.Logf("NO Credentials found [expected]")
	}

}
