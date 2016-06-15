package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "fmt"

func PutItem(i ModelBase) error {
	db := dynamodb.New(session.New(
		&aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String("http://127.0.0.1:8000")}))

	attMap, err := encode(i)
	if nil != err {
		return err
	}
	params := &dynamodb.PutItemInput{
		Item:         attMap,
		TableName:    aws.String(i.GetTableName()),
		ReturnValues: aws.String("NONE"),
	}
	resp, err := db.PutItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return err
	}

	// Pretty-print the response data.
	fmt.Println(resp)

	return nil
}
