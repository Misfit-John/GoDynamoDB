package main

import "testing"
import "fmt"
import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/aws"

func TestDynamoGetTableAPI(t *testing.T) {
	config := aws.NewConfig().WithRegion("us-east-1").WithEndpoint("http://127.0.0.1:8000")

	db := dynamodb.New(session.New(), config)

	params := &dynamodb.ScanInput{
		TableName: aws.String("Test"), // Required
		AttributesToGet: []*string{
			aws.String("id"), // Required
		},
		ConditionalOperator:    aws.String("ConditionalOperator"),
		ConsistentRead:         aws.Bool(true),
		IndexName:              aws.String("IndexName"),
		Limit:                  aws.Int64(1),
		ProjectionExpression:   aws.String("ProjectionExpression"),
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
		Segment:                aws.Int64(1),
		Select:                 aws.String("Select"),
		TotalSegments:          aws.Int64(1),
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
