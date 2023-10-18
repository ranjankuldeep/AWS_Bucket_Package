package awsbucket

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// An error message to be displayed to standard output
func exitError(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, msg+"\n", args...)
	os.Exit(1)
}
func ListBuckets(region string) {
	// creating the session
	session, err := session.NewSession(&aws.Config{Region: aws.String(region)})

	svc := s3.New(session)

	//For listing all the buckets out there
	result, err := svc.ListBuckets(nil)
	if err != nil {
		exitError("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
