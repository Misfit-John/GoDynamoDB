package main

import "testing"
import "fmt"
import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/aws"

func TestDynamoGetTableAPI(t *testing.T) {

	db := dynamodb.New(session.New(
		&aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String("http://127.0.0.1:8000")}))

	params := &dynamodb.ScanInput{
		TableName: aws.String("Test"), // Required
		AttributesToGet: []*string{
			aws.String("id"), // Required
		},
		Limit: aws.Int64(1),
	}
	resp, err := db.Scan(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
