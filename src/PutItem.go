package GoDynamoDB

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"
import "fmt"

func (db GoDynamoDB) PutItem(i ModelBase) error {
	attMap, err := encode(i)
	if nil != err {
		return err
	}
	params := &dynamodb.PutItemInput{
		Item:         attMap,
		TableName:    aws.String(i.GetTableName()),
		ReturnValues: aws.String("NONE"),
	}
	resp, err := db.db.PutItem(params)

	if err != nil {
		return err
	}

	return nil
}
